package events

import (
	"context"
)

func newSubconf(opts ...SubOpt) *subconf { _ = "STUB: not implemented"; return nil }

type subconf struct {
	buffer int
}

// SubOpt for changing subscribe options.
type SubOpt func(*subconf)

// WithBuffer changes subscription buffer size.
func WithBuffer(n int) SubOpt { _ = "STUB: not implemented"; return *new(SubOpt) }

func subscribe[T any](matcher func(*T) bool, opts ...SubOpt) (*BufferedSubscription[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subscribe to the objects of type T.
func Subscribe[T any](opts ...SubOpt) (*BufferedSubscription[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SubscribeMatched subscribes and filters results before adding them to the Out channel.
func SubscribeMatched[T any](matcher func(*T) bool, opts ...SubOpt) (*BufferedSubscription[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BufferedSubscription is meant to be used by API subscribers.
//
// Majority of the events are published on the critical consensus path, and they
// can't block consensus progress if consumer is slow.
// To account for slow consumer each subscription maintains an internal buffer
// with specific events, if this buffer overflows consumer should be dropped
// and stream restarted.
type BufferedSubscription[T any] struct {
	cancel func()
	result chan T
	full   chan struct{}
}

// Close non-blocking close of the subscription.
func (sub *BufferedSubscription[T]) Close() {
	_ = "STUB: not implemented"

	// Out is a channel with subscription results.
	return
}

func (sub *BufferedSubscription[T]) Out() <-chan T {
	_ = "STUB: not implemented"

	// Full is closed if subscriptions buffer overflows.
	return nil
}

func (sub *BufferedSubscription[T]) Full() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (sub *BufferedSubscription[T]) run(ctx context.Context, s Subscription, matcher func(*T) bool) {
	_ = "STUB: not implemented"
	return
}
