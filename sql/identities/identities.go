package identities

import (
	"context"
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

// SetMalicious records identity as malicious.
func SetMalicious(db sql.Executor, nodeID types.NodeID, proof []byte, received time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// IsMalicious returns true if identity is known to be malicious.
func IsMalicious(db sql.Executor, nodeID types.NodeID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetBlobSizes returns the sizes of the blobs corresponding to malfeasance proofs for the
// specified identities. For non-existent proofs, the corresponding items are set to -1.
func GetBlobSizes(db sql.Executor, ids [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadMalfeasanceBlob returns the malfeasance proof in raw bytes for the given identity.
func LoadMalfeasanceBlob(_ context.Context, db sql.Executor, nodeID []byte, blob *sql.Blob) error {
	_ = "STUB: not implemented"
	return nil
}

func IterateOps(
	db sql.Executor,
	operations builder.Operations,
	fn func(types.NodeID, []byte, time.Time) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// AllMalicious retrieves malicious node IDs from the database.
func AllMalicious(db sql.Executor) ([]types.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CountMalicious returns the number of malicious nodes.
func CountMalicious(db sql.Executor) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
