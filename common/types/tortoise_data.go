package types

import "go.uber.org/zap/zapcore"

type BallotTortoiseData struct {
	ID            BallotID       `json:"id"`
	Smesher       NodeID         `json:"node"`
	Layer         LayerID        `json:"lid"`
	Eligibilities uint32         `json:"elig"`
	AtxID         ATXID          `json:"atxid"`
	Opinion       Opinion        `json:"opinion"`
	EpochData     *ReferenceData `json:"epochdata"`
	Ref           *BallotID      `json:"ref"`
	Malicious     bool           `json:"mal"`
}

func (b *BallotTortoiseData) SetMalicious() { _ = "STUB: not implemented"; return }

func (b *BallotTortoiseData) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ReferenceData struct {
	Beacon        Beacon `json:"beacon"`
	Eligibilities uint32 `json:"elig"`
}

func (r *ReferenceData) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
