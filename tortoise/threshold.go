package tortoise

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
)

const (
	// By assumption adversarial weight can't be larger than 1/3.
	adversarialWeightFraction = 3
	// Nodes should not be on different sides of the local threshold if they receive different adversarial votes.
	localThresholdFraction = 3
	// For completeness:
	// global threshold is set in a such way so that if adversary
	// cancels their weight (adversarialWeightFraction) - honest nodes should still cross local threshold.
)

func getMedian(heights []uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func computeExpectedWeight(epochs map[types.EpochID]*epochInfo, target, last types.LayerID) weight {
	_ = "STUB: not implemented"
	// all layers after target are voting on the target layer
	// therefore expected weight for the target layer is a sum of all weights
	// within (target, last]
	return *new(weight)
}

// computeGlobalThreshold computes global threshold based on the expected weight.
func computeGlobalThreshold(
	config Config,
	localThreshold weight,
	epochs map[types.EpochID]*epochInfo,
	target, processed, last types.LayerID,
) weight {
	_ = "STUB: not implemented"
	return *new(weight)
}

func computeExpectedWeightInWindow(
	config Config,
	epochs map[types.EpochID]*epochInfo,
	target, processed, last types.LayerID,
) weight {
	_ = "STUB: not implemented"
	return *new(weight)
}

func crossesThreshold(w, t weight) sign { _ = "STUB: not implemented"; return *new(sign) }
