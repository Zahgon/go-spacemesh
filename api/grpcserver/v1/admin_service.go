package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/peerinfo"
	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	chunksize      = 1024
	defaultNumAtxs = 4
)

// AdminService exposes endpoints for node administration.
type AdminService struct {
	db      sql.StateDatabase
	dataDir string
	recover func()
	p       peers
}

// NewAdminService creates a new admin grpc service.
func NewAdminService(db sql.StateDatabase, dataDir string, p peers) *AdminService {
	_ = "STUB: not implemented"
	return nil
}

// Allow time for the response to be sent.

// RegisterService registers this service with a grpc server instance.
func (a *AdminService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (a *AdminService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of this service.
func (a *AdminService) String() string { _ = "STUB: not implemented"; return "" }

func (a *AdminService) CheckpointStream(
	req *pb.CheckpointStreamRequest,
	stream pb.AdminService_CheckpointStreamServer,
) error {
	_ = "STUB: not implemented"
	// checkpoint data can be more than 4MB, it can cause stress
	// - on the client side (default limit on the receiving end)
	// - locally as the node already loads db query result in memory
	return nil
}

func (a *AdminService) Recover(ctx context.Context, _ *pb.RecoverRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AdminService) EventsStream(_ *pb.EventStreamRequest, stream pb.AdminService_EventsStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

// send empty header after subscribing to the channel.
// this is optional but allows subscriber to wait until stream is fully initialized.

func (a *AdminService) PeerInfoStream(_ *emptypb.Empty, stream pb.AdminService_PeerInfoStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

// There is no guarantee that the peers originally returned will still
// be connected by the time we call ConnectedPeerInfo.

func connKind(kind peerinfo.Kind) pb.ConnectionInfo_Kind {
	_ = "STUB: not implemented"
	return *new(pb.ConnectionInfo_Kind)
}

func peerStats(stats p2p.PeerRequestStats) *pb.PeerRequestStats {
	_ = "STUB: not implemented"
	return nil
}
