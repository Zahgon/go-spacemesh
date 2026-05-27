package peerinfo

import (
	"context"
	"sync"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/protocol/holepunch"
	ma "github.com/multiformats/go-multiaddr"
	"golang.org/x/sync/errgroup"
)

type Kind string

const (
	KindUknown            Kind = ""
	KindInbound           Kind = "inbound"
	KindOutbound          Kind = "outbound"
	KindHolePunchInbound  Kind = "hp_inbound"
	KindHolePunchOutbound Kind = "hp_outbound"
	KindHolePunchUnknown  Kind = "hp_unknown"
	KindRelayInbound      Kind = "relay_in"
	KindRelayOutbound     Kind = "relay_out"

	bpsInterval1 = 10 * time.Second
	bpsInterval2 = 5 * time.Minute

	totalProto = "__total__"
	otherProto = "__other__"
)

type PeerRequestStats struct {
	mtx          sync.Mutex
	successCount int
	failureCount int
	duration     time.Duration
}

func (ps *PeerRequestStats) SuccessCount() int { _ = "STUB: not implemented"; return 0 }

func (ps *PeerRequestStats) FailureCount() int { _ = "STUB: not implemented"; return 0 }

func (ps *PeerRequestStats) Latency() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (ps *PeerRequestStats) RequestDone(took time.Duration, success bool) {
	_ = "STUB: not implemented"
	return
}

type DataStats struct {
	mtx sync.Mutex
	// [0] is the current value
	// [1] is prev value for rate 1
	// [2] is prev value for rate 2
	bytesSent     [3]int64
	bytesReceived [3]int64
	// [0] is for the shorter-interval rate, [1] is for the longer-interval rate
	recvRate [2]int64
	sendRate [2]int64
}

func rateIndex(which int) int { _ = "STUB: not implemented"; return 0 }

func bpsInterval(which int) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (ds *DataStats) Tick(which int) { _ = "STUB: not implemented"; return }

func (ds *DataStats) RecvRate(which int) int64 { _ = "STUB: not implemented"; return 0 }

func (ds *DataStats) SendRate(which int) int64 { _ = "STUB: not implemented"; return 0 }

func (ds *DataStats) RecordSent(n int64) { _ = "STUB: not implemented"; return }

func (ds *DataStats) RecordReceived(n int64) { _ = "STUB: not implemented"; return }

func (ds *DataStats) BytesSent() int64 { _ = "STUB: not implemented"; return 0 }

func (ds *DataStats) BytesReceived() int64 { _ = "STUB: not implemented"; return 0 }

type Info struct {
	DataStats
	connKinds   sync.Map
	ClientStats PeerRequestStats
	ServerStats PeerRequestStats
}

func (i *Info) Kind(c network.Conn) Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (i *Info) SetKind(c network.Conn, k Kind) { _ = "STUB: not implemented"; return }

//go:generate mockgen -typed -package=peerinfo -destination=./mocks/mocks.go -source=./peerinfo.go

// PeerInfo provides peer-related connection status and statistics.
type PeerInfo interface {
	// EnsurePeerInfo returns Info structure for the specific peers.
	// If there's no such structure assigned for the peer yet, it creates one. The
	// Info structure will be removed when all connections to the peer are closed.
	EnsurePeerInfo(p peer.ID) *Info
	// RecordReceived records that n bytes of data has been received from peer p
	// via protocol proto.
	RecordReceived(n int64, proto protocol.ID, p peer.ID)
	// RecordReceived records that n bytes of data has been sent to the peer p
	// via protocol proto.
	RecordSent(n int64, proto protocol.ID, p peer.ID)
	// Protocols returns the list of protocols used so far, in no particular order.
	Protocols() []protocol.ID
	// EnsureProtoStats returns DataStats structure for the specified protocol,
	// allocating one if it doesn't exist yet.
	EnsureProtoStats(proto protocol.ID) *DataStats
}

type PeerInfoTracker struct {
	mtx        sync.Mutex
	info       map[peer.ID]*Info
	protoStats map[protocol.ID]*DataStats
	clock      clockwork.Clock
	syncOnce   sync.Once
	stop       context.CancelFunc
	eg         errgroup.Group
}

type Opt func(t *PeerInfoTracker)

func withClock(clock clockwork.Clock) Opt { _ = "STUB: not implemented"; return *new(Opt) }

var _ network.Notifiee = &PeerInfoTracker{}

func NewPeerInfoTracker(opts ...Opt) *PeerInfoTracker { _ = "STUB: not implemented"; return nil }

func (t *PeerInfoTracker) Start(p2pNet network.Network) { _ = "STUB: not implemented"; return }

func (t *PeerInfoTracker) Stop() { _ = "STUB: not implemented"; return }

func (t *PeerInfoTracker) tick(which int) { _ = "STUB: not implemented"; return }

// Connected implements network.Notifiee.
func (t *PeerInfoTracker) Connected(_ network.Network, c network.Conn) {
	_ = "STUB: not implemented"
	return
}

// Disconnected implements network.Notifiee.
func (t *PeerInfoTracker) Disconnected(n network.Network, c network.Conn) {
	_ = "STUB: not implemented"
	return
}

// other connections exist

// Listen implements network.Notifiee.
func (*PeerInfoTracker) Listen(network.Network, ma.Multiaddr) {
	_ = "STUB: not implemented"

	// ListenClose implements network.Notifiee.
	return
}

func (*PeerInfoTracker) ListenClose(network.Network, ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (t *PeerInfoTracker) EnsurePeerInfo(p peer.ID) *Info { _ = "STUB: not implemented"; return nil }

func (t *PeerInfoTracker) EnsureProtoStats(proto protocol.ID) *DataStats {
	_ = "STUB: not implemented"
	return nil
}

func (t *PeerInfoTracker) RecordReceived(n int64, proto protocol.ID, p peer.ID) {
	_ = "STUB: not implemented"
	return
}

func (t *PeerInfoTracker) RecordSent(n int64, proto protocol.ID, p peer.ID) {
	_ = "STUB: not implemented"
	return
}

func (t *PeerInfoTracker) Protocols() []protocol.ID { _ = "STUB: not implemented"; return nil }

type HolePunchTracer struct {
	pi   PeerInfo
	next holepunch.MetricsTracer
}

var _ holepunch.MetricsTracer = &HolePunchTracer{}

func NewHolePunchTracer(pi PeerInfo, next holepunch.MetricsTracer) *HolePunchTracer {
	_ = "STUB: not implemented"
	return nil
}

// DirectDialFinished implements holepunch.MetricsTracer.
func (h *HolePunchTracer) DirectDialFinished(success bool) { _ = "STUB: not implemented"; return }

// HolePunchFinished implements holepunch.MetricsTracer.
func (h *HolePunchTracer) HolePunchFinished(
	side string,
	attemptNum int,
	theirAddrs, ourAddr []ma.Multiaddr,
	directConn network.ConnMultiaddrs,
) {
	_ = "STUB: not implemented"
	return
}
