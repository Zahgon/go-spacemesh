package sqlstore

import (
	"testing"

	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// CreateDB creates a test database. It is only used for testing.
func CreateDB(t *testing.T, keyLen int) sql.Database {
	_ = "STUB: not implemented"
	return *new(sql.Database)
}

func insertDBItems(t *testing.T, db sql.Database, content []rangesync.KeyBytes, cmd string) {
	_ = "STUB: not implemented"
	return
}

// InsertDBItems inserts items into a test database. It is only used for testing.
func InsertDBItems(t *testing.T, db sql.Database, content []rangesync.KeyBytes) {
	_ = "STUB: not implemented"
	return
}

// EnsureDBItems inserts items into a test database, skipping rows with keys that already exist.
// It is only used for testing.
func EnsureDBItems(t *testing.T, db sql.Database, content []rangesync.KeyBytes) {
	_ = "STUB: not implemented"
	return
}

// PopulateDB creates a test database and inserts items into it. It is only used for testing.
func PopulateDB(t *testing.T, keyLen int, content []rangesync.KeyBytes) sql.Database {
	_ = "STUB: not implemented"
	return *new(sql.Database)
}
