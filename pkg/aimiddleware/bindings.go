// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aimiddleware

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

// AiMiddlewareAiOpInfo is an auto generated low-level Go binding around an user-defined struct.
type AiMiddlewareAiOpInfo struct {
	MAiOp         AiMiddlewareMemoryAiOp
	AiOpHash      [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

// AiMiddlewareMemoryAiOp is an auto generated low-level Go binding around an user-defined struct.
type AiMiddlewareMemoryAiOp struct {
	Sender               common.Address
	Nonce                *big.Int
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	Paymaster            common.Address
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}

// AiOperation is an auto generated low-level Go binding around an user-defined struct.
type AiOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// IAiMiddlewareAiOpsPerAggregator is an auto generated low-level Go binding around an user-defined struct.
type IAiMiddlewareAiOpsPerAggregator struct {
	AiOps      []AiOperation
	Aggregator common.Address
	Signature  []byte
}

// IDepositManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IDepositManagerDepositInfo struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}

// AimiddlewareMetaData contains all meta data concerning the Aimiddleware contract.
var AimiddlewareMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"SIG_VALIDATION_FAILED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_validateSenderAndPaymaster\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStake\",\"inputs\":[{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositTo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"staked\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stake\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"withdrawTime\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAiOpHash\",\"inputs\":[{\"name\":\"aiOp\",\"type\":\"tuple\",\"internalType\":\"structAiOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositInfo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.DepositInfo\",\"components\":[{\"name\":\"deposit\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"staked\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stake\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"withdrawTime\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSenderAddress\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleAggregatedOps\",\"inputs\":[{\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\",\"internalType\":\"structIAiMiddleware.AiOpsPerAggregator[]\",\"components\":[{\"name\":\"aiOps\",\"type\":\"tuple[]\",\"internalType\":\"structAiOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"contractIAggregator\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleOps\",\"inputs\":[{\"name\":\"ops\",\"type\":\"tuple[]\",\"internalType\":\"structAiOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"incrementNonce\",\"inputs\":[{\"name\":\"key\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"innerHandleOp\",\"inputs\":[{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"opInfo\",\"type\":\"tuple\",\"internalType\":\"structAiMiddleware.AiOpInfo\",\"components\":[{\"name\":\"mAiOp\",\"type\":\"tuple\",\"internalType\":\"structAiMiddleware.MemoryAiOp\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymaster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"aiOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prefund\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preOpGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyValidator\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonceSequenceNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"simulateHandleOp\",\"inputs\":[{\"name\":\"op\",\"type\":\"tuple\",\"internalType\":\"structAiOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"targetCallData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"simulateValidation\",\"inputs\":[{\"name\":\"aiOp\",\"type\":\"tuple\",\"internalType\":\"structAiOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlockStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTo\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"withdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AccountDeployed\",\"inputs\":[{\"name\":\"aiOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"factory\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"paymaster\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AiOperationEvent\",\"inputs\":[{\"name\":\"aiOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymaster\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"actualGasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AiOperationRevertReason\",\"inputs\":[{\"name\":\"aiOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertReason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeforeAiopsExecution\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignatureAggregatorChanged\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeLocked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalStaked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeUnlocked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ExecutionResult\",\"inputs\":[{\"name\":\"preOpGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validAfter\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"validUntil\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"targetSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"targetResult\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"FailedOp\",\"inputs\":[{\"name\":\"opIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderAddressResult\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignatureValidationFailed\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ValidationResult\",\"inputs\":[{\"name\":\"returnInfo\",\"type\":\"tuple\",\"internalType\":\"structIAiMiddleware.ReturnInfo\",\"components\":[{\"name\":\"preOpGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prefund\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sigFailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"validAfter\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"validUntil\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"paymasterContext\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"senderInfo\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.StakeInfo\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"factoryInfo\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.StakeInfo\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"paymasterInfo\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.StakeInfo\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"ValidationResultWithAggregation\",\"inputs\":[{\"name\":\"returnInfo\",\"type\":\"tuple\",\"internalType\":\"structIAiMiddleware.ReturnInfo\",\"components\":[{\"name\":\"preOpGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prefund\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sigFailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"validAfter\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"validUntil\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"paymasterContext\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"senderInfo\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.StakeInfo\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"factoryInfo\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.StakeInfo\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"paymasterInfo\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.StakeInfo\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"aggregatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIAiMiddleware.AggregatorStakeInfo\",\"components\":[{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeInfo\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.StakeInfo\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]}]",
}

// AimiddlewareABI is the input ABI used to generate the binding from.
// Deprecated: Use AimiddlewareMetaData.ABI instead.
var AimiddlewareABI = AimiddlewareMetaData.ABI

// Aimiddleware is an auto generated Go binding around an Ethereum contract.
type Aimiddleware struct {
	AimiddlewareCaller     // Read-only binding to the contract
	AimiddlewareTransactor // Write-only binding to the contract
	AimiddlewareFilterer   // Log filterer for contract events
}

// AimiddlewareCaller is an auto generated read-only Go binding around an Ethereum contract.
type AimiddlewareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AimiddlewareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AimiddlewareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AimiddlewareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AimiddlewareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AimiddlewareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AimiddlewareSession struct {
	Contract     *Aimiddleware     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AimiddlewareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AimiddlewareCallerSession struct {
	Contract *AimiddlewareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AimiddlewareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AimiddlewareTransactorSession struct {
	Contract     *AimiddlewareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AimiddlewareRaw is an auto generated low-level Go binding around an Ethereum contract.
type AimiddlewareRaw struct {
	Contract *Aimiddleware // Generic contract binding to access the raw methods on
}

// AimiddlewareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AimiddlewareCallerRaw struct {
	Contract *AimiddlewareCaller // Generic read-only contract binding to access the raw methods on
}

// AimiddlewareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AimiddlewareTransactorRaw struct {
	Contract *AimiddlewareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAimiddleware creates a new instance of Aimiddleware, bound to a specific deployed contract.
func NewAimiddleware(address common.Address, backend bind.ContractBackend) (*Aimiddleware, error) {
	contract, err := bindAimiddleware(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aimiddleware{AimiddlewareCaller: AimiddlewareCaller{contract: contract}, AimiddlewareTransactor: AimiddlewareTransactor{contract: contract}, AimiddlewareFilterer: AimiddlewareFilterer{contract: contract}}, nil
}

// NewAimiddlewareCaller creates a new read-only instance of Aimiddleware, bound to a specific deployed contract.
func NewAimiddlewareCaller(address common.Address, caller bind.ContractCaller) (*AimiddlewareCaller, error) {
	contract, err := bindAimiddleware(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareCaller{contract: contract}, nil
}

// NewAimiddlewareTransactor creates a new write-only instance of Aimiddleware, bound to a specific deployed contract.
func NewAimiddlewareTransactor(address common.Address, transactor bind.ContractTransactor) (*AimiddlewareTransactor, error) {
	contract, err := bindAimiddleware(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareTransactor{contract: contract}, nil
}

// NewAimiddlewareFilterer creates a new log filterer instance of Aimiddleware, bound to a specific deployed contract.
func NewAimiddlewareFilterer(address common.Address, filterer bind.ContractFilterer) (*AimiddlewareFilterer, error) {
	contract, err := bindAimiddleware(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareFilterer{contract: contract}, nil
}

// bindAimiddleware binds a generic wrapper to an already deployed contract.
func bindAimiddleware(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AimiddlewareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aimiddleware *AimiddlewareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aimiddleware.Contract.AimiddlewareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aimiddleware *AimiddlewareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aimiddleware.Contract.AimiddlewareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aimiddleware *AimiddlewareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aimiddleware.Contract.AimiddlewareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aimiddleware *AimiddlewareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aimiddleware.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aimiddleware *AimiddlewareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aimiddleware.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aimiddleware *AimiddlewareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aimiddleware.Contract.contract.Transact(opts, method, params...)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Aimiddleware *AimiddlewareCaller) SIGVALIDATIONFAILED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aimiddleware.contract.Call(opts, &out, "SIG_VALIDATION_FAILED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Aimiddleware *AimiddlewareSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _Aimiddleware.Contract.SIGVALIDATIONFAILED(&_Aimiddleware.CallOpts)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_Aimiddleware *AimiddlewareCallerSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _Aimiddleware.Contract.SIGVALIDATIONFAILED(&_Aimiddleware.CallOpts)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Aimiddleware *AimiddlewareCaller) ValidateSenderAndPaymaster(opts *bind.CallOpts, initCode []byte, sender common.Address, paymasterAndData []byte) error {
	var out []interface{}
	err := _Aimiddleware.contract.Call(opts, &out, "_validateSenderAndPaymaster", initCode, sender, paymasterAndData)

	if err != nil {
		return err
	}

	return err

}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Aimiddleware *AimiddlewareSession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _Aimiddleware.Contract.ValidateSenderAndPaymaster(&_Aimiddleware.CallOpts, initCode, sender, paymasterAndData)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_Aimiddleware *AimiddlewareCallerSession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _Aimiddleware.Contract.ValidateSenderAndPaymaster(&_Aimiddleware.CallOpts, initCode, sender, paymasterAndData)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Aimiddleware *AimiddlewareCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aimiddleware.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Aimiddleware *AimiddlewareSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Aimiddleware.Contract.BalanceOf(&_Aimiddleware.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Aimiddleware *AimiddlewareCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Aimiddleware.Contract.BalanceOf(&_Aimiddleware.CallOpts, account)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Aimiddleware *AimiddlewareCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	var out []interface{}
	err := _Aimiddleware.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Deposit         *big.Int
		Staked          bool
		Stake           *big.Int
		UnstakeDelaySec uint32
		WithdrawTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Staked = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeDelaySec = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.WithdrawTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Aimiddleware *AimiddlewareSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _Aimiddleware.Contract.Deposits(&_Aimiddleware.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_Aimiddleware *AimiddlewareCallerSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _Aimiddleware.Contract.Deposits(&_Aimiddleware.CallOpts, arg0)
}

// GetAiOpHash is a free data retrieval call binding the contract method 0x6960c859.
//
// Solidity: function getAiOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) aiOp) view returns(bytes32)
func (_Aimiddleware *AimiddlewareCaller) GetAiOpHash(opts *bind.CallOpts, aiOp AiOperation) ([32]byte, error) {
	var out []interface{}
	err := _Aimiddleware.contract.Call(opts, &out, "getAiOpHash", aiOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAiOpHash is a free data retrieval call binding the contract method 0x6960c859.
//
// Solidity: function getAiOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) aiOp) view returns(bytes32)
func (_Aimiddleware *AimiddlewareSession) GetAiOpHash(aiOp AiOperation) ([32]byte, error) {
	return _Aimiddleware.Contract.GetAiOpHash(&_Aimiddleware.CallOpts, aiOp)
}

// GetAiOpHash is a free data retrieval call binding the contract method 0x6960c859.
//
// Solidity: function getAiOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) aiOp) view returns(bytes32)
func (_Aimiddleware *AimiddlewareCallerSession) GetAiOpHash(aiOp AiOperation) ([32]byte, error) {
	return _Aimiddleware.Contract.GetAiOpHash(&_Aimiddleware.CallOpts, aiOp)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Aimiddleware *AimiddlewareCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IDepositManagerDepositInfo, error) {
	var out []interface{}
	err := _Aimiddleware.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IDepositManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDepositManagerDepositInfo)).(*IDepositManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Aimiddleware *AimiddlewareSession) GetDepositInfo(account common.Address) (IDepositManagerDepositInfo, error) {
	return _Aimiddleware.Contract.GetDepositInfo(&_Aimiddleware.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_Aimiddleware *AimiddlewareCallerSession) GetDepositInfo(account common.Address) (IDepositManagerDepositInfo, error) {
	return _Aimiddleware.Contract.GetDepositInfo(&_Aimiddleware.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Aimiddleware *AimiddlewareCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Aimiddleware.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Aimiddleware *AimiddlewareSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _Aimiddleware.Contract.GetNonce(&_Aimiddleware.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Aimiddleware *AimiddlewareCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _Aimiddleware.Contract.GetNonce(&_Aimiddleware.CallOpts, sender, key)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Aimiddleware *AimiddlewareCaller) NonceSequenceNumber(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Aimiddleware.contract.Call(opts, &out, "nonceSequenceNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Aimiddleware *AimiddlewareSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Aimiddleware.Contract.NonceSequenceNumber(&_Aimiddleware.CallOpts, arg0, arg1)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_Aimiddleware *AimiddlewareCallerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Aimiddleware.Contract.NonceSequenceNumber(&_Aimiddleware.CallOpts, arg0, arg1)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Aimiddleware *AimiddlewareTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Aimiddleware *AimiddlewareSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Aimiddleware.Contract.AddStake(&_Aimiddleware.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Aimiddleware *AimiddlewareTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Aimiddleware.Contract.AddStake(&_Aimiddleware.TransactOpts, unstakeDelaySec)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Aimiddleware *AimiddlewareTransactor) DepositTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "depositTo", account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Aimiddleware *AimiddlewareSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.DepositTo(&_Aimiddleware.TransactOpts, account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_Aimiddleware *AimiddlewareTransactorSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.DepositTo(&_Aimiddleware.TransactOpts, account)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Aimiddleware *AimiddlewareTransactor) GetSenderAddress(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "getSenderAddress", initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Aimiddleware *AimiddlewareSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _Aimiddleware.Contract.GetSenderAddress(&_Aimiddleware.TransactOpts, initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _Aimiddleware.Contract.GetSenderAddress(&_Aimiddleware.TransactOpts, initCode)
}

// GetValidator is a paid mutator transaction binding the contract method 0x1195e07e.
//
// Solidity: function getValidator() returns(address)
func (_Aimiddleware *AimiddlewareTransactor) GetValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "getValidator")
}

// GetValidator is a paid mutator transaction binding the contract method 0x1195e07e.
//
// Solidity: function getValidator() returns(address)
func (_Aimiddleware *AimiddlewareSession) GetValidator() (*types.Transaction, error) {
	return _Aimiddleware.Contract.GetValidator(&_Aimiddleware.TransactOpts)
}

// GetValidator is a paid mutator transaction binding the contract method 0x1195e07e.
//
// Solidity: function getValidator() returns(address)
func (_Aimiddleware *AimiddlewareTransactorSession) GetValidator() (*types.Transaction, error) {
	return _Aimiddleware.Contract.GetValidator(&_Aimiddleware.TransactOpts)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Aimiddleware *AimiddlewareTransactor) HandleAggregatedOps(opts *bind.TransactOpts, opsPerAggregator []IAiMiddlewareAiOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "handleAggregatedOps", opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Aimiddleware *AimiddlewareSession) HandleAggregatedOps(opsPerAggregator []IAiMiddlewareAiOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.HandleAggregatedOps(&_Aimiddleware.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) HandleAggregatedOps(opsPerAggregator []IAiMiddlewareAiOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.HandleAggregatedOps(&_Aimiddleware.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Aimiddleware *AimiddlewareTransactor) HandleOps(opts *bind.TransactOpts, ops []AiOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Aimiddleware *AimiddlewareSession) HandleOps(ops []AiOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.HandleOps(&_Aimiddleware.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) HandleOps(ops []AiOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.HandleOps(&_Aimiddleware.TransactOpts, ops, beneficiary)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Aimiddleware *AimiddlewareTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Aimiddleware *AimiddlewareSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _Aimiddleware.Contract.IncrementNonce(&_Aimiddleware.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _Aimiddleware.Contract.IncrementNonce(&_Aimiddleware.TransactOpts, key)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Aimiddleware *AimiddlewareTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo AiMiddlewareAiOpInfo, context []byte) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Aimiddleware *AimiddlewareSession) InnerHandleOp(callData []byte, opInfo AiMiddlewareAiOpInfo, context []byte) (*types.Transaction, error) {
	return _Aimiddleware.Contract.InnerHandleOp(&_Aimiddleware.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_Aimiddleware *AimiddlewareTransactorSession) InnerHandleOp(callData []byte, opInfo AiMiddlewareAiOpInfo, context []byte) (*types.Transaction, error) {
	return _Aimiddleware.Contract.InnerHandleOp(&_Aimiddleware.TransactOpts, callData, opInfo, context)
}

// ModifyValidator is a paid mutator transaction binding the contract method 0x739429a5.
//
// Solidity: function modifyValidator(address _validator) returns()
func (_Aimiddleware *AimiddlewareTransactor) ModifyValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "modifyValidator", _validator)
}

// ModifyValidator is a paid mutator transaction binding the contract method 0x739429a5.
//
// Solidity: function modifyValidator(address _validator) returns()
func (_Aimiddleware *AimiddlewareSession) ModifyValidator(_validator common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.ModifyValidator(&_Aimiddleware.TransactOpts, _validator)
}

// ModifyValidator is a paid mutator transaction binding the contract method 0x739429a5.
//
// Solidity: function modifyValidator(address _validator) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) ModifyValidator(_validator common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.ModifyValidator(&_Aimiddleware.TransactOpts, _validator)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Aimiddleware *AimiddlewareTransactor) SimulateHandleOp(opts *bind.TransactOpts, op AiOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "simulateHandleOp", op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Aimiddleware *AimiddlewareSession) SimulateHandleOp(op AiOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Aimiddleware.Contract.SimulateHandleOp(&_Aimiddleware.TransactOpts, op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) SimulateHandleOp(op AiOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _Aimiddleware.Contract.SimulateHandleOp(&_Aimiddleware.TransactOpts, op, target, targetCallData)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) aiOp) returns()
func (_Aimiddleware *AimiddlewareTransactor) SimulateValidation(opts *bind.TransactOpts, aiOp AiOperation) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "simulateValidation", aiOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) aiOp) returns()
func (_Aimiddleware *AimiddlewareSession) SimulateValidation(aiOp AiOperation) (*types.Transaction, error) {
	return _Aimiddleware.Contract.SimulateValidation(&_Aimiddleware.TransactOpts, aiOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) aiOp) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) SimulateValidation(aiOp AiOperation) (*types.Transaction, error) {
	return _Aimiddleware.Contract.SimulateValidation(&_Aimiddleware.TransactOpts, aiOp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _owner) returns()
func (_Aimiddleware *AimiddlewareTransactor) TransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "transferOwnership", _owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _owner) returns()
func (_Aimiddleware *AimiddlewareSession) TransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.TransferOwnership(&_Aimiddleware.TransactOpts, _owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _owner) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) TransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.TransferOwnership(&_Aimiddleware.TransactOpts, _owner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Aimiddleware *AimiddlewareTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Aimiddleware *AimiddlewareSession) UnlockStake() (*types.Transaction, error) {
	return _Aimiddleware.Contract.UnlockStake(&_Aimiddleware.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Aimiddleware *AimiddlewareTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _Aimiddleware.Contract.UnlockStake(&_Aimiddleware.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Aimiddleware *AimiddlewareTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Aimiddleware *AimiddlewareSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.WithdrawStake(&_Aimiddleware.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Aimiddleware.Contract.WithdrawStake(&_Aimiddleware.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Aimiddleware *AimiddlewareTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Aimiddleware.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Aimiddleware *AimiddlewareSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Aimiddleware.Contract.WithdrawTo(&_Aimiddleware.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_Aimiddleware *AimiddlewareTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Aimiddleware.Contract.WithdrawTo(&_Aimiddleware.TransactOpts, withdrawAddress, withdrawAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aimiddleware *AimiddlewareTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aimiddleware.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aimiddleware *AimiddlewareSession) Receive() (*types.Transaction, error) {
	return _Aimiddleware.Contract.Receive(&_Aimiddleware.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aimiddleware *AimiddlewareTransactorSession) Receive() (*types.Transaction, error) {
	return _Aimiddleware.Contract.Receive(&_Aimiddleware.TransactOpts)
}

// AimiddlewareAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the Aimiddleware contract.
type AimiddlewareAccountDeployedIterator struct {
	Event *AimiddlewareAccountDeployed // Event containing the contract specifics and raw log

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
func (it *AimiddlewareAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareAccountDeployed)
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
		it.Event = new(AimiddlewareAccountDeployed)
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
func (it *AimiddlewareAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareAccountDeployed represents a AccountDeployed event raised by the Aimiddleware contract.
type AimiddlewareAccountDeployed struct {
	AiOpHash  [32]byte
	Sender    common.Address
	Factory   common.Address
	Paymaster common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed aiOpHash, address indexed sender, address factory, address paymaster)
func (_Aimiddleware *AimiddlewareFilterer) FilterAccountDeployed(opts *bind.FilterOpts, aiOpHash [][32]byte, sender []common.Address) (*AimiddlewareAccountDeployedIterator, error) {

	var aiOpHashRule []interface{}
	for _, aiOpHashItem := range aiOpHash {
		aiOpHashRule = append(aiOpHashRule, aiOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "AccountDeployed", aiOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareAccountDeployedIterator{contract: _Aimiddleware.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed aiOpHash, address indexed sender, address factory, address paymaster)
func (_Aimiddleware *AimiddlewareFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *AimiddlewareAccountDeployed, aiOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var aiOpHashRule []interface{}
	for _, aiOpHashItem := range aiOpHash {
		aiOpHashRule = append(aiOpHashRule, aiOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "AccountDeployed", aiOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareAccountDeployed)
				if err := _Aimiddleware.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
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

// ParseAccountDeployed is a log parse operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed aiOpHash, address indexed sender, address factory, address paymaster)
func (_Aimiddleware *AimiddlewareFilterer) ParseAccountDeployed(log types.Log) (*AimiddlewareAccountDeployed, error) {
	event := new(AimiddlewareAccountDeployed)
	if err := _Aimiddleware.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareAiOperationEventIterator is returned from FilterAiOperationEvent and is used to iterate over the raw logs and unpacked data for AiOperationEvent events raised by the Aimiddleware contract.
type AimiddlewareAiOperationEventIterator struct {
	Event *AimiddlewareAiOperationEvent // Event containing the contract specifics and raw log

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
func (it *AimiddlewareAiOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareAiOperationEvent)
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
		it.Event = new(AimiddlewareAiOperationEvent)
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
func (it *AimiddlewareAiOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareAiOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareAiOperationEvent represents a AiOperationEvent event raised by the Aimiddleware contract.
type AimiddlewareAiOperationEvent struct {
	AiOpHash      [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAiOperationEvent is a free log retrieval operation binding the contract event 0x7b592a4f684d032578f75dff2ca3d3d2aec981a6d0a782e1d127595a6511a4f1.
//
// Solidity: event AiOperationEvent(bytes32 indexed aiOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Aimiddleware *AimiddlewareFilterer) FilterAiOperationEvent(opts *bind.FilterOpts, aiOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (*AimiddlewareAiOperationEventIterator, error) {

	var aiOpHashRule []interface{}
	for _, aiOpHashItem := range aiOpHash {
		aiOpHashRule = append(aiOpHashRule, aiOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "AiOperationEvent", aiOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareAiOperationEventIterator{contract: _Aimiddleware.contract, event: "AiOperationEvent", logs: logs, sub: sub}, nil
}

// WatchAiOperationEvent is a free log subscription operation binding the contract event 0x7b592a4f684d032578f75dff2ca3d3d2aec981a6d0a782e1d127595a6511a4f1.
//
// Solidity: event AiOperationEvent(bytes32 indexed aiOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Aimiddleware *AimiddlewareFilterer) WatchAiOperationEvent(opts *bind.WatchOpts, sink chan<- *AimiddlewareAiOperationEvent, aiOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (event.Subscription, error) {

	var aiOpHashRule []interface{}
	for _, aiOpHashItem := range aiOpHash {
		aiOpHashRule = append(aiOpHashRule, aiOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "AiOperationEvent", aiOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareAiOperationEvent)
				if err := _Aimiddleware.contract.UnpackLog(event, "AiOperationEvent", log); err != nil {
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

// ParseAiOperationEvent is a log parse operation binding the contract event 0x7b592a4f684d032578f75dff2ca3d3d2aec981a6d0a782e1d127595a6511a4f1.
//
// Solidity: event AiOperationEvent(bytes32 indexed aiOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_Aimiddleware *AimiddlewareFilterer) ParseAiOperationEvent(log types.Log) (*AimiddlewareAiOperationEvent, error) {
	event := new(AimiddlewareAiOperationEvent)
	if err := _Aimiddleware.contract.UnpackLog(event, "AiOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareAiOperationRevertReasonIterator is returned from FilterAiOperationRevertReason and is used to iterate over the raw logs and unpacked data for AiOperationRevertReason events raised by the Aimiddleware contract.
type AimiddlewareAiOperationRevertReasonIterator struct {
	Event *AimiddlewareAiOperationRevertReason // Event containing the contract specifics and raw log

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
func (it *AimiddlewareAiOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareAiOperationRevertReason)
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
		it.Event = new(AimiddlewareAiOperationRevertReason)
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
func (it *AimiddlewareAiOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareAiOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareAiOperationRevertReason represents a AiOperationRevertReason event raised by the Aimiddleware contract.
type AimiddlewareAiOperationRevertReason struct {
	AiOpHash     [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAiOperationRevertReason is a free log retrieval operation binding the contract event 0xd5425d9883d2972240bba93dc9098c4f8a304a1a73d0c8b5a7276ce2994f9bb2.
//
// Solidity: event AiOperationRevertReason(bytes32 indexed aiOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Aimiddleware *AimiddlewareFilterer) FilterAiOperationRevertReason(opts *bind.FilterOpts, aiOpHash [][32]byte, sender []common.Address) (*AimiddlewareAiOperationRevertReasonIterator, error) {

	var aiOpHashRule []interface{}
	for _, aiOpHashItem := range aiOpHash {
		aiOpHashRule = append(aiOpHashRule, aiOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "AiOperationRevertReason", aiOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareAiOperationRevertReasonIterator{contract: _Aimiddleware.contract, event: "AiOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchAiOperationRevertReason is a free log subscription operation binding the contract event 0xd5425d9883d2972240bba93dc9098c4f8a304a1a73d0c8b5a7276ce2994f9bb2.
//
// Solidity: event AiOperationRevertReason(bytes32 indexed aiOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Aimiddleware *AimiddlewareFilterer) WatchAiOperationRevertReason(opts *bind.WatchOpts, sink chan<- *AimiddlewareAiOperationRevertReason, aiOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var aiOpHashRule []interface{}
	for _, aiOpHashItem := range aiOpHash {
		aiOpHashRule = append(aiOpHashRule, aiOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "AiOperationRevertReason", aiOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareAiOperationRevertReason)
				if err := _Aimiddleware.contract.UnpackLog(event, "AiOperationRevertReason", log); err != nil {
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

// ParseAiOperationRevertReason is a log parse operation binding the contract event 0xd5425d9883d2972240bba93dc9098c4f8a304a1a73d0c8b5a7276ce2994f9bb2.
//
// Solidity: event AiOperationRevertReason(bytes32 indexed aiOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_Aimiddleware *AimiddlewareFilterer) ParseAiOperationRevertReason(log types.Log) (*AimiddlewareAiOperationRevertReason, error) {
	event := new(AimiddlewareAiOperationRevertReason)
	if err := _Aimiddleware.contract.UnpackLog(event, "AiOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareBeforeAiopsExecutionIterator is returned from FilterBeforeAiopsExecution and is used to iterate over the raw logs and unpacked data for BeforeAiopsExecution events raised by the Aimiddleware contract.
type AimiddlewareBeforeAiopsExecutionIterator struct {
	Event *AimiddlewareBeforeAiopsExecution // Event containing the contract specifics and raw log

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
func (it *AimiddlewareBeforeAiopsExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareBeforeAiopsExecution)
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
		it.Event = new(AimiddlewareBeforeAiopsExecution)
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
func (it *AimiddlewareBeforeAiopsExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareBeforeAiopsExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareBeforeAiopsExecution represents a BeforeAiopsExecution event raised by the Aimiddleware contract.
type AimiddlewareBeforeAiopsExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeAiopsExecution is a free log retrieval operation binding the contract event 0xd22ebd07844ae9e9c8d9945236ef70e6c2f2f0c5bbeef9980c0384600d424243.
//
// Solidity: event BeforeAiopsExecution()
func (_Aimiddleware *AimiddlewareFilterer) FilterBeforeAiopsExecution(opts *bind.FilterOpts) (*AimiddlewareBeforeAiopsExecutionIterator, error) {

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "BeforeAiopsExecution")
	if err != nil {
		return nil, err
	}
	return &AimiddlewareBeforeAiopsExecutionIterator{contract: _Aimiddleware.contract, event: "BeforeAiopsExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeAiopsExecution is a free log subscription operation binding the contract event 0xd22ebd07844ae9e9c8d9945236ef70e6c2f2f0c5bbeef9980c0384600d424243.
//
// Solidity: event BeforeAiopsExecution()
func (_Aimiddleware *AimiddlewareFilterer) WatchBeforeAiopsExecution(opts *bind.WatchOpts, sink chan<- *AimiddlewareBeforeAiopsExecution) (event.Subscription, error) {

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "BeforeAiopsExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareBeforeAiopsExecution)
				if err := _Aimiddleware.contract.UnpackLog(event, "BeforeAiopsExecution", log); err != nil {
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

// ParseBeforeAiopsExecution is a log parse operation binding the contract event 0xd22ebd07844ae9e9c8d9945236ef70e6c2f2f0c5bbeef9980c0384600d424243.
//
// Solidity: event BeforeAiopsExecution()
func (_Aimiddleware *AimiddlewareFilterer) ParseBeforeAiopsExecution(log types.Log) (*AimiddlewareBeforeAiopsExecution, error) {
	event := new(AimiddlewareBeforeAiopsExecution)
	if err := _Aimiddleware.contract.UnpackLog(event, "BeforeAiopsExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Aimiddleware contract.
type AimiddlewareDepositedIterator struct {
	Event *AimiddlewareDeposited // Event containing the contract specifics and raw log

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
func (it *AimiddlewareDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareDeposited)
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
		it.Event = new(AimiddlewareDeposited)
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
func (it *AimiddlewareDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareDeposited represents a Deposited event raised by the Aimiddleware contract.
type AimiddlewareDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Aimiddleware *AimiddlewareFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*AimiddlewareDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareDepositedIterator{contract: _Aimiddleware.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Aimiddleware *AimiddlewareFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *AimiddlewareDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareDeposited)
				if err := _Aimiddleware.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_Aimiddleware *AimiddlewareFilterer) ParseDeposited(log types.Log) (*AimiddlewareDeposited, error) {
	event := new(AimiddlewareDeposited)
	if err := _Aimiddleware.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareSignatureAggregatorChangedIterator is returned from FilterSignatureAggregatorChanged and is used to iterate over the raw logs and unpacked data for SignatureAggregatorChanged events raised by the Aimiddleware contract.
type AimiddlewareSignatureAggregatorChangedIterator struct {
	Event *AimiddlewareSignatureAggregatorChanged // Event containing the contract specifics and raw log

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
func (it *AimiddlewareSignatureAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareSignatureAggregatorChanged)
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
		it.Event = new(AimiddlewareSignatureAggregatorChanged)
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
func (it *AimiddlewareSignatureAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareSignatureAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareSignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the Aimiddleware contract.
type AimiddlewareSignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignatureAggregatorChanged is a free log retrieval operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Aimiddleware *AimiddlewareFilterer) FilterSignatureAggregatorChanged(opts *bind.FilterOpts, aggregator []common.Address) (*AimiddlewareSignatureAggregatorChangedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareSignatureAggregatorChangedIterator{contract: _Aimiddleware.contract, event: "SignatureAggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchSignatureAggregatorChanged is a free log subscription operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Aimiddleware *AimiddlewareFilterer) WatchSignatureAggregatorChanged(opts *bind.WatchOpts, sink chan<- *AimiddlewareSignatureAggregatorChanged, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareSignatureAggregatorChanged)
				if err := _Aimiddleware.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
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

// ParseSignatureAggregatorChanged is a log parse operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_Aimiddleware *AimiddlewareFilterer) ParseSignatureAggregatorChanged(log types.Log) (*AimiddlewareSignatureAggregatorChanged, error) {
	event := new(AimiddlewareSignatureAggregatorChanged)
	if err := _Aimiddleware.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the Aimiddleware contract.
type AimiddlewareStakeLockedIterator struct {
	Event *AimiddlewareStakeLocked // Event containing the contract specifics and raw log

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
func (it *AimiddlewareStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareStakeLocked)
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
		it.Event = new(AimiddlewareStakeLocked)
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
func (it *AimiddlewareStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareStakeLocked represents a StakeLocked event raised by the Aimiddleware contract.
type AimiddlewareStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Aimiddleware *AimiddlewareFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*AimiddlewareStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareStakeLockedIterator{contract: _Aimiddleware.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Aimiddleware *AimiddlewareFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *AimiddlewareStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareStakeLocked)
				if err := _Aimiddleware.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_Aimiddleware *AimiddlewareFilterer) ParseStakeLocked(log types.Log) (*AimiddlewareStakeLocked, error) {
	event := new(AimiddlewareStakeLocked)
	if err := _Aimiddleware.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the Aimiddleware contract.
type AimiddlewareStakeUnlockedIterator struct {
	Event *AimiddlewareStakeUnlocked // Event containing the contract specifics and raw log

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
func (it *AimiddlewareStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareStakeUnlocked)
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
		it.Event = new(AimiddlewareStakeUnlocked)
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
func (it *AimiddlewareStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareStakeUnlocked represents a StakeUnlocked event raised by the Aimiddleware contract.
type AimiddlewareStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Aimiddleware *AimiddlewareFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*AimiddlewareStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareStakeUnlockedIterator{contract: _Aimiddleware.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Aimiddleware *AimiddlewareFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *AimiddlewareStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareStakeUnlocked)
				if err := _Aimiddleware.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
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

// ParseStakeUnlocked is a log parse operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_Aimiddleware *AimiddlewareFilterer) ParseStakeUnlocked(log types.Log) (*AimiddlewareStakeUnlocked, error) {
	event := new(AimiddlewareStakeUnlocked)
	if err := _Aimiddleware.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Aimiddleware contract.
type AimiddlewareStakeWithdrawnIterator struct {
	Event *AimiddlewareStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *AimiddlewareStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareStakeWithdrawn)
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
		it.Event = new(AimiddlewareStakeWithdrawn)
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
func (it *AimiddlewareStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareStakeWithdrawn represents a StakeWithdrawn event raised by the Aimiddleware contract.
type AimiddlewareStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Aimiddleware *AimiddlewareFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*AimiddlewareStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareStakeWithdrawnIterator{contract: _Aimiddleware.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Aimiddleware *AimiddlewareFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *AimiddlewareStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareStakeWithdrawn)
				if err := _Aimiddleware.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Aimiddleware *AimiddlewareFilterer) ParseStakeWithdrawn(log types.Log) (*AimiddlewareStakeWithdrawn, error) {
	event := new(AimiddlewareStakeWithdrawn)
	if err := _Aimiddleware.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AimiddlewareWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Aimiddleware contract.
type AimiddlewareWithdrawnIterator struct {
	Event *AimiddlewareWithdrawn // Event containing the contract specifics and raw log

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
func (it *AimiddlewareWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AimiddlewareWithdrawn)
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
		it.Event = new(AimiddlewareWithdrawn)
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
func (it *AimiddlewareWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AimiddlewareWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AimiddlewareWithdrawn represents a Withdrawn event raised by the Aimiddleware contract.
type AimiddlewareWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Aimiddleware *AimiddlewareFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*AimiddlewareWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &AimiddlewareWithdrawnIterator{contract: _Aimiddleware.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Aimiddleware *AimiddlewareFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *AimiddlewareWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Aimiddleware.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AimiddlewareWithdrawn)
				if err := _Aimiddleware.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_Aimiddleware *AimiddlewareFilterer) ParseWithdrawn(log types.Log) (*AimiddlewareWithdrawn, error) {
	event := new(AimiddlewareWithdrawn)
	if err := _Aimiddleware.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
