// Package gas implements helper functions for calculating EIP-4337 gas parameters.
package gas

import (
	"bytes"
	"math"
	"math/big"

	"github.com/AO-Metaplayer/aiops-bundler/internal/utils"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Overhead provides helper methods for calculating gas limits based on pre-defined parameters.
type Overhead struct {
	intrinsicFixed      float64
	perAiOpFixed        float64
	perAiOpMultiplier   float64
	zeroByte            float64
	nonZeroByte         float64
	minBundleSize       float64
	warmStorageRead     float64
	callWithValue       float64
	callOpcode          float64
	nonZeroValueStipend float64
	sanitizedPVG        *big.Int
	sanitizedVGL        *big.Int
	sanitizedCGL        *big.Int
	calcPVGFunc         CalcPreVerificationGasFunc
	pvgBufferFactor     int64
}

// NewDefaultOverhead returns an instance of Overhead using parameters defined by the Ethereum protocol.
func NewDefaultOverhead() *Overhead {
	return &Overhead{
		intrinsicFixed:      21000,
		perAiOpFixed:        22874,
		perAiOpMultiplier:   25,
		zeroByte:            4,
		nonZeroByte:         16,
		minBundleSize:       1,
		warmStorageRead:     100,
		callWithValue:       9000,
		callOpcode:          700,
		nonZeroValueStipend: 2300,
		sanitizedPVG:        big.NewInt(100000),
		sanitizedVGL:        big.NewInt(1000000),
		sanitizedCGL:        big.NewInt(1000000),
		calcPVGFunc:         calcPVGFuncNoop(),
		pvgBufferFactor:     0,
	}
}

// SetCalcPreVerificationGasFunc allows a custom function to be defined that can control how it calculates
// PVG. This is useful for networks that have different models for gas.
func (ov *Overhead) SetCalcPreVerificationGasFunc(fn CalcPreVerificationGasFunc) {
	ov.calcPVGFunc = fn
}

// SetPreVerificationGasBufferFactor defines the percentage to increase the preVerificationGas by during an
// estimation. This is useful for rollups that use 2D gas values where the L1 gas component is
// non-deterministic. This buffer accounts for any variability in-between eth_estimateAiOperationGas and
// eth_sendAiOperation. Defaults to 0.
func (ov *Overhead) SetPreVerificationGasBufferFactor(factor int64) {
	ov.pvgBufferFactor = factor
}

// CalcCallDataCost calculates the additional gas cost required to serialize the aiOp when making the
// transaction to submit the entire batch.
func (ov *Overhead) CalcCallDataCost(op *aiop.AiOperation) float64 {
	cost := float64(0)
	for _, b := range op.Pack() {
		if b == byte(0) {
			cost += ov.zeroByte
		} else {
			cost += ov.nonZeroByte
		}
	}

	return cost
}

// CalcPerAiOpCost calculates the gas overhead from processing a AiOperation's validation and execution
// phase. This overhead is not constant and is correlated to the number of 32 byte words in the AiOperation.
// It can be summarized in the equation perAiOpMultiplier * lenInWord + perAiOpFixed.
//
// Note: The constant values have been derived empirically by plotting the relationship between per aiOp
// overhead vs length in words with a sample size of 30.
func (ov *Overhead) CalcPerAiOpCost(op *aiop.AiOperation) float64 {
	opLen := math.Floor(float64(len(op.Pack())+31) / 32)
	cost := (ov.perAiOpMultiplier * opLen) + ov.perAiOpFixed

	return cost
}

// CalcPreVerificationGas returns an expected gas cost for processing a AiOperation from a batch.
func (ov *Overhead) CalcPreVerificationGas(op *aiop.AiOperation) (*big.Int, error) {
	// Sanitize fields to reduce as much variability due to length and zero bytes
	data, err := op.ToMap()
	if err != nil {
		return nil, err
	}
	data["preVerificationGas"] = hexutil.EncodeBig(ov.sanitizedPVG)
	data["verificationGasLimit"] = hexutil.EncodeBig(ov.sanitizedVGL)
	data["callGasLimit"] = hexutil.EncodeBig(ov.sanitizedCGL)
	data["signature"] = hexutil.Encode(bytes.Repeat([]byte{1}, len(op.Signature)))
	tmp, err := aiop.New(data)
	if err != nil {
		return nil, err
	}

	// Calculate the additional gas for adding this aiOp to a batch.
	batchOv := (ov.intrinsicFixed / ov.minBundleSize) + ov.CalcCallDataCost(tmp)

	// The total PVG is the sum of the batch overhead and the overhead for this aiOp's validation and
	// execution.
	pvg := batchOv + ov.CalcPerAiOpCost(tmp)
	static := big.NewInt(int64(math.Round(pvg)))

	// Use value from CalcPreVerificationGasFunc if set, otherwise return the static value.
	g, err := ov.calcPVGFunc(tmp, static)
	if err != nil {
		return nil, err
	}
	if g != nil {
		return g, nil
	}
	return static, nil
}

// CalcPreVerificationGasWithBuffer returns CalcPreVerificationGas increased by the set PVG buffer factor.
func (ov *Overhead) CalcPreVerificationGasWithBuffer(op *aiop.AiOperation) (*big.Int, error) {
	pvg, err := ov.CalcPreVerificationGas(op)
	if err != nil {
		return nil, err
	}
	return utils.AddBuffer(pvg, ov.pvgBufferFactor), nil
}

// NonZeroValueCall returns an expected gas cost of using the CALL opcode with non-zero value.
// See https://github.com/wolflo/evm-opcodes/blob/main/gas.md#aa-1-call.
func (ov *Overhead) NonZeroValueCall() *big.Int {
	return big.NewInt(
		int64(
			ov.callOpcode + ov.callWithValue + ov.warmStorageRead + ov.nonZeroValueStipend,
		),
	)
}
