package aiop_test

import (
	"math/big"
	"testing"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/testutils"
)

// TestAiOperationGetDynamicGasPrice verifies that (*AiOperation).GetDynamicGasPrice returns the correct
// effective gas price given a base fee.
func TestAiOperationGetDynamicGasPrice(t *testing.T) {
	bf := big.NewInt(3)
	op := testutils.MockValidInitAiOp()

	// basefee + MPF > MF
	want := big.NewInt(4)
	op.MaxFeePerGas = big.NewInt(4)
	op.MaxPriorityFeePerGas = big.NewInt(3)
	if op.GetDynamicGasPrice(bf).Cmp(want) != 0 {
		t.Fatalf("got %d, want %d", op.GetDynamicGasPrice(bf).Int64(), want.Int64())
	}

	// basefee + MPF == MF
	want = big.NewInt(5)
	op.MaxFeePerGas = big.NewInt(5)
	op.MaxPriorityFeePerGas = big.NewInt(2)
	if op.GetDynamicGasPrice(bf).Cmp(want) != 0 {
		t.Fatalf("got %d, want %d", op.GetDynamicGasPrice(bf).Int64(), want.Int64())
	}

	// basefee + MPF < MF
	want = big.NewInt(4)
	op.MaxFeePerGas = big.NewInt(6)
	op.MaxPriorityFeePerGas = big.NewInt(1)
	if op.GetDynamicGasPrice(bf).Cmp(want) != 0 {
		t.Fatalf("got %d, want %d", op.GetDynamicGasPrice(bf).Int64(), want.Int64())
	}
}

// TestAiOperationGetGasPriceNilBF verifies that (*AiOperation).GetDynamicGasPrice returns the correct
// value when basefee is nil.
func TestAiOperationGetGasPriceNilBF(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.MaxFeePerGas = big.NewInt(4)
	op.MaxPriorityFeePerGas = big.NewInt(3)
	if op.GetDynamicGasPrice(nil).Cmp(op.MaxPriorityFeePerGas) != 0 {
		t.Fatalf("got %d, want %d", op.GetDynamicGasPrice(nil).Int64(), op.MaxPriorityFeePerGas)
	}
}
