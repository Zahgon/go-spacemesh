package sqlstore

import (
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// dbSeq represents a sequence of IDs from a database table.
type dbSeq struct {
	// database
	db sql.Executor
	// starting point
	from rangesync.KeyBytes
	// table snapshot to use
	sts *SyncedTableSnapshot
	// currently used chunk size
	chunkSize int
	// timestamp used to fetch recent IDs
	// (nanoseconds since epoch, 0 if not in the "recent" mode)
	ts int64
	// maximum value for chunkSize
	maxChunkSize int
	// current chunk of items
	chunk []rangesync.KeyBytes
	// position within the current chunk
	pos int
	// lentgh of each key in bytes
	keyLen int
	// true if there is only a single chunk in the sequence.
	// It is set after loading the initial chunk and finding that it's the only one.
	singleChunk bool
}

// idsFromTable iterates over the id field values in a database table.
func idsFromTable(
	db sql.Executor,
	sts *SyncedTableSnapshot,
	from rangesync.KeyBytes,
	ts int64,
	chunkSize int,
	maxChunkSize int,
) rangesync.SeqResult {
	_ = "STUB: not implemented"
	return *new(rangesync.SeqResult)
}

// load makes sure the current chunk is loaded.
func (s *dbSeq) load() error { _ = "STUB: not implemented"; return nil }

// we have a single-chunk DB sequence, don't need to reload,
// just wrap around

// make sure the chunk is large enough

// if the chunk size was reduced due to a short chunk before wraparound, we need
// to extend it back

// we reuse existing slices when possible for retrieving new IDs

// empty chunk

// already wrapped around or started from 0,
// the set is empty

// wrap around

// short chunk means there are no more items after it,
// start the next chunk from 0

// wrapping around on an incomplete chunk that started
// from 0 means we have just a single chunk

// use last item incremented by 1 as the start of the next chunk

// inc may wrap around if it's 0xffff...fff, but it's fine

// if we wrapped around and the current chunk started from 0,
// we have just a single chunk

// iterate iterates over the table rows.
func (s *dbSeq) iterate(yield func(k rangesync.KeyBytes) bool) error {
	_ = "STUB: not implemented"
	return nil
}
