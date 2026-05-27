package model

import (
	"math/rand"
)

func newBeacon(rng *rand.Rand) *beacon { _ = "STUB: not implemented"; return nil }

// beacon outputs beacon at the last layer in epoch.
type beacon struct {
	rng *rand.Rand
}

// OnMessage ...
func (b *beacon) OnMessage(m Messenger, event Message) { _ = "STUB: not implemented"; return }

// first layer of the epoch
