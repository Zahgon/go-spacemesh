package tortoise

import (
	"encoding/json"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/common/types/result"
)

type output struct {
	Type  eventType       `json:"t"`
	Event json.RawMessage `json:"o"`
}

type tracer struct {
	logger *zap.Logger
}

func (t *tracer) On(event traceEvent) { _ = "STUB: not implemented"; return }

type TraceOpt func(*zap.Config)

func WithOutput(path string) TraceOpt { _ = "STUB: not implemented"; return *new(TraceOpt) }

func newTracer(opts ...TraceOpt) *tracer { _ = "STUB: not implemented"; return nil }

type traceRunner struct {
	opts          []Opt
	atxdata       *atxsdata.Data
	trt           *Tortoise
	pending       map[types.BallotID]*DecodedBallot
	assertOutputs bool
	assertErrors  bool
}

func RunTrace(path string, breakpoint func(), opts ...Opt) error {
	_ = "STUB: not implemented"
	return nil
}

type eventType = uint16

const (
	traceStart eventType = 1 + iota
	traceWeakCoin
	traceBeacon
	traceAtx
	traceBallot
	traceDecode
	traceStore
	traceEncode
	traceTally
	traceBlock
	traceHare
	traceUpdates
	traceApplied
	traceMalfeasance
	traceRecoveredBlocks
)

type traceEvent interface {
	Type() eventType
	New() traceEvent
	Run(*traceRunner) error
}

type ConfigTrace struct {
	Hdist                    uint32 `json:"hdist"`
	Zdist                    uint32 `json:"zdist"`
	WindowSize               uint32 `json:"window"`
	MaxExceptions            uint32 `json:"exceptions"`
	BadBeaconVoteDelayLayers uint32 `json:"delay"`
	LayerSize                uint32 `json:"layer-size"`
	EpochSize                uint32 `json:"epoch-size"` // this field is not set in the original config
	EffectiveGenesis         uint32 `json:"effective-genesis"`
}

func (c *ConfigTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (c *ConfigTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (c *ConfigTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type AtxTrace struct {
	ID          types.ATXID   `json:"id"`
	TargetEpoch types.EpochID `json:"target"`
	Atx         *atxsdata.ATX `json:",inline"`
}

func (a *AtxTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (a *AtxTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (a *AtxTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type WeakCoinTrace struct {
	Layer types.LayerID `json:"lid"`
	Coin  bool          `json:"coin"`
}

func (w *WeakCoinTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (w *WeakCoinTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (w *WeakCoinTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type BeaconTrace struct {
	Epoch  types.EpochID `json:"epoch"`
	Beacon types.Beacon  `json:"beacon"`
}

func (b *BeaconTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (b *BeaconTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (b *BeaconTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type BallotTrace struct {
	Ballot *types.BallotTortoiseData `json:",inline"`
}

func (b *BallotTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (b *BallotTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (b *BallotTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type DecodeBallotTrace struct {
	Ballot *types.BallotTortoiseData `json:",inline"`
	Error  string                    `json:"e"`

	// TODO(dshulyak) want to assert decoding results somehow
}

func (d *DecodeBallotTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (d *DecodeBallotTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (b *DecodeBallotTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type StoreBallotTrace struct {
	ID        types.BallotID `json:"id"`
	Malicious bool           `json:"mal"`
	Error     string         `json:"e,omitempty"`
}

func (s *StoreBallotTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (s *StoreBallotTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (s *StoreBallotTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type EncodeVotesTrace struct {
	Layer   types.LayerID  `json:"lid"`
	Opinion *types.Opinion `json:"opinion"`
	Error   string         `json:"e"`
}

func (e *EncodeVotesTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (e *EncodeVotesTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (e *EncodeVotesTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type TallyTrace struct {
	Layer types.LayerID `json:"lid"`
}

func (t *TallyTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (t *TallyTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (t *TallyTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type HareTrace struct {
	Layer types.LayerID `json:"lid"`
	Vote  types.BlockID `json:"vote"`
}

func (h *HareTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (h *HareTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (h *HareTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type UpdatesTrace struct {
	From    types.LayerID  `json:"from"`
	To      types.LayerID  `json:"to"`
	Error   string         `json:"e"`
	Results []result.Layer `json:"results"`
}

func (u *UpdatesTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (u *UpdatesTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (u *UpdatesTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type AppliedTrace struct {
	Layer   types.LayerID `json:"layer"`
	Opinion types.Hash32  `json:"opinion"`
	Result  bool          `json:"rst"`
}

func (a *AppliedTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (a *AppliedTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (a *AppliedTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type BlockTrace struct {
	Header types.BlockHeader `json:",inline"`
}

func (b *BlockTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (b *BlockTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (b *BlockTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type MalfeasanceTrace struct {
	ID types.NodeID `json:"id"`
}

func (m *MalfeasanceTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (m *MalfeasanceTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (m *MalfeasanceTrace) Run(r *traceRunner) error { _ = "STUB: not implemented"; return nil }

type headerWithValidity struct {
	Header types.BlockHeader `json:"header"`
	Valid  bool              `json:"valid"`
}

func newRecoveredBlocksTrace(
	layer types.LayerID,
	blocks map[types.BlockHeader]bool,
	hare *types.BlockID,
) *RecoveredBlocksTrace {
	_ = "STUB: not implemented"
	return nil
}

type RecoveredBlocksTrace struct {
	Layer  types.LayerID        `json:"layer"`
	Blocks []headerWithValidity `json:"blocks"`
	Hare   *types.BlockID       `json:"hare"`
}

func (r *RecoveredBlocksTrace) Type() eventType { _ = "STUB: not implemented"; return *new(eventType) }

func (r *RecoveredBlocksTrace) New() traceEvent { _ = "STUB: not implemented"; return *new(traceEvent) }

func (r *RecoveredBlocksTrace) Run(tr *traceRunner) error { _ = "STUB: not implemented"; return nil }

func assertErrors(err error, expect string) error { _ = "STUB: not implemented"; return nil }

func newEventEnum() eventEnum { _ = "STUB: not implemented"; return *new(eventEnum) }

type eventEnum struct {
	types map[eventType]traceEvent
}

func (e *eventEnum) Register(ev traceEvent) { _ = "STUB: not implemented"; return }

func (e *eventEnum) Decode(dec *json.Decoder) (traceEvent, error) {
	_ = "STUB: not implemented"
	return *new(traceEvent), nil
}
