package wire

import (
	"testing"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
)

type testAtxV2Opt func(*ActivationTxV2)

func WithMarriageCertificate(sig *signing.EdSigner, refAtx types.ATXID, atxPublisher types.NodeID) testAtxV2Opt {
	_ = "STUB: not implemented"
	return *new(testAtxV2Opt)
}

func WithMarriageATX(id types.ATXID) testAtxV2Opt {
	_ = "STUB: not implemented"
	return *new(testAtxV2Opt)
}

func WithPublishEpoch(epoch types.EpochID) testAtxV2Opt {
	_ = "STUB: not implemented"
	return *new(testAtxV2Opt)
}

func WithInitial(commitAtx types.ATXID, post PostV1) testAtxV2Opt {
	_ = "STUB: not implemented"
	return *new(testAtxV2Opt)
}

func WithPreviousATXs(atxs ...types.ATXID) testAtxV2Opt {
	_ = "STUB: not implemented"
	return *new(testAtxV2Opt)
}

func WithNIPost(opts ...testNIPostV2Opt) testAtxV2Opt {
	_ = "STUB: not implemented"
	return *new(testAtxV2Opt)
}

type testNIPostV2Opt func(*NIPostV2)

func WithNIPostChallenge(challenge types.Hash32) testNIPostV2Opt {
	_ = "STUB: not implemented"
	return *new(testNIPostV2Opt)
}

func WithNIPostMembershipProof(proof MerkleProofV2) testNIPostV2Opt {
	_ = "STUB: not implemented"
	return *new(testNIPostV2Opt)
}

func WithNIPostSubPost(subPost SubPostV2) testNIPostV2Opt {
	_ = "STUB: not implemented"
	return *new(testNIPostV2Opt)
}

// NewTestActivationTxV2 creates a new ActivationTxV2 with random values.
func NewTestActivationTxV2(tb testing.TB, opts ...testAtxV2Opt) *ActivationTxV2 {
	_ = "STUB: not implemented"
	return nil
}
