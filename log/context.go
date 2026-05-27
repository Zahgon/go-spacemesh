package log

import (
	"context"
)

type correlationIDType int

const (
	requestIDKey correlationIDType = iota
	sessionIDKey

	// Request and Session fields need to be separate so that each can have extra fields associated
	// with it and they don't overwrite each other.
	requestFieldsKey
	sessionFieldsKey
)

// withRequestID returns a context which knows its request ID.
// A request ID tracks the lifecycle of a single request across all execution contexts, including
// multiple goroutines, task queues, workers, etc. The canonical example is an incoming message
// received over the network. Rule of thumb: requests "traverse the heap" and may be passed from one
// session to another via channels.
// This requires a requestId string, and optionally, other LoggableFields that are added to
// context and printed in contextual logs.
func withRequestID(ctx context.Context, requestID string, fields ...LoggableField) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WithNewRequestID does the same thing as WithRequestID but generates a new, random requestId.
// It can be used when there isn't a single, clear, unique id associated with a request (e.g.,
// a block or tx hash).
func WithNewRequestID(ctx context.Context, fields ...LoggableField) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ExtractSessionID extracts the session id from a context object.
func ExtractSessionID(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// ExtractRequestID extracts the request id from a context object.
func ExtractRequestID(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// ExtractSessionFields extracts additional loggable fields associated with the session from a context object.
func ExtractSessionFields(ctx context.Context) (fields []LoggableField) {
	_ = "STUB: not implemented"
	return nil
}

// ExtractRequestFields extracts additional loggable fields associated with the request from a context object.
func ExtractRequestFields(ctx context.Context) (fields []LoggableField) {
	_ = "STUB: not implemented"
	return nil
}

// WithSessionID returns a context which knows its session ID
// A session ID tracks a single thread of execution. This may include multiple goroutines running
// concurrently, but it does not include asynchronous task execution such as task queues handled
// in separate threads. The canonical example is a single protocol routine, down to the point where
// an incoming request is either handled, or else handed off to another routine via a channel. Rule
// of thumb: sessions live entirely on the stack, and never on the heap (and should not be passed over
// channels).
// This requires a sessionId string, and optionally, other LoggableFields that are added to
// context and printed in contextual logs.
func WithSessionID(ctx context.Context, sessionID string, fields ...LoggableField) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WithNewSessionID does the same thing as WithSessionID but generates a new, random sessionId.
// It can be used when there isn't a single, clear, unique id associated with a session.
func WithNewSessionID(ctx context.Context, fields ...LoggableField) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func shortUUID() string {
	_ = "STUB: not implemented"
	// 4 first bytes from uuid in hex. before the first hyphen
	return ""
}
