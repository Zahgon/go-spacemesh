package multipeer

import (
	"context"
	"time"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/fetch/peers"
	"github.com/spacemeshos/go-spacemesh/p2p"
)

type syncResult struct {
	peer p2p.Peer
	err  error
}

// splitSync is a synchronization implementation that synchronizes the set against
// multiple peers in parallel, but splits the synchronization into ranges and assigns
// each range to a single peer.
// The splitting is done in such way that only maxDepth high bits of the key are non-zero,
// which helps with radix tree based OrderedSet implementation.
type splitSync struct {
	logger       *zap.Logger
	syncBase     SyncBase
	peers        *peers.Peers
	syncPeers    []p2p.Peer
	gracePeriod  time.Duration
	clock        clockwork.Clock
	sq           syncQueue
	resCh        chan syncResult
	slowRangeCh  chan *syncRange
	syncMap      map[p2p.Peer]*syncRange
	failedPeers  map[p2p.Peer]struct{}
	numRunning   int
	numRemaining int
	numPeers     int
	eg           *errgroup.Group
}

func newSplitSync(
	logger *zap.Logger,
	syncBase SyncBase,
	peers *peers.Peers,
	syncPeers []p2p.Peer,
	gracePeriod time.Duration,
	clock clockwork.Clock,
	keyLen, maxDepth int,
) *splitSync {
	_ = "STUB: not implemented"
	return nil
}

func (s *splitSync) nextPeer() p2p.Peer { _ = "STUB: not implemented"; return *new(p2p.Peer) }

func (s *splitSync) startPeerSync(ctx context.Context, p p2p.Peer, sr *syncRange) error {
	_ = "STUB: not implemented"
	return nil
}

// if another peer finishes its part early, let
// it pick up this range

func (s *splitSync) handleSyncResult(r syncResult) error { _ = "STUB: not implemented"; return nil }

// prioritize the syncRange for resync after failed
// sync with no active syncs remaining

func (s *splitSync) clearDeadPeers() { _ = "STUB: not implemented"; return }

func (s *splitSync) Sync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Push this syncRange to the back of the queue.
// There's some chance that the peer managed to complete
// the sync while the range was still sitting in the
// channel, so we double-check if it's done.

// Stop late peers that didn't manage to sync their ranges in time.
// The ranges were already reassigned to other peers and successfully
// synced by this point.

// If all the ranges are synced, the split sync is considered successful
// even if some peers failed to sync their ranges, so that these ranges
// got synced by other peers.
