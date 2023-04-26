package pulse_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/prysmaticlabs/go-bitfield"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/core/altair"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/core/helpers"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/core/signing"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/core/time"
	p2pType "github.com/prysmaticlabs/prysm/v4/beacon-chain/p2p/types"
	"github.com/prysmaticlabs/prysm/v4/config/params"
	types "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v4/crypto/bls"
	ethpb "github.com/prysmaticlabs/prysm/v4/proto/prysm/v1alpha1"
	"github.com/prysmaticlabs/prysm/v4/testing/require"
	"github.com/prysmaticlabs/prysm/v4/testing/util"
	"github.com/prysmaticlabs/prysm/v4/time/slots"
)

func TestPulseChainValidatorRewardsMock(t *testing.T) {
	cfg := params.MainnetConfig()
	cfg.BaseRewardFactor = 780000
	cfg.MinDepositAmount = 16000000000000000
	cfg.MaxEffectiveBalance = 32000000000000000
	params.OverrideBeaconConfig(cfg)
	// Vars
	numOfVals := uint64(18312)
	slotsPerEpoch := uint64(32)
	amountOfEpochsToProcess := 5

	// Generate genesis Beaconchain state with required num of vals and deposit amounts
	beaconState, privKeys := util.DeterministicGenesisStateBellatrix(t, numOfVals)
	validators, balance, err := altair.InitializePrecomputeValidators(context.Background(), beaconState)
	require.NoError(t, err)

	// Generate perfect participation object
	// All validators are participating 100% meaning no penalties accounted
	participation := make([]byte, numOfVals)
	inds := make([]uint64, numOfVals)
	for i := 0; i < len(participation); i++ {
		participation[i] = generateParticipation(params.BeaconConfig().TimelySourceFlagIndex, params.BeaconConfig().TimelyTargetFlagIndex, params.BeaconConfig().TimelyHeadFlagIndex)
		inds[i] = uint64(i)
	}

	// Loop for simulating slot progression
	// Sync rewards being processed per each slot
	// Epoch rewards being processed per each epoch
	var slotCounter uint64
	for i := 0; i < int(slotsPerEpoch)*amountOfEpochsToProcess; i++ {
		// Progress 1 slot each itteration
		require.NoError(t, beaconState.SetSlot(types.Slot(i+1)))
		slotCounter++

		// Assign sync committee
		committee, err := altair.NextSyncCommittee(context.Background(), beaconState)
		require.NoError(t, err)
		require.NoError(t, beaconState.SetCurrentSyncCommittee(committee))

		// Define signatures
		syncBits := bitfield.NewBitvector512()
		for i := range syncBits {
			syncBits[i] = 0xff
		}
		indices, err := altair.NextSyncCommitteeIndices(context.Background(), beaconState)
		require.NoError(t, err)
		ps := slots.PrevSlot(beaconState.Slot())
		pbr, err := helpers.BlockRootAtSlot(beaconState, ps)
		require.NoError(t, err)
		sigs := make([]bls.Signature, len(indices))
		for i, indice := range indices {
			b := p2pType.SSZBytes(pbr)
			sb, err := signing.ComputeDomainAndSign(beaconState, time.CurrentEpoch(beaconState), &b, params.BeaconConfig().DomainSyncCommittee, privKeys[indice])
			require.NoError(t, err)
			sig, err := bls.SignatureFromBytes(sb)
			require.NoError(t, err)
			sigs[i] = sig
		}
		aggregatedSig := bls.AggregateSignatures(sigs).Marshal()
		syncAggregate := &ethpb.SyncAggregate{
			SyncCommitteeBits:      syncBits,
			SyncCommitteeSignature: aggregatedSig,
		}

		// Process Sync aggregation and process rewards
		beaconState, _, err = altair.ProcessSyncAggregate(context.Background(), beaconState, syncAggregate)
		require.NoError(t, err)

		// If epoch passed - process participation and rewards
		if slotCounter == 32 {
			slotCounter = 0
			require.NoError(t, beaconState.SetCurrentParticipationBits(participation))
			require.NoError(t, beaconState.SetPreviousParticipationBits(participation))
			validators, balance, err = altair.ProcessEpochParticipation(context.Background(), beaconState, balance, validators)
			require.NoError(t, err)
			beaconState, err = altair.ProcessRewardsAndPenaltiesPrecompute(beaconState, balance, validators)
			require.NoError(t, err)
		}
	}

	// Count total rewards after simulation
	var totalRewards *big.Int = big.NewInt(0)
	bal := beaconState.Balances()
	for i := 0; i < len(bal); i++ {
		totalRewards.Add(totalRewards, new(big.Int).Sub(new(big.Int).SetUint64(bal[i]), new(big.Int).SetUint64(cfg.MaxEffectiveBalance)))
	}
	t.Log("totalRewards:", totalRewards)
}

// Helper function to generate participation
func generateParticipation(flags ...uint8) byte {
	b := byte(0)
	var err error
	for _, flag := range flags {
		b, err = altair.AddValidatorFlag(b, flag)
		if err != nil {
			return 0
		}
	}
	return b
}
