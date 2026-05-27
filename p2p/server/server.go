package server

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"

	"github.com/spacemeshos/go-spacemesh/p2p/peerinfo"
)

type DecayingTagSpec struct {
	Interval time.Duration `mapstructure:"interval"`
	Inc      int           `mapstructure:"inc"`
	Dec      int           `mapstructure:"dec"`
	Cap      int           `mapstructure:"cap"`
}

// ErrNotConnected is returned when peer is not connected.
var ErrNotConnected = errors.New("peer is not connected")

// Opt is a type to configure a server.
type Opt func(s *Server)

// WithTimeout configures stream timeout.
// The requests are terminated when no data is received or sent for
// the specified duration.
func WithTimeout(timeout time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithHardTimeout configures the hard timeout for requests.
// Requests are terminated if they take longer than the specified
// duration.
func WithHardTimeout(timeout time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithLog configures logger for the server.
func WithLog(log *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithRequestSizeLimit(limit int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithMetrics will enable metrics collection in the server.
func WithMetrics() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithQueueSize parametrize number of message that will be kept in queue
// and eventually processed by server. Otherwise stream is closed immediately.
//
// Size of the queue should be set to account for maximum expected latency, such as if expected latency is 10s
// and server processes 1000 requests per second size should be 100.
//
// Defaults to 100.
func WithQueueSize(size int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithRequestsPerInterval parametrizes server rate limit to limit maximum amount of bandwidth
// that this handler can consume.
//
// Defaults to 100 requests per second.
func WithRequestsPerInterval(n int, interval time.Duration) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithDecayingTag specifies P2P decaying tag that is applied to the peer when a request
// is being served.
func WithDecayingTag(tag DecayingTagSpec) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// Handler is a handler to be defined by the application.
type Handler func(context.Context, peer.ID, []byte) ([]byte, error)

// StreamHandler is a handler that writes the response to the stream directly instead of
// buffering the serialized representation.
type StreamHandler func(context.Context, peer.ID, []byte, io.ReadWriter) error

// StreamRequestCallback is a function that executes a streamed request.
type StreamRequestCallback func(context.Context, io.ReadWriter) error

// ServerError is used by the client (Request/StreamRequest) to represent an error
// returned by the server.
type ServerError struct {
	msg string
}

func NewServerError(msg string) *ServerError { _ = "STUB: not implemented"; return nil }

func (err *ServerError) Error() string { _ = "STUB: not implemented"; return "" }

//go:generate scalegen -types Response

// Response is a server response.
type Response struct {
	// keep in line with limit of ResponseMessage.Data in `fetch/wire_types.go`
	Data  []byte `scale:"max=272629760"` // 260 MiB > 8.0 mio ATX * 32 bytes per ID
	Error string `scale:"max=1024"`
}

// Server for the Handler.
type Server struct {
	logger              *zap.Logger
	protocol            string
	handler             StreamHandler
	timeout             time.Duration
	hardTimeout         time.Duration
	requestLimit        int
	queueSize           int
	requestsPerInterval int
	interval            time.Duration
	decayingTagSpec     *DecayingTagSpec
	decayingTag         connmgr.DecayingTag

	limit   *rate.Limiter
	sem     *semaphore.Weighted
	queue   chan request
	stopped chan struct{}

	metrics *tracker // metrics can be nil

	h Host
}

// New server for the handler.
func New(h Host, proto string, handler StreamHandler, opts ...Opt) *Server {
	_ = "STUB: not implemented"
	return nil
}

// at most s.queueSize requests block here, the others are dropped with the semaphore

type request struct {
	stream   network.Stream
	received time.Time
}

func (s *Server) peerInfo() peerinfo.PeerInfo {
	_ = "STUB: not implemented"
	return *new(peerinfo.PeerInfo)
}

func (s *Server) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) queueHandler(ctx context.Context, peer peer.ID, stream network.Stream) bool {
	_ = "STUB: not implemented"
	return false
}

// Request sends a binary request to the peer.
func (s *Server) Request(ctx context.Context, pid peer.ID, req []byte, extraProtocols ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ensure that a canceled context is returned as the right error

// StreamRequest sends a binary request to the peer. The response is read from the stream
// by the specified callback.
func (s *Server) StreamRequest(
	ctx context.Context,
	pid peer.ID,
	req []byte,
	callback StreamRequestCallback,
	extraProtocols ...string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) streamRequest(
	ctx context.Context,
	pid peer.ID,
	req []byte,
	extraProtocols ...string,
) (
	stm io.ReadWriteCloser,
	info *peerinfo.Info,
	err error,
) {
	_ = "STUB: not implemented"
	return *new(io.ReadWriteCloser), nil, nil
}

// NumAcceptedRequests returns the number of accepted requests for this server.
// It is used for testing.
func (s *Server) NumAcceptedRequests() int { _ = "STUB: not implemented"; return 0 }

func writeResponse(w io.Writer, resp *Response) error { _ = "STUB: not implemented"; return nil }

func WriteErrorResponse(w io.Writer, respErr error) error { _ = "STUB: not implemented"; return nil }

func ReadResponse(r io.Reader, toCall func(resLen uint32) (int, error)) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func WrapHandler(handler Handler) StreamHandler {
	_ = "STUB: not implemented"
	return *new(StreamHandler)
}
