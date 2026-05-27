package model

import (
	"math/rand"
	"testing"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/tortoise"
)

const (
	layerSize = 50
	units     = 10
)

func newCore(tb testing.TB, rng *rand.Rand, id string, logger *zap.Logger) *core {
	_ = "STUB: not implemented"
	return nil
}

// core state machine.
type core struct {
	id     string
	logger *zap.Logger
	rng    *rand.Rand

	cdb      *datastore.CachedDB
	atxdata  *atxsdata.Data
	tortoise *tortoise.Tortoise

	// generated on setup
	units  uint32
	signer *signing.EdSigner

	// set in the first layer of each epoch
	refBallot     *types.BallotID
	weight        uint64
	eligibilities uint32

	// set at the end of the epoch (MessageLayerEnd for the last layer of the epoch)
	atx types.ATXID
}

// OnMessage receive blocks, atx, input vector, beacon, coinflip and store them.
// Generate atx at the end of each epoch.
// Generate block at the start of every layer.
func (c *core) OnMessage(m Messenger, event Message) { _ = "STUB: not implemented"; return }

// TODO(dshulyak) produce ballot according to eligibilities
