package stateutil

import (
	"math/big"

	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v4/config/params"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	ethpb "github.com/prysmaticlabs/prysm/v4/proto/prysm/v1alpha1"
)

// UnrealizedCheckpointBalances returns the total current active balance, the
// total previous epoch correctly attested for target balance, and the total
// current epoch correctly attested for target balance. It takes the current and
// previous epoch participation bits as parameters so implicitly only works for
// beacon states post-Altair.
func UnrealizedCheckpointBalances(cp, pp []byte, validators []*ethpb.Validator, currentEpoch primitives.Epoch) (*big.Int, *big.Int, *big.Int, error) {
	targetIdx := params.BeaconConfig().TimelyTargetFlagIndex
	activeBalance := big.NewInt(0)
	currentTarget := big.NewInt(0)
	prevTarget := big.NewInt(0)
	if len(cp) < len(validators) || len(pp) < len(validators) {
		return big.NewInt(0), big.NewInt(0), big.NewInt(0), errors.New("participation does not match validator set")
	}

	for i, v := range validators {
		active := v.ActivationEpoch <= currentEpoch && currentEpoch < v.ExitEpoch
		if active && !v.Slashed {
			effectiveBalanceBig := new(big.Int).SetUint64(v.EffectiveBalance)
			activeBalance.Add(activeBalance, effectiveBalanceBig)
			if ((cp[i] >> targetIdx) & 1) == 1 {
				effectiveBalanceBig := new(big.Int).SetUint64(v.EffectiveBalance)
				currentTarget.Add(currentTarget, effectiveBalanceBig)
			}
			if ((pp[i] >> targetIdx) & 1) == 1 {
				effectiveBalanceBig := new(big.Int).SetUint64(v.EffectiveBalance)
				prevTarget.Add(prevTarget, effectiveBalanceBig)
			}
		}
	}
	return activeBalance, prevTarget, currentTarget, nil
}
