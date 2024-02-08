// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unistaking_contract

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
)

// UnistakingContractMetaData contains all meta data concerning the UnistakingContract contract.
var UnistakingContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Mintable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unmintable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintableGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"setMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UnistakingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use UnistakingContractMetaData.ABI instead.
var UnistakingContractABI = UnistakingContractMetaData.ABI

// UnistakingContract is an auto generated Go binding around an Ethereum contract.
type UnistakingContract struct {
	UnistakingContractCaller     // Read-only binding to the contract
	UnistakingContractTransactor // Write-only binding to the contract
	UnistakingContractFilterer   // Log filterer for contract events
}

// UnistakingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnistakingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnistakingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnistakingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnistakingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnistakingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnistakingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnistakingContractSession struct {
	Contract     *UnistakingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnistakingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnistakingContractCallerSession struct {
	Contract *UnistakingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UnistakingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnistakingContractTransactorSession struct {
	Contract     *UnistakingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UnistakingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnistakingContractRaw struct {
	Contract *UnistakingContract // Generic contract binding to access the raw methods on
}

// UnistakingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnistakingContractCallerRaw struct {
	Contract *UnistakingContractCaller // Generic read-only contract binding to access the raw methods on
}

// UnistakingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnistakingContractTransactorRaw struct {
	Contract *UnistakingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnistakingContract creates a new instance of UnistakingContract, bound to a specific deployed contract.
func NewUnistakingContract(address common.Address, backend bind.ContractBackend) (*UnistakingContract, error) {
	contract, err := bindUnistakingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnistakingContract{UnistakingContractCaller: UnistakingContractCaller{contract: contract}, UnistakingContractTransactor: UnistakingContractTransactor{contract: contract}, UnistakingContractFilterer: UnistakingContractFilterer{contract: contract}}, nil
}

// NewUnistakingContractCaller creates a new read-only instance of UnistakingContract, bound to a specific deployed contract.
func NewUnistakingContractCaller(address common.Address, caller bind.ContractCaller) (*UnistakingContractCaller, error) {
	contract, err := bindUnistakingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnistakingContractCaller{contract: contract}, nil
}

// NewUnistakingContractTransactor creates a new write-only instance of UnistakingContract, bound to a specific deployed contract.
func NewUnistakingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*UnistakingContractTransactor, error) {
	contract, err := bindUnistakingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnistakingContractTransactor{contract: contract}, nil
}

// NewUnistakingContractFilterer creates a new log filterer instance of UnistakingContract, bound to a specific deployed contract.
func NewUnistakingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*UnistakingContractFilterer, error) {
	contract, err := bindUnistakingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnistakingContractFilterer{contract: contract}, nil
}

// bindUnistakingContract binds a generic wrapper to an already deployed contract.
func bindUnistakingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnistakingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnistakingContract *UnistakingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnistakingContract.Contract.UnistakingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnistakingContract *UnistakingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.Contract.UnistakingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnistakingContract *UnistakingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnistakingContract.Contract.UnistakingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnistakingContract *UnistakingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnistakingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnistakingContract *UnistakingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnistakingContract *UnistakingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnistakingContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UnistakingContract *UnistakingContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UnistakingContract *UnistakingContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UnistakingContract.Contract.Allowance(&_UnistakingContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UnistakingContract *UnistakingContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UnistakingContract.Contract.Allowance(&_UnistakingContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UnistakingContract *UnistakingContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UnistakingContract *UnistakingContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UnistakingContract.Contract.BalanceOf(&_UnistakingContract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UnistakingContract *UnistakingContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UnistakingContract.Contract.BalanceOf(&_UnistakingContract.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_UnistakingContract *UnistakingContractCaller) BalanceOfAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "balanceOfAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_UnistakingContract *UnistakingContractSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _UnistakingContract.Contract.BalanceOfAt(&_UnistakingContract.CallOpts, account, snapshotId)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_UnistakingContract *UnistakingContractCallerSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _UnistakingContract.Contract.BalanceOfAt(&_UnistakingContract.CallOpts, account, snapshotId)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UnistakingContract *UnistakingContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UnistakingContract *UnistakingContractSession) Decimals() (uint8, error) {
	return _UnistakingContract.Contract.Decimals(&_UnistakingContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UnistakingContract *UnistakingContractCallerSession) Decimals() (uint8, error) {
	return _UnistakingContract.Contract.Decimals(&_UnistakingContract.CallOpts)
}

// MintableGroup is a free data retrieval call binding the contract method 0xec49c938.
//
// Solidity: function mintableGroup(address ) view returns(bool)
func (_UnistakingContract *UnistakingContractCaller) MintableGroup(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "mintableGroup", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintableGroup is a free data retrieval call binding the contract method 0xec49c938.
//
// Solidity: function mintableGroup(address ) view returns(bool)
func (_UnistakingContract *UnistakingContractSession) MintableGroup(arg0 common.Address) (bool, error) {
	return _UnistakingContract.Contract.MintableGroup(&_UnistakingContract.CallOpts, arg0)
}

// MintableGroup is a free data retrieval call binding the contract method 0xec49c938.
//
// Solidity: function mintableGroup(address ) view returns(bool)
func (_UnistakingContract *UnistakingContractCallerSession) MintableGroup(arg0 common.Address) (bool, error) {
	return _UnistakingContract.Contract.MintableGroup(&_UnistakingContract.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UnistakingContract *UnistakingContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UnistakingContract *UnistakingContractSession) Name() (string, error) {
	return _UnistakingContract.Contract.Name(&_UnistakingContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UnistakingContract *UnistakingContractCallerSession) Name() (string, error) {
	return _UnistakingContract.Contract.Name(&_UnistakingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnistakingContract *UnistakingContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnistakingContract *UnistakingContractSession) Owner() (common.Address, error) {
	return _UnistakingContract.Contract.Owner(&_UnistakingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnistakingContract *UnistakingContractCallerSession) Owner() (common.Address, error) {
	return _UnistakingContract.Contract.Owner(&_UnistakingContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UnistakingContract *UnistakingContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UnistakingContract *UnistakingContractSession) Paused() (bool, error) {
	return _UnistakingContract.Contract.Paused(&_UnistakingContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UnistakingContract *UnistakingContractCallerSession) Paused() (bool, error) {
	return _UnistakingContract.Contract.Paused(&_UnistakingContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UnistakingContract *UnistakingContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UnistakingContract *UnistakingContractSession) Symbol() (string, error) {
	return _UnistakingContract.Contract.Symbol(&_UnistakingContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UnistakingContract *UnistakingContractCallerSession) Symbol() (string, error) {
	return _UnistakingContract.Contract.Symbol(&_UnistakingContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UnistakingContract *UnistakingContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UnistakingContract *UnistakingContractSession) TotalSupply() (*big.Int, error) {
	return _UnistakingContract.Contract.TotalSupply(&_UnistakingContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UnistakingContract *UnistakingContractCallerSession) TotalSupply() (*big.Int, error) {
	return _UnistakingContract.Contract.TotalSupply(&_UnistakingContract.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_UnistakingContract *UnistakingContractCaller) TotalSupplyAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UnistakingContract.contract.Call(opts, &out, "totalSupplyAt", snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_UnistakingContract *UnistakingContractSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _UnistakingContract.Contract.TotalSupplyAt(&_UnistakingContract.CallOpts, snapshotId)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_UnistakingContract *UnistakingContractCallerSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _UnistakingContract.Contract.TotalSupplyAt(&_UnistakingContract.CallOpts, snapshotId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Approve(&_UnistakingContract.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Approve(&_UnistakingContract.TransactOpts, spender, amount)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] recipients, uint256[] amounts) returns()
func (_UnistakingContract *UnistakingContractTransactor) BatchTransfer(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "batchTransfer", recipients, amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] recipients, uint256[] amounts) returns()
func (_UnistakingContract *UnistakingContractSession) BatchTransfer(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.BatchTransfer(&_UnistakingContract.TransactOpts, recipients, amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] recipients, uint256[] amounts) returns()
func (_UnistakingContract *UnistakingContractTransactorSession) BatchTransfer(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.BatchTransfer(&_UnistakingContract.TransactOpts, recipients, amounts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_UnistakingContract *UnistakingContractTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_UnistakingContract *UnistakingContractSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Burn(&_UnistakingContract.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_UnistakingContract *UnistakingContractTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Burn(&_UnistakingContract.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_UnistakingContract *UnistakingContractTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_UnistakingContract *UnistakingContractSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.BurnFrom(&_UnistakingContract.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_UnistakingContract *UnistakingContractTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.BurnFrom(&_UnistakingContract.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_UnistakingContract *UnistakingContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_UnistakingContract *UnistakingContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.DecreaseAllowance(&_UnistakingContract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_UnistakingContract *UnistakingContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.DecreaseAllowance(&_UnistakingContract.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_UnistakingContract *UnistakingContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_UnistakingContract *UnistakingContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.IncreaseAllowance(&_UnistakingContract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_UnistakingContract *UnistakingContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.IncreaseAllowance(&_UnistakingContract.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_UnistakingContract *UnistakingContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_UnistakingContract *UnistakingContractSession) Initialize() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Initialize(&_UnistakingContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_UnistakingContract *UnistakingContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Initialize(&_UnistakingContract.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_UnistakingContract *UnistakingContractTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_UnistakingContract *UnistakingContractSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Mint(&_UnistakingContract.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_UnistakingContract *UnistakingContractTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Mint(&_UnistakingContract.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UnistakingContract *UnistakingContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UnistakingContract *UnistakingContractSession) Pause() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Pause(&_UnistakingContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UnistakingContract *UnistakingContractTransactorSession) Pause() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Pause(&_UnistakingContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnistakingContract *UnistakingContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnistakingContract *UnistakingContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnistakingContract.Contract.RenounceOwnership(&_UnistakingContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnistakingContract *UnistakingContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnistakingContract.Contract.RenounceOwnership(&_UnistakingContract.TransactOpts)
}

// SetMintable is a paid mutator transaction binding the contract method 0xf7eb06c4.
//
// Solidity: function setMintable(address account, bool allow) returns()
func (_UnistakingContract *UnistakingContractTransactor) SetMintable(opts *bind.TransactOpts, account common.Address, allow bool) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "setMintable", account, allow)
}

// SetMintable is a paid mutator transaction binding the contract method 0xf7eb06c4.
//
// Solidity: function setMintable(address account, bool allow) returns()
func (_UnistakingContract *UnistakingContractSession) SetMintable(account common.Address, allow bool) (*types.Transaction, error) {
	return _UnistakingContract.Contract.SetMintable(&_UnistakingContract.TransactOpts, account, allow)
}

// SetMintable is a paid mutator transaction binding the contract method 0xf7eb06c4.
//
// Solidity: function setMintable(address account, bool allow) returns()
func (_UnistakingContract *UnistakingContractTransactorSession) SetMintable(account common.Address, allow bool) (*types.Transaction, error) {
	return _UnistakingContract.Contract.SetMintable(&_UnistakingContract.TransactOpts, account, allow)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_UnistakingContract *UnistakingContractTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_UnistakingContract *UnistakingContractSession) Snapshot() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Snapshot(&_UnistakingContract.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_UnistakingContract *UnistakingContractTransactorSession) Snapshot() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Snapshot(&_UnistakingContract.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Transfer(&_UnistakingContract.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Transfer(&_UnistakingContract.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.TransferFrom(&_UnistakingContract.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_UnistakingContract *UnistakingContractTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnistakingContract.Contract.TransferFrom(&_UnistakingContract.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnistakingContract *UnistakingContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnistakingContract *UnistakingContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnistakingContract.Contract.TransferOwnership(&_UnistakingContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnistakingContract *UnistakingContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnistakingContract.Contract.TransferOwnership(&_UnistakingContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UnistakingContract *UnistakingContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UnistakingContract *UnistakingContractSession) Unpause() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Unpause(&_UnistakingContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UnistakingContract *UnistakingContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Unpause(&_UnistakingContract.TransactOpts)
}

// UnistakingContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UnistakingContract contract.
type UnistakingContractApprovalIterator struct {
	Event *UnistakingContractApproval // Event containing the contract specifics and raw log

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
func (it *UnistakingContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractApproval)
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
		it.Event = new(UnistakingContractApproval)
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
func (it *UnistakingContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractApproval represents a Approval event raised by the UnistakingContract contract.
type UnistakingContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UnistakingContract *UnistakingContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UnistakingContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UnistakingContractApprovalIterator{contract: _UnistakingContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UnistakingContract *UnistakingContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UnistakingContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractApproval)
				if err := _UnistakingContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UnistakingContract *UnistakingContractFilterer) ParseApproval(log types.Log) (*UnistakingContractApproval, error) {
	event := new(UnistakingContractApproval)
	if err := _UnistakingContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractMintableIterator is returned from FilterMintable and is used to iterate over the raw logs and unpacked data for Mintable events raised by the UnistakingContract contract.
type UnistakingContractMintableIterator struct {
	Event *UnistakingContractMintable // Event containing the contract specifics and raw log

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
func (it *UnistakingContractMintableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractMintable)
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
		it.Event = new(UnistakingContractMintable)
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
func (it *UnistakingContractMintableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractMintableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractMintable represents a Mintable event raised by the UnistakingContract contract.
type UnistakingContractMintable struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintable is a free log retrieval operation binding the contract event 0x8fc6bf72ae36e2e0d306c3344626c34ed9bbc8c82fe8289b2f79847094637551.
//
// Solidity: event Mintable(address account)
func (_UnistakingContract *UnistakingContractFilterer) FilterMintable(opts *bind.FilterOpts) (*UnistakingContractMintableIterator, error) {

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "Mintable")
	if err != nil {
		return nil, err
	}
	return &UnistakingContractMintableIterator{contract: _UnistakingContract.contract, event: "Mintable", logs: logs, sub: sub}, nil
}

// WatchMintable is a free log subscription operation binding the contract event 0x8fc6bf72ae36e2e0d306c3344626c34ed9bbc8c82fe8289b2f79847094637551.
//
// Solidity: event Mintable(address account)
func (_UnistakingContract *UnistakingContractFilterer) WatchMintable(opts *bind.WatchOpts, sink chan<- *UnistakingContractMintable) (event.Subscription, error) {

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "Mintable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractMintable)
				if err := _UnistakingContract.contract.UnpackLog(event, "Mintable", log); err != nil {
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

// ParseMintable is a log parse operation binding the contract event 0x8fc6bf72ae36e2e0d306c3344626c34ed9bbc8c82fe8289b2f79847094637551.
//
// Solidity: event Mintable(address account)
func (_UnistakingContract *UnistakingContractFilterer) ParseMintable(log types.Log) (*UnistakingContractMintable, error) {
	event := new(UnistakingContractMintable)
	if err := _UnistakingContract.contract.UnpackLog(event, "Mintable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UnistakingContract contract.
type UnistakingContractOwnershipTransferredIterator struct {
	Event *UnistakingContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UnistakingContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractOwnershipTransferred)
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
		it.Event = new(UnistakingContractOwnershipTransferred)
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
func (it *UnistakingContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractOwnershipTransferred represents a OwnershipTransferred event raised by the UnistakingContract contract.
type UnistakingContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnistakingContract *UnistakingContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnistakingContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnistakingContractOwnershipTransferredIterator{contract: _UnistakingContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnistakingContract *UnistakingContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnistakingContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractOwnershipTransferred)
				if err := _UnistakingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UnistakingContract *UnistakingContractFilterer) ParseOwnershipTransferred(log types.Log) (*UnistakingContractOwnershipTransferred, error) {
	event := new(UnistakingContractOwnershipTransferred)
	if err := _UnistakingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the UnistakingContract contract.
type UnistakingContractPausedIterator struct {
	Event *UnistakingContractPaused // Event containing the contract specifics and raw log

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
func (it *UnistakingContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractPaused)
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
		it.Event = new(UnistakingContractPaused)
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
func (it *UnistakingContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractPaused represents a Paused event raised by the UnistakingContract contract.
type UnistakingContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UnistakingContract *UnistakingContractFilterer) FilterPaused(opts *bind.FilterOpts) (*UnistakingContractPausedIterator, error) {

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &UnistakingContractPausedIterator{contract: _UnistakingContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UnistakingContract *UnistakingContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *UnistakingContractPaused) (event.Subscription, error) {

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractPaused)
				if err := _UnistakingContract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UnistakingContract *UnistakingContractFilterer) ParsePaused(log types.Log) (*UnistakingContractPaused, error) {
	event := new(UnistakingContractPaused)
	if err := _UnistakingContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the UnistakingContract contract.
type UnistakingContractSnapshotIterator struct {
	Event *UnistakingContractSnapshot // Event containing the contract specifics and raw log

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
func (it *UnistakingContractSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractSnapshot)
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
		it.Event = new(UnistakingContractSnapshot)
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
func (it *UnistakingContractSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractSnapshot represents a Snapshot event raised by the UnistakingContract contract.
type UnistakingContractSnapshot struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_UnistakingContract *UnistakingContractFilterer) FilterSnapshot(opts *bind.FilterOpts) (*UnistakingContractSnapshotIterator, error) {

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &UnistakingContractSnapshotIterator{contract: _UnistakingContract.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_UnistakingContract *UnistakingContractFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *UnistakingContractSnapshot) (event.Subscription, error) {

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractSnapshot)
				if err := _UnistakingContract.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_UnistakingContract *UnistakingContractFilterer) ParseSnapshot(log types.Log) (*UnistakingContractSnapshot, error) {
	event := new(UnistakingContractSnapshot)
	if err := _UnistakingContract.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UnistakingContract contract.
type UnistakingContractTransferIterator struct {
	Event *UnistakingContractTransfer // Event containing the contract specifics and raw log

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
func (it *UnistakingContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractTransfer)
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
		it.Event = new(UnistakingContractTransfer)
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
func (it *UnistakingContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractTransfer represents a Transfer event raised by the UnistakingContract contract.
type UnistakingContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UnistakingContract *UnistakingContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UnistakingContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnistakingContractTransferIterator{contract: _UnistakingContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UnistakingContract *UnistakingContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UnistakingContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractTransfer)
				if err := _UnistakingContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UnistakingContract *UnistakingContractFilterer) ParseTransfer(log types.Log) (*UnistakingContractTransfer, error) {
	event := new(UnistakingContractTransfer)
	if err := _UnistakingContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractUnmintableIterator is returned from FilterUnmintable and is used to iterate over the raw logs and unpacked data for Unmintable events raised by the UnistakingContract contract.
type UnistakingContractUnmintableIterator struct {
	Event *UnistakingContractUnmintable // Event containing the contract specifics and raw log

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
func (it *UnistakingContractUnmintableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractUnmintable)
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
		it.Event = new(UnistakingContractUnmintable)
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
func (it *UnistakingContractUnmintableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractUnmintableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractUnmintable represents a Unmintable event raised by the UnistakingContract contract.
type UnistakingContractUnmintable struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnmintable is a free log retrieval operation binding the contract event 0xb4fa1ea6b9ba3d53e901dbf2b924fff22ac6e0658eecffa8dbd101ff61fc839d.
//
// Solidity: event Unmintable(address account)
func (_UnistakingContract *UnistakingContractFilterer) FilterUnmintable(opts *bind.FilterOpts) (*UnistakingContractUnmintableIterator, error) {

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "Unmintable")
	if err != nil {
		return nil, err
	}
	return &UnistakingContractUnmintableIterator{contract: _UnistakingContract.contract, event: "Unmintable", logs: logs, sub: sub}, nil
}

// WatchUnmintable is a free log subscription operation binding the contract event 0xb4fa1ea6b9ba3d53e901dbf2b924fff22ac6e0658eecffa8dbd101ff61fc839d.
//
// Solidity: event Unmintable(address account)
func (_UnistakingContract *UnistakingContractFilterer) WatchUnmintable(opts *bind.WatchOpts, sink chan<- *UnistakingContractUnmintable) (event.Subscription, error) {

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "Unmintable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractUnmintable)
				if err := _UnistakingContract.contract.UnpackLog(event, "Unmintable", log); err != nil {
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

// ParseUnmintable is a log parse operation binding the contract event 0xb4fa1ea6b9ba3d53e901dbf2b924fff22ac6e0658eecffa8dbd101ff61fc839d.
//
// Solidity: event Unmintable(address account)
func (_UnistakingContract *UnistakingContractFilterer) ParseUnmintable(log types.Log) (*UnistakingContractUnmintable, error) {
	event := new(UnistakingContractUnmintable)
	if err := _UnistakingContract.contract.UnpackLog(event, "Unmintable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the UnistakingContract contract.
type UnistakingContractUnpausedIterator struct {
	Event *UnistakingContractUnpaused // Event containing the contract specifics and raw log

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
func (it *UnistakingContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractUnpaused)
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
		it.Event = new(UnistakingContractUnpaused)
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
func (it *UnistakingContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractUnpaused represents a Unpaused event raised by the UnistakingContract contract.
type UnistakingContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UnistakingContract *UnistakingContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*UnistakingContractUnpausedIterator, error) {

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &UnistakingContractUnpausedIterator{contract: _UnistakingContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UnistakingContract *UnistakingContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *UnistakingContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractUnpaused)
				if err := _UnistakingContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UnistakingContract *UnistakingContractFilterer) ParseUnpaused(log types.Log) (*UnistakingContractUnpaused, error) {
	event := new(UnistakingContractUnpaused)
	if err := _UnistakingContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
