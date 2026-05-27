package beacon

import (
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type proposalList []Proposal

func (hl proposalList) sort() []Proposal { _ = "STUB: not implemented"; return nil }

func (hl proposalList) hash() types.Hash32 { _ = "STUB: not implemented"; return *new(types.Hash32) }

// an error is never returned: https://golang.org/pkg/hash/#Hash

func (hl proposalList) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
