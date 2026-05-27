package rangesync

import (
	"context"
	"errors"
	"io"
	"sync/atomic"

	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/codec"
)

const (
	// TODO: currently, in RangeSetReconciler, the reconciliation process may block
	// indefinitely if the send queue in wireConduit overflows, causing connection to
	// time out, after which reconciliation is interrupted.
	// A way to partly mitigate this issue would be the following:
	// 1. Invoking Receive() immediately on any incoming set items.  It may help that
	//    the set is only actually modified via Add() when handling items associated
	//    with Recent messages.
	// 2. Branch out into more goroutines when handling incoming messages upon sending
	//    being blocked. This way, we'll allow the remote side to receive some
	//    messages by handling the messages it sent us and unblocking its send queue.
	// The OrderedSet is only added to by RangeSetReconciler when receiving recent
	// items. Receive() semantics should be updated so that Receive() being called on
	// OrderedSet's copies' does the same as Receive() being called on the original
	// OrderedSet. After these changes, it should be easy enough to parallelize
	// RangeSetReconciler's message handling as needed, passing copies of OrderedSet
	// to the new goroutines.
	sendQueueSize = 200000
)

var (
	ErrTrafficLimitExceeded = errors.New("sync traffic limit exceeded")
	ErrMessageLimitExceeded = errors.New("sync message limit exceeded")
)

// wireConduit is an implementation of the Conduit interface that sends and receives
// messages over a stream represented by an io.ReadWriter.
type wireConduit struct {
	stream     io.ReadWriter
	cfg        RangeSetReconcilerConfig
	eg         errgroup.Group
	sendCh     chan SyncMessage
	stopCh     chan struct{}
	nBytesSent atomic.Int64
	nBytesRecv atomic.Int64
	nMsgsSent  atomic.Int64
	nMsgsRecv  atomic.Int64
}

var _ Conduit = &wireConduit{}

// startWireConduit sets up a new wireConduit using the given context, stream and options.
func startWireConduit(ctx context.Context, s io.ReadWriter, cfg RangeSetReconcilerConfig) *wireConduit {
	_ = "STUB: not implemented"
	return nil
}

func (c *wireConduit) closeStream() { _ = "STUB: not implemented"; return }

// Stop stops the wireConduit's background sender, but doesn't wait for it to finish
// sending pending messages.
func (c *wireConduit) Stop() { _ = "STUB: not implemented"; return }

// if there was in error, there's no point in waiting for the send
// goroutine to finish, so we interrupt it by closing the stream

// End stops the wireConduit's background sender, waiting for it to finish sending pending
// messages.
func (c *wireConduit) End() { _ = "STUB: not implemented"; return }

// checkLimits checks if the traffic or message limits have been exceeded.
func (c *wireConduit) checkLimits() error { _ = "STUB: not implemented"; return nil }

// NextMessage implements Conduit.
func (c *wireConduit) NextMessage() (SyncMessage, error) {
	_ = "STUB: not implemented"
	return *new(SyncMessage), nil
}

func (c *wireConduit) nextMessage() (SyncMessage, int, error) {
	_ = "STUB: not implemented"
	return *new(SyncMessage), 0, nil
}

// Send implements Conduit.
func (c *wireConduit) Send(m SyncMessage) error { _ = "STUB: not implemented"; return nil }

// bytesSent returns the total number of bytes sent.
func (c *wireConduit) bytesSent() int { _ = "STUB: not implemented"; return 0 }

// bytesReceived returns the total number of bytes received.
func (c *wireConduit) bytesReceived() int { _ = "STUB: not implemented"; return 0 }

// messagesSent returns the total number of messages sent.
func (c *wireConduit) messagesSent() int { _ = "STUB: not implemented"; return 0 }

// messagesReceived returns the total number of messages received.
func (c *wireConduit) messagesReceived() int { _ = "STUB: not implemented"; return 0 }

func writeMessage(w io.Writer, m SyncMessage) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeMessage[T any, PT interface {
	SyncMessage
	codec.Decodable
	*T
}](r io.Reader) (SyncMessage, int, error) {
	_ = "STUB: not implemented"
	return *new(SyncMessage), 0, nil
}
