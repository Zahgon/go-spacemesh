package rewards

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/sql/builder"
)

const fullQuery = `select pubkey, coinbase, layer, total_reward, layer_reward from rewards`

type decoderCallback func(*types.Reward, error) bool

func decoder(fn decoderCallback) sql.Decoder { _ = "STUB: not implemented"; return *new(sql.Decoder) }

// Add reward to the database.
func Add(db sql.Executor, reward *types.Reward) error { _ = "STUB: not implemented"; return nil }

// Revert the rewards to the specified layer.
func Revert(db sql.Executor, revertTo types.LayerID) error { _ = "STUB: not implemented"; return nil }

// ListByCoinbase lists rewards from all layers for the coinbase address.
func ListByCoinbase(db sql.Executor, coinbase types.Address) (rst []*types.Reward, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IterateRewardsOps(
	db sql.Executor,
	operations builder.Operations,
	fn func(reward *types.Reward) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}
