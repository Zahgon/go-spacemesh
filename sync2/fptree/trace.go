package fptree

import (
	"fmt"

	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

type shortened rangesync.KeyBytes

// trace represents a logging facility for tracing FPTree operations, using indentation to
// show their nested structure.
type trace struct {
	traceEnabled bool
	traceStack   []string
}

func (t *trace) out(msg string) { _ = "STUB: not implemented"; return }

// enter marks the entry to a function, printing the log message with the given format
// string and arguments.
func (t *trace) enter(format string, args ...any) { _ = "STUB: not implemented"; return }

// leave marks the exit from a function, printing the results of the function call
// together with the same log message contents which was used in the corresponding enter
// call.
func (t *trace) leave(results ...any) { _ = "STUB: not implemented"; return }

// log prints a log message with the given format string and arguments.
func (t *trace) log(format string, args ...any) { _ = "STUB: not implemented"; return }

// seqFormatter is a lazy formatter for SeqResult.
type seqFormatter struct {
	sr rangesync.SeqResult
}

// String implements fmt.Stringer.
func (f seqFormatter) String() string { _ = "STUB: not implemented"; return "" }

// formatSeqResult returns a fmt.Stringer for the SeqResult that
// formats the sequence result lazily.
func formatSeqResult(sr rangesync.SeqResult) fmt.Stringer {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer)
}

func preprocessTraceArgs(args []any) []any { _ = "STUB: not implemented"; return nil }
