package checks

import (
	"errors"
	"math/big"
	"testing"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/testutils"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
)

func TestNoPendingOps(t *testing.T) {
	penOps := []*aiop.AiOperation{}
	op := testutils.MockValidInitAiOp()
	err := ValidatePendingOps(
		op,
		penOps,
	)

	if err != nil {
		t.Fatalf("got err %v, want nil", err)
	}
}

func TestPendingOpsWithNewOp(t *testing.T) {
	penOp := testutils.MockValidInitAiOp()
	penOps := []*aiop.AiOperation{penOp}
	op := testutils.MockValidInitAiOp()
	op.Nonce = big.NewInt(1)
	err := ValidatePendingOps(
		op,
		penOps,
	)

	if err != nil {
		t.Fatalf("got err %v, want nil", err)
	}
}

func TestPendingOpsWithNoGasFeeReplacement(t *testing.T) {
	penOp := testutils.MockValidInitAiOp()
	penOps := []*aiop.AiOperation{penOp}
	op := testutils.MockValidInitAiOp()
	err := ValidatePendingOps(
		op,
		penOps,
	)

	if !errors.Is(err, ErrReplacementOpUnderpriced) {
		t.Fatalf("got %v, want ErrReplacementOpUnderpriced", err)
	}
}

func TestPendingOpsWithOnlyMaxFeeReplacement(t *testing.T) {
	penOp := testutils.MockValidInitAiOp()
	penOps := []*aiop.AiOperation{penOp}
	op := testutils.MockValidInitAiOp()
	op.MaxFeePerGas, _ = calcNewThresholds(op.MaxFeePerGas, op.MaxPriorityFeePerGas)
	err := ValidatePendingOps(
		op,
		penOps,
	)

	if !errors.Is(err, ErrReplacementOpUnderpriced) {
		t.Fatalf("got %v, want ErrReplacementOpUnderpriced", err)
	}
}

func TestPendingOpsWithOnlyMaxPriorityFeeReplacement(t *testing.T) {
	penOp := testutils.MockValidInitAiOp()
	penOps := []*aiop.AiOperation{penOp}
	op := testutils.MockValidInitAiOp()
	_, op.MaxPriorityFeePerGas = calcNewThresholds(op.MaxFeePerGas, op.MaxPriorityFeePerGas)
	err := ValidatePendingOps(
		op,
		penOps,
	)

	if !errors.Is(err, ErrReplacementOpUnderpriced) {
		t.Fatalf("got %v, want ErrReplacementOpUnderpriced", err)
	}
}

func TestPendingOpsWithOkGasFeeReplacement(t *testing.T) {
	penOp := testutils.MockValidInitAiOp()
	penOps := []*aiop.AiOperation{penOp}
	op := testutils.MockValidInitAiOp()
	op.MaxFeePerGas, op.MaxPriorityFeePerGas = calcNewThresholds(op.MaxFeePerGas, op.MaxPriorityFeePerGas)
	err := ValidatePendingOps(
		op,
		penOps,
	)

	if err != nil {
		t.Fatalf("got err %v, want nil", err)
	}
}
