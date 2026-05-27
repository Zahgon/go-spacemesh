package malfeasance

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
)

type Publisher struct {
	logger    *zap.Logger
	cdb       *datastore.CachedDB
	tortoise  tortoise
	sync      syncer
	publisher pubsub.Publisher
}

func NewPublisher(
	logger *zap.Logger,
	cdb *datastore.CachedDB,
	sync syncer,
	tortoise tortoise,
	publisher pubsub.Publisher,
) *Publisher {
	_ = "STUB: not implemented"
	return nil
}

// Publishes a malfeasance proof to the network.
func (p *Publisher) PublishProof(ctx context.Context, smesherID types.NodeID, proof *wire.MalfeasanceProof) error {
	_ = "STUB: not implemented"
	return nil
}

// Only gossip the proof if we are synced (to not spam the network with proofs others probably already have).
