package blocks

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/hare4"
	"github.com/spacemeshos/go-spacemesh/proposals/store"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

// Generator generates a block from proposals.
type Generator struct {
	logger *zap.Logger
	cfg    Config
	once   sync.Once
	eg     errgroup.Group
	stop   func()

	db        sql.StateDatabase
	atxs      *atxsdata.Data
	proposals *store.Store
	msh       meshProvider
	executor  executor
	fetcher   system.ProposalFetcher
	cert      certifier
	patrol    layerPatrol

	hareCh           <-chan hare4.ConsensusOutput
	optimisticOutput map[types.LayerID]*proposalMetadata
}

// Config is the config for Generator.
type Config struct {
	GenBlockInterval   time.Duration
	BlockGasLimit      uint64
	OptFilterThreshold int
}

func defaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// GeneratorOpt for configuring Generator.
type GeneratorOpt func(*Generator)

// WithConfig defines cfg for Generator.
func WithConfig(cfg Config) GeneratorOpt { _ = "STUB: not implemented"; return *new(GeneratorOpt) }

// WithGeneratorLogger defines logger for Generator.
func WithGeneratorLogger(logger *zap.Logger) GeneratorOpt {
	_ = "STUB: not implemented"
	return *new(GeneratorOpt)
}

// WithHareOutputChan sets the chan to listen to hare output.
func WithHareOutputChan(ch <-chan hare4.ConsensusOutput) GeneratorOpt {
	_ = "STUB: not implemented"
	return *new(GeneratorOpt)
}

// NewGenerator creates new block generator.
func NewGenerator(
	db sql.StateDatabase,
	atxs *atxsdata.Data,
	proposals *store.Store,
	exec executor,
	m meshProvider,
	f system.ProposalFetcher,
	c certifier,
	p layerPatrol,
	opts ...GeneratorOpt,
) *Generator {
	_ = "STUB: not implemented"
	return nil
}

// Start starts listening to hare output.
func (g *Generator) Start(ctx context.Context) {
	_ = "STUB: not implemented"

	// TODO(mafa): fix this
	return
}

// nolint:fatcontext

// Stop stops listening to hare output.
func (g *Generator) Stop() { _ = "STUB: not implemented"; return }

func (g *Generator) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (g *Generator) processHareOutput(ctx context.Context, out hare4.ConsensusOutput) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetch proposals from peers if not locally available

// now all proposals should be in the local store

func (g *Generator) processOptimisticLayers(max types.LayerID) { _ = "STUB: not implemented"; return }

func (g *Generator) saveAndCertify(ctx context.Context, lid types.LayerID, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genBlockOptimistic(ctx context.Context, md *proposalMetadata) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
