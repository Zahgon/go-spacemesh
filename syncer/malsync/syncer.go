package malsync

import (
	"context"
	"time"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type counter interface {
	Inc()
}

type noCounter struct{}

func (noCounter) Inc() { _ = "STUB: not implemented"; return }

type Opt func(*Syncer)

func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithPeerErrMetric(counter counter) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func withClock(clock clockwork.Clock) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type Config struct {
	// IDRequestInterval specifies the interval for malfeasance proof id requests to the network.
	IDRequestInterval time.Duration `mapstructure:"id-request-interval"`

	// MalfeasanceIDPeers is the number of peers to fetch node IDs for malfeasance proofs from.
	MalfeasanceIDPeers int `mapstructure:"malfeasance-id-peers"`

	// Minimum number of peers to sync against for initial sync to be considered complete.
	MinSyncPeers int `mapstructure:"min-sync-peers"`

	// MaxEpochFraction specifies maximum fraction of epoch to expire before
	// synchronous malfeasance proof sync is needed upon startup.
	MaxEpochFraction float64 `mapstructure:"max-epoch-fraction"`

	// MaxBatchSize is the maximum number of node IDs to sync in a single request.
	MaxBatchSize int `mapstructure:"max-batch-size"`

	// RequestsLimit is the maximum number of requests for a single malfeasance proof.
	//
	// The purpose of it is to prevent peers from advertising invalid node ID and disappearing.
	// Which will make node ask other peers for invalid malfeasance proofs.
	// It will be reset to 0 once malfeasance proof is advertised again.
	RequestsLimit int `mapstructure:"requests-limit"`

	// RetryInterval specifies retry interval for the initial sync.
	RetryInterval time.Duration `mapstructure:"retry-interval"`
}

func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

type syncPeerSet map[p2p.Peer]struct{}

func (sps syncPeerSet) add(peer p2p.Peer) { _ = "STUB: not implemented"; return }

func (sps syncPeerSet) clear() { _ = "STUB: not implemented"; return }

func (sps syncPeerSet) updateFrom(other syncPeerSet) { _ = "STUB: not implemented"; return }

// syncState stores malfeasance sync state.
type syncState struct {
	limit        int
	initial      bool
	state        map[types.NodeID]int
	syncingPeers syncPeerSet
	syncedPeers  syncPeerSet
}

func newSyncState(limit int, initial bool) *syncState { _ = "STUB: not implemented"; return nil }

func (sst *syncState) done() { _ = "STUB: not implemented"; return }

func (sst *syncState) numSyncedPeers() int { _ = "STUB: not implemented"; return 0 }

func (sst *syncState) update(update malUpdate) { _ = "STUB: not implemented"; return }

func (sst *syncState) has(nodeID types.NodeID) bool { _ = "STUB: not implemented"; return false }

func (sst *syncState) failed(nodeID types.NodeID) {
	_ = "STUB: not implemented"
	// possibly temporary failure, count failed attempt
	return
}

func (sst *syncState) rejected(nodeID types.NodeID) {
	_ = "STUB: not implemented"
	// malfeasance proof didn't pass validation, no sense in requesting it anymore
	return
}

func (sst *syncState) downloaded(nodeID types.NodeID) { _ = "STUB: not implemented"; return }

func (sst *syncState) missing(max int, has func(nodeID types.NodeID) (bool, error)) ([]types.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// already downloaded

// unsuccessfully requested too many times

type Syncer struct {
	logger        *zap.Logger
	cfg           Config
	fetcher       fetcher
	db            sql.Executor
	localDB       sql.LocalDatabase
	clock         clockwork.Clock
	layerClock    layerClock
	peerErrMetric counter
}

func New(fetcher fetcher, db sql.Executor, localDB sql.LocalDatabase, layerClock layerClock, opts ...Opt) *Syncer {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) shouldSyncLegacy(epochStart, epochEnd time.Time) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Syncer) shouldSync(epochStart, epochEnd time.Time) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Syncer) downloadLegacy(parent context.Context, initial bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) download(parent context.Context, initial bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) downloadLegacyNodeIDs(ctx context.Context, initial bool, updates chan<- malUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(ivan4th) this has to be randomized in a followup
// when sync will be scheduled in advance, in order to smooth out request rate across the network

func (s *Syncer) downloadNodeIDs(ctx context.Context, initial bool, updates chan<- malUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(ivan4th) this has to be randomized in a followup
// when sync will be scheduled in advance, in order to smooth out request rate across the network

func (s *Syncer) updateLegacyState() error { _ = "STUB: not implemented"; return nil }

func (s *Syncer) updateState() error { _ = "STUB: not implemented"; return nil }

func (s *Syncer) downloadLegacyMalfeasanceProofs(ctx context.Context, initial bool, updates <-chan malUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

// If we have some hashes to fetch already, don't wait for
// another update

// TODO(ivan4th): check multiple node IDs at once in a single SQL query

func (s *Syncer) downloadMalfeasanceProofs(ctx context.Context, initial bool, updates <-chan malUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

// If we have some hashes to fetch already, don't wait for
// another update

// TODO(mafa): check multiple node IDs at once in a single SQL query

func (s *Syncer) EnsureLegacyInSync(ctx context.Context, epochStart, epochEnd time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) EnsureInSync(ctx context.Context, epochStart, epochEnd time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Syncer) DownloadLoop(parent context.Context, malSyncStart types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

// if we don't have a malSyncStart epoch, just start legacy sync

// wait until first mal2 epoch

type malUpdate struct {
	peer    p2p.Peer
	nodeIDs []types.NodeID
}
