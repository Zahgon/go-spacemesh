package syncer

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/activation"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/fetch"
	"github.com/spacemeshos/go-spacemesh/fetch/peers"
	"github.com/spacemeshos/go-spacemesh/mesh"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/sync2"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
	"github.com/spacemeshos/go-spacemesh/syncer/atxsync"
	"github.com/spacemeshos/go-spacemesh/syncer/malsync"
	"github.com/spacemeshos/go-spacemesh/system"
)

// Config is the config params for syncer.
type Config struct {
	Interval         time.Duration `mapstructure:"interval"`
	EpochEndFraction float64       `mapstructure:"epochendfraction"`
	HareDelayLayers  uint32        `mapstructure:"-"` // not configurable, overwritten by tortoise.Config.Zdist
	SyncCertDistance uint32        `mapstructure:"-"` // not configurable, overwritten by tortoise.Config.Hdist
	// TallyVotesFrequency how often to tally votes during layers sync.
	// Setting this to 0.25 will tally votes after downloading data for quarter of the epoch.
	TallyVotesFrequency      float64
	MaxStaleDuration         time.Duration    `mapstructure:"maxstaleduration"`
	Standalone               bool             `mapstructure:"-"` // not configurable, overwritten by BaseCfg.Standalone
	GossipDuration           time.Duration    `mapstructure:"gossipduration"`
	DisableMeshAgreement     bool             `mapstructure:"disable-mesh-agreement"`
	OutOfSyncThresholdLayers uint32           `mapstructure:"out-of-sync-threshold"`
	AtxSync                  atxsync.Config   `mapstructure:"atx-sync"`
	MalSync                  malsync.Config   `mapstructure:"malfeasance-sync"`
	ReconcSync               ReconcSyncConfig `mapstructure:"reconc-sync"`
}

type ReconcSyncConfig struct {
	Enable            bool               `mapstructure:"enable"`
	EnableActiveSync  bool               `mapstructure:"enable-active-sync"`
	OldAtxSyncCfg     sync2.Config       `mapstructure:"old-atx-sync"`
	NewAtxSyncCfg     sync2.Config       `mapstructure:"new-atx-sync"`
	ParallelLoadLimit int                `mapstructure:"parallel-load-limit"`
	HardTimeout       time.Duration      `mapstructure:"hard-timeout"`
	ServerConfig      fetch.ServerConfig `mapstructure:"server-config"`
}

// DefaultConfig for the syncer.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type syncState uint32

const (
	// notSynced is the state where the node is outOfSyncThreshold layers or more behind the current layer.
	notSynced syncState = iota
	// gossipSync is the state in which a node listens to at least one full layer of gossip before participating
	// in the protocol. This is to protect the node from participating in the consensus without full information.
	// For example, when a node wakes up in the middle of layer N, since it didn't receive all relevant messages and
	// blocks of layer N, it shouldn't vote or produce blocks in layer N+1. It instead listens to gossip for all
	// through layer N+1 and starts producing blocks and participates in hare committee in layer N+2.
	gossipSync
	// synced is the state where the node is in sync with its peers.
	synced
)

func (s syncState) String() string { _ = "STUB: not implemented"; return "" }

var (
	errHareInCharge  = errors.New("hare in charge of layer")
	errATXsNotSynced = errors.New("ATX not synced")
)

// Option is a type to configure a syncer.
type Option func(*Syncer)

// WithConfig ...
func WithConfig(c Config) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogger ...
func WithLogger(l *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAtxVersions(v activation.AtxVersions) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func withDataFetcher(d fetchLogic) Option { _ = "STUB: not implemented"; return *new(Option) }

func withForkFinder(f forkFinder) Option { _ = "STUB: not implemented"; return *new(Option) }

func withAtxSyncerV2(asv2 multiEpochAtxSyncerV2) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Syncer is responsible to keep the node in sync with the network.
type Syncer struct {
	logger       *zap.Logger
	cfg          Config
	cdb          *datastore.CachedDB
	atxsyncer    atxSyncer
	malsyncer    malSyncer
	ticker       layerTicker
	mesh         *mesh.Mesh
	tortoise     system.Tortoise
	certHandler  certHandler
	dataFetcher  fetchLogic
	patrol       layerPatrol
	forkFinder   forkFinder
	syncOnce     sync.Once
	syncState    atomic.Value
	atxSyncState atomic.Value
	isBusy       atomic.Bool
	// syncedTargetTime is used to signal at which time we can set this node to synced state
	syncedTargetTime time.Time
	lastLayerSynced  atomic.Uint32
	lastEpochSynced  atomic.Uint32
	stateErr         atomic.Bool

	// malSyncStartEpoch is the epoch from which we start syncing malfeasance proofs through the new protocol.
	malSyncStartEpoch types.EpochID

	// backgroundSync always runs one sync operation in the background.
	backgroundSync struct {
		epoch  atomic.Uint32
		eg     errgroup.Group
		cancel context.CancelFunc
	}

	// malSync runs malfeasant identity sync in the background
	malSync struct {
		started bool
		eg      errgroup.Group
	}

	// awaitATXSyncedCh is the list of subscribers' channels to notify when this node enters ATX synced state
	awaitATXSyncedCh chan struct{}

	eg   errgroup.Group
	stop context.CancelFunc

	asv2       multiEpochAtxSyncerV2
	dispatcher *rangesync.Dispatcher
}

// NewSyncer creates a new Syncer instance.
func NewSyncer(
	cdb *datastore.CachedDB,
	ticker layerTicker,
	mesh *mesh.Mesh,
	tortoise system.Tortoise,
	fetcher fetcher,
	peerCache *peers.Peers,
	host host.Host,
	patrol layerPatrol,
	ch certHandler,
	atxSyncer atxSyncer,
	malSyncer malSyncer,
	opts ...Option,
) (*Syncer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close stops the syncing process and the goroutines syncer spawns.
func (s *Syncer) Close() { _ = "STUB: not implemented"; return }

// not started yet

// RegisterForATXSynced returns a channel for notification when the node enters ATX synced state.
func (s *Syncer) RegisterForATXSynced() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// ListenToGossip returns true if the node is listening to gossip for blocks/TXs data.
func (s *Syncer) ListenToGossip() bool { _ = "STUB: not implemented"; return false }

// ListenToATXGossip returns true if the node is listening to gossip for ATXs data.
func (s *Syncer) ListenToATXGossip() bool { _ = "STUB: not implemented"; return false }

// IsSynced returns true if the node is in synced state.
func (s *Syncer) IsSynced(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// Start starts the main sync loop that tries to sync data for every SyncInterval.
func (s *Syncer) Start() { _ = "STUB: not implemented"; return }

func (s *Syncer) setATXSynced() { _ = "STUB: not implemented"; return }

func (s *Syncer) getATXSyncState() syncState { _ = "STUB: not implemented"; return *new(syncState) }

func (s *Syncer) getSyncState() syncState { _ = "STUB: not implemented"; return *new(syncState) }

func (s *Syncer) setSyncState(ctx context.Context, newState syncState) {
	_ = "STUB: not implemented"
	return
}

// setSyncerBusy returns false if the syncer is already running a sync process.
// Otherwise it sets syncer to be busy and returns true.
func (s *Syncer) setSyncerBusy() bool { _ = "STUB: not implemented"; return false }

func (s *Syncer) setSyncerIdle() { _ = "STUB: not implemented"; return }

func (s *Syncer) setLastSyncedLayer(lid types.LayerID) { _ = "STUB: not implemented"; return }

func (s *Syncer) getLastSyncedLayer() types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

func (s *Syncer) setLastAtxEpoch(epoch types.EpochID) { _ = "STUB: not implemented"; return }

func (s *Syncer) lastAtxEpoch() types.EpochID {
	_ = "STUB: not implemented"
	return *new(types.EpochID)
}

// synchronize sync data up to the currentLayer-1 and wait for the layers to be validated.
// It returns false if the data sync failed.
func (s *Syncer) synchronize(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// at most one synchronize process can run at any time

// no need to worry about race condition for s.run. only one instance of synchronize can run at a time

// TODO
// https://github.com/spacemeshos/go-spacemesh/issues/3987

// check that we have any peers

// always sync to currentLayer-1 to reduce race with gossip and hare/tortoise

// BatchError spams too much, in case of no progress enable debug mode for sync

func (s *Syncer) ensureATXsInSync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// nothing to sync in epoch 0

// if we are not advanced enough sync previous epoch, otherwise start syncing activations published in this epoch

// on startup always download all activations that were published before current epoch

// ensureATXsInSyncV2 ensures that the ATXs are in sync and being synchronized
// continuously using syncv2.
func (s *Syncer) ensureATXsInSyncV2(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// nothing to sync in epoch 0

// ATXs are not in sync yet, to we need to sync them synchronously

// When active syncv2 is not enabled, this will only cause the per-epoch sync
// servers (multiplexed via dispatcher) to be activated, without attempting to
// initiate sync against the peers

func (s *Syncer) ensureMalfeasanceInSync(ctx context.Context) error {
	_ = "STUB: not implemented"
	// TODO: use syncv2 for malfeasance proofs:
	return nil
}

// Malfeasance proofs are synced after the actual ATXs.
// We set ATX synced status after both ATXs and malfeasance proofs
// are in sync.

func (s *Syncer) syncAtxAndMalfeasance(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If syncv2 is being used in server-only mode, we still need to run
// active syncv1.

func isTooFarBehind(
	ctx context.Context,
	logger *zap.Logger,
	current, lastSynced types.LayerID,
	outOfSyncThreshold uint32,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Syncer) setStateBeforeSync(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Syncer) dataSynced() bool { _ = "STUB: not implemented"; return false }

func (s *Syncer) setStateAfterSync(ctx context.Context, success bool) {
	_ = "STUB: not implemented"
	return
}

// for the gossipSync/notSynced states, we check if the mesh state is on target before we advance sync state.
// but for the synced state, we don't check the mesh state because gossip+hare+tortoise are in charge of
// advancing processed/verified layers.  syncer is just auxiliary that fetches data in case of a temporary
// network outage.

// push out the target synced layer

// if we have gossip-synced long enough, we are ready to participate in consensus

// wait till s.ticker.GetCurrentLayer() + numGossipSyncLayers to participate in consensus

func (s *Syncer) syncMalfeasance(parent context.Context, epoch types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(mafa): remove this check again when ATXv2 is live https://github.com/spacemeshos/go-spacemesh/issues/6716

func (s *Syncer) syncLayer(ctx context.Context, layerID types.LayerID, peers ...p2p.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

// fetching ATXs published the specified epoch.
func (s *Syncer) fetchATXsForEpoch(ctx context.Context, publish types.EpochID, background bool) error {
	_ = "STUB: not implemented"
	return nil
}

// waitBackgroundSync is a helper to wait for the background sync to finish.
func (s *Syncer) waitBackgroundSync() { _ = "STUB: not implemented"; return }
