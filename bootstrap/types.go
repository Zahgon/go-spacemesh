package bootstrap

import (
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type Update struct {
	Version string    `json:"version"`
	Data    InnerData `json:"data"`
}

type InnerData struct {
	Epoch EpochData `json:"epoch"`
}

type EpochData struct {
	ID        uint32   `json:"number"`
	Beacon    string   `json:"beacon"`
	ActiveSet []string `json:"activeSet"`
}

func (ed *EpochData) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type VerifiedUpdate struct {
	Data      *EpochOverride
	Persisted string
}

type EpochOverride struct {
	Epoch     types.EpochID
	Beacon    types.Beacon
	ActiveSet []types.ATXID
}

func (vd *VerifiedUpdate) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
