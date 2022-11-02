// Package precompute provides gathering of nicely-structured
// data important to feed into epoch processing, such as attesting
// records and balances, for faster computation.
package precompute

import (
	"context"
	"math/big"

	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/helpers"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/time"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/state"
	"github.com/prysmaticlabs/prysm/v3/config/params"
	"go.opencensus.io/trace"
)

// New gets called at the beginning of process epoch cycle to return
// pre computed instances of validators attesting records and total
// balances attested in an epoch.
func New(ctx context.Context, s state.BeaconState) ([]*Validator, *Balance, error) {
	ctx, span := trace.StartSpan(ctx, "precomputeEpoch.New")
	defer span.End()

	pValidators := make([]*Validator, s.NumValidators())
	pBal := &Balance{
		ActiveCurrentEpoch:         big.NewInt(0),
		ActivePrevEpoch:            big.NewInt(0),
		CurrentEpochAttested:       big.NewInt(0),
		CurrentEpochTargetAttested: big.NewInt(0),
		PrevEpochAttested:          big.NewInt(0),
		PrevEpochHeadAttested:      big.NewInt(0),
		PrevEpochTargetAttested:    big.NewInt(0),
	}

	currentEpoch := time.CurrentEpoch(s)
	prevEpoch := time.PrevEpoch(s)

	if err := s.ReadFromEveryValidator(func(idx int, val state.ReadOnlyValidator) error {
		// Was validator withdrawable or slashed
		withdrawable := prevEpoch+1 >= val.WithdrawableEpoch()
		effectiveBalance := val.EffectiveBalance()
		pVal := &Validator{
			IsSlashed:                    val.Slashed(),
			IsWithdrawableCurrentEpoch:   withdrawable,
			CurrentEpochEffectiveBalance: val.EffectiveBalance(),
		}
		// Was validator active current epoch
		if helpers.IsActiveValidatorUsingTrie(val, currentEpoch) {
			pVal.IsActiveCurrentEpoch = true
			pBal.ActiveCurrentEpoch = new(big.Int).Add(pBal.ActiveCurrentEpoch, new(big.Int).SetUint64(effectiveBalance))
		}
		// Was validator active previous epoch
		if helpers.IsActiveValidatorUsingTrie(val, prevEpoch) {
			pVal.IsActivePrevEpoch = true
			pBal.ActivePrevEpoch = new(big.Int).Add(pBal.ActivePrevEpoch, new(big.Int).SetUint64(effectiveBalance))
		}
		// Set inclusion slot and inclusion distance to be max, they will be compared and replaced
		// with the lower values
		pVal.InclusionSlot = params.BeaconConfig().FarFutureSlot
		pVal.InclusionDistance = params.BeaconConfig().FarFutureSlot

		pValidators[idx] = pVal
		return nil
	}); err != nil {
		return nil, nil, errors.Wrap(err, "failed to initialize precompute")
	}
	return pValidators, pBal, nil
}
