package proposals

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/proposals/util"
)

var (
	CalcEligibleLayer   = util.CalcEligibleLayer
	GetNumEligibleSlots = util.GetNumEligibleSlots
)

func MustGetNumEligibleSlots(
	weight, minWeight, totalWeight uint64,
	committeeSize, layersPerEpoch uint32,
) uint32 {
	_ = "STUB: not implemented"
	return 0
}

//go:generate scalegen -types VrfMessage

// VrfMessage is a verification message. It is the payload for the signature in `VotingEligibility`.
type VrfMessage struct {
	Type    types.EligibilityType // always types.EligibilityVoting
	Beacon  types.Beacon
	Epoch   types.EpochID
	Nonce   types.VRFPostIndex
	Counter uint32
}

// MustSerializeVRFMessage serializes a message for generating/verifying a VRF signature.
func MustSerializeVRFMessage(
	beacon types.Beacon,
	epoch types.EpochID,
	nonce types.VRFPostIndex,
	counter uint32,
) []byte {
	_ = "STUB: not implemented"
	return nil
}
