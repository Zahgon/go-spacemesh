package server

import (
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jonboulle/clockwork"
)

const (
	deadlineAdjusterChunkSize = 4096
)

type deadlineAdjusterError struct {
	what         string
	innerErr     error
	elapsed      time.Duration
	totalRead    int
	totalWritten int
	timeout      time.Duration
	hardTimeout  time.Duration
}

func (err *deadlineAdjusterError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (err *deadlineAdjusterError) Error() string { _ = "STUB: not implemented"; return "" }

type deadlineAdjuster struct {
	peerStream
	adjustMtx       sync.Mutex
	timeout         time.Duration
	hardTimeout     time.Duration
	totalRead       atomic.Int64
	totalWritten    atomic.Int64
	start           time.Time
	clock           clockwork.Clock
	chunkSize       int
	nextAdjustRead  int
	nextAdjustWrite int
	hardDeadline    time.Time
}

var _ io.ReadWriteCloser = &deadlineAdjuster{}

func newDeadlineAdjuster(stream peerStream, timeout, hardTimeout time.Duration) *deadlineAdjuster {
	_ = "STUB: not implemented"
	return nil
}

func (dadj *deadlineAdjuster) augmentError(what string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the stream. This method is safe to call multiple times.
func (dadj *deadlineAdjuster) Close() error {
	_ = "STUB: not implemented"
	// FIXME: unsure if this is really needed (inherited from the older Server code)
	return nil
}

func (dadj *deadlineAdjuster) adjust() error { _ = "STUB: not implemented"; return nil }

// emulate yamux timeout error

// Do not adjust the deadline too often

// We ignore the error returned by SetDeadline b/c the call
// doesn't work for mock hosts

func (dadj *deadlineAdjuster) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Short read, don't try to read more data

func (dadj *deadlineAdjuster) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadByte implements io.ByteReader, which is needed for varint.ReadUvarint, which is
// used to read request length.
func (dadj *deadlineAdjuster) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }
