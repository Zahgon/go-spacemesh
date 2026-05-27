package txs

import (
	"context"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
)

var (
	errWrongHash   = fmt.Errorf("%w: incorrect hash", pubsub.ErrValidationReject)
	errDuplicateTX = errors.New("tx already exists")
	errParse       = errors.New("failed to parse tx")
	errVerify      = errors.New("failed to verify tx")
)

// TxHandler handles the transactions received via gossip or sync.
type TxHandler struct {
	self   peer.ID
	logger *zap.Logger
	state  conservativeState
}

// NewTxHandler returns a new TxHandler.
func NewTxHandler(s conservativeState, id peer.ID, l *zap.Logger) *TxHandler {
	_ = "STUB: not implemented"
	return nil
}

func updateMetrics(err error, counter *prometheus.CounterVec) { _ = "STUB: not implemented"; return }

// HandleGossipTransaction handles data received on the transactions gossip channel.
func (th *TxHandler) HandleGossipTransaction(ctx context.Context, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleProposalTransaction handles data received on the transactions synced as a part of proposal.
func (th *TxHandler) HandleProposalTransaction(
	ctx context.Context,
	expHash types.Hash32,
	_ p2p.Peer,
	msg []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (th *TxHandler) VerifyAndCacheTx(ctx context.Context, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (th *TxHandler) verifyAndCache(ctx context.Context, expHash types.Hash32, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleBlockTransaction handles transactions received as a reference to a block.
func (th *TxHandler) HandleBlockTransaction(_ context.Context, expHash types.Hash32, _ p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
