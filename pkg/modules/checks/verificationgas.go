package checks

import (
	"fmt"
	"math/big"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/gas"
)

// ValidateVerificationGas checks that the verificationGasLimit is sufficiently low (<= MAX_VERIFICATION_GAS)
// and the preVerificationGas is sufficiently high (enough to pay for the calldata gas cost of serializing
// the AiOperation plus PRE_VERIFICATION_OVERHEAD_GAS).
func ValidateVerificationGas(op *aiop.AiOperation, ov *gas.Overhead, maxVerificationGas *big.Int) error {
	if op.VerificationGasLimit.Cmp(maxVerificationGas) > 0 {
		return fmt.Errorf(
			"verificationGasLimit: exceeds maxVerificationGas of %s",
			maxVerificationGas.String(),
		)
	}

	pvg, err := ov.CalcPreVerificationGas(op)
	if err != nil {
		return err
	}
	if op.PreVerificationGas.Cmp(pvg) < 0 {
		return fmt.Errorf("preVerificationGas: below expected gas of %s", pvg.String())
	}

	return nil
}
