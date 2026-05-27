package v1

import (
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/activation"
	"github.com/spacemeshos/go-spacemesh/common/types"
)

// PostService is a grpc server that PoST nodes can connect to in order to register.
// The bidirectional stream established between the node and the PoST node can be used
// to send challenges and receive proofs.
type PostService struct {
	log *zap.Logger

	clientMtx        sync.Mutex
	allowConnections bool
	client           map[types.NodeID]*postClient
	queryInterval    time.Duration
}

type postCommand struct {
	req  *pb.NodeRequest
	resp chan<- *pb.ServiceResponse
}

// RegisterService registers this service with a grpc server instance.
func (s *PostService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *PostService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of this service.
func (s *PostService) String() string { _ = "STUB: not implemented"; return "" }

type PostServiceOpt func(*PostService)

func PostServiceQueryInterval(interval time.Duration) PostServiceOpt {
	_ = "STUB: not implemented"
	return *new(PostServiceOpt)
}

// NewPostService creates a new instance of the post grpc service.
func NewPostService(log *zap.Logger, opts ...PostServiceOpt) *PostService {
	_ = "STUB: not implemented"
	return nil
}

// AllowConnections sets if the grpc service accepts new incoming connections from post services.
func (s *PostService) AllowConnections(allow bool) { _ = "STUB: not implemented"; return }

// connectionAllowed returns if the grpc service accepts new incoming connections from post services.
func (s *PostService) connectionAllowed() bool { _ = "STUB: not implemented"; return false }

// Register is called by the PoST service to connect with the node.
// It creates a bidirectional stream that is kept open until either side closes it.
// The other functions on this service are called by services of the node to send
// requests to the PoST node and receive responses.
func (s *PostService) Register(stream pb.PostService_RegisterServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *PostService) setConnection(nodeId types.NodeID, con chan postCommand) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *PostService) dropConnection(nodeId types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *PostService) Client(nodeId types.NodeID) (activation.PostClient, error) {
	_ = "STUB: not implemented"
	return *new(activation.PostClient), nil
}
