package p2p

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/protocol/identify"
	ma "github.com/multiformats/go-multiaddr"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	discovery "github.com/spacemeshos/go-spacemesh/p2p/dhtdiscovery"
	"github.com/spacemeshos/go-spacemesh/p2p/peerinfo"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
)

// Opt is for configuring Host.
type Opt func(fh *Host)

// WithLog configures logger for Host.
func WithLog(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithConfig sets Config for Host.
func WithConfig(cfg Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithNodeReporter updates reporter that is notified every time when
// node added or removed a peer.
func WithNodeReporter(reporter func() error) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithDirectNodes(direct map[peer.ID]struct{}) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithBootnodes(bootnodes map[peer.ID]struct{}) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithRelayCandidateChannel(relayCh chan<- peer.AddrInfo) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithPeerInfo(pi peerinfo.PeerInfo) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithIDService(idService identify.IDService) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// Host is a convenience wrapper for all p2p related functionality required to run
// a full spacemesh node.
type Host struct {
	eg     errgroup.Group
	ctx    context.Context
	cancel context.CancelFunc

	cfg    Config
	logger *zap.Logger

	closed struct {
		sync.Mutex
		closed bool
	}

	host.Host
	peerInfo peerinfo.PeerInfo
	pubsub.PubSub

	nodeReporter func() error

	discovery        *discovery.Discovery
	direct, bootnode map[peer.ID]struct{}
	relayCh          chan<- peer.AddrInfo

	natTypeSub event.Subscription
	natType    struct {
		sync.Mutex
		udpNATType network.NATDeviceType
		tcpNATType network.NATDeviceType
	}
	reachSub     event.Subscription
	reachability struct {
		sync.Mutex
		value network.Reachability
	}

	ping      *Ping
	idService identify.IDService
}

// Upgrade creates Host instance from host.Host.
func Upgrade(h host.Host, opts ...Opt) (*Host, error) { _ = "STUB: not implemented"; return nil, nil }

// If no IDService is provided, which may be the case in the tests,
// we can try to get it from the host, assuming it's a *basichost.BasicHost.
// *basichost.BasicHost is expected when libp2p mocknet is being used
// instead of libp2p.New().

// TBD: also protect ping

// GetPeers returns connected peers.
func (fh *Host) GetPeers() []Peer { _ = "STUB: not implemented"; return nil }

// Connected returns true if the specified peer is connected.
// Peers that only have transient connections to them aren't considered connected.
func (fh *Host) Connected(p Peer) bool { _ = "STUB: not implemented"; return false }

// ConnectedPeerInfo retrieves a peer info object for the given peer.ID, if the
// given peer is not connected then nil is returned.
func (fh *Host) ConnectedPeerInfo(id peer.ID) *PeerInfo { _ = "STUB: not implemented"; return nil }

// there's no sync between  Peers() and ConnsToPeer() so by the time we
// try to get the conns they may not exist.

// ProtocolDataStats returns per-protocol data stats.
func (fh *Host) ProtocolDataStats() map[protocol.ID]*peerinfo.DataStats {
	_ = "STUB: not implemented"
	return nil
}

// ListenAddresses returns the addresses on which this host listens.
func (fh *Host) ListenAddresses() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

// KnownAddresses returns the addresses by which the peers know this one.
func (fh *Host) KnownAddresses() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

// NATDeviceType returns NATDeviceType returns the NAT device types for
// UDP and TCP so far for this host.
func (fh *Host) NATDeviceType() (udpNATType, tcpNATType network.NATDeviceType) {
	_ = "STUB: not implemented"
	return *new(network.NATDeviceType), *new(network.NATDeviceType)
}

// Reachability returns reachability of the host (public, private, unknown).
func (fh *Host) Reachability() network.Reachability {
	_ = "STUB: not implemented"
	return *new(network.Reachability)
}

// DHTServerEnabled returns true if the server has DHT running in server mode.
func (fh *Host) DHTServerEnabled() bool { _ = "STUB: not implemented"; return false }

// NeedPeerDiscovery returns true if it makes sense to do additional
// discovery of non-DHT (NATed) peers.
func (fh *Host) NeedPeerDiscovery() bool {
	_ = "STUB: not implemented"
	// Once we get LowPeers, the discovery mechanism is no longer
	// needed
	return false
}

// Check if this is a public-reachable node which can reach
// nodes behind Cone NAT

// Check if we have Cone NAT for either TCP or UDP. If so,
// hole punching should work for other NATed nodes. Also, in
// case of an unknown NAT type, assume there's chance at hole
// punching

// Symmetric NAT for both TCP and UDP, hole punching will not
// work so we're not looking for NATed peers. Will only
// connect to the nodes with DHT Server mode

// HaveRelay returns true if this host can be used as a relay, that
// is, it supports relay service and has public reachability.
func (fh *Host) HaveRelay() bool { _ = "STUB: not implemented"; return false }

// PeerCount returns number of connected peers.
func (fh *Host) PeerCount() uint64 { _ = "STUB: not implemented"; return 0 }

// PeerProtocols returns the protocols supported by peer.
func (fh *Host) PeerProtocols(p Peer) ([]protocol.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ping returns Ping structure for this Host, if any PingPeers are
// specified in the config. Otherwise, it returns nil.
func (fh *Host) Ping() *Ping { _ = "STUB: not implemented"; return nil }

func (fh *Host) Start() error { _ = "STUB: not implemented"; return nil }

// Stop background workers and release external resources.
func (fh *Host) Stop() error { _ = "STUB: not implemented"; return nil }

func (fh *Host) trackNetEvents() error { _ = "STUB: not implemented"; return nil }

func (fh *Host) PeerInfo() peerinfo.PeerInfo {
	_ = "STUB: not implemented"
	return *

	// Identify ensures that the given peer is identified via libp2p identify protocol.
	// Identification is initiated after connecting to the peer, and the set of protocols for
	// the peer in the ProtoBook is not guaranteed to be correct until identification
	// finishes.
	// Note that the set of the protocols in the ProtoBook for a particular peer may also
	// change via a push identity notification when the peer adds a new handler via
	// SetStreamHandler (e.g. sets up a new Server).
	new(peerinfo.PeerInfo)
}

func (fh *Host) Identify(ctx context.Context, p peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// IDService.IdentifyConn is a no-op if the connection is already
// identified, but otherwise we need to wait for identification to finish
// to have proper set of protocols.
