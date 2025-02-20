package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/AO-Metaplayer/aiops-bundler/pkg/bundler"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/mempool"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/entities"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/signer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Debug exposes methods used for testing the bundler. These should not be made available in production.
type Debug struct {
	eoa          *signer.EOA
	eth          *ethclient.Client
	mempool      *mempool.Mempool
	rep          *entities.Reputation
	bundler      *bundler.Bundler
	chainID      *big.Int
	aimiddleware common.Address
	beneficiary  common.Address
}

func NewDebug(
	eoa *signer.EOA,
	eth *ethclient.Client,
	mempool *mempool.Mempool,
	rep *entities.Reputation,
	bundler *bundler.Bundler,
	chainID *big.Int,
	aimiddleware common.Address,
	beneficiary common.Address,
) *Debug {
	return &Debug{eoa, eth, mempool, rep, bundler, chainID, aimiddleware, beneficiary}
}

// ClearState clears the bundler mempool and reputation data of paymasters/accounts/factories/aggregators.
func (d *Debug) ClearState() (string, error) {
	if err := d.mempool.Clear(); err != nil {
		return "", err
	}

	return "ok", nil
}

// DumpMempool dumps the current AiOperations mempool in order of arrival.
func (d *Debug) DumpMempool(ep string) ([]map[string]any, error) {
	ops, err := d.mempool.Dump(common.HexToAddress(ep))
	if err != nil {
		return []map[string]any{}, err
	}

	res := []map[string]any{}
	for _, op := range ops {
		data, err := op.MarshalJSON()
		if err != nil {
			return []map[string]any{}, err
		}

		item := make(map[string]any)
		if err := json.Unmarshal(data, &item); err != nil {
			return []map[string]any{}, err
		}

		res = append(res, item)
	}

	return res, nil
}

// SendBundleNow forces the bundler to build and execute a bundle from the mempool as handleOps() transaction.
func (d *Debug) SendBundleNow() (string, error) {
	ctx, err := d.bundler.Process(d.aimiddleware)
	if err != nil {
		return "", err
	}
	if ctx == nil {
		return "", nil
	}

	hash, ok := ctx.Data["txn_hash"].(string)
	if !ok {
		return "", errors.New("txn_hash not in ctx Data")
	}
	return hash, nil
}

// SetBundlingMode allows the bundler to be stopped so that an explicit call to debug_bundler_sendBundleNow is
// required to send a bundle.
func (d *Debug) SetBundlingMode(mode string) (string, error) {
	switch mode {
	case "manual":
		d.bundler.Stop()
	case "auto":
		if err := d.bundler.Run(); err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("debug: unrecognized mode %s", mode)
	}

	return "ok", nil
}

// SetReputation allows the bundler to set the reputation of given addresses.
func (d *Debug) SetReputation(entries []any, ep string) (string, error) {
	roArr := []*entities.ReputationOverride{}
	for _, entry := range entries {
		b, err := json.Marshal(entry)
		if err != nil {
			return "", err
		}

		ro := &entities.ReputationOverride{}
		if err := json.Unmarshal(b, ro); err != nil {
			return "", err
		}

		roArr = append(roArr, ro)
	}
	if err := d.rep.Override(roArr); err != nil {
		return "", err
	}

	return "ok", nil
}

// DumpReputation returns the reputation data of all known addresses.
// TODO: Implement
func (d *Debug) DumpReputation(ep string) ([]map[string]any, error) {
	return []map[string]any{}, nil
}
