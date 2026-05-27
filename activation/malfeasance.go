package activation

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	multiATXs        = "atx"
	invalidPostIndex = "invalid_post_index"
	invalidPrevATX   = "invalid_prev_atx"
)

type MalfeasanceHandler struct {
	logger *zap.Logger
	db     sql.Executor

	edVerifier *signing.EdVerifier
}

func NewMalfeasanceHandler(
	db sql.Executor,
	logger *zap.Logger,
	edVerifier *signing.EdVerifier,
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

type InvalidPostIndexHandler struct {
	db sql.Executor

	edVerifier *signing.EdVerifier
	validator  nipostValidatorV1
}

func NewInvalidPostIndexHandler(
	db sql.Executor,
	edVerifier *signing.EdVerifier,
	validator nipostValidatorV1,
) *InvalidPostIndexHandler {
	_ = "STUB: not implemented"
	return nil
}

func (mh *InvalidPostIndexHandler) Info(data wire.ProofData) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mh *InvalidPostIndexHandler) Validate(ctx context.Context, data wire.ProofData) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

func (mh *InvalidPostIndexHandler) ReportLabel() string { _ = "STUB: not implemented"; return "" }

type InvalidPrevATXHandler struct {
	db sql.Executor

	edVerifier *signing.EdVerifier
}

func NewInvalidPrevATXHandler(db sql.Executor, edVerifier *signing.EdVerifier) *InvalidPrevATXHandler {
	_ = "STUB: not implemented"
	return nil
}

func (mh *InvalidPrevATXHandler) Info(data wire.ProofData) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mh *InvalidPrevATXHandler) Validate(ctx context.Context, data wire.ProofData) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

func (mh *InvalidPrevATXHandler) ReportLabel() string { _ = "STUB: not implemented"; return "" }
