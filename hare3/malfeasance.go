package hare3

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	hareEquivocate = "hare_eq"
)

type MalfeasanceHandler struct {
	logger *zap.Logger
	db     sql.Executor

	edVerifier *signing.EdVerifier
}

type MalfeasanceOpt func(*MalfeasanceHandler)

func WithMalfeasanceLogger(logger *zap.Logger) MalfeasanceOpt {
	_ = "STUB: not implemented"
	return *new(MalfeasanceOpt)
}

func NewMalfeasanceHandler(
	db sql.Executor,
	edVerifier *signing.EdVerifier,
	opt ...MalfeasanceOpt,
) *MalfeasanceHandler {
	_ = "STUB: not implemented"
	return nil
}

func (mh *MalfeasanceHandler) Info(data wire.ProofData) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mh *MalfeasanceHandler) Validate(ctx context.Context, data wire.ProofData) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

func (mh *MalfeasanceHandler) ReportLabel() string { _ = "STUB: not implemented"; return "" }
