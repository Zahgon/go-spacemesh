package ballotwriter

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

var writerDelay = 100 * time.Millisecond

type BallotWriter struct {
	db     sql.StateDatabase
	logger *zap.Logger

	atxMu sync.Mutex
	timer *time.Ticker

	ballotBatch       map[types.BallotID]*types.Ballot // the current
	ballotBatchResult *batchResult
}

func New(db sql.StateDatabase, logger *zap.Logger) *BallotWriter {
	_ = "STUB: not implemented"
	// create a stopped ticker that can be started later
	return nil
}

// Start the forever-loop that flushes the ballots to the DB
// at-least every `writerDelay`. The caller is responsible
// to call Start in a different goroutine.
func (w *BallotWriter) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

// we hang on to this lock for the entire duration of this select case branch
// as it simplifies the business logic.

// copy the result type

// Store a ballot. Will return the error encountered during the
// write to the db. May also return a context canceled error during
// shutdown.
func (w *BallotWriter) Store(b *types.Ballot) error { _ = "STUB: not implemented"; return nil }

type batchResult struct {
	doneC chan struct{}
	err   error
}
