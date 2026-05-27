package tortoise

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/common/types/result"
)

// Config for protocol parameters.
type Config struct {
	// how long we are waiting for a switch from verifying to full. relevant during rerun.
	Hdist                uint32 `mapstructure:"tortoise-hdist"`       // hare output look back distance
	Zdist                uint32 `mapstructure:"tortoise-zdist"`       // hare result wait distance
	WindowSize           uint32 `mapstructure:"tortoise-window-size"` // size of the tortoise sliding window (in layers)
	HistoricalWindowSize []WindowSizeInterval
	// ignored if candidate for base ballot has more than max exceptions
	MaxExceptions int `mapstructure:"tortoise-max-exceptions"`
	// number of layers to delay votes for blocks with bad beacon values during self-healing. ideally a full epoch.
	BadBeaconVoteDelayLayers uint32 `mapstructure:"tortoise-delay-layers"`
	// EnableTracer will write tortoise traces to the stderr.
	EnableTracer bool `mapstructure:"tortoise-enable-tracer"`
	// MinimalActiveSetWeight is a weight that will replace weight
	// recorded in the first ballot, if that weight is less than minimal
	// for purposes of eligibility computation.
	MinimalActiveSetWeight []types.EpochMinimalActiveWeight
	// CollectDetails sets numbers of layers to collect details.
	// Must be less than WindowSize.
	CollectDetails uint32 `mapstructure:"tortoise-collect-details"`
	LayerSize      uint32 `mapstructure:"-"` // not configurable, overwritten by BaseConfig.LayerAvgSize
}

type WindowSizeInterval struct {
	Start  types.LayerID
	End    types.LayerID
	Window uint32
}

// DefaultConfig for Tortoise.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// 100 layers of average size

func (c *Config) WindowSizeLayers(applied types.LayerID) types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

func (c *Config) WindowSizeEpochs(applied types.LayerID) types.EpochID {
	_ = "STUB: not implemented"
	return *new(types.EpochID)
}

// Tortoise is a thread safe verifying tortoise wrapper, it just locks all actions.
type Tortoise struct {
	logger *zap.Logger
	cfg    Config

	mu     sync.Mutex
	trtl   *turtle
	tracer *tracer
}

// Opt for configuring tortoise.
type Opt func(t *Tortoise)

// WithLogger defines logger for tortoise.
func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithConfig defines protocol parameters.
func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithTracer enables tracing of every call to the tortoise.
func WithTracer(opts ...TraceOpt) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// New creates Tortoise instance.
func New(atxdata *atxsdata.Data, opts ...Opt) (*Tortoise, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tortoise) RecoverFrom(lid types.LayerID, opinion, prev types.Hash32) {
	_ = "STUB: not implemented"
	return
}

// -1 so that iteration in tallyVotes starts from the target layer

// LatestComplete returns the latest verified layer.
func (t *Tortoise) LatestComplete() types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

func (t *Tortoise) OnWeakCoin(lid types.LayerID, coin bool) { _ = "STUB: not implemented"; return }

// OnMalfeasance registers node id as malfeasant.
// - ballots from this id will have zero weight
// - atxs - will not be counted towards global/local thresholds
// If node registers equivocating ballot/atx it should
// call OnMalfeasance before storing ballot/atx.
func (t *Tortoise) OnMalfeasance(id types.NodeID) { _ = "STUB: not implemented"; return }

func (t *Tortoise) OnBeacon(eid types.EpochID, beacon types.Beacon) {
	_ = "STUB: not implemented"
	return
}

type encodeConf struct {
	current *types.LayerID
}

// EncodeVotesOpts is for configuring EncodeVotes options.
type EncodeVotesOpts func(*encodeConf)

// EncodeVotesWithCurrent changes last known layer that will be used for encoding votes.
//
// NOTE(dshulyak) why do we need this?
// Tortoise computes threshold from last non-verified till the last known layer,
// since we don't download atxs before starting tortoise we won't be able to compute threshold
// based on the last clock layer (see https://github.com/spacemeshos/go-spacemesh/issues/3003)
func EncodeVotesWithCurrent(current types.LayerID) EncodeVotesOpts {
	_ = "STUB: not implemented"
	return *new(EncodeVotesOpts)
}

// EncodeVotes chooses a base ballot and creates a differences list. Needs the hare results for latest layers.
func (t *Tortoise) EncodeVotes(
	ctx context.Context,
	opts ...EncodeVotesOpts,
) (*types.Opinion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TallyVotes up to the specified layer.
func (t *Tortoise) TallyVotes(lid types.LayerID) { _ = "STUB: not implemented"; return }

// OnAtx is expected to be called before ballots that use this atx.
func (t *Tortoise) OnAtx(target types.EpochID, id types.ATXID, atx *atxsdata.ATX) {
	_ = "STUB: not implemented"
	return
}

// OnBlock updates tortoise with information that data is available locally.
func (t *Tortoise) OnBlock(header types.BlockHeader) { _ = "STUB: not implemented"; return }

// OnRecoveredBlocks uploads blocks to the state with all metadata.
//
// Implementation assumes that they will be uploaded in order.
func (t *Tortoise) OnRecoveredBlocks(lid types.LayerID, validity map[types.BlockHeader]bool, hare *types.BlockID) {
	_ = "STUB: not implemented"
	return
}

// OnBallot should be called every time new ballot is received.
// Dependencies (base ballot, ref ballot, active set and its own atx) must
// be processed before ballot.
func (t *Tortoise) OnBallot(ballot *types.BallotTortoiseData) { _ = "STUB: not implemented"; return }

// OnRecoveredBallot is called for ballots recovered from database.
//
// For recovered ballots base ballot is not required to be in state therefore
// opinion is not recomputed, but instead recovered from database state.
func (t *Tortoise) OnRecoveredBallot(ballot *types.BallotTortoiseData) {
	_ = "STUB: not implemented"
	return
}

// DecodedBallot created after unwrapping exceptions list and computing internal opinion.
type DecodedBallot struct {
	*types.BallotTortoiseData
	info *ballotInfo
	// after validation is finished we need to add new vote targets
	// for tortoise from the decoded votes. minHint identifies the boundary
	// until which we have to scan.
	minHint types.LayerID
}

type BallotData struct {
	ID            types.BallotID
	Layer         types.LayerID
	ATXID         types.ATXID
	Smesher       types.NodeID
	Beacon        types.Beacon
	Eligibilities uint32
}

func (t *Tortoise) GetBallot(id types.BallotID) *BallotData { _ = "STUB: not implemented"; return nil }

// DecodeBallot decodes ballot if it wasn't processed earlier.
func (t *Tortoise) DecodeBallot(ballot *types.BallotTortoiseData) (*DecodedBallot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tortoise) decodeBallot(ballot *types.BallotTortoiseData) (*DecodedBallot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StoreBallot stores previously decoded ballot.
func (t *Tortoise) StoreBallot(decoded *DecodedBallot) error { _ = "STUB: not implemented"; return nil }

// OnHareOutput should be called when hare terminated or certificate for a block
// was synced from a peer.
// This method is expected to be called any number of times, with layers in any order.
//
// This method should be called with EmptyBlockID if node received proof of malicious behavior,
// such as signing same block id by members of the same committee.
func (t *Tortoise) OnHareOutput(lid types.LayerID, bid types.BlockID) {
	_ = "STUB: not implemented"
	return
}

// GetMissingActiveSet returns unknown atxs from the original list. It is done for a specific epoch
// as active set atxs never cross epoch boundary.
func (t *Tortoise) GetMissingActiveSet(target types.EpochID, atxs []types.ATXID) []types.ATXID {
	_ = "STUB: not implemented"
	return nil
}

// OnApplied compares stored opinion with computed opinion and sets
// pending layer to the layer above equal layer.
// This method is meant to be used only in recovery from disk codepath.
func (t *Tortoise) OnApplied(lid types.LayerID, opinion types.Hash32) bool {
	_ = "STUB: not implemented"
	return false
}

// latestsResults returns at most N latest results from process layer.
//
// Private as it meant to be used for metering.
func (t *Tortoise) latestsResults(n uint32) []result.Layer { _ = "STUB: not implemented"; return nil }

// Updates returns list of layers where opinion was changed since previous call.
func (t *Tortoise) Updates() []result.Layer { _ = "STUB: not implemented"; return nil }

func (t *Tortoise) results(from, to types.LayerID) ([]result.Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Mode int

func (m Mode) String() string { _ = "STUB: not implemented"; return "" }

const (
	Verifying = 0
	Full      = 1
)

// Mode returns 0 for verifying.
func (t *Tortoise) Mode() Mode { _ = "STUB: not implemented"; return *new(Mode) }

// UpdateLastLayer updates last layer which is used for determining weight thresholds.
func (t *Tortoise) UpdateLastLayer(last types.LayerID) { _ = "STUB: not implemented"; return }

// UpdateVerified layers based on the previously known verified layer.
func (t *Tortoise) UpdateVerified(verified types.LayerID) { _ = "STUB: not implemented"; return }

func (t *Tortoise) WithinHdist(lid types.LayerID) bool { _ = "STUB: not implemented"; return false }
