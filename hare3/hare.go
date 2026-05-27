package hare3

import (
	"context"
	"sync"
	"time"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/hare4"
	"github.com/spacemeshos/go-spacemesh/layerpatrol"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/proposals/store"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

type CommitteeUpgrade struct {
	Layer types.LayerID
	Size  uint16
}

type Config struct {
	Enable           bool          `mapstructure:"enable"`
	EnableLayer      types.LayerID `mapstructure:"enable-layer"`
	DisableLayer     types.LayerID `mapstructure:"disable-layer"`
	Committee        uint16        `mapstructure:"committee"`
	CommitteeUpgrade *CommitteeUpgrade
	Leaders          uint16        `mapstructure:"leaders"`
	IterationsLimit  uint8         `mapstructure:"iterations-limit"`
	PreroundDelay    time.Duration `mapstructure:"preround-delay"`
	RoundDuration    time.Duration `mapstructure:"round-duration"`
	// LogStats if true will log iteration statistics with INFO level at the start of the next iteration.
	// This requires additional computation and should be used for debugging only.
	LogStats     bool   `mapstructure:"log-stats"`
	ProtocolName string `mapstructure:"protocolname"`
}

func (cfg *Config) CommitteeFor(layer types.LayerID) uint16 { _ = "STUB: not implemented"; return 0 }

func (cfg *Config) Validate(zdist time.Duration) error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// roundStart returns expected time for iter/round relative to
// layer start.
func (cfg *Config) roundStart(round IterRound) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func DefaultConfig() Config {
	_ = "STUB: not implemented"

	// NOTE(talm) We aim for a 2^{-40} error probability; if the population at large has a 2/3 honest majority,
	// we need a committee of size ~800 to guarantee this error rate (at least,
	// this is what the Chernoff bound gives you; the actual value is a bit lower,
	// so we can probably get away with a smaller committee). For a committee of size 400,
	// the Chernoff bound gives 2^{-20} probability of a dishonest majority when 1/3 of the population is dishonest.
	return *new(Config)
}

// can be bumped to 3.1 when oracle upgrades

type ConsensusOutput struct {
	Layer     types.LayerID
	Proposals []types.ProposalID
}

type WeakCoinOutput struct {
	Layer types.LayerID
	Coin  bool
}

type Opt func(*Hare)

func WithWallClock(clock clockwork.Clock) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithTracer(tracer Tracer) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithResultsChan overrides the default result channel with a different one.
// This is only needed for the migration period between hare3 and hare4.
func WithResultsChan(c chan hare4.ConsensusOutput) Opt { _ = "STUB: not implemented"; return *new(Opt) }

type nodeClock interface {
	AwaitLayer(types.LayerID) <-chan struct{}
	CurrentLayer() types.LayerID
	LayerToTime(types.LayerID) time.Time
}

func New(
	nodeClock nodeClock,
	pubsub pubsub.PublishSubscriber,
	db sql.StateDatabase,
	atxsdata *atxsdata.Data,
	proposals *store.Store,
	verifier *signing.EdVerifier,
	oracle oracle,
	sync system.SyncStateProvider,
	patrol *layerpatrol.LayerPatrol,
	opts ...Opt,
) *Hare {
	_ = "STUB: not implemented"
	return nil
}

type Hare struct {
	// state
	ctx      context.Context
	cancel   context.CancelFunc
	eg       errgroup.Group
	results  chan hare4.ConsensusOutput
	coins    chan hare4.WeakCoinOutput
	mu       sync.Mutex
	signers  map[string]*signing.EdSigner
	sessions map[types.LayerID]*protocol

	// options
	config    Config
	log       *zap.Logger
	wallClock clockwork.Clock

	// dependencies
	nodeClock nodeClock
	pubsub    pubsub.PublishSubscriber
	db        sql.StateDatabase
	atxsdata  *atxsdata.Data
	proposals *store.Store
	verifier  *signing.EdVerifier
	oracle    *legacyOracle
	sync      system.SyncStateProvider
	patrol    *layerpatrol.LayerPatrol
	tracer    Tracer
}

func (h *Hare) Register(sig *signing.EdSigner) { _ = "STUB: not implemented"; return }

func (h *Hare) Results() <-chan hare4.ConsensusOutput { _ = "STUB: not implemented"; return nil }

func (h *Hare) Coins() <-chan hare4.WeakCoinOutput { _ = "STUB: not implemented"; return nil }

func (h *Hare) Start() { _ = "STUB: not implemented"; return }

func (h *Hare) Running() int { _ = "STUB: not implemented"; return 0 }

func (h *Hare) Handler(ctx context.Context, _ p2p.Peer, buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Hare) onLayer(layer types.LayerID) { _ = "STUB: not implemented"; return }

// signer can't join mid session

// if terminated successfully it will notify block generator
// and it will have to CompleteHare

func (h *Hare) run(session *session) error {
	_ = "STUB: not implemented"
	// oracle may load non-negligible amount of data from disk
	// we do it before preround starts, so that load can have some slack time
	// before it needs to be used in validation
	return nil
}

// initial set is not needed if node is not active in preround

// we are logging stats 1 network delay after new iteration start
// so that we can receive notify messages from previous iteration

func (h *Hare) onOutput(session *session, ir IterRound, out output) error {
	_ = "STUB: not implemented"
	return nil
}

// shallow copy

func (h *Hare) selectProposals(session *session) []types.ProposalID {
	_ = "STUB: not implemented"
	return nil
}

// if atx is not registered for identity we will get sql.ErrNotFound

// double check that a single smesher is not included twice
// theoretically it should never happen as it is covered
// by the malicious check above.

// does not vote for future proposal

func (h *Hare) IsKnown(layer types.LayerID, proposal types.ProposalID) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *Hare) OnProposal(p *types.Proposal) error { _ = "STUB: not implemented"; return nil }

func (h *Hare) Stop() { _ = "STUB: not implemented"; return }

type session struct {
	proto   *protocol
	lid     types.LayerID
	beacon  types.Beacon
	signers []*signing.EdSigner
	vrfs    []*types.HareEligibility
}
