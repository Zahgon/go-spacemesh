package events

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
)

// EventMalfeasance includes the malfeasance proof.
type EventMalfeasance struct {
	Smesher types.NodeID
}

// SubscribeMalfeasance subscribes malfeasance events.
func SubscribeMalfeasance() Subscription { _ = "STUB: not implemented"; return *new(Subscription) }

// ReportMalfeasance reports a malfeasance proof.
func ReportMalfeasance(nodeID types.NodeID) { _ = "STUB: not implemented"; return }
