package store

import (
	"errors"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

var (
	ErrNotFound       = errors.New("proposal not found")
	ErrLayerEvicted   = errors.New("layer evicted")
	ErrProposalExists = errors.New("proposal already exists")
)

type Store struct {
	// number of layers to keep
	capacity types.LayerID
	logger   *zap.Logger

	// guards access to data and evicted
	mu      sync.RWMutex
	evicted *types.LayerID
	data    map[types.LayerID]layerData
}

type layerData struct {
	proposals map[types.ProposalID]*types.Proposal
	metric    prometheus.Counter
}

type StoreOption func(*Store)

func WithCapacity(capacity uint32) StoreOption { _ = "STUB: not implemented"; return *new(StoreOption) }

func WithEvictedLayer(layer types.LayerID) StoreOption {
	_ = "STUB: not implemented"
	return *new(StoreOption)
}

func WithLogger(logger *zap.Logger) StoreOption {
	_ = "STUB: not implemented"
	return *new(StoreOption)
}

func New(opts ...StoreOption) *Store { _ = "STUB: not implemented"; return nil }

func (s *Store) IsEvicted(layer types.LayerID) bool { _ = "STUB: not implemented"; return false }

func (s *Store) OnLayer(layer types.LayerID) { _ = "STUB: not implemented"; return }

func (s *Store) Add(p *types.Proposal) error { _ = "STUB: not implemented"; return nil }

func (s *Store) Get(layer types.LayerID, id types.ProposalID) *types.Proposal {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) GetForLayer(layer types.LayerID) []*types.Proposal {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) GetMany(layer types.LayerID, pids ...types.ProposalID) []*types.Proposal {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) Has(id types.ProposalID) bool { _ = "STUB: not implemented"; return false }

func (s *Store) getByID(id types.ProposalID) (*types.Proposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetBlobSize(id types.ProposalID) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Store) GetBlob(id types.ProposalID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
