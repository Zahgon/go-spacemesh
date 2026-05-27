package activation

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/activation/wire"
	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	mwire "github.com/spacemeshos/go-spacemesh/malfeasance/wire"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

type nipostValidatorV1 interface {
	InitialNIPostChallengeV1(challenge *wire.NIPostChallengeV1, atxs atxProvider, goldenATXID types.ATXID) error
	NIPostChallengeV1(challenge *wire.NIPostChallengeV1, previous *types.ActivationTx, nodeID types.NodeID) error
	NIPost(
		ctx context.Context,
		nodeId types.NodeID,
		commitmentAtxId types.ATXID,
		NIPost *types.NIPost,
		expectedChallenge types.Hash32,
		numUnits uint32,
		opts ...validatorOption,
	) (uint64, error)

	NumUnits(cfg *PostConfig, numUnits uint32) error

	IsVerifyingFullPost() bool

	Post(
		ctx context.Context,
		nodeId types.NodeID,
		commitmentAtxId types.ATXID,
		post *types.Post,
		metadata *types.PostMetadata,
		numUnits uint32,
		opts ...validatorOption,
	) error

	VRFNonce(nodeId types.NodeID, commitmentAtxId types.ATXID, vrfNonce, labelsPerUnit uint64, numUnits uint32) error
	PositioningAtx(id types.ATXID, atxs atxProvider, goldenATXID types.ATXID, pubEpoch types.EpochID) error
}

// HandlerV1 processes ATXs version 1.
type HandlerV1 struct {
	local           p2p.Peer
	cdb             *datastore.CachedDB
	atxsdata        *atxsdata.Data
	edVerifier      *signing.EdVerifier
	clock           layerClock
	tickSize        uint64
	goldenATXID     types.ATXID
	nipostValidator nipostValidatorV1
	beacon          atxReceiver
	tortoise        system.Tortoise
	logger          *zap.Logger
	fetcher         system.Fetcher
	malPublisher    legacyMalfeasancePublisher
	malPublisher2   atxMalfeasancePublisher
}

func (h *HandlerV1) syntacticallyValidate(ctx context.Context, atx *wire.ActivationTxV1) error {
	_ = "STUB: not implemented"
	return nil
}

// Use the NIPost's Post metadata, while overriding the challenge to a zero challenge,
// as expected from the initial Post.

// Obtain the commitment ATX ID for the given ATX.
func (h *HandlerV1) commitment(atx *wire.ActivationTxV1) (types.ATXID, error) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

func (h *HandlerV1) syntacticallyValidateDeps(
	ctx context.Context,
	watx *wire.ActivationTxV1,
	received time.Time,
) (*types.ActivationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use the local peer ID as seed for random subset

func (h *HandlerV1) validateNonInitialAtx(
	ctx context.Context,
	atx *wire.ActivationTxV1,
	previous *types.ActivationTx,
	commitment types.ATXID,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// cacheAtx caches the atx in the atxsdata cache.
// Returns the cached ATX or nil if the epoch was already evicted.
func (h *HandlerV1) cacheAtx(atx *types.ActivationTx, malicious bool) *atxsdata.ATX {
	_ = "STUB: not implemented"
	return nil
}

// checkDoublePublish verifies if a node has already published an ATX in the same epoch.
func (h *HandlerV1) checkDoublePublish(
	ctx context.Context,
	tx sql.Executor,
	atx *wire.ActivationTxV1,
	peer peer.ID,
) (*mwire.MalfeasanceProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no ATX previously published for this epoch, or we are handling the same ATX again

// if we land here we tried to publish 2 ATXs in the same epoch
// don't punish ourselves but fail validation and thereby the handling of the incoming ATX

// checkWrongPrevAtx verifies if the previous ATX referenced in the ATX is correct.
func (h *HandlerV1) checkWrongPrevAtx(
	ctx context.Context,
	tx sql.Executor,
	atx *wire.ActivationTxV1,
	peer peer.ID,
) (*mwire.MalfeasanceProof, wire.Proof, error) {
	_ = "STUB: not implemented"
	return nil, *new(wire.Proof), nil
}

// if we land here we tried to publish an ATX with a wrong prevATX

// We retrieved the same ATX, which means this ATX is already in the DB.
// We don't need to look for a different ATX with the same previous ATX
// because if there are already 2 with the same previous ATX, the
// malfeasance proof was already generated.

func (h *HandlerV1) checkMalicious(
	ctx context.Context,
	tx sql.Transaction,
	watx *wire.ActivationTxV1,
	peer peer.ID,
) (*mwire.MalfeasanceProof, wire.Proof, error) {
	_ = "STUB: not implemented"
	return nil, *new(wire.Proof), nil
}

// storeAtx stores an ATX and notifies subscribers of the ATXID.
func (h *HandlerV1) storeAtx(
	ctx context.Context,
	atx *types.ActivationTx,
	watx *wire.ActivationTxV1,
	peer peer.ID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// legacy malfeasance proofs
// new malfeasance proofs

// new legacy malfeasance proof for identity created, publish proof (= persist and gossip)

// new malfeasance proof for identity created, publish proof (= persist and gossip)

func (h *HandlerV1) processATX(
	ctx context.Context,
	peer p2p.Peer,
	watx *wire.ActivationTxV1,
	received time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// registerHashes registers that the given peer should be asked for
// the hashes of the poet proof and ATXs.
func (h *HandlerV1) registerHashes(peer p2p.Peer, poetRef types.Hash32, atxIDs []types.ATXID) {
	_ = "STUB: not implemented"
	return
}

// fetchReferences makes sure that the referenced poet proof and ATXs are available.
func (h *HandlerV1) fetchReferences(ctx context.Context, poetRef types.Hash32, atxIDs []types.ATXID) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect unique dependencies of an ATX.
// Filters out EmptyATXID and the golden ATX.
func collectAtxDeps(goldenAtxId types.ATXID, atx *wire.ActivationTxV1) (types.Hash32, []types.ATXID) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), nil
}

// Obtain the signature of the given ATX.
func atxSignature(ctx context.Context, db sql.Executor, id types.ATXID) (types.EdSignature, error) {
	_ = "STUB: not implemented"
	return *new(types.EdSignature), nil
}

// An empty blob indicates a golden ATX (after a checkpoint-recovery).

// only needed for V1 ATXs
