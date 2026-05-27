// Package layerpatrol keeps parties informed about the general progress of each layer.
package layerpatrol

import (
	"sync"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

const bufferSize = uint32(100)

// LayerPatrol keeps progress of each layer.
type LayerPatrol struct {
	mu          sync.Mutex
	oldestLayer types.LayerID
	runByHare   map[types.LayerID]struct{}
}

// New returns an instance of LayerPatrol.
func New() *LayerPatrol { _ = "STUB: not implemented"; return nil }

// SetHareInCharge sets the layer's validation to be triggered by hare (as opposed to syncer).
func (lp *LayerPatrol) SetHareInCharge(layerID types.LayerID) { _ = "STUB: not implemented"; return }

// IsHareInCharge returns true if the hare is set to handle the validation of the specified layer.
func (lp *LayerPatrol) IsHareInCharge(layerID types.LayerID) bool {
	_ = "STUB: not implemented"
	return false
}

// CompleteHare is called by hare instance that completed this layer.
func (lp *LayerPatrol) CompleteHare(layerID types.LayerID) { _ = "STUB: not implemented"; return }
