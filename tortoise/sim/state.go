package sim

import (
	"testing"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
)

func newState(tb testing.TB, logger *zap.Logger, conf config, atxdata *atxsdata.Data) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// State of the node.
type State struct {
	logger *zap.Logger

	DB      *datastore.CachedDB
	Atxdata *atxsdata.Data
}

// OnBeacon callback to store generated beacon.
func (s *State) OnBeacon(eid types.EpochID, beacon types.Beacon) { _ = "STUB: not implemented"; return }

// OnActivationTx callback to store activation transaction.
func (s *State) OnActivationTx(atx *types.ActivationTx) {
	_ = "STUB: not implemented"
	// TODO: consider using actual values for malicious if needed
	return
}

// OnBallot callback to store ballot.
func (s *State) OnBallot(ballot *types.Ballot) { _ = "STUB: not implemented"; return }

// OnBlock callback to store block.
func (s *State) OnBlock(block *types.Block) { _ = "STUB: not implemented"; return }

// OnHareOutput callback to store hare output.
func (s *State) OnHareOutput(lid types.LayerID, bid types.BlockID) {
	_ = "STUB: not implemented"
	return
}

// OnCoinflip callback to store coinflip.
func (s *State) OnCoinflip(lid types.LayerID, coinflip bool) { _ = "STUB: not implemented"; return }
