package layers

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

// SetWeakCoin for the layer.
func SetWeakCoin(db sql.Executor, lid types.LayerID, weakcoin bool) error {
	_ = "STUB: not implemented"
	return nil
}

// GetWeakCoin for layer.
func GetWeakCoin(db sql.Executor, lid types.LayerID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SetApplied for the layer to a block id.
func SetApplied(db sql.Executor, lid types.LayerID, applied types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

// UnsetAppliedFrom updates the applied block to nil for layer >= `lid`.
func UnsetAppliedFrom(db sql.Executor, lid types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateStateHash for the layer.
func UpdateStateHash(db sql.Executor, lid types.LayerID, hash types.Hash32) error {
	_ = "STUB: not implemented"
	return nil
}

// GetLatestStateHash loads latest state hash.
func GetLatestStateHash(db sql.Executor) (rst types.Hash32, err error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), nil
}

// GetStateHash loads state hash for the layer.
func GetStateHash(db sql.Executor, lid types.LayerID) (rst types.Hash32, err error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), nil
}

// GetApplied for the applied block for layer.
func GetApplied(db sql.Executor, lid types.LayerID) (rst types.BlockID, err error) {
	_ = "STUB: not implemented"
	return *new(types.BlockID), nil
}

func FirstAppliedInEpoch(db sql.Executor, epoch types.EpochID) (types.BlockID, error) {
	_ = "STUB: not implemented"
	return *new(types.BlockID), nil
}

// GetLastApplied for the applied block for layer.
func GetLastApplied(db sql.Executor) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}

// SetProcessed sets a layer processed.
func SetProcessed(db sql.Executor, lid types.LayerID) error { _ = "STUB: not implemented"; return nil }

// GetProcessed gets the highest layer processed.
func GetProcessed(db sql.Executor) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}

// SetMeshHash sets the aggregated hash up to the specified layer.
func SetMeshHash(db sql.Executor, lid types.LayerID, aggHash types.Hash32) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAggregatedHash for layer.
func GetAggregatedHash(db sql.Executor, lid types.LayerID) (types.Hash32, error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), nil
}

func IterateAggHashes(
	db sql.Executor,
	from, to types.LayerID,
	by uint32,
	callback func(total int, id types.Hash32) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// last layer is not a multiple of By, so we need to add it

func GetAggHashes(db sql.Executor, from, to types.LayerID, by uint32) (hashes []types.Hash32, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Layer struct {
	Id             types.LayerID
	WeakCoin       bool
	Processed      bool
	AppliedBlock   types.BlockID
	StateHash      types.Hash32
	AggregatedHash types.Hash32
	Block          *types.Block
}

func IterateLayersWithBlockOps(
	db sql.Executor,
	operations builder.Operations,
	fn func(layer *Layer) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func Get(
	db sql.Executor,
	lid types.LayerID,
) (*Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
