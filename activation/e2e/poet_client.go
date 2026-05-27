package activation

import (
	"context"
	"sync"
	"time"

	"github.com/spacemeshos/go-spacemesh/activation"
	"github.com/spacemeshos/go-spacemesh/common/types"
)

type TestPoet struct {
	mu      sync.Mutex
	round   int
	poetCfg activation.PoetConfig

	expectedMembers int
	registrations   chan []byte
}

func NewTestPoetClient(expectedMembers int, poetCfg activation.PoetConfig) *TestPoet {
	_ = "STUB: not implemented"
	return nil
}

func (p *TestPoet) Id() []byte { _ = "STUB: not implemented"; return nil }

func (p *TestPoet) Address() string { _ = "STUB: not implemented"; return "" }

func (p *TestPoet) PowParams(ctx context.Context) (*activation.PoetPowParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SAFE to be called concurrently.
func (p *TestPoet) Submit(
	_ context.Context,
	_ time.Time,
	_, challenge []byte,
	_ types.EdSignature,
	_ types.NodeID,
	_ activation.PoetAuth,
) (*types.PoetRound, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *TestPoet) Info(_ context.Context) (*types.PoetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build a proof.
//
// Waits for the expected number of registrations to be submitted
// before starting to build the proof.
//
// NOT safe to be called concurrently.
func (p *TestPoet) Proof(ctx context.Context, roundID string) (*types.PoetProofMessage, []types.Hash32, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
