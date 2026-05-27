package result

import (
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type Layer struct {
	Layer    types.LayerID `json:"lid"`
	Verified bool          `json:"v"`
	Opinion  types.Hash32  `json:"opinion"`
	Blocks   []Block       `json:"blocks"`
}

// FirstValid returns first block that crossed positive tortoise threshold,
// or if layer didn't accumulate enough weight yet - use hare result.
func (l *Layer) FirstValid() types.BlockID { _ = "STUB: not implemented"; return *new(types.BlockID) }

func (l Layer) String() string { _ = "STUB: not implemented"; return "" }

func (l *Layer) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type Block struct {
	Header  types.Vote `json:"header"`
	Valid   bool       `json:"v"`
	Local   bool       `json:"l"` // set to true if block crossed local threshold
	Invalid bool       `json:"i"`
	Hare    bool       `json:"h"`
	Data    bool       `json:"d"`
}

func (b *Block) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
