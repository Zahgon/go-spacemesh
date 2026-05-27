package sim

import (
	"math/rand"
	"testing"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/signing"
)

// GenOpt for configuring Generator.
type GenOpt func(*Generator)

// WithSeed configures seed for Generator. By default 0 is used.
func WithSeed(seed int64) GenOpt { _ = "STUB: not implemented"; return *new(GenOpt) }

// WithLayerSize configures average layer size.
func WithLayerSize(size uint32) GenOpt { _ = "STUB: not implemented"; return *new(GenOpt) }

// WithLogger configures logger.
func WithLogger(logger *zap.Logger) GenOpt { _ = "STUB: not implemented"; return *new(GenOpt) }

// WithPath configures path for persistent databases.
func WithPath(path string) GenOpt { _ = "STUB: not implemented"; return *new(GenOpt) }

// WithStates creates n states.
func WithStates(n int) GenOpt { _ = "STUB: not implemented"; return *new(GenOpt) }

func withRng(rng *rand.Rand) GenOpt { _ = "STUB: not implemented"; return *new(GenOpt) }

func withConf(conf config) GenOpt { _ = "STUB: not implemented"; return *new(GenOpt) }

type config struct {
	Path           string
	LayerSize      uint32
	LayersPerEpoch uint32
	StateInstances int
	WindowSize     uint32
}

func defaults() config { _ = "STUB: not implemented"; return *new(config) }

// New creates Generator instance.
func New(tb testing.TB, opts ...GenOpt) *Generator { _ = "STUB: not implemented"; return nil }

// TODO support multiple persist states.

// Generator for layers of blocks.
type Generator struct {
	tb testing.TB

	logger *zap.Logger
	rng    *rand.Rand
	conf   config

	states []State

	nextLayer types.LayerID
	// key is when to return => value is the layer to return
	reordered map[types.LayerID]types.LayerID
	layers    []*types.Layer
	units     [2]int

	activations []*types.ActivationTx
	ticksRange  [2]int
	ticks       []uint64
	prevHeight  []uint64

	keys []*signing.EdSigner
}

// SetupOpt configures setup.
type SetupOpt func(g *setupConf)

// WithSetupMinerRange number of miners will be selected between low and high values.
func WithSetupMinerRange(low, high int) SetupOpt { _ = "STUB: not implemented"; return *new(SetupOpt) }

// WithSetupUnitsRange adjusts units of the ATXs, which will directly affect block weight.
func WithSetupUnitsRange(low, high int) SetupOpt { _ = "STUB: not implemented"; return *new(SetupOpt) }

// WithSetupTicksRange configures range of atxs, that will be randomly chosen by atxs.
func WithSetupTicksRange(low, high int) SetupOpt { _ = "STUB: not implemented"; return *new(SetupOpt) }

// WithSetupTicks configures ticks for every atx.
func WithSetupTicks(ticks ...uint64) SetupOpt { _ = "STUB: not implemented"; return *new(SetupOpt) }

type setupConf struct {
	Miners     [2]int
	UnitsRange [2]int
	TicksRange [2]int
	Ticks      []uint64
}

func defaultSetupConf() setupConf { _ = "STUB: not implemented"; return *new(setupConf) }

// GetState at index.
func (g *Generator) GetState(i int) State { _ = "STUB: not implemented"; return *new(State) }

func (g *Generator) addState(state State) { _ = "STUB: not implemented"; return }

func (g *Generator) popState(i int) State { _ = "STUB: not implemented"; return *new(State) }

// Setup should be called before running Next.
func (g *Generator) Setup(opts ...SetupOpt) { _ = "STUB: not implemented"; return }

func (g *Generator) generateAtxs() { _ = "STUB: not implemented"; return }

// Layer returns generated layer.
func (g *Generator) Layer(i int) *types.Layer { _ = "STUB: not implemented"; return nil }
