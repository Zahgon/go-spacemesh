package handshake

import (
	"context"
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"
	ma "github.com/multiformats/go-multiaddr"
	"go.uber.org/zap"
)

const (
	handshakeTimeout       = 5 * time.Second
	handshakeAttempts      = 3
	handshakeRetryInterval = 1 * time.Second
	maxCookieSize          = 64
)

var cookieStreamPrefix = []byte{
	0x21, 0x23, 0x42, 0x42,
}

// NetworkCookie specifies a sequence of bytes that can be used to
// prevent peers from different networks from communicating with each
// other.
type NetworkCookie []byte

// NoNetworkCookie represents an empty NetworkCookie.
var NoNetworkCookie NetworkCookie = nil

// Empty returns true if the network cookie is empty.
func (nc NetworkCookie) Empty() bool { _ = "STUB: not implemented"; return false }

// String returns string representation of the NetworkCookie, which is
// a hex string.
func (nc NetworkCookie) String() string { _ = "STUB: not implemented"; return "" }

// Equal returns true if this cookie is the same as the other cookie.
func (nc NetworkCookie) Equal(other NetworkCookie) bool { _ = "STUB: not implemented"; return false }

// Option specifies a handshake option.
type Option func(w *transportWrapper)

// WithLog specifies a logger for handshake.
func WithLog(logger *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLog specifies handshake timeout.
func WithTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAttempts specifies handshake retry count.
func WithAttempts(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

type transportWrapper struct {
	transport.Transport
	nc       NetworkCookie
	logger   *zap.Logger
	timeout  time.Duration
	attempts int
}

var (
	_ transport.Transport = &transportWrapper{}
	_ io.Closer           = &transportWrapper{}
)

// MaybeWrapTransport adds a handshake wrapper around the Transport if
// network cookie nc is not empty, otherwise it just returns the
// transport.
func MaybeWrapTransport(t transport.Transport, nc NetworkCookie, opts ...Option) transport.Transport {
	_ = "STUB: not implemented"
	return *new(transport.Transport)
}

func (tr *transportWrapper) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (tr *transportWrapper) handshake(ctx context.Context, c transport.CapableConn) (retry bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (tr *transportWrapper) Listen(laddr ma.Multiaddr) (transport.Listener, error) {
	_ = "STUB: not implemented"
	return *new(transport.Listener), nil
}

func (tr *transportWrapper) Close() error { _ = "STUB: not implemented"; return nil }

type listenerWrapper struct {
	transport.Listener
	nc       NetworkCookie
	logger   *zap.Logger
	timeout  time.Duration
	attempts int
}

var _ transport.Listener = &listenerWrapper{}

func (l *listenerWrapper) Accept() (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (l *listenerWrapper) handshake(c transport.CapableConn) (retry bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Write our cookie even in case of mismatch so that the
// peer sees the reason for disconnect immediately

func writeCookie(w io.Writer, nc NetworkCookie) error { _ = "STUB: not implemented"; return nil }

func readCookie(r io.Reader) (nc NetworkCookie, mayHaveCookie bool, err error) {
	_ = "STUB: not implemented"
	return *new(NetworkCookie), false, nil
}
