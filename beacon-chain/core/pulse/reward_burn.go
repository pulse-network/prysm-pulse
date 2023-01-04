// Package pulse implements the PulseChain fork
package pulse

import (
	"github.com/prysmaticlabs/prysm/v4/config/params"
	"github.com/sirupsen/logrus"
)

// Applies the PulseChain burn to a pending validator reward.
func ApplyBurn(baseReward uint64) uint64 {
	secondsPerSlot := params.BeaconConfig().SecondsPerSlot

	// First we compensate for the increased block frequency.
	afterBurn := baseReward * secondsPerSlot / 12

	// Then we burn an additional 25%.
	afterBurn = afterBurn * 3 / 4

	logrus.WithFields(logrus.Fields{
		"baseReward": baseReward,
		"afterBurn":  afterBurn,
	}).Debug("Applied PulseChain Burn 🔥")
	return afterBurn
}
