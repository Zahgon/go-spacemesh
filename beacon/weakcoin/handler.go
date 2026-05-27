package weakcoin

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/p2p"
)

// HandleProposal defines method to handle Beacon Weak Coin Messages from gossip.
func (wc *WeakCoin) HandleProposal(ctx context.Context, peer p2p.Peer, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (wc *WeakCoin) receiveMessage(ctx context.Context, message Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (wc *WeakCoin) isNextRound(epoch types.EpochID, round types.RoundID) bool {
	_ = "STUB: not implemented"
	return false
}

// after completed epoch but haven't started the new one

// after started epoch but didn't start the round
