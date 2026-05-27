package model

import (
	"math/rand"
)

func newHare(rng *rand.Rand) *hare { _ = "STUB: not implemented"; return nil }

// hare is an instance of the hare consensus.
// At the end of each layer it outputs input vector and coinflip events.
type hare struct {
	rng *rand.Rand
}

// OnMessage produces blocks.
func (h *hare) OnMessage(m Messenger, event Message) { _ = "STUB: not implemented"; return }

// head and tails are at equal probability.
