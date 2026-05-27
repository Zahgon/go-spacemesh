package blocks

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

var (
	errMalformedData = fmt.Errorf("%w: malformed data", pubsub.ErrValidationReject)
	errWrongHash     = fmt.Errorf("%w: incorrect hash", pubsub.ErrValidationReject)
	errDuplicateTX   = errors.New("duplicate TxID in proposal")
)

// Handler processes Block fetched from peers during sync.
type Handler struct {
	logger *zap.Logger

	fetcher  system.Fetcher
	db       sql.StateDatabase
	tortoise tortoiseProvider
	mesh     meshProvider
}

// Opt for configuring BlockHandler.
type Opt func(*Handler)

// WithLogger defines logger for Handler.
func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// NewHandler creates new Handler.
func NewHandler(
	f system.Fetcher,
	db sql.StateDatabase,
	tortoise tortoiseProvider,
	m meshProvider,
	opts ...Opt,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

// HandleSyncedBlock handles Block data from sync.
func (h *Handler) HandleSyncedBlock(ctx context.Context, expHash types.Hash32, peer p2p.Peer, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// set the block ID when received

// ValidateRewards syntactically validates rewards.
func ValidateRewards(rewards []types.AnyReward) error { _ = "STUB: not implemented"; return nil }

func (h *Handler) checkTransactions(ctx context.Context, peer p2p.Peer, b *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func toAtxIDs(rewards []types.AnyReward) []types.ATXID { _ = "STUB: not implemented"; return nil }
