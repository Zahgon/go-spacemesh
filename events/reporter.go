package events

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/event"
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// Subscription is a subscription to events.
// Consumer must be aware that publish will block if subscription is not read fast enough.
type Subscription = event.Subscription

var (
	mu sync.RWMutex
	// reporter is the event reporter singleton.
	reporter *EventReporter
)

// InitializeReporter initializes the event reporting interface with
// a nonzero channel buffer. This is useful for testing, where we want reporting to
// block.
func InitializeReporter() { _ = "STUB: not implemented"; return }

// EventHook returns hook for logger.
func EventHook() func(entry zapcore.Entry) error { _ = "STUB: not implemented"; return nil }

// If we report anything less than this we'll end up in an infinite loop

// TODO(nkryuchkov): consider returning an error and log outside the function

// ReportNewTx dispatches incoming events to the reporter singleton.
func ReportNewTx(layerID types.LayerID, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// ReportNewActivation reports a new activation.
func ReportNewActivation(activation *types.ActivationTx) error {
	_ = "STUB: not implemented"
	return nil
}

// ReportRewardReceived reports a new reward.
func ReportRewardReceived(r types.Reward) error { _ = "STUB: not implemented"; return nil }

// ReportLayerUpdate reports a new layer, or an update to an existing layer.
func ReportLayerUpdate(layer LayerUpdate) error { _ = "STUB: not implemented"; return nil }

// ReportError reports an error.
func ReportError(err NodeError) error { _ = "STUB: not implemented"; return nil }

// ReportNodeStatusUpdate reports an update to the node status. It just
// pings the listener to notify them that there is an update; the listener
// is responsible for fetching the new status details. This is because
// status contains disparate information coming from different services,
// and the listener already knows how to gather that information so there
// is no point in duplicating that logic here.
// Note: There is some overlap with channelLayer here, as a new latest
// or verified layer should be sent over that channel as well. However,
// that happens inside the Mesh, at the source. It doesn't currently
// happen here because the status update includes only a layer ID, not
// full layer data, and the Reporter currently has no way to retrieve
// full layer data.
func ReportNodeStatusUpdate() error { _ = "STUB: not implemented"; return nil }

// ReportResult reports creation or receipt of a new tx receipt.
func ReportResult(rst types.TransactionWithResult) error { _ = "STUB: not implemented"; return nil }

// ReportAccountUpdate reports an account whose data has been updated.
func ReportAccountUpdate(a types.Address) error { _ = "STUB: not implemented"; return nil }

// SubscribeTxs subscribes to new transactions.
func SubscribeTxs() (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

// SubscribeActivations subscribes to activations.
func SubscribeActivations() (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

// SubscribeLayers subscribes to all layer data.
func SubscribeLayers() (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

// SubscribeErrors subscribes to node errors.
func SubscribeErrors() (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

// SubscribeStatus subscribes to node status messages.
func SubscribeStatus() (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

// SubscribeAccount subscribes to account data updates.
func SubscribeAccount() (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

// SubscribeRewards subscribes to rewards.
func SubscribeRewards() (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

// SubscribeToLayers is used to track and report automatically every time a
// new layer is reached.
func SubscribeToLayers(ticker LayerClock) { _ = "STUB: not implemented"; return }

// This will block, so run in a goroutine

// TODO(nkryuchkov): consider returning an error and log outside the function

func SubscribeUserEvents(opts ...SubOpt) (*BufferedSubscription[UserEvent], *Ring[UserEvent], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// The status of a layer
// TODO: this list is woefully inadequate and does not map to reality.
// See https://github.com/spacemeshos/api/issues/144.
const (
	LayerStatusTypeUnknown   = iota
	LayerStatusTypeApproved  // approved by Hare
	LayerStatusTypeConfirmed // confirmed by Tortoise
	LayerStatusTypeApplied   // applied to state
)

// LayerUpdate packages up a layer with its status (which a layer does not ordinarily contain).
type LayerUpdate struct {
	LayerID types.LayerID
	Status  int
}

func (lu LayerUpdate) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// NodeError represents an internal error to be reported.
type NodeError struct {
	Msg   string
	Trace string
	Level zapcore.Level
}

// TxReceipt represents a transaction receipt.
type TxReceipt struct {
	ID      types.TransactionID
	Result  int
	GasUsed uint64
	Fee     uint64
	Layer   types.LayerID
	Index   uint32
	Address types.Address
}

// Transaction wraps a tx with its layer ID and validity info.
type Transaction struct {
	Transaction *types.Transaction
	LayerID     types.LayerID
	Valid       bool
}

// ActivationTx wraps *types.ActivationTx.
type ActivationTx struct {
	*types.ActivationTx
}

// Status indicates status change event.
type Status struct{}

// Account wraps account address.
type Account struct {
	types.Address
}

// EventReporter is the struct that receives incoming events and dispatches them.
type EventReporter struct {
	bus                event.Bus
	transactionEmitter event.Emitter
	activationEmitter  event.Emitter
	layerEmitter       event.Emitter
	errorEmitter       event.Emitter
	statusEmitter      event.Emitter
	accountEmitter     event.Emitter
	rewardEmitter      event.Emitter
	resultsEmitter     event.Emitter
	proposalsEmitter   event.Emitter
	malfeasanceEmitter event.Emitter
	events             struct {
		sync.Mutex
		buf     *Ring[UserEvent]
		emitter event.Emitter
	}
	stopChan chan struct{}
}

func (r *EventReporter) emitUserEvent(ev UserEvent) error { _ = "STUB: not implemented"; return nil }

func (r *EventReporter) subUserEvents(opts ...SubOpt) (*BufferedSubscription[UserEvent], *Ring[UserEvent], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func newEventReporter() *EventReporter { _ = "STUB: not implemented"; return nil }

// CloseEventReporter shuts down the event reporting service and closes open channels.
func CloseEventReporter() { _ = "STUB: not implemented"; return }

func newRing[T any](size int) *Ring[T] { _ = "STUB: not implemented"; return nil }

// Ring is an insert only buffer.
type Ring[T any] struct {
	data        []T
	first, last int
}

func (r *Ring[T]) insert(value T) { _ = "STUB: not implemented"; return }

func (r *Ring[T]) Copy() *Ring[T] { _ = "STUB: not implemented"; return nil }

func (r *Ring[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (r *Ring[T]) Iterate(iter func(val T) bool) { _ = "STUB: not implemented"; return }
