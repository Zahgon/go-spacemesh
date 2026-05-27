package beacon

import (
	"go.uber.org/zap/zapcore"
)

type proposalSet map[Proposal]struct{}

func (p proposalSet) sorted() proposalList { _ = "STUB: not implemented"; return *new(proposalList) }

func (p proposalSet) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
