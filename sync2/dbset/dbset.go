package dbset

import (
	"sync"
	"time"

	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sync2/fptree"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
	"github.com/spacemeshos/go-spacemesh/sync2/sqlstore"
)

// DBSet is an implementation of rangesync.OrderedSet that uses an SQL database
// as its backing store. It uses an FPTree to perform efficient range queries.
type DBSet struct {
	loadMtx  sync.Mutex
	db       sql.Executor
	ft       *fptree.FPTree
	st       *sqlstore.SyncedTable
	snapshot *sqlstore.SyncedTableSnapshot
	dbStore  *fptree.DBBackedStore
	keyLen   int
	maxDepth int
	received map[string]struct{}
}

var _ rangesync.OrderedSet = &DBSet{}

// NewDBSet creates a new DBSet.
func NewDBSet(
	db sql.Executor,
	st *sqlstore.SyncedTable,
	keyLen, maxDepth int,
) *DBSet {
	_ = "STUB: not implemented"
	return nil
}

func (d *DBSet) handleIDfromDB(stmt *sql.Statement) bool { _ = "STUB: not implemented"; return false }

// Loaded returns true if the DBSet is loaded.
// Implements rangesync.OrderedSet.
func (d *DBSet) Loaded() bool { _ = "STUB: not implemented"; return false }

// EnsureLoaded ensures that the DBSet is loaded and ready to be used.
// Implements rangesync.OrderedSet.
func (d *DBSet) EnsureLoaded() error { _ = "STUB: not implemented"; return nil }

// Received returns a sequence of all items that have been received.
// Implements rangesync.OrderedSet.
func (d *DBSet) Received() rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// Add adds an item to the DBSet.
// Implements rangesync.OrderedSet.
func (d *DBSet) Add(k rangesync.KeyBytes) error { _ = "STUB: not implemented"; return nil }

// Receive handles a newly received item, arranging for it to be returned as part of the
// sequence returned by Received.
// Implements rangesync.OrderedSet.
func (d *DBSet) Receive(k rangesync.KeyBytes) error { _ = "STUB: not implemented"; return nil }

// RangeInfo returns information about the range of items in the DBSet.
// Implements rangesync.OrderedSet.
func (d *DBSet) RangeInfo(x, y rangesync.KeyBytes) (rangesync.RangeInfo, error) {
	_ = "STUB: not implemented"
	return *new(rangesync.RangeInfo), nil
}

// SplitRange splits the range of items in the DBSet into two parts,
// returning information about eachn part and the middle item.
// Implements rangesync.OrderedSet.
func (d *DBSet) SplitRange(x, y rangesync.KeyBytes, count int) (rangesync.SplitInfo, error) {
	_ = "STUB: not implemented"
	return *new(rangesync.SplitInfo), nil
}

// SetInfo returns RangeInfo for the whole DBSet.
// Implements rangesync.OrderedSet.
func (d *DBSet) SetInfo() (rangesync.RangeInfo, error) {
	_ = "STUB: not implemented"
	return *new(rangesync.RangeInfo), nil
}

// Items returns a sequence of all items in the DBSet.
// Implements rangesync.OrderedSet.
func (d *DBSet) Items() rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// Empty returns true if the DBSet is empty.
// Implements rangesync.OrderedSet.
func (d *DBSet) Empty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Advance advances the DBSet to the latest state of the underlying database table.
func (d *DBSet) Advance() error { _ = "STUB: not implemented"; return nil }

// WithCopy invokes the specified function, passing it a temporary copy of the DBSet.
// Implements rangesync.OrderedSet.
func (d *DBSet) WithCopy(toCall func(rangesync.OrderedSet) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Has returns true if the DBSet contains the given item.
// Implements rangesync.OrderedSet.
func (d *DBSet) Has(k rangesync.KeyBytes) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// checkKey may have false positives, but not false negatives, and it's much
// faster than querying the database

// Recent returns a sequence of items that have been added to the DBSet since the given time.
// Implements rangesync.OrderedSet.
func (d *DBSet) Recent(since time.Time) (rangesync.SeqResult, int) {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult), 0
}

func (d *DBSet) release() { _ = "STUB: not implemented"; return }
