package sql

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	sqlite "github.com/go-llsqlite/crawshaw"
	"github.com/go-llsqlite/crawshaw/sqlitex"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	// ErrClosed is returned if database is closed.
	ErrClosed = errors.New("database closed")
	// ErrNoConnection is returned if pooled connection is not available.
	ErrNoConnection = errors.New("database: no free connection")
	// ErrNotFound is returned if requested record is not found.
	ErrNotFound = errors.New("database: not found")
	// ErrObjectExists is returned if database constraints didn't allow to insert an object.
	ErrObjectExists = errors.New("database: object exists")
	// ErrConflict is returned if database constraints didn't allow to update an object.
	ErrConflict = errors.New("database: conflict")
	// ErrTooNew is returned if database version is newer than expected.
	ErrTooNew = errors.New("database version is too new")
	// ErrOldSchema is returned when the database version differs from the expected one
	// and migrations are disabled.
	ErrOldSchema = errors.New("old database version")
)

const (
	beginDefault   = "BEGIN;"
	beginImmediate = "BEGIN IMMEDIATE;"
)

// Statement is an sqlite statement.
type Statement = sqlite.Stmt

// Encoder for parameters.
// Both positional parameters:
// select block from blocks where id = ?1;
//
// and named parameters are supported:
// select blocks from blocks where id = @id;
//
// For complete information see https://www.sqlite.org/c3ref/bind_blob.html.
type Encoder func(*Statement)

// Decoder for sqlite rows.
type Decoder func(*Statement) bool

func defaultConf() *conf { _ = "STUB: not implemented"; return nil }

type conf struct {
	uri                        string
	enableMigrations           bool
	forceFresh                 bool
	forceMigrations            bool
	connections                int
	vacuumState                int
	enableLatency              bool
	cache                      bool
	cacheSizes                 map[QueryCacheKind]int
	logger                     *zap.Logger
	schema                     *Schema
	allowSchemaDrift           bool
	checkSchemaDrift           bool
	temp                       bool
	handleIncompleteMigrations bool
	exclusive                  bool
	readOnly                   bool
	dbName                     string
	connIdleTimeout            time.Duration
}

// WithConnections overwrites number of pooled connections.
func WithConnections(n int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithLogger specifies logger for the database.
func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithMigrationsDisabled disables migrations for the database.
// The migrations are enabled by default.
func WithMigrationsDisabled() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithVacuumState will execute vacuum if database version before the migration was less or equal to the provided value.
func WithVacuumState(i int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithLatencyMetering enables metric that track latency for every database query.
// Note that it will be a significant amount of data, and should not be enabled on
// multiple nodes by default.
func WithLatencyMetering(enable bool) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithQueryCache enables in-memory caching of results of some queries.
func WithQueryCache(enable bool) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithQueryCacheSizes sets query cache sizes for the specified cache kinds.
func WithQueryCacheSizes(sizes map[QueryCacheKind]int) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithForceMigrations forces database to run all the migrations instead
// of using a schema snapshot in case of a fresh database.
func WithForceMigrations(force bool) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithDatabaseSchema specifies database schema script.
func WithDatabaseSchema(schema *Schema) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithAllowSchemaDrift prevents Open from failing upon schema drift when schema drift
// checks are enabled. A warning is printed instead.
func WithAllowSchemaDrift(allow bool) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithNoCheckSchemaDrift disables schema drift checks.
func WithNoCheckSchemaDrift() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func withForceFresh() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithTemp specifies temporary database mode.
// For the temporary database, the migrations are always run in place, and vacuuming is
// nover done.  PRAGMA journal_mode=OFF and PRAGMA synchronous=OFF are used.
func WithTemp() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithDBName sets the name of the database which is used for metrics.
func WithDBName(name string) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithConnIdleTimeout sets idle timeout for connections from the pool
// which are acquired upon first statement executed against a Connection
// passed to the callback of Database.WithConnection. After the timeout,
// the connection is released back to the pool until the next statement.
func WithConnIdleTimeout(timeout time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func withDisableIncompleteMigrationHandling() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithExclusive specifies that the database is to be open in exclusive mode.
// This means that no other processes can open the database at the same time.
// If the database is already open by any process, this Open will fail.
// Any subsequent attempts by other processes to open the database will fail until this db
// handle is closed.
// In Exclusive mode, the database supports just one concurrent connection.
func WithExclusive() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithReadOnly specifies that the database is to be open in read-only mode.
func WithReadOnly() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// Opt for configuring database.
type Opt func(c *conf)

// OpenInMemory creates an in-memory database.
func OpenInMemory(opts ...Opt) (*sqliteDatabase, error) { _ = "STUB: not implemented"; return nil, nil }

// Unique uri is needed to avoid sharing the same in-memory database,
// while allowing multiple connections to the same database.

// InMemory creates an in-memory database for testing and panics if
// there's an error.
func InMemory(opts ...Opt) *sqliteDatabase { _ = "STUB: not implemented"; return nil }

// InMemoryTest returns an in-mem database for testing and ensures database is closed during `tb.Cleanup`.
func InMemoryTest(tb testing.TB, opts ...Opt) *sqliteDatabase {
	_ = "STUB: not implemented"
	// When using empty DB schema, we don't want to check for schema drift due to
	// "PRAGMA user_version = 0;" in the initial schema retrieved from the DB.
	return nil
}

// Open database with options.
//
// Database is opened in WAL mode and pragma synchronous=normal.
// https://sqlite.org/wal.html
// https://www.sqlite.org/pragma.html#pragma_synchronous
func Open(uri string, opts ...Opt) (*sqliteDatabase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openDB(config *conf) (db *sqliteDatabase, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note that SQLITE_OPEN_WAL is not handled by SQLITE api itself,
// but rather by the crawshaw library which executes
// PRAGMA journal_mode=WAL in this case.
// We don't want it for temporary databases as they're not
// using any journal

// If something goes wrong, close the database even in case of a
// panic. This is important for tests that verify incomplete migration.

// In case of VACUUM INTO based migration, prepareDB may close this database and
// open another one.

func prepareDB(logger *zap.Logger, db *sqliteDatabase, config *conf, freshDB bool) (*sqliteDatabase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Temporary database is used for migration and is deleted if migrations
// fail, so we make it faster by disabling journaling and synchronous
// writes.

// ensureDBSchemaUpToDate may replace the original database and open the new one,
// in which case the original db is already closed but we must close the new one.
// If there are migrations to be done in place without vacuuming,
// the original db is returned and we must close it if there's an error.

// ok

func ensureDBSchemaUpToDate(logger *zap.Logger, db *sqliteDatabase, config *conf) (*sqliteDatabase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Temporary database, do migrations without transactions
// and sync afterwards

func Version(uri string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// deleteDB deletes the database at the specified path by removing /path/to/DB* files.
// If the database doesn't exist, no error is returned.
// In addition to what DROP DATABASE does, this also removes the migration marker file.
func deleteDB(path string) error {
	_ = "STUB: not implemented"
	// https://www.sqlite.org/tempfiles.html plus marker *_done
	return nil
}

// moveMigratedDB runs "VACUUM INTO" on the database at fromPath and
// replaces the database at toPath with the vacuumed one. The database
// at fromPath is deleted after the operation.
func moveMigratedDB(config *conf, fromPath, toPath string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Try to open the temporary migrated DB in exclusive mode before deleting the
// original one.
// If the temporary DB is being copied to the original path by another
// process, this will fail and the original database will not be deleted.
// We don't use the proper database schema here because the temporary DB
// may have been created with a different set of migrations.

// Open the freshly vacuumed DB in exclusive mode to avoid race condition when
// another process also tries to vacuum the temporary DB into the original path
// after we close the temporary DB.

func dbMigrationPaths(uri string) (dbPath, migratedPath string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// handleIncompleteCopyMigration handles incomplete copy-based migrations.
// It only works for 'file:' URIs, doing nothing for other URIs.
// It first checks if there's a copy of the database with "_migrate" suffix.
// If it's there, it checks if the migration is complete by checking if
// DBNAME_migrate_done file exists. It it doesn't, the migration is considered
// incomplete and the migrated database is removed. If DBNAME_migrate_done
// file exists, the migration is finalized by running "VACUUM INTO" on the
// migrated database and replacing the original, after which the migrated
// database is deleted.
func handleIncompleteCopyMigration(config *conf) error { _ = "STUB: not implemented"; return nil }

// no migration in progress

// incomplete migration, delete the temporary DB to start over
// after that

// the migration is complete except for the last step

// Interceptor is invoked on every query after it's added to a database using
// PushIntercept. The query will fail if Interceptor returns an error.
type Interceptor func(query string) error

// Database represents a database.
type Database interface {
	Executor
	QueryCache

	// Close closes the database.
	Close() error

	// QueryCount returns the number of queries executed on the database.
	QueryCount() int

	// QueryCache returns the query cache for this database, if it's present,
	// or nil otherwise.
	QueryCache() QueryCache

	// Tx creates deferred sqlite transaction.
	//
	// Deferred transactions are not started until the first statement.
	// Transaction may be started in read mode and automatically upgraded to write mode
	// after one of the write statements.
	//
	// https://www.sqlite.org/lang_transaction.html
	Tx() (Transaction, error)

	// WithTx starts a new transaction and passes it to the exec function.
	// It then commits the transaction if the exec function doesn't return an error,
	// and rolls it back otherwise.
	WithTx(exec func(Transaction) error) error

	// TxImmediate begins a new immediate transaction on the database, that is,
	// a transaction that starts a write immediately without waiting for a write
	// statement.
	// The transaction returned from this function must always be released by calling
	// its Release method. Release rolls back the transaction if it hasn't been
	// committed.
	TxImmediate() (Transaction, error)

	// WithTxImmediate starts a new immediate transaction and passes it to the exec
	// function.
	// An immediate transaction is started immediately, without waiting for a write
	// statement.
	// It then commits the transaction if the exec function doesn't return an error,
	// and rolls it back otherwise.
	WithTxImmediate(exec func(Transaction) error) error

	// WithConnection executes the provided function with a connection from the
	// database pool.
	// If many queries are to be executed in a row, but there's no need for an
	// explicit transaction which may be long-running and thus block
	// WAL checkpointing, it may be preferable to use a single connection for
	// it to avoid database pool overhead.
	// The connection is released back to the pool after the function returns.
	WithConnection(exec func(Executor) error) error

	// Intercept adds an interceptor function to the database. The interceptor
	// functions are invoked upon each query on the database, including queries
	// executed within transactions.
	// The query will fail if the interceptor returns an error.
	// The interceptor can later be removed using RemoveInterceptor with the same key.
	Intercept(key string, fn Interceptor)

	// RemoveInterceptor removes the interceptor function with specified key from the database.
	RemoveInterceptor(key string)
}

// Transaction represents a transaction.
type Transaction interface {
	Executor
	// Commit commits the transaction.
	Commit() error
	// Release releases the transaction. If the transaction hasn't been committed,
	// it's rolled back.
	Release() error
}

type sqliteDatabase struct {
	*queryCache
	pool *sqlitex.Pool

	closed   bool
	closeMux sync.Mutex

	latency    *prometheus.HistogramVec
	queryCount atomic.Int64

	interceptMtx sync.Mutex
	interceptors map[string]Interceptor

	connWaitLatency prometheus.Observer
	poolUsage       prometheus.Gauge

	connIdleTimeout time.Duration
}

var _ Database = &sqliteDatabase{}

func (db *sqliteDatabase) getConn() *sqlite.Conn { _ = "STUB: not implemented"; return nil }

func (db *sqliteDatabase) putConn(conn *sqlite.Conn) { _ = "STUB: not implemented"; return }

func (db *sqliteDatabase) getTx(initstmt string) (*sqliteTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *sqliteDatabase) withTx(initstmt string, exec func(Transaction) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (db *sqliteDatabase) startExclusive() error { _ = "STUB: not implemented"; return nil }

// We don't need to wait for long if the database is busy

// From SQLite docs:
// When the locking-mode is set to EXCLUSIVE, the database connection
// never releases file-locks. The first time the database is read in
// EXCLUSIVE mode, a shared lock is obtained and held. The first time the
// database is written, an exclusive lock is obtained and held.

// We need to perform a transaction to have the database actually locked.
// From SQLite docs, regarding BEGIN EXCLUSIVE / BEGIN IMMEDIATE:
// EXCLUSIVE is similar to IMMEDIATE in that a write transaction is
// started immediately. EXCLUSIVE and IMMEDIATE are the same in WAL mode,
// but in other journaling modes, EXCLUSIVE prevents other database
// connections from reading the database while the transaction is
// underway.

// Tx implements Database.
func (db *sqliteDatabase) Tx() (Transaction, error) {
	_ = "STUB: not implemented"
	return *new(Transaction), nil
}

// WithTx implements Database.
func (db *sqliteDatabase) WithTx(exec func(Transaction) error) error {
	_ = "STUB: not implemented"
	return nil
}

// TxImmediate implements Database.
func (db *sqliteDatabase) TxImmediate() (Transaction, error) {
	_ = "STUB: not implemented"
	return *new(Transaction), nil
}

// WithTxImmediate implements Database.
func (db *sqliteDatabase) WithTxImmediate(exec func(Transaction) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *sqliteDatabase) runInterceptors(query string) error {
	_ = "STUB: not implemented"
	return nil
}

// Exec implements Executor.
//
// If you care about atomicity of the operation (for example writing rewards to multiple accounts)
// Tx should be used. Otherwise sqlite will not guarantee that all side-effects of operations are
// applied to the database if machine crashes.
//
// Note that Exec will block until database is closed or statement has finished.
// If application needs to control statement execution lifetime use one of the transaction.
func (db *sqliteDatabase) Exec(query string, encoder Encoder, decoder Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close implements Database.
func (db *sqliteDatabase) Close() error { _ = "STUB: not implemented"; return nil }

// WithConnection implements Database.
func (db *sqliteDatabase) WithConnection(toCall func(Executor) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Intercept adds an interceptor function to the database. The interceptor functions
// are invoked upon each query. The query will fail if the interceptor returns an error.
// The interceptor can later be removed using RemoveInterceptor with the same key.
func (db *sqliteDatabase) Intercept(key string, fn Interceptor) { _ = "STUB: not implemented"; return }

// PopIntercept removes the interceptor function with specified key from the database.
// If there's no such interceptor, the function does nothing.
func (db *sqliteDatabase) RemoveInterceptor(key string) { _ = "STUB: not implemented"; return }

// vacuumInto runs VACUUM INTO on the database and saves the vacuumed
// database at toPath.
func (db *sqliteDatabase) vacuumInto(toPath string) error { _ = "STUB: not implemented"; return nil }

// copyMigrateDB performs a copy-based migration of the database.
// The source database is always closed by this function.
// Upon success, the migrated database is opened.
func (db *sqliteDatabase) copyMigrateDB(config *conf) (finalDB *sqliteDatabase, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Before we start the migration, re-open the database in exclusive mode
// so that no other connections will be able to use it.
// This will fail if another process is already using this database.

// instead of just copying the source database to the temporary migration DB, use VACUUM INTO.
// This is somewhat slower but achieves two goals:
// 1. The lock is held on the source database while it's being copied
// 2. If the source database has a lot of free pages for whatever reason, those
// are not copied, saving disk space

// Opening the temporary migrated DB runs the actual migrations on it.
// We disable vacuuming here because we're going to vacuum the temporary DB
// into the original one.

// Make sure the temporary DB is fully synced to the disk before creating the marker file.
// We don't need wal_checkpoint(TRUNCATE) here as we're going to delete the temporary DB.

// Create the marker file to indicate that the migration is complete and make sure
// the file is written to the disk before closing the database.
// We could create a table in the temporary database instead of the marker file,
// but as the temporary database is opened without PRAGMA journal_mode=OFF
// and PRAGMA synchronous=OFF, it may become corrupt in case of a crash or power
// outage, so we avoid trying to open it.

// The errors returned by createMarkerFile are already descriptive enough
// so no need to augment them

// At this point, the temporary database is complete and should not be deleted
// until we copy it to the original database location.

// We only close the source database at the end of the migration process
// so that the lock is held. There's a possibility that right after we
// close the source database, another process will see the migrated database
// and the marker file and will try to open the migrated database. If the

// Delete the original database. VACUUM INTO will fail if the destination
// database exists.

// Overwrite the original database with the migrated one.
// The lock is held on the temporary DB during this, preventing concurrent
// go-spacemesh instances to attempt the same operation.

// Open the final DB in the exclusive mode before deleting the source DB, so one of the locks
// is always held. The migrations are already run, so we're disabling them.

// Now we can delete the temporary DB and the marker file.

// If we were not intending to open the database in exclusive mode,
// reopen it in the normal mode

func createMarkerFile(basePath string) error { _ = "STUB: not implemented"; return nil }

// QueryCount returns the number of queries executed, including failed
// queries, but not counting transaction start / commit / rollback.
func (db *sqliteDatabase) QueryCount() int { _ = "STUB: not implemented"; return 0 }

// Return database's QueryCache.
func (db *sqliteDatabase) QueryCache() QueryCache {
	_ = "STUB: not implemented"
	return *new(QueryCache)
}

func exec(conn *sqlite.Conn, query string, encoder Encoder, decoder Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// exhaust iterator

// sqliteTx is wrapper for database transaction.
type sqliteTx struct {
	*queryCache
	db        *sqliteDatabase
	conn      *sqlite.Conn
	committed bool
	err       error
}

func (tx *sqliteTx) begin(initstmt string) error { _ = "STUB: not implemented"; return nil }

// Commit transaction.
func (tx *sqliteTx) Commit() error { _ = "STUB: not implemented"; return nil }

// Release transaction. Every transaction that was created must be released.
func (tx *sqliteTx) Release() error { _ = "STUB: not implemented"; return nil }

// Exec query.
func (tx *sqliteTx) Exec(query string, encoder Encoder, decoder Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// lazyConn is a connection that is acquired lazily from the pool, that is, upon the first
// query, and released after a certain period of inactivity.
type lazyConn struct {
	*queryCache
	db      *sqliteDatabase
	getConn func() *sqlite.Conn
	eg      errgroup.Group
	conn    *sqlite.Conn
	timer   *time.Timer
	doneCh  chan struct{}
	connMtx sync.Mutex
}

func newLazyConn(db *sqliteDatabase) *lazyConn { _ = "STUB: not implemented"; return nil }

func (c *lazyConn) ensureConn() *sqlite.Conn { _ = "STUB: not implemented"; return nil }

// Although TryLock docs say that it's not recommended to use it
// in most cases, this use case is justified.
// If the mutex is already locked here, this means that an SQL
// statement is being executed on the connection, after which
// the idle timer will be restarted, or the connection is currently
// being released.

func (c *lazyConn) releaseConn() { _ = "STUB: not implemented"; return }

func (c *lazyConn) release() {
	_ = "STUB: not implemented"
	// Lock the mutex so that we don't get concurrent releaseConn() from the timer handler.
	return
}

func (c *lazyConn) Exec(query string, encoder Encoder, decoder Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func mapSqliteError(err error) error { _ = "STUB: not implemented"; return nil }

// TODO: we probably should check if there was indeed a context that was
// canceled

// Blob represents a binary blob data. It can be reused efficiently
// across multiple data retrieval operations, minimizing reallocations
// of the underlying byte slice.
type Blob struct {
	Bytes []byte
}

// Resize the underlying byte slice to the specified size.
// The returned slice has length equal n, but it might have a larger capacity.
// Warning: it is not guaranteed to keep the old data.
func (b *Blob) Resize(n int) { _ = "STUB: not implemented"; return }

func (b *Blob) FromColumn(stmt *Statement, col int) { _ = "STUB: not implemented"; return }

// GetBlobSizes returns a slice containing the sizes of blobs
// corresponding to the specified ids. For non-existent ids the
// corresponding value is -1.
func GetBlobSizes(db Executor, cmd string, ids [][]byte) (sizes []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadBlob loads an encoded blob.
func LoadBlob(db Executor, cmd string, id []byte, blob *Blob) error {
	_ = "STUB: not implemented"
	return nil
}

// IsNull returns true if the specified result column is null.
func IsNull(stmt *Statement, col int) bool { _ = "STUB: not implemented"; return false }

// StateDatabase is a Database used for Spacemesh state.
type StateDatabase interface {
	Database
	IsStateDatabase()
}

// LocalDatabase is a Database used for local node data.
type LocalDatabase interface {
	Database
	IsLocalDatabase()
}
