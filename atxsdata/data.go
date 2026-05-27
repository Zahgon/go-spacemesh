package atxsdata

import (
	"iter"
	"sync"
	"sync/atomic"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
)

// SAFETY: all exported fields are read-only and are safe to read concurrently.
// Thanks to the fact that ATX is immutable, it is safe to return a pointer to it.
type ATX struct {
	Node               types.NodeID
	Coinbase           types.Address
	Weight             uint64
	BaseHeight, Height uint64
	Nonce              types.VRFPostIndex
}

func New() *Data { _ = "STUB: not implemented"; return nil }

type Data struct {
	evicted atomic.Uint32

	mu        sync.RWMutex
	malicious map[types.NodeID]struct{}
	epochs    map[types.EpochID]map[types.ATXID]*ATX

	signers map[types.NodeID]struct{}
	managed map[types.EpochID]map[types.NodeID]types.ATXID
}

func (d *Data) Register(sig *signing.EdSigner) { _ = "STUB: not implemented"; return }

// update quick access for newly registered signer

func (d *Data) Evicted() types.EpochID { _ = "STUB: not implemented"; return *new(types.EpochID) }

func (d *Data) IsEvicted(epoch types.EpochID) bool { _ = "STUB: not implemented"; return false }

// EvictEpoch is a notification for cache to evict epochs that are not useful
// to keep in memory.
func (d *Data) EvictEpoch(evict types.EpochID) { _ = "STUB: not implemented"; return }

// AddFromHeader extracts relevant fields from an ActivationTx and adds them together with nonce and malicious flag.
// Returns the ATX that was added to the store (if any) or `nil` if it wasn't.
func (d *Data) AddFromAtx(atx *types.ActivationTx, malicious bool) *ATX {
	_ = "STUB: not implemented"
	return nil
}

// Add adds ATX data to the store.
// Returns whether the ATX was added to the store.
func (d *Data) AddAtx(target types.EpochID, id types.ATXID, atx *ATX) bool {
	_ = "STUB: not implemented"
	return false
}

// Add adds ATX data to the store.
// Returns the ATX that was added to the store (if any) or `nil` if it wasn't.
func (d *Data) Add(
	epoch types.EpochID,
	node types.NodeID,
	coinbase types.Address,
	atxid types.ATXID,
	weight, baseHeight, height uint64,
	nonce types.VRFPostIndex,
	malicious bool,
) *ATX {
	_ = "STUB: not implemented"
	return nil
}

func (d *Data) IsMalicious(node types.NodeID) bool { _ = "STUB: not implemented"; return false }

func (d *Data) SetMalicious(node types.NodeID) { _ = "STUB: not implemented"; return }

func (d *Data) MaliciousIdentities() iter.Seq[types.NodeID] { _ = "STUB: not implemented"; return nil }

// Get returns atx data.
// SAFETY: The returned pointer MUST NOT be modified.
func (d *Data) Get(epoch types.EpochID, atx types.ATXID) *ATX {
	_ = "STUB: not implemented"
	return nil
}

// GetByEpochAndNodeID returns atx data by epoch and node id. This query will be slow for nodeIDs that
// are not managed by the node.
func (d *Data) GetByEpochAndNodeID(epoch types.EpochID, node types.NodeID) (types.ATXID, *ATX) {
	_ = "STUB: not implemented"
	return *new(types.ATXID), nil
}

func (d *Data) Size(target types.EpochID) int { _ = "STUB: not implemented"; return 0 }

type lockGuard struct{}

// AtxFilter is a function that filters atxs.
// The `lockGuard` prevents using the filter functions outside of the allowed context
// to prevent data races.
type AtxFilter func(*Data, *ATX, lockGuard) bool

func NotMalicious(d *Data, atx *ATX, _ lockGuard) bool { _ = "STUB: not implemented"; return false }

// IterateInEpoch calls `fn` for every ATX in epoch.
// If filters are provided, only atxs that pass all filters are returned.
// SAFETY: The returned pointer MUST NOT be modified.
func (d *Data) IterateInEpoch(epoch types.EpochID, fn func(types.ATXID, *ATX), filters ...AtxFilter) {
	_ = "STUB: not implemented"
	return
}

func (d *Data) iterateInEpoch(epoch types.EpochID, fn func(types.ATXID, *ATX), filters ...AtxFilter) {
	_ = "STUB: not implemented"
	return
}

func (d *Data) IterateHighTicksInEpoch(target types.EpochID, fn func(types.ATXID) bool) {
	_ = "STUB: not implemented"
	return
}

// FindHighestHonest looks for a heightest ATX in the most recent 2 epochs.
func (d *Data) FindHighestHonest() types.ATXID { _ = "STUB: not implemented"; return *new(types.ATXID) }

func (d *Data) MissingInEpoch(epoch types.EpochID, atxs []types.ATXID) []types.ATXID {
	_ = "STUB: not implemented"
	return nil
}

// WeightForSet computes total weight of atxs in the set and returned array with
// atxs in the set that weren't used.
func (d *Data) WeightForSet(epoch types.EpochID, set []types.ATXID) (uint64, []bool) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO(dshulyak) bitfield is a perfect fit here
