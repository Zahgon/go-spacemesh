package hare3

import "github.com/spacemeshos/go-spacemesh/common/types"

type Tracer interface {
	OnStart(types.LayerID)
	OnStop(types.LayerID)
	OnActive([]*types.HareEligibility)
	OnMessageSent(*Message)
	OnMessageReceived(*Message)
}

var _ Tracer = noopTracer{}

type noopTracer struct{}

func (noopTracer) OnStart(types.LayerID) { _ = "STUB: not implemented"; return }

func (noopTracer) OnStop(types.LayerID) { _ = "STUB: not implemented"; return }

func (noopTracer) OnActive([]*types.HareEligibility) { _ = "STUB: not implemented"; return }

func (noopTracer) OnMessageSent(*Message) { _ = "STUB: not implemented"; return }

func (noopTracer) OnMessageReceived(*Message) { _ = "STUB: not implemented"; return }
