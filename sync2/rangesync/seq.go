package rangesync

import (
	"iter"

	"go.uber.org/zap/zapcore"
)

// Seq represents an ordered sequence of elements.
// Most sequences are finite. Infinite sequences are explicitly mentioned in the
// documentation of functions/methods that return them.
type Seq iter.Seq[KeyBytes]

var _ zapcore.ArrayMarshaler = Seq(nil)

// First returns the first element from the sequence, if any.
// If the sequence is empty, it returns nil.
func (s Seq) First() KeyBytes { _ = "STUB: not implemented"; return *new(KeyBytes) }

// FirstN returns the first n elements from the sequence.
func (s Seq) FirstN(n int) []KeyBytes { _ = "STUB: not implemented"; return nil }

// Collect returns all elements in the sequence as a slice.
// It may not be very efficient due to reallocations, and thus it should only be used for
// small sequences or for testing.
func (s Seq) Collect() []KeyBytes { _ = "STUB: not implemented"; return nil }

// MarshalLogArray implements zapcore.ArrayMarshaler.
func (s Seq) MarshalLogArray(enc zapcore.ArrayEncoder) error { _ = "STUB: not implemented"; return nil }

// Limit limits sequence to n elements.
func (s Seq) Limit(n int) Seq { _ = "STUB: not implemented"; return *new(Seq) }

// ensure reusability

// EmptySeq returns an empty sequence.
func EmptySeq() Seq { _ = "STUB: not implemented"; return *new(Seq) }

// SeqErrorFunc is a function that returns an error that happened during iteration, if
// any.
type SeqErrorFunc func() error

// NoSeqError is a SeqErrorFunc that always returns nil (no error).
var NoSeqError SeqErrorFunc = func() error { return nil }

// SeqError returns a SeqErrorFunc that always returns the given error.
func SeqError(err error) SeqErrorFunc { _ = "STUB: not implemented"; return *new(SeqErrorFunc) }

// SeqResult represents the result of a function that returns a sequence.
// Error method most be called to check if an error occurred after
// processing the sequence.
// Error is reset at the beginning of each Seq call (iteration over the sequence).
type SeqResult struct {
	Seq   Seq
	Error SeqErrorFunc
}

// MarshalLogArray implements zapcore.ArrayMarshaler.
func (s SeqResult) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
	// never returns an error
}

// First returns the first element from the result's sequence, if any.
// If the sequence is empty, it returns nil.
func (s SeqResult) First() (KeyBytes, error) { _ = "STUB: not implemented"; return *new(KeyBytes), nil }

// FirstN returns the first n elements from the result's sequence.
func (s SeqResult) FirstN(n int) ([]KeyBytes, error) { _ = "STUB: not implemented"; return nil, nil }

// Collect returns all elements in the result's sequence as a slice.
// It may not be very efficient due to reallocations, and thus it should only be used for
// small sequences or for testing.
func (s SeqResult) Collect() ([]KeyBytes, error) { _ = "STUB: not implemented"; return nil, nil }

// IsEmpty returns true if the sequence in SeqResult is empty.
// It also checks for errors.
func (s SeqResult) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Limit limits SeqResult to n elements.
func (s SeqResult) Limit(n int) SeqResult { _ = "STUB: not implemented"; return *new(SeqResult) }

// EmptySeqResult returns an empty sequence result.
func EmptySeqResult() SeqResult { _ = "STUB: not implemented"; return *new(SeqResult) }

// ErrorSeqResult returns a sequence result with an empty sequence and an error.
func ErrorSeqResult(err error) SeqResult { _ = "STUB: not implemented"; return *new(SeqResult) }

// MakeSeqResult makes a SeqResult out of a slice.
func MakeSeqResult(items []KeyBytes) SeqResult { _ = "STUB: not implemented"; return *new(SeqResult) }
