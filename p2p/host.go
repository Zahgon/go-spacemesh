package p2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-pubsub/timecache"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/host/autorelay"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/p2p/handshake"
)

// DefaultConfig config.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// 128K

// localhost

// private networks

// link local

// localhost

// ULA reserved

// link local

// Defaults taken from libp2p

const (
	PublicReachability  = "public"
	PrivateReachability = "private"
)

// Config for all things related to p2p layer.
type Config struct {
	DataDir            string        `mapstructure:"-"` // not configurable, overwritten by "`BaseConfig.DataDir()`/p2p"
	LogLevel           zapcore.Level `mapstructure:"-"` // not configurable, overwritten by LoggerConfig.P2PLoggerLevel
	GracePeersShutdown time.Duration `mapstructure:"gracepeersshutdown"`
	MaxMessageSize     int           `mapstructure:"maxmessagesize"`

	// see https://lwn.net/Articles/542629/ for reuseport explanation
	DisableReusePort            bool               `mapstructure:"disable-reuseport"`
	DisableNatPort              bool               `mapstructure:"disable-natport"`
	DisableConnectionManager    bool               `mapstructure:"disable-connection-manager"`
	DisableResourceManager      bool               `mapstructure:"disable-resource-manager"`
	DisableDHT                  bool               `mapstructure:"disable-dht"`
	DisablePubSub               bool               `mapstructure:"disable-pubsub"`
	Flood                       bool               `mapstructure:"flood"`
	Listen                      AddressList        `mapstructure:"listen"`
	Bootnodes                   []string           `mapstructure:"bootnodes"`
	Direct                      []string           `mapstructure:"direct"`
	MinPeers                    int                `mapstructure:"min-peers"`
	LowPeers                    int                `mapstructure:"low-peers"`
	HighPeers                   int                `mapstructure:"high-peers"`
	InboundFraction             float64            `mapstructure:"inbound-fraction"`
	OutboundFraction            float64            `mapstructure:"outbound-fraction"`
	AutoscalePeers              bool               `mapstructure:"autoscale-peers"`
	AdvertiseAddress            AddressList        `mapstructure:"advertise-address"`
	AcceptQueue                 int                `mapstructure:"p2p-accept-queue"`
	Metrics                     bool               `mapstructure:"p2p-metrics"`
	Bootnode                    bool               `mapstructure:"p2p-bootnode"`
	ForceReachability           string             `mapstructure:"p2p-reachability"`
	ForceDHTServer              bool               `mapstructure:"force-dht-server"`
	EnableHolepunching          bool               `mapstructure:"p2p-holepunching"`
	PrivateNetwork              bool               `mapstructure:"p2p-private-network"`
	RelayServer                 RelayServer        `mapstructure:"relay-server"`
	IP4Blocklist                []string           `mapstructure:"ip4-blocklist"`
	IP6Blocklist                []string           `mapstructure:"ip6-blocklist"`
	GossipQueueSize             int                `mapstructure:"gossip-queue-size"`
	GossipPeerOutboundQueueSize int                `mapstructure:"gossip-peer-outbound-queue-size"`
	GossipValidationThrottle    int                `mapstructure:"gossip-validation-throttle"`
	GossipAtxValidationThrottle int                `mapstructure:"gossip-atx-validation-throttle"`
	GossipEvictionStrategy      timecache.Strategy `mapstructure:"gossip-eviction-strategy"`
	PingPeers                   []string           `mapstructure:"ping-peers"`
	PingInterval                time.Duration      `mapstructure:"ping-interval"`
	Relay                       bool               `mapstructure:"relay"`
	StaticRelays                []string           `mapstructure:"static-relays"`
	EnableTCPTransport          bool               `mapstructure:"enable-tcp-transport"`
	EnableQUICTransport         bool               `mapstructure:"enable-quic-transport"`
	EnableRoutingDiscovery      bool               `mapstructure:"enable-routing-discovery"`
	RoutingDiscoveryAdvertise   bool               `mapstructure:"routing-discovery-advertise"`
	DiscoveryTimings            DiscoveryTimings   `mapstructure:"discovery-timings"`
	AutoNATServer               AutoNATServer      `mapstructure:"auto-nat-server"`
}

type DiscoveryTimings struct {
	AdvertiseDelay          time.Duration `mapstructure:"advertise-delay"`
	AdvertiseInterval       time.Duration `mapstructure:"advertise-interval"`
	AdvertiseIntervalSpread time.Duration `mapstructure:"advertise-interval-spread"`
	AdvertiseRetryDelay     time.Duration `mapstructure:"advertise-retry-delay"`
	FindPeersRetryDelay     time.Duration `mapstructure:"find-peers-retry-delay"`
	MinBackoff              time.Duration `mapstructure:"min-backoff"`
	MaxBackoff              time.Duration `mapstructure:"max-backoff"`
	MinConnBackoff          time.Duration `mapstructure:"min-conn-backoff"`
	MaxConnBackoff          time.Duration `mapstructure:"max-conn-backoff"`
	DialTimeout             time.Duration `mapstructure:"dial-timeout"`
}

type AutoNATServer struct {
	GlobalMax   int           `mapstructure:"global-max"`
	PeerMax     int           `mapstructure:"peer-max"`
	ResetPeriod time.Duration `mapstructure:"reset-period"`
}

type DeprecatedMaxReservationsPerPeer struct{}

// DeprecatedMsg implements Deprecated interface.
func (DeprecatedMaxReservationsPerPeer) DeprecatedMsg() string {
	_ = "STUB: not implemented"
	return ""
}

type RelayServer struct {
	Enable                 bool                             `mapstructure:"enable"`
	Reservations           int                              `mapstructure:"reservations"`
	TTL                    time.Duration                    `mapstructure:"ttl"`
	ConnDurationLimit      time.Duration                    `mapstructure:"conn-duration-limit"`
	ConnDataLimit          int64                            `mapstructure:"conn-data-limit"`
	MaxCircuits            int                              `mapstructure:"max-circuits"`
	BufferSize             int                              `mapstructure:"buffer-size"`
	MaxReservationsPerPeer DeprecatedMaxReservationsPerPeer `mapstructure:"max-reservations-per-peer"`
	MaxReservationsPerIP   int                              `mapstructure:"max-reservations-per-ip"`
	MaxReservationsPerASN  int                              `mapstructure:"max-reservations-per-asn"`
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// New initializes libp2p host configured for spacemesh.
func New(
	logger *zap.Logger,
	cfg Config,
	prologue []byte,
	quicNetCookie handshake.NetworkCookie,
	opts ...Opt,
) (*Host, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// should be NOT exposed in the config

// leaves a small room for outbound connections in order to
// reduce risk of network isolation

// Obtain the IDService via fx dependency injection.
// This function is always called by libp2p.New().

// TODO(dshulyak) this is small mess. refactor to avoid this patching
// both New and Upgrade should use options.

// AutoStart initializes a new host and starts it.
func AutoStart(ctx context.Context,
	logger *zap.Logger,
	cfg Config,
	prologue []byte,
	quicNetCookie handshake.NetworkCookie,
	opts ...Opt,
) (*Host, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setupResourcesManager(hostcfg Config) func(cfg *libp2p.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func parseIntoAddr(nodes []string) ([]peer.AddrInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func relayPeerSource(logger *zap.Logger) (autorelay.PeerSource, chan<- peer.AddrInfo) {
	_ = "STUB: not implemented"
	return *new(autorelay.PeerSource), nil
}
