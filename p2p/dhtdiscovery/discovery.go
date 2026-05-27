package discovery

import (
	"context"
	"sync"
	"time"

	levelds "github.com/ipfs/go-ds-leveldb"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	p2pdisc "github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	backoff "github.com/libp2p/go-libp2p/p2p/discovery/backoff"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	discoveryNS         = "spacemesh-disc"
	relayNS             = "spacemesh-disc-relay"
	protocolPrefix      = "/spacekad"
	ProtocolID          = protocolPrefix + "/kad/1.0.0"
	advRecheckInterval  = time.Minute
	discRecheckInterval = 10 * time.Second
)

type Opt func(*Discovery)

func WithPeriod(period time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithTimeout(timeout time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithBootstrapDuration(bootstrapDuration time.Duration) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithMinPeers(minPeers int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithBootnodes(bootnodes []peer.AddrInfo) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithBackup(backup []peer.AddrInfo) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithMode(mode dht.ModeOpt) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func Private() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithDir(path string) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func DisableDHT() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithRelayCandidateChannel(relayCh chan<- peer.AddrInfo) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func EnableRoutingDiscovery() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func AdvertiseForPeerDiscovery() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithAdvertiseRetryDelay(value time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithAdvertiseDelay(value time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithAdvertiseInterval(value time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithAdvertiseIntervalSpread(value time.Duration) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithFindPeersRetryDelay(value time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithDiscoveryBackoff(backoff bool) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithDiscoveryBackoffTimings(minBackoff, maxBackoff time.Duration) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithConnBackoffTimings(minBackoff, maxBackoff, dialTimeout time.Duration) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

type DiscoveryHost interface {
	host.Host
	NeedPeerDiscovery() bool
	HaveRelay() bool
}

func New(h DiscoveryHost, opts ...Opt) (*Discovery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Discovery struct {
	public                 bool
	mode                   dht.ModeOpt
	disableDht             bool
	dir                    string
	relayCh                chan<- peer.AddrInfo
	enableRoutingDiscovery bool
	advertise              bool

	logger *zap.Logger
	eg     errgroup.Group
	cancel context.CancelFunc

	h         DiscoveryHost
	dhtLock   sync.Mutex
	dht       *dht.IpfsDHT
	datastore *levelds.Datastore
	disc      p2pdisc.Discovery

	// how often to check if we have enough peers
	period time.Duration
	// timeout used for connections
	timeout                 time.Duration
	bootstrapDuration       time.Duration
	minPeers                int
	backup, bootnodes       []peer.AddrInfo
	advertiseDelay          time.Duration
	advertiseInterval       time.Duration
	advertiseIntervalSpread time.Duration
	advertiseRetryDelay     time.Duration
	findPeersRetryDelay     time.Duration
	minBackoff              time.Duration
	maxBackoff              time.Duration
	minConnBackoff          time.Duration
	maxConnBackoff          time.Duration
	dialTimeout             time.Duration
	discBackoff             bool
	connBackoff             *backoff.BackoffConnector
}

func (d *Discovery) FindPeer(ctx context.Context, p peer.ID) (peer.AddrInfo, error) {
	_ = "STUB: not implemented"
	return *new(peer.AddrInfo), nil
}

func (d *Discovery) Start() error { _ = "STUB: not implemented"; return nil }

func (d *Discovery) Stop() { _ = "STUB: not implemented"; return }

func (d *Discovery) bootstrap(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *Discovery) connect(ctx context.Context, eg *errgroup.Group, nodes []peer.AddrInfo) {
	_ = "STUB: not implemented"
	return
}

func (d *Discovery) setupDHT(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *Discovery) ensureAtLeastMinPeers(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// trigger bootstrap when node starts immediately
// TODO: connectedF, disconnectedF: track enough rendezvous peers

// once got enough peers no need to keep, they are either already connected or unavailable

// no reason to spend more resources if we got enough from backup

func (d *Discovery) advInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (d *Discovery) advertiseNS(ctx context.Context, ns string, active func() bool) error {
	_ = "STUB: not implemented"
	return nil
}

// At this moment, advertisement is not needed.
// There was previous delay already, just need to recheck if we need
// to start advertising again

func (d *Discovery) discoverPeers(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *Discovery) discoverRelays(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *Discovery) findPeersContinuously(
	ctx context.Context,
	ns string,
	active func() bool,
) <-chan peer.AddrInfo {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discovery) findPeers(
	ctx context.Context,
	ns string,
	active func() bool,
	checkInterval time.Duration,
	out chan<- peer.AddrInfo,
) (cont bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Note: ticker is not started if active is nil.

// skip self

type rngSource struct{}

func (r rngSource) Int63() int64 { _ = "STUB: not implemented"; return 0 }

func (r rngSource) Seed(seed int64) { _ = "STUB: not implemented"; return }
