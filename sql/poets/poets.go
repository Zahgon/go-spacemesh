package poets

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

// Has checks if a PoET exists by the given ref.
func Has(db sql.Executor, ref types.PoetProofRef) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetBlobSizes returns the sizes of the blobs corresponding to PoETs with specified
// refs. For non-existent PoETs, the corresponding items are set to -1.
func GetBlobSizes(db sql.Executor, refs [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadBlob loads PoET as an encoded blob, ready to be sent over the wire.
func LoadBlob(ctx context.Context, db sql.Executor, ref []byte, blob *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

// Get gets a PoET for a given ref.
func Get(db sql.Executor, ref types.PoetProofRef) (poet []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds a poet for a given ref.
func Add(db sql.Executor, ref types.PoetProofRef, poet, serviceID []byte, roundID string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetRef gets a PoET ref for a given service ID and round ID.
func GetRef(db sql.Executor, poetID []byte, roundID string) (ref types.PoetProofRef, err error) {
	_ = "STUB: not implemented"
	return *new(types.PoetProofRef), nil
}
