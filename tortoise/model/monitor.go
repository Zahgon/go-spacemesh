package model

import (
	"testing"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// Monitor is an interface for monitors.
type Monitor interface {
	OnEvent(Event)
	Test()
}

// Event is received by monitors, and used for validation.
type Event any

// EventVerified is raised when tortoise outputs verified layer.
type EventVerified struct {
	ID       string
	Layer    types.LayerID
	Verified types.LayerID
	Revert   bool
}

func newVerifiedMonitor(tb testing.TB, genesis types.LayerID) *verifiedMonitor {
	_ = "STUB: not implemented"
	return nil
}

type verifiedMonitor struct {
	tb       testing.TB
	last     types.LayerID
	genesis  types.LayerID
	verified map[string]types.LayerID
}

func (m *verifiedMonitor) OnEvent(event Event) { _ = "STUB: not implemented"; return }

func (m *verifiedMonitor) Test() { _ = "STUB: not implemented"; return }
