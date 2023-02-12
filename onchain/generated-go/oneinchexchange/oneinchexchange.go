// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinchexchange

import (
	"math/big"
	"strings"

	ethereum "github.com/ledgerwatch/erigon"
	"github.com/ledgerwatch/erigon/accounts/abi"
	"github.com/ledgerwatch/erigon/accounts/abi/bind"
	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AggregationRouterV3SwapDescription is an auto generated low-level Go binding around an user-defined struct.
type AggregationRouterV3SwapDescription struct {
	SrcToken        libcommon.Address
	DstToken        libcommon.Address
	SrcReceiver     libcommon.Address
	DstReceiver     libcommon.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// OneinchexchangeABI is the input ABI used to generate the binding from.
const OneinchexchangeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structAggregationRouterV3.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"discountedSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chiSpent\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structAggregationRouterV3.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pools\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"unoswapWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Oneinchexchange is an auto generated Go binding around an Ethereum contract.
type Oneinchexchange struct {
	OneinchexchangeCaller     // Read-only binding to the contract
	OneinchexchangeTransactor // Write-only binding to the contract
	OneinchexchangeFilterer   // Log filterer for contract events
}

// OneinchexchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneinchexchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchexchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneinchexchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchexchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneinchexchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchexchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneinchexchangeSession struct {
	Contract     *Oneinchexchange  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneinchexchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneinchexchangeCallerSession struct {
	Contract *OneinchexchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// OneinchexchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneinchexchangeTransactorSession struct {
	Contract     *OneinchexchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OneinchexchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneinchexchangeRaw struct {
	Contract *Oneinchexchange // Generic contract binding to access the raw methods on
}

// OneinchexchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneinchexchangeCallerRaw struct {
	Contract *OneinchexchangeCaller // Generic read-only contract binding to access the raw methods on
}

// OneinchexchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneinchexchangeTransactorRaw struct {
	Contract *OneinchexchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneinchexchange creates a new instance of Oneinchexchange, bound to a specific deployed contract.
func NewOneinchexchange(address libcommon.Address, backend bind.ContractBackend) (*Oneinchexchange, error) {
	contract, err := bindOneinchexchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oneinchexchange{OneinchexchangeCaller: OneinchexchangeCaller{contract: contract}, OneinchexchangeTransactor: OneinchexchangeTransactor{contract: contract}, OneinchexchangeFilterer: OneinchexchangeFilterer{contract: contract}}, nil
}

// NewOneinchexchangeCaller creates a new read-only instance of Oneinchexchange, bound to a specific deployed contract.
func NewOneinchexchangeCaller(address libcommon.Address, caller bind.ContractCaller) (*OneinchexchangeCaller, error) {
	contract, err := bindOneinchexchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneinchexchangeCaller{contract: contract}, nil
}

// NewOneinchexchangeTransactor creates a new write-only instance of Oneinchexchange, bound to a specific deployed contract.
func NewOneinchexchangeTransactor(address libcommon.Address, transactor bind.ContractTransactor) (*OneinchexchangeTransactor, error) {
	contract, err := bindOneinchexchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneinchexchangeTransactor{contract: contract}, nil
}

// NewOneinchexchangeFilterer creates a new log filterer instance of Oneinchexchange, bound to a specific deployed contract.
func NewOneinchexchangeFilterer(address libcommon.Address, filterer bind.ContractFilterer) (*OneinchexchangeFilterer, error) {
	contract, err := bindOneinchexchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneinchexchangeFilterer{contract: contract}, nil
}

// bindOneinchexchange binds a generic wrapper to an already deployed contract.
func bindOneinchexchange(address libcommon.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneinchexchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinchexchange *OneinchexchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oneinchexchange.Contract.OneinchexchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinchexchange *OneinchexchangeRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Oneinchexchange.Contract.OneinchexchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinchexchange *OneinchexchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _Oneinchexchange.Contract.OneinchexchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinchexchange *OneinchexchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oneinchexchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinchexchange *OneinchexchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Oneinchexchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinchexchange *OneinchexchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _Oneinchexchange.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinchexchange *OneinchexchangeCaller) Owner(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _Oneinchexchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinchexchange *OneinchexchangeSession) Owner() (libcommon.Address, error) {
	return _Oneinchexchange.Contract.Owner(&_Oneinchexchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinchexchange *OneinchexchangeCallerSession) Owner() (libcommon.Address, error) {
	return _Oneinchexchange.Contract.Owner(&_Oneinchexchange.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinchexchange *OneinchexchangeTransactor) Destroy(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Oneinchexchange.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinchexchange *OneinchexchangeSession) Destroy() (types.Transaction, error) {
	return _Oneinchexchange.Contract.Destroy(&_Oneinchexchange.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinchexchange *OneinchexchangeTransactorSession) Destroy() (types.Transaction, error) {
	return _Oneinchexchange.Contract.Destroy(&_Oneinchexchange.TransactOpts)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_Oneinchexchange *OneinchexchangeTransactor) DiscountedSwap(opts *bind.TransactOpts, caller libcommon.Address, desc AggregationRouterV3SwapDescription, data []byte) (types.Transaction, error) {
	return _Oneinchexchange.contract.Transact(opts, "discountedSwap", caller, desc, data)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_Oneinchexchange *OneinchexchangeSession) DiscountedSwap(caller libcommon.Address, desc AggregationRouterV3SwapDescription, data []byte) (types.Transaction, error) {
	return _Oneinchexchange.Contract.DiscountedSwap(&_Oneinchexchange.TransactOpts, caller, desc, data)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_Oneinchexchange *OneinchexchangeTransactorSession) DiscountedSwap(caller libcommon.Address, desc AggregationRouterV3SwapDescription, data []byte) (types.Transaction, error) {
	return _Oneinchexchange.Contract.DiscountedSwap(&_Oneinchexchange.TransactOpts, caller, desc, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinchexchange *OneinchexchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Oneinchexchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinchexchange *OneinchexchangeSession) RenounceOwnership() (types.Transaction, error) {
	return _Oneinchexchange.Contract.RenounceOwnership(&_Oneinchexchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinchexchange *OneinchexchangeTransactorSession) RenounceOwnership() (types.Transaction, error) {
	return _Oneinchexchange.Contract.RenounceOwnership(&_Oneinchexchange.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinchexchange *OneinchexchangeTransactor) RescueFunds(opts *bind.TransactOpts, token libcommon.Address, amount *big.Int) (types.Transaction, error) {
	return _Oneinchexchange.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinchexchange *OneinchexchangeSession) RescueFunds(token libcommon.Address, amount *big.Int) (types.Transaction, error) {
	return _Oneinchexchange.Contract.RescueFunds(&_Oneinchexchange.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinchexchange *OneinchexchangeTransactorSession) RescueFunds(token libcommon.Address, amount *big.Int) (types.Transaction, error) {
	return _Oneinchexchange.Contract.RescueFunds(&_Oneinchexchange.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_Oneinchexchange *OneinchexchangeTransactor) Swap(opts *bind.TransactOpts, caller libcommon.Address, desc AggregationRouterV3SwapDescription, data []byte) (types.Transaction, error) {
	return _Oneinchexchange.contract.Transact(opts, "swap", caller, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_Oneinchexchange *OneinchexchangeSession) Swap(caller libcommon.Address, desc AggregationRouterV3SwapDescription, data []byte) (types.Transaction, error) {
	return _Oneinchexchange.Contract.Swap(&_Oneinchexchange.TransactOpts, caller, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_Oneinchexchange *OneinchexchangeTransactorSession) Swap(caller libcommon.Address, desc AggregationRouterV3SwapDescription, data []byte) (types.Transaction, error) {
	return _Oneinchexchange.Contract.Swap(&_Oneinchexchange.TransactOpts, caller, desc, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinchexchange *OneinchexchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner libcommon.Address) (types.Transaction, error) {
	return _Oneinchexchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinchexchange *OneinchexchangeSession) TransferOwnership(newOwner libcommon.Address) (types.Transaction, error) {
	return _Oneinchexchange.Contract.TransferOwnership(&_Oneinchexchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinchexchange *OneinchexchangeTransactorSession) TransferOwnership(newOwner libcommon.Address) (types.Transaction, error) {
	return _Oneinchexchange.Contract.TransferOwnership(&_Oneinchexchange.TransactOpts, newOwner)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeTransactor) Unoswap(opts *bind.TransactOpts, srcToken libcommon.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (types.Transaction, error) {
	return _Oneinchexchange.contract.Transact(opts, "unoswap", srcToken, amount, minReturn, arg3)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeSession) Unoswap(srcToken libcommon.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (types.Transaction, error) {
	return _Oneinchexchange.Contract.Unoswap(&_Oneinchexchange.TransactOpts, srcToken, amount, minReturn, arg3)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeTransactorSession) Unoswap(srcToken libcommon.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (types.Transaction, error) {
	return _Oneinchexchange.Contract.Unoswap(&_Oneinchexchange.TransactOpts, srcToken, amount, minReturn, arg3)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeTransactor) UnoswapWithPermit(opts *bind.TransactOpts, srcToken libcommon.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (types.Transaction, error) {
	return _Oneinchexchange.contract.Transact(opts, "unoswapWithPermit", srcToken, amount, minReturn, pools, permit)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeSession) UnoswapWithPermit(srcToken libcommon.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (types.Transaction, error) {
	return _Oneinchexchange.Contract.UnoswapWithPermit(&_Oneinchexchange.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeTransactorSession) UnoswapWithPermit(srcToken libcommon.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (types.Transaction, error) {
	return _Oneinchexchange.Contract.UnoswapWithPermit(&_Oneinchexchange.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinchexchange *OneinchexchangeTransactor) Receive(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Oneinchexchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinchexchange *OneinchexchangeSession) Receive() (types.Transaction, error) {
	return _Oneinchexchange.Contract.Receive(&_Oneinchexchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinchexchange *OneinchexchangeTransactorSession) Receive() (types.Transaction, error) {
	return _Oneinchexchange.Contract.Receive(&_Oneinchexchange.TransactOpts)
}

// OneinchexchangeErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the Oneinchexchange contract.
type OneinchexchangeErrorIterator struct {
	Event *OneinchexchangeError // Event containing the contract specifics and raw log

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
func (it *OneinchexchangeErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchexchangeError)
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
		it.Event = new(OneinchexchangeError)
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
func (it *OneinchexchangeErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchexchangeErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchexchangeError represents a Error event raised by the Oneinchexchange contract.
type OneinchexchangeError struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Oneinchexchange *OneinchexchangeFilterer) FilterError(opts *bind.FilterOpts) (*OneinchexchangeErrorIterator, error) {

	logs, sub, err := _Oneinchexchange.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &OneinchexchangeErrorIterator{contract: _Oneinchexchange.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Oneinchexchange *OneinchexchangeFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *OneinchexchangeError) (event.Subscription, error) {

	logs, sub, err := _Oneinchexchange.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchexchangeError)
				if err := _Oneinchexchange.contract.UnpackLog(event, "Error", log); err != nil {
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

// ParseError is a log parse operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Oneinchexchange *OneinchexchangeFilterer) ParseError(log types.Log) (*OneinchexchangeError, error) {
	event := new(OneinchexchangeError)
	if err := _Oneinchexchange.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneinchexchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oneinchexchange contract.
type OneinchexchangeOwnershipTransferredIterator struct {
	Event *OneinchexchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneinchexchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchexchangeOwnershipTransferred)
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
		it.Event = new(OneinchexchangeOwnershipTransferred)
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
func (it *OneinchexchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchexchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchexchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Oneinchexchange contract.
type OneinchexchangeOwnershipTransferred struct {
	PreviousOwner libcommon.Address
	NewOwner      libcommon.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinchexchange *OneinchexchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []libcommon.Address, newOwner []libcommon.Address) (*OneinchexchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinchexchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneinchexchangeOwnershipTransferredIterator{contract: _Oneinchexchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinchexchange *OneinchexchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneinchexchangeOwnershipTransferred, previousOwner []libcommon.Address, newOwner []libcommon.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinchexchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchexchangeOwnershipTransferred)
				if err := _Oneinchexchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinchexchange *OneinchexchangeFilterer) ParseOwnershipTransferred(log types.Log) (*OneinchexchangeOwnershipTransferred, error) {
	event := new(OneinchexchangeOwnershipTransferred)
	if err := _Oneinchexchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneinchexchangeSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Oneinchexchange contract.
type OneinchexchangeSwappedIterator struct {
	Event *OneinchexchangeSwapped // Event containing the contract specifics and raw log

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
func (it *OneinchexchangeSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchexchangeSwapped)
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
		it.Event = new(OneinchexchangeSwapped)
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
func (it *OneinchexchangeSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchexchangeSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchexchangeSwapped represents a Swapped event raised by the Oneinchexchange contract.
type OneinchexchangeSwapped struct {
	Sender       libcommon.Address
	SrcToken     libcommon.Address
	DstToken     libcommon.Address
	DstReceiver  libcommon.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeFilterer) FilterSwapped(opts *bind.FilterOpts) (*OneinchexchangeSwappedIterator, error) {

	logs, sub, err := _Oneinchexchange.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &OneinchexchangeSwappedIterator{contract: _Oneinchexchange.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *OneinchexchangeSwapped) (event.Subscription, error) {

	logs, sub, err := _Oneinchexchange.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchexchangeSwapped)
				if err := _Oneinchexchange.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Oneinchexchange *OneinchexchangeFilterer) ParseSwapped(log types.Log) (*OneinchexchangeSwapped, error) {
	event := new(OneinchexchangeSwapped)
	if err := _Oneinchexchange.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
