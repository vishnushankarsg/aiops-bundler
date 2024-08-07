package mempool

import (
	"sync"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wangjia184/sortedset"
)

type set struct {
	all      *sortedset.SortedSet
	entities map[common.Address]*sortedset.SortedSet
}

func (s *set) getEntitiesSortedSet(entity common.Address) *sortedset.SortedSet {
	if _, ok := s.entities[entity]; !ok {
		s.entities[entity] = sortedset.New()
	}

	return s.entities[entity]
}

type aiOpQueues struct {
	setsByAiMiddleware sync.Map
}

func (q *aiOpQueues) getAiMiddlewareSet(aiMiddleware common.Address) *set {
	val, ok := q.setsByAiMiddleware.Load(aiMiddleware)
	if !ok {
		val = &set{
			all:      sortedset.New(),
			entities: make(map[common.Address]*sortedset.SortedSet),
		}
		q.setsByAiMiddleware.Store(aiMiddleware, val)
	}

	return val.(*set)
}

func (q *aiOpQueues) AddOp(aiMiddleware common.Address, op *aiop.AiOperation) {
	eps := q.getAiMiddlewareSet(aiMiddleware)
	key := string(getUniqueKey(aiMiddleware, op.Sender, op.Nonce))

	eps.all.AddOrUpdate(key, sortedset.SCORE(eps.all.GetCount()), op)
	eps.getEntitiesSortedSet(op.Sender).AddOrUpdate(key, sortedset.SCORE(op.Nonce.Int64()), op)
	if factory := op.GetFactory(); factory != common.HexToAddress("0x") {
		fss := eps.getEntitiesSortedSet(factory)
		fss.AddOrUpdate(key, sortedset.SCORE(fss.GetCount()), op)
	}
	if paymaster := op.GetPaymaster(); paymaster != common.HexToAddress("0x") {
		pss := eps.getEntitiesSortedSet(paymaster)
		pss.AddOrUpdate(key, sortedset.SCORE(pss.GetCount()), op)
	}
}

func (q *aiOpQueues) GetOps(aiMiddleware common.Address, entity common.Address) []*aiop.AiOperation {
	eps := q.getAiMiddlewareSet(aiMiddleware)
	ess := eps.getEntitiesSortedSet(entity)
	nodes := ess.GetByRankRange(-1, -ess.GetCount(), false)
	batch := []*aiop.AiOperation{}
	for _, n := range nodes {
		batch = append(batch, n.Value.(*aiop.AiOperation))
	}

	return batch
}

func (q *aiOpQueues) All(aiMiddleware common.Address) []*aiop.AiOperation {
	eps := q.getAiMiddlewareSet(aiMiddleware)
	nodes := eps.all.GetByRankRange(1, -1, false)
	batch := []*aiop.AiOperation{}
	for _, n := range nodes {
		batch = append(batch, n.Value.(*aiop.AiOperation))
	}

	return batch
}

func (q *aiOpQueues) RemoveOps(aiMiddleware common.Address, ops ...*aiop.AiOperation) {
	eps := q.getAiMiddlewareSet(aiMiddleware)
	for _, op := range ops {
		key := string(getUniqueKey(aiMiddleware, op.Sender, op.Nonce))
		eps.all.Remove(key)
		eps.getEntitiesSortedSet(op.Sender).Remove(key)
		eps.getEntitiesSortedSet(op.GetFactory()).Remove(key)
		eps.getEntitiesSortedSet(op.GetPaymaster()).Remove(key)
	}
}

func newAiOpQueue() *aiOpQueues {
	return &aiOpQueues{}
}
