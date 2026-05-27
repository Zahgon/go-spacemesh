package metrics

import (
	"sync"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/spacemeshos/go-spacemesh/metrics"
)

var (
	totalPeers       = metrics.NewGauge("total_peers", subsystem, "Total number of peers", nil)
	peersPerProtocol = metrics.NewGauge(
		"peers_per_protocol",
		subsystem,
		"Number of peers per protocol",
		[]string{"protocol"},
	)
)

var (
	// ProcessedMessagesDuration in nanoseconds to process a message. Labeled by protocol and result.
	ProcessedMessagesDuration = metrics.NewHistogramWithBuckets(
		"processed_messages_duration",
		subsystem,
		"Duration in nanoseconds to process a message",
		[]string{"protocol", "result"},
		prometheus.ExponentialBuckets(1_000_000, 4, 10),
	)
	deliveredMessagesBytes = metrics.NewCounter(
		"delivered_messages_bytes",
		subsystem,
		"Total amount of delivered payloads (doesn't count gossipsub metadata)",
		[]string{"protocol"},
	)
	receivedMessagesBytes = metrics.NewCounter(
		"received_messages_bytes",
		subsystem,
		"Total amount of received payloads (doesn't count gossipsub metadata)",
		[]string{"protocol"},
	)
	deliveredMessagesCount = metrics.NewCounter(
		"delivered_messages_count",
		subsystem,
		"Total number of delivered messages",
		[]string{"protocol"},
	)
	receivedMessagesCount = metrics.NewCounter(
		"received_messages_count",
		subsystem,
		"Total amount of received messages",
		[]string{"protocol"},
	)
	queueFullCount = metrics.NewCounter(
		"queue_full",
		subsystem,
		"counter of dropped messages because queue was full",
		[]string{"protocol"},
	)
	throttledCount = metrics.NewCounter(
		"throttled",
		subsystem,
		"counter of dropped messages because of throttling",
		[]string{"protocol"},
	)
	rejectedCount = metrics.NewCounter(
		"rejected",
		subsystem,
		"counter of dropped messages for any other reason",
		[]string{"protocol"},
	)
)

// GossipCollector pubsub.RawTracer implementation
// total number of peers
// number of peers per each gossip protocol.
type GossipCollector struct {
	peers struct {
		sync.Mutex
		m map[peer.ID]protocol.ID
	}
}

// NewGoSIPCollector creates a new GossipCollector.
func NewGoSIPCollector() *GossipCollector { _ = "STUB: not implemented"; return nil }

// AddPeer is invoked when a new peer is added.
func (g *GossipCollector) AddPeer(id peer.ID, proto protocol.ID) { _ = "STUB: not implemented"; return }

// RemovePeer is invoked when a peer is removed.
func (g *GossipCollector) RemovePeer(id peer.ID) { _ = "STUB: not implemented"; return }

// Join is invoked when a new topic is joined.
func (g *GossipCollector) Join(string) {
	_ = "STUB: not implemented"

	// Leave is invoked when a topic is abandoned.
	return
}

func (g *GossipCollector) Leave(string) {
	_ = "STUB: not implemented"

	// Graft is invoked when a new peer is grafted on the mesh (gossipsub).
	return
}

func (g *GossipCollector) Graft(peer.ID, string) {
	_ = "STUB: not implemented"

	// Prune is invoked when a peer is pruned from the message (gossipsub).
	return
}

func (g *GossipCollector) Prune(peer.ID, string) {
	_ = "STUB: not implemented"

	// ValidateMessage is invoked when a message first enters the validation pipeline.
	return
}

func (g *GossipCollector) ValidateMessage(msg *pubsub.Message) { _ = "STUB: not implemented"; return }

// DeliverMessage is invoked when a message is delivered.
func (g *GossipCollector) DeliverMessage(msg *pubsub.Message) { _ = "STUB: not implemented"; return }

// RejectMessage is invoked when a message is Rejected or Ignored.
// The reason argument can be one of the named strings Reject*.
func (g *GossipCollector) RejectMessage(msg *pubsub.Message, reason string) {
	_ = "STUB: not implemented"
	return
}

// DuplicateMessage is invoked when a duplicate message is dropped.
func (g *GossipCollector) DuplicateMessage(msg *pubsub.Message) { _ = "STUB: not implemented"; return }

// ThrottlePeer is invoked when a peer is throttled by the peer gater.
func (g *GossipCollector) ThrottlePeer(peer.ID) {
	_ = "STUB: not implemented"

	// RecvRPC is invoked when an incoming RPC is received.
	return
}

func (g *GossipCollector) RecvRPC(*pubsub.RPC) {
	_ = "STUB: not implemented"

	// SendRPC is invoked when a RPC is sent.
	return
}

func (g *GossipCollector) SendRPC(*pubsub.RPC, peer.ID) {
	_ = "STUB: not implemented"

	// DropRPC is invoked when an outbound RPC is dropped, typically because of a queue full.
	return
}

func (g *GossipCollector) DropRPC(*pubsub.RPC, peer.ID) {
	_ = "STUB: not implemented"

	// UndeliverableMessage is invoked when the consumer of Subscribe is not reading messages fast enough and
	// the pressure release mechanism trigger, dropping messages.
	return
}

func (g *GossipCollector) UndeliverableMessage(*pubsub.Message) { _ = "STUB: not implemented"; return }
