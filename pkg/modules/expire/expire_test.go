package expire

import (
	"testing"
	"time"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/testutils"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/modules"
	"github.com/ethereum/go-ethereum/common"
)

// TestDropExpired calls (*ExpireHandler).DropExpired and verifies that it marks old AiOperations for
// pending removal.
func TestDropExpired(t *testing.T) {
	exp := New(time.Second * 30)
	op1 := testutils.MockValidInitAiOp()
	op2 := testutils.MockValidInitAiOp()
	op2.CallData = common.Hex2Bytes("0xdead")
	exp.seenAt = map[common.Hash]time.Time{
		op1.GetAiOpHash(testutils.ValidAddress1, common.Big1): time.Now().Add(time.Second * -45),
		op2.GetAiOpHash(testutils.ValidAddress1, common.Big1): time.Now().Add(time.Second * -15),
	}

	ctx := modules.NewBatchHandlerContext(
		[]*aiop.AiOperation{op1, op2},
		testutils.ValidAddress1,
		testutils.ChainID,
		nil,
		nil,
		nil,
	)
	if err := exp.DropExpired()(ctx); err != nil {
		t.Fatalf("got %v, want nil", err)
	} else if len(ctx.Batch) != 1 {
		t.Fatalf("got batch length %d, want 1", len(ctx.Batch))
	} else if len(ctx.PendingRemoval) != 1 {
		t.Fatalf("got pending removal length %d, want 1", len(ctx.Batch))
	} else if !testutils.IsOpsEqual(ctx.Batch[0], op2) {
		t.Fatal("incorrect batch: Dropped legit op")
	} else if !testutils.IsOpsEqual(ctx.PendingRemoval[0].Op, op1) {
		t.Fatal("incorrect pending removal: Didn't drop bad op")
	}

}
