package beacon

import (
	"math/big"
	"testing"
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// Config is the configuration of the beacon.
type Config struct {
	// Security parameter (for calculating ATX threshold)
	Kappa int `mapstructure:"beacon-kappa"`
	// Ratio of dishonest spacetime (for calculating ATX threshold). Should be a string representing a rational number.
	Q big.Rat `mapstructure:"beacon-q"`
	// Amount of rounds in every epoch
	RoundsNumber types.RoundID `mapstructure:"beacon-rounds-number"`
	// Grace period duration
	GracePeriodDuration time.Duration `mapstructure:"beacon-grace-period-duration"`
	// Proposal phase duration
	ProposalDuration time.Duration `mapstructure:"beacon-proposal-duration"`
	// First voting round duration
	FirstVotingRoundDuration time.Duration `mapstructure:"beacon-first-voting-round-duration"`
	// Voting round duration
	VotingRoundDuration time.Duration `mapstructure:"beacon-voting-round-duration"`
	// Weak coin round duration
	WeakCoinRoundDuration time.Duration `mapstructure:"beacon-weak-coin-round-duration"`
	// Ratio of votes for reaching consensus
	Theta big.Rat `mapstructure:"beacon-theta"`
	// Maximum allowed number of votes to be sent
	VotesLimit uint32 `mapstructure:"beacon-votes-limit"`
	// Numbers of layers to wait before determining beacon values from ballots when the node didn't participate
	// in previous epoch.
	BeaconSyncWeightUnits int `mapstructure:"beacon-sync-weight-units"`
}

// DefaultConfig returns the default configuration for the beacon.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// TODO: around 100, find the calculation in the forum
// at least 1 cluster of 800 weight units

// UnitTestConfig returns the unit test configuration for the beacon.
func UnitTestConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// NodeSimUnitTestConfig returns configuration for the beacon the unit tests with node simulation .
func NodeSimUnitTestConfig(tb testing.TB) Config { _ = "STUB: not implemented"; return *new(Config) }
