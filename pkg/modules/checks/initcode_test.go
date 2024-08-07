package checks

import (
	"testing"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/testutils"
)

func TestInitCodeDNE(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.InitCode = []byte{}
	err := ValidateInitCode(op)

	if err != nil {
		t.Fatalf(`got err %v, want nil`, err)
	}
}

func TestInitCodeContainsBadAddress(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	op.InitCode = []byte("1234")
	err := ValidateInitCode(op)

	if err == nil {
		t.Fatalf("got nil, want err")
	}
}

func TestInitCodeExists(t *testing.T) {
	op := testutils.MockValidInitAiOp()
	err := ValidateInitCode(op)

	if err != nil {
		t.Fatalf(`got err %v, want nil`, err)
	}
}
