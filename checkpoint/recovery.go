package checkpoint

import (
	"context"

	"github.com/spf13/afero"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/atxs"
	"github.com/spacemeshos/go-spacemesh/sql/marriage"
)

const recoveryDir = "recovery"

type Config struct {
	Uri     string `mapstructure:"recovery-uri"`
	Restore uint32 `mapstructure:"recovery-layer"`

	// set to false if atxs are not compatible before and after the checkpoint recovery.
	PreserveOwnAtx bool `mapstructure:"preserve-own-atx"`
}

func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type RecoverConfig struct {
	GoldenAtx   types.ATXID
	DataDir     string
	DbFile      string
	LocalDbFile string
	NodeIDs     []types.NodeID // IDs to preserve own ATXs
	Uri         string
	Restore     types.LayerID
}

func (c *RecoverConfig) DbPath() string { _ = "STUB: not implemented"; return "" }

func RecoveryDir(dataDir string) string { _ = "STUB: not implemented"; return "" }

func RecoveryFilename(dataDir, base string, restore types.LayerID) string {
	_ = "STUB: not implemented"
	return ""
}

func copyToLocalFile(
	ctx context.Context,
	logger *zap.Logger,
	fs afero.Fs,
	dataDir, uri string,
	restore types.LayerID,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type AtxDep struct {
	ID           types.ATXID
	PublishEpoch types.EpochID
	Blob         []byte
}

type PreservedData struct {
	Deps   []*AtxDep
	Proofs []*types.PoetProofMessage
}

func Recover(
	ctx context.Context,
	logger *zap.Logger,
	fs afero.Fs,
	cfg *RecoverConfig,
) (*PreservedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RecoverWithDb(
	ctx context.Context,
	logger *zap.Logger,
	db sql.StateDatabase,
	localDB sql.LocalDatabase,
	fs afero.Fs,
	cfg *RecoverConfig,
) (*PreservedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type recoveryData struct {
	accounts  []*types.Account
	atxs      []*atxs.CheckpointAtx
	marriages []marriage.Info
}

func RecoverFromLocalFile(
	ctx context.Context,
	logger *zap.Logger,
	db sql.StateDatabase,
	localDB sql.LocalDatabase,
	fs afero.Fs,
	cfg *RecoverConfig,
	file string,
) (*PreservedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// continue to recover from checkpoint despite failure to preserve own atx

// sort ATXs them by publishEpoch and then by ID

// all is ready. backup the old data and create new.

func checkpointData(fs afero.Fs, file string, newGenesis types.LayerID) (*recoveryData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectOwnAtxDeps(
	logger *zap.Logger,
	db sql.StateDatabase,
	localDB sql.LocalDatabase,
	nodeID types.NodeID,
	goldenATX types.ATXID,
	data *recoveryData,
) (map[types.ATXID]*AtxDep, map[types.PoetProofRef]*types.PoetProofMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// check for if smesher is building any atx

// any previous atx in nipost should already be captured earlier
// we only care about positioning atx here

func collectDeps(
	db sql.StateDatabase,
	ref types.ATXID,
	all map[types.ATXID]struct{},
) (map[types.ATXID]*AtxDep, map[types.PoetProofRef]*types.PoetProofMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func collect(
	db sql.StateDatabase,
	ref types.ATXID,
	all map[types.ATXID]struct{},
	deps map[types.ATXID]*AtxDep,
) error {
	_ = "STUB: not implemented"
	return nil
}

func poetProofs(
	db sql.StateDatabase,
	atxIds map[types.ATXID]*AtxDep,
) (map[types.PoetProofRef]*types.PoetProofMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
