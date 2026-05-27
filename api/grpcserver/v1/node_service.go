package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// NodeService is a grpc server that provides the NodeService, which exposes node-related
// data such as node status, software version, errors, etc. It can also be used to start
// the sync process, or to shut down the node.
type NodeService struct {
	mesh        meshAPI
	genTime     genesisTimeAPI
	peerCounter peerCounter
	syncer      syncer
	appVersion  string
	appCommit   string
}

// RegisterService registers this service with a grpc server instance.
func (s *NodeService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *NodeService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of this service.
func (s *NodeService) String() string { _ = "STUB: not implemented"; return "" }

// NewNodeService creates a new grpc service using config data.
func NewNodeService(
	peers peerCounter,
	msh meshAPI,
	genTime genesisTimeAPI,
	syncer syncer,
	appVersion string,
	appCommit string,
) *NodeService {
	_ = "STUB: not implemented"
	return nil
}

// Echo returns the response for an echo api request. It's used for E2E tests.
func (s *NodeService) Echo(_ context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Version returns the version of the node software as a semver string.
func (s *NodeService) Version(context.Context, *emptypb.Empty) (*pb.VersionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build returns the build of the node software.
func (s *NodeService) Build(context.Context, *emptypb.Empty) (*pb.BuildResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Status returns a status object providing information about the connected peers, sync status,
// current and verified layer.
func (s *NodeService) Status(ctx context.Context, _ *pb.StatusRequest) (*pb.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// number of connected peers
// whether the node is synced
// latest layer we saw from the network
// current layer, based on time
// latest verified layer

func (s *NodeService) NodeInfo(context.Context, *emptypb.Empty) (*pb.NodeInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *NodeService) getLayers() (curLayer, latestLayer, verifiedLayer uint32) {
	_ = "STUB: not implemented"
	// We cannot get meaningful data from the mesh during the genesis epochs since there are no blocks in these
	// epochs, so just return the current layer instead
	return 0, 0, 0
}

// STREAMS

// StatusStream exposes a stream of node status updates.
func (s *NodeService) StatusStream(_ *pb.StatusStreamRequest, stream pb.NodeService_StatusStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

// statusCh works a bit differently than the other streams. It doesn't actually
// send us data. Instead, it just notifies us that there's new data to be read.

// number of connected peers
// whether the node is synced
// latest layer we saw from the network
// current layer, based on time
// latest verified layer

// ErrorStream exposes a stream of node errors.
func (s *NodeService) ErrorStream(_ *pb.ErrorStreamRequest, stream pb.NodeService_ErrorStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert internal error level into level understood by the API.
func convertErrorLevel(level zapcore.Level) pb.LogLevel {
	_ = "STUB: not implemented"
	return *new(pb.LogLevel)
}
