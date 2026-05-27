package multipeer

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/fetch/peers"
	"github.com/spacemeshos/go-spacemesh/p2p"
)

const (
	Protocol = "sync/2"
)

type syncability struct {
	// peers that were probed successfully
	syncable []p2p.Peer
	// peers that have enough items for split sync
	splitSyncable []p2p.Peer
	// Number of peers that are similar enough to this one for full sync
	nearFullCount int
	// Number of peers that are in sync with this one
	inSyncCount int
}

type runner struct {
	mpr *MultiPeerReconciler
}

var _ syncRunner = &runner{}

func (r *runner) SplitSync(ctx context.Context, syncPeers []p2p.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) FullSync(ctx context.Context, syncPeers []p2p.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiPeerReconcilerConfig contains the configuration for a MultiPeerReconciler.
type MultiPeerReconcilerConfig struct {
	// Number of peers to pick for synchronization.
	// Synchronization will still happen if fewer peers are available.
	SyncPeerCount uint `mapstructure:"sync-peer-count"`
	// Minimum number of peers for the split sync to happen.
	MinSplitSyncPeers uint `mapstructure:"min-split-sync-peers"`
	// Minimum number of items that a peer must have to be eligible for split sync
	// (subrange-per-peer).
	MinSplitSyncCount uint `mapstructure:"min-split-sync-count"`
	// Maximum approximate size of symmetric difference between the local set and the
	// remote one for the sets to be considered "mostly in sync", so that full sync is
	// preferred to split sync.
	MaxFullDiff uint `mapstructure:"max-full-diff"`
	// Maximum number of items that a peer can have less than the local set for it to
	// be considered for synchronization.
	MaxSyncDiff uint `mapstructure:"max-sync-diff"`
	// Minimum fraction (0..1) of "mostly synced" peers starting with which full sync
	// is used instead of split sync.
	MinCompleteFraction float64 `mapstructure:"min-complete-fraction"`
	// Interval between syncs.
	SyncInterval time.Duration `mapstructure:"sync-interval"`
	// Interval spread factor for split sync.
	// The actual interval will be SyncInterval * (1 + (random[0..2]*SplitSyncIntervalSpread-1)).
	// So, if you want the actual interval to be with the range of SplitInterval ± 25%,
	// set this to 0.25.
	SyncIntervalSpread float64 `mapstructure:"sync-interval-spread"`
	// Interval between retries after a failed sync.
	RetryInterval time.Duration `mapstructure:"retry-interval"`
	// Interval between rechecking for peers after no synchronization peers were
	// found.
	NoPeersRecheckInterval time.Duration `mapstructure:"no-peers-recheck-interval"`
	// Grace period for split sync peers.
	// If a peer doesn't complete syncing its range within the specified duration
	// during split sync, its range is assigned additionally to another quicker
	// peer. The sync against the "slow" peer is NOT stopped immediately after that.
	SplitSyncGracePeriod time.Duration `mapstructure:"split-sync-grace-period"`
	// Minimum number of full syncs that must have happened within the
	// fullSyncednessPeriod for the node to be considered fully synced
	MinFullSyncednessCount uint `mapstructure:"min-full-syncedness-count"`
	// Duration within which the minimum number of full syncs must have happened for
	// the node to be considered fully synced.
	FullSyncednessPeriod time.Duration `mapstructure:"full-syncedness-count"`
}

func (cfg *MultiPeerReconcilerConfig) Validate(logger *zap.Logger) bool {
	_ = "STUB: not implemented"
	// Join the errors together so that the user doesn't have to fix one at a time.
	return false
}

// DefaultConfig returns the default configuration for the MultiPeerReconciler.
func DefaultConfig() MultiPeerReconcilerConfig {
	_ = "STUB: not implemented"
	return *new(MultiPeerReconcilerConfig)
}

// MultiPeerReconciler reconciles the local set against multiple remote sets.
type MultiPeerReconciler struct {
	logger         *zap.Logger
	cfg            MultiPeerReconcilerConfig
	syncBase       SyncBase
	peers          *peers.Peers
	clock          clockwork.Clock
	keyLen         int
	maxDepth       int
	runner         syncRunner
	sl             *syncList
	syncCycleCount atomic.Uint32
}

func newMultiPeerReconciler(
	logger *zap.Logger,
	cfg MultiPeerReconcilerConfig,
	syncBase SyncBase,
	peers *peers.Peers,
	keyLen, maxDepth int,
	syncRunner syncRunner,
	clock clockwork.Clock,
) *MultiPeerReconciler {
	_ = "STUB: not implemented"
	return nil
}

// NewMultiPeerReconciler creates a new MultiPeerReconciler.
func NewMultiPeerReconciler(
	logger *zap.Logger,
	cfg MultiPeerReconcilerConfig,
	syncBase SyncBase,
	peers *peers.Peers,
	keyLen, maxDepth int,
) *MultiPeerReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (mpr *MultiPeerReconciler) probePeers(ctx context.Context, syncPeers []p2p.Peer) (syncability, error) {
	_ = "STUB: not implemented"
	return *new(syncability), nil
}

// We need to close probeCh for the loop below to terminate, and we must do that
// only after all the goroutines above have finished.

// We do not consider peers with substantially fewer items than the local
// set for active sync. It's these peers' responsibility to request sync
// against this node.

func (mpr *MultiPeerReconciler) needSplitSync(s syncability) bool {
	_ = "STUB: not implemented"
	return false
}

// enough peers are close to this one according to minhash score, can do
// full sync

// would be nice to do split sync, but not enough peers for that

func (mpr *MultiPeerReconciler) fullSync(ctx context.Context, syncPeers []p2p.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

// failing to sync against a particular peer is not considered
// a fatal sync failure, so we just log the error

func (mpr *MultiPeerReconciler) syncOnce(ctx context.Context, lastWasSplit bool) (full bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// probePeers doesn't return transient errors, sync must stop if it failed

// We try to sync against peers which are not in full sync with this one.
// If there are no such peers, but there are some which are in sync with
// us, we consider this node to be in sync.

// Run runs the MultiPeerReconciler.
func (mpr *MultiPeerReconciler) Run(ctx context.Context, kickCh chan struct{}) error {
	_ = "STUB: not implemented"
	// The point of using split sync, which syncs different key ranges against
	// different peers, vs full sync which syncs the full key range against different
	// peers, is:
	// 1. Avoid getting too many range splits and thus network transfer overhead
	// 2. Avoid fetching same keys from multiple peers
	return nil
}

// States:
// A. Wait. Pause for sync interval
//    Timeout => A
// B. No peers -> do nothing.
//    Got any peers => C
// C. Low on peers. Wait for more to appear
//    Lost all peers => B
//    Got enough peers => D
//    Timeout => D
// D. Probe the peers. Use successfully probed ones in states E/F
//      Drop failed peers from the peer set while polling.
//    All probes failed => B
//    Last sync was split sync => E
//    N of peers < minSplitSyncPeers => E
//    All are low on count (minSplitSyncCount) => F
//    Enough peers (minCompleteFraction) with diffSize <= maxFullDiff => E
//      diffSize = (1-sim)*localItemCount
//    Otherwise => F
// E. Full sync. Run full syncs against each peer
//    All syncs completed (success / fail) => A
// F. Bounded sync. Subdivide the range by peers and start syncs.
//      Use peers with > minSplitSyncCount
//      Wait for all the syncs to complete/fail
//    All syncs completed (success / fail) => A

// Split sync needs to be followed by a full sync.
// Don't wait to have sync move forward quicker.
// In most cases, the full sync will be very quick.

// Synced returns true if the node is considered synced, that is, the specified
// number of full syncs has happened within the specified duration of time.
func (mpr *MultiPeerReconciler) Synced() bool { _ = "STUB: not implemented"; return false }

// SyncCycleCount returns the number of sync cycles that have happened,
// no matter if they were successful or not.
func (mpr *MultiPeerReconciler) SyncCycleCount() int { _ = "STUB: not implemented"; return 0 }
