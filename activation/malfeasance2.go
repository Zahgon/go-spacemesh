package activation

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/activation/wire"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type MalfeasanceHandlerV2 struct {
	logger *zap.Logger
	db     sql.StateDatabase

	malPublisher malfeasancePublisher
	edVerifier   *signing.EdVerifier
	validator    nipostValidatorV2

	signersMtx sync.Mutex
	signers    map[types.NodeID]*signing.EdSigner
}

func NewMalfeasanceHandlerV2(
	logger *zap.Logger,
	db sql.StateDatabase,
	malPublisher malfeasancePublisher,
	edVerifier *signing.EdVerifier,
	validator nipostValidatorV2,
) *MalfeasanceHandlerV2 {
	_ = "STUB: not implemented"
	return nil
}

func (p *MalfeasanceHandlerV2) Register(sig *signing.EdSigner) { _ = "STUB: not implemented"; return }

// Publish publishes an ATX proof by encoding it and sending it to the malfeasance publisher.
func (p *MalfeasanceHandlerV2) Publish(ctx context.Context, nodeID types.NodeID, proof wire.Proof) error {
	_ = "STUB: not implemented"
	return nil
}

// do not publish proofs against one self

// for now we only have one version

func (p *MalfeasanceHandlerV2) Regossip(ctx context.Context, nodeID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// do not publish proofs against one self

func (mh *MalfeasanceHandlerV2) decodeProof(data []byte) (wire.Proof, error) {
	_ = "STUB: not implemented"
	return *new(wire.Proof), nil
}

func (mh *MalfeasanceHandlerV2) Validate(ctx context.Context, data []byte) (types.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeID), nil
}

func (mh *MalfeasanceHandlerV2) Info(data []byte) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mh *MalfeasanceHandlerV2) ReportLabels(data []byte) []string {
	_ = "STUB: not implemented"
	return nil
}

func (mh *MalfeasanceHandlerV2) PostIndex(
	ctx context.Context,
	smesherID types.NodeID,
	commitment types.ATXID,
	post *types.Post,
	challenge []byte,
	numUnits uint32,
	idx int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (mh *MalfeasanceHandlerV2) Signature(d signing.Domain, nodeID types.NodeID, m []byte, sig types.EdSignature) bool {
	_ = "STUB: not implemented"
	return false
}

func (mh *MalfeasanceHandlerV2) IdentityExists(nodeID types.NodeID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
