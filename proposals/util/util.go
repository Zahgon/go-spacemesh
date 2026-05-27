package util

import (
	"errors"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

var (
	// ErrZeroTotalWeight is returned when zero total epoch weight is used when calculating eligible slots.
	ErrZeroTotalWeight = errors.New("zero total weight not allowed")

	ErrBadBallotData = errors.New("bad ballot data")
)

// NOTE(dshulyak) this was moved into a separate module to disentangle dependencies
// between tortoise and proposals. tests were kept in the proposals/util_test.go
// i will refactor them in a followup

// CalcEligibleLayer calculates the eligible layer from the VRF signature.
func CalcEligibleLayer(epochNumber types.EpochID, layersPerEpoch uint32, vrfSig types.VrfSignature) types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

// GetNumEligibleSlots calculates the number of eligible slots for a smesher in an epoch.
func GetNumEligibleSlots(weight, minWeight, totalWeight uint64, committeeSize, layersPerEpoch uint32) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO: numEligible could overflow uint64 if weight is very large
