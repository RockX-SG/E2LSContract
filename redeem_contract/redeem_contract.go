// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redeem_contract

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

// DelayRedeemRouterDelayedRedeem is an auto generated low-level Go binding around an user-defined struct.
type DelayRedeemRouterDelayedRedeem struct {
	Amount    *big.Int
	CreatedAt uint32
	Token     common.Address
}

// DelayRedeemRouterUserDelayedRedeems is an auto generated low-level Go binding around an user-defined struct.
type DelayRedeemRouterUserDelayedRedeems struct {
	DelayedRedeemsCompleted *big.Int
	DelayedRedeems          []DelayRedeemRouterDelayedRedeem
}

// RedeemContractMetaData contains all meta data concerning the RedeemContract contract.
var RedeemContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"BlacklistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"BlacklistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"BtclistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"BtclistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemFee\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delayedRedeemsCompleted\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsPrincipalClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delayedRedeemsCompleted\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsPrincipalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ManagementFeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"MaxQuotaSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"RateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"RedeemDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"RedeemFeeRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"RedeemPrincipalDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"WhitelistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"WhitelistEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"WhitelistRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_REDEEM_FEE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXCHANGE_RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DAILY_REDEEM_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REDEEM_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_BTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDEEM_FEE_RATE_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_A_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"addToBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"addToBtclist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"canClaimDelayedRedeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"canClaimDelayedRedeemPrincipal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDelayedRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxNumberOfDelayedRedeemsToClaim\",\"type\":\"uint256\"}],\"name\":\"claimDelayedRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxNumberOfDelayedRedeemsToClaim\",\"type\":\"uint256\"}],\"name\":\"claimPrincipals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPrincipals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createDelayedRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimableUserDelayedRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getCurrentCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserDelayedRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_redeemDelay\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_whitelistEnabled\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isBtclisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastRebaseTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managementFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxQuotas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"pauseTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"quotaBases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"quotaRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemPrincipalDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"removeFromBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeFromBtclist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_quotas\",\"type\":\"uint256[]\"}],\"name\":\"setMaxQuotaForTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_quotas\",\"type\":\"uint256[]\"}],\"name\":\"setQuotaRates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDelay\",\"type\":\"uint256\"}],\"name\":\"setRedeemDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setRedeemFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDelay\",\"type\":\"uint256\"}],\"name\":\"setRedeemPrincipalDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDebts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCleared\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"unpauseTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"userDelayedRedeemByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayedRedeemsCompleted\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"delayedRedeems\",\"type\":\"tuple[]\"}],\"internalType\":\"structDelayRedeemRouter.UserDelayedRedeems\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userRedeemsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"withdrawManagementFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RedeemContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RedeemContractMetaData.ABI instead.
var RedeemContractABI = RedeemContractMetaData.ABI

// RedeemContract is an auto generated Go binding around an Ethereum contract.
type RedeemContract struct {
	RedeemContractCaller     // Read-only binding to the contract
	RedeemContractTransactor // Write-only binding to the contract
	RedeemContractFilterer   // Log filterer for contract events
}

// RedeemContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedeemContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedeemContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedeemContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedeemContractSession struct {
	Contract     *RedeemContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedeemContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedeemContractCallerSession struct {
	Contract *RedeemContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RedeemContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedeemContractTransactorSession struct {
	Contract     *RedeemContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RedeemContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedeemContractRaw struct {
	Contract *RedeemContract // Generic contract binding to access the raw methods on
}

// RedeemContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedeemContractCallerRaw struct {
	Contract *RedeemContractCaller // Generic read-only contract binding to access the raw methods on
}

// RedeemContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedeemContractTransactorRaw struct {
	Contract *RedeemContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedeemContract creates a new instance of RedeemContract, bound to a specific deployed contract.
func NewRedeemContract(address common.Address, backend bind.ContractBackend) (*RedeemContract, error) {
	contract, err := bindRedeemContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RedeemContract{RedeemContractCaller: RedeemContractCaller{contract: contract}, RedeemContractTransactor: RedeemContractTransactor{contract: contract}, RedeemContractFilterer: RedeemContractFilterer{contract: contract}}, nil
}

// NewRedeemContractCaller creates a new read-only instance of RedeemContract, bound to a specific deployed contract.
func NewRedeemContractCaller(address common.Address, caller bind.ContractCaller) (*RedeemContractCaller, error) {
	contract, err := bindRedeemContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemContractCaller{contract: contract}, nil
}

// NewRedeemContractTransactor creates a new write-only instance of RedeemContract, bound to a specific deployed contract.
func NewRedeemContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RedeemContractTransactor, error) {
	contract, err := bindRedeemContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemContractTransactor{contract: contract}, nil
}

// NewRedeemContractFilterer creates a new log filterer instance of RedeemContract, bound to a specific deployed contract.
func NewRedeemContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RedeemContractFilterer, error) {
	contract, err := bindRedeemContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedeemContractFilterer{contract: contract}, nil
}

// bindRedeemContract binds a generic wrapper to an already deployed contract.
func bindRedeemContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedeemContract *RedeemContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedeemContract.Contract.RedeemContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedeemContract *RedeemContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemContract.Contract.RedeemContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedeemContract *RedeemContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedeemContract.Contract.RedeemContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedeemContract *RedeemContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedeemContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedeemContract *RedeemContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedeemContract *RedeemContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedeemContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemContract *RedeemContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemContract *RedeemContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RedeemContract.Contract.DEFAULTADMINROLE(&_RedeemContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemContract *RedeemContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RedeemContract.Contract.DEFAULTADMINROLE(&_RedeemContract.CallOpts)
}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) DEFAULTREDEEMFEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "DEFAULT_REDEEM_FEE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_RedeemContract *RedeemContractSession) DEFAULTREDEEMFEERATE() (*big.Int, error) {
	return _RedeemContract.Contract.DEFAULTREDEEMFEERATE(&_RedeemContract.CallOpts)
}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) DEFAULTREDEEMFEERATE() (*big.Int, error) {
	return _RedeemContract.Contract.DEFAULTREDEEMFEERATE(&_RedeemContract.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) EXCHANGERATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "EXCHANGE_RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_RedeemContract *RedeemContractSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _RedeemContract.Contract.EXCHANGERATEBASE(&_RedeemContract.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _RedeemContract.Contract.EXCHANGERATEBASE(&_RedeemContract.CallOpts)
}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) MAXDAILYREDEEMCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "MAX_DAILY_REDEEM_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_RedeemContract *RedeemContractSession) MAXDAILYREDEEMCAP() (*big.Int, error) {
	return _RedeemContract.Contract.MAXDAILYREDEEMCAP(&_RedeemContract.CallOpts)
}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) MAXDAILYREDEEMCAP() (*big.Int, error) {
	return _RedeemContract.Contract.MAXDAILYREDEEMCAP(&_RedeemContract.CallOpts)
}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) MAXREDEEMDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "MAX_REDEEM_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_RedeemContract *RedeemContractSession) MAXREDEEMDELAY() (*big.Int, error) {
	return _RedeemContract.Contract.MAXREDEEMDELAY(&_RedeemContract.CallOpts)
}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) MAXREDEEMDELAY() (*big.Int, error) {
	return _RedeemContract.Contract.MAXREDEEMDELAY(&_RedeemContract.CallOpts)
}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_RedeemContract *RedeemContractCaller) NATIVEBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "NATIVE_BTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_RedeemContract *RedeemContractSession) NATIVEBTC() (common.Address, error) {
	return _RedeemContract.Contract.NATIVEBTC(&_RedeemContract.CallOpts)
}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_RedeemContract *RedeemContractCallerSession) NATIVEBTC() (common.Address, error) {
	return _RedeemContract.Contract.NATIVEBTC(&_RedeemContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RedeemContract *RedeemContractCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RedeemContract *RedeemContractSession) PAUSERROLE() ([32]byte, error) {
	return _RedeemContract.Contract.PAUSERROLE(&_RedeemContract.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RedeemContract *RedeemContractCallerSession) PAUSERROLE() ([32]byte, error) {
	return _RedeemContract.Contract.PAUSERROLE(&_RedeemContract.CallOpts)
}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) REDEEMFEERATERANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "REDEEM_FEE_RATE_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_RedeemContract *RedeemContractSession) REDEEMFEERATERANGE() (*big.Int, error) {
	return _RedeemContract.Contract.REDEEMFEERATERANGE(&_RedeemContract.CallOpts)
}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) REDEEMFEERATERANGE() (*big.Int, error) {
	return _RedeemContract.Contract.REDEEMFEERATERANGE(&_RedeemContract.CallOpts)
}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) SECONDSINADAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "SECONDS_IN_A_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_RedeemContract *RedeemContractSession) SECONDSINADAY() (*big.Int, error) {
	return _RedeemContract.Contract.SECONDSINADAY(&_RedeemContract.CallOpts)
}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) SECONDSINADAY() (*big.Int, error) {
	return _RedeemContract.Contract.SECONDSINADAY(&_RedeemContract.CallOpts)
}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_RedeemContract *RedeemContractCaller) CanClaimDelayedRedeem(opts *bind.CallOpts, user common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "canClaimDelayedRedeem", user, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_RedeemContract *RedeemContractSession) CanClaimDelayedRedeem(user common.Address, index *big.Int) (bool, error) {
	return _RedeemContract.Contract.CanClaimDelayedRedeem(&_RedeemContract.CallOpts, user, index)
}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) CanClaimDelayedRedeem(user common.Address, index *big.Int) (bool, error) {
	return _RedeemContract.Contract.CanClaimDelayedRedeem(&_RedeemContract.CallOpts, user, index)
}

// CanClaimDelayedRedeemPrincipal is a free data retrieval call binding the contract method 0x7ac221fd.
//
// Solidity: function canClaimDelayedRedeemPrincipal(address user, uint256 index) view returns(bool)
func (_RedeemContract *RedeemContractCaller) CanClaimDelayedRedeemPrincipal(opts *bind.CallOpts, user common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "canClaimDelayedRedeemPrincipal", user, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimDelayedRedeemPrincipal is a free data retrieval call binding the contract method 0x7ac221fd.
//
// Solidity: function canClaimDelayedRedeemPrincipal(address user, uint256 index) view returns(bool)
func (_RedeemContract *RedeemContractSession) CanClaimDelayedRedeemPrincipal(user common.Address, index *big.Int) (bool, error) {
	return _RedeemContract.Contract.CanClaimDelayedRedeemPrincipal(&_RedeemContract.CallOpts, user, index)
}

// CanClaimDelayedRedeemPrincipal is a free data retrieval call binding the contract method 0x7ac221fd.
//
// Solidity: function canClaimDelayedRedeemPrincipal(address user, uint256 index) view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) CanClaimDelayedRedeemPrincipal(user common.Address, index *big.Int) (bool, error) {
	return _RedeemContract.Contract.CanClaimDelayedRedeemPrincipal(&_RedeemContract.CallOpts, user, index)
}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_RedeemContract *RedeemContractCaller) GetClaimableUserDelayedRedeems(opts *bind.CallOpts, user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "getClaimableUserDelayedRedeems", user)

	if err != nil {
		return *new([]DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DelayRedeemRouterDelayedRedeem)).(*[]DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_RedeemContract *RedeemContractSession) GetClaimableUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _RedeemContract.Contract.GetClaimableUserDelayedRedeems(&_RedeemContract.CallOpts, user)
}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_RedeemContract *RedeemContractCallerSession) GetClaimableUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _RedeemContract.Contract.GetClaimableUserDelayedRedeems(&_RedeemContract.CallOpts, user)
}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_RedeemContract *RedeemContractCaller) GetCurrentCap(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "getCurrentCap", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_RedeemContract *RedeemContractSession) GetCurrentCap(token common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.GetCurrentCap(&_RedeemContract.CallOpts, token)
}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) GetCurrentCap(token common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.GetCurrentCap(&_RedeemContract.CallOpts, token)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemContract *RedeemContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemContract *RedeemContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RedeemContract.Contract.GetRoleAdmin(&_RedeemContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemContract *RedeemContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RedeemContract.Contract.GetRoleAdmin(&_RedeemContract.CallOpts, role)
}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_RedeemContract *RedeemContractCaller) GetUserDelayedRedeems(opts *bind.CallOpts, user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "getUserDelayedRedeems", user)

	if err != nil {
		return *new([]DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DelayRedeemRouterDelayedRedeem)).(*[]DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_RedeemContract *RedeemContractSession) GetUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _RedeemContract.Contract.GetUserDelayedRedeems(&_RedeemContract.CallOpts, user)
}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_RedeemContract *RedeemContractCallerSession) GetUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _RedeemContract.Contract.GetUserDelayedRedeems(&_RedeemContract.CallOpts, user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemContract *RedeemContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemContract *RedeemContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RedeemContract.Contract.HasRole(&_RedeemContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RedeemContract.Contract.HasRole(&_RedeemContract.CallOpts, role, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_RedeemContract *RedeemContractCaller) IsBlacklisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "isBlacklisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_RedeemContract *RedeemContractSession) IsBlacklisted(account common.Address) (bool, error) {
	return _RedeemContract.Contract.IsBlacklisted(&_RedeemContract.CallOpts, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) IsBlacklisted(account common.Address) (bool, error) {
	return _RedeemContract.Contract.IsBlacklisted(&_RedeemContract.CallOpts, account)
}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_RedeemContract *RedeemContractCaller) IsBtclisted(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "isBtclisted", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_RedeemContract *RedeemContractSession) IsBtclisted(token common.Address) (bool, error) {
	return _RedeemContract.Contract.IsBtclisted(&_RedeemContract.CallOpts, token)
}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) IsBtclisted(token common.Address) (bool, error) {
	return _RedeemContract.Contract.IsBtclisted(&_RedeemContract.CallOpts, token)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_RedeemContract *RedeemContractCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_RedeemContract *RedeemContractSession) IsWhitelisted(account common.Address) (bool, error) {
	return _RedeemContract.Contract.IsWhitelisted(&_RedeemContract.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _RedeemContract.Contract.IsWhitelisted(&_RedeemContract.CallOpts, account)
}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_RedeemContract *RedeemContractCaller) LastRebaseTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "lastRebaseTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_RedeemContract *RedeemContractSession) LastRebaseTimestamps(arg0 common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.LastRebaseTimestamps(&_RedeemContract.CallOpts, arg0)
}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) LastRebaseTimestamps(arg0 common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.LastRebaseTimestamps(&_RedeemContract.CallOpts, arg0)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) ManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "managementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_RedeemContract *RedeemContractSession) ManagementFee() (*big.Int, error) {
	return _RedeemContract.Contract.ManagementFee(&_RedeemContract.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) ManagementFee() (*big.Int, error) {
	return _RedeemContract.Contract.ManagementFee(&_RedeemContract.CallOpts)
}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_RedeemContract *RedeemContractCaller) MaxQuotas(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "maxQuotas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_RedeemContract *RedeemContractSession) MaxQuotas(arg0 common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.MaxQuotas(&_RedeemContract.CallOpts, arg0)
}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) MaxQuotas(arg0 common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.MaxQuotas(&_RedeemContract.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RedeemContract *RedeemContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RedeemContract *RedeemContractSession) Paused() (bool, error) {
	return _RedeemContract.Contract.Paused(&_RedeemContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) Paused() (bool, error) {
	return _RedeemContract.Contract.Paused(&_RedeemContract.CallOpts)
}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_RedeemContract *RedeemContractCaller) QuotaBases(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "quotaBases", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_RedeemContract *RedeemContractSession) QuotaBases(arg0 common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.QuotaBases(&_RedeemContract.CallOpts, arg0)
}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) QuotaBases(arg0 common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.QuotaBases(&_RedeemContract.CallOpts, arg0)
}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_RedeemContract *RedeemContractCaller) QuotaRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "quotaRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_RedeemContract *RedeemContractSession) QuotaRates(arg0 common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.QuotaRates(&_RedeemContract.CallOpts, arg0)
}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) QuotaRates(arg0 common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.QuotaRates(&_RedeemContract.CallOpts, arg0)
}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) RedeemDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "redeemDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_RedeemContract *RedeemContractSession) RedeemDelay() (*big.Int, error) {
	return _RedeemContract.Contract.RedeemDelay(&_RedeemContract.CallOpts)
}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) RedeemDelay() (*big.Int, error) {
	return _RedeemContract.Contract.RedeemDelay(&_RedeemContract.CallOpts)
}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) RedeemFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "redeemFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_RedeemContract *RedeemContractSession) RedeemFeeRate() (*big.Int, error) {
	return _RedeemContract.Contract.RedeemFeeRate(&_RedeemContract.CallOpts)
}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) RedeemFeeRate() (*big.Int, error) {
	return _RedeemContract.Contract.RedeemFeeRate(&_RedeemContract.CallOpts)
}

// RedeemPrincipalDelay is a free data retrieval call binding the contract method 0x83c0894b.
//
// Solidity: function redeemPrincipalDelay() view returns(uint256)
func (_RedeemContract *RedeemContractCaller) RedeemPrincipalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "redeemPrincipalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemPrincipalDelay is a free data retrieval call binding the contract method 0x83c0894b.
//
// Solidity: function redeemPrincipalDelay() view returns(uint256)
func (_RedeemContract *RedeemContractSession) RedeemPrincipalDelay() (*big.Int, error) {
	return _RedeemContract.Contract.RedeemPrincipalDelay(&_RedeemContract.CallOpts)
}

// RedeemPrincipalDelay is a free data retrieval call binding the contract method 0x83c0894b.
//
// Solidity: function redeemPrincipalDelay() view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) RedeemPrincipalDelay() (*big.Int, error) {
	return _RedeemContract.Contract.RedeemPrincipalDelay(&_RedeemContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemContract *RedeemContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemContract *RedeemContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RedeemContract.Contract.SupportsInterface(&_RedeemContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RedeemContract.Contract.SupportsInterface(&_RedeemContract.CallOpts, interfaceId)
}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_RedeemContract *RedeemContractCaller) TokenDebts(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "tokenDebts", arg0)

	outstruct := new(struct {
		TotalDebts   *big.Int
		TotalCleared *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalDebts = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCleared = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_RedeemContract *RedeemContractSession) TokenDebts(arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	return _RedeemContract.Contract.TokenDebts(&_RedeemContract.CallOpts, arg0)
}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_RedeemContract *RedeemContractCallerSession) TokenDebts(arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	return _RedeemContract.Contract.TokenDebts(&_RedeemContract.CallOpts, arg0)
}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_RedeemContract *RedeemContractCaller) UniBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "uniBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_RedeemContract *RedeemContractSession) UniBTC() (common.Address, error) {
	return _RedeemContract.Contract.UniBTC(&_RedeemContract.CallOpts)
}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_RedeemContract *RedeemContractCallerSession) UniBTC() (common.Address, error) {
	return _RedeemContract.Contract.UniBTC(&_RedeemContract.CallOpts)
}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_RedeemContract *RedeemContractCaller) UserDelayedRedeemByIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "userDelayedRedeemByIndex", user, index)

	if err != nil {
		return *new(DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new(DelayRedeemRouterDelayedRedeem)).(*DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_RedeemContract *RedeemContractSession) UserDelayedRedeemByIndex(user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	return _RedeemContract.Contract.UserDelayedRedeemByIndex(&_RedeemContract.CallOpts, user, index)
}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_RedeemContract *RedeemContractCallerSession) UserDelayedRedeemByIndex(user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	return _RedeemContract.Contract.UserDelayedRedeemByIndex(&_RedeemContract.CallOpts, user, index)
}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_RedeemContract *RedeemContractCaller) UserRedeems(opts *bind.CallOpts, user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "userRedeems", user)

	if err != nil {
		return *new(DelayRedeemRouterUserDelayedRedeems), err
	}

	out0 := *abi.ConvertType(out[0], new(DelayRedeemRouterUserDelayedRedeems)).(*DelayRedeemRouterUserDelayedRedeems)

	return out0, err

}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_RedeemContract *RedeemContractSession) UserRedeems(user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	return _RedeemContract.Contract.UserRedeems(&_RedeemContract.CallOpts, user)
}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_RedeemContract *RedeemContractCallerSession) UserRedeems(user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	return _RedeemContract.Contract.UserRedeems(&_RedeemContract.CallOpts, user)
}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_RedeemContract *RedeemContractCaller) UserRedeemsLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "userRedeemsLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_RedeemContract *RedeemContractSession) UserRedeemsLength(user common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.UserRedeemsLength(&_RedeemContract.CallOpts, user)
}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_RedeemContract *RedeemContractCallerSession) UserRedeemsLength(user common.Address) (*big.Int, error) {
	return _RedeemContract.Contract.UserRedeemsLength(&_RedeemContract.CallOpts, user)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_RedeemContract *RedeemContractCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_RedeemContract *RedeemContractSession) Vault() (common.Address, error) {
	return _RedeemContract.Contract.Vault(&_RedeemContract.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_RedeemContract *RedeemContractCallerSession) Vault() (common.Address, error) {
	return _RedeemContract.Contract.Vault(&_RedeemContract.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_RedeemContract *RedeemContractCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RedeemContract.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_RedeemContract *RedeemContractSession) WhitelistEnabled() (bool, error) {
	return _RedeemContract.Contract.WhitelistEnabled(&_RedeemContract.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_RedeemContract *RedeemContractCallerSession) WhitelistEnabled() (bool, error) {
	return _RedeemContract.Contract.WhitelistEnabled(&_RedeemContract.CallOpts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractTransactor) AddToBlacklist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "addToBlacklist", _accounts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractSession) AddToBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.AddToBlacklist(&_RedeemContract.TransactOpts, _accounts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractTransactorSession) AddToBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.AddToBlacklist(&_RedeemContract.TransactOpts, _accounts)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_RedeemContract *RedeemContractTransactor) AddToBtclist(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "addToBtclist", _tokens)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_RedeemContract *RedeemContractSession) AddToBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.AddToBtclist(&_RedeemContract.TransactOpts, _tokens)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_RedeemContract *RedeemContractTransactorSession) AddToBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.AddToBtclist(&_RedeemContract.TransactOpts, _tokens)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractTransactor) AddToWhitelist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "addToWhitelist", _accounts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractSession) AddToWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.AddToWhitelist(&_RedeemContract.TransactOpts, _accounts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractTransactorSession) AddToWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.AddToWhitelist(&_RedeemContract.TransactOpts, _accounts)
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_RedeemContract *RedeemContractTransactor) ClaimDelayedRedeems(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "claimDelayedRedeems")
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_RedeemContract *RedeemContractSession) ClaimDelayedRedeems() (*types.Transaction, error) {
	return _RedeemContract.Contract.ClaimDelayedRedeems(&_RedeemContract.TransactOpts)
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_RedeemContract *RedeemContractTransactorSession) ClaimDelayedRedeems() (*types.Transaction, error) {
	return _RedeemContract.Contract.ClaimDelayedRedeems(&_RedeemContract.TransactOpts)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_RedeemContract *RedeemContractTransactor) ClaimDelayedRedeems0(opts *bind.TransactOpts, maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "claimDelayedRedeems0", maxNumberOfDelayedRedeemsToClaim)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_RedeemContract *RedeemContractSession) ClaimDelayedRedeems0(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.ClaimDelayedRedeems0(&_RedeemContract.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_RedeemContract *RedeemContractTransactorSession) ClaimDelayedRedeems0(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.ClaimDelayedRedeems0(&_RedeemContract.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0x7849ace2.
//
// Solidity: function claimPrincipals(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_RedeemContract *RedeemContractTransactor) ClaimPrincipals(opts *bind.TransactOpts, maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "claimPrincipals", maxNumberOfDelayedRedeemsToClaim)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0x7849ace2.
//
// Solidity: function claimPrincipals(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_RedeemContract *RedeemContractSession) ClaimPrincipals(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.ClaimPrincipals(&_RedeemContract.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0x7849ace2.
//
// Solidity: function claimPrincipals(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_RedeemContract *RedeemContractTransactorSession) ClaimPrincipals(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.ClaimPrincipals(&_RedeemContract.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimPrincipals0 is a paid mutator transaction binding the contract method 0xa60c7b3a.
//
// Solidity: function claimPrincipals() returns()
func (_RedeemContract *RedeemContractTransactor) ClaimPrincipals0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "claimPrincipals0")
}

// ClaimPrincipals0 is a paid mutator transaction binding the contract method 0xa60c7b3a.
//
// Solidity: function claimPrincipals() returns()
func (_RedeemContract *RedeemContractSession) ClaimPrincipals0() (*types.Transaction, error) {
	return _RedeemContract.Contract.ClaimPrincipals0(&_RedeemContract.TransactOpts)
}

// ClaimPrincipals0 is a paid mutator transaction binding the contract method 0xa60c7b3a.
//
// Solidity: function claimPrincipals() returns()
func (_RedeemContract *RedeemContractTransactorSession) ClaimPrincipals0() (*types.Transaction, error) {
	return _RedeemContract.Contract.ClaimPrincipals0(&_RedeemContract.TransactOpts)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_RedeemContract *RedeemContractTransactor) CreateDelayedRedeem(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "createDelayedRedeem", token, amount)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_RedeemContract *RedeemContractSession) CreateDelayedRedeem(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.CreateDelayedRedeem(&_RedeemContract.TransactOpts, token, amount)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_RedeemContract *RedeemContractTransactorSession) CreateDelayedRedeem(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.CreateDelayedRedeem(&_RedeemContract.TransactOpts, token, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.GrantRole(&_RedeemContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.GrantRole(&_RedeemContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_RedeemContract *RedeemContractTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "initialize", _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_RedeemContract *RedeemContractSession) Initialize(_defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _RedeemContract.Contract.Initialize(&_RedeemContract.TransactOpts, _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_RedeemContract *RedeemContractTransactorSession) Initialize(_defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _RedeemContract.Contract.Initialize(&_RedeemContract.TransactOpts, _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RedeemContract *RedeemContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RedeemContract *RedeemContractSession) Pause() (*types.Transaction, error) {
	return _RedeemContract.Contract.Pause(&_RedeemContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RedeemContract *RedeemContractTransactorSession) Pause() (*types.Transaction, error) {
	return _RedeemContract.Contract.Pause(&_RedeemContract.TransactOpts)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_RedeemContract *RedeemContractTransactor) PauseTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "pauseTokens", _tokens)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_RedeemContract *RedeemContractSession) PauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.PauseTokens(&_RedeemContract.TransactOpts, _tokens)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_RedeemContract *RedeemContractTransactorSession) PauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.PauseTokens(&_RedeemContract.TransactOpts, _tokens)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractTransactor) RemoveFromBlacklist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "removeFromBlacklist", _accounts)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractSession) RemoveFromBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RemoveFromBlacklist(&_RedeemContract.TransactOpts, _accounts)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractTransactorSession) RemoveFromBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RemoveFromBlacklist(&_RedeemContract.TransactOpts, _accounts)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_RedeemContract *RedeemContractTransactor) RemoveFromBtclist(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "removeFromBtclist", _tokens)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_RedeemContract *RedeemContractSession) RemoveFromBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RemoveFromBtclist(&_RedeemContract.TransactOpts, _tokens)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_RedeemContract *RedeemContractTransactorSession) RemoveFromBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RemoveFromBtclist(&_RedeemContract.TransactOpts, _tokens)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "removeFromWhitelist", _accounts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractSession) RemoveFromWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RemoveFromWhitelist(&_RedeemContract.TransactOpts, _accounts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_RedeemContract *RedeemContractTransactorSession) RemoveFromWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RemoveFromWhitelist(&_RedeemContract.TransactOpts, _accounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RenounceRole(&_RedeemContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RenounceRole(&_RedeemContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RevokeRole(&_RedeemContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemContract *RedeemContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.RevokeRole(&_RedeemContract.TransactOpts, role, account)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_RedeemContract *RedeemContractTransactor) SetMaxQuotaForTokens(opts *bind.TransactOpts, _tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "setMaxQuotaForTokens", _tokens, _quotas)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_RedeemContract *RedeemContractSession) SetMaxQuotaForTokens(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetMaxQuotaForTokens(&_RedeemContract.TransactOpts, _tokens, _quotas)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_RedeemContract *RedeemContractTransactorSession) SetMaxQuotaForTokens(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetMaxQuotaForTokens(&_RedeemContract.TransactOpts, _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_RedeemContract *RedeemContractTransactor) SetQuotaRates(opts *bind.TransactOpts, _tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "setQuotaRates", _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_RedeemContract *RedeemContractSession) SetQuotaRates(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetQuotaRates(&_RedeemContract.TransactOpts, _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_RedeemContract *RedeemContractTransactorSession) SetQuotaRates(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetQuotaRates(&_RedeemContract.TransactOpts, _tokens, _quotas)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_RedeemContract *RedeemContractTransactor) SetRedeemDelay(opts *bind.TransactOpts, _newDelay *big.Int) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "setRedeemDelay", _newDelay)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_RedeemContract *RedeemContractSession) SetRedeemDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetRedeemDelay(&_RedeemContract.TransactOpts, _newDelay)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_RedeemContract *RedeemContractTransactorSession) SetRedeemDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetRedeemDelay(&_RedeemContract.TransactOpts, _newDelay)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_RedeemContract *RedeemContractTransactor) SetRedeemFeeRate(opts *bind.TransactOpts, _newFeeRate *big.Int) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "setRedeemFeeRate", _newFeeRate)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_RedeemContract *RedeemContractSession) SetRedeemFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetRedeemFeeRate(&_RedeemContract.TransactOpts, _newFeeRate)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_RedeemContract *RedeemContractTransactorSession) SetRedeemFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetRedeemFeeRate(&_RedeemContract.TransactOpts, _newFeeRate)
}

// SetRedeemPrincipalDelay is a paid mutator transaction binding the contract method 0x98c0e0f0.
//
// Solidity: function setRedeemPrincipalDelay(uint256 _newDelay) returns()
func (_RedeemContract *RedeemContractTransactor) SetRedeemPrincipalDelay(opts *bind.TransactOpts, _newDelay *big.Int) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "setRedeemPrincipalDelay", _newDelay)
}

// SetRedeemPrincipalDelay is a paid mutator transaction binding the contract method 0x98c0e0f0.
//
// Solidity: function setRedeemPrincipalDelay(uint256 _newDelay) returns()
func (_RedeemContract *RedeemContractSession) SetRedeemPrincipalDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetRedeemPrincipalDelay(&_RedeemContract.TransactOpts, _newDelay)
}

// SetRedeemPrincipalDelay is a paid mutator transaction binding the contract method 0x98c0e0f0.
//
// Solidity: function setRedeemPrincipalDelay(uint256 _newDelay) returns()
func (_RedeemContract *RedeemContractTransactorSession) SetRedeemPrincipalDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetRedeemPrincipalDelay(&_RedeemContract.TransactOpts, _newDelay)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_RedeemContract *RedeemContractTransactor) SetWhitelistEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "setWhitelistEnabled", _enabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_RedeemContract *RedeemContractSession) SetWhitelistEnabled(_enabled bool) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetWhitelistEnabled(&_RedeemContract.TransactOpts, _enabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_RedeemContract *RedeemContractTransactorSession) SetWhitelistEnabled(_enabled bool) (*types.Transaction, error) {
	return _RedeemContract.Contract.SetWhitelistEnabled(&_RedeemContract.TransactOpts, _enabled)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RedeemContract *RedeemContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RedeemContract *RedeemContractSession) Unpause() (*types.Transaction, error) {
	return _RedeemContract.Contract.Unpause(&_RedeemContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RedeemContract *RedeemContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _RedeemContract.Contract.Unpause(&_RedeemContract.TransactOpts)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_RedeemContract *RedeemContractTransactor) UnpauseTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "unpauseTokens", _tokens)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_RedeemContract *RedeemContractSession) UnpauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.UnpauseTokens(&_RedeemContract.TransactOpts, _tokens)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_RedeemContract *RedeemContractTransactorSession) UnpauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.UnpauseTokens(&_RedeemContract.TransactOpts, _tokens)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_RedeemContract *RedeemContractTransactor) WithdrawManagementFee(opts *bind.TransactOpts, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _RedeemContract.contract.Transact(opts, "withdrawManagementFee", _amount, _recipient)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_RedeemContract *RedeemContractSession) WithdrawManagementFee(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.WithdrawManagementFee(&_RedeemContract.TransactOpts, _amount, _recipient)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_RedeemContract *RedeemContractTransactorSession) WithdrawManagementFee(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _RedeemContract.Contract.WithdrawManagementFee(&_RedeemContract.TransactOpts, _amount, _recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RedeemContract *RedeemContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RedeemContract *RedeemContractSession) Receive() (*types.Transaction, error) {
	return _RedeemContract.Contract.Receive(&_RedeemContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RedeemContract *RedeemContractTransactorSession) Receive() (*types.Transaction, error) {
	return _RedeemContract.Contract.Receive(&_RedeemContract.TransactOpts)
}

// RedeemContractBlacklistAddedIterator is returned from FilterBlacklistAdded and is used to iterate over the raw logs and unpacked data for BlacklistAdded events raised by the RedeemContract contract.
type RedeemContractBlacklistAddedIterator struct {
	Event *RedeemContractBlacklistAdded // Event containing the contract specifics and raw log

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
func (it *RedeemContractBlacklistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractBlacklistAdded)
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
		it.Event = new(RedeemContractBlacklistAdded)
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
func (it *RedeemContractBlacklistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractBlacklistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractBlacklistAdded represents a BlacklistAdded event raised by the RedeemContract contract.
type RedeemContractBlacklistAdded struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlacklistAdded is a free log retrieval operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) FilterBlacklistAdded(opts *bind.FilterOpts) (*RedeemContractBlacklistAddedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "BlacklistAdded")
	if err != nil {
		return nil, err
	}
	return &RedeemContractBlacklistAddedIterator{contract: _RedeemContract.contract, event: "BlacklistAdded", logs: logs, sub: sub}, nil
}

// WatchBlacklistAdded is a free log subscription operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) WatchBlacklistAdded(opts *bind.WatchOpts, sink chan<- *RedeemContractBlacklistAdded) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "BlacklistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractBlacklistAdded)
				if err := _RedeemContract.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
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

// ParseBlacklistAdded is a log parse operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) ParseBlacklistAdded(log types.Log) (*RedeemContractBlacklistAdded, error) {
	event := new(RedeemContractBlacklistAdded)
	if err := _RedeemContract.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractBlacklistRemovedIterator is returned from FilterBlacklistRemoved and is used to iterate over the raw logs and unpacked data for BlacklistRemoved events raised by the RedeemContract contract.
type RedeemContractBlacklistRemovedIterator struct {
	Event *RedeemContractBlacklistRemoved // Event containing the contract specifics and raw log

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
func (it *RedeemContractBlacklistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractBlacklistRemoved)
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
		it.Event = new(RedeemContractBlacklistRemoved)
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
func (it *RedeemContractBlacklistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractBlacklistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractBlacklistRemoved represents a BlacklistRemoved event raised by the RedeemContract contract.
type RedeemContractBlacklistRemoved struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlacklistRemoved is a free log retrieval operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) FilterBlacklistRemoved(opts *bind.FilterOpts) (*RedeemContractBlacklistRemovedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "BlacklistRemoved")
	if err != nil {
		return nil, err
	}
	return &RedeemContractBlacklistRemovedIterator{contract: _RedeemContract.contract, event: "BlacklistRemoved", logs: logs, sub: sub}, nil
}

// WatchBlacklistRemoved is a free log subscription operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) WatchBlacklistRemoved(opts *bind.WatchOpts, sink chan<- *RedeemContractBlacklistRemoved) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "BlacklistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractBlacklistRemoved)
				if err := _RedeemContract.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
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

// ParseBlacklistRemoved is a log parse operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) ParseBlacklistRemoved(log types.Log) (*RedeemContractBlacklistRemoved, error) {
	event := new(RedeemContractBlacklistRemoved)
	if err := _RedeemContract.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractBtclistAddedIterator is returned from FilterBtclistAdded and is used to iterate over the raw logs and unpacked data for BtclistAdded events raised by the RedeemContract contract.
type RedeemContractBtclistAddedIterator struct {
	Event *RedeemContractBtclistAdded // Event containing the contract specifics and raw log

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
func (it *RedeemContractBtclistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractBtclistAdded)
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
		it.Event = new(RedeemContractBtclistAdded)
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
func (it *RedeemContractBtclistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractBtclistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractBtclistAdded represents a BtclistAdded event raised by the RedeemContract contract.
type RedeemContractBtclistAdded struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBtclistAdded is a free log retrieval operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) FilterBtclistAdded(opts *bind.FilterOpts) (*RedeemContractBtclistAddedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "BtclistAdded")
	if err != nil {
		return nil, err
	}
	return &RedeemContractBtclistAddedIterator{contract: _RedeemContract.contract, event: "BtclistAdded", logs: logs, sub: sub}, nil
}

// WatchBtclistAdded is a free log subscription operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) WatchBtclistAdded(opts *bind.WatchOpts, sink chan<- *RedeemContractBtclistAdded) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "BtclistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractBtclistAdded)
				if err := _RedeemContract.contract.UnpackLog(event, "BtclistAdded", log); err != nil {
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

// ParseBtclistAdded is a log parse operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) ParseBtclistAdded(log types.Log) (*RedeemContractBtclistAdded, error) {
	event := new(RedeemContractBtclistAdded)
	if err := _RedeemContract.contract.UnpackLog(event, "BtclistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractBtclistRemovedIterator is returned from FilterBtclistRemoved and is used to iterate over the raw logs and unpacked data for BtclistRemoved events raised by the RedeemContract contract.
type RedeemContractBtclistRemovedIterator struct {
	Event *RedeemContractBtclistRemoved // Event containing the contract specifics and raw log

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
func (it *RedeemContractBtclistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractBtclistRemoved)
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
		it.Event = new(RedeemContractBtclistRemoved)
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
func (it *RedeemContractBtclistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractBtclistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractBtclistRemoved represents a BtclistRemoved event raised by the RedeemContract contract.
type RedeemContractBtclistRemoved struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBtclistRemoved is a free log retrieval operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) FilterBtclistRemoved(opts *bind.FilterOpts) (*RedeemContractBtclistRemovedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "BtclistRemoved")
	if err != nil {
		return nil, err
	}
	return &RedeemContractBtclistRemovedIterator{contract: _RedeemContract.contract, event: "BtclistRemoved", logs: logs, sub: sub}, nil
}

// WatchBtclistRemoved is a free log subscription operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) WatchBtclistRemoved(opts *bind.WatchOpts, sink chan<- *RedeemContractBtclistRemoved) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "BtclistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractBtclistRemoved)
				if err := _RedeemContract.contract.UnpackLog(event, "BtclistRemoved", log); err != nil {
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

// ParseBtclistRemoved is a log parse operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) ParseBtclistRemoved(log types.Log) (*RedeemContractBtclistRemoved, error) {
	event := new(RedeemContractBtclistRemoved)
	if err := _RedeemContract.contract.UnpackLog(event, "BtclistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractDelayedRedeemCreatedIterator is returned from FilterDelayedRedeemCreated and is used to iterate over the raw logs and unpacked data for DelayedRedeemCreated events raised by the RedeemContract contract.
type RedeemContractDelayedRedeemCreatedIterator struct {
	Event *RedeemContractDelayedRedeemCreated // Event containing the contract specifics and raw log

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
func (it *RedeemContractDelayedRedeemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractDelayedRedeemCreated)
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
		it.Event = new(RedeemContractDelayedRedeemCreated)
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
func (it *RedeemContractDelayedRedeemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractDelayedRedeemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractDelayedRedeemCreated represents a DelayedRedeemCreated event raised by the RedeemContract contract.
type RedeemContractDelayedRedeemCreated struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Index     *big.Int
	RedeemFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemCreated is a free log retrieval operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_RedeemContract *RedeemContractFilterer) FilterDelayedRedeemCreated(opts *bind.FilterOpts) (*RedeemContractDelayedRedeemCreatedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "DelayedRedeemCreated")
	if err != nil {
		return nil, err
	}
	return &RedeemContractDelayedRedeemCreatedIterator{contract: _RedeemContract.contract, event: "DelayedRedeemCreated", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemCreated is a free log subscription operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_RedeemContract *RedeemContractFilterer) WatchDelayedRedeemCreated(opts *bind.WatchOpts, sink chan<- *RedeemContractDelayedRedeemCreated) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "DelayedRedeemCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractDelayedRedeemCreated)
				if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemCreated", log); err != nil {
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

// ParseDelayedRedeemCreated is a log parse operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_RedeemContract *RedeemContractFilterer) ParseDelayedRedeemCreated(log types.Log) (*RedeemContractDelayedRedeemCreated, error) {
	event := new(RedeemContractDelayedRedeemCreated)
	if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractDelayedRedeemsClaimedIterator is returned from FilterDelayedRedeemsClaimed and is used to iterate over the raw logs and unpacked data for DelayedRedeemsClaimed events raised by the RedeemContract contract.
type RedeemContractDelayedRedeemsClaimedIterator struct {
	Event *RedeemContractDelayedRedeemsClaimed // Event containing the contract specifics and raw log

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
func (it *RedeemContractDelayedRedeemsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractDelayedRedeemsClaimed)
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
		it.Event = new(RedeemContractDelayedRedeemsClaimed)
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
func (it *RedeemContractDelayedRedeemsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractDelayedRedeemsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractDelayedRedeemsClaimed represents a DelayedRedeemsClaimed event raised by the RedeemContract contract.
type RedeemContractDelayedRedeemsClaimed struct {
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsClaimed is a free log retrieval operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_RedeemContract *RedeemContractFilterer) FilterDelayedRedeemsClaimed(opts *bind.FilterOpts) (*RedeemContractDelayedRedeemsClaimedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "DelayedRedeemsClaimed")
	if err != nil {
		return nil, err
	}
	return &RedeemContractDelayedRedeemsClaimedIterator{contract: _RedeemContract.contract, event: "DelayedRedeemsClaimed", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsClaimed is a free log subscription operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_RedeemContract *RedeemContractFilterer) WatchDelayedRedeemsClaimed(opts *bind.WatchOpts, sink chan<- *RedeemContractDelayedRedeemsClaimed) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "DelayedRedeemsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractDelayedRedeemsClaimed)
				if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemsClaimed", log); err != nil {
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

// ParseDelayedRedeemsClaimed is a log parse operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_RedeemContract *RedeemContractFilterer) ParseDelayedRedeemsClaimed(log types.Log) (*RedeemContractDelayedRedeemsClaimed, error) {
	event := new(RedeemContractDelayedRedeemsClaimed)
	if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractDelayedRedeemsCompletedIterator is returned from FilterDelayedRedeemsCompleted and is used to iterate over the raw logs and unpacked data for DelayedRedeemsCompleted events raised by the RedeemContract contract.
type RedeemContractDelayedRedeemsCompletedIterator struct {
	Event *RedeemContractDelayedRedeemsCompleted // Event containing the contract specifics and raw log

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
func (it *RedeemContractDelayedRedeemsCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractDelayedRedeemsCompleted)
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
		it.Event = new(RedeemContractDelayedRedeemsCompleted)
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
func (it *RedeemContractDelayedRedeemsCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractDelayedRedeemsCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractDelayedRedeemsCompleted represents a DelayedRedeemsCompleted event raised by the RedeemContract contract.
type RedeemContractDelayedRedeemsCompleted struct {
	Recipient               common.Address
	BurnedAmount            *big.Int
	DelayedRedeemsCompleted *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsCompleted is a free log retrieval operation binding the contract event 0xe4ba4789a7bd26025ed5932a86697793be11feeb86193a97dd9ef9849531ee60.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted)
func (_RedeemContract *RedeemContractFilterer) FilterDelayedRedeemsCompleted(opts *bind.FilterOpts) (*RedeemContractDelayedRedeemsCompletedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "DelayedRedeemsCompleted")
	if err != nil {
		return nil, err
	}
	return &RedeemContractDelayedRedeemsCompletedIterator{contract: _RedeemContract.contract, event: "DelayedRedeemsCompleted", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsCompleted is a free log subscription operation binding the contract event 0xe4ba4789a7bd26025ed5932a86697793be11feeb86193a97dd9ef9849531ee60.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted)
func (_RedeemContract *RedeemContractFilterer) WatchDelayedRedeemsCompleted(opts *bind.WatchOpts, sink chan<- *RedeemContractDelayedRedeemsCompleted) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "DelayedRedeemsCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractDelayedRedeemsCompleted)
				if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemsCompleted", log); err != nil {
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

// ParseDelayedRedeemsCompleted is a log parse operation binding the contract event 0xe4ba4789a7bd26025ed5932a86697793be11feeb86193a97dd9ef9849531ee60.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted)
func (_RedeemContract *RedeemContractFilterer) ParseDelayedRedeemsCompleted(log types.Log) (*RedeemContractDelayedRedeemsCompleted, error) {
	event := new(RedeemContractDelayedRedeemsCompleted)
	if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemsCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractDelayedRedeemsPrincipalClaimedIterator is returned from FilterDelayedRedeemsPrincipalClaimed and is used to iterate over the raw logs and unpacked data for DelayedRedeemsPrincipalClaimed events raised by the RedeemContract contract.
type RedeemContractDelayedRedeemsPrincipalClaimedIterator struct {
	Event *RedeemContractDelayedRedeemsPrincipalClaimed // Event containing the contract specifics and raw log

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
func (it *RedeemContractDelayedRedeemsPrincipalClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractDelayedRedeemsPrincipalClaimed)
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
		it.Event = new(RedeemContractDelayedRedeemsPrincipalClaimed)
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
func (it *RedeemContractDelayedRedeemsPrincipalClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractDelayedRedeemsPrincipalClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractDelayedRedeemsPrincipalClaimed represents a DelayedRedeemsPrincipalClaimed event raised by the RedeemContract contract.
type RedeemContractDelayedRedeemsPrincipalClaimed struct {
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsPrincipalClaimed is a free log retrieval operation binding the contract event 0xbf84d30fb64eb92b5961bba3a9c507953f2ea0fec68d7b78601c455825fb334a.
//
// Solidity: event DelayedRedeemsPrincipalClaimed(address recipient, address token, uint256 claimedAmount)
func (_RedeemContract *RedeemContractFilterer) FilterDelayedRedeemsPrincipalClaimed(opts *bind.FilterOpts) (*RedeemContractDelayedRedeemsPrincipalClaimedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "DelayedRedeemsPrincipalClaimed")
	if err != nil {
		return nil, err
	}
	return &RedeemContractDelayedRedeemsPrincipalClaimedIterator{contract: _RedeemContract.contract, event: "DelayedRedeemsPrincipalClaimed", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsPrincipalClaimed is a free log subscription operation binding the contract event 0xbf84d30fb64eb92b5961bba3a9c507953f2ea0fec68d7b78601c455825fb334a.
//
// Solidity: event DelayedRedeemsPrincipalClaimed(address recipient, address token, uint256 claimedAmount)
func (_RedeemContract *RedeemContractFilterer) WatchDelayedRedeemsPrincipalClaimed(opts *bind.WatchOpts, sink chan<- *RedeemContractDelayedRedeemsPrincipalClaimed) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "DelayedRedeemsPrincipalClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractDelayedRedeemsPrincipalClaimed)
				if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemsPrincipalClaimed", log); err != nil {
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

// ParseDelayedRedeemsPrincipalClaimed is a log parse operation binding the contract event 0xbf84d30fb64eb92b5961bba3a9c507953f2ea0fec68d7b78601c455825fb334a.
//
// Solidity: event DelayedRedeemsPrincipalClaimed(address recipient, address token, uint256 claimedAmount)
func (_RedeemContract *RedeemContractFilterer) ParseDelayedRedeemsPrincipalClaimed(log types.Log) (*RedeemContractDelayedRedeemsPrincipalClaimed, error) {
	event := new(RedeemContractDelayedRedeemsPrincipalClaimed)
	if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemsPrincipalClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractDelayedRedeemsPrincipalCompletedIterator is returned from FilterDelayedRedeemsPrincipalCompleted and is used to iterate over the raw logs and unpacked data for DelayedRedeemsPrincipalCompleted events raised by the RedeemContract contract.
type RedeemContractDelayedRedeemsPrincipalCompletedIterator struct {
	Event *RedeemContractDelayedRedeemsPrincipalCompleted // Event containing the contract specifics and raw log

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
func (it *RedeemContractDelayedRedeemsPrincipalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractDelayedRedeemsPrincipalCompleted)
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
		it.Event = new(RedeemContractDelayedRedeemsPrincipalCompleted)
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
func (it *RedeemContractDelayedRedeemsPrincipalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractDelayedRedeemsPrincipalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractDelayedRedeemsPrincipalCompleted represents a DelayedRedeemsPrincipalCompleted event raised by the RedeemContract contract.
type RedeemContractDelayedRedeemsPrincipalCompleted struct {
	Recipient               common.Address
	PrincipalAmount         *big.Int
	DelayedRedeemsCompleted *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsPrincipalCompleted is a free log retrieval operation binding the contract event 0xd8bb463f2d1d5c972542f10f4a33f937c1042d40970dc2f0a9f927e4f18b89c7.
//
// Solidity: event DelayedRedeemsPrincipalCompleted(address recipient, uint256 principalAmount, uint256 delayedRedeemsCompleted)
func (_RedeemContract *RedeemContractFilterer) FilterDelayedRedeemsPrincipalCompleted(opts *bind.FilterOpts) (*RedeemContractDelayedRedeemsPrincipalCompletedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "DelayedRedeemsPrincipalCompleted")
	if err != nil {
		return nil, err
	}
	return &RedeemContractDelayedRedeemsPrincipalCompletedIterator{contract: _RedeemContract.contract, event: "DelayedRedeemsPrincipalCompleted", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsPrincipalCompleted is a free log subscription operation binding the contract event 0xd8bb463f2d1d5c972542f10f4a33f937c1042d40970dc2f0a9f927e4f18b89c7.
//
// Solidity: event DelayedRedeemsPrincipalCompleted(address recipient, uint256 principalAmount, uint256 delayedRedeemsCompleted)
func (_RedeemContract *RedeemContractFilterer) WatchDelayedRedeemsPrincipalCompleted(opts *bind.WatchOpts, sink chan<- *RedeemContractDelayedRedeemsPrincipalCompleted) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "DelayedRedeemsPrincipalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractDelayedRedeemsPrincipalCompleted)
				if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemsPrincipalCompleted", log); err != nil {
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

// ParseDelayedRedeemsPrincipalCompleted is a log parse operation binding the contract event 0xd8bb463f2d1d5c972542f10f4a33f937c1042d40970dc2f0a9f927e4f18b89c7.
//
// Solidity: event DelayedRedeemsPrincipalCompleted(address recipient, uint256 principalAmount, uint256 delayedRedeemsCompleted)
func (_RedeemContract *RedeemContractFilterer) ParseDelayedRedeemsPrincipalCompleted(log types.Log) (*RedeemContractDelayedRedeemsPrincipalCompleted, error) {
	event := new(RedeemContractDelayedRedeemsPrincipalCompleted)
	if err := _RedeemContract.contract.UnpackLog(event, "DelayedRedeemsPrincipalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RedeemContract contract.
type RedeemContractInitializedIterator struct {
	Event *RedeemContractInitialized // Event containing the contract specifics and raw log

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
func (it *RedeemContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractInitialized)
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
		it.Event = new(RedeemContractInitialized)
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
func (it *RedeemContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractInitialized represents a Initialized event raised by the RedeemContract contract.
type RedeemContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedeemContract *RedeemContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*RedeemContractInitializedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RedeemContractInitializedIterator{contract: _RedeemContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedeemContract *RedeemContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RedeemContractInitialized) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractInitialized)
				if err := _RedeemContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RedeemContract *RedeemContractFilterer) ParseInitialized(log types.Log) (*RedeemContractInitialized, error) {
	event := new(RedeemContractInitialized)
	if err := _RedeemContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractManagementFeeWithdrawnIterator is returned from FilterManagementFeeWithdrawn and is used to iterate over the raw logs and unpacked data for ManagementFeeWithdrawn events raised by the RedeemContract contract.
type RedeemContractManagementFeeWithdrawnIterator struct {
	Event *RedeemContractManagementFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *RedeemContractManagementFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractManagementFeeWithdrawn)
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
		it.Event = new(RedeemContractManagementFeeWithdrawn)
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
func (it *RedeemContractManagementFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractManagementFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractManagementFeeWithdrawn represents a ManagementFeeWithdrawn event raised by the RedeemContract contract.
type RedeemContractManagementFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterManagementFeeWithdrawn is a free log retrieval operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_RedeemContract *RedeemContractFilterer) FilterManagementFeeWithdrawn(opts *bind.FilterOpts) (*RedeemContractManagementFeeWithdrawnIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "ManagementFeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &RedeemContractManagementFeeWithdrawnIterator{contract: _RedeemContract.contract, event: "ManagementFeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchManagementFeeWithdrawn is a free log subscription operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_RedeemContract *RedeemContractFilterer) WatchManagementFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *RedeemContractManagementFeeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "ManagementFeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractManagementFeeWithdrawn)
				if err := _RedeemContract.contract.UnpackLog(event, "ManagementFeeWithdrawn", log); err != nil {
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

// ParseManagementFeeWithdrawn is a log parse operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_RedeemContract *RedeemContractFilterer) ParseManagementFeeWithdrawn(log types.Log) (*RedeemContractManagementFeeWithdrawn, error) {
	event := new(RedeemContractManagementFeeWithdrawn)
	if err := _RedeemContract.contract.UnpackLog(event, "ManagementFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractMaxQuotaSetIterator is returned from FilterMaxQuotaSet and is used to iterate over the raw logs and unpacked data for MaxQuotaSet events raised by the RedeemContract contract.
type RedeemContractMaxQuotaSetIterator struct {
	Event *RedeemContractMaxQuotaSet // Event containing the contract specifics and raw log

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
func (it *RedeemContractMaxQuotaSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractMaxQuotaSet)
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
		it.Event = new(RedeemContractMaxQuotaSet)
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
func (it *RedeemContractMaxQuotaSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractMaxQuotaSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractMaxQuotaSet represents a MaxQuotaSet event raised by the RedeemContract contract.
type RedeemContractMaxQuotaSet struct {
	Token         common.Address
	PreviousQuota *big.Int
	NewQuota      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxQuotaSet is a free log retrieval operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_RedeemContract *RedeemContractFilterer) FilterMaxQuotaSet(opts *bind.FilterOpts) (*RedeemContractMaxQuotaSetIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "MaxQuotaSet")
	if err != nil {
		return nil, err
	}
	return &RedeemContractMaxQuotaSetIterator{contract: _RedeemContract.contract, event: "MaxQuotaSet", logs: logs, sub: sub}, nil
}

// WatchMaxQuotaSet is a free log subscription operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_RedeemContract *RedeemContractFilterer) WatchMaxQuotaSet(opts *bind.WatchOpts, sink chan<- *RedeemContractMaxQuotaSet) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "MaxQuotaSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractMaxQuotaSet)
				if err := _RedeemContract.contract.UnpackLog(event, "MaxQuotaSet", log); err != nil {
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

// ParseMaxQuotaSet is a log parse operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_RedeemContract *RedeemContractFilterer) ParseMaxQuotaSet(log types.Log) (*RedeemContractMaxQuotaSet, error) {
	event := new(RedeemContractMaxQuotaSet)
	if err := _RedeemContract.contract.UnpackLog(event, "MaxQuotaSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RedeemContract contract.
type RedeemContractPausedIterator struct {
	Event *RedeemContractPaused // Event containing the contract specifics and raw log

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
func (it *RedeemContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractPaused)
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
		it.Event = new(RedeemContractPaused)
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
func (it *RedeemContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractPaused represents a Paused event raised by the RedeemContract contract.
type RedeemContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RedeemContract *RedeemContractFilterer) FilterPaused(opts *bind.FilterOpts) (*RedeemContractPausedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RedeemContractPausedIterator{contract: _RedeemContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RedeemContract *RedeemContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RedeemContractPaused) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractPaused)
				if err := _RedeemContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_RedeemContract *RedeemContractFilterer) ParsePaused(log types.Log) (*RedeemContractPaused, error) {
	event := new(RedeemContractPaused)
	if err := _RedeemContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractRateSetIterator is returned from FilterRateSet and is used to iterate over the raw logs and unpacked data for RateSet events raised by the RedeemContract contract.
type RedeemContractRateSetIterator struct {
	Event *RedeemContractRateSet // Event containing the contract specifics and raw log

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
func (it *RedeemContractRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractRateSet)
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
		it.Event = new(RedeemContractRateSet)
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
func (it *RedeemContractRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractRateSet represents a RateSet event raised by the RedeemContract contract.
type RedeemContractRateSet struct {
	Token         common.Address
	PreviousQuota *big.Int
	NewQuota      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRateSet is a free log retrieval operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_RedeemContract *RedeemContractFilterer) FilterRateSet(opts *bind.FilterOpts) (*RedeemContractRateSetIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "RateSet")
	if err != nil {
		return nil, err
	}
	return &RedeemContractRateSetIterator{contract: _RedeemContract.contract, event: "RateSet", logs: logs, sub: sub}, nil
}

// WatchRateSet is a free log subscription operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_RedeemContract *RedeemContractFilterer) WatchRateSet(opts *bind.WatchOpts, sink chan<- *RedeemContractRateSet) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "RateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractRateSet)
				if err := _RedeemContract.contract.UnpackLog(event, "RateSet", log); err != nil {
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

// ParseRateSet is a log parse operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_RedeemContract *RedeemContractFilterer) ParseRateSet(log types.Log) (*RedeemContractRateSet, error) {
	event := new(RedeemContractRateSet)
	if err := _RedeemContract.contract.UnpackLog(event, "RateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractRedeemDelaySetIterator is returned from FilterRedeemDelaySet and is used to iterate over the raw logs and unpacked data for RedeemDelaySet events raised by the RedeemContract contract.
type RedeemContractRedeemDelaySetIterator struct {
	Event *RedeemContractRedeemDelaySet // Event containing the contract specifics and raw log

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
func (it *RedeemContractRedeemDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractRedeemDelaySet)
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
		it.Event = new(RedeemContractRedeemDelaySet)
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
func (it *RedeemContractRedeemDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractRedeemDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractRedeemDelaySet represents a RedeemDelaySet event raised by the RedeemContract contract.
type RedeemContractRedeemDelaySet struct {
	PreviousDelay *big.Int
	NewDelay      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRedeemDelaySet is a free log retrieval operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_RedeemContract *RedeemContractFilterer) FilterRedeemDelaySet(opts *bind.FilterOpts) (*RedeemContractRedeemDelaySetIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "RedeemDelaySet")
	if err != nil {
		return nil, err
	}
	return &RedeemContractRedeemDelaySetIterator{contract: _RedeemContract.contract, event: "RedeemDelaySet", logs: logs, sub: sub}, nil
}

// WatchRedeemDelaySet is a free log subscription operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_RedeemContract *RedeemContractFilterer) WatchRedeemDelaySet(opts *bind.WatchOpts, sink chan<- *RedeemContractRedeemDelaySet) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "RedeemDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractRedeemDelaySet)
				if err := _RedeemContract.contract.UnpackLog(event, "RedeemDelaySet", log); err != nil {
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

// ParseRedeemDelaySet is a log parse operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_RedeemContract *RedeemContractFilterer) ParseRedeemDelaySet(log types.Log) (*RedeemContractRedeemDelaySet, error) {
	event := new(RedeemContractRedeemDelaySet)
	if err := _RedeemContract.contract.UnpackLog(event, "RedeemDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractRedeemFeeRateSetIterator is returned from FilterRedeemFeeRateSet and is used to iterate over the raw logs and unpacked data for RedeemFeeRateSet events raised by the RedeemContract contract.
type RedeemContractRedeemFeeRateSetIterator struct {
	Event *RedeemContractRedeemFeeRateSet // Event containing the contract specifics and raw log

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
func (it *RedeemContractRedeemFeeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractRedeemFeeRateSet)
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
		it.Event = new(RedeemContractRedeemFeeRateSet)
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
func (it *RedeemContractRedeemFeeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractRedeemFeeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractRedeemFeeRateSet represents a RedeemFeeRateSet event raised by the RedeemContract contract.
type RedeemContractRedeemFeeRateSet struct {
	PreviousFeeRate *big.Int
	NewFeeRate      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRedeemFeeRateSet is a free log retrieval operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_RedeemContract *RedeemContractFilterer) FilterRedeemFeeRateSet(opts *bind.FilterOpts) (*RedeemContractRedeemFeeRateSetIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "RedeemFeeRateSet")
	if err != nil {
		return nil, err
	}
	return &RedeemContractRedeemFeeRateSetIterator{contract: _RedeemContract.contract, event: "RedeemFeeRateSet", logs: logs, sub: sub}, nil
}

// WatchRedeemFeeRateSet is a free log subscription operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_RedeemContract *RedeemContractFilterer) WatchRedeemFeeRateSet(opts *bind.WatchOpts, sink chan<- *RedeemContractRedeemFeeRateSet) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "RedeemFeeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractRedeemFeeRateSet)
				if err := _RedeemContract.contract.UnpackLog(event, "RedeemFeeRateSet", log); err != nil {
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

// ParseRedeemFeeRateSet is a log parse operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_RedeemContract *RedeemContractFilterer) ParseRedeemFeeRateSet(log types.Log) (*RedeemContractRedeemFeeRateSet, error) {
	event := new(RedeemContractRedeemFeeRateSet)
	if err := _RedeemContract.contract.UnpackLog(event, "RedeemFeeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractRedeemPrincipalDelaySetIterator is returned from FilterRedeemPrincipalDelaySet and is used to iterate over the raw logs and unpacked data for RedeemPrincipalDelaySet events raised by the RedeemContract contract.
type RedeemContractRedeemPrincipalDelaySetIterator struct {
	Event *RedeemContractRedeemPrincipalDelaySet // Event containing the contract specifics and raw log

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
func (it *RedeemContractRedeemPrincipalDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractRedeemPrincipalDelaySet)
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
		it.Event = new(RedeemContractRedeemPrincipalDelaySet)
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
func (it *RedeemContractRedeemPrincipalDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractRedeemPrincipalDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractRedeemPrincipalDelaySet represents a RedeemPrincipalDelaySet event raised by the RedeemContract contract.
type RedeemContractRedeemPrincipalDelaySet struct {
	PreviousDelay *big.Int
	NewDelay      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRedeemPrincipalDelaySet is a free log retrieval operation binding the contract event 0xcf8211b2d9a5296d32fee575ddbc4623ed83733cae205f8c1046b6eaf48dd7b2.
//
// Solidity: event RedeemPrincipalDelaySet(uint256 previousDelay, uint256 newDelay)
func (_RedeemContract *RedeemContractFilterer) FilterRedeemPrincipalDelaySet(opts *bind.FilterOpts) (*RedeemContractRedeemPrincipalDelaySetIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "RedeemPrincipalDelaySet")
	if err != nil {
		return nil, err
	}
	return &RedeemContractRedeemPrincipalDelaySetIterator{contract: _RedeemContract.contract, event: "RedeemPrincipalDelaySet", logs: logs, sub: sub}, nil
}

// WatchRedeemPrincipalDelaySet is a free log subscription operation binding the contract event 0xcf8211b2d9a5296d32fee575ddbc4623ed83733cae205f8c1046b6eaf48dd7b2.
//
// Solidity: event RedeemPrincipalDelaySet(uint256 previousDelay, uint256 newDelay)
func (_RedeemContract *RedeemContractFilterer) WatchRedeemPrincipalDelaySet(opts *bind.WatchOpts, sink chan<- *RedeemContractRedeemPrincipalDelaySet) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "RedeemPrincipalDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractRedeemPrincipalDelaySet)
				if err := _RedeemContract.contract.UnpackLog(event, "RedeemPrincipalDelaySet", log); err != nil {
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

// ParseRedeemPrincipalDelaySet is a log parse operation binding the contract event 0xcf8211b2d9a5296d32fee575ddbc4623ed83733cae205f8c1046b6eaf48dd7b2.
//
// Solidity: event RedeemPrincipalDelaySet(uint256 previousDelay, uint256 newDelay)
func (_RedeemContract *RedeemContractFilterer) ParseRedeemPrincipalDelaySet(log types.Log) (*RedeemContractRedeemPrincipalDelaySet, error) {
	event := new(RedeemContractRedeemPrincipalDelaySet)
	if err := _RedeemContract.contract.UnpackLog(event, "RedeemPrincipalDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RedeemContract contract.
type RedeemContractRoleAdminChangedIterator struct {
	Event *RedeemContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RedeemContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractRoleAdminChanged)
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
		it.Event = new(RedeemContractRoleAdminChanged)
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
func (it *RedeemContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractRoleAdminChanged represents a RoleAdminChanged event raised by the RedeemContract contract.
type RedeemContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RedeemContract *RedeemContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RedeemContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RedeemContractRoleAdminChangedIterator{contract: _RedeemContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RedeemContract *RedeemContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RedeemContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractRoleAdminChanged)
				if err := _RedeemContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_RedeemContract *RedeemContractFilterer) ParseRoleAdminChanged(log types.Log) (*RedeemContractRoleAdminChanged, error) {
	event := new(RedeemContractRoleAdminChanged)
	if err := _RedeemContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RedeemContract contract.
type RedeemContractRoleGrantedIterator struct {
	Event *RedeemContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *RedeemContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractRoleGranted)
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
		it.Event = new(RedeemContractRoleGranted)
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
func (it *RedeemContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractRoleGranted represents a RoleGranted event raised by the RedeemContract contract.
type RedeemContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemContract *RedeemContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RedeemContractRoleGrantedIterator, error) {

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

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RedeemContractRoleGrantedIterator{contract: _RedeemContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemContract *RedeemContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RedeemContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractRoleGranted)
				if err := _RedeemContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RedeemContract *RedeemContractFilterer) ParseRoleGranted(log types.Log) (*RedeemContractRoleGranted, error) {
	event := new(RedeemContractRoleGranted)
	if err := _RedeemContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RedeemContract contract.
type RedeemContractRoleRevokedIterator struct {
	Event *RedeemContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RedeemContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractRoleRevoked)
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
		it.Event = new(RedeemContractRoleRevoked)
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
func (it *RedeemContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractRoleRevoked represents a RoleRevoked event raised by the RedeemContract contract.
type RedeemContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemContract *RedeemContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RedeemContractRoleRevokedIterator, error) {

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

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RedeemContractRoleRevokedIterator{contract: _RedeemContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemContract *RedeemContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RedeemContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractRoleRevoked)
				if err := _RedeemContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RedeemContract *RedeemContractFilterer) ParseRoleRevoked(log types.Log) (*RedeemContractRoleRevoked, error) {
	event := new(RedeemContractRoleRevoked)
	if err := _RedeemContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractTokensPausedIterator is returned from FilterTokensPaused and is used to iterate over the raw logs and unpacked data for TokensPaused events raised by the RedeemContract contract.
type RedeemContractTokensPausedIterator struct {
	Event *RedeemContractTokensPaused // Event containing the contract specifics and raw log

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
func (it *RedeemContractTokensPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractTokensPaused)
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
		it.Event = new(RedeemContractTokensPaused)
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
func (it *RedeemContractTokensPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractTokensPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractTokensPaused represents a TokensPaused event raised by the RedeemContract contract.
type RedeemContractTokensPaused struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensPaused is a free log retrieval operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) FilterTokensPaused(opts *bind.FilterOpts) (*RedeemContractTokensPausedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "TokensPaused")
	if err != nil {
		return nil, err
	}
	return &RedeemContractTokensPausedIterator{contract: _RedeemContract.contract, event: "TokensPaused", logs: logs, sub: sub}, nil
}

// WatchTokensPaused is a free log subscription operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) WatchTokensPaused(opts *bind.WatchOpts, sink chan<- *RedeemContractTokensPaused) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "TokensPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractTokensPaused)
				if err := _RedeemContract.contract.UnpackLog(event, "TokensPaused", log); err != nil {
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

// ParseTokensPaused is a log parse operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) ParseTokensPaused(log types.Log) (*RedeemContractTokensPaused, error) {
	event := new(RedeemContractTokensPaused)
	if err := _RedeemContract.contract.UnpackLog(event, "TokensPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractTokensUnpausedIterator is returned from FilterTokensUnpaused and is used to iterate over the raw logs and unpacked data for TokensUnpaused events raised by the RedeemContract contract.
type RedeemContractTokensUnpausedIterator struct {
	Event *RedeemContractTokensUnpaused // Event containing the contract specifics and raw log

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
func (it *RedeemContractTokensUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractTokensUnpaused)
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
		it.Event = new(RedeemContractTokensUnpaused)
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
func (it *RedeemContractTokensUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractTokensUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractTokensUnpaused represents a TokensUnpaused event raised by the RedeemContract contract.
type RedeemContractTokensUnpaused struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensUnpaused is a free log retrieval operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) FilterTokensUnpaused(opts *bind.FilterOpts) (*RedeemContractTokensUnpausedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "TokensUnpaused")
	if err != nil {
		return nil, err
	}
	return &RedeemContractTokensUnpausedIterator{contract: _RedeemContract.contract, event: "TokensUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokensUnpaused is a free log subscription operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) WatchTokensUnpaused(opts *bind.WatchOpts, sink chan<- *RedeemContractTokensUnpaused) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "TokensUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractTokensUnpaused)
				if err := _RedeemContract.contract.UnpackLog(event, "TokensUnpaused", log); err != nil {
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

// ParseTokensUnpaused is a log parse operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_RedeemContract *RedeemContractFilterer) ParseTokensUnpaused(log types.Log) (*RedeemContractTokensUnpaused, error) {
	event := new(RedeemContractTokensUnpaused)
	if err := _RedeemContract.contract.UnpackLog(event, "TokensUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RedeemContract contract.
type RedeemContractUnpausedIterator struct {
	Event *RedeemContractUnpaused // Event containing the contract specifics and raw log

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
func (it *RedeemContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractUnpaused)
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
		it.Event = new(RedeemContractUnpaused)
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
func (it *RedeemContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractUnpaused represents a Unpaused event raised by the RedeemContract contract.
type RedeemContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RedeemContract *RedeemContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RedeemContractUnpausedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RedeemContractUnpausedIterator{contract: _RedeemContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RedeemContract *RedeemContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RedeemContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractUnpaused)
				if err := _RedeemContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_RedeemContract *RedeemContractFilterer) ParseUnpaused(log types.Log) (*RedeemContractUnpaused, error) {
	event := new(RedeemContractUnpaused)
	if err := _RedeemContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractWhitelistAddedIterator is returned from FilterWhitelistAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdded events raised by the RedeemContract contract.
type RedeemContractWhitelistAddedIterator struct {
	Event *RedeemContractWhitelistAdded // Event containing the contract specifics and raw log

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
func (it *RedeemContractWhitelistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractWhitelistAdded)
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
		it.Event = new(RedeemContractWhitelistAdded)
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
func (it *RedeemContractWhitelistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractWhitelistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractWhitelistAdded represents a WhitelistAdded event raised by the RedeemContract contract.
type RedeemContractWhitelistAdded struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdded is a free log retrieval operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) FilterWhitelistAdded(opts *bind.FilterOpts) (*RedeemContractWhitelistAddedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "WhitelistAdded")
	if err != nil {
		return nil, err
	}
	return &RedeemContractWhitelistAddedIterator{contract: _RedeemContract.contract, event: "WhitelistAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdded is a free log subscription operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) WatchWhitelistAdded(opts *bind.WatchOpts, sink chan<- *RedeemContractWhitelistAdded) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "WhitelistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractWhitelistAdded)
				if err := _RedeemContract.contract.UnpackLog(event, "WhitelistAdded", log); err != nil {
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

// ParseWhitelistAdded is a log parse operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) ParseWhitelistAdded(log types.Log) (*RedeemContractWhitelistAdded, error) {
	event := new(RedeemContractWhitelistAdded)
	if err := _RedeemContract.contract.UnpackLog(event, "WhitelistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractWhitelistEnabledSetIterator is returned from FilterWhitelistEnabledSet and is used to iterate over the raw logs and unpacked data for WhitelistEnabledSet events raised by the RedeemContract contract.
type RedeemContractWhitelistEnabledSetIterator struct {
	Event *RedeemContractWhitelistEnabledSet // Event containing the contract specifics and raw log

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
func (it *RedeemContractWhitelistEnabledSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractWhitelistEnabledSet)
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
		it.Event = new(RedeemContractWhitelistEnabledSet)
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
func (it *RedeemContractWhitelistEnabledSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractWhitelistEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractWhitelistEnabledSet represents a WhitelistEnabledSet event raised by the RedeemContract contract.
type RedeemContractWhitelistEnabledSet struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistEnabledSet is a free log retrieval operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_RedeemContract *RedeemContractFilterer) FilterWhitelistEnabledSet(opts *bind.FilterOpts) (*RedeemContractWhitelistEnabledSetIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "WhitelistEnabledSet")
	if err != nil {
		return nil, err
	}
	return &RedeemContractWhitelistEnabledSetIterator{contract: _RedeemContract.contract, event: "WhitelistEnabledSet", logs: logs, sub: sub}, nil
}

// WatchWhitelistEnabledSet is a free log subscription operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_RedeemContract *RedeemContractFilterer) WatchWhitelistEnabledSet(opts *bind.WatchOpts, sink chan<- *RedeemContractWhitelistEnabledSet) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "WhitelistEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractWhitelistEnabledSet)
				if err := _RedeemContract.contract.UnpackLog(event, "WhitelistEnabledSet", log); err != nil {
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

// ParseWhitelistEnabledSet is a log parse operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_RedeemContract *RedeemContractFilterer) ParseWhitelistEnabledSet(log types.Log) (*RedeemContractWhitelistEnabledSet, error) {
	event := new(RedeemContractWhitelistEnabledSet)
	if err := _RedeemContract.contract.UnpackLog(event, "WhitelistEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemContractWhitelistRemovedIterator is returned from FilterWhitelistRemoved and is used to iterate over the raw logs and unpacked data for WhitelistRemoved events raised by the RedeemContract contract.
type RedeemContractWhitelistRemovedIterator struct {
	Event *RedeemContractWhitelistRemoved // Event containing the contract specifics and raw log

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
func (it *RedeemContractWhitelistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemContractWhitelistRemoved)
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
		it.Event = new(RedeemContractWhitelistRemoved)
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
func (it *RedeemContractWhitelistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemContractWhitelistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemContractWhitelistRemoved represents a WhitelistRemoved event raised by the RedeemContract contract.
type RedeemContractWhitelistRemoved struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistRemoved is a free log retrieval operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) FilterWhitelistRemoved(opts *bind.FilterOpts) (*RedeemContractWhitelistRemovedIterator, error) {

	logs, sub, err := _RedeemContract.contract.FilterLogs(opts, "WhitelistRemoved")
	if err != nil {
		return nil, err
	}
	return &RedeemContractWhitelistRemovedIterator{contract: _RedeemContract.contract, event: "WhitelistRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistRemoved is a free log subscription operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) WatchWhitelistRemoved(opts *bind.WatchOpts, sink chan<- *RedeemContractWhitelistRemoved) (event.Subscription, error) {

	logs, sub, err := _RedeemContract.contract.WatchLogs(opts, "WhitelistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemContractWhitelistRemoved)
				if err := _RedeemContract.contract.UnpackLog(event, "WhitelistRemoved", log); err != nil {
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

// ParseWhitelistRemoved is a log parse operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_RedeemContract *RedeemContractFilterer) ParseWhitelistRemoved(log types.Log) (*RedeemContractWhitelistRemoved, error) {
	event := new(RedeemContractWhitelistRemoved)
	if err := _RedeemContract.contract.UnpackLog(event, "WhitelistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
