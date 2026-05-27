// Package miner is responsible for creating valid blocks that contain valid activation transactions and transactions
package miner

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/semaphore"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
	"github.com/spacemeshos/go-spacemesh/tortoise"
)

var errAtxNotAvailable = errors.New("atx not available")

//go:generate mockgen -typed -package=mocks -destination=./mocks/mocks.go -source=./proposal_builder.go

type conservativeState interface {
	SelectProposalTXs(types.LayerID, int) []types.TransactionID
}

type votesEncoder interface {
	LatestComplete() types.LayerID
	TallyVotes(types.LayerID)
	EncodeVotes(context.Context, ...tortoise.EncodeVotesOpts) (*types.Opinion, error)
}

type layerClock interface {
	AwaitLayer(layerID types.LayerID) <-chan struct{}
	CurrentLayer() types.LayerID
	LayerToTime(types.LayerID) time.Time
}

type atxsData interface {
	Get(types.EpochID, types.ATXID) *atxsdata.ATX
	GetByEpochAndNodeID(types.EpochID, types.NodeID) (types.ATXID, *atxsdata.ATX)
}

// ProposalBuilder builds Proposals for a miner.
type ProposalBuilder struct {
	logger *zap.Logger
	cfg    config

	db        sql.StateDatabase
	localdb   sql.Executor
	atxsdata  atxsData
	clock     layerClock
	publisher pubsub.Publisher
	conState  conservativeState
	tortoise  votesEncoder
	syncer    system.SyncStateProvider
	activeGen *activeSetGenerator

	signers struct {
		mu      sync.Mutex
		signers map[types.NodeID]*signerSession
	}
	shared sharedSession

	limiter *semaphore.Weighted
}

type signerSession struct {
	signer  *signing.EdSigner
	log     *zap.Logger
	session session
	latency latencyTracker
}

// shared data for all signers in the epoch.
type sharedSession struct {
	epoch  types.EpochID
	beacon types.Beacon
	active struct {
		id     types.Hash32
		set    types.ATXIDList
		weight uint64
	}
}

// session per every signing key for the whole epoch.
type session struct {
	epoch         types.EpochID
	atx           types.ATXID
	atxWeight     uint64
	ref           types.BallotID
	beacon        types.Beacon
	prev          types.LayerID
	nonce         types.VRFPostIndex
	eligibilities struct {
		proofs map[types.LayerID][]types.VotingEligibility
		slots  uint32
	}
}

func (s *session) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Sort the layer map to log the layer data in order

// config defines configuration for the ProposalBuilder.
type config struct {
	layerSize          uint32
	layersPerEpoch     uint32
	hdist              uint32
	networkDelay       time.Duration
	workersLimit       int
	minActiveSetWeight []types.EpochMinimalActiveWeight
	// used to determine whether a node has enough information on the active set this epoch
	goodAtxPercent int
	activeSet      ActiveSetPreparation
}

// ActiveSetPreparation is a configuration to enable computation of activeset in advance.
type ActiveSetPreparation struct {
	// Window describes how much in advance the active set should be prepared.
	Window time.Duration `mapstructure:"window"`
	// RetryInterval describes how often the active set is retried.
	RetryInterval time.Duration `mapstructure:"retry-interval"`
	// Tries describes how many times the active set is retried.
	Tries int `mapstructure:"tries"`
}

func DefaultActiveSetPreparation() ActiveSetPreparation {
	_ = "STUB: not implemented"
	return *new(ActiveSetPreparation)
}

func (c *config) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Opt for configuring ProposalBuilder.
type Opt func(h *ProposalBuilder)

// WithLayerSize defines the average number of proposal per layer.
func WithLayerSize(size uint32) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithWorkersLimit configures parallelization factor for builder operation when working with
// more than one signer.
func WithWorkersLimit(limit int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithLayerPerEpoch defines the number of layers per epoch.
func WithLayerPerEpoch(layers uint32) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithMinimalActiveSetWeight(weight []types.EpochMinimalActiveWeight) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithLogger defines the logger.
func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithHdist(dist uint32) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithNetworkDelay(delay time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithMinGoodAtxPercent(percent int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithSigners guarantees that builder will start execution with provided list of signers.
// Should be after logging.
func WithSigners(signers ...*signing.EdSigner) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithActivesetPreparation overwrites configuration for activeset preparation.
func WithActivesetPreparation(prep ActiveSetPreparation) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// New creates a struct of block builder type.
func New(
	clock layerClock,
	db sql.StateDatabase,
	localdb sql.Executor,
	atxsdata atxsData,
	publisher pubsub.Publisher,
	trtl votesEncoder,
	syncer system.SyncStateProvider,
	conState conservativeState,
	opts ...Opt,
) *ProposalBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (pb *ProposalBuilder) Register(sig *signing.EdSigner) { _ = "STUB: not implemented"; return }

// Start the loop that listens to layers and build proposals.
func (pb *ProposalBuilder) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (pb *ProposalBuilder) runPrepare(ctx context.Context, current types.LayerID) {
	_ = "STUB: not implemented"
	return
}

// Only output the mesh hash in the proposal when the following conditions are met:
// - tortoise has verified every layer i < N-hdist.
// - the node has hare output for every layer i such that N-hdist <= i <= N.
// This is done such that when the node is generating the block based on hare output,
// it can do optimistic filtering if the majority of the proposals agreed on the mesh hash.
func (pb *ProposalBuilder) decideMeshHash(ctx context.Context, current types.LayerID) types.Hash32 {
	_ = "STUB: not implemented"
	return *new(types.Hash32)
}

func (pb *ProposalBuilder) UpdateActiveSet(target types.EpochID, set []types.ATXID) {
	_ = "STUB: not implemented"
	return
}

func (pb *ProposalBuilder) initSharedData(ctx context.Context, current types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// Ideally we only persist the active set when we are actually eligible with at least one identity in at least one
// layer, but since at the moment we use a bootstrapped activeset, `activesets.Has` will always return
// true anyways.
//
// Additionally all activesets that are older than 2 epochs are deleted at the beginning of an epoch anyway, but
// maybe we should revisit this when activesets are no longer bootstrapped.

func (pb *ProposalBuilder) initSignerData(ss *signerSession, lid types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pb *ProposalBuilder) build(ctx context.Context, lid types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// don't accept registration in the middle of computing proposals

// TODO(dshulyak) get rid from the EncodeVotesWithCurrent option in a followup
// there are some dependencies in the tests

// Two stage pipeline, with the stages running in parallel.
// 1. Initializes signers. Runs limited number of goroutines because the initialization is CPU and DB bound.
// 2. Collects eligible signers' sessions from the stage 1 and creates and publishes proposals.

// Used to pass eligible singers from stage 1 → 2.
// Buffered with capacity for all signers so that writes don't block.

// Stage 1
// We use a semaphore instead of eg.SetLimit so that the stage 2 starts immediately after
// scheduling all signers in the stage 1. Otherwise, stage 2 would wait for all stage 1
// goroutines to at least start, which is not what we want. We want to start stage 2 as soon as possible.

// won't block

// Stage 2

func createProposal(
	session *session,
	beacon types.Beacon,
	activeset types.ATXIDList,
	signer *signing.EdSigner,
	lid types.LayerID,
	txs []types.TransactionID,
	opinion *types.Opinion,
	eligibility []types.VotingEligibility,
	meshHash types.Hash32,
) *types.Proposal {
	_ = "STUB: not implemented"
	return nil
}

// calcEligibilityProofs calculates the eligibility proofs of proposals for the miner in the given epoch
// and returns the proofs along with the epoch's active set.
func calcEligibilityProofs(
	signer *signing.VRFSigner,
	epoch types.EpochID,
	beacon types.Beacon,
	nonce types.VRFPostIndex,
	slots uint32,
	layersPerEpoch uint32,
) map[types.LayerID][]types.VotingEligibility {
	_ = "STUB: not implemented"
	return nil
}
