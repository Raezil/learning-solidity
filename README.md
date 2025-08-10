<p align="center">
  <img width="1024" height="1024" alt="logo" src="https://github.com/user-attachments/assets/39fd81db-691e-4ef4-baf6-13caf0981f82" />
</p>
A tiny, runnable CLI dApp that deploys and uses a Token + Marketplace + Interactor using your abigen bindings.

## Prereqs
- Go 1.22+
- A JSON-RPC node (e.g., Anvil): `anvil --chain-id 31337`
- Export a funded private key from the node

## Setup
```bash
cd market-dapp
go mod tidy
```

## Env
```bash
export RPC_URL=http://127.0.0.1:8545
export PRIVATE_KEY=<hex-no-0x>
export CHAIN_ID=31337
```

## Deploy
```bash
go run . deploy-token
# copy token address
go run . deploy-marketplace --token <TOKEN_ADDR>
# copy marketplace address
go run . deploy-interactor --marketplace <MARKETPLACE_ADDR>
```

## Use
```bash
# Who am I
go run . whoami

# Approve token spending for marketplace (via Interactor or directly)
go run . approve --interactor <INTERACTOR_ADDR> --amount 1000000000000000000

# List an item (price is an integer according to your Token's unit semantics)
go run . list --interactor <INTERACTOR_ADDR> --price 250000000000000000

# Buy (send ETH only if your Marketplace requires it)
go run . buy --interactor <INTERACTOR_ADDR> --id 0 --value 0

# Inspect profile and listing
go run . profile --interactor <INTERACTOR_ADDR>
go run . listing --interactor <INTERACTOR_ADDR> --id 0
```

> Adjust amounts to your token economics. The CLI uses `bind.NewKeyedTransactorWithChainID`
> with dynamic fees suggested by the node.
