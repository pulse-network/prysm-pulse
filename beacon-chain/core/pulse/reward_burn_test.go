package pulse

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v3/config/params"
	"github.com/prysmaticlabs/prysm/v3/testing/require"
)

func TestApplyBurn(t *testing.T) {
	t.Run("Test burn with various slot times", func(t *testing.T) {
		beforeBurn := uint64(1000000)

		// Default 12 second slots => 25% general burn.
		afterBurn := ApplyBurn(beforeBurn)
		require.Equal(t, uint64(750000), afterBurn)

		// 6 second slots => 50% burn then 25% general burn.
		params.BeaconConfig().SecondsPerSlot = 6
		afterBurn = ApplyBurn(beforeBurn)
		require.Equal(t, uint64(375000), afterBurn)

		// 3 second slots => 75% burn then 25% general burn.
		params.BeaconConfig().SecondsPerSlot = 3
		afterBurn = ApplyBurn(beforeBurn)
		require.Equal(t, uint64(187500), afterBurn)
	})
}
