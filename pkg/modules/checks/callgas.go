package checks

import (
	"fmt"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/gas"
)

// ValidateCallGasLimit checks the callGasLimit is at least the cost of a CALL with non-zero value.
func ValidateCallGasLimit(op *aiop.AiOperation, ov *gas.Overhead) error {
	cg := ov.NonZeroValueCall()
	if op.CallGasLimit.Cmp(cg) < 0 {
		return fmt.Errorf("callGasLimit: below expected gas of %s", cg.String())
	}

	return nil
}
