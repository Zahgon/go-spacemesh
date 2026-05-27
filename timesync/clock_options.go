package timesync

import (
	"time"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"
)

type option struct {
	clock         clockwork.Clock
	genesisTime   time.Time
	layerDuration time.Duration
	tickInterval  time.Duration

	log *zap.Logger
}

func (o *option) validate() error { _ = "STUB: not implemented"; return nil }

type OptionFunc func(*option) error

// withClock specifies which clock the NodeClock should use. Defaults to the real clock.
func withClock(clock clockwork.Clock) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithGenesisTime sets the genesis time for the NodeClock.
func WithGenesisTime(genesis time.Time) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithLayerDuration sets the layer duration for the NodeClock.
func WithLayerDuration(d time.Duration) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

func WithTickInterval(d time.Duration) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithLogger sets the logger for the NodeClock.
func WithLogger(logger *zap.Logger) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }
