package sqlstore

import (
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// max chunk size to use for dbSeq.
const sqlMaxChunkSize = 1024

// SQLIDStore is an implementation of IDStore that is based on a database table snapshot.
type SQLIDStore struct {
	db     sql.Executor
	sts    *SyncedTableSnapshot
	keyLen int
}

var _ IDStore = &SQLIDStore{}

// NewSQLIDStore creates a new SQLIDStore.
func NewSQLIDStore(db sql.Executor, sts *SyncedTableSnapshot, keyLen int) *SQLIDStore {
	_ = "STUB: not implemented"
	return nil
}

// Clone creates a new SQLIDStore that shares the same database connection and table snapshot.
// Implements IDStore.
func (s *SQLIDStore) Clone() IDStore { _ = "STUB: not implemented"; return *new(IDStore) }

// RegisterKey is a no-op for SQLIDStore, as the underlying table is never immediately
// updated upon receiving new items.
// Implements IDStore.
func (s *SQLIDStore) RegisterKey(k rangesync.KeyBytes) error {
	_ = "STUB: not implemented"
	// should be registered by the handler code
	return nil
}

// All returns all IDs in the store.
// The sequence in SeqResult returned by All is either empty or infinite.
// Implements IDStore.
func (s *SQLIDStore) All() rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// From returns IDs in the store starting from the given key.
// The sequence in SeqResult returned by From is either empty or infinite.
// Implements IDStore.
func (s *SQLIDStore) From(from rangesync.KeyBytes, sizeHint int) rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// Since returns IDs in the store starting from the given key and timestamp.
// The sequence in SeqResult returned by Since is either empty or infinite.
func (s *SQLIDStore) Since(from rangesync.KeyBytes, since int64) (rangesync.SeqResult, int) {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult), 0
}

// Sets the table snapshot to use for the store.
func (s *SQLIDStore) SetSnapshot(sts *SyncedTableSnapshot) {
	_ = "STUB: not implemented"

	// Release is a no-op for SQLIDStore.
	// Implements IDStore.
	return
}

func (s *SQLIDStore) Release() { _ = "STUB: not implemented"; return }
