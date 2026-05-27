package transactions

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

// ResultsFilter applies filter on transaction results query.
type ResultsFilter struct {
	Address    *types.Address
	Start, End *types.LayerID
	TID        *types.TransactionID
}

func (f *ResultsFilter) query() string { _ = "STUB: not implemented"; return "" }

func (f *ResultsFilter) binding(stmt *sql.Statement) { _ = "STUB: not implemented"; return }

// IterateResults allows to control iteration by the output of `fn`.
func IterateResults(db sql.Executor, filter ResultsFilter, fn func(*types.TransactionWithResult) bool) error {
	_ = "STUB: not implemented"
	return nil
}
