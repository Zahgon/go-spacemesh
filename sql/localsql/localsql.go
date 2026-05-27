package localsql

import (
	"embed"
	"testing"

	"github.com/spacemeshos/go-spacemesh/sql"
)

//go:generate go run ../schemagen -dbtype local -output schema/schema.sql

//go:embed schema/schema.sql
var schemaScript string

//go:embed schema/migrations/*.sql
var migrations embed.FS

type database struct {
	sql.Database
}

var _ sql.LocalDatabase = &database{}

func (d *database) IsLocalDatabase() {
	_ = "STUB: not implemented"

	// Schema returns the schema for the local database.
	return
}

func Schema() (*sql.Schema, error) { _ = "STUB: not implemented"; return nil, nil }

// NOTE: coded state migrations can be added here
// They can be a part of this localsql package

// Open opens a local database.
func Open(uri string, opts ...sql.Opt) (*database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Open opens an in-memory local database.
func InMemory(opts ...sql.Opt) *database { _ = "STUB: not implemented"; return nil }

// InMemoryTest returns an in-mem database for testing and ensures database is closed during `tb.Cleanup`.
func InMemoryTest(tb testing.TB, opts ...sql.Opt) sql.LocalDatabase {
	_ = "STUB: not implemented"
	return *new(sql.LocalDatabase)
}
