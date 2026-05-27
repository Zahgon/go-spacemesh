// Package mesh defines the main store point for all the persisted mesh objects
// such as ATXs, ballots and blocks.
package mesh

import (
	"context"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/common/types/result"
	"github.com/spacemeshos/go-spacemesh/malfeasance/wire"
	"github.com/spacemeshos/go-spacemesh/mesh/ballotwriter"
	"github.com/spacemeshos/go-spacemesh/sql"
	"github.com/spacemeshos/go-spacemesh/system"
)

// Mesh is the logic layer above our mesh.DB database.
type Mesh struct {
	logger   *zap.Logger
	cdb      sql.StateDatabase
	atxsdata *atxsdata.Data

	executor *Executor
	conState conservativeState
	trtl     system.Tortoise

	missingBlocks chan []types.BlockID

	mu sync.Mutex
	// latestLayer is the latest layer this node had seen from blocks
	latestLayer atomic.Value
	// latestLayerInState is the latest layer whose contents have been applied to the state
	latestLayerInState atomic.Value
	// processedLayer is the latest layer whose votes have been processed
	processedLayer      atomic.Value
	nextProcessedLayers map[types.LayerID]struct{}
	maxProcessedLayer   types.LayerID

	ballotWriter *ballotwriter.BallotWriter
}

// NewMesh creates a new instant of a mesh.
func NewMesh(
	db sql.StateDatabase,
	atxsdata *atxsdata.Data,
	trtl system.Tortoise,
	exec *Executor,
	state conservativeState,
	logger *zap.Logger,
) (*Mesh, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Mesh) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (msh *Mesh) recoverFromDB(latest types.LayerID) { _ = "STUB: not implemented"; return }

// LatestLayerInState returns the latest layer we applied to state.
func (msh *Mesh) LatestLayerInState() types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

// MissingBlocks returns single consumer channel.
// Consumer by contract is responsible for downloading missing blocks.
func (msh *Mesh) MissingBlocks() <-chan []types.BlockID { _ = "STUB: not implemented"; return nil }

// LatestLayer - returns the latest layer we saw from the network.
func (msh *Mesh) LatestLayer() types.LayerID { _ = "STUB: not implemented"; return *new(types.LayerID) }

// MeshHash returns the aggregated mesh hash at the specified layer.
func (msh *Mesh) MeshHash(lid types.LayerID) (types.Hash32, error) {
	_ = "STUB: not implemented"
	return *new(types.Hash32), nil
}

// setLatestLayer sets the latest layer we saw from the network.
func (msh *Mesh) setLatestLayer(lid types.LayerID) { _ = "STUB: not implemented"; return }

// GetLayer returns GetLayer i from the database.
func (msh *Mesh) GetLayer(lid types.LayerID) (*types.Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLayerVerified returns the verified, canonical block for a layer (or none for an empty layer).
func (msh *Mesh) GetLayerVerified(lid types.LayerID) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProcessedLayer returns the last processed layer ID.
func (msh *Mesh) ProcessedLayer() types.LayerID {
	_ = "STUB: not implemented"
	return *new(types.LayerID)
}

func (msh *Mesh) setProcessedLayer(layerID types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// ensureStateConsistent finds first layer where applied doesn't match
// the block in consensus, and reverts state before that layer
// if such layer is found.
func (msh *Mesh) ensureStateConsistent(ctx context.Context, results []result.Layer) error {
	_ = "STUB: not implemented"
	return nil
}

// ProcessLayer reads latest consensus results and ensures that vm state
// is consistent with results.
// It is safe to call after optimistically executing the block.
func (msh *Mesh) ProcessLayer(ctx context.Context, lid types.LayerID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(dshulyak) https://github.com/spacemeshos/go-spacemesh/issues/4425

// apply what we were able to download, as it will allow to prune some of the data from tortoise

func filterMissing(results []result.Layer, next types.LayerID) ([]result.Layer, []types.BlockID) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (msh *Mesh) applyResults(ctx context.Context, results []result.Layer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(mafa): sometimes this fails because the block executed references a tx that is not in the DB
// maybe in that case the node should try to fetch the missing txs and retry executing the block?
//
// https://github.com/spacemeshos/go-spacemesh/issues/5616

// tortoise will evict layer when OnApplied is called.
// however we also apply layers before contextual validity was determined (e.g block.Valid field)
// in such case we would apply block because of hare, and then we may evict event when block.Valid was set
// but before it was saved to database

func (msh *Mesh) saveHareOutput(ctx context.Context, lid types.LayerID, bid types.BlockID) error {
	_ = "STUB: not implemented"
	return nil
}

// check if a certificate has been generated or sync'ed.
// - node generated the certificate when it collected enough certify messages
// - hare outputs are processed in layer order. i.e. when hare fails for a previous layer N,
//   a later layer N+x has to wait for syncer to fetch block certificate in layer N for the mesh
//   to make progress. therefore, by the time layer N+x is processed, syncer may already have
//   sync'ed the certificate for this layer

// otherwise always notify tortoise about hare output

// ProcessLayerPerHareOutput receives hare output once it finishes running for a given layer.
func (msh *Mesh) ProcessLayerPerHareOutput(
	ctx context.Context,
	layerID types.LayerID,
	blockID types.BlockID,
	executed bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (msh *Mesh) setLatestLayerInState(lyr types.LayerID) { _ = "STUB: not implemented"; return }

// SetZeroBlockLayer advances the latest layer in the network with a layer
// that has no data.
func (msh *Mesh) SetZeroBlockLayer(ctx context.Context, lid types.LayerID) {
	_ = "STUB: not implemented"
	return

	// AddTXsFromProposal adds the TXs in a Proposal into the database.
}

func (msh *Mesh) AddTXsFromProposal(
	ctx context.Context,
	layerID types.LayerID,
	proposalID types.ProposalID,
	txIDs []types.TransactionID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// AddBallot to the mesh.
func (msh *Mesh) AddBallot(
	ctx context.Context,
	ballot *types.Ballot,
) (*wire.MalfeasanceProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if this ballot was the one that turned the identity to be malicious
// we call the hooks to notify tortoise and atxsdata

// so this is a bit of double work (getting the malfeasance proof right after
// we stored it), BUT, probably negligible considering the amount of malicious
// identities we have. The other way to go about this is to allocate an individual channel
// for every write (so that every writer gets its own response with the potential proof)
// However I find that more costly than this approach.

// AddBlockWithTXs adds the block and its TXs in into the database.
func (msh *Mesh) AddBlockWithTXs(ctx context.Context, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// add block to the tortoise before storing it
// otherwise fetcher will not wait until data is stored in the tortoise

// GetRewardsByCoinbase retrieves account's rewards by the coinbase address.
func (msh *Mesh) GetRewardsByCoinbase(coinbase types.Address) ([]*types.Reward, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MissingBlocksError struct {
	Blocks []types.BlockID
}

func (e *MissingBlocksError) Error() string { _ = "STUB: not implemented"; return "" }
