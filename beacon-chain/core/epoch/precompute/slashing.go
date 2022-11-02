package precompute

import (
	"math/big"

	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/helpers"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/time"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/state"
	"github.com/prysmaticlabs/prysm/v3/config/params"
	types "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
)

// ProcessSlashingsPrecompute processes the slashed validators during epoch processing.
// This is an optimized version by passing in precomputed total epoch balances.
func ProcessSlashingsPrecompute(s state.BeaconState, pBal *Balance) error {
	currentEpoch := time.CurrentEpoch(s)
	exitLength := params.BeaconConfig().EpochsPerSlashingsVector

	// Compute the sum of state slashings
	slashings := s.Slashings()
	totalSlashing := uint64(0)
	for _, slashing := range slashings {
		totalSlashing += slashing
	}

	var minSlashing *big.Int = big.NewInt(0)
	totalSlashingTimesPropSlashMul := new(big.Int).SetUint64(totalSlashing * params.BeaconConfig().ProportionalSlashingMultiplier)

	if totalSlashingTimesPropSlashMul.Cmp(pBal.ActiveCurrentEpoch) == -1 {
		minSlashing = totalSlashingTimesPropSlashMul
	} else {
		minSlashing = pBal.ActiveCurrentEpoch
	}
	epochToWithdraw := currentEpoch + exitLength/2

	var hasSlashing bool
	// Iterate through validator list in state, stop until a validator satisfies slashing condition of current epoch.
	err := s.ReadFromEveryValidator(func(idx int, val state.ReadOnlyValidator) error {
		correctEpoch := epochToWithdraw == val.WithdrawableEpoch()
		if val.Slashed() && correctEpoch {
			hasSlashing = true
		}
		return nil
	})
	if err != nil {
		return err
	}
	// Exit early if there's no meaningful slashing to process.
	if !hasSlashing {
		return nil
	}

	increment := new(big.Int).SetUint64(params.BeaconConfig().EffectiveBalanceIncrement)
	validatorFunc := func(idx int, val *ethpb.Validator) (bool, *ethpb.Validator, error) {
		correctEpoch := epochToWithdraw == val.WithdrawableEpoch
		if val.Slashed && correctEpoch {
			valEffectiveBal := new(big.Int).SetUint64(val.EffectiveBalance)
			penaltyNumerator := new(big.Int).Mul(new(big.Int).Div(valEffectiveBal, increment), minSlashing)
			penalty := new(big.Int).Mul(new(big.Int).Div(penaltyNumerator, pBal.ActiveCurrentEpoch), increment).Uint64()
			if err := helpers.DecreaseBalance(s, types.ValidatorIndex(idx), penalty); err != nil {
				return false, val, err
			}
			return true, val, nil
		}
		return false, val, nil
	}

	return s.ApplyToEveryValidator(validatorFunc)
}
