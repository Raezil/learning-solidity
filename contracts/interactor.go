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

// InteractorMetaData contains all meta data concerning the Interactor contract.
var InteractorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"getAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"contractMarketplace\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161069838038061069883398101604081905261002e916100e5565b5f80546001600160a01b0319166001600160a01b03831690811790915560408051637e062a3560e11b8152905163fc0c546a916004808201926020929091908290030181865afa158015610084573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100a891906100e5565b600180546001600160a01b0319166001600160a01b039290921691909117905550610107565b6001600160a01b03811681146100e2575f5ffd5b50565b5f602082840312156100f5575f5ffd5b8151610100816100ce565b9392505050565b610584806101145f395ff3fe608060405260043610610079575f3560e01c806355a373d61161004c57806355a373d614610144578063abc8c7af1461017b578063e7fb74c714610199578063eb5a662e146101ac575f5ffd5b80630f53a4701461007d578063107a274a146100be5780632ce38190146101045780632eb50ef714610125575b5f5ffd5b348015610088575f5ffd5b5061009c61009736600461048b565b6101d9565b604080519283526001600160a01b039091166020830152015b60405180910390f35b3480156100c9575f5ffd5b506100dd6100d83660046104ad565b610251565b604080516001600160a01b03909416845260208401929092521515908201526060016100b5565b34801561010f575f5ffd5b5061012361011e3660046104ad565b6102cd565b005b348015610130575f5ffd5b5061012361013f3660046104ad565b610329565b34801561014f575f5ffd5b50600154610163906001600160a01b031681565b6040516001600160a01b0390911681526020016100b5565b348015610186575f5ffd5b505f54610163906001600160a01b031681565b6101236101a73660046104ad565b610359565b3480156101b7575f5ffd5b506101cb6101c636600461048b565b6103fc565b6040519081526020016100b5565b60015460405163bbe1562760e01b81526001600160a01b0383811660048301525f92839291169063bbe15627906024016040805180830381865afa158015610223573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061024791906104c4565b9094909350915050565b5f805460405163083d13a560e11b815260048101849052829182916001600160a01b039091169063107a274a90602401606060405180830381865afa15801561029c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102c091906104f3565b9196909550909350915050565b5f546040516302ce381960e41b8152600481018390526001600160a01b0390911690632ce38190906024015b5f604051808303815f87803b158015610310575f5ffd5b505af1158015610322573d5f5f3e3d5ffd5b5050505050565b5f54604051632eb50ef760e01b8152600481018390526001600160a01b0390911690632eb50ef7906024016102f9565b5f341161039e5760405162461bcd60e51b815260206004820152600f60248201526e53656e642045544820746f2062757960881b604482015260640160405180910390fd5b5f54604051636cb504a560e11b8152600481018390526001600160a01b039091169063d96a094a9034906024015f604051808303818588803b1580156103e2575f5ffd5b505af11580156103f4573d5f5f3e3d5ffd5b505050505050565b5f8054604051636eb1769f60e11b81526001600160a01b0384811660048301523060248301529091169063dd62ed3e90604401602060405180830381865afa15801561044a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046e9190610537565b92915050565b6001600160a01b0381168114610488575f5ffd5b50565b5f6020828403121561049b575f5ffd5b81356104a681610474565b9392505050565b5f602082840312156104bd575f5ffd5b5035919050565b5f5f604083850312156104d5575f5ffd5b825160208401519092506104e881610474565b809150509250929050565b5f5f5f60608486031215610505575f5ffd5b835161051081610474565b602085015160408601519194509250801515811461052c575f5ffd5b809150509250925092565b5f60208284031215610547575f5ffd5b505191905056fea264697066735822122087233b1050d0a1f386f2dde30c26be6cbcd5ab61a21a0ed6f4b6180701c31df064736f6c634300081e0033",
}

// InteractorABI is the input ABI used to generate the binding from.
// Deprecated: Use InteractorMetaData.ABI instead.
var InteractorABI = InteractorMetaData.ABI

// InteractorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InteractorMetaData.Bin instead.
var InteractorBin = InteractorMetaData.Bin

// DeployInteractor deploys a new Ethereum contract, binding an instance of Interactor to it.
func DeployInteractor(auth *bind.TransactOpts, backend bind.ContractBackend, _marketplaceAddr common.Address) (common.Address, *types.Transaction, *Interactor, error) {
	parsed, err := InteractorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InteractorBin), backend, _marketplaceAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Interactor{InteractorCaller: InteractorCaller{contract: contract}, InteractorTransactor: InteractorTransactor{contract: contract}, InteractorFilterer: InteractorFilterer{contract: contract}}, nil
}

// Interactor is an auto generated Go binding around an Ethereum contract.
type Interactor struct {
	InteractorCaller     // Read-only binding to the contract
	InteractorTransactor // Write-only binding to the contract
	InteractorFilterer   // Log filterer for contract events
}

// InteractorCaller is an auto generated read-only Go binding around an Ethereum contract.
type InteractorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InteractorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InteractorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InteractorSession struct {
	Contract     *Interactor       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InteractorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InteractorCallerSession struct {
	Contract *InteractorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// InteractorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InteractorTransactorSession struct {
	Contract     *InteractorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InteractorRaw is an auto generated low-level Go binding around an Ethereum contract.
type InteractorRaw struct {
	Contract *Interactor // Generic contract binding to access the raw methods on
}

// InteractorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InteractorCallerRaw struct {
	Contract *InteractorCaller // Generic read-only contract binding to access the raw methods on
}

// InteractorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InteractorTransactorRaw struct {
	Contract *InteractorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInteractor creates a new instance of Interactor, bound to a specific deployed contract.
func NewInteractor(address common.Address, backend bind.ContractBackend) (*Interactor, error) {
	contract, err := bindInteractor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Interactor{InteractorCaller: InteractorCaller{contract: contract}, InteractorTransactor: InteractorTransactor{contract: contract}, InteractorFilterer: InteractorFilterer{contract: contract}}, nil
}

// NewInteractorCaller creates a new read-only instance of Interactor, bound to a specific deployed contract.
func NewInteractorCaller(address common.Address, caller bind.ContractCaller) (*InteractorCaller, error) {
	contract, err := bindInteractor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InteractorCaller{contract: contract}, nil
}

// NewInteractorTransactor creates a new write-only instance of Interactor, bound to a specific deployed contract.
func NewInteractorTransactor(address common.Address, transactor bind.ContractTransactor) (*InteractorTransactor, error) {
	contract, err := bindInteractor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InteractorTransactor{contract: contract}, nil
}

// NewInteractorFilterer creates a new log filterer instance of Interactor, bound to a specific deployed contract.
func NewInteractorFilterer(address common.Address, filterer bind.ContractFilterer) (*InteractorFilterer, error) {
	contract, err := bindInteractor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InteractorFilterer{contract: contract}, nil
}

// bindInteractor binds a generic wrapper to an already deployed contract.
func bindInteractor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InteractorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interactor *InteractorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interactor.Contract.InteractorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interactor *InteractorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interactor.Contract.InteractorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interactor *InteractorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interactor.Contract.InteractorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interactor *InteractorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interactor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interactor *InteractorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interactor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interactor *InteractorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interactor.Contract.contract.Transact(opts, method, params...)
}

// GetAllowance is a free data retrieval call binding the contract method 0xeb5a662e.
//
// Solidity: function getAllowance(address seller) view returns(uint256)
func (_Interactor *InteractorCaller) GetAllowance(opts *bind.CallOpts, seller common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interactor.contract.Call(opts, &out, "getAllowance", seller)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllowance is a free data retrieval call binding the contract method 0xeb5a662e.
//
// Solidity: function getAllowance(address seller) view returns(uint256)
func (_Interactor *InteractorSession) GetAllowance(seller common.Address) (*big.Int, error) {
	return _Interactor.Contract.GetAllowance(&_Interactor.CallOpts, seller)
}

// GetAllowance is a free data retrieval call binding the contract method 0xeb5a662e.
//
// Solidity: function getAllowance(address seller) view returns(uint256)
func (_Interactor *InteractorCallerSession) GetAllowance(seller common.Address) (*big.Int, error) {
	return _Interactor.Contract.GetAllowance(&_Interactor.CallOpts, seller)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns(address seller, uint256 price, bool active)
func (_Interactor *InteractorCaller) GetListing(opts *bind.CallOpts, listingId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	var out []interface{}
	err := _Interactor.contract.Call(opts, &out, "getListing", listingId)

	outstruct := new(struct {
		Seller common.Address
		Price  *big.Int
		Active bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns(address seller, uint256 price, bool active)
func (_Interactor *InteractorSession) GetListing(listingId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	return _Interactor.Contract.GetListing(&_Interactor.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns(address seller, uint256 price, bool active)
func (_Interactor *InteractorCallerSession) GetListing(listingId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	return _Interactor.Contract.GetListing(&_Interactor.CallOpts, listingId)
}

// GetProfile is a free data retrieval call binding the contract method 0x0f53a470.
//
// Solidity: function getProfile(address user) view returns(uint256 balance, address to)
func (_Interactor *InteractorCaller) GetProfile(opts *bind.CallOpts, user common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	var out []interface{}
	err := _Interactor.contract.Call(opts, &out, "getProfile", user)

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
func (_Interactor *InteractorSession) GetProfile(user common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	return _Interactor.Contract.GetProfile(&_Interactor.CallOpts, user)
}

// GetProfile is a free data retrieval call binding the contract method 0x0f53a470.
//
// Solidity: function getProfile(address user) view returns(uint256 balance, address to)
func (_Interactor *InteractorCallerSession) GetProfile(user common.Address) (struct {
	Balance *big.Int
	To      common.Address
}, error) {
	return _Interactor.Contract.GetProfile(&_Interactor.CallOpts, user)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Interactor *InteractorCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interactor.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Interactor *InteractorSession) Marketplace() (common.Address, error) {
	return _Interactor.Contract.Marketplace(&_Interactor.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Interactor *InteractorCallerSession) Marketplace() (common.Address, error) {
	return _Interactor.Contract.Marketplace(&_Interactor.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Interactor *InteractorCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interactor.contract.Call(opts, &out, "tokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Interactor *InteractorSession) TokenContract() (common.Address, error) {
	return _Interactor.Contract.TokenContract(&_Interactor.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Interactor *InteractorCallerSession) TokenContract() (common.Address, error) {
	return _Interactor.Contract.TokenContract(&_Interactor.CallOpts)
}

// ApproveMarketplace is a paid mutator transaction binding the contract method 0x2eb50ef7.
//
// Solidity: function approveMarketplace(uint256 amount) returns()
func (_Interactor *InteractorTransactor) ApproveMarketplace(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Interactor.contract.Transact(opts, "approveMarketplace", amount)
}

// ApproveMarketplace is a paid mutator transaction binding the contract method 0x2eb50ef7.
//
// Solidity: function approveMarketplace(uint256 amount) returns()
func (_Interactor *InteractorSession) ApproveMarketplace(amount *big.Int) (*types.Transaction, error) {
	return _Interactor.Contract.ApproveMarketplace(&_Interactor.TransactOpts, amount)
}

// ApproveMarketplace is a paid mutator transaction binding the contract method 0x2eb50ef7.
//
// Solidity: function approveMarketplace(uint256 amount) returns()
func (_Interactor *InteractorTransactorSession) ApproveMarketplace(amount *big.Int) (*types.Transaction, error) {
	return _Interactor.Contract.ApproveMarketplace(&_Interactor.TransactOpts, amount)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 listingId) payable returns()
func (_Interactor *InteractorTransactor) BuyItem(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _Interactor.contract.Transact(opts, "buyItem", listingId)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 listingId) payable returns()
func (_Interactor *InteractorSession) BuyItem(listingId *big.Int) (*types.Transaction, error) {
	return _Interactor.Contract.BuyItem(&_Interactor.TransactOpts, listingId)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 listingId) payable returns()
func (_Interactor *InteractorTransactorSession) BuyItem(listingId *big.Int) (*types.Transaction, error) {
	return _Interactor.Contract.BuyItem(&_Interactor.TransactOpts, listingId)
}

// ListItem is a paid mutator transaction binding the contract method 0x2ce38190.
//
// Solidity: function listItem(uint256 price) returns()
func (_Interactor *InteractorTransactor) ListItem(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Interactor.contract.Transact(opts, "listItem", price)
}

// ListItem is a paid mutator transaction binding the contract method 0x2ce38190.
//
// Solidity: function listItem(uint256 price) returns()
func (_Interactor *InteractorSession) ListItem(price *big.Int) (*types.Transaction, error) {
	return _Interactor.Contract.ListItem(&_Interactor.TransactOpts, price)
}

// ListItem is a paid mutator transaction binding the contract method 0x2ce38190.
//
// Solidity: function listItem(uint256 price) returns()
func (_Interactor *InteractorTransactorSession) ListItem(price *big.Int) (*types.Transaction, error) {
	return _Interactor.Contract.ListItem(&_Interactor.TransactOpts, price)
}
