package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/libp2p/go-libp2p/core/network"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/spacemeshos/go-spacemesh/events"
	"github.com/spacemeshos/go-spacemesh/sql"
)

// DebugService exposes global state data, output from the STF.
type DebugService struct {
	db       sql.StateDatabase
	conState conservativeState
	netInfo  networkInfo
	oracle   oracle
	loggers  map[string]*zap.AtomicLevel
}

// RegisterService registers this service with a grpc server instance.
func (d *DebugService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (d *DebugService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of this service.
func (d *DebugService) String() string { _ = "STUB: not implemented"; return "" }

// NewDebugService creates a new grpc service using config data.
func NewDebugService(db sql.StateDatabase, conState conservativeState, host networkInfo, oracle oracle,
	loggers map[string]*zap.AtomicLevel,
) *DebugService {
	_ = "STUB: not implemented"
	return nil
}

// Accounts returns current counter and balance for all accounts.
func (d *DebugService) Accounts(ctx context.Context, in *pb.AccountsRequest) (*pb.AccountsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Address is bech32 string, not a 0x hex string

// NetworkInfo query provides NetworkInfoResponse.
func (d *DebugService) NetworkInfo(ctx context.Context, _ *emptypb.Empty) (*pb.NetworkInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ActiveSet query provides hare active set for the specified epoch.
func (d *DebugService) ActiveSet(ctx context.Context, req *pb.ActiveSetRequest) (*pb.ActiveSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProposalsStream streams all proposals confirmed by hare.
func (d *DebugService) ProposalsStream(_ *emptypb.Empty, stream pb.DebugService_ProposalsStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

// send empty header after subscribing to the channel.
// this is optional but allows subscriber to wait until stream is fully initialized.

func (d *DebugService) ChangeLogLevel(ctx context.Context, req *pb.ChangeLogLevelRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func castEventProposal(ev *events.EventProposal) *pb.Proposal {
	_ = "STUB: not implemented"
	return nil
}

func convertNATType(natType network.NATDeviceType) pb.NetworkInfoResponse_NATType {
	_ = "STUB: not implemented"
	return *new(pb.NetworkInfoResponse_NATType)
}

func convertReachability(r network.Reachability) pb.NetworkInfoResponse_Reachability {
	_ = "STUB: not implemented"
	return *new(pb.NetworkInfoResponse_Reachability)
}
