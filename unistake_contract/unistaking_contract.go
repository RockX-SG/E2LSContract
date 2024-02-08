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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_UnistakingContract *UnistakingContractTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_UnistakingContract *UnistakingContractSession) Admin() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Admin(&_UnistakingContract.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_UnistakingContract *UnistakingContractTransactorSession) Admin() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Admin(&_UnistakingContract.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UnistakingContract *UnistakingContractTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UnistakingContract *UnistakingContractSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _UnistakingContract.Contract.ChangeAdmin(&_UnistakingContract.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UnistakingContract *UnistakingContractTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _UnistakingContract.Contract.ChangeAdmin(&_UnistakingContract.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_UnistakingContract *UnistakingContractTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_UnistakingContract *UnistakingContractSession) Implementation() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Implementation(&_UnistakingContract.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_UnistakingContract *UnistakingContractTransactorSession) Implementation() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Implementation(&_UnistakingContract.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UnistakingContract *UnistakingContractTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UnistakingContract *UnistakingContractSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UnistakingContract.Contract.UpgradeTo(&_UnistakingContract.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UnistakingContract *UnistakingContractTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UnistakingContract.Contract.UpgradeTo(&_UnistakingContract.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UnistakingContract *UnistakingContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UnistakingContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UnistakingContract *UnistakingContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UnistakingContract.Contract.UpgradeToAndCall(&_UnistakingContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UnistakingContract *UnistakingContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UnistakingContract.Contract.UpgradeToAndCall(&_UnistakingContract.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UnistakingContract *UnistakingContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UnistakingContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UnistakingContract *UnistakingContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Fallback(&_UnistakingContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UnistakingContract *UnistakingContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UnistakingContract.Contract.Fallback(&_UnistakingContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UnistakingContract *UnistakingContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnistakingContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UnistakingContract *UnistakingContractSession) Receive() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Receive(&_UnistakingContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UnistakingContract *UnistakingContractTransactorSession) Receive() (*types.Transaction, error) {
	return _UnistakingContract.Contract.Receive(&_UnistakingContract.TransactOpts)
}

// UnistakingContractAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UnistakingContract contract.
type UnistakingContractAdminChangedIterator struct {
	Event *UnistakingContractAdminChanged // Event containing the contract specifics and raw log

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
func (it *UnistakingContractAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractAdminChanged)
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
		it.Event = new(UnistakingContractAdminChanged)
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
func (it *UnistakingContractAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractAdminChanged represents a AdminChanged event raised by the UnistakingContract contract.
type UnistakingContractAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UnistakingContract *UnistakingContractFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UnistakingContractAdminChangedIterator, error) {

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UnistakingContractAdminChangedIterator{contract: _UnistakingContract.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UnistakingContract *UnistakingContractFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UnistakingContractAdminChanged) (event.Subscription, error) {

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractAdminChanged)
				if err := _UnistakingContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UnistakingContract *UnistakingContractFilterer) ParseAdminChanged(log types.Log) (*UnistakingContractAdminChanged, error) {
	event := new(UnistakingContractAdminChanged)
	if err := _UnistakingContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the UnistakingContract contract.
type UnistakingContractBeaconUpgradedIterator struct {
	Event *UnistakingContractBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UnistakingContractBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractBeaconUpgraded)
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
		it.Event = new(UnistakingContractBeaconUpgraded)
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
func (it *UnistakingContractBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractBeaconUpgraded represents a BeaconUpgraded event raised by the UnistakingContract contract.
type UnistakingContractBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UnistakingContract *UnistakingContractFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UnistakingContractBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UnistakingContractBeaconUpgradedIterator{contract: _UnistakingContract.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UnistakingContract *UnistakingContractFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UnistakingContractBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractBeaconUpgraded)
				if err := _UnistakingContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UnistakingContract *UnistakingContractFilterer) ParseBeaconUpgraded(log types.Log) (*UnistakingContractBeaconUpgraded, error) {
	event := new(UnistakingContractBeaconUpgraded)
	if err := _UnistakingContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnistakingContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UnistakingContract contract.
type UnistakingContractUpgradedIterator struct {
	Event *UnistakingContractUpgraded // Event containing the contract specifics and raw log

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
func (it *UnistakingContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnistakingContractUpgraded)
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
		it.Event = new(UnistakingContractUpgraded)
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
func (it *UnistakingContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnistakingContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnistakingContractUpgraded represents a Upgraded event raised by the UnistakingContract contract.
type UnistakingContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UnistakingContract *UnistakingContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UnistakingContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UnistakingContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UnistakingContractUpgradedIterator{contract: _UnistakingContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UnistakingContract *UnistakingContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UnistakingContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UnistakingContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnistakingContractUpgraded)
				if err := _UnistakingContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UnistakingContract *UnistakingContractFilterer) ParseUpgraded(log types.Log) (*UnistakingContractUpgraded, error) {
	event := new(UnistakingContractUpgraded)
	if err := _UnistakingContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
