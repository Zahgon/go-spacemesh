package hare3

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
)

type oracle interface {
	Validate(context.Context, types.LayerID, uint32, int, types.NodeID, types.VrfSignature, uint16) (bool, error)
	CalcEligibility(context.Context, types.LayerID, uint32, int, types.NodeID, types.VrfSignature) (uint16, error)
}

type legacyOracle struct {
	log    *zap.Logger
	oracle oracle
	config Config
}

func (lg *legacyOracle) validate(msg *Message) grade { _ = "STUB: not implemented"; return *new(grade) }

func (lg *legacyOracle) active(
	signer *signing.EdSigner,
	beacon types.Beacon,
	layer types.LayerID,
	ir IterRound,
) *types.HareEligibility {
	_ = "STUB: not implemented"
	return nil
}
