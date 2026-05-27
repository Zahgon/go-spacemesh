package sync2

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/fetch/peers"
	"github.com/spacemeshos/go-spacemesh/sync2/multipeer"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// Config contains the configuration for the P2PHashSync.
type Config struct {
	rangesync.RangeSetReconcilerConfig  `mapstructure:",squash"`
	multipeer.MultiPeerReconcilerConfig `mapstructure:",squash"`
	TrafficLimit                        int           `mapstructure:"traffic-limit"`
	MessageLimit                        int           `mapstructure:"message-limit"`
	MaxDepth                            uint          `mapstructure:"max-depth"`
	BatchSize                           uint          `mapstructure:"batch-size"`
	MaxAttempts                         uint          `mapstructure:"max-attempts"`
	MaxBatchRetries                     uint          `mapstructure:"max-batch-retries"`
	FailedBatchDelay                    time.Duration `mapstructure:"failed-batch-delay"`
	AdvanceInterval                     time.Duration `mapstructure:"advance-interval"`
}

func (cfg *Config) Validate(logger *zap.Logger) bool { _ = "STUB: not implemented"; return false }

// always invoke Validate to log validation errors

// DefaultConfig returns the default configuration for the P2PHashSync.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// P2PHashSync is handles the synchronization of a local OrderedSet against other peers.
type P2PHashSync struct {
	logger           *zap.Logger
	cfg              Config
	enableActiveSync bool
	os               rangesync.OrderedSet
	syncBase         multipeer.SyncBase
	reconciler       *multipeer.MultiPeerReconciler
	cancel           context.CancelFunc
	eg               errgroup.Group
	startOnce        sync.Once
	running          atomic.Bool
	kickCh           chan struct{}
}

// NewP2PHashSync creates a new P2PHashSync.
func NewP2PHashSync(
	logger *zap.Logger,
	d *rangesync.Dispatcher,
	name string,
	os rangesync.OrderedSet,
	keyLen int,
	peers *peers.Peers,
	handler multipeer.SyncKeyHandler,
	cfg Config,
	enableActiveSync bool,
) (*P2PHashSync, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set returns the OrderedSet that is being synchronized.
func (s *P2PHashSync) Set() rangesync.OrderedSet {
	_ = "STUB: not implemented"

	// Load loads the OrderedSet from the underlying storage.
	return *new(rangesync.OrderedSet)
}

func (s *P2PHashSync) Load() error { _ = "STUB: not implemented"; return nil }

// We pre-load the set to avoid waiting for it to load during a
// sync request

func (s *P2PHashSync) start() (isWaiting bool) { _ = "STUB: not implemented"; return false }

// Start starts the multi-peer reconciler if it is not already running.
func (s *P2PHashSync) Start() {
	_ = "STUB: not implemented"

	// StartAndSync starts the multi-peer reconciler if it is not already running, and waits
	// until the local OrderedSet is in sync with the peers.
	return
}

func (s *P2PHashSync) StartAndSync(ctx context.Context) error {
	_ = "STUB: not implemented"

	// If the multipeer reconciler is waiting for sync, we kick it to start
	// the sync so as not to wait for the next scheduled sync interval.
	return nil
}

// Stop stops the multi-peer reconciler.
func (s *P2PHashSync) Stop() { _ = "STUB: not implemented"; return }

// Synced returns true if the local OrderedSet is in sync with the peers, as determined by
// the multi-peer reconciler.
func (s *P2PHashSync) Synced() bool { _ = "STUB: not implemented"; return false }

var errStopped = errors.New("syncer stopped")

// WaitForSync waits until the local OrderedSet is in sync with the peers.
func (s *P2PHashSync) WaitForSync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// SyncCycleCount returns the number of sync cycles that have happened,
// no matter if they were successful or not.
func (s *P2PHashSync) SyncCycleCount() int { _ = "STUB: not implemented"; return 0 }
