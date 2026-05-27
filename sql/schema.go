package sql

import (
	"io"

	"go.uber.org/zap"
)

const (
	SchemaPath        = "schema/schema.sql"
	UpdatedSchemaPath = "schema/schema.sql.updated"
)

// LoadDBSchemaScript retrieves the database schema as text.
func LoadDBSchemaScript(db Executor) (string, error) { _ = "STUB: not implemented"; return "", nil }

// The following SQL query ensures that tables are listed first,
// ordered by name, and then all other objects, ordered by their table name
// and then by their own name.

// On Windows, the result contains extra carriage returns

// Schema represents database schema.
type Schema struct {
	Script        string
	Migrations    MigrationList
	skipMigration map[int]struct{}
}

// Diff diffs the database schema against the actual schema.
// If there's no differences, it returns an empty string.
func (s *Schema) Diff(actualScript string) string {
	_ = "STUB: not implemented"
	// If the difference is only in whitespaces, consider the schemas equal.
	return ""
}

// WriteToFile writes the schema to the corresponding updated schema file.
func (s *Schema) WriteToFile(basedir string) error { _ = "STUB: not implemented"; return nil }

// SkipMigrations skips the specified migrations.
func (s *Schema) SkipMigrations(i ...int) { _ = "STUB: not implemented"; return }

// Apply applies the schema to the database.
func (s *Schema) Apply(db Database) error { _ = "STUB: not implemented"; return nil }

func (s *Schema) CheckDBVersion(logger *zap.Logger, db Database) (before, after int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (s *Schema) setVersion(db Executor, version int) error {
	_ = "STUB: not implemented"
	// binding values in pragma statement is not allowed
	return nil
}

// Migrate performs database migration. In case if migrations are disabled, the database
// version is checked but no migrations are run, and if the database is too old and
// migrations are disabled, an error is returned.
func (s *Schema) Migrate(logger *zap.Logger, db Database, before, vacuumState int) error {
	_ = "STUB: not implemented"
	return nil
}

// MigrateTempDB performs database migration on the temporary database.
// It doesn't use transactions and the temporary database should be considered
// invalid and discarded if it fails.
// The database is switched into synchronous mode with WAL journal enabled and
// synced after the migrations are completed before setting the database version,
// which triggers file sync.
func (s *Schema) MigrateTempDB(logger *zap.Logger, db Database, before int) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't set the version here as if any migration fails,
// the temporary database is considered invalid and should be discarded.

// Enable WAL journal and synchronous mode to ensure the database is synced

// This should trigger file sync

// SchemaGenOpt represents a schema generator option.
type SchemaGenOpt func(g *SchemaGen)

func withDefaultOut(w io.Writer) SchemaGenOpt { _ = "STUB: not implemented"; return *new(SchemaGenOpt) }

// SchemaGen generates database schema files.
type SchemaGen struct {
	logger     *zap.Logger
	schema     *Schema
	defaultOut io.Writer
}

// NewSchemaGen creates a new SchemaGen instance.
func NewSchemaGen(logger *zap.Logger, schema *Schema, opts ...SchemaGenOpt) *SchemaGen {
	_ = "STUB: not implemented"
	return nil
}

// Generate generates database schema and writes it to the specified file.
// If an empty string is specified as outputFile, os.Stdout is used for output.
func (g *SchemaGen) Generate(outputFile string) error { _ = "STUB: not implemented"; return nil }

func logQueryInterceptor(logger *zap.Logger) Interceptor {
	_ = "STUB: not implemented"
	return *new(Interceptor)
}
