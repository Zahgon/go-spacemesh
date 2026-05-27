package recovery

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

// SetCheckpoint for the node.
func SetCheckpoint(db sql.Executor, restore types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckpointInfo for the node.
func CheckpointInfo(db sql.Executor) (types.LayerID, error) {
	_ = "STUB: not implemented"
	return *new(types.LayerID), nil
}
