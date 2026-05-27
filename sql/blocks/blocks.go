package blocks

import (
	"context"
	"errors"
	"io"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	valid   = 1
	invalid = -1
)

var ErrValidityNotDecided = errors.New("block validity undecided")

func decodeBlock(reader io.Reader, id types.BlockID) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add block to the database.
func Add(db sql.Executor, block *types.Block) error { _ = "STUB: not implemented"; return nil }

// this is actually should encode block

// Has a block in the database.
func Has(db sql.Executor, id types.BlockID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetBlobSizes returns the sizes of the blobs corresponding to blocks with specified
// ids. For non-existent ballots, the corresponding items are set to -1.
func GetBlobSizes(db sql.Executor, ids [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadBlob loads block as an encoded blob, ready to be sent over the wire.
func LoadBlob(ctx context.Context, db sql.Executor, id []byte, b *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

// Get block with id from database.
func Get(db sql.Executor, id types.BlockID) (rst *types.Block, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LastValid(db sql.Executor) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *

	// it doesn't use max(layer) in order to get rows == 0 when there are no layers.
	// aggregation always returns rows == 1, hence the check below doesn't work.
	new(types.LayerID), nil
}

func UpdateValid(db sql.Executor, id types.BlockID, valid bool) error {
	_ = "STUB: not implemented"
	return nil
}

// SetValid updates verified status for a block.
func SetValid(db sql.Executor, id types.BlockID) error { _ = "STUB: not implemented"; return nil }

// SetInvalid updates blocks to an invalid status.
func SetInvalid(db sql.Executor, id types.BlockID) error { _ = "STUB: not implemented"; return nil }

func setValidity(db sql.Executor, id types.BlockID, validity int8) error {
	_ = "STUB: not implemented"
	return nil
}

// IsValid returns true if block is verified.
func IsValid(db sql.Executor, id types.BlockID) (rst bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetLayer returns the layer of a block.
func GetLayer(db sql.Executor, id types.BlockID) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}

// Layer returns full body blocks for layer.
func Layer(db sql.Executor, lid types.LayerID) ([]*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IDsInLayer returns list of block ids in the layer.
func IDsInLayer(db sql.Executor, lid types.LayerID) ([]types.BlockID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ContextualValidity returns tuples with block id and contextual validity for all blocks in the layer.
func ContextualValidity(db sql.Executor, lid types.LayerID) ([]types.BlockContextualValidity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
