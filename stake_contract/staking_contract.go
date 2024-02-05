// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking_contract

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

// StakingContractMetaData contains all meta data concerning the StakingContract contract.
var StakingContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AutoCompoundToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"BalanceSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creditor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountEther\",\"type\":\"uint256\"}],\"name\":\"DebtQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DepositContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAccountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"milli\",\"type\":\"uint256\"}],\"name\":\"ManagerFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ManagerFeeWithdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RedeemContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RestakingAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RevenueAccounted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextValidatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stoppedCount\",\"type\":\"uint256\"}],\"name\":\"ValidatorSlashedStopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stoppedCount\",\"type\":\"uint256\"}],\"name\":\"ValidatorStopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"WhiteListToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawCredential\",\"type\":\"bytes32\"}],\"name\":\"WithdrawCredentialSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"XETHContractSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"checkDebt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDepositContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountedBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountedManagerRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountedUserRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDebtQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingEthers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getQuota\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecentReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecentSlashed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idx_to\",\"type\":\"uint256\"}],\"name\":\"getRegisteredValidators\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"stopped\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idx_to\",\"type\":\"uint256\"}],\"name\":\"getRegisteredValidatorsV2\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"stopped\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"restaking\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportedValidatorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportedValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStoppedValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVectorClock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"restakingContract_\",\"type\":\"address\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhiteListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_aliveValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"clock\",\"type\":\"bytes32\"}],\"name\":\"pushBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_aliveValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"clock\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"pushBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethersToRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"redeemFromValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"burned\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"registerRestakingValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"registerValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"oldpubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"replaceValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"oldpubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"restaking\",\"type\":\"bool\"}],\"name\":\"replaceValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethDepositContract\",\"type\":\"address\"}],\"name\":\"setETHDepositContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"milli\",\"type\":\"uint256\"}],\"name\":\"setManagerFeeShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_redeemContract\",\"type\":\"address\"}],\"name\":\"setRedeemContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"withdrawalCredentials_\",\"type\":\"bytes32\"}],\"name\":\"setWithdrawCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xETHContract\",\"type\":\"address\"}],\"name\":\"setXETHContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPhase\",\"type\":\"uint256\"}],\"name\":\"switchPhase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleAutoCompound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"toggleWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_stoppedPubKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"clock\",\"type\":\"bytes32\"}],\"name\":\"validatorSlashedStop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_stoppedPubKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"clock\",\"type\":\"bytes32\"}],\"name\":\"validatorStopped\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawManagerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xETHAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingContractMetaData.ABI instead.
var StakingContractABI = StakingContractMetaData.ABI

// StakingContract is an auto generated Go binding around an Ethereum contract.
type StakingContract struct {
	StakingContractCaller     // Read-only binding to the contract
	StakingContractTransactor // Write-only binding to the contract
	StakingContractFilterer   // Log filterer for contract events
}

// StakingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingContractSession struct {
	Contract     *StakingContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingContractCallerSession struct {
	Contract *StakingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StakingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingContractTransactorSession struct {
	Contract     *StakingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StakingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingContractRaw struct {
	Contract *StakingContract // Generic contract binding to access the raw methods on
}

// StakingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingContractCallerRaw struct {
	Contract *StakingContractCaller // Generic read-only contract binding to access the raw methods on
}

// StakingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingContractTransactorRaw struct {
	Contract *StakingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingContract creates a new instance of StakingContract, bound to a specific deployed contract.
func NewStakingContract(address common.Address, backend bind.ContractBackend) (*StakingContract, error) {
	contract, err := bindStakingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingContract{StakingContractCaller: StakingContractCaller{contract: contract}, StakingContractTransactor: StakingContractTransactor{contract: contract}, StakingContractFilterer: StakingContractFilterer{contract: contract}}, nil
}

// NewStakingContractCaller creates a new read-only instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractCaller(address common.Address, caller bind.ContractCaller) (*StakingContractCaller, error) {
	contract, err := bindStakingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingContractCaller{contract: contract}, nil
}

// NewStakingContractTransactor creates a new write-only instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingContractTransactor, error) {
	contract, err := bindStakingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingContractTransactor{contract: contract}, nil
}

// NewStakingContractFilterer creates a new log filterer instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingContractFilterer, error) {
	contract, err := bindStakingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingContractFilterer{contract: contract}, nil
}

// bindStakingContract binds a generic wrapper to an already deployed contract.
func bindStakingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingContract *StakingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingContract.Contract.StakingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingContract *StakingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.Contract.StakingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingContract *StakingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingContract.Contract.StakingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingContract *StakingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingContract *StakingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingContract *StakingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingContract.Contract.DEFAULTADMINROLE(&_StakingContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingContract.Contract.DEFAULTADMINROLE(&_StakingContract.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingContract *StakingContractCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingContract *StakingContractSession) DEPOSITSIZE() (*big.Int, error) {
	return _StakingContract.Contract.DEPOSITSIZE(&_StakingContract.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _StakingContract.Contract.DEPOSITSIZE(&_StakingContract.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) MANAGERROLE() ([32]byte, error) {
	return _StakingContract.Contract.MANAGERROLE(&_StakingContract.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) MANAGERROLE() ([32]byte, error) {
	return _StakingContract.Contract.MANAGERROLE(&_StakingContract.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) ORACLEROLE() ([32]byte, error) {
	return _StakingContract.Contract.ORACLEROLE(&_StakingContract.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) ORACLEROLE() ([32]byte, error) {
	return _StakingContract.Contract.ORACLEROLE(&_StakingContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) PAUSERROLE() ([32]byte, error) {
	return _StakingContract.Contract.PAUSERROLE(&_StakingContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) PAUSERROLE() ([32]byte, error) {
	return _StakingContract.Contract.PAUSERROLE(&_StakingContract.CallOpts)
}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCaller) REGISTRYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "REGISTRY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractSession) REGISTRYROLE() ([32]byte, error) {
	return _StakingContract.Contract.REGISTRYROLE(&_StakingContract.CallOpts)
}

// REGISTRYROLE is a free data retrieval call binding the contract method 0x42f1e879.
//
// Solidity: function REGISTRY_ROLE() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) REGISTRYROLE() ([32]byte, error) {
	return _StakingContract.Contract.REGISTRYROLE(&_StakingContract.CallOpts)
}

// CheckDebt is a free data retrieval call binding the contract method 0xc8c3df4a.
//
// Solidity: function checkDebt(uint256 index) view returns(address account, uint256 amount)
func (_StakingContract *StakingContractCaller) CheckDebt(opts *bind.CallOpts, index *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "checkDebt", index)

	outstruct := new(struct {
		Account common.Address
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CheckDebt is a free data retrieval call binding the contract method 0xc8c3df4a.
//
// Solidity: function checkDebt(uint256 index) view returns(address account, uint256 amount)
func (_StakingContract *StakingContractSession) CheckDebt(index *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	return _StakingContract.Contract.CheckDebt(&_StakingContract.CallOpts, index)
}

// CheckDebt is a free data retrieval call binding the contract method 0xc8c3df4a.
//
// Solidity: function checkDebt(uint256 index) view returns(address account, uint256 amount)
func (_StakingContract *StakingContractCallerSession) CheckDebt(index *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	return _StakingContract.Contract.CheckDebt(&_StakingContract.CallOpts, index)
}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_StakingContract *StakingContractCaller) CurrentReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "currentReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_StakingContract *StakingContractSession) CurrentReserve() (*big.Int, error) {
	return _StakingContract.Contract.CurrentReserve(&_StakingContract.CallOpts)
}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) CurrentReserve() (*big.Int, error) {
	return _StakingContract.Contract.CurrentReserve(&_StakingContract.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_StakingContract *StakingContractCaller) DebtOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "debtOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_StakingContract *StakingContractSession) DebtOf(account common.Address) (*big.Int, error) {
	return _StakingContract.Contract.DebtOf(&_StakingContract.CallOpts, account)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_StakingContract *StakingContractCallerSession) DebtOf(account common.Address) (*big.Int, error) {
	return _StakingContract.Contract.DebtOf(&_StakingContract.CallOpts, account)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingContract *StakingContractCaller) EthDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "ethDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingContract *StakingContractSession) EthDepositContract() (common.Address, error) {
	return _StakingContract.Contract.EthDepositContract(&_StakingContract.CallOpts)
}

// EthDepositContract is a free data retrieval call binding the contract method 0x3884545d.
//
// Solidity: function ethDepositContract() view returns(address)
func (_StakingContract *StakingContractCallerSession) EthDepositContract() (common.Address, error) {
	return _StakingContract.Contract.EthDepositContract(&_StakingContract.CallOpts)
}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256)
func (_StakingContract *StakingContractCaller) ExchangeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "exchangeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256)
func (_StakingContract *StakingContractSession) ExchangeRatio() (*big.Int, error) {
	return _StakingContract.Contract.ExchangeRatio(&_StakingContract.CallOpts)
}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) ExchangeRatio() (*big.Int, error) {
	return _StakingContract.Contract.ExchangeRatio(&_StakingContract.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(int256)
func (_StakingContract *StakingContractCaller) GetAccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getAccountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(int256)
func (_StakingContract *StakingContractSession) GetAccountedBalance() (*big.Int, error) {
	return _StakingContract.Contract.GetAccountedBalance(&_StakingContract.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(int256)
func (_StakingContract *StakingContractCallerSession) GetAccountedBalance() (*big.Int, error) {
	return _StakingContract.Contract.GetAccountedBalance(&_StakingContract.CallOpts)
}

// GetAccountedManagerRevenue is a free data retrieval call binding the contract method 0x08b84c0c.
//
// Solidity: function getAccountedManagerRevenue() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetAccountedManagerRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getAccountedManagerRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedManagerRevenue is a free data retrieval call binding the contract method 0x08b84c0c.
//
// Solidity: function getAccountedManagerRevenue() view returns(uint256)
func (_StakingContract *StakingContractSession) GetAccountedManagerRevenue() (*big.Int, error) {
	return _StakingContract.Contract.GetAccountedManagerRevenue(&_StakingContract.CallOpts)
}

// GetAccountedManagerRevenue is a free data retrieval call binding the contract method 0x08b84c0c.
//
// Solidity: function getAccountedManagerRevenue() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetAccountedManagerRevenue() (*big.Int, error) {
	return _StakingContract.Contract.GetAccountedManagerRevenue(&_StakingContract.CallOpts)
}

// GetAccountedUserRevenue is a free data retrieval call binding the contract method 0x61c993c5.
//
// Solidity: function getAccountedUserRevenue() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetAccountedUserRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getAccountedUserRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedUserRevenue is a free data retrieval call binding the contract method 0x61c993c5.
//
// Solidity: function getAccountedUserRevenue() view returns(uint256)
func (_StakingContract *StakingContractSession) GetAccountedUserRevenue() (*big.Int, error) {
	return _StakingContract.Contract.GetAccountedUserRevenue(&_StakingContract.CallOpts)
}

// GetAccountedUserRevenue is a free data retrieval call binding the contract method 0x61c993c5.
//
// Solidity: function getAccountedUserRevenue() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetAccountedUserRevenue() (*big.Int, error) {
	return _StakingContract.Contract.GetAccountedUserRevenue(&_StakingContract.CallOpts)
}

// GetCurrentDebts is a free data retrieval call binding the contract method 0x8b0bfd35.
//
// Solidity: function getCurrentDebts() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetCurrentDebts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getCurrentDebts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentDebts is a free data retrieval call binding the contract method 0x8b0bfd35.
//
// Solidity: function getCurrentDebts() view returns(uint256)
func (_StakingContract *StakingContractSession) GetCurrentDebts() (*big.Int, error) {
	return _StakingContract.Contract.GetCurrentDebts(&_StakingContract.CallOpts)
}

// GetCurrentDebts is a free data retrieval call binding the contract method 0x8b0bfd35.
//
// Solidity: function getCurrentDebts() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetCurrentDebts() (*big.Int, error) {
	return _StakingContract.Contract.GetCurrentDebts(&_StakingContract.CallOpts)
}

// GetDebtQueue is a free data retrieval call binding the contract method 0xdc3fc3b2.
//
// Solidity: function getDebtQueue() view returns(uint256 first, uint256 last)
func (_StakingContract *StakingContractCaller) GetDebtQueue(opts *bind.CallOpts) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getDebtQueue")

	outstruct := new(struct {
		First *big.Int
		Last  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.First = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDebtQueue is a free data retrieval call binding the contract method 0xdc3fc3b2.
//
// Solidity: function getDebtQueue() view returns(uint256 first, uint256 last)
func (_StakingContract *StakingContractSession) GetDebtQueue() (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _StakingContract.Contract.GetDebtQueue(&_StakingContract.CallOpts)
}

// GetDebtQueue is a free data retrieval call binding the contract method 0xdc3fc3b2.
//
// Solidity: function getDebtQueue() view returns(uint256 first, uint256 last)
func (_StakingContract *StakingContractCallerSession) GetDebtQueue() (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _StakingContract.Contract.GetDebtQueue(&_StakingContract.CallOpts)
}

// GetNextValidatorId is a free data retrieval call binding the contract method 0xda863b3b.
//
// Solidity: function getNextValidatorId() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetNextValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getNextValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextValidatorId is a free data retrieval call binding the contract method 0xda863b3b.
//
// Solidity: function getNextValidatorId() view returns(uint256)
func (_StakingContract *StakingContractSession) GetNextValidatorId() (*big.Int, error) {
	return _StakingContract.Contract.GetNextValidatorId(&_StakingContract.CallOpts)
}

// GetNextValidatorId is a free data retrieval call binding the contract method 0xda863b3b.
//
// Solidity: function getNextValidatorId() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetNextValidatorId() (*big.Int, error) {
	return _StakingContract.Contract.GetNextValidatorId(&_StakingContract.CallOpts)
}

// GetPendingEthers is a free data retrieval call binding the contract method 0x64363f2b.
//
// Solidity: function getPendingEthers() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetPendingEthers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getPendingEthers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingEthers is a free data retrieval call binding the contract method 0x64363f2b.
//
// Solidity: function getPendingEthers() view returns(uint256)
func (_StakingContract *StakingContractSession) GetPendingEthers() (*big.Int, error) {
	return _StakingContract.Contract.GetPendingEthers(&_StakingContract.CallOpts)
}

// GetPendingEthers is a free data retrieval call binding the contract method 0x64363f2b.
//
// Solidity: function getPendingEthers() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetPendingEthers() (*big.Int, error) {
	return _StakingContract.Contract.GetPendingEthers(&_StakingContract.CallOpts)
}

// GetQuota is a free data retrieval call binding the contract method 0x43a2a302.
//
// Solidity: function getQuota(address account) view returns(uint256)
func (_StakingContract *StakingContractCaller) GetQuota(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getQuota", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuota is a free data retrieval call binding the contract method 0x43a2a302.
//
// Solidity: function getQuota(address account) view returns(uint256)
func (_StakingContract *StakingContractSession) GetQuota(account common.Address) (*big.Int, error) {
	return _StakingContract.Contract.GetQuota(&_StakingContract.CallOpts, account)
}

// GetQuota is a free data retrieval call binding the contract method 0x43a2a302.
//
// Solidity: function getQuota(address account) view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetQuota(account common.Address) (*big.Int, error) {
	return _StakingContract.Contract.GetQuota(&_StakingContract.CallOpts, account)
}

// GetRecentReceived is a free data retrieval call binding the contract method 0xe08f2d89.
//
// Solidity: function getRecentReceived() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetRecentReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRecentReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecentReceived is a free data retrieval call binding the contract method 0xe08f2d89.
//
// Solidity: function getRecentReceived() view returns(uint256)
func (_StakingContract *StakingContractSession) GetRecentReceived() (*big.Int, error) {
	return _StakingContract.Contract.GetRecentReceived(&_StakingContract.CallOpts)
}

// GetRecentReceived is a free data retrieval call binding the contract method 0xe08f2d89.
//
// Solidity: function getRecentReceived() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetRecentReceived() (*big.Int, error) {
	return _StakingContract.Contract.GetRecentReceived(&_StakingContract.CallOpts)
}

// GetRecentSlashed is a free data retrieval call binding the contract method 0xaba52536.
//
// Solidity: function getRecentSlashed() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetRecentSlashed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRecentSlashed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecentSlashed is a free data retrieval call binding the contract method 0xaba52536.
//
// Solidity: function getRecentSlashed() view returns(uint256)
func (_StakingContract *StakingContractSession) GetRecentSlashed() (*big.Int, error) {
	return _StakingContract.Contract.GetRecentSlashed(&_StakingContract.CallOpts)
}

// GetRecentSlashed is a free data retrieval call binding the contract method 0xaba52536.
//
// Solidity: function getRecentSlashed() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetRecentSlashed() (*big.Int, error) {
	return _StakingContract.Contract.GetRecentSlashed(&_StakingContract.CallOpts)
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xf22abf37.
//
// Solidity: function getRegisteredValidators(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped)
func (_StakingContract *StakingContractCaller) GetRegisteredValidators(opts *bind.CallOpts, idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
}, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRegisteredValidators", idx_from, idx_to)

	outstruct := new(struct {
		Pubkeys    [][]byte
		Signatures [][]byte
		Stopped    []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkeys = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.Signatures = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	outstruct.Stopped = *abi.ConvertType(out[2], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xf22abf37.
//
// Solidity: function getRegisteredValidators(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped)
func (_StakingContract *StakingContractSession) GetRegisteredValidators(idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
}, error) {
	return _StakingContract.Contract.GetRegisteredValidators(&_StakingContract.CallOpts, idx_from, idx_to)
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xf22abf37.
//
// Solidity: function getRegisteredValidators(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped)
func (_StakingContract *StakingContractCallerSession) GetRegisteredValidators(idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
}, error) {
	return _StakingContract.Contract.GetRegisteredValidators(&_StakingContract.CallOpts, idx_from, idx_to)
}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetRegisteredValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRegisteredValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractSession) GetRegisteredValidatorsCount() (*big.Int, error) {
	return _StakingContract.Contract.GetRegisteredValidatorsCount(&_StakingContract.CallOpts)
}

// GetRegisteredValidatorsCount is a free data retrieval call binding the contract method 0x30b12c8d.
//
// Solidity: function getRegisteredValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetRegisteredValidatorsCount() (*big.Int, error) {
	return _StakingContract.Contract.GetRegisteredValidatorsCount(&_StakingContract.CallOpts)
}

// GetRegisteredValidatorsV2 is a free data retrieval call binding the contract method 0x3bf39dca.
//
// Solidity: function getRegisteredValidatorsV2(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped, bool[] restaking)
func (_StakingContract *StakingContractCaller) GetRegisteredValidatorsV2(opts *bind.CallOpts, idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
	Restaking  []bool
}, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRegisteredValidatorsV2", idx_from, idx_to)

	outstruct := new(struct {
		Pubkeys    [][]byte
		Signatures [][]byte
		Stopped    []bool
		Restaking  []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkeys = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.Signatures = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	outstruct.Stopped = *abi.ConvertType(out[2], new([]bool)).(*[]bool)
	outstruct.Restaking = *abi.ConvertType(out[3], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetRegisteredValidatorsV2 is a free data retrieval call binding the contract method 0x3bf39dca.
//
// Solidity: function getRegisteredValidatorsV2(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped, bool[] restaking)
func (_StakingContract *StakingContractSession) GetRegisteredValidatorsV2(idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
	Restaking  []bool
}, error) {
	return _StakingContract.Contract.GetRegisteredValidatorsV2(&_StakingContract.CallOpts, idx_from, idx_to)
}

// GetRegisteredValidatorsV2 is a free data retrieval call binding the contract method 0x3bf39dca.
//
// Solidity: function getRegisteredValidatorsV2(uint256 idx_from, uint256 idx_to) view returns(bytes[] pubkeys, bytes[] signatures, bool[] stopped, bool[] restaking)
func (_StakingContract *StakingContractCallerSession) GetRegisteredValidatorsV2(idx_from *big.Int, idx_to *big.Int) (struct {
	Pubkeys    [][]byte
	Signatures [][]byte
	Stopped    []bool
	Restaking  []bool
}, error) {
	return _StakingContract.Contract.GetRegisteredValidatorsV2(&_StakingContract.CallOpts, idx_from, idx_to)
}

// GetReportedValidatorBalance is a free data retrieval call binding the contract method 0xcdb54a1b.
//
// Solidity: function getReportedValidatorBalance() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetReportedValidatorBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getReportedValidatorBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportedValidatorBalance is a free data retrieval call binding the contract method 0xcdb54a1b.
//
// Solidity: function getReportedValidatorBalance() view returns(uint256)
func (_StakingContract *StakingContractSession) GetReportedValidatorBalance() (*big.Int, error) {
	return _StakingContract.Contract.GetReportedValidatorBalance(&_StakingContract.CallOpts)
}

// GetReportedValidatorBalance is a free data retrieval call binding the contract method 0xcdb54a1b.
//
// Solidity: function getReportedValidatorBalance() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetReportedValidatorBalance() (*big.Int, error) {
	return _StakingContract.Contract.GetReportedValidatorBalance(&_StakingContract.CallOpts)
}

// GetReportedValidators is a free data retrieval call binding the contract method 0x6a42602c.
//
// Solidity: function getReportedValidators() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetReportedValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getReportedValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportedValidators is a free data retrieval call binding the contract method 0x6a42602c.
//
// Solidity: function getReportedValidators() view returns(uint256)
func (_StakingContract *StakingContractSession) GetReportedValidators() (*big.Int, error) {
	return _StakingContract.Contract.GetReportedValidators(&_StakingContract.CallOpts)
}

// GetReportedValidators is a free data retrieval call binding the contract method 0x6a42602c.
//
// Solidity: function getReportedValidators() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetReportedValidators() (*big.Int, error) {
	return _StakingContract.Contract.GetReportedValidators(&_StakingContract.CallOpts)
}

// GetRewardDebts is a free data retrieval call binding the contract method 0xa065913f.
//
// Solidity: function getRewardDebts() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetRewardDebts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRewardDebts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardDebts is a free data retrieval call binding the contract method 0xa065913f.
//
// Solidity: function getRewardDebts() view returns(uint256)
func (_StakingContract *StakingContractSession) GetRewardDebts() (*big.Int, error) {
	return _StakingContract.Contract.GetRewardDebts(&_StakingContract.CallOpts)
}

// GetRewardDebts is a free data retrieval call binding the contract method 0xa065913f.
//
// Solidity: function getRewardDebts() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetRewardDebts() (*big.Int, error) {
	return _StakingContract.Contract.GetRewardDebts(&_StakingContract.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingContract *StakingContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingContract *StakingContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingContract.Contract.GetRoleAdmin(&_StakingContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingContract.Contract.GetRoleAdmin(&_StakingContract.CallOpts, role)
}

// GetStoppedValidatorsCount is a free data retrieval call binding the contract method 0x49557df9.
//
// Solidity: function getStoppedValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetStoppedValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getStoppedValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStoppedValidatorsCount is a free data retrieval call binding the contract method 0x49557df9.
//
// Solidity: function getStoppedValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractSession) GetStoppedValidatorsCount() (*big.Int, error) {
	return _StakingContract.Contract.GetStoppedValidatorsCount(&_StakingContract.CallOpts)
}

// GetStoppedValidatorsCount is a free data retrieval call binding the contract method 0x49557df9.
//
// Solidity: function getStoppedValidatorsCount() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetStoppedValidatorsCount() (*big.Int, error) {
	return _StakingContract.Contract.GetStoppedValidatorsCount(&_StakingContract.CallOpts)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_StakingContract *StakingContractCaller) GetTotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getTotalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_StakingContract *StakingContractSession) GetTotalStaked() (*big.Int, error) {
	return _StakingContract.Contract.GetTotalStaked(&_StakingContract.CallOpts)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) GetTotalStaked() (*big.Int, error) {
	return _StakingContract.Contract.GetTotalStaked(&_StakingContract.CallOpts)
}

// GetVectorClock is a free data retrieval call binding the contract method 0xecacf56d.
//
// Solidity: function getVectorClock() view returns(bytes32)
func (_StakingContract *StakingContractCaller) GetVectorClock(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "getVectorClock")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVectorClock is a free data retrieval call binding the contract method 0xecacf56d.
//
// Solidity: function getVectorClock() view returns(bytes32)
func (_StakingContract *StakingContractSession) GetVectorClock() ([32]byte, error) {
	return _StakingContract.Contract.GetVectorClock(&_StakingContract.CallOpts)
}

// GetVectorClock is a free data retrieval call binding the contract method 0xecacf56d.
//
// Solidity: function getVectorClock() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) GetVectorClock() ([32]byte, error) {
	return _StakingContract.Contract.GetVectorClock(&_StakingContract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingContract *StakingContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingContract *StakingContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingContract.Contract.HasRole(&_StakingContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingContract *StakingContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingContract.Contract.HasRole(&_StakingContract.CallOpts, role, account)
}

// IsWhiteListed is a free data retrieval call binding the contract method 0x6f9170f6.
//
// Solidity: function isWhiteListed(address account) view returns(bool)
func (_StakingContract *StakingContractCaller) IsWhiteListed(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "isWhiteListed", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhiteListed is a free data retrieval call binding the contract method 0x6f9170f6.
//
// Solidity: function isWhiteListed(address account) view returns(bool)
func (_StakingContract *StakingContractSession) IsWhiteListed(account common.Address) (bool, error) {
	return _StakingContract.Contract.IsWhiteListed(&_StakingContract.CallOpts, account)
}

// IsWhiteListed is a free data retrieval call binding the contract method 0x6f9170f6.
//
// Solidity: function isWhiteListed(address account) view returns(bool)
func (_StakingContract *StakingContractCallerSession) IsWhiteListed(account common.Address) (bool, error) {
	return _StakingContract.Contract.IsWhiteListed(&_StakingContract.CallOpts, account)
}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_StakingContract *StakingContractCaller) ManagerFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "managerFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_StakingContract *StakingContractSession) ManagerFeeShare() (*big.Int, error) {
	return _StakingContract.Contract.ManagerFeeShare(&_StakingContract.CallOpts)
}

// ManagerFeeShare is a free data retrieval call binding the contract method 0xe43a4954.
//
// Solidity: function managerFeeShare() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) ManagerFeeShare() (*big.Int, error) {
	return _StakingContract.Contract.ManagerFeeShare(&_StakingContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractSession) Paused() (bool, error) {
	return _StakingContract.Contract.Paused(&_StakingContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractCallerSession) Paused() (bool, error) {
	return _StakingContract.Contract.Paused(&_StakingContract.CallOpts)
}

// RedeemContract is a free data retrieval call binding the contract method 0x7a4473e1.
//
// Solidity: function redeemContract() view returns(address)
func (_StakingContract *StakingContractCaller) RedeemContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "redeemContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedeemContract is a free data retrieval call binding the contract method 0x7a4473e1.
//
// Solidity: function redeemContract() view returns(address)
func (_StakingContract *StakingContractSession) RedeemContract() (common.Address, error) {
	return _StakingContract.Contract.RedeemContract(&_StakingContract.CallOpts)
}

// RedeemContract is a free data retrieval call binding the contract method 0x7a4473e1.
//
// Solidity: function redeemContract() view returns(address)
func (_StakingContract *StakingContractCallerSession) RedeemContract() (common.Address, error) {
	return _StakingContract.Contract.RedeemContract(&_StakingContract.CallOpts)
}

// RestakingContract is a free data retrieval call binding the contract method 0xf1f3b3e7.
//
// Solidity: function restakingContract() view returns(address)
func (_StakingContract *StakingContractCaller) RestakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "restakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RestakingContract is a free data retrieval call binding the contract method 0xf1f3b3e7.
//
// Solidity: function restakingContract() view returns(address)
func (_StakingContract *StakingContractSession) RestakingContract() (common.Address, error) {
	return _StakingContract.Contract.RestakingContract(&_StakingContract.CallOpts)
}

// RestakingContract is a free data retrieval call binding the contract method 0xf1f3b3e7.
//
// Solidity: function restakingContract() view returns(address)
func (_StakingContract *StakingContractCallerSession) RestakingContract() (common.Address, error) {
	return _StakingContract.Contract.RestakingContract(&_StakingContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingContract *StakingContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingContract *StakingContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingContract.Contract.SupportsInterface(&_StakingContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingContract *StakingContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingContract.Contract.SupportsInterface(&_StakingContract.CallOpts, interfaceId)
}

// WithdrawalCredentials is a free data retrieval call binding the contract method 0x4cd79e0a.
//
// Solidity: function withdrawalCredentials() view returns(bytes32)
func (_StakingContract *StakingContractCaller) WithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "withdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawalCredentials is a free data retrieval call binding the contract method 0x4cd79e0a.
//
// Solidity: function withdrawalCredentials() view returns(bytes32)
func (_StakingContract *StakingContractSession) WithdrawalCredentials() ([32]byte, error) {
	return _StakingContract.Contract.WithdrawalCredentials(&_StakingContract.CallOpts)
}

// WithdrawalCredentials is a free data retrieval call binding the contract method 0x4cd79e0a.
//
// Solidity: function withdrawalCredentials() view returns(bytes32)
func (_StakingContract *StakingContractCallerSession) WithdrawalCredentials() ([32]byte, error) {
	return _StakingContract.Contract.WithdrawalCredentials(&_StakingContract.CallOpts)
}

// XETHAddress is a free data retrieval call binding the contract method 0xb181033a.
//
// Solidity: function xETHAddress() view returns(address)
func (_StakingContract *StakingContractCaller) XETHAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "xETHAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XETHAddress is a free data retrieval call binding the contract method 0xb181033a.
//
// Solidity: function xETHAddress() view returns(address)
func (_StakingContract *StakingContractSession) XETHAddress() (common.Address, error) {
	return _StakingContract.Contract.XETHAddress(&_StakingContract.CallOpts)
}

// XETHAddress is a free data retrieval call binding the contract method 0xb181033a.
//
// Solidity: function xETHAddress() view returns(address)
func (_StakingContract *StakingContractCallerSession) XETHAddress() (common.Address, error) {
	return _StakingContract.Contract.XETHAddress(&_StakingContract.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.GrantRole(&_StakingContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.GrantRole(&_StakingContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingContract *StakingContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingContract *StakingContractSession) Initialize() (*types.Transaction, error) {
	return _StakingContract.Contract.Initialize(&_StakingContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakingContract *StakingContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _StakingContract.Contract.Initialize(&_StakingContract.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address restakingContract_) returns()
func (_StakingContract *StakingContractTransactor) InitializeV2(opts *bind.TransactOpts, restakingContract_ common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "initializeV2", restakingContract_)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address restakingContract_) returns()
func (_StakingContract *StakingContractSession) InitializeV2(restakingContract_ common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.InitializeV2(&_StakingContract.TransactOpts, restakingContract_)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address restakingContract_) returns()
func (_StakingContract *StakingContractTransactorSession) InitializeV2(restakingContract_ common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.InitializeV2(&_StakingContract.TransactOpts, restakingContract_)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_StakingContract *StakingContractTransactor) Mint(opts *bind.TransactOpts, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "mint", minToMint, deadline)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_StakingContract *StakingContractSession) Mint(minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Mint(&_StakingContract.TransactOpts, minToMint, deadline)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_StakingContract *StakingContractTransactorSession) Mint(minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Mint(&_StakingContract.TransactOpts, minToMint, deadline)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractSession) Pause() (*types.Transaction, error) {
	return _StakingContract.Contract.Pause(&_StakingContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractTransactorSession) Pause() (*types.Transaction, error) {
	return _StakingContract.Contract.Pause(&_StakingContract.TransactOpts)
}

// PushBeacon is a paid mutator transaction binding the contract method 0x496bf5bb.
//
// Solidity: function pushBeacon(uint256 _aliveValidators, uint256 , bytes32 clock) returns()
func (_StakingContract *StakingContractTransactor) PushBeacon(opts *bind.TransactOpts, _aliveValidators *big.Int, arg1 *big.Int, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "pushBeacon", _aliveValidators, arg1, clock)
}

// PushBeacon is a paid mutator transaction binding the contract method 0x496bf5bb.
//
// Solidity: function pushBeacon(uint256 _aliveValidators, uint256 , bytes32 clock) returns()
func (_StakingContract *StakingContractSession) PushBeacon(_aliveValidators *big.Int, arg1 *big.Int, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.PushBeacon(&_StakingContract.TransactOpts, _aliveValidators, arg1, clock)
}

// PushBeacon is a paid mutator transaction binding the contract method 0x496bf5bb.
//
// Solidity: function pushBeacon(uint256 _aliveValidators, uint256 , bytes32 clock) returns()
func (_StakingContract *StakingContractTransactorSession) PushBeacon(_aliveValidators *big.Int, arg1 *big.Int, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.PushBeacon(&_StakingContract.TransactOpts, _aliveValidators, arg1, clock)
}

// PushBeacon0 is a paid mutator transaction binding the contract method 0x93454817.
//
// Solidity: function pushBeacon(uint256 _aliveValidators, uint256 , bytes32 clock, uint256 limit) returns()
func (_StakingContract *StakingContractTransactor) PushBeacon0(opts *bind.TransactOpts, _aliveValidators *big.Int, arg1 *big.Int, clock [32]byte, limit *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "pushBeacon0", _aliveValidators, arg1, clock, limit)
}

// PushBeacon0 is a paid mutator transaction binding the contract method 0x93454817.
//
// Solidity: function pushBeacon(uint256 _aliveValidators, uint256 , bytes32 clock, uint256 limit) returns()
func (_StakingContract *StakingContractSession) PushBeacon0(_aliveValidators *big.Int, arg1 *big.Int, clock [32]byte, limit *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.PushBeacon0(&_StakingContract.TransactOpts, _aliveValidators, arg1, clock, limit)
}

// PushBeacon0 is a paid mutator transaction binding the contract method 0x93454817.
//
// Solidity: function pushBeacon(uint256 _aliveValidators, uint256 , bytes32 clock, uint256 limit) returns()
func (_StakingContract *StakingContractTransactorSession) PushBeacon0(_aliveValidators *big.Int, arg1 *big.Int, clock [32]byte, limit *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.PushBeacon0(&_StakingContract.TransactOpts, _aliveValidators, arg1, clock, limit)
}

// RedeemFromValidators is a paid mutator transaction binding the contract method 0xf5404d60.
//
// Solidity: function redeemFromValidators(uint256 ethersToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_StakingContract *StakingContractTransactor) RedeemFromValidators(opts *bind.TransactOpts, ethersToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "redeemFromValidators", ethersToRedeem, maxToBurn, deadline)
}

// RedeemFromValidators is a paid mutator transaction binding the contract method 0xf5404d60.
//
// Solidity: function redeemFromValidators(uint256 ethersToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_StakingContract *StakingContractSession) RedeemFromValidators(ethersToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.RedeemFromValidators(&_StakingContract.TransactOpts, ethersToRedeem, maxToBurn, deadline)
}

// RedeemFromValidators is a paid mutator transaction binding the contract method 0xf5404d60.
//
// Solidity: function redeemFromValidators(uint256 ethersToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_StakingContract *StakingContractTransactorSession) RedeemFromValidators(ethersToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.RedeemFromValidators(&_StakingContract.TransactOpts, ethersToRedeem, maxToBurn, deadline)
}

// RegisterRestakingValidators is a paid mutator transaction binding the contract method 0x4e8ab1a1.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_StakingContract *StakingContractTransactor) RegisterRestakingValidators(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "registerRestakingValidators", pubkeys, signatures)
}

// RegisterRestakingValidators is a paid mutator transaction binding the contract method 0x4e8ab1a1.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_StakingContract *StakingContractSession) RegisterRestakingValidators(pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterRestakingValidators(&_StakingContract.TransactOpts, pubkeys, signatures)
}

// RegisterRestakingValidators is a paid mutator transaction binding the contract method 0x4e8ab1a1.
//
// Solidity: function registerRestakingValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_StakingContract *StakingContractTransactorSession) RegisterRestakingValidators(pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterRestakingValidators(&_StakingContract.TransactOpts, pubkeys, signatures)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x226f2645.
//
// Solidity: function registerValidator(bytes pubkey, bytes signature) returns()
func (_StakingContract *StakingContractTransactor) RegisterValidator(opts *bind.TransactOpts, pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "registerValidator", pubkey, signature)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x226f2645.
//
// Solidity: function registerValidator(bytes pubkey, bytes signature) returns()
func (_StakingContract *StakingContractSession) RegisterValidator(pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterValidator(&_StakingContract.TransactOpts, pubkey, signature)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x226f2645.
//
// Solidity: function registerValidator(bytes pubkey, bytes signature) returns()
func (_StakingContract *StakingContractTransactorSession) RegisterValidator(pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterValidator(&_StakingContract.TransactOpts, pubkey, signature)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xe7efb747.
//
// Solidity: function registerValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_StakingContract *StakingContractTransactor) RegisterValidators(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "registerValidators", pubkeys, signatures)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xe7efb747.
//
// Solidity: function registerValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_StakingContract *StakingContractSession) RegisterValidators(pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterValidators(&_StakingContract.TransactOpts, pubkeys, signatures)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xe7efb747.
//
// Solidity: function registerValidators(bytes[] pubkeys, bytes[] signatures) returns()
func (_StakingContract *StakingContractTransactorSession) RegisterValidators(pubkeys [][]byte, signatures [][]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.RegisterValidators(&_StakingContract.TransactOpts, pubkeys, signatures)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.RenounceRole(&_StakingContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.RenounceRole(&_StakingContract.TransactOpts, role, account)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0xc3d51709.
//
// Solidity: function replaceValidator(bytes oldpubkey, bytes pubkey, bytes signature) returns()
func (_StakingContract *StakingContractTransactor) ReplaceValidator(opts *bind.TransactOpts, oldpubkey []byte, pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "replaceValidator", oldpubkey, pubkey, signature)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0xc3d51709.
//
// Solidity: function replaceValidator(bytes oldpubkey, bytes pubkey, bytes signature) returns()
func (_StakingContract *StakingContractSession) ReplaceValidator(oldpubkey []byte, pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ReplaceValidator(&_StakingContract.TransactOpts, oldpubkey, pubkey, signature)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0xc3d51709.
//
// Solidity: function replaceValidator(bytes oldpubkey, bytes pubkey, bytes signature) returns()
func (_StakingContract *StakingContractTransactorSession) ReplaceValidator(oldpubkey []byte, pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ReplaceValidator(&_StakingContract.TransactOpts, oldpubkey, pubkey, signature)
}

// ReplaceValidators is a paid mutator transaction binding the contract method 0xfc65260a.
//
// Solidity: function replaceValidators(bytes[] oldpubkeys, bytes[] pubkeys, bytes[] signatures, bool restaking) returns()
func (_StakingContract *StakingContractTransactor) ReplaceValidators(opts *bind.TransactOpts, oldpubkeys [][]byte, pubkeys [][]byte, signatures [][]byte, restaking bool) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "replaceValidators", oldpubkeys, pubkeys, signatures, restaking)
}

// ReplaceValidators is a paid mutator transaction binding the contract method 0xfc65260a.
//
// Solidity: function replaceValidators(bytes[] oldpubkeys, bytes[] pubkeys, bytes[] signatures, bool restaking) returns()
func (_StakingContract *StakingContractSession) ReplaceValidators(oldpubkeys [][]byte, pubkeys [][]byte, signatures [][]byte, restaking bool) (*types.Transaction, error) {
	return _StakingContract.Contract.ReplaceValidators(&_StakingContract.TransactOpts, oldpubkeys, pubkeys, signatures, restaking)
}

// ReplaceValidators is a paid mutator transaction binding the contract method 0xfc65260a.
//
// Solidity: function replaceValidators(bytes[] oldpubkeys, bytes[] pubkeys, bytes[] signatures, bool restaking) returns()
func (_StakingContract *StakingContractTransactorSession) ReplaceValidators(oldpubkeys [][]byte, pubkeys [][]byte, signatures [][]byte, restaking bool) (*types.Transaction, error) {
	return _StakingContract.Contract.ReplaceValidators(&_StakingContract.TransactOpts, oldpubkeys, pubkeys, signatures, restaking)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.RevokeRole(&_StakingContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingContract *StakingContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.RevokeRole(&_StakingContract.TransactOpts, role, account)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_StakingContract *StakingContractTransactor) SetETHDepositContract(opts *bind.TransactOpts, _ethDepositContract common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "setETHDepositContract", _ethDepositContract)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_StakingContract *StakingContractSession) SetETHDepositContract(_ethDepositContract common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetETHDepositContract(&_StakingContract.TransactOpts, _ethDepositContract)
}

// SetETHDepositContract is a paid mutator transaction binding the contract method 0x91b66caa.
//
// Solidity: function setETHDepositContract(address _ethDepositContract) returns()
func (_StakingContract *StakingContractTransactorSession) SetETHDepositContract(_ethDepositContract common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetETHDepositContract(&_StakingContract.TransactOpts, _ethDepositContract)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_StakingContract *StakingContractTransactor) SetManagerFeeShare(opts *bind.TransactOpts, milli *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "setManagerFeeShare", milli)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_StakingContract *StakingContractSession) SetManagerFeeShare(milli *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.SetManagerFeeShare(&_StakingContract.TransactOpts, milli)
}

// SetManagerFeeShare is a paid mutator transaction binding the contract method 0x755d7dd3.
//
// Solidity: function setManagerFeeShare(uint256 milli) returns()
func (_StakingContract *StakingContractTransactorSession) SetManagerFeeShare(milli *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.SetManagerFeeShare(&_StakingContract.TransactOpts, milli)
}

// SetRedeemContract is a paid mutator transaction binding the contract method 0x6bb022c8.
//
// Solidity: function setRedeemContract(address _redeemContract) returns()
func (_StakingContract *StakingContractTransactor) SetRedeemContract(opts *bind.TransactOpts, _redeemContract common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "setRedeemContract", _redeemContract)
}

// SetRedeemContract is a paid mutator transaction binding the contract method 0x6bb022c8.
//
// Solidity: function setRedeemContract(address _redeemContract) returns()
func (_StakingContract *StakingContractSession) SetRedeemContract(_redeemContract common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetRedeemContract(&_StakingContract.TransactOpts, _redeemContract)
}

// SetRedeemContract is a paid mutator transaction binding the contract method 0x6bb022c8.
//
// Solidity: function setRedeemContract(address _redeemContract) returns()
func (_StakingContract *StakingContractTransactorSession) SetRedeemContract(_redeemContract common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetRedeemContract(&_StakingContract.TransactOpts, _redeemContract)
}

// SetWithdrawCredential is a paid mutator transaction binding the contract method 0xd7d25461.
//
// Solidity: function setWithdrawCredential(bytes32 withdrawalCredentials_) returns()
func (_StakingContract *StakingContractTransactor) SetWithdrawCredential(opts *bind.TransactOpts, withdrawalCredentials_ [32]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "setWithdrawCredential", withdrawalCredentials_)
}

// SetWithdrawCredential is a paid mutator transaction binding the contract method 0xd7d25461.
//
// Solidity: function setWithdrawCredential(bytes32 withdrawalCredentials_) returns()
func (_StakingContract *StakingContractSession) SetWithdrawCredential(withdrawalCredentials_ [32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.SetWithdrawCredential(&_StakingContract.TransactOpts, withdrawalCredentials_)
}

// SetWithdrawCredential is a paid mutator transaction binding the contract method 0xd7d25461.
//
// Solidity: function setWithdrawCredential(bytes32 withdrawalCredentials_) returns()
func (_StakingContract *StakingContractTransactorSession) SetWithdrawCredential(withdrawalCredentials_ [32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.SetWithdrawCredential(&_StakingContract.TransactOpts, withdrawalCredentials_)
}

// SetXETHContractAddress is a paid mutator transaction binding the contract method 0x72cda182.
//
// Solidity: function setXETHContractAddress(address _xETHContract) returns()
func (_StakingContract *StakingContractTransactor) SetXETHContractAddress(opts *bind.TransactOpts, _xETHContract common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "setXETHContractAddress", _xETHContract)
}

// SetXETHContractAddress is a paid mutator transaction binding the contract method 0x72cda182.
//
// Solidity: function setXETHContractAddress(address _xETHContract) returns()
func (_StakingContract *StakingContractSession) SetXETHContractAddress(_xETHContract common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetXETHContractAddress(&_StakingContract.TransactOpts, _xETHContract)
}

// SetXETHContractAddress is a paid mutator transaction binding the contract method 0x72cda182.
//
// Solidity: function setXETHContractAddress(address _xETHContract) returns()
func (_StakingContract *StakingContractTransactorSession) SetXETHContractAddress(_xETHContract common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.SetXETHContractAddress(&_StakingContract.TransactOpts, _xETHContract)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_StakingContract *StakingContractTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_StakingContract *StakingContractSession) Stake() (*types.Transaction, error) {
	return _StakingContract.Contract.Stake(&_StakingContract.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_StakingContract *StakingContractTransactorSession) Stake() (*types.Transaction, error) {
	return _StakingContract.Contract.Stake(&_StakingContract.TransactOpts)
}

// SwitchPhase is a paid mutator transaction binding the contract method 0x63e54717.
//
// Solidity: function switchPhase(uint256 newPhase) returns()
func (_StakingContract *StakingContractTransactor) SwitchPhase(opts *bind.TransactOpts, newPhase *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "switchPhase", newPhase)
}

// SwitchPhase is a paid mutator transaction binding the contract method 0x63e54717.
//
// Solidity: function switchPhase(uint256 newPhase) returns()
func (_StakingContract *StakingContractSession) SwitchPhase(newPhase *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.SwitchPhase(&_StakingContract.TransactOpts, newPhase)
}

// SwitchPhase is a paid mutator transaction binding the contract method 0x63e54717.
//
// Solidity: function switchPhase(uint256 newPhase) returns()
func (_StakingContract *StakingContractTransactorSession) SwitchPhase(newPhase *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.SwitchPhase(&_StakingContract.TransactOpts, newPhase)
}

// SyncBalance is a paid mutator transaction binding the contract method 0xfd9c652b.
//
// Solidity: function syncBalance() returns()
func (_StakingContract *StakingContractTransactor) SyncBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "syncBalance")
}

// SyncBalance is a paid mutator transaction binding the contract method 0xfd9c652b.
//
// Solidity: function syncBalance() returns()
func (_StakingContract *StakingContractSession) SyncBalance() (*types.Transaction, error) {
	return _StakingContract.Contract.SyncBalance(&_StakingContract.TransactOpts)
}

// SyncBalance is a paid mutator transaction binding the contract method 0xfd9c652b.
//
// Solidity: function syncBalance() returns()
func (_StakingContract *StakingContractTransactorSession) SyncBalance() (*types.Transaction, error) {
	return _StakingContract.Contract.SyncBalance(&_StakingContract.TransactOpts)
}

// ToggleAutoCompound is a paid mutator transaction binding the contract method 0x89d00678.
//
// Solidity: function toggleAutoCompound() returns()
func (_StakingContract *StakingContractTransactor) ToggleAutoCompound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "toggleAutoCompound")
}

// ToggleAutoCompound is a paid mutator transaction binding the contract method 0x89d00678.
//
// Solidity: function toggleAutoCompound() returns()
func (_StakingContract *StakingContractSession) ToggleAutoCompound() (*types.Transaction, error) {
	return _StakingContract.Contract.ToggleAutoCompound(&_StakingContract.TransactOpts)
}

// ToggleAutoCompound is a paid mutator transaction binding the contract method 0x89d00678.
//
// Solidity: function toggleAutoCompound() returns()
func (_StakingContract *StakingContractTransactorSession) ToggleAutoCompound() (*types.Transaction, error) {
	return _StakingContract.Contract.ToggleAutoCompound(&_StakingContract.TransactOpts)
}

// ToggleWhiteList is a paid mutator transaction binding the contract method 0x4f431a6f.
//
// Solidity: function toggleWhiteList(address account) returns()
func (_StakingContract *StakingContractTransactor) ToggleWhiteList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "toggleWhiteList", account)
}

// ToggleWhiteList is a paid mutator transaction binding the contract method 0x4f431a6f.
//
// Solidity: function toggleWhiteList(address account) returns()
func (_StakingContract *StakingContractSession) ToggleWhiteList(account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.ToggleWhiteList(&_StakingContract.TransactOpts, account)
}

// ToggleWhiteList is a paid mutator transaction binding the contract method 0x4f431a6f.
//
// Solidity: function toggleWhiteList(address account) returns()
func (_StakingContract *StakingContractTransactorSession) ToggleWhiteList(account common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.ToggleWhiteList(&_StakingContract.TransactOpts, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractSession) Unpause() (*types.Transaction, error) {
	return _StakingContract.Contract.Unpause(&_StakingContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakingContract.Contract.Unpause(&_StakingContract.TransactOpts)
}

// ValidatorSlashedStop is a paid mutator transaction binding the contract method 0x5662f706.
//
// Solidity: function validatorSlashedStop(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_StakingContract *StakingContractTransactor) ValidatorSlashedStop(opts *bind.TransactOpts, _stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "validatorSlashedStop", _stoppedPubKeys, clock)
}

// ValidatorSlashedStop is a paid mutator transaction binding the contract method 0x5662f706.
//
// Solidity: function validatorSlashedStop(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_StakingContract *StakingContractSession) ValidatorSlashedStop(_stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ValidatorSlashedStop(&_StakingContract.TransactOpts, _stoppedPubKeys, clock)
}

// ValidatorSlashedStop is a paid mutator transaction binding the contract method 0x5662f706.
//
// Solidity: function validatorSlashedStop(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_StakingContract *StakingContractTransactorSession) ValidatorSlashedStop(_stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ValidatorSlashedStop(&_StakingContract.TransactOpts, _stoppedPubKeys, clock)
}

// ValidatorStopped is a paid mutator transaction binding the contract method 0x99629f58.
//
// Solidity: function validatorStopped(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_StakingContract *StakingContractTransactor) ValidatorStopped(opts *bind.TransactOpts, _stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "validatorStopped", _stoppedPubKeys, clock)
}

// ValidatorStopped is a paid mutator transaction binding the contract method 0x99629f58.
//
// Solidity: function validatorStopped(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_StakingContract *StakingContractSession) ValidatorStopped(_stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ValidatorStopped(&_StakingContract.TransactOpts, _stoppedPubKeys, clock)
}

// ValidatorStopped is a paid mutator transaction binding the contract method 0x99629f58.
//
// Solidity: function validatorStopped(bytes[] _stoppedPubKeys, bytes32 clock) returns()
func (_StakingContract *StakingContractTransactorSession) ValidatorStopped(_stoppedPubKeys [][]byte, clock [32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ValidatorStopped(&_StakingContract.TransactOpts, _stoppedPubKeys, clock)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address to) returns()
func (_StakingContract *StakingContractTransactor) WithdrawManagerFee(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "withdrawManagerFee", amount, to)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address to) returns()
func (_StakingContract *StakingContractSession) WithdrawManagerFee(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.WithdrawManagerFee(&_StakingContract.TransactOpts, amount, to)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address to) returns()
func (_StakingContract *StakingContractTransactorSession) WithdrawManagerFee(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.WithdrawManagerFee(&_StakingContract.TransactOpts, amount, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingContract *StakingContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingContract *StakingContractSession) Receive() (*types.Transaction, error) {
	return _StakingContract.Contract.Receive(&_StakingContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingContract *StakingContractTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingContract.Contract.Receive(&_StakingContract.TransactOpts)
}

// StakingContractAutoCompoundToggleIterator is returned from FilterAutoCompoundToggle and is used to iterate over the raw logs and unpacked data for AutoCompoundToggle events raised by the StakingContract contract.
type StakingContractAutoCompoundToggleIterator struct {
	Event *StakingContractAutoCompoundToggle // Event containing the contract specifics and raw log

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
func (it *StakingContractAutoCompoundToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractAutoCompoundToggle)
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
		it.Event = new(StakingContractAutoCompoundToggle)
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
func (it *StakingContractAutoCompoundToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractAutoCompoundToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractAutoCompoundToggle represents a AutoCompoundToggle event raised by the StakingContract contract.
type StakingContractAutoCompoundToggle struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAutoCompoundToggle is a free log retrieval operation binding the contract event 0x0e1bdbe9c15a50fbc7d124adb288d201acc7b9a8a535b003ca6ba80b7c9f5335.
//
// Solidity: event AutoCompoundToggle(bool enabled)
func (_StakingContract *StakingContractFilterer) FilterAutoCompoundToggle(opts *bind.FilterOpts) (*StakingContractAutoCompoundToggleIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "AutoCompoundToggle")
	if err != nil {
		return nil, err
	}
	return &StakingContractAutoCompoundToggleIterator{contract: _StakingContract.contract, event: "AutoCompoundToggle", logs: logs, sub: sub}, nil
}

// WatchAutoCompoundToggle is a free log subscription operation binding the contract event 0x0e1bdbe9c15a50fbc7d124adb288d201acc7b9a8a535b003ca6ba80b7c9f5335.
//
// Solidity: event AutoCompoundToggle(bool enabled)
func (_StakingContract *StakingContractFilterer) WatchAutoCompoundToggle(opts *bind.WatchOpts, sink chan<- *StakingContractAutoCompoundToggle) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "AutoCompoundToggle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractAutoCompoundToggle)
				if err := _StakingContract.contract.UnpackLog(event, "AutoCompoundToggle", log); err != nil {
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

// ParseAutoCompoundToggle is a log parse operation binding the contract event 0x0e1bdbe9c15a50fbc7d124adb288d201acc7b9a8a535b003ca6ba80b7c9f5335.
//
// Solidity: event AutoCompoundToggle(bool enabled)
func (_StakingContract *StakingContractFilterer) ParseAutoCompoundToggle(log types.Log) (*StakingContractAutoCompoundToggle, error) {
	event := new(StakingContractAutoCompoundToggle)
	if err := _StakingContract.contract.UnpackLog(event, "AutoCompoundToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractBalanceSyncedIterator is returned from FilterBalanceSynced and is used to iterate over the raw logs and unpacked data for BalanceSynced events raised by the StakingContract contract.
type StakingContractBalanceSyncedIterator struct {
	Event *StakingContractBalanceSynced // Event containing the contract specifics and raw log

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
func (it *StakingContractBalanceSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractBalanceSynced)
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
		it.Event = new(StakingContractBalanceSynced)
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
func (it *StakingContractBalanceSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractBalanceSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractBalanceSynced represents a BalanceSynced event raised by the StakingContract contract.
type StakingContractBalanceSynced struct {
	Diff *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBalanceSynced is a free log retrieval operation binding the contract event 0xe7948c33eb604391785037114655100edf93283c25b69884e9238ae197f07817.
//
// Solidity: event BalanceSynced(uint256 diff)
func (_StakingContract *StakingContractFilterer) FilterBalanceSynced(opts *bind.FilterOpts) (*StakingContractBalanceSyncedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "BalanceSynced")
	if err != nil {
		return nil, err
	}
	return &StakingContractBalanceSyncedIterator{contract: _StakingContract.contract, event: "BalanceSynced", logs: logs, sub: sub}, nil
}

// WatchBalanceSynced is a free log subscription operation binding the contract event 0xe7948c33eb604391785037114655100edf93283c25b69884e9238ae197f07817.
//
// Solidity: event BalanceSynced(uint256 diff)
func (_StakingContract *StakingContractFilterer) WatchBalanceSynced(opts *bind.WatchOpts, sink chan<- *StakingContractBalanceSynced) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "BalanceSynced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractBalanceSynced)
				if err := _StakingContract.contract.UnpackLog(event, "BalanceSynced", log); err != nil {
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

// ParseBalanceSynced is a log parse operation binding the contract event 0xe7948c33eb604391785037114655100edf93283c25b69884e9238ae197f07817.
//
// Solidity: event BalanceSynced(uint256 diff)
func (_StakingContract *StakingContractFilterer) ParseBalanceSynced(log types.Log) (*StakingContractBalanceSynced, error) {
	event := new(StakingContractBalanceSynced)
	if err := _StakingContract.contract.UnpackLog(event, "BalanceSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractDebtQueuedIterator is returned from FilterDebtQueued and is used to iterate over the raw logs and unpacked data for DebtQueued events raised by the StakingContract contract.
type StakingContractDebtQueuedIterator struct {
	Event *StakingContractDebtQueued // Event containing the contract specifics and raw log

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
func (it *StakingContractDebtQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractDebtQueued)
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
		it.Event = new(StakingContractDebtQueued)
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
func (it *StakingContractDebtQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractDebtQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractDebtQueued represents a DebtQueued event raised by the StakingContract contract.
type StakingContractDebtQueued struct {
	Creditor    common.Address
	AmountEther *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDebtQueued is a free log retrieval operation binding the contract event 0x889f3b08db1ce169841b4f7e2aadfe4298088b1bce57b31fecae3260dd435829.
//
// Solidity: event DebtQueued(address creditor, uint256 amountEther)
func (_StakingContract *StakingContractFilterer) FilterDebtQueued(opts *bind.FilterOpts) (*StakingContractDebtQueuedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "DebtQueued")
	if err != nil {
		return nil, err
	}
	return &StakingContractDebtQueuedIterator{contract: _StakingContract.contract, event: "DebtQueued", logs: logs, sub: sub}, nil
}

// WatchDebtQueued is a free log subscription operation binding the contract event 0x889f3b08db1ce169841b4f7e2aadfe4298088b1bce57b31fecae3260dd435829.
//
// Solidity: event DebtQueued(address creditor, uint256 amountEther)
func (_StakingContract *StakingContractFilterer) WatchDebtQueued(opts *bind.WatchOpts, sink chan<- *StakingContractDebtQueued) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "DebtQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractDebtQueued)
				if err := _StakingContract.contract.UnpackLog(event, "DebtQueued", log); err != nil {
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

// ParseDebtQueued is a log parse operation binding the contract event 0x889f3b08db1ce169841b4f7e2aadfe4298088b1bce57b31fecae3260dd435829.
//
// Solidity: event DebtQueued(address creditor, uint256 amountEther)
func (_StakingContract *StakingContractFilterer) ParseDebtQueued(log types.Log) (*StakingContractDebtQueued, error) {
	event := new(StakingContractDebtQueued)
	if err := _StakingContract.contract.UnpackLog(event, "DebtQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractDepositContractSetIterator is returned from FilterDepositContractSet and is used to iterate over the raw logs and unpacked data for DepositContractSet events raised by the StakingContract contract.
type StakingContractDepositContractSetIterator struct {
	Event *StakingContractDepositContractSet // Event containing the contract specifics and raw log

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
func (it *StakingContractDepositContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractDepositContractSet)
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
		it.Event = new(StakingContractDepositContractSet)
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
func (it *StakingContractDepositContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractDepositContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractDepositContractSet represents a DepositContractSet event raised by the StakingContract contract.
type StakingContractDepositContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDepositContractSet is a free log retrieval operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_StakingContract *StakingContractFilterer) FilterDepositContractSet(opts *bind.FilterOpts) (*StakingContractDepositContractSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "DepositContractSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractDepositContractSetIterator{contract: _StakingContract.contract, event: "DepositContractSet", logs: logs, sub: sub}, nil
}

// WatchDepositContractSet is a free log subscription operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_StakingContract *StakingContractFilterer) WatchDepositContractSet(opts *bind.WatchOpts, sink chan<- *StakingContractDepositContractSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "DepositContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractDepositContractSet)
				if err := _StakingContract.contract.UnpackLog(event, "DepositContractSet", log); err != nil {
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

// ParseDepositContractSet is a log parse operation binding the contract event 0x1781ac9526b978975dba0fd26a33e044a55a7ace054a3ee7efa5f8459513bead.
//
// Solidity: event DepositContractSet(address addr)
func (_StakingContract *StakingContractFilterer) ParseDepositContractSet(log types.Log) (*StakingContractDepositContractSet, error) {
	event := new(StakingContractDepositContractSet)
	if err := _StakingContract.contract.UnpackLog(event, "DepositContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingContract contract.
type StakingContractInitializedIterator struct {
	Event *StakingContractInitialized // Event containing the contract specifics and raw log

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
func (it *StakingContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractInitialized)
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
		it.Event = new(StakingContractInitialized)
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
func (it *StakingContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractInitialized represents a Initialized event raised by the StakingContract contract.
type StakingContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingContract *StakingContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingContractInitializedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingContractInitializedIterator{contract: _StakingContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingContract *StakingContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingContractInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractInitialized)
				if err := _StakingContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseInitialized(log types.Log) (*StakingContractInitialized, error) {
	event := new(StakingContractInitialized)
	if err := _StakingContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractManagerAccountSetIterator is returned from FilterManagerAccountSet and is used to iterate over the raw logs and unpacked data for ManagerAccountSet events raised by the StakingContract contract.
type StakingContractManagerAccountSetIterator struct {
	Event *StakingContractManagerAccountSet // Event containing the contract specifics and raw log

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
func (it *StakingContractManagerAccountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractManagerAccountSet)
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
		it.Event = new(StakingContractManagerAccountSet)
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
func (it *StakingContractManagerAccountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractManagerAccountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractManagerAccountSet represents a ManagerAccountSet event raised by the StakingContract contract.
type StakingContractManagerAccountSet struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAccountSet is a free log retrieval operation binding the contract event 0x3aa9ccc8285bfd4a1ad6b4cbde39a3b91da051237323010b14c1fa4130be9b7f.
//
// Solidity: event ManagerAccountSet(address account)
func (_StakingContract *StakingContractFilterer) FilterManagerAccountSet(opts *bind.FilterOpts) (*StakingContractManagerAccountSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "ManagerAccountSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractManagerAccountSetIterator{contract: _StakingContract.contract, event: "ManagerAccountSet", logs: logs, sub: sub}, nil
}

// WatchManagerAccountSet is a free log subscription operation binding the contract event 0x3aa9ccc8285bfd4a1ad6b4cbde39a3b91da051237323010b14c1fa4130be9b7f.
//
// Solidity: event ManagerAccountSet(address account)
func (_StakingContract *StakingContractFilterer) WatchManagerAccountSet(opts *bind.WatchOpts, sink chan<- *StakingContractManagerAccountSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "ManagerAccountSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractManagerAccountSet)
				if err := _StakingContract.contract.UnpackLog(event, "ManagerAccountSet", log); err != nil {
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

// ParseManagerAccountSet is a log parse operation binding the contract event 0x3aa9ccc8285bfd4a1ad6b4cbde39a3b91da051237323010b14c1fa4130be9b7f.
//
// Solidity: event ManagerAccountSet(address account)
func (_StakingContract *StakingContractFilterer) ParseManagerAccountSet(log types.Log) (*StakingContractManagerAccountSet, error) {
	event := new(StakingContractManagerAccountSet)
	if err := _StakingContract.contract.UnpackLog(event, "ManagerAccountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractManagerFeeSetIterator is returned from FilterManagerFeeSet and is used to iterate over the raw logs and unpacked data for ManagerFeeSet events raised by the StakingContract contract.
type StakingContractManagerFeeSetIterator struct {
	Event *StakingContractManagerFeeSet // Event containing the contract specifics and raw log

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
func (it *StakingContractManagerFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractManagerFeeSet)
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
		it.Event = new(StakingContractManagerFeeSet)
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
func (it *StakingContractManagerFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractManagerFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractManagerFeeSet represents a ManagerFeeSet event raised by the StakingContract contract.
type StakingContractManagerFeeSet struct {
	Milli *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeSet is a free log retrieval operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_StakingContract *StakingContractFilterer) FilterManagerFeeSet(opts *bind.FilterOpts) (*StakingContractManagerFeeSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "ManagerFeeSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractManagerFeeSetIterator{contract: _StakingContract.contract, event: "ManagerFeeSet", logs: logs, sub: sub}, nil
}

// WatchManagerFeeSet is a free log subscription operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_StakingContract *StakingContractFilterer) WatchManagerFeeSet(opts *bind.WatchOpts, sink chan<- *StakingContractManagerFeeSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "ManagerFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractManagerFeeSet)
				if err := _StakingContract.contract.UnpackLog(event, "ManagerFeeSet", log); err != nil {
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

// ParseManagerFeeSet is a log parse operation binding the contract event 0x4de90ec86e1bc56c192e2399bacbd10bdaba720caca606354d66c5cb33d6802b.
//
// Solidity: event ManagerFeeSet(uint256 milli)
func (_StakingContract *StakingContractFilterer) ParseManagerFeeSet(log types.Log) (*StakingContractManagerFeeSet, error) {
	event := new(StakingContractManagerFeeSet)
	if err := _StakingContract.contract.UnpackLog(event, "ManagerFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractManagerFeeWithdrawedIterator is returned from FilterManagerFeeWithdrawed and is used to iterate over the raw logs and unpacked data for ManagerFeeWithdrawed events raised by the StakingContract contract.
type StakingContractManagerFeeWithdrawedIterator struct {
	Event *StakingContractManagerFeeWithdrawed // Event containing the contract specifics and raw log

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
func (it *StakingContractManagerFeeWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractManagerFeeWithdrawed)
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
		it.Event = new(StakingContractManagerFeeWithdrawed)
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
func (it *StakingContractManagerFeeWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractManagerFeeWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractManagerFeeWithdrawed represents a ManagerFeeWithdrawed event raised by the StakingContract contract.
type StakingContractManagerFeeWithdrawed struct {
	Amount *big.Int
	Arg1   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeWithdrawed is a free log retrieval operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address arg1)
func (_StakingContract *StakingContractFilterer) FilterManagerFeeWithdrawed(opts *bind.FilterOpts) (*StakingContractManagerFeeWithdrawedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return &StakingContractManagerFeeWithdrawedIterator{contract: _StakingContract.contract, event: "ManagerFeeWithdrawed", logs: logs, sub: sub}, nil
}

// WatchManagerFeeWithdrawed is a free log subscription operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address arg1)
func (_StakingContract *StakingContractFilterer) WatchManagerFeeWithdrawed(opts *bind.WatchOpts, sink chan<- *StakingContractManagerFeeWithdrawed) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractManagerFeeWithdrawed)
				if err := _StakingContract.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
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

// ParseManagerFeeWithdrawed is a log parse operation binding the contract event 0x2425aa1fadefc5c570850aa9c9e3dfa4fc6b43ccd1c05b47db38dd6518a743b3.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, address arg1)
func (_StakingContract *StakingContractFilterer) ParseManagerFeeWithdrawed(log types.Log) (*StakingContractManagerFeeWithdrawed, error) {
	event := new(StakingContractManagerFeeWithdrawed)
	if err := _StakingContract.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingContract contract.
type StakingContractPausedIterator struct {
	Event *StakingContractPaused // Event containing the contract specifics and raw log

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
func (it *StakingContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractPaused)
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
		it.Event = new(StakingContractPaused)
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
func (it *StakingContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractPaused represents a Paused event raised by the StakingContract contract.
type StakingContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingContract *StakingContractFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingContractPausedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingContractPausedIterator{contract: _StakingContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingContract *StakingContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingContractPaused) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractPaused)
				if err := _StakingContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParsePaused(log types.Log) (*StakingContractPaused, error) {
	event := new(StakingContractPaused)
	if err := _StakingContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRedeemContractSetIterator is returned from FilterRedeemContractSet and is used to iterate over the raw logs and unpacked data for RedeemContractSet events raised by the StakingContract contract.
type StakingContractRedeemContractSetIterator struct {
	Event *StakingContractRedeemContractSet // Event containing the contract specifics and raw log

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
func (it *StakingContractRedeemContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRedeemContractSet)
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
		it.Event = new(StakingContractRedeemContractSet)
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
func (it *StakingContractRedeemContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRedeemContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRedeemContractSet represents a RedeemContractSet event raised by the StakingContract contract.
type StakingContractRedeemContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRedeemContractSet is a free log retrieval operation binding the contract event 0x4d3d3f5d1d8423a4e40b3fcfccaf0c1b18ca550a8592d5e38a61765931a9cf78.
//
// Solidity: event RedeemContractSet(address addr)
func (_StakingContract *StakingContractFilterer) FilterRedeemContractSet(opts *bind.FilterOpts) (*StakingContractRedeemContractSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RedeemContractSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractRedeemContractSetIterator{contract: _StakingContract.contract, event: "RedeemContractSet", logs: logs, sub: sub}, nil
}

// WatchRedeemContractSet is a free log subscription operation binding the contract event 0x4d3d3f5d1d8423a4e40b3fcfccaf0c1b18ca550a8592d5e38a61765931a9cf78.
//
// Solidity: event RedeemContractSet(address addr)
func (_StakingContract *StakingContractFilterer) WatchRedeemContractSet(opts *bind.WatchOpts, sink chan<- *StakingContractRedeemContractSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RedeemContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRedeemContractSet)
				if err := _StakingContract.contract.UnpackLog(event, "RedeemContractSet", log); err != nil {
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

// ParseRedeemContractSet is a log parse operation binding the contract event 0x4d3d3f5d1d8423a4e40b3fcfccaf0c1b18ca550a8592d5e38a61765931a9cf78.
//
// Solidity: event RedeemContractSet(address addr)
func (_StakingContract *StakingContractFilterer) ParseRedeemContractSet(log types.Log) (*StakingContractRedeemContractSet, error) {
	event := new(StakingContractRedeemContractSet)
	if err := _StakingContract.contract.UnpackLog(event, "RedeemContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRestakingAddressSetIterator is returned from FilterRestakingAddressSet and is used to iterate over the raw logs and unpacked data for RestakingAddressSet events raised by the StakingContract contract.
type StakingContractRestakingAddressSetIterator struct {
	Event *StakingContractRestakingAddressSet // Event containing the contract specifics and raw log

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
func (it *StakingContractRestakingAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRestakingAddressSet)
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
		it.Event = new(StakingContractRestakingAddressSet)
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
func (it *StakingContractRestakingAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRestakingAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRestakingAddressSet represents a RestakingAddressSet event raised by the StakingContract contract.
type StakingContractRestakingAddressSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRestakingAddressSet is a free log retrieval operation binding the contract event 0x834a16436cbc9545643638861484439875a76defa59e448db456c67d36e4e417.
//
// Solidity: event RestakingAddressSet(address addr)
func (_StakingContract *StakingContractFilterer) FilterRestakingAddressSet(opts *bind.FilterOpts) (*StakingContractRestakingAddressSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RestakingAddressSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractRestakingAddressSetIterator{contract: _StakingContract.contract, event: "RestakingAddressSet", logs: logs, sub: sub}, nil
}

// WatchRestakingAddressSet is a free log subscription operation binding the contract event 0x834a16436cbc9545643638861484439875a76defa59e448db456c67d36e4e417.
//
// Solidity: event RestakingAddressSet(address addr)
func (_StakingContract *StakingContractFilterer) WatchRestakingAddressSet(opts *bind.WatchOpts, sink chan<- *StakingContractRestakingAddressSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RestakingAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRestakingAddressSet)
				if err := _StakingContract.contract.UnpackLog(event, "RestakingAddressSet", log); err != nil {
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

// ParseRestakingAddressSet is a log parse operation binding the contract event 0x834a16436cbc9545643638861484439875a76defa59e448db456c67d36e4e417.
//
// Solidity: event RestakingAddressSet(address addr)
func (_StakingContract *StakingContractFilterer) ParseRestakingAddressSet(log types.Log) (*StakingContractRestakingAddressSet, error) {
	event := new(StakingContractRestakingAddressSet)
	if err := _StakingContract.contract.UnpackLog(event, "RestakingAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRevenueAccountedIterator is returned from FilterRevenueAccounted and is used to iterate over the raw logs and unpacked data for RevenueAccounted events raised by the StakingContract contract.
type StakingContractRevenueAccountedIterator struct {
	Event *StakingContractRevenueAccounted // Event containing the contract specifics and raw log

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
func (it *StakingContractRevenueAccountedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRevenueAccounted)
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
		it.Event = new(StakingContractRevenueAccounted)
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
func (it *StakingContractRevenueAccountedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRevenueAccountedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRevenueAccounted represents a RevenueAccounted event raised by the StakingContract contract.
type StakingContractRevenueAccounted struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevenueAccounted is a free log retrieval operation binding the contract event 0x82f24840c0f58d92529afdd441950ddc6e8f2d60138d4458a8d74ba367540cda.
//
// Solidity: event RevenueAccounted(uint256 amount)
func (_StakingContract *StakingContractFilterer) FilterRevenueAccounted(opts *bind.FilterOpts) (*StakingContractRevenueAccountedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RevenueAccounted")
	if err != nil {
		return nil, err
	}
	return &StakingContractRevenueAccountedIterator{contract: _StakingContract.contract, event: "RevenueAccounted", logs: logs, sub: sub}, nil
}

// WatchRevenueAccounted is a free log subscription operation binding the contract event 0x82f24840c0f58d92529afdd441950ddc6e8f2d60138d4458a8d74ba367540cda.
//
// Solidity: event RevenueAccounted(uint256 amount)
func (_StakingContract *StakingContractFilterer) WatchRevenueAccounted(opts *bind.WatchOpts, sink chan<- *StakingContractRevenueAccounted) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RevenueAccounted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRevenueAccounted)
				if err := _StakingContract.contract.UnpackLog(event, "RevenueAccounted", log); err != nil {
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

// ParseRevenueAccounted is a log parse operation binding the contract event 0x82f24840c0f58d92529afdd441950ddc6e8f2d60138d4458a8d74ba367540cda.
//
// Solidity: event RevenueAccounted(uint256 amount)
func (_StakingContract *StakingContractFilterer) ParseRevenueAccounted(log types.Log) (*StakingContractRevenueAccounted, error) {
	event := new(StakingContractRevenueAccounted)
	if err := _StakingContract.contract.UnpackLog(event, "RevenueAccounted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingContract contract.
type StakingContractRoleAdminChangedIterator struct {
	Event *StakingContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRoleAdminChanged)
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
		it.Event = new(StakingContractRoleAdminChanged)
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
func (it *StakingContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRoleAdminChanged represents a RoleAdminChanged event raised by the StakingContract contract.
type StakingContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingContract *StakingContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractRoleAdminChangedIterator{contract: _StakingContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingContract *StakingContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRoleAdminChanged)
				if err := _StakingContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseRoleAdminChanged(log types.Log) (*StakingContractRoleAdminChanged, error) {
	event := new(StakingContractRoleAdminChanged)
	if err := _StakingContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingContract contract.
type StakingContractRoleGrantedIterator struct {
	Event *StakingContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRoleGranted)
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
		it.Event = new(StakingContractRoleGranted)
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
func (it *StakingContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRoleGranted represents a RoleGranted event raised by the StakingContract contract.
type StakingContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingContract *StakingContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingContractRoleGrantedIterator, error) {

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

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractRoleGrantedIterator{contract: _StakingContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingContract *StakingContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRoleGranted)
				if err := _StakingContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseRoleGranted(log types.Log) (*StakingContractRoleGranted, error) {
	event := new(StakingContractRoleGranted)
	if err := _StakingContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingContract contract.
type StakingContractRoleRevokedIterator struct {
	Event *StakingContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRoleRevoked)
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
		it.Event = new(StakingContractRoleRevoked)
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
func (it *StakingContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRoleRevoked represents a RoleRevoked event raised by the StakingContract contract.
type StakingContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingContract *StakingContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingContractRoleRevokedIterator, error) {

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

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractRoleRevokedIterator{contract: _StakingContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingContract *StakingContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRoleRevoked)
				if err := _StakingContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseRoleRevoked(log types.Log) (*StakingContractRoleRevoked, error) {
	event := new(StakingContractRoleRevoked)
	if err := _StakingContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingContract contract.
type StakingContractUnpausedIterator struct {
	Event *StakingContractUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractUnpaused)
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
		it.Event = new(StakingContractUnpaused)
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
func (it *StakingContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractUnpaused represents a Unpaused event raised by the StakingContract contract.
type StakingContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingContract *StakingContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingContractUnpausedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingContractUnpausedIterator{contract: _StakingContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingContract *StakingContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractUnpaused)
				if err := _StakingContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StakingContract *StakingContractFilterer) ParseUnpaused(log types.Log) (*StakingContractUnpaused, error) {
	event := new(StakingContractUnpaused)
	if err := _StakingContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractValidatorActivatedIterator is returned from FilterValidatorActivated and is used to iterate over the raw logs and unpacked data for ValidatorActivated events raised by the StakingContract contract.
type StakingContractValidatorActivatedIterator struct {
	Event *StakingContractValidatorActivated // Event containing the contract specifics and raw log

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
func (it *StakingContractValidatorActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractValidatorActivated)
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
		it.Event = new(StakingContractValidatorActivated)
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
func (it *StakingContractValidatorActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractValidatorActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractValidatorActivated represents a ValidatorActivated event raised by the StakingContract contract.
type StakingContractValidatorActivated struct {
	NextValidatorId *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorActivated is a free log retrieval operation binding the contract event 0xe2a191ee805447bcf5adabadd39cb816b1b46de1364263aef69980bdafd8370f.
//
// Solidity: event ValidatorActivated(uint256 nextValidatorId)
func (_StakingContract *StakingContractFilterer) FilterValidatorActivated(opts *bind.FilterOpts) (*StakingContractValidatorActivatedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "ValidatorActivated")
	if err != nil {
		return nil, err
	}
	return &StakingContractValidatorActivatedIterator{contract: _StakingContract.contract, event: "ValidatorActivated", logs: logs, sub: sub}, nil
}

// WatchValidatorActivated is a free log subscription operation binding the contract event 0xe2a191ee805447bcf5adabadd39cb816b1b46de1364263aef69980bdafd8370f.
//
// Solidity: event ValidatorActivated(uint256 nextValidatorId)
func (_StakingContract *StakingContractFilterer) WatchValidatorActivated(opts *bind.WatchOpts, sink chan<- *StakingContractValidatorActivated) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "ValidatorActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractValidatorActivated)
				if err := _StakingContract.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
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

// ParseValidatorActivated is a log parse operation binding the contract event 0xe2a191ee805447bcf5adabadd39cb816b1b46de1364263aef69980bdafd8370f.
//
// Solidity: event ValidatorActivated(uint256 nextValidatorId)
func (_StakingContract *StakingContractFilterer) ParseValidatorActivated(log types.Log) (*StakingContractValidatorActivated, error) {
	event := new(StakingContractValidatorActivated)
	if err := _StakingContract.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractValidatorSlashedStoppedIterator is returned from FilterValidatorSlashedStopped and is used to iterate over the raw logs and unpacked data for ValidatorSlashedStopped events raised by the StakingContract contract.
type StakingContractValidatorSlashedStoppedIterator struct {
	Event *StakingContractValidatorSlashedStopped // Event containing the contract specifics and raw log

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
func (it *StakingContractValidatorSlashedStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractValidatorSlashedStopped)
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
		it.Event = new(StakingContractValidatorSlashedStopped)
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
func (it *StakingContractValidatorSlashedStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractValidatorSlashedStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractValidatorSlashedStopped represents a ValidatorSlashedStopped event raised by the StakingContract contract.
type StakingContractValidatorSlashedStopped struct {
	StoppedCount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashedStopped is a free log retrieval operation binding the contract event 0x0c5c941f6f7b3b7b1624e8a6738cdc96705c25510de1f11ff37671927a5c47c0.
//
// Solidity: event ValidatorSlashedStopped(uint256 stoppedCount)
func (_StakingContract *StakingContractFilterer) FilterValidatorSlashedStopped(opts *bind.FilterOpts) (*StakingContractValidatorSlashedStoppedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "ValidatorSlashedStopped")
	if err != nil {
		return nil, err
	}
	return &StakingContractValidatorSlashedStoppedIterator{contract: _StakingContract.contract, event: "ValidatorSlashedStopped", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashedStopped is a free log subscription operation binding the contract event 0x0c5c941f6f7b3b7b1624e8a6738cdc96705c25510de1f11ff37671927a5c47c0.
//
// Solidity: event ValidatorSlashedStopped(uint256 stoppedCount)
func (_StakingContract *StakingContractFilterer) WatchValidatorSlashedStopped(opts *bind.WatchOpts, sink chan<- *StakingContractValidatorSlashedStopped) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "ValidatorSlashedStopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractValidatorSlashedStopped)
				if err := _StakingContract.contract.UnpackLog(event, "ValidatorSlashedStopped", log); err != nil {
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

// ParseValidatorSlashedStopped is a log parse operation binding the contract event 0x0c5c941f6f7b3b7b1624e8a6738cdc96705c25510de1f11ff37671927a5c47c0.
//
// Solidity: event ValidatorSlashedStopped(uint256 stoppedCount)
func (_StakingContract *StakingContractFilterer) ParseValidatorSlashedStopped(log types.Log) (*StakingContractValidatorSlashedStopped, error) {
	event := new(StakingContractValidatorSlashedStopped)
	if err := _StakingContract.contract.UnpackLog(event, "ValidatorSlashedStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractValidatorStoppedIterator is returned from FilterValidatorStopped and is used to iterate over the raw logs and unpacked data for ValidatorStopped events raised by the StakingContract contract.
type StakingContractValidatorStoppedIterator struct {
	Event *StakingContractValidatorStopped // Event containing the contract specifics and raw log

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
func (it *StakingContractValidatorStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractValidatorStopped)
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
		it.Event = new(StakingContractValidatorStopped)
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
func (it *StakingContractValidatorStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractValidatorStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractValidatorStopped represents a ValidatorStopped event raised by the StakingContract contract.
type StakingContractValidatorStopped struct {
	StoppedCount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorStopped is a free log retrieval operation binding the contract event 0xf25558665a382a9abb684f20b20021df5923b51485bbf2829ff0089b5b271410.
//
// Solidity: event ValidatorStopped(uint256 stoppedCount)
func (_StakingContract *StakingContractFilterer) FilterValidatorStopped(opts *bind.FilterOpts) (*StakingContractValidatorStoppedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "ValidatorStopped")
	if err != nil {
		return nil, err
	}
	return &StakingContractValidatorStoppedIterator{contract: _StakingContract.contract, event: "ValidatorStopped", logs: logs, sub: sub}, nil
}

// WatchValidatorStopped is a free log subscription operation binding the contract event 0xf25558665a382a9abb684f20b20021df5923b51485bbf2829ff0089b5b271410.
//
// Solidity: event ValidatorStopped(uint256 stoppedCount)
func (_StakingContract *StakingContractFilterer) WatchValidatorStopped(opts *bind.WatchOpts, sink chan<- *StakingContractValidatorStopped) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "ValidatorStopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractValidatorStopped)
				if err := _StakingContract.contract.UnpackLog(event, "ValidatorStopped", log); err != nil {
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

// ParseValidatorStopped is a log parse operation binding the contract event 0xf25558665a382a9abb684f20b20021df5923b51485bbf2829ff0089b5b271410.
//
// Solidity: event ValidatorStopped(uint256 stoppedCount)
func (_StakingContract *StakingContractFilterer) ParseValidatorStopped(log types.Log) (*StakingContractValidatorStopped, error) {
	event := new(StakingContractValidatorStopped)
	if err := _StakingContract.contract.UnpackLog(event, "ValidatorStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractWhiteListToggleIterator is returned from FilterWhiteListToggle and is used to iterate over the raw logs and unpacked data for WhiteListToggle events raised by the StakingContract contract.
type StakingContractWhiteListToggleIterator struct {
	Event *StakingContractWhiteListToggle // Event containing the contract specifics and raw log

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
func (it *StakingContractWhiteListToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractWhiteListToggle)
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
		it.Event = new(StakingContractWhiteListToggle)
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
func (it *StakingContractWhiteListToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractWhiteListToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractWhiteListToggle represents a WhiteListToggle event raised by the StakingContract contract.
type StakingContractWhiteListToggle struct {
	Account common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhiteListToggle is a free log retrieval operation binding the contract event 0x96993f732e425c7615ed977e16e6e83d84bb45203ca23885018335b66c2e85be.
//
// Solidity: event WhiteListToggle(address account, bool enabled)
func (_StakingContract *StakingContractFilterer) FilterWhiteListToggle(opts *bind.FilterOpts) (*StakingContractWhiteListToggleIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "WhiteListToggle")
	if err != nil {
		return nil, err
	}
	return &StakingContractWhiteListToggleIterator{contract: _StakingContract.contract, event: "WhiteListToggle", logs: logs, sub: sub}, nil
}

// WatchWhiteListToggle is a free log subscription operation binding the contract event 0x96993f732e425c7615ed977e16e6e83d84bb45203ca23885018335b66c2e85be.
//
// Solidity: event WhiteListToggle(address account, bool enabled)
func (_StakingContract *StakingContractFilterer) WatchWhiteListToggle(opts *bind.WatchOpts, sink chan<- *StakingContractWhiteListToggle) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "WhiteListToggle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractWhiteListToggle)
				if err := _StakingContract.contract.UnpackLog(event, "WhiteListToggle", log); err != nil {
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

// ParseWhiteListToggle is a log parse operation binding the contract event 0x96993f732e425c7615ed977e16e6e83d84bb45203ca23885018335b66c2e85be.
//
// Solidity: event WhiteListToggle(address account, bool enabled)
func (_StakingContract *StakingContractFilterer) ParseWhiteListToggle(log types.Log) (*StakingContractWhiteListToggle, error) {
	event := new(StakingContractWhiteListToggle)
	if err := _StakingContract.contract.UnpackLog(event, "WhiteListToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractWithdrawCredentialSetIterator is returned from FilterWithdrawCredentialSet and is used to iterate over the raw logs and unpacked data for WithdrawCredentialSet events raised by the StakingContract contract.
type StakingContractWithdrawCredentialSetIterator struct {
	Event *StakingContractWithdrawCredentialSet // Event containing the contract specifics and raw log

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
func (it *StakingContractWithdrawCredentialSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractWithdrawCredentialSet)
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
		it.Event = new(StakingContractWithdrawCredentialSet)
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
func (it *StakingContractWithdrawCredentialSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractWithdrawCredentialSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractWithdrawCredentialSet represents a WithdrawCredentialSet event raised by the StakingContract contract.
type StakingContractWithdrawCredentialSet struct {
	WithdrawCredential [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCredentialSet is a free log retrieval operation binding the contract event 0x690facbfacb53c9319489117a6ac422718b5cb059a6ffade4871ff10f6f9aee9.
//
// Solidity: event WithdrawCredentialSet(bytes32 withdrawCredential)
func (_StakingContract *StakingContractFilterer) FilterWithdrawCredentialSet(opts *bind.FilterOpts) (*StakingContractWithdrawCredentialSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "WithdrawCredentialSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractWithdrawCredentialSetIterator{contract: _StakingContract.contract, event: "WithdrawCredentialSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawCredentialSet is a free log subscription operation binding the contract event 0x690facbfacb53c9319489117a6ac422718b5cb059a6ffade4871ff10f6f9aee9.
//
// Solidity: event WithdrawCredentialSet(bytes32 withdrawCredential)
func (_StakingContract *StakingContractFilterer) WatchWithdrawCredentialSet(opts *bind.WatchOpts, sink chan<- *StakingContractWithdrawCredentialSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "WithdrawCredentialSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractWithdrawCredentialSet)
				if err := _StakingContract.contract.UnpackLog(event, "WithdrawCredentialSet", log); err != nil {
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

// ParseWithdrawCredentialSet is a log parse operation binding the contract event 0x690facbfacb53c9319489117a6ac422718b5cb059a6ffade4871ff10f6f9aee9.
//
// Solidity: event WithdrawCredentialSet(bytes32 withdrawCredential)
func (_StakingContract *StakingContractFilterer) ParseWithdrawCredentialSet(log types.Log) (*StakingContractWithdrawCredentialSet, error) {
	event := new(StakingContractWithdrawCredentialSet)
	if err := _StakingContract.contract.UnpackLog(event, "WithdrawCredentialSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractXETHContractSetIterator is returned from FilterXETHContractSet and is used to iterate over the raw logs and unpacked data for XETHContractSet events raised by the StakingContract contract.
type StakingContractXETHContractSetIterator struct {
	Event *StakingContractXETHContractSet // Event containing the contract specifics and raw log

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
func (it *StakingContractXETHContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractXETHContractSet)
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
		it.Event = new(StakingContractXETHContractSet)
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
func (it *StakingContractXETHContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractXETHContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractXETHContractSet represents a XETHContractSet event raised by the StakingContract contract.
type StakingContractXETHContractSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXETHContractSet is a free log retrieval operation binding the contract event 0x65f24f264bd70348b9888b2fbf27cd04f9e2fb0fcc50dde8fb36483576ef32a3.
//
// Solidity: event XETHContractSet(address addr)
func (_StakingContract *StakingContractFilterer) FilterXETHContractSet(opts *bind.FilterOpts) (*StakingContractXETHContractSetIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "XETHContractSet")
	if err != nil {
		return nil, err
	}
	return &StakingContractXETHContractSetIterator{contract: _StakingContract.contract, event: "XETHContractSet", logs: logs, sub: sub}, nil
}

// WatchXETHContractSet is a free log subscription operation binding the contract event 0x65f24f264bd70348b9888b2fbf27cd04f9e2fb0fcc50dde8fb36483576ef32a3.
//
// Solidity: event XETHContractSet(address addr)
func (_StakingContract *StakingContractFilterer) WatchXETHContractSet(opts *bind.WatchOpts, sink chan<- *StakingContractXETHContractSet) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "XETHContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractXETHContractSet)
				if err := _StakingContract.contract.UnpackLog(event, "XETHContractSet", log); err != nil {
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

// ParseXETHContractSet is a log parse operation binding the contract event 0x65f24f264bd70348b9888b2fbf27cd04f9e2fb0fcc50dde8fb36483576ef32a3.
//
// Solidity: event XETHContractSet(address addr)
func (_StakingContract *StakingContractFilterer) ParseXETHContractSet(log types.Log) (*StakingContractXETHContractSet, error) {
	event := new(StakingContractXETHContractSet)
	if err := _StakingContract.contract.UnpackLog(event, "XETHContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
