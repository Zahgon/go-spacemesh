package checkpoint

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net/url"

	"github.com/spf13/afero"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

var (
	ErrCheckpointNotFound    = errors.New("checkpoint not found")
	ErrUrlSchemeNotSupported = errors.New("url scheme not supported")
)

type RecoveryFile struct {
	file    afero.File
	fwriter *bufio.Writer
	path    string
}

func NewRecoveryFile(aferoFs afero.Fs, path string) (*RecoveryFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rf *RecoveryFile) Copy(fs afero.Fs, src io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (rf *RecoveryFile) Save(fs afero.Fs) error { _ = "STUB: not implemented"; return nil }

func ValidateSchema(data []byte) error { _ = "STUB: not implemented"; return nil }

func CopyFile(fs afero.Fs, src, dst string) error { _ = "STUB: not implemented"; return nil }

func httpToLocalFile(ctx context.Context, resource *url.URL, fs afero.Fs, dst string) error {
	_ = "STUB: not implemented"
	return nil
}

func backupRecovery(fs afero.Fs, recoveryDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func backupOldDb(fs afero.Fs, srcDir, dbFile string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// sqlite create .sql, .sql-shm and .sql-wal files.

func poetProofRefs(ctx context.Context, db sql.Executor, id types.ATXID) ([]types.PoetProofRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
