package proposals

import (
	"context"
	"errors"
	"fmt"
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
	"github.com/spacemeshos/go-spacemesh/tortoise"
)

var (
	errMalformedData         = fmt.Errorf("%w: malformed data", pubsub.ErrValidationReject)
	errWrongHash             = fmt.Errorf("%w: incorrect hash", pubsub.ErrValidationReject)
	errInitialize            = fmt.Errorf("%w: failed to initialize", pubsub.ErrValidationReject)
	errInvalidATXID          = fmt.Errorf("%w: ballot has invalid ATXID", pubsub.ErrValidationReject)
	errDoubleVoting          = fmt.Errorf("%w: ballot doubly-voted in same layer", pubsub.ErrValidationReject)
	errConflictingExceptions = fmt.Errorf("%w: conflicting exceptions", pubsub.ErrValidationReject)
	errExceptionsOverflow    = fmt.Errorf("%w: too many exceptions", pubsub.ErrValidationReject)
	errDuplicateTX           = fmt.Errorf("%w: duplicate TxID in proposal", pubsub.ErrValidationReject)
	errKnownProposal         = errors.New("known proposal")
	errKnownBallot           = errors.New("known ballot")
	errMaliciousBallot       = errors.New("malicious ballot")
)

// Handler processes Proposal from gossip and, if deems it valid, propagates it to peers.
type Handler struct {
	logger *zap.Logger
	cfg    Config

	db                sql.StateDatabase
	atxsdata          *atxsdata.Data
	activeSets        *lru.Cache[types.Hash32, uint64]
	edVerifier        *signing.EdVerifier
	publisher         pubsub.Publisher
	fetcher           system.Fetcher
	mesh              meshProvider
	validator         eligibilityValidator
	tortoise          tortoiseProvider
	weightCalcLock    sync.Mutex
	pendingWeightCalc map[types.Hash32][]chan uint64
	clock             layerClock

	proposals proposalsConsumer
}

// Config defines configuration for the handler.
type Config struct {
	LayerSize              uint32
	LayersPerEpoch         uint32
	GoldenATXID            types.ATXID
	MaxExceptions          int
	Hdist                  uint32
	MinimalActiveSetWeight []types.EpochMinimalActiveWeight
}

// defaultConfig for BlockHandler.
func defaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// Opt for configuring Handler.
type Opt func(h *Handler)

// withValidator defines eligibility Validator.
func withValidator(v eligibilityValidator) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithLogger defines logger for Handler.
func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithConfig defines protocol parameters.
func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// NewHandler creates new Handler.
func NewHandler(
	db sql.StateDatabase,
	atxsdata *atxsdata.Data,
	proposals proposalsConsumer,
	edVerifier *signing.EdVerifier,
	p pubsub.Publisher,
	f system.Fetcher,
	bc system.BeaconCollector,
	m meshProvider,
	tortoise tortoiseProvider,
	verifier vrfVerifier,
	clock layerClock,
	opts ...Opt,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

// HandleSyncedBallot handles Ballot data from sync.
func (h *Handler) HandleSyncedBallot(ctx context.Context, expHash types.Hash32, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// set the ballot and smesher ID when received

func (h *Handler) HandleActiveSet(ctx context.Context, id types.Hash32, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) handleSet(ctx context.Context, id types.Hash32, set types.EpochActiveSet) error {
	_ = "STUB: not implemented"
	return nil
}

// active set is invalid unless all activations that it references are from the correct epoch

// collectHashes gathers all hashes in a proposal or ballot.
func collectHashes(a any) []types.Hash32 { _ = "STUB: not implemented"; return nil }

// HandleSyncedProposal handles Proposal data from sync.
func (h *Handler) HandleSyncedProposal(ctx context.Context, expHash types.Hash32, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleProposal is the gossip receiver for Proposal.
func (h *Handler) HandleProposal(ctx context.Context, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleProposal is the gossip receiver for Proposal.
func (h *Handler) handleProposal(ctx context.Context, expHash types.Hash32, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// set the proposal ID when received

// FIXME: how to handle proposals from malicious identity?

// broadcast malfeasance proof last as the verification of the proof will take place
// in the same goroutine

func (h *Handler) setProposalBeacon(p *types.Proposal) error { _ = "STUB: not implemented"; return nil }

func (h *Handler) processBallot(ctx context.Context, b *types.Ballot) (*wire.MalfeasanceProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the active set is not needed anymore

func (h *Handler) checkBallotSyntacticValidity(ctx context.Context, b *types.Ballot) (*tortoise.DecodedBallot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ballot can be decoded only if all dependencies (ballots, atxs) were downloaded
// and added to the tortoise.

// note that computed opinion has to match signed opinion, otherwise it is unknown
// if attached votes struct was modified
//
// TODO this check can work only on the list with decoded votes, otherwise
// otherwise it validates only diff, which is easy to bypass

func (h *Handler) getActiveSetWeight(ctx context.Context, id types.Hash32) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// The calculation is running or the activeset is being fetched,
// subscribe.
// Avoid any blocking on the channel by making it buffered, also so that
// we don't have to wait on it in case the context is canceled

// need to wait for the calculation which is already running to finish

// Channel closed, fetch / calculation failed.
// The actual error will be logged by the initiator of the
// initial fetch / calculation, let's not make an
// impression it happened multiple times and use a simpler
// message

// mark calculation as running

// this is guaranteed not to block b/c each channel is buffered

// totalWeight will be sent to the subscribers

func (h *Handler) checkBallotDataIntegrity(ctx context.Context, b *types.Ballot) (uint64, error) {
	_ = "STUB: not implemented"
	// nolint:nestif
	return 0, nil
}

// this is the smesher's first Ballot in this epoch, should contain EpochData

// download activesets in the previous epoch too

func (h *Handler) checkVotesConsistency(ctx context.Context, b *types.Ballot) error {
	_ = "STUB: not implemented"
	return nil
}

// a ballot should not vote for multiple blocks in the same layer within hdist,
// since hare only output a single block each layer and miner should vote according
// to the hare output within hdist of the current layer when producing a ballot.

// already voted for a block in this layer

// a ballot should not vote support and against on the same block.

// a ballot should not abstain on a layer that it voted for/against on block in that layer.

func (h *Handler) checkBallotDataAvailability(ctx context.Context, b *types.Ballot) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) checkTransactions(ctx context.Context, p *types.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

func reportProposalMetrics(p *types.Proposal) { _ = "STUB: not implemented"; return }

func reportVotesMetrics(b *types.Ballot) { _ = "STUB: not implemented"; return }
