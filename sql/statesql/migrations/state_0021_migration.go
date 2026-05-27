package migrations

import (
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type migration0021 struct {
	batch int
}

var _ sql.Migration = &migration0021{}

func New0021Migration(batch int) *migration0021 { _ = "STUB: not implemented"; return nil }

func (*migration0021) Name() string { _ = "STUB: not implemented"; return "" }

func (*migration0021) Order() int { _ = "STUB: not implemented"; return 0 }

func (*migration0021) Rollback() error { _ = "STUB: not implemented"; return nil }

func (m *migration0021) Apply(db sql.Executor, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

const tableSQL = `CREATE TABLE posts (
    atxid CHAR(32) NOT NULL,
    pubkey CHAR(32) NOT NULL,
    prev_atxid CHAR(32),
    prev_atx_index INT,
    units INT NOT NULL
);`

func (m *migration0021) applySql(db sql.Executor) error { _ = "STUB: not implemented"; return nil }

func (m *migration0021) applyIndices(db sql.Executor) error { _ = "STUB: not implemented"; return nil }

type update struct {
	id    types.NodeID
	prev  types.ATXID
	units uint32
}

func (m *migration0021) processBatch(db sql.Executor, offset, size int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *migration0021) applyPendingUpdates(db sql.Executor, updates map[types.ATXID]*update) error {
	_ = "STUB: not implemented"
	return nil
}

func processATX(blob types.AtxBlob) (*update, error) {
	_ = "STUB: not implemented"
	// The migration adding the version column does not set it to 1 for existing ATXs.
	// Thus, both values 0 and 1 mean V1.
	return nil, nil
}

func setPost(db sql.Executor, atxID, prev types.ATXID, prevIndex int, id types.NodeID, units uint32) error {
	_ = "STUB: not implemented"
	return nil
}
