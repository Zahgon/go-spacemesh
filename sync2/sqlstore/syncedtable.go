package sqlstore

import (
	"errors"
	"sync"

	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/expr"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// Binder is a function that binds filter expression parameters to a SQL statement.
type Binder func(s *sql.Statement)

// SyncedTable represents a table that can be used with SQLIDStore.
type SyncedTable struct {
	// The name of the table.
	TableName string
	// The name of the ID column.
	IDColumn string
	// The name of the timestamp column.
	TimestampColumn string
	// The filter expression.
	Filter expr.Expr
	// The binder function for the bind parameters appearing in the filter expression.
	Binder  Binder
	mtx     sync.Mutex
	queries map[string]string
}

func (st *SyncedTable) cacheQuery(name string, gen func() expr.Statement) string {
	_ = "STUB: not implemented"
	return ""
}

func (st *SyncedTable) exec(
	db sql.Executor,
	name string,
	gen func() expr.Statement,
	enc sql.Encoder,
	dec sql.Decoder,
) error {
	_ = "STUB: not implemented"
	return nil
}

// genSelectMaxRowID generates a SELECT statement that returns the maximum rowid in the
// table.
func (st *SyncedTable) genSelectMaxRowID() expr.Statement {
	_ = "STUB: not implemented"
	return *new(expr.Statement)
}

// rowIDCutoff returns an expression that represents a rowid cutoff, that is, limits the
// rowid to be less than or equal to a bind parameter.
func (st *SyncedTable) rowIDCutoff() expr.Expr { _ = "STUB: not implemented"; return *new(expr.Expr) }

// timestampCutoff returns an expression that represents a timestamp cutoff, that is, limits the
// timestamp to be greater than or equal to a bind parameter.
func (st *SyncedTable) timestampCutoff() expr.Expr {
	_ = "STUB: not implemented"
	return *new(expr.Expr)
}

// genSelectAll generates a SELECT statement that returns all the rows in the table
// satisfying the filter expression and the rowid cutoff.
func (st *SyncedTable) genSelectAll() expr.Statement {
	_ = "STUB: not implemented"
	return *new(expr.Statement)
}

// genCount generates a SELECT statement that returns the number of rows in the table
// satisfying the filter expression and the rowid cutoff.
func (st *SyncedTable) genCount() expr.Statement {
	_ = "STUB: not implemented"
	return *new(expr.Statement)
}

// genSelectAllSinceSnapshot generates a SELECT statement that returns all the rows in the
// table satisfying the filter expression that have rowid between the specified min and
// max parameter values, inclusive.
func (st *SyncedTable) genSelectAllSinceSnapshot() expr.Statement {
	_ = "STUB: not implemented"
	return *new(expr.Statement)
}

// genSelectRange generates a SELECT statement that returns the rows in the table
// satisfying the filter expression, the rowid cutoff and the specified ID range.
func (st *SyncedTable) genSelectRange() expr.Statement {
	_ = "STUB: not implemented"
	return *new(expr.Statement)
}

// genRecentCount generates a SELECT statement that returns the number of rows in the table
// added starting with the specified timestamp, taking into account the filter expression
// and the rowid cutoff.
func (st *SyncedTable) genRecentCount() expr.Statement {
	_ = "STUB: not implemented"
	return *new(expr.Statement)
}

// genRecentCount generates a SELECT statement that returns the rows in the table added
// starting with the specified timestamp, taking into account the filter expression and
// the rowid cutoff.
func (st *SyncedTable) genSelectRecent() expr.Statement {
	_ = "STUB: not implemented"
	return *new(expr.Statement)
}

// loadMaxRowID returns the max rowid in the table.
func (st *SyncedTable) loadMaxRowID(db sql.Executor) (maxRowID int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Snapshot creates a snapshot of the table based on its current max rowid value.
func (st *SyncedTable) Snapshot(db sql.Executor) (*SyncedTableSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SyncedTableSnapshot represents a snapshot of an append-only table.
// The snapshotting is relies on rowids of the table rows never decreasing
// as new rows are added.
// Each snapshot inherits filter expression from the table, so all the rows relevant to
// the snapshot are always filtered using that expression, if it's specified.
type SyncedTableSnapshot struct {
	*SyncedTable
	maxRowID int64
}

// Load loads all the table rows belonging to a snapshot.
func (sts *SyncedTableSnapshot) Load(
	db sql.Executor,
	dec func(stmt *sql.Statement) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadCount returns the number of rows in the snapshot.
func (sts *SyncedTableSnapshot) LoadCount(
	db sql.Executor,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// LoadSinceSnapshot loads rows added since the specified previous snapshot.
func (sts *SyncedTableSnapshot) LoadSinceSnapshot(
	db sql.Executor,
	prev *SyncedTableSnapshot,
	dec func(stmt *sql.Statement) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadRange loads ids starting from the specified one.
// limit specifies the maximum number of ids to load.
func (sts *SyncedTableSnapshot) LoadRange(
	db sql.Executor,
	fromID rangesync.KeyBytes,
	limit int,
	dec func(stmt *sql.Statement) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

var errNoTimestampColumn = errors.New("no timestamp column")

// LoadRecentCount returns the number of rows added since the specified timestamp.
func (sts *SyncedTableSnapshot) LoadRecentCount(
	db sql.Executor,
	since int64,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// LoadRecent loads rows added since the specified timestamp.
func (sts *SyncedTableSnapshot) LoadRecent(
	db sql.Executor,
	fromID rangesync.KeyBytes,
	limit int,
	since int64,
	dec func(stmt *sql.Statement) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}
