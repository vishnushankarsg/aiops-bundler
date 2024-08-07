package filter

import (
	"context"
	"math/big"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func filterAiOperationEvent(
	eth *ethclient.Client,
	aiOpHash string,
	aiMiddleware common.Address,
	blkRange uint64,
) (*aimiddleware.AimiddlewareAiOperationEventIterator, error) {
	ep, err := aimiddleware.NewAimiddleware(aiMiddleware, eth)
	if err != nil {
		return nil, err
	}
	bn, err := eth.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	toBlk := big.NewInt(0).SetUint64(bn)
	startBlk := big.NewInt(0)
	subBlkRange := big.NewInt(0).Sub(toBlk, big.NewInt(0).SetUint64(blkRange))
	if subBlkRange.Cmp(startBlk) > 0 {
		startBlk = subBlkRange
	}

	return ep.FilterAiOperationEvent(
		&bind.FilterOpts{Start: startBlk.Uint64()},
		[][32]byte{common.HexToHash(aiOpHash)},
		[]common.Address{},
		[]common.Address{},
	)
}
