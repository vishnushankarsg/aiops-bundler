package checks

import (
	"testing"

	"github.com/AO-Metaplayer/aiops-bundler/internal/testutils"
)

// TestNilPaymasterAndData calls checks.ValidatePaymasterAndData with no paymaster set. Expects nil.
func TestNilPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.PaymasterAndData = []byte{}
	err := ValidatePaymasterAndData(op, testutils.NonStakedZeroDepositInfo, testutils.MockGetCodeZero)

	if err != nil {
		t.Fatalf("got err %v, want nil", err)
	}
}

// TestBadPaymasterAndData calls checks.ValidatePaymasterAndData with invalid paymaster set. Expects error.
func TestBadPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.PaymasterAndData = []byte("1234")
	err := ValidatePaymasterAndData(op, testutils.NonStakedZeroDepositInfo, testutils.MockGetCodeZero)

	if err == nil {
		t.Fatal("got nil, want err")
	}
}

// TestZeroByteCodePaymasterAndData calls checks.ValidatePaymasterAndData with paymaster contract not
// deployed. Expects error.
func TestZeroByteCodePaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.NonStakedZeroDepositInfo, testutils.MockGetCodeZero)

	if err == nil {
		t.Fatal("got nil, want err")
	}
}

// TestNonStakedZeroDepositPaymasterAndData calls checks.ValidatePaymasterAndData with paymaster that is not
// staked with zero deposit. Expects error.
func TestNonStakedZeroDepositPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.NonStakedZeroDepositInfo, testutils.MockGetCode)

	if err == nil {
		t.Fatal("got nil, want err")
	}
}

// TestZeroDepositPaymasterAndData calls checks.ValidatePaymasterAndData with paymaster that is staked but
// with zero deposit. Expects error.
func TestZeroDepositPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.StakedZeroDepositInfo, testutils.MockGetCode)

	if err == nil {
		t.Fatal("got nil, want err")
	}
}

// TestNotStakedPaymasterAndData calls checks.ValidatePaymasterAndData with paymaster that is not staked and
// has sufficient deposit. Expects nil.
func TestNotStakedPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.NonStakedDepositInfo, testutils.MockGetCode)

	if err != nil {
		t.Fatalf("got %v, want nil", err)
	}
}

// TestPaymasterAndData calls checks.ValidatePaymasterAndData with paymaster that is staked and has sufficient
// deposit. Expects nil.
func TestPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.StakedDepositInfo, testutils.MockGetCode)

	if err != nil {
		t.Fatalf("got err %v, want nil", err)
	}
}
