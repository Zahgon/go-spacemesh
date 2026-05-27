package syncer

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/fetch"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/sql"
)

var (
	ErrPeerMeshChangedMidSession = errors.New("peer mesh changed mid session")
	ErrNodeMeshChangedMidSession = errors.New("node mesh changed mid session")
)

type layerHash struct {
	layer   types.LayerID
	hash    types.Hash32
	created time.Time
}

// boundary is used to define the to and from layers in the hash mesh requests to peers.
// The hashes of the boundary layers are known to the node and are used to double-check that
// the peer has not changed its opinions on those layers.
// If the boundary hashes change during a fork-finding session, the session is aborted.
type boundary struct {
	from, to *layerHash
}

func (b *boundary) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ForkFinder struct {
	logger           *zap.Logger
	db               sql.Executor
	fetcher          fetcher
	maxStaleDuration time.Duration

	mu          sync.Mutex
	agreedPeers map[p2p.Peer]*layerHash
	// used to make sure we only resync based on the same layer hash once across runs.
	resynced map[types.LayerID]map[types.Hash32]time.Time
}

func NewForkFinder(lg *zap.Logger, db sql.Executor, f fetcher, maxStale time.Duration) *ForkFinder {
	_ = "STUB: not implemented"
	return nil
}

// Purge cached agreements with peers.
func (ff *ForkFinder) Purge(all bool, toPurge ...p2p.Peer) { _ = "STUB: not implemented"; return }

// NumPeersCached returns the number of peer agreement cached.
func (ff *ForkFinder) NumPeersCached() int { _ = "STUB: not implemented"; return 0 }

func (ff *ForkFinder) AddResynced(lid types.LayerID, hash types.Hash32) {
	_ = "STUB: not implemented"
	return
}

func (ff *ForkFinder) NeedResync(lid types.LayerID, hash types.Hash32) bool {
	_ = "STUB: not implemented"
	return false
}

// FindFork finds the point of divergence in layer opinions between the node and the specified peer
// from a given disagreed layer.
func (ff *ForkFinder) FindFork(
	ctx context.Context,
	peer p2p.Peer,
	diffLid types.LayerID,
	diffHash types.Hash32,
) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}

// every layer hash is different/same from node's. this can only happen when
// the node's local hashes change while the mesh hash request is running

// UpdateAgreement updates the layer at which the peer agreed with the node.
func (ff *ForkFinder) UpdateAgreement(
	peer p2p.Peer,
	lid types.LayerID,
	hash types.Hash32,
	created time.Time,
) {
	_ = "STUB: not implemented"
	return
}

func (ff *ForkFinder) updateAgreement(peer p2p.Peer, update *layerHash, created time.Time) {
	_ = "STUB: not implemented"
	return
}

// unconditional update instead of comparing layers because peers can change its opinions on historical layers.

// setupBoundary sets up the boundary for the hash requests.
// - boundary.from contains the latest layer node and peer agree on hash.
// - boundary.to contains the oldest layer node and peer disagree on hash.
func (ff *ForkFinder) setupBoundary(peer p2p.Peer, oldestDiff *layerHash) (*boundary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// double check if the node still has the same hash

// Get layer hashes from peers in the range defined by the boundary.
// If the number of hashes is less than maxHashesInReq, then request every hash.
// Otherwise, set appropriate params such that the number of hashes requested is maxHashesInReq
// while ensuring hashes for the boundary layers are requested.
func (ff *ForkFinder) sendRequest(
	ctx context.Context,
	logger *zap.Logger,
	peer p2p.Peer,
	bnd *boundary,
) (*fetch.MeshHashes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
