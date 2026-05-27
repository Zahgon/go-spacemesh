package checkpoint

import (
	"context"

	"github.com/spf13/afero"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	SchemaVersion = "https://spacemesh.io/checkpoint.schema.json.1.0"

	CommandString = "grpcurl -plaintext -d '%s' 0.0.0.0:9093 spacemesh.v1.AdminService.CheckpointStream"

	checkpointDir = "checkpoint"
	schemaFile    = "schema.json"
	dirPerm       = 0o700
)

func checkpointDB(
	ctx context.Context,
	db sql.StateDatabase,
	snapshot types.LayerID,
	numAtxs int,
) (*types.Checkpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// collect marriage ATXs

func Generate(
	ctx context.Context,
	fs afero.Fs,
	db sql.StateDatabase,
	dataDir string,
	snapshot types.LayerID,
	numAtxs int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// one writer persist the checkpoint data, one returning result to caller.

func SelfCheckpointFilename(dataDir string, snapshot types.LayerID) string {
	_ = "STUB: not implemented"
	return ""
}
