package config

import (
	"github.com/spacemeshos/go-spacemesh/timesync/peersync"
)

// TimeConfig specifies the timesync params for ntp.
type TimeConfig struct {
	Peersync peersync.Config `mapstructure:"peersync"`
}

// DefaultConfig defines the default timesync configuration.
func DefaultConfig() TimeConfig {
	_ = "STUB: not implemented"
	// TimeConfigValues defines default values for all time and ntp related params.
	return *new(TimeConfig)
}
