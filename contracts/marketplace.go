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

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Listed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextListingId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060405161067a38038061067a833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b6105f4806100865f395ff3fe608060405260043610610079575f3560e01c8063d96a094a1161004c578063d96a094a1461016d578063dd62ed3e14610180578063de74e57b146101b6578063fc0c546a146101fe575f5ffd5b8063107a274a1461007d5780632ce381901461010a5780632eb50ef71461012b578063aaccf1ec1461014a575b5f5ffd5b348015610088575f5ffd5b506100de610097366004610537565b5f90815260026020818152604092839020835160608101855281546001600160a01b03168082526001830154938201849052919093015460ff161515929093018290529192565b604080516001600160a01b03909416845260208401929092521515908201526060015b60405180910390f35b348015610115575f5ffd5b50610129610124366004610537565b610234565b005b348015610136575f5ffd5b50610129610145366004610537565b610322565b348015610155575f5ffd5b5061015f60015481565b604051908152602001610101565b61012961017b366004610537565b610376565b34801561018b575f5ffd5b5061015f61019a366004610569565b600360209081525f928352604080842090915290825290205481565b3480156101c1575f5ffd5b506100de6101d0366004610537565b600260208190525f91825260409091208054600182015491909201546001600160a01b039092169160ff1683565b348015610209575f5ffd5b505f5461021c906001600160a01b031681565b6040516001600160a01b039091168152602001610101565b5f81116102725760405162461bcd60e51b8152602060048201526007602482015266050726963653e360cc1b60448201526064015b60405180910390fd5b60408051606081018252338082526020808301858152600184860181815281545f908152600280865290889020965187546001600160a01b0319166001600160a01b039091161787559251868301555194909101805460ff191694151594909417909355915492518481529092917f50955776c5778c3b7d968d86d8c51fb6b29a7a74c20866b533268e209fc08343910160405180910390a360018054905f61031a8361059a565b919050555050565b335f8181526003602090815260408083203080855290835292819020859055518481529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350565b5f8181526002602081905260409091209081015460ff166103c65760405162461bcd60e51b815260206004820152600a6024820152694e6f742061637469766560b01b6044820152606401610269565b600181015481546001600160a01b03165f90815260036020908152604080832030845290915290205410156104505760405162461bcd60e51b815260206004820152602a60248201527f4d61726b6574706c616365206e6f7420617070726f76656420666f7220656e6f60448201526975676820746f6b656e7360b01b6064820152608401610269565b5f54600182015460405163b7760c8f60e01b815260048101919091523360248201526001600160a01b039091169063b7760c8f906044015f604051808303815f87803b15801561049e575f5ffd5b505af11580156104b0573d5f5f3e3d5ffd5b505082546040516001600160a01b0390911692503480156108fc029250905f818181858888f193505050501580156104ea573d5f5f3e3d5ffd5b5060028101805460ff191690556001810154604051908152339083907fd2728f908c7e0feb83c6278798370fcb86b62f236c9dbf1a3f541096c21590409060200160405180910390a35050565b5f60208284031215610547575f5ffd5b5035919050565b80356001600160a01b0381168114610564575f5ffd5b919050565b5f5f6040838503121561057a575f5ffd5b6105838361054e565b91506105916020840161054e565b90509250929050565b5f600182016105b757634e487b7160e01b5f52601160045260245ffd5b506001019056fea2646970667358221220f2b0fbde2ccdbad884d29d35f8d2bb9bc7ef445dedd9c8b7b6dbde6fe6f3902f64736f6c634300081e0033",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// MarketplaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketplaceMetaData.Bin instead.
var MarketplaceBin = MarketplaceMetaData.Bin

// DeployMarketplace deploys a new Ethereum contract, binding an instance of Marketplace to it.
func DeployMarketplace(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Marketplace, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketplaceBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Marketplace *MarketplaceSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Marketplace.Contract.Allowance(&_Marketplace.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Marketplace.Contract.Allowance(&_Marketplace.CallOpts, arg0, arg1)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns(address seller, uint256 price, bool active)
func (_Marketplace *MarketplaceCaller) GetListing(opts *bind.CallOpts, listingId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListing", listingId)

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
func (_Marketplace *MarketplaceSession) GetListing(listingId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	return _Marketplace.Contract.GetListing(&_Marketplace.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns(address seller, uint256 price, bool active)
func (_Marketplace *MarketplaceCallerSession) GetListing(listingId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	return _Marketplace.Contract.GetListing(&_Marketplace.CallOpts, listingId)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint256 price, bool active)
func (_Marketplace *MarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "listings", arg0)

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

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint256 price, bool active)
func (_Marketplace *MarketplaceSession) Listings(arg0 *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint256 price, bool active)
func (_Marketplace *MarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Active bool
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// NextListingId is a free data retrieval call binding the contract method 0xaaccf1ec.
//
// Solidity: function nextListingId() view returns(uint256)
func (_Marketplace *MarketplaceCaller) NextListingId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nextListingId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextListingId is a free data retrieval call binding the contract method 0xaaccf1ec.
//
// Solidity: function nextListingId() view returns(uint256)
func (_Marketplace *MarketplaceSession) NextListingId() (*big.Int, error) {
	return _Marketplace.Contract.NextListingId(&_Marketplace.CallOpts)
}

// NextListingId is a free data retrieval call binding the contract method 0xaaccf1ec.
//
// Solidity: function nextListingId() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) NextListingId() (*big.Int, error) {
	return _Marketplace.Contract.NextListingId(&_Marketplace.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Marketplace *MarketplaceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Marketplace *MarketplaceSession) Token() (common.Address, error) {
	return _Marketplace.Contract.Token(&_Marketplace.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Token() (common.Address, error) {
	return _Marketplace.Contract.Token(&_Marketplace.CallOpts)
}

// ApproveMarketplace is a paid mutator transaction binding the contract method 0x2eb50ef7.
//
// Solidity: function approveMarketplace(uint256 amount) returns()
func (_Marketplace *MarketplaceTransactor) ApproveMarketplace(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "approveMarketplace", amount)
}

// ApproveMarketplace is a paid mutator transaction binding the contract method 0x2eb50ef7.
//
// Solidity: function approveMarketplace(uint256 amount) returns()
func (_Marketplace *MarketplaceSession) ApproveMarketplace(amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.ApproveMarketplace(&_Marketplace.TransactOpts, amount)
}

// ApproveMarketplace is a paid mutator transaction binding the contract method 0x2eb50ef7.
//
// Solidity: function approveMarketplace(uint256 amount) returns()
func (_Marketplace *MarketplaceTransactorSession) ApproveMarketplace(amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.ApproveMarketplace(&_Marketplace.TransactOpts, amount)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 listingId) payable returns()
func (_Marketplace *MarketplaceTransactor) Buy(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buy", listingId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 listingId) payable returns()
func (_Marketplace *MarketplaceSession) Buy(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, listingId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 listingId) payable returns()
func (_Marketplace *MarketplaceTransactorSession) Buy(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, listingId)
}

// ListItem is a paid mutator transaction binding the contract method 0x2ce38190.
//
// Solidity: function listItem(uint256 price) returns()
func (_Marketplace *MarketplaceTransactor) ListItem(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "listItem", price)
}

// ListItem is a paid mutator transaction binding the contract method 0x2ce38190.
//
// Solidity: function listItem(uint256 price) returns()
func (_Marketplace *MarketplaceSession) ListItem(price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.ListItem(&_Marketplace.TransactOpts, price)
}

// ListItem is a paid mutator transaction binding the contract method 0x2ce38190.
//
// Solidity: function listItem(uint256 price) returns()
func (_Marketplace *MarketplaceTransactorSession) ListItem(price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.ListItem(&_Marketplace.TransactOpts, price)
}

// MarketplaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Marketplace contract.
type MarketplaceApprovalIterator struct {
	Event *MarketplaceApproval // Event containing the contract specifics and raw log

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
func (it *MarketplaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceApproval)
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
		it.Event = new(MarketplaceApproval)
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
func (it *MarketplaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceApproval represents a Approval event raised by the Marketplace contract.
type MarketplaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Marketplace *MarketplaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MarketplaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceApprovalIterator{contract: _Marketplace.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Marketplace *MarketplaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MarketplaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceApproval)
				if err := _Marketplace.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Marketplace *MarketplaceFilterer) ParseApproval(log types.Log) (*MarketplaceApproval, error) {
	event := new(MarketplaceApproval)
	if err := _Marketplace.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the Marketplace contract.
type MarketplaceBoughtIterator struct {
	Event *MarketplaceBought // Event containing the contract specifics and raw log

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
func (it *MarketplaceBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBought)
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
		it.Event = new(MarketplaceBought)
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
func (it *MarketplaceBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBought represents a Bought event raised by the Marketplace contract.
type MarketplaceBought struct {
	ListingId *big.Int
	Buyer     common.Address
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0xd2728f908c7e0feb83c6278798370fcb86b62f236c9dbf1a3f541096c2159040.
//
// Solidity: event Bought(uint256 indexed listingId, address indexed buyer, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterBought(opts *bind.FilterOpts, listingId []*big.Int, buyer []common.Address) (*MarketplaceBoughtIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Bought", listingIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceBoughtIterator{contract: _Marketplace.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0xd2728f908c7e0feb83c6278798370fcb86b62f236c9dbf1a3f541096c2159040.
//
// Solidity: event Bought(uint256 indexed listingId, address indexed buyer, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *MarketplaceBought, listingId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Bought", listingIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBought)
				if err := _Marketplace.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0xd2728f908c7e0feb83c6278798370fcb86b62f236c9dbf1a3f541096c2159040.
//
// Solidity: event Bought(uint256 indexed listingId, address indexed buyer, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParseBought(log types.Log) (*MarketplaceBought, error) {
	event := new(MarketplaceBought)
	if err := _Marketplace.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListedIterator is returned from FilterListed and is used to iterate over the raw logs and unpacked data for Listed events raised by the Marketplace contract.
type MarketplaceListedIterator struct {
	Event *MarketplaceListed // Event containing the contract specifics and raw log

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
func (it *MarketplaceListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListed)
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
		it.Event = new(MarketplaceListed)
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
func (it *MarketplaceListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListed represents a Listed event raised by the Marketplace contract.
type MarketplaceListed struct {
	ListingId *big.Int
	Seller    common.Address
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListed is a free log retrieval operation binding the contract event 0x50955776c5778c3b7d968d86d8c51fb6b29a7a74c20866b533268e209fc08343.
//
// Solidity: event Listed(uint256 indexed listingId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterListed(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address) (*MarketplaceListedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Listed", listingIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListedIterator{contract: _Marketplace.contract, event: "Listed", logs: logs, sub: sub}, nil
}

// WatchListed is a free log subscription operation binding the contract event 0x50955776c5778c3b7d968d86d8c51fb6b29a7a74c20866b533268e209fc08343.
//
// Solidity: event Listed(uint256 indexed listingId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchListed(opts *bind.WatchOpts, sink chan<- *MarketplaceListed, listingId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Listed", listingIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListed)
				if err := _Marketplace.contract.UnpackLog(event, "Listed", log); err != nil {
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

// ParseListed is a log parse operation binding the contract event 0x50955776c5778c3b7d968d86d8c51fb6b29a7a74c20866b533268e209fc08343.
//
// Solidity: event Listed(uint256 indexed listingId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParseListed(log types.Log) (*MarketplaceListed, error) {
	event := new(MarketplaceListed)
	if err := _Marketplace.contract.UnpackLog(event, "Listed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
