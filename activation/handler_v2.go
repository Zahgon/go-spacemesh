package activation

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/activation/wire"
	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

var errAtxNotV2 = errors.New("ATX is not V2")

type nipostValidatorV2 interface {
	IsVerifyingFullPost() bool
	VRFNonceV2(smesherID types.NodeID, commitment types.ATXID, vrfNonce uint64, numUnits uint32) error
	PostV2(
		ctx context.Context,
		smesherID types.NodeID,
		commitment types.ATXID,
		post *types.Post,
		challenge []byte,
		numUnits uint32,
		opts ...validatorOption,
	) error

	PoetMembership(
		ctx context.Context,
		membership *types.MultiMerkleProof,
		postChallenge types.Hash32,
		poetChallenges [][]byte,
	) (uint64, error)
}

type HandlerV2 struct {
	local           p2p.Peer
	cdb             *datastore.CachedDB
	atxsdata        *atxsdata.Data
	edVerifier      *signing.EdVerifier
	clock           layerClock
	tickSize        uint64
	goldenATXID     types.ATXID
	nipostValidator nipostValidatorV2
	beacon          atxReceiver
	tortoise        system.Tortoise
	logger          *zap.Logger
	fetcher         system.Fetcher
	malPublisher    atxMalfeasancePublisher
}

func (h *HandlerV2) processATX(
	ctx context.Context,
	peer p2p.Peer,
	watx *wire.ActivationTxV2,
	received time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Syntactically validate an ATX.
func (h *HandlerV2) syntacticallyValidate(ctx context.Context, atx *wire.ActivationTxV2) error {
	_ = "STUB: not implemented"
	return nil
}

// Marriage ATX must contain a self-signed certificate.
// It's identified by having ReferenceAtx == EmptyATXID.

// Merged ATX

// Solo chained (non-initial) ATX

// registerHashes registers that the given peer should be asked for
// the hashes of the poet proofs and ATXs.
func (h *HandlerV2) registerHashes(peer p2p.Peer, poetRefs []types.Hash32, atxIDs []types.ATXID) {
	_ = "STUB: not implemented"
	return
}

// fetchReferences makes sure that the referenced poet proof and ATXs are available.
func (h *HandlerV2) fetchReferences(ctx context.Context, poetRefs []types.Hash32, atxIDs []types.ATXID) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect unique dependencies of an ATX.
// Filters out EmptyATXID and the golden ATX.
func (h *HandlerV2) collectAtxDeps(atx *wire.ActivationTxV2) ([]types.Hash32, []types.ATXID) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate the previous ATX for the given PoST and return the effective numunits.
func (h *HandlerV2) validatePreviousAtx(
	id types.NodeID,
	post *wire.SubPostV2,
	prevAtxs []*types.ActivationTx,
) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *HandlerV2) validateCommitmentAtx(golden, commitmentAtxId types.ATXID, publish types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

// validate positioning ATX and return its tick height.
func (h *HandlerV2) validatePositioningAtx(publish types.EpochID, golden, positioning types.ATXID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type marriageInfo struct {
	id        types.NodeID
	signature types.EdSignature
}

// Validate marriages and return married IDs.
// Note: The order of returned IDs is important and must match the order of the marriage certificates.
// The MarriageIndex in PoST proof matches the index in this marriage slice.
func (h *HandlerV2) validateMarriages(atx *wire.ActivationTxV2) ([]marriageInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate marriage ATX and return the full equivocation set.
func (h *HandlerV2) equivocationSet(atx *wire.ActivationTxV2) ([]types.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type idData struct {
	previous      types.ATXID
	previousIndex int
	units         uint32
}

type activationTx struct {
	*wire.ActivationTxV2
	ticks          uint64
	weight         uint64
	effectiveUnits uint32
	ids            map[types.NodeID]idData
	marriages      []marriageInfo
}

type nipostSize struct {
	units uint32
	ticks uint64
}

func (n *nipostSize) addUnits(units uint32) error { _ = "STUB: not implemented"; return nil }

type nipostSizes []*nipostSize

func (n nipostSizes) minTicks() uint64 { _ = "STUB: not implemented"; return 0 }

func (n nipostSizes) sumUp() (units uint32, weight uint64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (h *HandlerV2) verifyIncludedIDsUniqueness(atx *wire.ActivationTxV2) error {
	_ = "STUB: not implemented"
	return nil
}

// Syntactically validate the ATX with its dependencies.
func (h *HandlerV2) syntacticallyValidateDeps(
	ctx context.Context,
	atx *wire.ActivationTxV2,
	peer p2p.Peer,
) (*activationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate previous ATXs

// validate poet membership proofs

// verify PoET memberships in a single go

// validate all NIPoSTs

func (h *HandlerV2) validatePost(
	ctx context.Context,
	nodeID types.NodeID,
	atx *wire.ActivationTxV2,
	peer p2p.Peer,
	commitment types.ATXID,
	challenge types.Hash32,
	post wire.SubPostV2,
	nipostIndex int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// check if post contains at least one valid label

// TODO(mafa): checkpoints need to include all marriage ATXs in full to be able to create malfeasance proofs
// like this one (but also others)
//
// see https://github.com/spacemeshos/go-spacemesh/issues/6435

func (h *HandlerV2) checkMalicious(
	ctx context.Context,
	tx sql.Transaction,
	watx *activationTx,
	peer p2p.Peer,
) (wire.Proof, types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(wire.Proof), *new(types.NodeID), nil
}

func (h *HandlerV2) fetchWireAtx(
	ctx context.Context,
	tx sql.Executor,
	id types.ATXID,
) (*wire.ActivationTxV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HandlerV2) checkDoubleMarry(
	ctx context.Context,
	tx sql.Transaction,
	atx *activationTx,
	peer p2p.Peer,
) (wire.Proof, types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(wire.Proof), *new(types.NodeID), nil
}

func (h *HandlerV2) checkDoubleMerge(
	ctx context.Context,
	tx sql.Transaction,
	atx *activationTx,
	peer p2p.Peer,
) (wire.Proof, types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(wire.Proof), *new(types.NodeID), nil
}

// TODO(mafa): during syntactical validation we should check if a merged ATX is targeting a checkpointed epoch
// merged ATXs need to be checkpointed with their marriage ATXs
// if there is a collision (i.e. the new ATX references the same marriage ATX as a golden ATX) it should be
// considered syntactically invalid
//
// see https://github.com/spacemeshos/go-spacemesh/issues/6434

// TODO(mafa): checkpoints need to include all marriage ATXs in full to be able to create malfeasance proofs
// like this one (but also others)
//
// see https://github.com/spacemeshos/go-spacemesh/issues/6435

func (h *HandlerV2) checkPrevAtx(
	ctx context.Context,
	tx sql.Transaction,
	atx *activationTx,
	peer p2p.Peer,
) (wire.Proof, types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(wire.Proof), *new(types.NodeID), nil
}

// we have at least one v2 ATX (the one we are validating right now) so we only need one
// v1 ATX to create the proof if no other v2 ATXs are found

// prefer creating a proof with 2 ATXs of version 2

// no ATXv2 found, create a proof with an ATXv1

// Store an ATX in the DB.
func (h *HandlerV2) storeAtx(ctx context.Context, atx *types.ActivationTx, watx *activationTx, peer p2p.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

// malfeasance check happens at the end of storing the ATX because storing updates the marriage set
// that is needed for the malfeasance proof

// marriage set of known malicious smesher has changed, force re-gossip of proof

// new malfeasance proof for identity created, publish proof (gossip is decided by publisher)

// cacheAtx caches the atx in the atxsdata cache.
// Returns the cached ATX or nil if the epoch was already evicted.
func (h *HandlerV2) cacheAtx(atx *types.ActivationTx, malicious bool) *atxsdata.ATX {
	_ = "STUB: not implemented"
	return nil
}
