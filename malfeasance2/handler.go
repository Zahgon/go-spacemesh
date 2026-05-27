package malfeasance2

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

var (
	ErrMalformedData  = fmt.Errorf("%w: malformed data", pubsub.ErrValidationReject)
	ErrWrongHash      = fmt.Errorf("%w: incorrect hash", pubsub.ErrValidationReject)
	ErrUnknownVersion = fmt.Errorf("%w: unknown version", pubsub.ErrValidationReject)
	ErrUnknownDomain  = fmt.Errorf("%w: unknown domain", pubsub.ErrValidationReject)
)

type Handler struct {
	logger   *zap.Logger
	db       sql.StateDatabase
	self     p2p.Peer
	nodeIDs  []types.NodeID
	fetcher  system.Fetcher
	tortoise tortoise

	handlers map[ProofDomain]MalfeasanceHandler

	// metrics
	numProofs        *prometheus.CounterVec
	numInvalidProofs *prometheus.CounterVec
	numMalformed     prometheus.Counter
}

func NewHandler(
	db sql.StateDatabase,
	lg *zap.Logger,
	self p2p.Peer,
	nodeIDs []types.NodeID,
	fetcher system.Fetcher,
	tortoise tortoise,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) RegisterHandler(malfeasanceType ProofDomain, handler MalfeasanceHandler) {
	_ = "STUB: not implemented"
	return
}

func (h *Handler) countProof(mp MalfeasanceProof) { _ = "STUB: not implemented"; return }

func (h *Handler) countInvalidProof(mp MalfeasanceProof) { _ = "STUB: not implemented"; return }

func (h *Handler) reportMalfeasance(smesher types.NodeID) { _ = "STUB: not implemented"; return }

func (h *Handler) Info(ctx context.Context, nodeID types.NodeID) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Handler) HandleSynced(ctx context.Context, expHash types.Hash32, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// we log & return because libp2p will ignore the message if we return an error,
// but only log "validation ignored" instead of the error we return

func (h *Handler) HandleGossip(ctx context.Context, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil

	// ignore messages from self, we already validate and persist proofs when publishing
}

func (h *Handler) handleProof(ctx context.Context, peer p2p.Peer, proof MalfeasanceProof) ([]types.NodeID, error) {
	_ = "STUB: not implemented"
	return nil,

		// unsupported proof version
		nil
}

// unknown proof domain

// smesher is not married

// ensure that the ID for which the proof was created is the first in the equivocation set

func (h *Handler) fetchReferences(ctx context.Context, peer p2p.Peer, atxIDs []types.ATXID) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) storeProof(ctx context.Context, nodeIDs []types.NodeID, proof []byte, domain ProofDomain) error {
	_ = "STUB: not implemented"
	// Persisting the proof in the DB has to be done within a transaction to ensure consistency. The ATX handler could
	// update the (or merge multiple) marriage set in parallel, so we need to make sure data is consistent while we
	// update the malfeasance table.
	return nil
}

// smesher is not married
