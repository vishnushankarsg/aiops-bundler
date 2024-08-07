package stake

import (
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetStakeFunc provides a general interface for retrieving the AiMiddleware stake for a given address.
type GetStakeFunc = func(aiMiddleware, entity common.Address) (*aimiddleware.IDepositManagerDepositInfo, error)

func GetStakeFuncNoop() GetStakeFunc {
	return func(aiMiddleware, entity common.Address) (*aimiddleware.IDepositManagerDepositInfo, error) {
		return &aimiddleware.IDepositManagerDepositInfo{}, nil
	}
}

// GetStakeWithEthClient returns a GetStakeFunc that relies on an eth client to get stake info from the
// AiMiddleware.
func GetStakeWithEthClient(eth *ethclient.Client) GetStakeFunc {
	return func(aiMiddleware, addr common.Address) (*aimiddleware.IDepositManagerDepositInfo, error) {
		if addr == common.HexToAddress("0x") {
			return nil, nil
		}

		ep, err := aimiddleware.NewAimiddleware(aiMiddleware, eth)
		if err != nil {
			return nil, err
		}

		dep, err := ep.GetDepositInfo(nil, addr)
		if err != nil {
			return nil, err
		}

		return &dep, nil
	}
}
