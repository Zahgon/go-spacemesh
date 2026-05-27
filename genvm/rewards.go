package vm

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
)

func (v *VM) addRewards(
	layer types.LayerID,
	ss *core.StagedCache,
	fees uint64,
	blockRewards []types.CoinbaseReward,
) ([]types.Reward, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
