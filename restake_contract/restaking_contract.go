// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package restaking_contract

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

// BeaconChainProofsStateRootProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

// RestakingContractMetaData contains all meta data concerning the RestakingContract contract.
var RestakingContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"podId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callDelegationManager\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxNumberOfWithdrawalsToClaim\",\"type\":\"uint256\"}],\"name\":\"claimDelayedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createPod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedWithdrawalRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPod\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingWithdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getPod\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"initializeV3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"podOwners\",\"outputs\":[{\"internalType\":\"contractIPodOwner\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"upgradeBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"podId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"oracleTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"name\":\"stateRootProof\",\"type\":\"tuple\"},{\"internalType\":\"uint40[]\",\"name\":\"validatorIndices\",\"type\":\"uint40[]\"},{\"internalType\":\"bytes[]\",\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"validatorFields\",\"type\":\"bytes32[][]\"}],\"name\":\"verifyWithdrawalCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawBeforeRestaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RestakingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RestakingContractMetaData.ABI instead.
var RestakingContractABI = RestakingContractMetaData.ABI

// RestakingContract is an auto generated Go binding around an Ethereum contract.
type RestakingContract struct {
	RestakingContractCaller     // Read-only binding to the contract
	RestakingContractTransactor // Write-only binding to the contract
	RestakingContractFilterer   // Log filterer for contract events
}

// RestakingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestakingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestakingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestakingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestakingContractSession struct {
	Contract     *RestakingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RestakingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestakingContractCallerSession struct {
	Contract *RestakingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RestakingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestakingContractTransactorSession struct {
	Contract     *RestakingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RestakingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestakingContractRaw struct {
	Contract *RestakingContract // Generic contract binding to access the raw methods on
}

// RestakingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestakingContractCallerRaw struct {
	Contract *RestakingContractCaller // Generic read-only contract binding to access the raw methods on
}

// RestakingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestakingContractTransactorRaw struct {
	Contract *RestakingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestakingContract creates a new instance of RestakingContract, bound to a specific deployed contract.
func NewRestakingContract(address common.Address, backend bind.ContractBackend) (*RestakingContract, error) {
	contract, err := bindRestakingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestakingContract{RestakingContractCaller: RestakingContractCaller{contract: contract}, RestakingContractTransactor: RestakingContractTransactor{contract: contract}, RestakingContractFilterer: RestakingContractFilterer{contract: contract}}, nil
}

// NewRestakingContractCaller creates a new read-only instance of RestakingContract, bound to a specific deployed contract.
func NewRestakingContractCaller(address common.Address, caller bind.ContractCaller) (*RestakingContractCaller, error) {
	contract, err := bindRestakingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestakingContractCaller{contract: contract}, nil
}

// NewRestakingContractTransactor creates a new write-only instance of RestakingContract, bound to a specific deployed contract.
func NewRestakingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RestakingContractTransactor, error) {
	contract, err := bindRestakingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestakingContractTransactor{contract: contract}, nil
}

// NewRestakingContractFilterer creates a new log filterer instance of RestakingContract, bound to a specific deployed contract.
func NewRestakingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RestakingContractFilterer, error) {
	contract, err := bindRestakingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestakingContractFilterer{contract: contract}, nil
}

// bindRestakingContract binds a generic wrapper to an already deployed contract.
func bindRestakingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestakingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestakingContract *RestakingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RestakingContract.Contract.RestakingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestakingContract *RestakingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestakingContract.Contract.RestakingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestakingContract *RestakingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestakingContract.Contract.RestakingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestakingContract *RestakingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RestakingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestakingContract *RestakingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestakingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestakingContract *RestakingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestakingContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RestakingContract *RestakingContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RestakingContract *RestakingContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RestakingContract.Contract.DEFAULTADMINROLE(&_RestakingContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RestakingContract *RestakingContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RestakingContract.Contract.DEFAULTADMINROLE(&_RestakingContract.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_RestakingContract *RestakingContractCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_RestakingContract *RestakingContractSession) OPERATORROLE() ([32]byte, error) {
	return _RestakingContract.Contract.OPERATORROLE(&_RestakingContract.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_RestakingContract *RestakingContractCallerSession) OPERATORROLE() ([32]byte, error) {
	return _RestakingContract.Contract.OPERATORROLE(&_RestakingContract.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_RestakingContract *RestakingContractCaller) Beacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "beacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_RestakingContract *RestakingContractSession) Beacon() (common.Address, error) {
	return _RestakingContract.Contract.Beacon(&_RestakingContract.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_RestakingContract *RestakingContractCallerSession) Beacon() (common.Address, error) {
	return _RestakingContract.Contract.Beacon(&_RestakingContract.CallOpts)
}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_RestakingContract *RestakingContractCaller) DelayedWithdrawalRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "delayedWithdrawalRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_RestakingContract *RestakingContractSession) DelayedWithdrawalRouter() (common.Address, error) {
	return _RestakingContract.Contract.DelayedWithdrawalRouter(&_RestakingContract.CallOpts)
}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_RestakingContract *RestakingContractCallerSession) DelayedWithdrawalRouter() (common.Address, error) {
	return _RestakingContract.Contract.DelayedWithdrawalRouter(&_RestakingContract.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_RestakingContract *RestakingContractCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_RestakingContract *RestakingContractSession) DelegationManager() (common.Address, error) {
	return _RestakingContract.Contract.DelegationManager(&_RestakingContract.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_RestakingContract *RestakingContractCallerSession) DelegationManager() (common.Address, error) {
	return _RestakingContract.Contract.DelegationManager(&_RestakingContract.CallOpts)
}

// EigenPod is a free data retrieval call binding the contract method 0xa3aae136.
//
// Solidity: function eigenPod() view returns(address)
func (_RestakingContract *RestakingContractCaller) EigenPod(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "eigenPod")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPod is a free data retrieval call binding the contract method 0xa3aae136.
//
// Solidity: function eigenPod() view returns(address)
func (_RestakingContract *RestakingContractSession) EigenPod() (common.Address, error) {
	return _RestakingContract.Contract.EigenPod(&_RestakingContract.CallOpts)
}

// EigenPod is a free data retrieval call binding the contract method 0xa3aae136.
//
// Solidity: function eigenPod() view returns(address)
func (_RestakingContract *RestakingContractCallerSession) EigenPod() (common.Address, error) {
	return _RestakingContract.Contract.EigenPod(&_RestakingContract.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_RestakingContract *RestakingContractCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_RestakingContract *RestakingContractSession) EigenPodManager() (common.Address, error) {
	return _RestakingContract.Contract.EigenPodManager(&_RestakingContract.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_RestakingContract *RestakingContractCallerSession) EigenPodManager() (common.Address, error) {
	return _RestakingContract.Contract.EigenPodManager(&_RestakingContract.CallOpts)
}

// GetPendingWithdrawalAmount is a free data retrieval call binding the contract method 0x99cba074.
//
// Solidity: function getPendingWithdrawalAmount() view returns(uint256)
func (_RestakingContract *RestakingContractCaller) GetPendingWithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "getPendingWithdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingWithdrawalAmount is a free data retrieval call binding the contract method 0x99cba074.
//
// Solidity: function getPendingWithdrawalAmount() view returns(uint256)
func (_RestakingContract *RestakingContractSession) GetPendingWithdrawalAmount() (*big.Int, error) {
	return _RestakingContract.Contract.GetPendingWithdrawalAmount(&_RestakingContract.CallOpts)
}

// GetPendingWithdrawalAmount is a free data retrieval call binding the contract method 0x99cba074.
//
// Solidity: function getPendingWithdrawalAmount() view returns(uint256)
func (_RestakingContract *RestakingContractCallerSession) GetPendingWithdrawalAmount() (*big.Int, error) {
	return _RestakingContract.Contract.GetPendingWithdrawalAmount(&_RestakingContract.CallOpts)
}

// GetPod is a free data retrieval call binding the contract method 0x6a93ee0f.
//
// Solidity: function getPod(uint256 i) view returns(address)
func (_RestakingContract *RestakingContractCaller) GetPod(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "getPod", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPod is a free data retrieval call binding the contract method 0x6a93ee0f.
//
// Solidity: function getPod(uint256 i) view returns(address)
func (_RestakingContract *RestakingContractSession) GetPod(i *big.Int) (common.Address, error) {
	return _RestakingContract.Contract.GetPod(&_RestakingContract.CallOpts, i)
}

// GetPod is a free data retrieval call binding the contract method 0x6a93ee0f.
//
// Solidity: function getPod(uint256 i) view returns(address)
func (_RestakingContract *RestakingContractCallerSession) GetPod(i *big.Int) (common.Address, error) {
	return _RestakingContract.Contract.GetPod(&_RestakingContract.CallOpts, i)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RestakingContract *RestakingContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RestakingContract *RestakingContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RestakingContract.Contract.GetRoleAdmin(&_RestakingContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RestakingContract *RestakingContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RestakingContract.Contract.GetRoleAdmin(&_RestakingContract.CallOpts, role)
}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_RestakingContract *RestakingContractCaller) GetTotalPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "getTotalPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_RestakingContract *RestakingContractSession) GetTotalPods() (*big.Int, error) {
	return _RestakingContract.Contract.GetTotalPods(&_RestakingContract.CallOpts)
}

// GetTotalPods is a free data retrieval call binding the contract method 0xd2197bd3.
//
// Solidity: function getTotalPods() view returns(uint256)
func (_RestakingContract *RestakingContractCallerSession) GetTotalPods() (*big.Int, error) {
	return _RestakingContract.Contract.GetTotalPods(&_RestakingContract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RestakingContract *RestakingContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RestakingContract *RestakingContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RestakingContract.Contract.HasRole(&_RestakingContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RestakingContract *RestakingContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RestakingContract.Contract.HasRole(&_RestakingContract.CallOpts, role, account)
}

// PodOwners is a free data retrieval call binding the contract method 0xa1dccd11.
//
// Solidity: function podOwners(uint256 ) view returns(address)
func (_RestakingContract *RestakingContractCaller) PodOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "podOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PodOwners is a free data retrieval call binding the contract method 0xa1dccd11.
//
// Solidity: function podOwners(uint256 ) view returns(address)
func (_RestakingContract *RestakingContractSession) PodOwners(arg0 *big.Int) (common.Address, error) {
	return _RestakingContract.Contract.PodOwners(&_RestakingContract.CallOpts, arg0)
}

// PodOwners is a free data retrieval call binding the contract method 0xa1dccd11.
//
// Solidity: function podOwners(uint256 ) view returns(address)
func (_RestakingContract *RestakingContractCallerSession) PodOwners(arg0 *big.Int) (common.Address, error) {
	return _RestakingContract.Contract.PodOwners(&_RestakingContract.CallOpts, arg0)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_RestakingContract *RestakingContractCaller) StakingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "stakingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_RestakingContract *RestakingContractSession) StakingAddress() (common.Address, error) {
	return _RestakingContract.Contract.StakingAddress(&_RestakingContract.CallOpts)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_RestakingContract *RestakingContractCallerSession) StakingAddress() (common.Address, error) {
	return _RestakingContract.Contract.StakingAddress(&_RestakingContract.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_RestakingContract *RestakingContractCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_RestakingContract *RestakingContractSession) StrategyManager() (common.Address, error) {
	return _RestakingContract.Contract.StrategyManager(&_RestakingContract.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_RestakingContract *RestakingContractCallerSession) StrategyManager() (common.Address, error) {
	return _RestakingContract.Contract.StrategyManager(&_RestakingContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RestakingContract *RestakingContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RestakingContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RestakingContract *RestakingContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RestakingContract.Contract.SupportsInterface(&_RestakingContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RestakingContract *RestakingContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RestakingContract.Contract.SupportsInterface(&_RestakingContract.CallOpts, interfaceId)
}

// CallDelegationManager is a paid mutator transaction binding the contract method 0x37a26eef.
//
// Solidity: function callDelegationManager(uint256 podId, bytes data) returns(bytes)
func (_RestakingContract *RestakingContractTransactor) CallDelegationManager(opts *bind.TransactOpts, podId *big.Int, data []byte) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "callDelegationManager", podId, data)
}

// CallDelegationManager is a paid mutator transaction binding the contract method 0x37a26eef.
//
// Solidity: function callDelegationManager(uint256 podId, bytes data) returns(bytes)
func (_RestakingContract *RestakingContractSession) CallDelegationManager(podId *big.Int, data []byte) (*types.Transaction, error) {
	return _RestakingContract.Contract.CallDelegationManager(&_RestakingContract.TransactOpts, podId, data)
}

// CallDelegationManager is a paid mutator transaction binding the contract method 0x37a26eef.
//
// Solidity: function callDelegationManager(uint256 podId, bytes data) returns(bytes)
func (_RestakingContract *RestakingContractTransactorSession) CallDelegationManager(podId *big.Int, data []byte) (*types.Transaction, error) {
	return _RestakingContract.Contract.CallDelegationManager(&_RestakingContract.TransactOpts, podId, data)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_RestakingContract *RestakingContractTransactor) ClaimDelayedWithdrawals(opts *bind.TransactOpts, maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "claimDelayedWithdrawals", maxNumberOfWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_RestakingContract *RestakingContractSession) ClaimDelayedWithdrawals(maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _RestakingContract.Contract.ClaimDelayedWithdrawals(&_RestakingContract.TransactOpts, maxNumberOfWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_RestakingContract *RestakingContractTransactorSession) ClaimDelayedWithdrawals(maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _RestakingContract.Contract.ClaimDelayedWithdrawals(&_RestakingContract.TransactOpts, maxNumberOfWithdrawalsToClaim)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns()
func (_RestakingContract *RestakingContractTransactor) CreatePod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "createPod")
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns()
func (_RestakingContract *RestakingContractSession) CreatePod() (*types.Transaction, error) {
	return _RestakingContract.Contract.CreatePod(&_RestakingContract.TransactOpts)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns()
func (_RestakingContract *RestakingContractTransactorSession) CreatePod() (*types.Transaction, error) {
	return _RestakingContract.Contract.CreatePod(&_RestakingContract.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) returns(bytes)
func (_RestakingContract *RestakingContractTransactor) Execute(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "execute", target, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) returns(bytes)
func (_RestakingContract *RestakingContractSession) Execute(target common.Address, data []byte) (*types.Transaction, error) {
	return _RestakingContract.Contract.Execute(&_RestakingContract.TransactOpts, target, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) returns(bytes)
func (_RestakingContract *RestakingContractTransactorSession) Execute(target common.Address, data []byte) (*types.Transaction, error) {
	return _RestakingContract.Contract.Execute(&_RestakingContract.TransactOpts, target, data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.GrantRole(&_RestakingContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.GrantRole(&_RestakingContract.TransactOpts, role, account)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address impl) returns()
func (_RestakingContract *RestakingContractTransactor) InitializeV3(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "initializeV3", impl)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address impl) returns()
func (_RestakingContract *RestakingContractSession) InitializeV3(impl common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.InitializeV3(&_RestakingContract.TransactOpts, impl)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address impl) returns()
func (_RestakingContract *RestakingContractTransactorSession) InitializeV3(impl common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.InitializeV3(&_RestakingContract.TransactOpts, impl)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.RenounceRole(&_RestakingContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.RenounceRole(&_RestakingContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.RevokeRole(&_RestakingContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RestakingContract *RestakingContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.RevokeRole(&_RestakingContract.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns()
func (_RestakingContract *RestakingContractTransactor) Transfer(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "transfer", target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns()
func (_RestakingContract *RestakingContractSession) Transfer(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RestakingContract.Contract.Transfer(&_RestakingContract.TransactOpts, target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns()
func (_RestakingContract *RestakingContractTransactorSession) Transfer(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RestakingContract.Contract.Transfer(&_RestakingContract.TransactOpts, target, amount)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_RestakingContract *RestakingContractTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_RestakingContract *RestakingContractSession) Update() (*types.Transaction, error) {
	return _RestakingContract.Contract.Update(&_RestakingContract.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_RestakingContract *RestakingContractTransactorSession) Update() (*types.Transaction, error) {
	return _RestakingContract.Contract.Update(&_RestakingContract.TransactOpts)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x1bce4583.
//
// Solidity: function upgradeBeacon(address impl) returns()
func (_RestakingContract *RestakingContractTransactor) UpgradeBeacon(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "upgradeBeacon", impl)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x1bce4583.
//
// Solidity: function upgradeBeacon(address impl) returns()
func (_RestakingContract *RestakingContractSession) UpgradeBeacon(impl common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.UpgradeBeacon(&_RestakingContract.TransactOpts, impl)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x1bce4583.
//
// Solidity: function upgradeBeacon(address impl) returns()
func (_RestakingContract *RestakingContractTransactorSession) UpgradeBeacon(impl common.Address) (*types.Transaction, error) {
	return _RestakingContract.Contract.UpgradeBeacon(&_RestakingContract.TransactOpts, impl)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc125d9d7.
//
// Solidity: function verifyWithdrawalCredentials(uint256 podId, uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_RestakingContract *RestakingContractTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, podId *big.Int, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "verifyWithdrawalCredentials", podId, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc125d9d7.
//
// Solidity: function verifyWithdrawalCredentials(uint256 podId, uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_RestakingContract *RestakingContractSession) VerifyWithdrawalCredentials(podId *big.Int, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _RestakingContract.Contract.VerifyWithdrawalCredentials(&_RestakingContract.TransactOpts, podId, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc125d9d7.
//
// Solidity: function verifyWithdrawalCredentials(uint256 podId, uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_RestakingContract *RestakingContractTransactorSession) VerifyWithdrawalCredentials(podId *big.Int, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _RestakingContract.Contract.VerifyWithdrawalCredentials(&_RestakingContract.TransactOpts, podId, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_RestakingContract *RestakingContractTransactor) WithdrawBeforeRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestakingContract.contract.Transact(opts, "withdrawBeforeRestaking")
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_RestakingContract *RestakingContractSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _RestakingContract.Contract.WithdrawBeforeRestaking(&_RestakingContract.TransactOpts)
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_RestakingContract *RestakingContractTransactorSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _RestakingContract.Contract.WithdrawBeforeRestaking(&_RestakingContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RestakingContract *RestakingContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestakingContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RestakingContract *RestakingContractSession) Receive() (*types.Transaction, error) {
	return _RestakingContract.Contract.Receive(&_RestakingContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RestakingContract *RestakingContractTransactorSession) Receive() (*types.Transaction, error) {
	return _RestakingContract.Contract.Receive(&_RestakingContract.TransactOpts)
}

// RestakingContractClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the RestakingContract contract.
type RestakingContractClaimedIterator struct {
	Event *RestakingContractClaimed // Event containing the contract specifics and raw log

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
func (it *RestakingContractClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingContractClaimed)
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
		it.Event = new(RestakingContractClaimed)
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
func (it *RestakingContractClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingContractClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingContractClaimed represents a Claimed event raised by the RestakingContract contract.
type RestakingContractClaimed struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 amount)
func (_RestakingContract *RestakingContractFilterer) FilterClaimed(opts *bind.FilterOpts) (*RestakingContractClaimedIterator, error) {

	logs, sub, err := _RestakingContract.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &RestakingContractClaimedIterator{contract: _RestakingContract.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 amount)
func (_RestakingContract *RestakingContractFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *RestakingContractClaimed) (event.Subscription, error) {

	logs, sub, err := _RestakingContract.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingContractClaimed)
				if err := _RestakingContract.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 amount)
func (_RestakingContract *RestakingContractFilterer) ParseClaimed(log types.Log) (*RestakingContractClaimed, error) {
	event := new(RestakingContractClaimed)
	if err := _RestakingContract.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RestakingContract contract.
type RestakingContractInitializedIterator struct {
	Event *RestakingContractInitialized // Event containing the contract specifics and raw log

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
func (it *RestakingContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingContractInitialized)
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
		it.Event = new(RestakingContractInitialized)
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
func (it *RestakingContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingContractInitialized represents a Initialized event raised by the RestakingContract contract.
type RestakingContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RestakingContract *RestakingContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*RestakingContractInitializedIterator, error) {

	logs, sub, err := _RestakingContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RestakingContractInitializedIterator{contract: _RestakingContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RestakingContract *RestakingContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RestakingContractInitialized) (event.Subscription, error) {

	logs, sub, err := _RestakingContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingContractInitialized)
				if err := _RestakingContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RestakingContract *RestakingContractFilterer) ParseInitialized(log types.Log) (*RestakingContractInitialized, error) {
	event := new(RestakingContractInitialized)
	if err := _RestakingContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingContractPendingIterator is returned from FilterPending and is used to iterate over the raw logs and unpacked data for Pending events raised by the RestakingContract contract.
type RestakingContractPendingIterator struct {
	Event *RestakingContractPending // Event containing the contract specifics and raw log

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
func (it *RestakingContractPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingContractPending)
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
		it.Event = new(RestakingContractPending)
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
func (it *RestakingContractPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingContractPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingContractPending represents a Pending event raised by the RestakingContract contract.
type RestakingContractPending struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPending is a free log retrieval operation binding the contract event 0x4e2c5ab35c5a6c864cb37f0d4b556e39e83aa85e744306ad9fd3ad7cc933028b.
//
// Solidity: event Pending(uint256 amount)
func (_RestakingContract *RestakingContractFilterer) FilterPending(opts *bind.FilterOpts) (*RestakingContractPendingIterator, error) {

	logs, sub, err := _RestakingContract.contract.FilterLogs(opts, "Pending")
	if err != nil {
		return nil, err
	}
	return &RestakingContractPendingIterator{contract: _RestakingContract.contract, event: "Pending", logs: logs, sub: sub}, nil
}

// WatchPending is a free log subscription operation binding the contract event 0x4e2c5ab35c5a6c864cb37f0d4b556e39e83aa85e744306ad9fd3ad7cc933028b.
//
// Solidity: event Pending(uint256 amount)
func (_RestakingContract *RestakingContractFilterer) WatchPending(opts *bind.WatchOpts, sink chan<- *RestakingContractPending) (event.Subscription, error) {

	logs, sub, err := _RestakingContract.contract.WatchLogs(opts, "Pending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingContractPending)
				if err := _RestakingContract.contract.UnpackLog(event, "Pending", log); err != nil {
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

// ParsePending is a log parse operation binding the contract event 0x4e2c5ab35c5a6c864cb37f0d4b556e39e83aa85e744306ad9fd3ad7cc933028b.
//
// Solidity: event Pending(uint256 amount)
func (_RestakingContract *RestakingContractFilterer) ParsePending(log types.Log) (*RestakingContractPending, error) {
	event := new(RestakingContractPending)
	if err := _RestakingContract.contract.UnpackLog(event, "Pending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RestakingContract contract.
type RestakingContractRoleAdminChangedIterator struct {
	Event *RestakingContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RestakingContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingContractRoleAdminChanged)
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
		it.Event = new(RestakingContractRoleAdminChanged)
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
func (it *RestakingContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingContractRoleAdminChanged represents a RoleAdminChanged event raised by the RestakingContract contract.
type RestakingContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RestakingContract *RestakingContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RestakingContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RestakingContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RestakingContractRoleAdminChangedIterator{contract: _RestakingContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RestakingContract *RestakingContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RestakingContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RestakingContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingContractRoleAdminChanged)
				if err := _RestakingContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RestakingContract *RestakingContractFilterer) ParseRoleAdminChanged(log types.Log) (*RestakingContractRoleAdminChanged, error) {
	event := new(RestakingContractRoleAdminChanged)
	if err := _RestakingContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RestakingContract contract.
type RestakingContractRoleGrantedIterator struct {
	Event *RestakingContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *RestakingContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingContractRoleGranted)
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
		it.Event = new(RestakingContractRoleGranted)
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
func (it *RestakingContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingContractRoleGranted represents a RoleGranted event raised by the RestakingContract contract.
type RestakingContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RestakingContract *RestakingContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RestakingContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RestakingContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RestakingContractRoleGrantedIterator{contract: _RestakingContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RestakingContract *RestakingContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RestakingContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RestakingContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingContractRoleGranted)
				if err := _RestakingContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RestakingContract *RestakingContractFilterer) ParseRoleGranted(log types.Log) (*RestakingContractRoleGranted, error) {
	event := new(RestakingContractRoleGranted)
	if err := _RestakingContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RestakingContract contract.
type RestakingContractRoleRevokedIterator struct {
	Event *RestakingContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RestakingContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingContractRoleRevoked)
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
		it.Event = new(RestakingContractRoleRevoked)
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
func (it *RestakingContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingContractRoleRevoked represents a RoleRevoked event raised by the RestakingContract contract.
type RestakingContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RestakingContract *RestakingContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RestakingContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RestakingContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RestakingContractRoleRevokedIterator{contract: _RestakingContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RestakingContract *RestakingContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RestakingContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RestakingContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingContractRoleRevoked)
				if err := _RestakingContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RestakingContract *RestakingContractFilterer) ParseRoleRevoked(log types.Log) (*RestakingContractRoleRevoked, error) {
	event := new(RestakingContractRoleRevoked)
	if err := _RestakingContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
