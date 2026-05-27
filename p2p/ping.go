package p2p

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/routing"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	ma "github.com/multiformats/go-multiaddr"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	pingProtectTag = "spacemesh-ping"
)

type pingStat struct {
	numSuccess int
	numFail    int
}

type Ping struct {
	sync.Mutex
	logger   *zap.Logger
	h        host.Host
	peers    []peer.ID
	pr       routing.PeerRouting
	interval time.Duration
	stats    map[peer.ID]*pingStat
	cancel   context.CancelFunc
	eg       errgroup.Group
}

type PingOpt func(p *Ping)

func WithPingInterval(d time.Duration) PingOpt { _ = "STUB: not implemented"; return *new(PingOpt) }

func NewPing(logger *zap.Logger, h host.Host, peers []peer.ID, pr routing.PeerRouting, opts ...PingOpt) *Ping {
	_ = "STUB: not implemented"
	return nil
}

func (p *Ping) Start() { _ = "STUB: not implemented"; return }

func (p *Ping) Stop() { _ = "STUB: not implemented"; return }

func (p *Ping) doPing(ctx context.Context, peerID peer.ID) (<-chan ping.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runForPeer runs ping for the specific peer till an error occurs
// or the context is canceled.
func (p *Ping) runForPeer(ctx context.Context, peerID peer.ID, addrs []ma.Multiaddr) error {
	_ = "STUB: not implemented"
	// go-libp2p Ping tends to stop working after an error is received on the channel.
	// This behavior is probably not intended and may change later.
	// So, we create a child context here and cancel it when we receive an error
	// on the channel returned by Ping, and then after a delay we restart
	// Ping for this peer.
	return nil
}

func (p *Ping) startForPeer(ctx context.Context, peerID peer.ID) { _ = "STUB: not implemented"; return }

func (p *Ping) record(peerID peer.ID, success bool) { _ = "STUB: not implemented"; return }

func (p *Ping) Stats(peerID peer.ID) (numSuccess, numFail int) {
	_ = "STUB: not implemented"
	return 0, 0
}
