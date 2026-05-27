package events

import (
	"time"

	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type UserEvent struct {
	Event *pb.Event
}

func EmitBeacon(epoch types.EpochID, beacon types.Beacon) { _ = "STUB: not implemented"; return }

func EmitInitStart(nodeID types.NodeID, commitment types.ATXID) { _ = "STUB: not implemented"; return }

func EmitInitFailure(nodeID types.NodeID, commitment types.ATXID, err error) {
	_ = "STUB: not implemented"
	return
}

func EmitInitComplete(nodeID types.NodeID) { _ = "STUB: not implemented"; return }

func EmitRegisteredInPoet(nodeID types.NodeID, url, roundID string) {
	_ = "STUB: not implemented"
	return
}

func EmitProofDownloadedFromPoet(url, roundID string, ticks uint64) {
	_ = "STUB: not implemented"
	return
}

func EmitBestProofSelected(nodeID types.NodeID, url, roundID string, ticks uint64) {
	_ = "STUB: not implemented"
	return
}

// Deprecation. Will be removed soon in favor of EmitWaitingForPoETRegistrationWindow.
func EmitPoetWaitRound(nodeID types.NodeID, current, publish types.EpochID, wait time.Time) {
	_ = "STUB: not implemented"
	return
}

// nolint:staticcheck // SA1019 (deprecated)

func EmitWaitingForPoETRegistrationWindow(nodeID types.NodeID, current, publish types.EpochID, roundEnd time.Time) {
	_ = "STUB: not implemented"
	return
}

// Deprecation. Will be removed soon in favor of EmitWaitingForPoETRegistrationWindow.
func EmitPoetWaitProof(nodeID types.NodeID, publish types.EpochID, wait time.Time) {
	_ = "STUB: not implemented"
	return
}

// nolint:staticcheck // SA1019 (deprecated)

func EmitWaitingForPoETRoundEnd(nodeID types.NodeID, publish types.EpochID, roundEnd time.Time) {
	_ = "STUB: not implemented"
	return
}

func EmitPostServiceStarted() { _ = "STUB: not implemented"; return }

func EmitPostServiceStopped() { _ = "STUB: not implemented"; return }

func EmitPostStart(nodeID types.NodeID, challenge []byte) { _ = "STUB: not implemented"; return }

func EmitPostComplete(nodeID types.NodeID, challenge []byte) { _ = "STUB: not implemented"; return }

func EmitPostFailure(nodeID types.NodeID) { _ = "STUB: not implemented"; return }

func EmitInvalidPostProof(nodeID types.NodeID) { _ = "STUB: not implemented"; return }

func EmitAtxPublished(
	nodeID types.NodeID,
	current, target types.EpochID,
	atxID types.ATXID,
	wait time.Time,
) {
	_ = "STUB: not implemented"
	return
}

func EmitEligibilities(
	nodeID types.NodeID,
	epoch types.EpochID,
	beacon types.Beacon,
	atxID types.ATXID,
	activeSetSize uint32,
	eligibilities map[types.LayerID][]types.VotingEligibility,
) {
	_ = "STUB: not implemented"
	return
}

func castEligibilities(proofs map[types.LayerID][]types.VotingEligibility) []*pb.ProposalEligibility {
	_ = "STUB: not implemented"
	return nil
}

func EmitProposal(nodeID types.NodeID, layer types.LayerID, proposal types.ProposalID) {
	_ = "STUB: not implemented"
	return
}

func EmitOwnMalfeasanceProof(nodeID types.NodeID) { _ = "STUB: not implemented"; return }

func emitUserEvent(help string, failure bool, details pb.IsEventDetails) {
	_ = "STUB: not implemented"
	return
}
