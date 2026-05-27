package fixture

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/common/types/result"
)

func RLayers(layers ...result.Layer) []result.Layer { _ = "STUB: not implemented"; return nil }

func RLayerNonFinal(lid types.LayerID, blocks ...result.Block) result.Layer {
	_ = "STUB: not implemented"
	return *new(result.Layer)
}

func ROpinion(lid types.LayerID, opinion types.Hash32, blocks ...result.Block) result.Layer {
	_ = "STUB: not implemented"
	return *new(result.Layer)
}

func RLayer(lid types.LayerID, blocks ...result.Block) result.Layer {
	_ = "STUB: not implemented"
	return *new(result.Layer)
}

type RBlockOpt func(*result.Block)

func Hare() RBlockOpt { _ = "STUB: not implemented"; return *new(RBlockOpt) }

func Valid() RBlockOpt { _ = "STUB: not implemented"; return *new(RBlockOpt) }

func Invalid() RBlockOpt { _ = "STUB: not implemented"; return *new(RBlockOpt) }

func Data() RBlockOpt { _ = "STUB: not implemented"; return *new(RBlockOpt) }

func Good() RBlockOpt { _ = "STUB: not implemented"; return *new(RBlockOpt) }

func RBlock(id types.BlockID, opts ...RBlockOpt) result.Block {
	_ = "STUB: not implemented"
	return *new(result.Block)
}

func IDGen(sid string) (id types.BlockID) { _ = "STUB: not implemented"; return *new(types.BlockID) }
