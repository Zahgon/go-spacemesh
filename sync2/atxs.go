package sync2

import (
	"context"

	"github.com/jonboulle/clockwork"
	"github.com/libp2p/go-libp2p/core/host"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/fetch/peers"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/server"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sync2/multipeer"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
	"github.com/spacemeshos/go-spacemesh/sync2/sqlstore"
)

const (
	proto = "sync/2"
)

type ATXHandler struct {
	logger *zap.Logger
	f      Fetcher
	clock  clockwork.Clock
	cfg    Config
}

var _ multipeer.SyncKeyHandler = &ATXHandler{}

func NewATXHandler(
	logger *zap.Logger,
	f Fetcher,
	cfg Config,
	clock clockwork.Clock,
) *ATXHandler {
	_ = "STUB: not implemented"
	return nil
}

type commitState struct {
	state         map[types.ATXID]uint
	total         int
	numDownloaded int
	items         []types.ATXID
}

func (h *ATXHandler) setupState(
	peer p2p.Peer,
	base rangesync.OrderedSet,
	received rangesync.SeqResult,
) (*commitState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *ATXHandler) getAtxs(ctx context.Context, cs *commitState) (bool, error) {
	_ = "STUB: not implemented"
	return false,
		// reuse the slice to reduce allocations
		nil
}

func (h *ATXHandler) Commit(
	ctx context.Context,
	peer p2p.Peer,
	base rangesync.OrderedSet,
	received rangesync.SeqResult,
) error {
	_ = "STUB: not implemented"
	return nil
}

type MultiEpochATXSyncer struct {
	logger            *zap.Logger
	oldCfg            Config
	newCfg            Config
	parallelLoadLimit int
	hss               HashSyncSource
	newEpoch          types.EpochID
	atxSyncers        []HashSync
}

func NewMultiEpochATXSyncer(
	logger *zap.Logger,
	hss HashSyncSource,
	oldCfg, newCfg Config,
	parallelLoadLimit int,
) (*MultiEpochATXSyncer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MultiEpochATXSyncer) load(newEpoch types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

// EnsureSync ensures that ATX sync is active for all the epochs up to and including
// currentEpoch, and that all ATXs are
// synced up to and including lastWaitEpoch.
// If newEpoch argument is non-zero, faster but less memory efficient sync is used for
// that epoch, based on the newCfg (larger maxDepth).
// For other epochs, oldCfg is used which corresponds to slower but more memory efficient
// sync (smaller maxDepth).
// It returns the last epoch that was synced synchronously.
func (s *MultiEpochATXSyncer) EnsureSync(
	ctx context.Context,
	lastWaitEpoch, newEpoch types.EpochID,
) (lastSynced types.EpochID, err error) {
	_ = "STUB: not implemented"
	return *new(types.EpochID), nil
}

// Stop stops all ATX syncers.
func (s *MultiEpochATXSyncer) Stop() { _ = "STUB: not implemented"; return }

func atxsTable(epoch types.EpochID) *sqlstore.SyncedTable { _ = "STUB: not implemented"; return nil }

func NewATXSyncer(
	logger *zap.Logger,
	d *rangesync.Dispatcher,
	name string,
	cfg Config,
	db sql.Database,
	f Fetcher,
	peers *peers.Peers,
	epoch types.EpochID,
	enableActiveSync bool,
) (*P2PHashSync, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDispatcher(logger *zap.Logger, host host.Host, opts []server.Opt) *rangesync.Dispatcher {
	_ = "STUB: not implemented"
	return nil
}

type ATXSyncSource struct {
	logger           *zap.Logger
	d                *rangesync.Dispatcher
	db               sql.Database
	f                Fetcher
	peers            *peers.Peers
	enableActiveSync bool
}

var _ HashSyncSource = &ATXSyncSource{}

func NewATXSyncSource(
	logger *zap.Logger,
	d *rangesync.Dispatcher,
	db sql.Database,
	f Fetcher,
	peers *peers.Peers,
	enableActiveSync bool,
) *ATXSyncSource {
	_ = "STUB: not implemented"
	return nil
}

// CreateHashSync implements HashSyncSource.
func (as *ATXSyncSource) CreateHashSync(name string, cfg Config, epoch types.EpochID) (HashSync, error) {
	_ = "STUB: not implemented"
	return *new(HashSync), nil
}
