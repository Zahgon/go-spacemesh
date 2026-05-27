package v2alpha1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	spacemeshv2alpha1 "github.com/spacemeshos/api/release/go/spacemesh/v2alpha1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/timesync"
)

// nodePeerCounter is an api to get current peer count.
type nodePeerCounter interface {
	PeerCount() uint64
}

// nodeMeshAPI is an api for getting mesh status.
type nodeMeshAPI interface {
	LatestLayer() types.LayerID
	LatestLayerInState() types.LayerID
	ProcessedLayer() types.LayerID
}

func NewNodeService(peers nodePeerCounter, msh nodeMeshAPI, clock *timesync.NodeClock, syncer syncer) *NodeService {
	_ = "STUB: not implemented"
	return nil
}

type NodeService struct {
	mesh        nodeMeshAPI
	clock       *timesync.NodeClock
	peerCounter nodePeerCounter
	syncer      syncer
}

func (s *NodeService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *NodeService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the service name.
func (s *NodeService) String() string { _ = "STUB: not implemented"; return "" }

func (s *NodeService) Status(ctx context.Context, _ *spacemeshv2alpha1.NodeStatusRequest) (
	*spacemeshv2alpha1.NodeStatusResponse, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// latest layer node has seen from blocks
// last layer node has applied to the state
// last layer whose votes have been processed
// current layer, based on clock time
