package malfeasance

import (
	"context"
	"errors"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
)

var (
	errKnownProof    = errors.New("known proof")
	errMalformedData = fmt.Errorf("%w: malformed data", pubsub.ErrValidationReject)
	errWrongHash     = fmt.Errorf("%w: incorrect hash", pubsub.ErrValidationReject)
	errUnknownProof  = fmt.Errorf("%w: unknown proof type", pubsub.ErrValidationReject)
	errInvalidProof  = fmt.Errorf("%w: invalid proof", pubsub.ErrValidationReject)
)

type MalfeasanceType byte

const (
	MultipleATXs     MalfeasanceType = MalfeasanceType(wire.MultipleATXs)
	MultipleBallots                  = MalfeasanceType(wire.MultipleBallots)
	HareEquivocation                 = MalfeasanceType(wire.HareEquivocation)
	InvalidPostIndex                 = MalfeasanceType(wire.InvalidPostIndex)
	InvalidPrevATX                   = MalfeasanceType(wire.InvalidPrevATX)
)

// Handler processes MalfeasanceProof from gossip and, if deems it valid, propagates it to peers.
type Handler struct {
	logger   *zap.Logger
	cdb      *datastore.CachedDB
	self     p2p.Peer
	nodeIDs  []types.NodeID
	tortoise tortoise

	handlers map[MalfeasanceType]MalfeasanceHandler

	// metrics
	numProofs        *prometheus.CounterVec
	numInvalidProofs *prometheus.CounterVec
	numMalformed     prometheus.Counter
}

func NewHandler(
	cdb *datastore.CachedDB,
	lg *zap.Logger,
	self p2p.Peer,
	nodeIDs []types.NodeID,
	tortoise tortoise,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) RegisterHandler(malfeasanceType MalfeasanceType, handler MalfeasanceHandler) {
	_ = "STUB: not implemented"
	return
}

func (h *Handler) reportMalfeasance(smesher types.NodeID) { _ = "STUB: not implemented"; return }

func (h *Handler) countProof(mp *wire.MalfeasanceProof) { _ = "STUB: not implemented"; return }

func (h *Handler) countInvalidProof(p *wire.MalfeasanceProof) { _ = "STUB: not implemented"; return }

func (h *Handler) Info(ctx context.Context, nodeID types.NodeID) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for malfeasance V1 there are no domains

// HandleSynced is the sync validator for MalfeasanceProof.
func (h *Handler) HandleSynced(
	ctx context.Context,
	expHash types.Hash32,
	peer p2p.Peer,
	data []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// we log & return because libp2p will ignore the message if we return an error,
// but only log "validation ignored" instead of the error we return

// HandleGossip is the gossip receiver for MalfeasanceGossip.
func (h *Handler) HandleGossip(ctx context.Context, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// node saves malfeasance proof eagerly/atomically with the malicious data.
// it has validated the proof before saving to db.

func (h *Handler) validateAndSave(ctx context.Context, p *wire.MalfeasanceProof) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

func (h *Handler) Validate(ctx context.Context, p *wire.MalfeasanceProof) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}
