package pubsub

import (
	"context"
	"sync"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"go.uber.org/zap"
)

type PubSub interface {
	Register(topic string, handler GossipHandler, opts ...ValidatorOpt)
	Publish(ctx context.Context, topic string, msg []byte) error
	ProtocolPeers(protocol string) []peer.ID
}

type NullPubSub struct{}

var _ PubSub = &NullPubSub{}

// Register implements PubSub.
func (*NullPubSub) Register(
	topic string,
	handler func(context.Context, peer.ID, []byte) error,
	opts ...pubsub.ValidatorOpt,
) {
	_ = "STUB: not implemented"

	// Publish implements PubSub.
	return
}

func (*NullPubSub) Publish(ctx context.Context, topic string, msg []byte) error {
	_ = "STUB: not implemented"

	// ProtocolPeers implements PubSub.
	return nil
}

func (*NullPubSub) ProtocolPeers(protocol string) []peer.ID {
	_ = "STUB: not implemented"

	// GossipPubSub is a spacemesh-specific wrapper around gossip protocol.
	return nil
}

type GossipPubSub struct {
	logger *zap.Logger
	pubsub *pubsub.PubSub
	host   host.Host

	mu     sync.RWMutex
	topics map[string]*pubsub.Topic
}

var _ PubSub = &GossipPubSub{}

// Register handler for topic.
func (ps *GossipPubSub) Register(topic string, handler GossipHandler, opts ...ValidatorOpt) {
	_ = "STUB: not implemented"
	return
}

// Drop peers on ValidationRejectErr

// Publish message to the topic.
func (ps *GossipPubSub) Publish(ctx context.Context, topic string, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// ProtocolPeers returns list of peers that are communicating in a given protocol.
func (ps *GossipPubSub) ProtocolPeers(protocol string) []peer.ID {
	_ = "STUB: not implemented"
	return nil
}
