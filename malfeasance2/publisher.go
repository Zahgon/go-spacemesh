package malfeasance2

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type Publisher struct {
	logger    *zap.Logger
	db        sql.StateDatabase
	sync      syncer
	tortoise  tortoise
	publisher pubsub.Publisher
}

func NewPublisher(
	logger *zap.Logger,
	db sql.StateDatabase,
	sync syncer,
	tortoise tortoise,
	publisher pubsub.Publisher,
) *Publisher {
	_ = "STUB: not implemented"
	return nil
}

func (p *Publisher) PublishATXProof(ctx context.Context, nodeID types.NodeID, proof []byte, allowNoRefATXs bool) error {
	_ = "STUB: not implemented"
	// whether to publish the proof
	return nil
}

// Persisting the proof in the DB has to be done within a transaction to ensure consistency. The ATX handler could
// update the (or merge multiple) marriage set in parallel, so we need to make sure data is consistent while we
// update the malfeasance table.

// smesher is not married

// no ATXs found for this node, but we allow it

// ATX found

// smesher is married

// Combine IDs from the present equivocation set for atx.SmesherID and IDs in atx.Marriages.

// already handled

// all smeshers were already marked as malicious - no gossip to void spamming the network

func (p *Publisher) Regossip(ctx context.Context, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// smesher is not married

// smesher is married

func (p *Publisher) publish(
	ctx context.Context,
	nodeID []types.NodeID,
	refATXs []types.ATXID,
	proof []byte,
	domain ProofDomain,
) error {
	_ = "STUB: not implemented"
	// Only gossip the proof if we are synced (to not spam the network with proofs others probably already have).
	return nil
}

func (p *Publisher) ProofByID(ctx context.Context, nodeID types.NodeID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// smesher is not married

// smesher is married
