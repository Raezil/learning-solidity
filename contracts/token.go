// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"NotEnoughFunds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWasTransfered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"createOrUpdateProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"profiles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"updateProfileBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTo\",\"type\":\"address\"}],\"name\":\"updateProfileTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506104a68061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c806347e7ef241161006357806347e7ef24146101a2578063586d524a146101b5578063b7760c8f146101c8578063bbe15627146101db578063f58f34711461020a575f5ffd5b806308a0dfa6146100945780630f53a470146100d457806312065fe01461012b578063288a18fe14610148575b5f5ffd5b6100d26100a2366004610346565b6001600160a01b039182165f90815260208190526040902060010180546001600160a01b03191691909216179055565b005b6101096100e2366004610377565b6001600160a01b039081165f90815260208190526040902080546001909101549092911690565b604080519283526001600160a01b039091166020830152015b60405180910390f35b335f90815260208190526040902054604051908152602001610122565b6100d2610156366004610397565b6040805180820182529283526001600160a01b0391821660208085019182529483165f908152948590529320915182559151600190910180546001600160a01b03191691909216179055565b6100d26101b03660046103d0565b610233565b6100d26101c33660046103f8565b610263565b6100d26101d636600461040f565b610295565b6101096101e9366004610377565b5f60208190529081526040902080546001909101546001600160a01b031682565b6100d26102183660046103d0565b6001600160a01b039091165f90815260208190526040902055565b6001600160a01b0382165f908152602081905260408120805483929061025a908490610444565b90915550505050565b80610279335f9081526020819052604090205490565b610283919061045d565b335f9081526020819052604090205550565b335f90815260208190526040902054828110156102d3576040516311920a6d60e31b8152600481018490526024810182905260440160405180910390fd5b6102dc83610263565b6102e68284610233565b6040518381526001600160a01b0383169033907f0b36a9283d9787f888df0e5d4959e6c09dfad91ccf85a3ec88d17e562c0f17559060200160405180910390a3505050565b80356001600160a01b0381168114610341575f5ffd5b919050565b5f5f60408385031215610357575f5ffd5b6103608361032b565b915061036e6020840161032b565b90509250929050565b5f60208284031215610387575f5ffd5b6103908261032b565b9392505050565b5f5f5f606084860312156103a9575f5ffd5b6103b28461032b565b9250602084013591506103c76040850161032b565b90509250925092565b5f5f604083850312156103e1575f5ffd5b6103ea8361032b565b946020939093013593505050565b5f60208284031215610408575f5ffd5b5035919050565b5f5f60408385031215610420575f5ffd5b8235915061036e6020840161032b565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561045757610457610430565b92915050565b818103818111156104575761045761043056fea26469706673582212205a67a1c81014ed427d097f7967fca86c4527ea4333a34ca07a762d3c3d349bed64736f6c634300081e0033",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenMetaData.Bin instead.
var TokenBin = TokenMetaData.Bin

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token *TokenCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token *TokenSession) GetBalance() (*big.Int, error) {
	return _Token.Contract.GetBalance(&_Token.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Token *TokenCallerSession) GetBalance() (*big.Int, error) {
	return _Token.Contract.GetBalance(&_Token.CallOpts)
}

// GetProfile is a free data retrieval call binding the contract method 0x0f53a470.
//
// Solidity: function getProfile(address user) view returns(uint256 balance, address to)
func (_Token *TokenCaller) GetProfile(opts *bind.CallOpts, user common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getProfile", user)

	outstruct := new(struct {
		Balance *big.Int
		To      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.To = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetProfile is a free data retrieval call binding the contract method 0x0f53a470.
//
// Solidity: function getProfile(address user) view returns(uint256 balance, address to)
func (_Token *TokenSession) GetProfile(user common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	return _Token.Contract.GetProfile(&_Token.CallOpts, user)
}

// GetProfile is a free data retrieval call binding the contract method 0x0f53a470.
//
// Solidity: function getProfile(address user) view returns(uint256 balance, address to)
func (_Token *TokenCallerSession) GetProfile(user common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	return _Token.Contract.GetProfile(&_Token.CallOpts, user)
}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(uint256 balance, address to)
func (_Token *TokenCaller) Profiles(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "profiles", arg0)

	outstruct := new(struct {
		Balance *big.Int
		To      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.To = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(uint256 balance, address to)
func (_Token *TokenSession) Profiles(arg0 common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	return _Token.Contract.Profiles(&_Token.CallOpts, arg0)
}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(uint256 balance, address to)
func (_Token *TokenCallerSession) Profiles(arg0 common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	return _Token.Contract.Profiles(&_Token.CallOpts, arg0)
}

// CreateOrUpdateProfile is a paid mutator transaction binding the contract method 0x288a18fe.
//
// Solidity: function createOrUpdateProfile(address user, uint256 balance, address to) returns()
func (_Token *TokenTransactor) CreateOrUpdateProfile(opts *bind.TransactOpts, user common.Address, balance *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createOrUpdateProfile", user, balance, to)
}

// CreateOrUpdateProfile is a paid mutator transaction binding the contract method 0x288a18fe.
//
// Solidity: function createOrUpdateProfile(address user, uint256 balance, address to) returns()
func (_Token *TokenSession) CreateOrUpdateProfile(user common.Address, balance *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.CreateOrUpdateProfile(&_Token.TransactOpts, user, balance, to)
}

// CreateOrUpdateProfile is a paid mutator transaction binding the contract method 0x288a18fe.
//
// Solidity: function createOrUpdateProfile(address user, uint256 balance, address to) returns()
func (_Token *TokenTransactorSession) CreateOrUpdateProfile(user common.Address, balance *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.CreateOrUpdateProfile(&_Token.TransactOpts, user, balance, to)
}

// Deduct is a paid mutator transaction binding the contract method 0x586d524a.
//
// Solidity: function deduct(uint256 amount) returns()
func (_Token *TokenTransactor) Deduct(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "deduct", amount)
}

// Deduct is a paid mutator transaction binding the contract method 0x586d524a.
//
// Solidity: function deduct(uint256 amount) returns()
func (_Token *TokenSession) Deduct(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deduct(&_Token.TransactOpts, amount)
}

// Deduct is a paid mutator transaction binding the contract method 0x586d524a.
//
// Solidity: function deduct(uint256 amount) returns()
func (_Token *TokenTransactorSession) Deduct(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deduct(&_Token.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns()
func (_Token *TokenTransactor) Deposit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "deposit", to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns()
func (_Token *TokenSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns()
func (_Token *TokenTransactorSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(uint256 amount, address to) returns()
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", amount, to)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(uint256 amount, address to) returns()
func (_Token *TokenSession) Transfer(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, amount, to)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(uint256 amount, address to) returns()
func (_Token *TokenTransactorSession) Transfer(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, amount, to)
}

// UpdateProfileBalance is a paid mutator transaction binding the contract method 0xf58f3471.
//
// Solidity: function updateProfileBalance(address user, uint256 newBalance) returns()
func (_Token *TokenTransactor) UpdateProfileBalance(opts *bind.TransactOpts, user common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateProfileBalance", user, newBalance)
}

// UpdateProfileBalance is a paid mutator transaction binding the contract method 0xf58f3471.
//
// Solidity: function updateProfileBalance(address user, uint256 newBalance) returns()
func (_Token *TokenSession) UpdateProfileBalance(user common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateProfileBalance(&_Token.TransactOpts, user, newBalance)
}

// UpdateProfileBalance is a paid mutator transaction binding the contract method 0xf58f3471.
//
// Solidity: function updateProfileBalance(address user, uint256 newBalance) returns()
func (_Token *TokenTransactorSession) UpdateProfileBalance(user common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateProfileBalance(&_Token.TransactOpts, user, newBalance)
}

// UpdateProfileTo is a paid mutator transaction binding the contract method 0x08a0dfa6.
//
// Solidity: function updateProfileTo(address user, address newTo) returns()
func (_Token *TokenTransactor) UpdateProfileTo(opts *bind.TransactOpts, user common.Address, newTo common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateProfileTo", user, newTo)
}

// UpdateProfileTo is a paid mutator transaction binding the contract method 0x08a0dfa6.
//
// Solidity: function updateProfileTo(address user, address newTo) returns()
func (_Token *TokenSession) UpdateProfileTo(user common.Address, newTo common.Address) (*types.Transaction, error) {
	return _Token.Contract.UpdateProfileTo(&_Token.TransactOpts, user, newTo)
}

// UpdateProfileTo is a paid mutator transaction binding the contract method 0x08a0dfa6.
//
// Solidity: function updateProfileTo(address user, address newTo) returns()
func (_Token *TokenTransactorSession) UpdateProfileTo(user common.Address, newTo common.Address) (*types.Transaction, error) {
	return _Token.Contract.UpdateProfileTo(&_Token.TransactOpts, user, newTo)
}

// TokenTokenWasTransferedIterator is returned from FilterTokenWasTransfered and is used to iterate over the raw logs and unpacked data for TokenWasTransfered events raised by the Token contract.
type TokenTokenWasTransferedIterator struct {
	Event *TokenTokenWasTransfered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenTokenWasTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTokenWasTransfered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenTokenWasTransfered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenTokenWasTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTokenWasTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTokenWasTransfered represents a TokenWasTransfered event raised by the Token contract.
type TokenTokenWasTransfered struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWasTransfered is a free log retrieval operation binding the contract event 0x0b36a9283d9787f888df0e5d4959e6c09dfad91ccf85a3ec88d17e562c0f1755.
//
// Solidity: event TokenWasTransfered(address indexed from, address indexed to, uint256 amount)
func (_Token *TokenFilterer) FilterTokenWasTransfered(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTokenWasTransferedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "TokenWasTransfered", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTokenWasTransferedIterator{contract: _Token.contract, event: "TokenWasTransfered", logs: logs, sub: sub}, nil
}

// WatchTokenWasTransfered is a free log subscription operation binding the contract event 0x0b36a9283d9787f888df0e5d4959e6c09dfad91ccf85a3ec88d17e562c0f1755.
//
// Solidity: event TokenWasTransfered(address indexed from, address indexed to, uint256 amount)
func (_Token *TokenFilterer) WatchTokenWasTransfered(opts *bind.WatchOpts, sink chan<- *TokenTokenWasTransfered, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "TokenWasTransfered", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTokenWasTransfered)
				if err := _Token.contract.UnpackLog(event, "TokenWasTransfered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenWasTransfered is a log parse operation binding the contract event 0x0b36a9283d9787f888df0e5d4959e6c09dfad91ccf85a3ec88d17e562c0f1755.
//
// Solidity: event TokenWasTransfered(address indexed from, address indexed to, uint256 amount)
func (_Token *TokenFilterer) ParseTokenWasTransfered(log types.Log) (*TokenTokenWasTransfered, error) {
	event := new(TokenTokenWasTransfered)
	if err := _Token.contract.UnpackLog(event, "TokenWasTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
