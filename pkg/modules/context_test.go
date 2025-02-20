package modules

import (
	"math/big"
	"testing"

	"github.com/AO-Metaplayer/aiops-bundler/internal/testutils"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aimiddleware"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aimiddleware/stake"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/mempool"
	"github.com/ethereum/go-ethereum/common"
)

func TestNoPendingOps(t *testing.T) {
	db := testutils.DBMock()
	defer db.Close()
	mem, _ := mempool.New(db)
	op := testutils.MockValidInitAiOp()
	op.InitCode = []byte{}
	op.PaymasterAndData = []byte{}

	ctx, err := NewAiOpHandlerContext(
		op,
		testutils.ValidAddress5,
		testutils.ChainID,
		mem,
		stake.GetStakeFuncNoop(),
	)
	if err != nil {
		t.Fatalf("init failed: %v", err)
	} else if pso := ctx.GetPendingSenderOps(); len(pso) != 0 {
		t.Fatalf("pending sender ops: want 0, got %d", len(pso))
	} else if pfo := ctx.GetPendingFactoryOps(); len(pfo) != 0 {
		t.Fatalf("pending factory ops: want 0, got %d", len(pfo))
	} else if ppo := ctx.GetPendingPaymasterOps(); len(ppo) != 0 {
		t.Fatalf("pending paymaster ops: want 0, got %d", len(ppo))
	}
}

func TestGetPendingSenderOps(t *testing.T) {
	db := testutils.DBMock()
	defer db.Close()
	mem, _ := mempool.New(db)
	op := testutils.MockValidInitAiOp()
	op.InitCode = []byte{}
	op.PaymasterAndData = []byte{}

	penOp1 := testutils.MockValidInitAiOp()
	_ = mem.AddOp(testutils.ValidAddress5, penOp1)

	penOp2 := testutils.MockValidInitAiOp()
	penOp2.Nonce = big.NewInt(0).Add(penOp1.Nonce, common.Big1)
	_ = mem.AddOp(testutils.ValidAddress5, penOp2)

	penOp3 := testutils.MockValidInitAiOp()
	penOp3.Nonce = big.NewInt(0).Add(penOp2.Nonce, common.Big1)
	_ = mem.AddOp(testutils.ValidAddress5, penOp3)

	ctx, err := NewAiOpHandlerContext(
		op,
		testutils.ValidAddress5,
		testutils.ChainID,
		mem,
		stake.GetStakeFuncNoop(),
	)
	if err != nil {
		t.Fatalf("init failed: %v", err)
	}

	expectedPenOps := []*aiop.AiOperation{penOp3, penOp2, penOp1}
	penOps := ctx.GetPendingSenderOps()
	if len(penOps) != len(expectedPenOps) {
		t.Fatalf("got length %d, want %d", len(penOps), len(expectedPenOps))
	}

	for i, penOp := range penOps {
		if !testutils.IsOpsEqual(penOp, expectedPenOps[i]) {
			t.Fatalf("ops not equal: %s", testutils.GetOpsDiff(penOp, expectedPenOps[i]))
		}
	}
}

func TestGetPendingFactoryOps(t *testing.T) {
	db := testutils.DBMock()
	defer db.Close()
	mem, _ := mempool.New(db)
	op := testutils.MockValidInitAiOp()
	op.InitCode = testutils.ValidAddress4.Bytes()
	op.PaymasterAndData = []byte{}

	penOp1 := testutils.MockValidInitAiOp()
	penOp1.Sender = testutils.ValidAddress1
	penOp1.InitCode = testutils.ValidAddress4.Bytes()
	_ = mem.AddOp(testutils.ValidAddress5, penOp1)

	penOp2 := testutils.MockValidInitAiOp()
	penOp2.Sender = testutils.ValidAddress2
	penOp2.InitCode = testutils.ValidAddress4.Bytes()
	_ = mem.AddOp(testutils.ValidAddress5, penOp2)

	penOp3 := testutils.MockValidInitAiOp()
	penOp3.Sender = testutils.ValidAddress3
	penOp3.InitCode = testutils.ValidAddress4.Bytes()
	_ = mem.AddOp(testutils.ValidAddress5, penOp3)

	ctx, err := NewAiOpHandlerContext(
		op,
		testutils.ValidAddress5,
		testutils.ChainID,
		mem,
		stake.GetStakeFuncNoop(),
	)
	if err != nil {
		t.Fatalf("init failed: %v", err)
	}

	expectedPenOps := []*aiop.AiOperation{penOp3, penOp2, penOp1}
	penOps := ctx.GetPendingFactoryOps()
	if len(penOps) != len(expectedPenOps) {
		t.Fatalf("got length %d, want %d", len(penOps), len(expectedPenOps))
	}

	for i, penOp := range penOps {
		if !testutils.IsOpsEqual(penOp, expectedPenOps[i]) {
			t.Fatalf("ops not equal: %s", testutils.GetOpsDiff(penOp, expectedPenOps[i]))
		}
	}
}

func TestGetPendingPaymasterOps(t *testing.T) {
	db := testutils.DBMock()
	defer db.Close()
	mem, _ := mempool.New(db)
	op := testutils.MockValidInitAiOp()
	op.InitCode = []byte{}
	op.PaymasterAndData = testutils.ValidAddress4.Bytes()

	penOp1 := testutils.MockValidInitAiOp()
	penOp1.Sender = testutils.ValidAddress1
	penOp1.PaymasterAndData = testutils.ValidAddress4.Bytes()
	_ = mem.AddOp(testutils.ValidAddress5, penOp1)

	penOp2 := testutils.MockValidInitAiOp()
	penOp2.Sender = testutils.ValidAddress2
	penOp2.PaymasterAndData = testutils.ValidAddress4.Bytes()
	_ = mem.AddOp(testutils.ValidAddress5, penOp2)

	penOp3 := testutils.MockValidInitAiOp()
	penOp3.Sender = testutils.ValidAddress3
	penOp3.PaymasterAndData = testutils.ValidAddress4.Bytes()
	_ = mem.AddOp(testutils.ValidAddress5, penOp3)

	ctx, err := NewAiOpHandlerContext(
		op,
		testutils.ValidAddress5,
		testutils.ChainID,
		mem,
		stake.GetStakeFuncNoop(),
	)
	if err != nil {
		t.Fatalf("init failed: %v", err)
	}

	expectedPenOps := []*aiop.AiOperation{penOp3, penOp2, penOp1}
	penOps := ctx.GetPendingPaymasterOps()
	if len(penOps) != len(expectedPenOps) {
		t.Fatalf("got length %d, want %d", len(penOps), len(expectedPenOps))
	}

	for i, penOp := range penOps {
		if !testutils.IsOpsEqual(penOp, expectedPenOps[i]) {
			t.Fatalf("ops not equal: %s", testutils.GetOpsDiff(penOp, expectedPenOps[i]))
		}
	}
}

func TestNilDepositInfo(t *testing.T) {
	db := testutils.DBMock()
	defer db.Close()
	mem, _ := mempool.New(db)
	op := testutils.MockValidInitAiOp()
	op.InitCode = []byte{}
	op.PaymasterAndData = []byte{}

	ctx, err := NewAiOpHandlerContext(
		op,
		testutils.ValidAddress5,
		testutils.ChainID,
		mem,
		func(aiMiddleware, entity common.Address) (*aimiddleware.IDepositManagerDepositInfo, error) {
			if entity == op.Sender {
				return testutils.NonStakedZeroDepositInfo, nil
			}
			return nil, nil
		},
	)
	if err != nil {
		t.Fatalf("init failed: %v", err)
	} else if fd := ctx.GetFactoryDepositInfo(); fd != nil {
		t.Fatalf("factory: want nil, got %v", fd)
	} else if pd := ctx.GetPaymasterDepositInfo(); pd != nil {
		t.Fatalf("paymaster: want nil, got %v", pd)
	}
}

func TestGetSenderDepositInfo(t *testing.T) {
	db := testutils.DBMock()
	defer db.Close()
	mem, _ := mempool.New(db)
	op := testutils.MockValidInitAiOp()
	op.InitCode = []byte{}
	op.PaymasterAndData = []byte{}

	ctx, err := NewAiOpHandlerContext(
		op,
		testutils.ValidAddress5,
		testutils.ChainID,
		mem,
		func(aiMiddleware, entity common.Address) (*aimiddleware.IDepositManagerDepositInfo, error) {
			if entity == op.Sender {
				return testutils.NonStakedZeroDepositInfo, nil
			}
			return nil, nil
		},
	)
	if err != nil {
		t.Fatalf("init failed: %v", err)
	} else if dep := ctx.GetSenderDepositInfo(); dep != testutils.NonStakedZeroDepositInfo {
		t.Fatalf("want %p, got %p", testutils.NonStakedZeroDepositInfo, dep)
	}
}

func TestGetFactoryDepositInfo(t *testing.T) {
	db := testutils.DBMock()
	defer db.Close()
	mem, _ := mempool.New(db)
	op := testutils.MockValidInitAiOp()
	op.InitCode = testutils.ValidAddress1.Bytes()
	op.PaymasterAndData = []byte{}

	ctx, err := NewAiOpHandlerContext(
		op,
		testutils.ValidAddress5,
		testutils.ChainID,
		mem,
		func(aiMiddleware, entity common.Address) (*aimiddleware.IDepositManagerDepositInfo, error) {
			if entity == testutils.ValidAddress1 {
				return testutils.NonStakedZeroDepositInfo, nil
			}
			return nil, nil
		},
	)
	if err != nil {
		t.Fatalf("init failed: %v", err)
	} else if dep := ctx.GetFactoryDepositInfo(); dep != testutils.NonStakedZeroDepositInfo {
		t.Fatalf("want %p, got %p", testutils.NonStakedZeroDepositInfo, dep)
	}
}

func TestGetPaymasterDepositInfo(t *testing.T) {
	db := testutils.DBMock()
	defer db.Close()
	mem, _ := mempool.New(db)
	op := testutils.MockValidInitAiOp()
	op.InitCode = []byte{}
	op.PaymasterAndData = testutils.ValidAddress1.Bytes()

	ctx, err := NewAiOpHandlerContext(
		op,
		testutils.ValidAddress5,
		testutils.ChainID,
		mem,
		func(aiMiddleware, entity common.Address) (*aimiddleware.IDepositManagerDepositInfo, error) {
			if entity == testutils.ValidAddress1 {
				return testutils.NonStakedZeroDepositInfo, nil
			}
			return nil, nil
		},
	)
	if err != nil {
		t.Fatalf("init failed: %v", err)
	} else if dep := ctx.GetPaymasterDepositInfo(); dep != testutils.NonStakedZeroDepositInfo {
		t.Fatalf("want %p, got %p", testutils.NonStakedZeroDepositInfo, dep)
	}
}
