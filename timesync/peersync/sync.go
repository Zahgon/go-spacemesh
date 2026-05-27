package peersync

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/p2p"
)

const (
	protocolName = "/peersync/1.0/"
)

var (
	// errPeersNotSynced returned if system clock is out of sync with peers clock for configured period of time.
	errPeersNotSynced = errors.New("timesync: peers are not time synced")
	// errTimesyncFailed returned if we weren't able to collect enough clock samples from peers.
	errTimesyncFailed = errors.New("timesync: failed request")
)

//go:generate mockgen -typed -package=mocks -destination=./mocks/mocks.go -source=./sync.go

// Time provides interface for current time.
type Time interface {
	Now() time.Time
}

type systemTime struct{}

func (s systemTime) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type getPeers interface {
	GetPeers() []p2p.Peer
}

//go:generate scalegen -types Request,Response

// Request is a sync request.
type Request struct {
	ID uint64
}

// Response is a sync response.
type Response struct {
	ID        uint64
	Timestamp uint64
}

// DefaultConfig for Sync.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// Config for Sync.
type Config struct {
	Disable            bool          `mapstructure:"disable"`
	RoundRetryInterval time.Duration `mapstructure:"round-retry-interval"`
	RoundInterval      time.Duration `mapstructure:"round-interval"`
	RoundTimeout       time.Duration `mapstructure:"round-timeout"`
	MaxClockOffset     time.Duration `mapstructure:"max-clock-offset"`
	MaxOffsetErrors    int           `mapstructure:"max-offset-errors"`
	RequiredResponses  int           `mapstructure:"required-responses"`
}

// Option to modify Sync behavior.
type Option func(*Sync)

// WithTime modifies source of time used in Sync.
func WithTime(t Time) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLog modifies Log used in Sync.
func WithLog(lg *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConfig modifies config used in Sync.
func WithConfig(config Config) Option { _ = "STUB: not implemented"; return *new(Option) }

// New creates Sync instance and returns pointer.
func New(h host.Host, peers getPeers, opts ...Option) *Sync { _ = "STUB: not implemented"; return nil }

// Sync manages background worker that compares peers time with system time.
type Sync struct {
	config Config
	log    *zap.Logger
	time   Time
	h      host.Host
	peers  getPeers

	eg errgroup.Group

	cancelMtx sync.Mutex
	cancel    context.CancelFunc
}

func (s *Sync) streamHandler(stream network.Stream) { _ = "STUB: not implemented"; return }

// Start background workers.
func (s *Sync) Start() { _ = "STUB: not implemented"; return }

// already started

// Stop background workers.
func (s *Sync) Stop() { _ = "STUB: not implemented"; return }

// already stopped

// Wait will return first error that is returned by background workers.
func (s *Sync) Wait() error { _ = "STUB: not implemented"; return nil }

func (s *Sync) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// GetOffset computes offset from received response. The method is stateless and safe to use concurrently.
func (s *Sync) GetOffset(ctx context.Context, id uint64, prs []p2p.Peer) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
