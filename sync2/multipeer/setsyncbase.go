package multipeer

import (
	"context"
	"io"
	"sync"

	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// SetSyncBase is a synchronization base which holds the original OrderedSet.
// For each peer, a Syncer is derived from the base which is used to synchronize against
// that peer only. This way, there's no propagation of any keys for which the actual data
// has not been yet received and validated.
type SetSyncBase struct {
	mtx     sync.Mutex
	ps      PairwiseSyncer
	os      rangesync.OrderedSet
	handler SyncKeyHandler
}

var _ SyncBase = &SetSyncBase{}

// NewSetSyncBase creates a new SetSyncBase.
func NewSetSyncBase(
	ps PairwiseSyncer,
	os rangesync.OrderedSet,
	handler SyncKeyHandler,
) *SetSyncBase {
	_ = "STUB: not implemented"
	return nil
}

// Count implements SyncBase.
func (ssb *SetSyncBase) Count() (int, error) {
	_ = "STUB: not implemented"
	// In most cases ssb.os.SetInfo will not access the database, so we're not holding
	// the lock for long here.
	return 0, nil
}

// Advance implements SyncBase.
func (ssb *SetSyncBase) Advance() error { _ = "STUB: not implemented"; return nil }

func (ssb *SetSyncBase) syncPeer(
	ctx context.Context,
	p p2p.Peer,
	toCall func(rangesync.OrderedSet) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (ssb *SetSyncBase) Sync(ctx context.Context, p p2p.Peer, x, y rangesync.KeyBytes) error {
	_ = "STUB: not implemented"
	return nil
}

func (ssb *SetSyncBase) Serve(ctx context.Context, p p2p.Peer, stream io.ReadWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// Probe implements SyncBase.
func (ssb *SetSyncBase) Probe(ctx context.Context, p p2p.Peer) (pr rangesync.ProbeResult, err error) {
	_ = "STUB: not implemented"
	// Use a snapshot of the store to avoid holding the mutex for a long time
	return *new(rangesync.ProbeResult), nil
}
