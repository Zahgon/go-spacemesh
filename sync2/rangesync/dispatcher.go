package rangesync

import (
	"context"
	"io"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/server"
)

// Handler is a function that handles a request for a Dispatcher.
type Handler func(context.Context, p2p.Peer, io.ReadWriter) error

// Dispatcher multiplexes a P2P Server to multiple set reconcilers.
type Dispatcher struct {
	*server.Server
	mtx      sync.Mutex
	logger   *zap.Logger
	handlers map[string]Handler
}

// NewDispatcher creates a new Dispatcher.
func NewDispatcher(logger *zap.Logger) *Dispatcher { _ = "STUB: not implemented"; return nil }

// SetupServer creates a new P2P Server for the Dispatcher.
func (d *Dispatcher) SetupServer(
	host host.Host,
	proto string,
	opts ...server.Opt,
) *server.Server {
	_ = "STUB: not implemented"
	return nil
}

// Register registers a handler with a Dispatcher.
func (d *Dispatcher) Register(name string, h Handler) { _ = "STUB: not implemented"; return }

// Dispatch dispatches a request to a handler.
func (d *Dispatcher) Dispatch(
	ctx context.Context,
	peer p2p.Peer,
	req []byte,
	stream io.ReadWriter,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}
