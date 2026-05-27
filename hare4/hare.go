package hare4

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/layerpatrol"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/p2p/server"
	"github.com/spacemeshos/go-spacemesh/proposals/store"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

const (
	PROTOCOL_NAME     = "hare4/full_exchange"
	MAX_EXCHANGE_SIZE = 1_000_000 // protect against a malicious allocation of too much space.
)

var (
	errNoLayerProposals     = errors.New("no proposals for layer")
	errCannotMatchProposals = errors.New("cannot match proposals to compacted form")
	errResponseTooBig       = errors.New("response too big")
	errCannotFindProposal   = errors.New("cannot find proposal")
	errNoEligibilityProofs  = errors.New("no eligibility proofs")
	fetchFullTimeout        = 5 * time.Second
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

type ConsensusOutput struct {
	Layer     types.LayerID
	Proposals []types.ProposalID
}

type WeakCoinOutput struct {
	Layer types.LayerID
	Coin  bool
}

type Opt func(*Hare)

func WithServer(s streamRequester) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithWallClock(clock clockwork.Clock) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithTracer(tracer Tracer) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithResultsChan overrides the default result channel with a different one.
// This is only needed for the migration period between hare3 and hare4.
func WithResultsChan(c chan ConsensusOutput) Opt { _ = "STUB: not implemented"; return *new(Opt) }

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
	verifier verifier,
	oracle oracle,
	sync system.SyncStateProvider,
	patrol *layerpatrol.LayerPatrol,
	host server.Host,
	opts ...Opt,
) *Hare {
	_ = "STUB: not implemented"
	return nil
}

type Hare struct {
	// state
	ctx          context.Context
	cancel       context.CancelFunc
	eg           errgroup.Group
	results      chan ConsensusOutput
	coins        chan WeakCoinOutput
	mu           sync.Mutex
	signers      map[string]*signing.EdSigner
	sessions     map[types.LayerID]*protocol
	messageCache map[types.Hash32]Message

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
	verifier  verifier
	oracle    *legacyOracle
	sync      system.SyncStateProvider
	patrol    *layerpatrol.LayerPatrol
	p2p       streamRequester
	tracer    Tracer
}

func (h *Hare) Register(sig *signing.EdSigner) { _ = "STUB: not implemented"; return }

func (h *Hare) Results() <-chan ConsensusOutput { _ = "STUB: not implemented"; return nil }

func (h *Hare) Coins() <-chan WeakCoinOutput { _ = "STUB: not implemented"; return nil }

func (h *Hare) Start() { _ = "STUB: not implemented"; return }

func (h *Hare) Running() int { _ = "STUB: not implemented"; return 0 }

// fetchFull will fetch the full list of proposal IDs from the provided peer.
func (h *Hare) fetchFull(ctx context.Context, peer p2p.Peer, msgId types.Hash32) (
	[]types.ProposalID, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Hare) handleProposalsStream(ctx context.Context, _ p2p.Peer, msg []byte, s io.ReadWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// reconstructProposals tries to reconstruct the full list of proposals from a peer based on a delivered
// set of compact IDs.
func (h *Hare) reconstructProposals(ctx context.Context, peer p2p.Peer, msgId types.Hash32, msg *Message) error {
	_ = "STUB: not implemented"
	return nil
}

// item is both not taken and equals to looked up ID

// try to see if we can match it to the proposals we have
// if we do, add the proposal ID to the list of hashes in the proposals on the message

// if we can't find it, we can already assume that we cannot match the whole
// set and just fail fast

// this will force the calling context to do a fetchFull

// sort the found proposals and unset the compact proposals
// field before trying to check the signature
// since it would add unnecessary data to the hasher

func (h *Hare) Handler(ctx context.Context, peer p2p.Peer, buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// this will mutate the message to conform to the (hopefully)
// original sent message for signature validation to occur

// we might have a bad signature because of a local hash collision
// of a proposal that has the same short hash that the node sent us.
// in this case we try to ask for a full exchange of all full proposal
// ids and try to validate again

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

// OnProposal is a hook which gets called when we get a proposal.
func (h *Hare) OnProposal(p *types.Proposal) error { _ = "STUB: not implemented"; return nil }

// cleanMessageCache cleans old cached preround messages
// once the layers become irrelevant.
func (h *Hare) cleanMessageCache(l types.LayerID) { _ = "STUB: not implemented"; return }

// mark key for deletion

func (h *Hare) Stop() { _ = "STUB: not implemented"; return }

type session struct {
	proto   *protocol
	lid     types.LayerID
	beacon  types.Beacon
	signers []*signing.EdSigner
	vrfs    []*types.HareEligibility
}

func (h *Hare) compactProposals(layer types.LayerID,
	proposals []*types.Proposal,
) []types.CompactProposalID {
	_ = "STUB: not implemented"
	return nil
}

func (h *Hare) compactProposalIds(layer types.LayerID,
	proposals []types.ProposalID,
) ([]types.CompactProposalID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we must handle this explicitly or we risk a panic on
// a nil slice access below

type proposalTuple struct {
	id      types.ProposalID
	compact types.CompactProposalID
}
