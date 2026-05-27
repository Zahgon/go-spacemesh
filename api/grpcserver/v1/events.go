package v1

import (
	"context"

	"github.com/libp2p/go-libp2p/core/event"
)

const subscriptionChanBufSize = 1 << 16

// Errors for cases with a full event buffer.
var (
	errTxBufferFull          = "tx buffer is full"
	errLayerBufferFull       = "layer buffer is full"
	errAccountBufferFull     = "account buffer is full"
	errRewardsBufferFull     = "rewards buffer is full"
	errActivationsBufferFull = "activations buffer is full"
	errStatusBufferFull      = "status buffer is full"
	errErrorsBufferFull      = "errors buffer is full"
)

func consumeEvents[T any](
	ctx context.Context,
	subscription event.Subscription,
) (out <-chan T, bufFull <-chan struct{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

func closeSubscription(accountSubscription event.Subscription) { _ = "STUB: not implemented"; return }
