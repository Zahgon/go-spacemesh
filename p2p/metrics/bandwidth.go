package metrics

import (
	"github.com/libp2p/go-libp2p/core/metrics"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"

	prometheusMetrics "github.com/spacemeshos/go-spacemesh/metrics"
	"github.com/spacemeshos/go-spacemesh/p2p/peerinfo"
)

const (
	incoming = "incoming"
	outgoing = "outgoing"
	// subsystem shared by all metrics exposed by this package.
	subsystem = "p2p"
)

var (
	totalIn             = prometheusMetrics.NewCounter("total_send", subsystem, "Total bytes sent", nil)
	totalOut            = prometheusMetrics.NewCounter("total_recv", subsystem, "Total bytes received", nil)
	messagesPerProtocol = prometheusMetrics.NewCounter(
		"messages_per_protocol",
		subsystem,
		"Number of messages sent per protocol",
		[]string{"protocol", "direction"},
	)
	trafficPerProtocol = prometheusMetrics.NewCounter(
		"traffic_per_protocol",
		subsystem,
		"Traffic sent per protocol",
		[]string{"protocol", "direction"},
	)
)

// BandwidthCollector implement metrics.Reporter
// that keeps track of the number of messages sent and received per protocol.
type BandwidthCollector struct {
	pi peerinfo.PeerInfo
}

// NewBandwidthCollector creates a new BandwidthCollector.
func NewBandwidthCollector(pi peerinfo.PeerInfo) *BandwidthCollector {
	_ = "STUB: not implemented"
	return nil
}

// LogSentMessageStream logs the message node sent to the peer.
func (b *BandwidthCollector) LogSentMessageStream(size int64, proto protocol.ID, p peer.ID) {
	_ = "STUB: not implemented"
	return
}

// LogRecvMessageStream logs the message that node received from the peer.
func (b *BandwidthCollector) LogRecvMessageStream(size int64, proto protocol.ID, p peer.ID) {
	_ = "STUB: not implemented"
	return
}

// LogSentMessage  logs the message sent to the peer.
func (b *BandwidthCollector) LogSentMessage(int64) {
	_ = "STUB: not implemented"

	// LogRecvMessage logs the message received from the peer.
	return
}

func (b *BandwidthCollector) LogRecvMessage(int64) {
	_ = "STUB: not implemented"

	// GetBandwidthForPeer mock returns the bandwidth for a given peer.
	return
}

func (b *BandwidthCollector) GetBandwidthForPeer(peer.ID) metrics.Stats {
	_ = "STUB: not implemented"
	return *

	// GetBandwidthForProtocol mock returns the bandwidth for a given protocol.
	new(metrics.Stats)
}

func (b *BandwidthCollector) GetBandwidthForProtocol(protocol.ID) metrics.Stats {
	_ = "STUB: not implemented"
	return *

	// GetBandwidthTotals returns mock the total bandwidth used by the node.
	new(metrics.Stats)
}

func (b *BandwidthCollector) GetBandwidthTotals() metrics.Stats {
	_ = "STUB: not implemented"
	return *

	// GetBandwidthByPeer mock returns the bandwidth for a given peer.
	new(metrics.Stats)
}

func (b *BandwidthCollector) GetBandwidthByPeer() map[peer.ID]metrics.Stats {
	_ = "STUB: not implemented"

	// GetBandwidthByProtocol mock returns the bandwidth for a given protocol.
	return nil
}

func (b *BandwidthCollector) GetBandwidthByProtocol() map[protocol.ID]metrics.Stats {
	_ = "STUB: not implemented"
	return nil
}
