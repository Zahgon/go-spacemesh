package v1

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
)

// MeshService exposes mesh data such as accounts, blocks, and transactions.
type MeshService struct {
	cdb            *datastore.CachedDB
	mesh           meshAPI // Mesh
	conState       conservativeState
	genTime        genesisTimeAPI
	layersPerEpoch uint32
	genesisID      types.Hash20
	layerDuration  time.Duration
	layerAvgSize   uint32
	txsPerProposal uint32
}

// RegisterService registers this service with a grpc server instance.
func (s *MeshService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *MeshService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of this service.
func (s *MeshService) String() string { _ = "STUB: not implemented"; return "" }

// NewMeshService creates a new service using config data.
func NewMeshService(
	cdb *datastore.CachedDB,
	msh meshAPI,
	cstate conservativeState,
	genTime genesisTimeAPI,
	layersPerEpoch uint32,
	genesisID types.Hash20,
	layerDuration time.Duration,
	layerAvgSize,
	txsPerProposal uint32,
) *MeshService {
	_ = "STUB: not implemented"
	return nil
}

// GenesisTime returns the network genesis time as UNIX time.
func (s *MeshService) GenesisTime(context.Context, *pb.GenesisTimeRequest) (*pb.GenesisTimeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CurrentLayer returns the current layer number.
func (s *MeshService) CurrentLayer(context.Context, *pb.CurrentLayerRequest) (*pb.CurrentLayerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CurrentEpoch returns the current epoch number.
func (s *MeshService) CurrentEpoch(context.Context, *pb.CurrentEpochRequest) (*pb.CurrentEpochResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenesisID returns the network ID.
func (s *MeshService) GenesisID(context.Context, *pb.GenesisIDRequest) (*pb.GenesisIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EpochNumLayers returns the number of layers per epoch (a network parameter).
func (s *MeshService) EpochNumLayers(context.Context, *pb.EpochNumLayersRequest) (*pb.EpochNumLayersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LayerDuration returns the layer duration in seconds (a network parameter).
func (s *MeshService) LayerDuration(context.Context, *pb.LayerDurationRequest) (*pb.LayerDurationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MaxTransactionsPerSecond returns the max number of tx per sec (a network parameter).
func (s *MeshService) MaxTransactionsPerSecond(
	context.Context,
	*pb.MaxTransactionsPerSecondRequest,
) (*pb.MaxTransactionsPerSecondResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QUERIES

func (s *MeshService) getFilteredTransactions(
	from types.LayerID,
	address types.Address,
) ([]*types.MeshTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountMeshDataQuery returns account data.
func (s *MeshService) AccountMeshDataQuery(
	ctx context.Context,
	in *pb.AccountMeshDataQueryRequest,
) (*pb.AccountMeshDataQueryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the filter flags

// Gather transaction data

// MAX RESULTS, OFFSET
// There is some code duplication here as this is implemented in other Query endpoints,
// but without generics, there's no clean way to do this for different types.

// Skip to offset, don't send more than max results

// If the offset is too high there is nothing to return (this is not an error)

// If the max results is too high, trim it. If MaxResults is zero, that means unlimited
// (since we have no way to distinguish between zero and its not being provided).

func convertLayerID(l types.LayerID) *pb.LayerNumber { _ = "STUB: not implemented"; return nil }

func castTransaction(t *types.Transaction) *pb.Transaction { _ = "STUB: not implemented"; return nil }

func convertActivation(a *types.ActivationTx, previous []types.ATXID) *pb.Activation {
	_ = "STUB: not implemented"
	return nil
}

// nolint:staticcheck // SA1019 (deprecated)

func (s *MeshService) readLayer(
	ctx context.Context,
	layerID types.LayerID,
	layerStatus pb.Layer_LayerStatus,
) (*pb.Layer, error) {
	_ = "STUB: not implemented"
	// Populate with what we already know
	return nil, nil
}

// read the canonical block for this layer

// Be careful with how we handle missing layers here.
// A layer that's newer than the currentLayer (defined above)
// is clearly an input error. A missing layer that's older than
// lastValidLayer is clearly an internal error. A missing layer
// between these two is a gray area: do we define this as an
// internal or an input error? For now, all missing layers produce
// internal errors.

// This is expected. We can only retrieve state root for a layer that was applied to state,
// which only happens after it's approved/confirmed.

// This is expected. We can only retrieve state root for a layer that was applied to state,
// which only happens after it's approved/confirmed.

// LayersQuery returns all mesh data, layer by layer.
func (s *MeshService) LayersQuery(ctx context.Context, in *pb.LayersQueryRequest) (*pb.LayersQueryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the latest layers that passed both consensus engines.

// First check if the layer passed the Hare, then check if it passed the Tortoise.
// It may be either, or both, but Tortoise always takes precedence.

// Be careful with how we handle missing layers here.
// A layer that's newer than the currentLayer (defined above)
// is clearly an input error. A missing layer that's older than
// lastValidLayer is clearly an internal error. A missing layer
// between these two is a gray area: do we define this as an
// internal or an input error? For now, all missing layers produce
// internal errors.

// STREAMS

// AccountMeshDataStream exposes a stream of transactions and activations for an account.
func (s *MeshService) AccountMeshDataStream(
	in *pb.AccountMeshDataStreamRequest,
	stream pb.MeshService_AccountMeshDataStreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the filter flags

// Subscribe to the stream of transactions and activations

// Apply address filter

// Apply address filter

// LayerStream exposes a stream of all mesh data per layer.
func (s *MeshService) LayerStream(_ *pb.LayerStreamRequest, stream pb.MeshService_LayerStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func convertLayerStatus(in int) pb.Layer_LayerStatus {
	_ = "STUB: not implemented"
	return *new(pb.Layer_LayerStatus)
}

func (s *MeshService) EpochStream(req *pb.EpochStreamRequest, stream pb.MeshService_EpochStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MeshService) MalfeasanceQuery(
	ctx context.Context,
	req *pb.MalfeasanceRequest,
) (*pb.MalfeasanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MeshService) MalfeasanceStream(
	req *pb.MalfeasanceStreamRequest,
	stream pb.MeshService_MalfeasanceStreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// first serve those already existed locally.
