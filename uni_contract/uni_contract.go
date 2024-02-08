// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniEth_contract

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

// UniEthContractMetaData contains all meta data concerning the UniEthContract contract.
var UniEthContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniEthContractABI is the input ABI used to generate the binding from.
// Deprecated: Use UniEthContractMetaData.ABI instead.
var UniEthContractABI = UniEthContractMetaData.ABI

// UniEthContract is an auto generated Go binding around an Ethereum contract.
type UniEthContract struct {
	UniEthContractCaller     // Read-only binding to the contract
	UniEthContractTransactor // Write-only binding to the contract
	UniEthContractFilterer   // Log filterer for contract events
}

// UniEthContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniEthContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniEthContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniEthContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniEthContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniEthContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniEthContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniEthContractSession struct {
	Contract     *UniEthContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniEthContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniEthContractCallerSession struct {
	Contract *UniEthContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UniEthContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniEthContractTransactorSession struct {
	Contract     *UniEthContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniEthContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniEthContractRaw struct {
	Contract *UniEthContract // Generic contract binding to access the raw methods on
}

// UniEthContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniEthContractCallerRaw struct {
	Contract *UniEthContractCaller // Generic read-only contract binding to access the raw methods on
}

// UniEthContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniEthContractTransactorRaw struct {
	Contract *UniEthContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniEthContract creates a new instance of UniEthContract, bound to a specific deployed contract.
func NewUniEthContract(address common.Address, backend bind.ContractBackend) (*UniEthContract, error) {
	contract, err := bindUniEthContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniEthContract{UniEthContractCaller: UniEthContractCaller{contract: contract}, UniEthContractTransactor: UniEthContractTransactor{contract: contract}, UniEthContractFilterer: UniEthContractFilterer{contract: contract}}, nil
}

// NewUniEthContractCaller creates a new read-only instance of UniEthContract, bound to a specific deployed contract.
func NewUniEthContractCaller(address common.Address, caller bind.ContractCaller) (*UniEthContractCaller, error) {
	contract, err := bindUniEthContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniEthContractCaller{contract: contract}, nil
}

// NewUniEthContractTransactor creates a new write-only instance of UniEthContract, bound to a specific deployed contract.
func NewUniEthContractTransactor(address common.Address, transactor bind.ContractTransactor) (*UniEthContractTransactor, error) {
	contract, err := bindUniEthContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniEthContractTransactor{contract: contract}, nil
}

// NewUniEthContractFilterer creates a new log filterer instance of UniEthContract, bound to a specific deployed contract.
func NewUniEthContractFilterer(address common.Address, filterer bind.ContractFilterer) (*UniEthContractFilterer, error) {
	contract, err := bindUniEthContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniEthContractFilterer{contract: contract}, nil
}

// bindUniEthContract binds a generic wrapper to an already deployed contract.
func bindUniEthContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniEthContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniEthContract *UniEthContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniEthContract.Contract.UniEthContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniEthContract *UniEthContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniEthContract.Contract.UniEthContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniEthContract *UniEthContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniEthContract.Contract.UniEthContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniEthContract *UniEthContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniEthContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniEthContract *UniEthContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniEthContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniEthContract *UniEthContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniEthContract.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_UniEthContract *UniEthContractTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniEthContract.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_UniEthContract *UniEthContractSession) Admin() (*types.Transaction, error) {
	return _UniEthContract.Contract.Admin(&_UniEthContract.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_UniEthContract *UniEthContractTransactorSession) Admin() (*types.Transaction, error) {
	return _UniEthContract.Contract.Admin(&_UniEthContract.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UniEthContract *UniEthContractTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _UniEthContract.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UniEthContract *UniEthContractSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _UniEthContract.Contract.ChangeAdmin(&_UniEthContract.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_UniEthContract *UniEthContractTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _UniEthContract.Contract.ChangeAdmin(&_UniEthContract.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_UniEthContract *UniEthContractTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniEthContract.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_UniEthContract *UniEthContractSession) Implementation() (*types.Transaction, error) {
	return _UniEthContract.Contract.Implementation(&_UniEthContract.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_UniEthContract *UniEthContractTransactorSession) Implementation() (*types.Transaction, error) {
	return _UniEthContract.Contract.Implementation(&_UniEthContract.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UniEthContract *UniEthContractTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UniEthContract.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UniEthContract *UniEthContractSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UniEthContract.Contract.UpgradeTo(&_UniEthContract.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UniEthContract *UniEthContractTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UniEthContract.Contract.UpgradeTo(&_UniEthContract.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UniEthContract *UniEthContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UniEthContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UniEthContract *UniEthContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UniEthContract.Contract.UpgradeToAndCall(&_UniEthContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UniEthContract *UniEthContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UniEthContract.Contract.UpgradeToAndCall(&_UniEthContract.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UniEthContract *UniEthContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UniEthContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UniEthContract *UniEthContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UniEthContract.Contract.Fallback(&_UniEthContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UniEthContract *UniEthContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UniEthContract.Contract.Fallback(&_UniEthContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniEthContract *UniEthContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniEthContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniEthContract *UniEthContractSession) Receive() (*types.Transaction, error) {
	return _UniEthContract.Contract.Receive(&_UniEthContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniEthContract *UniEthContractTransactorSession) Receive() (*types.Transaction, error) {
	return _UniEthContract.Contract.Receive(&_UniEthContract.TransactOpts)
}

// UniEthContractAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UniEthContract contract.
type UniEthContractAdminChangedIterator struct {
	Event *UniEthContractAdminChanged // Event containing the contract specifics and raw log

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
func (it *UniEthContractAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniEthContractAdminChanged)
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
		it.Event = new(UniEthContractAdminChanged)
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
func (it *UniEthContractAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniEthContractAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniEthContractAdminChanged represents a AdminChanged event raised by the UniEthContract contract.
type UniEthContractAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UniEthContract *UniEthContractFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UniEthContractAdminChangedIterator, error) {

	logs, sub, err := _UniEthContract.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UniEthContractAdminChangedIterator{contract: _UniEthContract.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UniEthContract *UniEthContractFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UniEthContractAdminChanged) (event.Subscription, error) {

	logs, sub, err := _UniEthContract.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniEthContractAdminChanged)
				if err := _UniEthContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_UniEthContract *UniEthContractFilterer) ParseAdminChanged(log types.Log) (*UniEthContractAdminChanged, error) {
	event := new(UniEthContractAdminChanged)
	if err := _UniEthContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniEthContractBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the UniEthContract contract.
type UniEthContractBeaconUpgradedIterator struct {
	Event *UniEthContractBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UniEthContractBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniEthContractBeaconUpgraded)
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
		it.Event = new(UniEthContractBeaconUpgraded)
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
func (it *UniEthContractBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniEthContractBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniEthContractBeaconUpgraded represents a BeaconUpgraded event raised by the UniEthContract contract.
type UniEthContractBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UniEthContract *UniEthContractFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UniEthContractBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UniEthContract.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UniEthContractBeaconUpgradedIterator{contract: _UniEthContract.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UniEthContract *UniEthContractFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UniEthContractBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UniEthContract.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniEthContractBeaconUpgraded)
				if err := _UniEthContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_UniEthContract *UniEthContractFilterer) ParseBeaconUpgraded(log types.Log) (*UniEthContractBeaconUpgraded, error) {
	event := new(UniEthContractBeaconUpgraded)
	if err := _UniEthContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniEthContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UniEthContract contract.
type UniEthContractUpgradedIterator struct {
	Event *UniEthContractUpgraded // Event containing the contract specifics and raw log

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
func (it *UniEthContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniEthContractUpgraded)
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
		it.Event = new(UniEthContractUpgraded)
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
func (it *UniEthContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniEthContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniEthContractUpgraded represents a Upgraded event raised by the UniEthContract contract.
type UniEthContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UniEthContract *UniEthContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UniEthContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UniEthContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UniEthContractUpgradedIterator{contract: _UniEthContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UniEthContract *UniEthContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UniEthContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UniEthContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniEthContractUpgraded)
				if err := _UniEthContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_UniEthContract *UniEthContractFilterer) ParseUpgraded(log types.Log) (*UniEthContractUpgraded, error) {
	event := new(UniEthContractUpgraded)
	if err := _UniEthContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
