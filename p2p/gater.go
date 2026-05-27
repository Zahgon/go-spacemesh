package p2p

import (
	"net"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/control"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

var _ connmgr.ConnectionGater = (*gater)(nil)

func newGater(cfg Config) (*gater, error) {
	_ = "STUB: not implemented"
	// leaves a small room for outbound connections in order to
	// reduce risk of network isolation
	return nil, nil
}

type gater struct {
	h                 host.Host
	inbound, outbound int
	direct            map[peer.ID]struct{}
	ip4blocklist      []*net.IPNet
	ip6blocklist      []*net.IPNet
}

func (g *gater) updateHost(h host.Host) { _ = "STUB: not implemented"; return }

func (g *gater) InterceptPeerDial(pid peer.ID) bool { _ = "STUB: not implemented"; return false }

func (g *gater) InterceptAddrDial(pid peer.ID, m ma.Multiaddr) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *gater) InterceptAccept(n network.ConnMultiaddrs) bool {
	_ = "STUB: not implemented"
	return false
}

func (*gater) InterceptSecured(_ network.Direction, _ peer.ID, _ network.ConnMultiaddrs) bool {
	_ = "STUB: not implemented"
	return false
}

func (*gater) InterceptUpgraded(_ network.Conn) (allow bool, reason control.DisconnectReason) {
	_ = "STUB: not implemented"
	return false, *new(control.DisconnectReason)
}

func (g *gater) allowed(m ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func parseCIDR(cidrs []string) ([]*net.IPNet, error) { _ = "STUB: not implemented"; return nil, nil }

func inAddrRange(ip net.IP, ipnets []*net.IPNet) bool { _ = "STUB: not implemented"; return false }
