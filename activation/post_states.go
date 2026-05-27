package activation

import (
	"sync"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type postStates struct {
	log    *zap.Logger
	mu     sync.RWMutex
	states map[types.NodeID]types.PostState
}

func NewPostStates(log *zap.Logger) *postStates { _ = "STUB: not implemented"; return nil }

func (s *postStates) Set(id types.NodeID, state types.PostState) { _ = "STUB: not implemented"; return }

func (s *postStates) Get() map[types.NodeID]types.PostState { _ = "STUB: not implemented"; return nil }
