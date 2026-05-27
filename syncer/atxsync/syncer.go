package atxsync

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/fetch"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/atxsync"
	"github.com/spacemeshos/go-spacemesh/system"
)

//go:generate mockgen -typed -package=mocks -destination=./mocks/mocks.go -source=./syncer.go

type fetcher interface {
	SelectBestShuffled(int) []p2p.Peer
	PeerEpochInfo(context.Context, p2p.Peer, types.EpochID) (*fetch.EpochData, error)
	system.AtxFetcher
}

type Opt func(*Syncer)

func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type Config struct {
	// EpochInfoInterval between epoch info requests to the network.
	EpochInfoInterval time.Duration `mapstructure:"epoch-info-request-interval"`
	// EpochInfoPeers is the number of peers we will ask for epoch info, every epoch info requests interval.
	EpochInfoPeers int `mapstructure:"epoch-info-peers"`

	// RequestsLimit is the maximum number of requests for single activation.
	//
	// The purpose of it is to prevent peers from advertising invalid atx and disappearing.
	// Which will make node ask other peers for invalid atx.
	// It will be reset to 0 once atx advertised again.
	RequestsLimit int `mapstructure:"requests-limit"`

	// AtxsBatch is the maximum number of atxs to sync in a single request.
	AtxsBatch int `mapstructure:"atxs-batch"`

	// ProgressFraction will report progress every fraction from total is downloaded.
	ProgressFraction float64 `mapstructure:"progress-every-fraction"`
	// ProgressInterval will report progress every interval.
	ProgressInterval time.Duration `mapstructure:"progress-on-time"`
}

func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func New(fetcher fetcher, db sql.Executor, localdb sql.LocalDatabase, opts ...Opt) *Syncer {
	_ = "STUB: not implemented"
	return nil
}

type Syncer struct {
	logger  *zap.Logger
	cfg     Config
	fetcher fetcher
	db      sql.Executor
	localdb sql.LocalDatabase
}

func (s *Syncer) Download(parent context.Context, publish types.EpochID, downloadUntil time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// in case of immediate we will request epoch info without waiting EpochInfoInterval

// termination requires two conditions:
// - epoch info has to be successfully downloaded close to or after the epoch start
// - all atxs from that epoch have to be downloaded or they are unavailable.
//   atx is unavailable if it was requested more than RequestsLimit times, and no peer provided it.

func (s *Syncer) downloadEpochInfo(
	ctx context.Context,
	publish types.EpochID,
	immediate bool,
	updates chan<- epochUpdate,
) error {
	_ = "STUB: not implemented"
	return nil
}

// not really immediate, to avoid an endless loop that doesn't wait between requests

// randomize interval to avoid sync spikes

// do not run it concurrently, epoch info is large and will continue to grow

// adding hashes to fetcher is not useful as they overflow the cache and are not used
// so we switch to asking best peers immediately

// after first success switch to requests after interval

func (s *Syncer) downloadAtxs(
	ctx context.Context,
	publish types.EpochID,
	downloadUntil time.Time,
	state atxsync.EpochSyncState,
	updates <-chan epochUpdate,
) error {
	_ = "STUB: not implemented"
	return nil
}

// waiting for update if there is nothing to download

// otherwise check updates periodically but don't stop downloading

// drop from memory if we already persisted info that this atx is not available

// if atx invalid there is no pointing in re-downloading it again

type epochUpdate struct {
	time   time.Time
	update atxsync.EpochSyncState
}
