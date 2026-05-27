package timesync

import (
	"time"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// LayerConverter is a converter between time to layer ID struct.
type LayerConverter struct {
	duration time.Duration // the layer duration, assumed to be > 0
	genesis  time.Time     // the genesis time
}

// TimeToLayer returns the layer of the provided time.
func (lc LayerConverter) TimeToLayer(t time.Time) types.LayerID {
	_ = "STUB: not implemented"
	return *
	// the genesis is in the future
	new(types.LayerID)
}

// LayerToTime returns the time of the provided layer.
func (lc LayerConverter) LayerToTime(id types.LayerID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
