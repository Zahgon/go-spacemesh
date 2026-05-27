package timesync

import (
	"sync"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/metrics"
)

var tickDistance = metrics.NewHistogramWithBuckets(
	"tick_distance",
	"clock",
	"distance between layer ticks",
	[]string{},
	prometheus.ExponentialBuckets(1, 2, 10),
).WithLabelValues()

// NodeClock is the struct holding a real clock.
type NodeClock struct {
	LayerConverter // layer conversions provider

	clock        clockwork.Clock // provides the time
	genesis      time.Time
	tickInterval time.Duration

	mu            sync.Mutex    // protects the following fields
	lastTicked    types.LayerID // track last ticked layer
	minLayer      types.LayerID // track earliest layer that has a channel waiting for tick
	layerChannels map[types.LayerID]chan struct{}

	stop chan struct{}
	once sync.Once

	log *zap.Logger
	eg  errgroup.Group
}

// NewClock return TimeClock struct that notifies tickInterval has passed.
func NewClock(opts ...OptionFunc) (*NodeClock, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *NodeClock) startClock() error { _ = "STUB: not implemented"; return nil }

// GenesisTime returns at which time this clock has started (used to calculate current tick).
func (t *NodeClock) GenesisTime() time.Time {
	_ = "STUB: not implemented"

	// Close closes the clock ticker.
	return *new(time.Time)
}

func (t *NodeClock) Close() { _ = "STUB: not implemented"; return }

// tick processes the current tick. It iterates over all layers that have passed since the last tick and notifies
// listeners that are awaiting these layers.
func (t *NodeClock) tick() { _ = "STUB: not implemented"; return }

// don't warn right after fresh startup

// close await channel for prev layers

// CurrentLayer gets the current layer.
func (t *NodeClock) CurrentLayer() types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

// AwaitLayer returns a channel that will be signaled when layer id layerID was ticked by the clock,
// or if this layer has passed while sleeping. It does so by closing the returned channel.
func (t *NodeClock) AwaitLayer(layerID types.LayerID) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// passed the time of layerID
