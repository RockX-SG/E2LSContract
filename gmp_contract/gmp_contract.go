// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gmp_contract

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

// GmpContractMetaData contains all meta data concerning the GmpContract contract.
var GmpContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gateway_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasService_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unieth_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gasService\",\"outputs\":[{\"internalType\":\"contractIAxelarGasService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIAxelarGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniEth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"destinationChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationAddress\",\"type\":\"string\"}],\"name\":\"updateExchangeRatio\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// GmpContractABI is the input ABI used to generate the binding from.
// Deprecated: Use GmpContractMetaData.ABI instead.
var GmpContractABI = GmpContractMetaData.ABI

// GmpContract is an auto generated Go binding around an Ethereum contract.
type GmpContract struct {
	GmpContractCaller     // Read-only binding to the contract
	GmpContractTransactor // Write-only binding to the contract
	GmpContractFilterer   // Log filterer for contract events
}

// GmpContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type GmpContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmpContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GmpContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmpContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GmpContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmpContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GmpContractSession struct {
	Contract     *GmpContract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GmpContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GmpContractCallerSession struct {
	Contract *GmpContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GmpContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GmpContractTransactorSession struct {
	Contract     *GmpContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GmpContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type GmpContractRaw struct {
	Contract *GmpContract // Generic contract binding to access the raw methods on
}

// GmpContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GmpContractCallerRaw struct {
	Contract *GmpContractCaller // Generic read-only contract binding to access the raw methods on
}

// GmpContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GmpContractTransactorRaw struct {
	Contract *GmpContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGmpContract creates a new instance of GmpContract, bound to a specific deployed contract.
func NewGmpContract(address common.Address, backend bind.ContractBackend) (*GmpContract, error) {
	contract, err := bindGmpContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GmpContract{GmpContractCaller: GmpContractCaller{contract: contract}, GmpContractTransactor: GmpContractTransactor{contract: contract}, GmpContractFilterer: GmpContractFilterer{contract: contract}}, nil
}

// NewGmpContractCaller creates a new read-only instance of GmpContract, bound to a specific deployed contract.
func NewGmpContractCaller(address common.Address, caller bind.ContractCaller) (*GmpContractCaller, error) {
	contract, err := bindGmpContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GmpContractCaller{contract: contract}, nil
}

// NewGmpContractTransactor creates a new write-only instance of GmpContract, bound to a specific deployed contract.
func NewGmpContractTransactor(address common.Address, transactor bind.ContractTransactor) (*GmpContractTransactor, error) {
	contract, err := bindGmpContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GmpContractTransactor{contract: contract}, nil
}

// NewGmpContractFilterer creates a new log filterer instance of GmpContract, bound to a specific deployed contract.
func NewGmpContractFilterer(address common.Address, filterer bind.ContractFilterer) (*GmpContractFilterer, error) {
	contract, err := bindGmpContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GmpContractFilterer{contract: contract}, nil
}

// bindGmpContract binds a generic wrapper to an already deployed contract.
func bindGmpContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GmpContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmpContract *GmpContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmpContract.Contract.GmpContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmpContract *GmpContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmpContract.Contract.GmpContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmpContract *GmpContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmpContract.Contract.GmpContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmpContract *GmpContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmpContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmpContract *GmpContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmpContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmpContract *GmpContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmpContract.Contract.contract.Transact(opts, method, params...)
}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_GmpContract *GmpContractCaller) GasService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmpContract.contract.Call(opts, &out, "gasService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_GmpContract *GmpContractSession) GasService() (common.Address, error) {
	return _GmpContract.Contract.GasService(&_GmpContract.CallOpts)
}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_GmpContract *GmpContractCallerSession) GasService() (common.Address, error) {
	return _GmpContract.Contract.GasService(&_GmpContract.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_GmpContract *GmpContractCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmpContract.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_GmpContract *GmpContractSession) Gateway() (common.Address, error) {
	return _GmpContract.Contract.Gateway(&_GmpContract.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_GmpContract *GmpContractCallerSession) Gateway() (common.Address, error) {
	return _GmpContract.Contract.Gateway(&_GmpContract.CallOpts)
}

// UniEth is a free data retrieval call binding the contract method 0xf44cc416.
//
// Solidity: function uniEth() view returns(address)
func (_GmpContract *GmpContractCaller) UniEth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmpContract.contract.Call(opts, &out, "uniEth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniEth is a free data retrieval call binding the contract method 0xf44cc416.
//
// Solidity: function uniEth() view returns(address)
func (_GmpContract *GmpContractSession) UniEth() (common.Address, error) {
	return _GmpContract.Contract.UniEth(&_GmpContract.CallOpts)
}

// UniEth is a free data retrieval call binding the contract method 0xf44cc416.
//
// Solidity: function uniEth() view returns(address)
func (_GmpContract *GmpContractCallerSession) UniEth() (common.Address, error) {
	return _GmpContract.Contract.UniEth(&_GmpContract.CallOpts)
}

// UpdateExchangeRatio is a paid mutator transaction binding the contract method 0x6b8a169d.
//
// Solidity: function updateExchangeRatio(string destinationChain, string destinationAddress) payable returns()
func (_GmpContract *GmpContractTransactor) UpdateExchangeRatio(opts *bind.TransactOpts, destinationChain string, destinationAddress string) (*types.Transaction, error) {
	return _GmpContract.contract.Transact(opts, "updateExchangeRatio", destinationChain, destinationAddress)
}

// UpdateExchangeRatio is a paid mutator transaction binding the contract method 0x6b8a169d.
//
// Solidity: function updateExchangeRatio(string destinationChain, string destinationAddress) payable returns()
func (_GmpContract *GmpContractSession) UpdateExchangeRatio(destinationChain string, destinationAddress string) (*types.Transaction, error) {
	return _GmpContract.Contract.UpdateExchangeRatio(&_GmpContract.TransactOpts, destinationChain, destinationAddress)
}

// UpdateExchangeRatio is a paid mutator transaction binding the contract method 0x6b8a169d.
//
// Solidity: function updateExchangeRatio(string destinationChain, string destinationAddress) payable returns()
func (_GmpContract *GmpContractTransactorSession) UpdateExchangeRatio(destinationChain string, destinationAddress string) (*types.Transaction, error) {
	return _GmpContract.Contract.UpdateExchangeRatio(&_GmpContract.TransactOpts, destinationChain, destinationAddress)
}

// GmpContractSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the GmpContract contract.
type GmpContractSentIterator struct {
	Event *GmpContractSent // Event containing the contract specifics and raw log

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
func (it *GmpContractSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmpContractSent)
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
		it.Event = new(GmpContractSent)
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
func (it *GmpContractSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmpContractSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmpContractSent represents a Sent event raised by the GmpContract contract.
type GmpContractSent struct {
	Ratio *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0xff9eb9d26c4f039c77fb69f72cdcd62f2942afe4d95c7b00919c9b6a6bf003f1.
//
// Solidity: event Sent(uint256 ratio)
func (_GmpContract *GmpContractFilterer) FilterSent(opts *bind.FilterOpts) (*GmpContractSentIterator, error) {

	logs, sub, err := _GmpContract.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &GmpContractSentIterator{contract: _GmpContract.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0xff9eb9d26c4f039c77fb69f72cdcd62f2942afe4d95c7b00919c9b6a6bf003f1.
//
// Solidity: event Sent(uint256 ratio)
func (_GmpContract *GmpContractFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *GmpContractSent) (event.Subscription, error) {

	logs, sub, err := _GmpContract.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmpContractSent)
				if err := _GmpContract.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0xff9eb9d26c4f039c77fb69f72cdcd62f2942afe4d95c7b00919c9b6a6bf003f1.
//
// Solidity: event Sent(uint256 ratio)
func (_GmpContract *GmpContractFilterer) ParseSent(log types.Log) (*GmpContractSent, error) {
	event := new(GmpContractSent)
	if err := _GmpContract.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
