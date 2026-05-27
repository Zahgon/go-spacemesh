package proposals

import (
	"context"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/system"
)

// Validator validates the eligibility of a Ballot.
// The validation focuses on eligibility only and assumes the Ballot to be valid otherwise.
type Validator struct {
	minActiveSetWeight []types.EpochMinimalActiveWeight
	avgLayerSize       uint32
	layersPerEpoch     uint32
	tortoise           tortoiseProvider
	atxsdata           *atxsdata.Data
	clock              layerClock
	beacons            system.BeaconCollector
	logger             *zap.Logger
	vrfVerifier        vrfVerifier
}

// ValidatorOpt for configuring Validator.
type ValidatorOpt func(h *Validator)

// NewEligibilityValidator returns a new EligibilityValidator.
func NewEligibilityValidator(
	avgLayerSize, layersPerEpoch uint32,
	minActiveSetWeight []types.EpochMinimalActiveWeight,
	clock layerClock,
	tortoise tortoiseProvider,
	atxsdata *atxsdata.Data,
	bc system.BeaconCollector,
	lg *zap.Logger,
	vrfVerifier vrfVerifier,
	opts ...ValidatorOpt,
) *Validator {
	_ = "STUB: not implemented"
	return nil
}

// CheckEligibility checks that a ballot is eligible in the layer that it specifies.
func (v *Validator) CheckEligibility(ctx context.Context, ballot *types.Ballot, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// validateReference executed for reference ballots in latest epoch.
func (v *Validator) validateReference(
	ballot *types.Ballot,
	weight, totalWeight uint64,
) (*types.EpochData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateSecondary executed for non-reference ballots in latest epoch and all ballots in past epochs.
func (v *Validator) validateSecondary(ballot *types.Ballot) (*types.EpochData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
