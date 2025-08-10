package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Env struct {
	RPCURL  string
	PrivKey string
	ChainID *big.Int
}

func mustEnv() Env {
	rpc := getenv("RPC_URL", "http://127.0.0.1:8545")
	pk := os.Getenv("PRIVATE_KEY")
	cid := getenv("CHAIN_ID", "31337")
	chainID, ok := new(big.Int).SetString(cid, 10)
	if !ok {
		log.Fatalf("invalid CHAIN_ID: %s", cid)
	}
	if pk == "" {
		log.Fatalf("set PRIVATE_KEY (hex, no 0x)")
	}
	return Env{RPCURL: rpc, PrivKey: pk, ChainID: chainID}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func dial(rpc string) *ethclient.Client {
	cli, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	return cli
}

func authFromPK(privHex string, chainID *big.Int) (*bind.TransactOpts, *ecdsa.PrivateKey) {
	if strings.HasPrefix(privHex, "0x") {
		privHex = privHex[2:]
	}
	key, err := crypto.HexToECDSA(privHex)
	if err != nil {
		log.Fatalf("bad PRIVATE_KEY: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("transactor: %v", err)
	}
	auth.Context = context.Background()
	// Leave gas fields nil to use dynamic fees suggested by the node.
	return auth, key
}

func waitMined(cli *ethclient.Client, tx *types.Transaction, what string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	rec, err := bind.WaitMined(ctx, cli, tx)
	if err != nil {
		log.Fatalf("wait mined: %v", err)
	}
	if rec.Status != types.ReceiptStatusSuccessful {
		log.Fatalf("tx failed: %s (status %d)", tx.Hash(), rec.Status)
	}
	fmt.Printf("%s mined in block %d (gas used %d)\n", what, rec.BlockNumber.Uint64(), rec.GasUsed)
}

// ---- Commands ----

func cmdDeployToken(args []string) {
	fs := flag.NewFlagSet("deploy-token", flag.ExitOnError)
	_ = fs.Parse(args)
	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, key := authFromPK(env.PrivKey, env.ChainID)

	addr, tx, token, err := contracts.DeployToken(auth, cli)
	if err != nil {
		log.Fatalf("deploy token: %v", err)
	}
	fmt.Printf("Token tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Token deployment")
	fmt.Printf("Token deployed at: %s\n", addr.Hex())

	bal, err := token.GetBalance(&bind.CallOpts{})
	if err == nil {
		fmt.Printf("Token initial balance (contract): %s\n", bal.String())
	}
	fmt.Printf("Deployer: %s\n", crypto.PubkeyToAddress(key.PublicKey).Hex())
}

func cmdDeployMarketplace(args []string) {
	fs := flag.NewFlagSet("deploy-marketplace", flag.ExitOnError)
	tokenAddr := fs.String("token", "", "deployed Token address")
	_ = fs.Parse(args)
	if *tokenAddr == "" {
		log.Fatalf("--token is required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	tAddr := common.HexToAddress(*tokenAddr)
	addr, tx, _, err := contracts.DeployMarketplace(auth, cli, tAddr)
	if err != nil {
		log.Fatalf("deploy marketplace: %v", err)
	}
	fmt.Printf("Marketplace tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Marketplace deployment")
	fmt.Printf("Marketplace deployed at: %s\n", addr.Hex())
}

func cmdDeployInteractor(args []string) {
	fs := flag.NewFlagSet("deploy-interactor", flag.ExitOnError)
	marketAddr := fs.String("marketplace", "", "deployed Marketplace address")
	_ = fs.Parse(args)
	if *marketAddr == "" {
		log.Fatalf("--marketplace is required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	mAddr := common.HexToAddress(*marketAddr)
	addr, tx, _, err := contracts.DeployInteractor(auth, cli, mAddr)
	if err != nil {
		log.Fatalf("deploy interactor: %v", err)
	}
	fmt.Printf("Interactor tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Interactor deployment")
	fmt.Printf("Interactor deployed at: %s\n", addr.Hex())
}

func cmdApprove(args []string) {
	fs := flag.NewFlagSet("approve", flag.ExitOnError)
	interactor := fs.String("interactor", "", "Interactor address (preferred)")
	marketplace := fs.String("marketplace", "", "Marketplace address (direct)")
	amount := fs.String("amount", "1000000000000000000", "amount in wei-like units of your Token (uint)") // adapt to your token semantics
	_ = fs.Parse(args)
	if *interactor == "" && *marketplace == "" {
		log.Fatalf("--interactor or --marketplace is required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	amt, ok := new(big.Int).SetString(*amount, 10)
	if !ok {
		log.Fatalf("bad --amount")
	}

	if *interactor != "" {
		ia := common.HexToAddress(*interactor)
		inter, err := contracts.NewInteractor(ia, cli)
		if err != nil {
			log.Fatalf("bind interactor: %v", err)
		}
		tx, err := inter.ApproveMarketplace(auth, amt) // via Interactor
		if err != nil {
			log.Fatalf("approve via interactor: %v", err)
		}
		fmt.Printf("Approve tx: %s\n", tx.Hash())
		waitMined(cli, tx, "Approve marketplace")
		return
	}
	// Fallback: direct call to Marketplace.approveMarketplace
	ma := common.HexToAddress(*marketplace)
	mkt, err := contracts.NewMarketplace(ma, cli)
	if err != nil {
		log.Fatalf("bind marketplace: %v", err)
	}
	tx, err := mkt.ApproveMarketplace(auth, amt)
	if err != nil {
		log.Fatalf("approve via marketplace: %v", err)
	}
	fmt.Printf("Approve tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Approve marketplace")
}

func cmdList(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	interactor := fs.String("interactor", "", "Interactor address")
	price := fs.String("price", "100000000000000000", "listing price (uint)")
	_ = fs.Parse(args)
	if *interactor == "" {
		log.Fatalf("--interactor is required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	ia := common.HexToAddress(*interactor)
	inter, err := contracts.NewInteractor(ia, cli)
	if err != nil {
		log.Fatalf("bind interactor: %v", err)
	}
	p, ok := new(big.Int).SetString(*price, 10)
	if !ok {
		log.Fatalf("bad --price")
	}
	tx, err := inter.ListItem(auth, p)
	if err != nil {
		log.Fatalf("listItem: %v", err)
	}
	fmt.Printf("ListItem tx: %s\n", tx.Hash())
	waitMined(cli, tx, "List item")
}

func cmdBuy(args []string) {
	fs := flag.NewFlagSet("buy", flag.ExitOnError)
	interactor := fs.String("interactor", "", "Interactor address")
	id := fs.Uint64("id", 0, "listing id")
	value := fs.String("value", "0", "native ETH to send (if required by your contract)")
	_ = fs.Parse(args)
	if *interactor == "" {
		log.Fatalf("--interactor is required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	ia := common.HexToAddress(*interactor)
	inter, err := contracts.NewInteractor(ia, cli)
	if err != nil {
		log.Fatalf("bind interactor: %v", err)
	}
	// Set value if Marketplace.buy is payable in your ABI (it is), but Interactor may forward ETH too.
	val, ok := new(big.Int).SetString(*value, 10)
	if !ok {
		log.Fatalf("bad --value")
	}
	auth.Value = val
	tx, err := inter.BuyItem(auth, new(big.Int).SetUint64(*id))
	if err != nil {
		log.Fatalf("buyItem: %v", err)
	}
	fmt.Printf("BuyItem tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Buy item")
}

func cmdProfile(args []string) {
	fs := flag.NewFlagSet("profile", flag.ExitOnError)
	interactor := fs.String("interactor", "", "Interactor address")
	user := fs.String("user", "", "user address (defaults to signer)")
	_ = fs.Parse(args)
	if *interactor == "" {
		log.Fatalf("--interactor is required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)
	_, key := authFromPK(env.PrivKey, env.ChainID)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	if *user != "" {
		addr = common.HexToAddress(*user)
	}

	ia := common.HexToAddress(*interactor)
	inter, err := contracts.NewInteractor(ia, cli)
	if err != nil {
		log.Fatalf("bind interactor: %v", err)
	}
	prof, err := inter.GetProfile(&bind.CallOpts{}, addr)
	if err != nil {
		log.Fatalf("getProfile: %v", err)
	}
	fmt.Printf("Profile(%s): balance=%s, to=%s\n", addr.Hex(), prof.Balance.String(), prof.To.Hex())
}

func cmdListing(args []string) {
	fs := flag.NewFlagSet("listing", flag.ExitOnError)
	interactor := fs.String("interactor", "", "Interactor address")
	id := fs.Uint64("id", 0, "listing id")
	_ = fs.Parse(args)
	if *interactor == "" {
		log.Fatalf("--interactor is required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)

	ia := common.HexToAddress(*interactor)
	inter, err := contracts.NewInteractor(ia, cli)
	if err != nil {
		log.Fatalf("bind interactor: %v", err)
	}
	li, err := inter.GetListing(&bind.CallOpts{}, new(big.Int).SetUint64(*id))
	if err != nil {
		log.Fatalf("getListing: %v", err)
	}
	fmt.Printf("Listing %d => seller=%s price=%s active=%v\n", *id, li.Seller.Hex(), li.Price.String(), li.Active)
}

func cmdWhoami(args []string) {
	env := mustEnv()
	_, key := authFromPK(env.PrivKey, env.ChainID)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println(addr.Hex())
}

func cmdSetBalance(args []string) {
	fs := flag.NewFlagSet("set-balance", flag.ExitOnError)
	tokenAddr := fs.String("token", "", "Token contract address")
	user := fs.String("user", "", "user address")
	balance := fs.String("balance", "", "new balance (uint)")
	_ = fs.Parse(args)
	if *tokenAddr == "" || *user == "" || *balance == "" {
		log.Fatalf("--token, --user and --balance are required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	t := common.HexToAddress(*tokenAddr)
	token, err := contracts.NewToken(t, cli)
	if err != nil {
		log.Fatalf("bind token: %v", err)
	}

	nb, ok := new(big.Int).SetString(*balance, 10)
	if !ok {
		log.Fatalf("bad --balance")
	}

	ua := common.HexToAddress(*user)
	tx, err := token.UpdateProfileBalance(auth, ua, nb)
	if err != nil {
		log.Fatalf("updateProfileBalance: %v", err)
	}
	fmt.Printf("UpdateProfileBalance tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Update profile balance")
}

func cmdSetProfile(args []string) {
	fs := flag.NewFlagSet("set-profile", flag.ExitOnError)
	tokenAddr := fs.String("token", "", "Token contract address")
	user := fs.String("user", "", "user address")
	balance := fs.String("balance", "", "balance (uint)")
	to := fs.String("to", "", "to address")
	_ = fs.Parse(args)
	if *tokenAddr == "" || *user == "" || *balance == "" || *to == "" {
		log.Fatalf("--token, --user, --balance and --to are required")
	}
	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	t := common.HexToAddress(*tokenAddr)
	token, err := contracts.NewToken(t, cli)
	if err != nil {
		log.Fatalf("bind token: %v", err)
	}

	nb, ok := new(big.Int).SetString(*balance, 10)
	if !ok {
		log.Fatalf("bad --balance")
	}

	ua := common.HexToAddress(*user)
	toa := common.HexToAddress(*to)
	tx, err := token.CreateOrUpdateProfile(auth, ua, nb, toa)
	if err != nil {
		log.Fatalf("createOrUpdateProfile: %v", err)
	}
	fmt.Printf("CreateOrUpdateProfile tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Create/Update profile")
}

func cmdGetUser(args []string) {
	fs := flag.NewFlagSet("get-user", flag.ExitOnError)
	interactor := fs.String("interactor", "", "Interactor address (preferred)")
	tokenAddr := fs.String("token", "", "Token contract address (optional)")
	user := fs.String("user", "", "user address (defaults to signer)")
	_ = fs.Parse(args)

	env := mustEnv()
	cli := dial(env.RPCURL)
	_, key := authFromPK(env.PrivKey, env.ChainID)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	if *user != "" {
		addr = common.HexToAddress(*user)
	}

	// Prefer Interactor if provided
	if *interactor != "" {
		ia := common.HexToAddress(*interactor)
		inter, err := contracts.NewInteractor(ia, cli)
		if err != nil {
			log.Fatalf("bind interactor: %v", err)
		}
		prof, err := inter.GetProfile(&bind.CallOpts{}, addr)
		if err != nil {
			log.Fatalf("getProfile via interactor: %v", err)
		}
		fmt.Printf("User %s\nbalance=%s\nto=%s\n", addr.Hex(), prof.Balance.String(), prof.To.Hex())
		return
	}

	// Else, try Token if provided
	if *tokenAddr != "" {
		ta := common.HexToAddress(*tokenAddr)
		token, err := contracts.NewToken(ta, cli)
		if err != nil {
			log.Fatalf("bind token: %v", err)
		}
		prof, err := token.GetProfile(&bind.CallOpts{}, addr)
		if err != nil {
			log.Fatalf("getProfile via token: %v", err)
		}
		fmt.Printf("User %s\nbalance=%s\nto=%s\n", addr.Hex(), prof.Balance.String(), prof.To.Hex())
		return
	}

	// Fallback: just print the signer address
	fmt.Println(addr.Hex())
}

func cmdGetToken(args []string) {
	fs := flag.NewFlagSet("get-token", flag.ExitOnError)
	interactor := fs.String("interactor", "", "Interactor address (required)")
	_ = fs.Parse(args)
	if *interactor == "" {
		log.Fatalf("--interactor is required")
	}

	env := mustEnv()
	cli := dial(env.RPCURL)

	ia := common.HexToAddress(*interactor)
	inter, err := contracts.NewInteractor(ia, cli)
	if err != nil {
		log.Fatalf("bind interactor: %v", err)
	}
	addr, err := inter.TokenContract(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("tokenContract(): %v", err)
	}
	fmt.Println(addr.Hex())
}

func cmdNextListingID(args []string) {
	fs := flag.NewFlagSet("next-listing-id", flag.ExitOnError)
	marketplace := fs.String("marketplace", "", "Marketplace address (required)")
	_ = fs.Parse(args)
	if *marketplace == "" {
		log.Fatalf("--marketplace is required")
	}

	env := mustEnv()
	cli := dial(env.RPCURL)

	ma := common.HexToAddress(*marketplace)
	mkt, err := contracts.NewMarketplace(ma, cli)
	if err != nil {
		log.Fatalf("bind marketplace: %v", err)
	}
	id, err := mkt.NextListingId(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("nextListingId(): %v", err)
	}
	fmt.Println(id.String())
}
func usage() {
	fmt.Println(`market-dapp — minimal CLI to deploy & use a Token/Marketplace/Interactor dApp

Environment:
  RPC_URL     — JSON-RPC endpoint (default http://127.0.0.1:8545)
  PRIVATE_KEY — hex private key (no 0x)
  CHAIN_ID    — chain id (default 31337)

Commands:
  get-token         --interactor <ADDR>
  next-listing-id   --marketplace <ADDR>
  get-user          [--interactor <ADDR> | --token <ADDR>] [--user <ADDR>]
  set-balance       --token <ADDR> --user <ADDR> --balance <N>
  set-profile       --token <ADDR> --user <ADDR> --balance <N> --to <ADDR>
  whoami
  deploy-token
  deploy-marketplace --token <TOKEN_ADDR>
  deploy-interactor  --marketplace <MARKETPLACE_ADDR>
  approve           --interactor <ADDR> --amount <N>   # or: --marketplace <ADDR>
  list              --interactor <ADDR> --price <N>
  buy               --interactor <ADDR> --id <N> [--value <ETH_WEI>]
  profile           --interactor <ADDR> [--user <ADDR>]
  listing           --interactor <ADDR> --id <N>
`)
}
func cmdListDirect(args []string) {
	fs := flag.NewFlagSet("list-direct", flag.ExitOnError)
	marketplace := fs.String("marketplace", "", "Marketplace address (required)")
	price := fs.String("price", "", "listing price (uint)")
	_ = fs.Parse(args)
	if *marketplace == "" || *price == "" {
		log.Fatalf("--marketplace and --price are required")
	}

	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	ma := common.HexToAddress(*marketplace)
	mkt, err := contracts.NewMarketplace(ma, cli)
	if err != nil {
		log.Fatalf("bind marketplace: %v", err)
	}

	p, ok := new(big.Int).SetString(*price, 10)
	if !ok {
		log.Fatalf("bad --price")
	}

	tx, err := mkt.ListItem(auth, p)
	if err != nil {
		log.Fatalf("marketplace.listItem: %v", err)
	}
	fmt.Printf("ListItem (direct) tx: %s\n", tx.Hash())
	waitMined(cli, tx, "List item (direct)")
}

func cmdBuyDirect(args []string) {
	fs := flag.NewFlagSet("buy-direct", flag.ExitOnError)
	marketplace := fs.String("marketplace", "", "Marketplace address (required)")
	id := fs.Uint64("id", 0, "listing id")
	value := fs.String("value", "", "ETH value to send (wei)")
	_ = fs.Parse(args)
	if *marketplace == "" || *value == "" {
		log.Fatalf("--marketplace and --value are required")
	}

	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	ma := common.HexToAddress(*marketplace)
	mkt, err := contracts.NewMarketplace(ma, cli)
	if err != nil {
		log.Fatalf("bind marketplace: %v", err)
	}

	val, ok := new(big.Int).SetString(*value, 10)
	if !ok {
		log.Fatalf("bad --value")
	}

	auth.Value = val
	tx, err := mkt.Buy(auth, new(big.Int).SetUint64(*id))
	if err != nil {
		log.Fatalf("marketplace.buy: %v", err)
	}
	fmt.Printf("Buy (direct) tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Buy (direct)")
}

func cmdBuyDirectAtPrice(args []string) {
	fs := flag.NewFlagSet("buy-direct-at-price", flag.ExitOnError)
	marketplace := fs.String("marketplace", "", "Marketplace address (required)")
	id := fs.Uint64("id", 0, "listing id")
	_ = fs.Parse(args)
	if *marketplace == "" {
		log.Fatalf("--marketplace is required")
	}

	env := mustEnv()
	cli := dial(env.RPCURL)
	auth, _ := authFromPK(env.PrivKey, env.ChainID)

	ma := common.HexToAddress(*marketplace)
	mkt, err := contracts.NewMarketplace(ma, cli)
	if err != nil {
		log.Fatalf("bind marketplace: %v", err)
	}

	li, err := mkt.GetListing(&bind.CallOpts{}, new(big.Int).SetUint64(*id))
	if err != nil {
		log.Fatalf("getListing (direct): %v", err)
	}
	if !li.Active {
		log.Fatalf("listing %d is not active", *id)
	}

	auth.Value = li.Price
	tx, err := mkt.Buy(auth, new(big.Int).SetUint64(*id))
	if err != nil {
		log.Fatalf("marketplace.buy (at price): %v", err)
	}
	fmt.Printf("Buy (direct at price) tx: %s\n", tx.Hash())
	waitMined(cli, tx, "Buy (direct at price)")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	switch os.Args[1] {
	case "get-token":
		cmdGetToken(os.Args[2:])
	case "next-listing-id":
		cmdNextListingID(os.Args[2:])
	case "get-user", "get-profile":
		cmdGetUser(os.Args[2:])
	case "set-balance":
		cmdSetBalance(os.Args[2:])
	case "set-profile":
		cmdSetProfile(os.Args[2:])
	case "whoami":
		cmdWhoami(os.Args[2:])
	case "deploy-token":
		cmdDeployToken(os.Args[2:])
	case "deploy-marketplace":
		cmdDeployMarketplace(os.Args[2:])
	case "deploy-interactor":
		cmdDeployInteractor(os.Args[2:])
	case "approve":
		cmdApprove(os.Args[2:])
	case "list":
		cmdList(os.Args[2:])
	case "buy":
		cmdBuy(os.Args[2:])
	case "profile":
		cmdProfile(os.Args[2:])
	case "listing":
		cmdListing(os.Args[2:])
	case "list-direct":
		cmdListDirect(os.Args[2:])
	case "buy-direct":
		cmdBuyDirect(os.Args[2:])
	case "buy-direct-at-price":
		cmdBuyDirectAtPrice(os.Args[2:])

	default:
		usage()
	}
}
