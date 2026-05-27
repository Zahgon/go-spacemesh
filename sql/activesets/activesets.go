package activesets

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	CacheKindActiveSetBlob sql.QueryCacheKind = "activeset-blob"
)

func Add(db sql.Executor, id types.Hash32, set *types.EpochActiveSet) error {
	_ = "STUB: not implemented"
	return nil
}

func Get(db sql.Executor, id types.Hash32) (*types.EpochActiveSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlobSizes returns the sizes of the blobs corresponding to activesets with specified
// ids. For non-existent activesets, the corresponding items are set to -1.
func GetBlobSizes(db sql.Executor, ids [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadBlob loads activesets as an encoded blob, ready to be sent over the wire.
func LoadBlob(ctx context.Context, db sql.Executor, id []byte, blob *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

func getBlob(ctx context.Context, db sql.Executor, id []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteBeforeEpoch(db sql.Executor, epoch types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

func Has(db sql.Executor, id types.Hash32) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
