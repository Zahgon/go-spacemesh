package v1

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/spacemeshos/go-spacemesh/activation"
	"github.com/spacemeshos/go-spacemesh/signing"
)

// SmesherService exposes endpoints to manage smeshing.
type SmesherService struct {
	smeshingProvider activation.SmeshingProvider
	postSupervisor   postSupervisor
	grpcPostService  grpcPostService

	streamInterval time.Duration
	cmdCfg         *activation.PostSupervisorConfig
	postOpts       activation.PostSetupOpts
	sig            *signing.EdSigner
}

// RegisterService registers this service with a grpc server instance.
func (s *SmesherService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *SmesherService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of this service.
func (s *SmesherService) String() string { _ = "STUB: not implemented"; return "" }

// NewSmesherService creates a new grpc service using config data.
func NewSmesherService(
	smeshing activation.SmeshingProvider,
	postSupervisor postSupervisor,
	grpcPostService grpcPostService,
	streamInterval time.Duration,
	postOpts activation.PostSetupOpts,
	sig *signing.EdSigner,
) *SmesherService {
	_ = "STUB: not implemented"
	return nil
}

// SetPostServiceConfig sets the post supervisor config.
func (s *SmesherService) SetPostServiceConfig(cfg activation.PostSupervisorConfig) {
	_ = "STUB: not implemented"

	// IsSmeshing reports whether the node is smeshing.
	return
}

func (s *SmesherService) IsSmeshing(context.Context, *emptypb.Empty) (*pb.IsSmeshingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StartSmeshing requests that the node begin smeshing.
func (s *SmesherService) StartSmeshing(
	ctx context.Context,
	in *pb.StartSmeshingRequest,
) (*pb.StartSmeshingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SmesherService) postSetupOpts(in *pb.PostSetupOpts) (activation.PostSetupOpts, error) {
	_ = "STUB: not implemented"
	return *new(activation.PostSetupOpts), nil
}

// Overlay default with api provided opts

// StopSmeshing requests that the node stop smeshing.
func (s *SmesherService) StopSmeshing(
	ctx context.Context,
	in *pb.StopSmeshingRequest,
) (*pb.StopSmeshingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SmesherService) SmesherIDs(context.Context, *emptypb.Empty) (*pb.SmesherIDsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Coinbase returns the current coinbase setting of this node.
func (s *SmesherService) Coinbase(context.Context, *emptypb.Empty) (*pb.CoinbaseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCoinbase sets the current coinbase setting of this node.
func (s *SmesherService) SetCoinbase(_ context.Context, in *pb.SetCoinbaseRequest) (*pb.SetCoinbaseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PostSetupStatus returns post data status.
func (s *SmesherService) PostSetupStatus(ctx context.Context, _ *emptypb.Empty) (*pb.PostSetupStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PostSetupStatusStream exposes a stream of status updates during post setup.
func (s *SmesherService) PostSetupStatusStream(
	_ *emptypb.Empty,
	stream pb.SmesherService_PostSetupStatusStreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PostSetupProviders returns a list of available Post setup compute providers.
func (s *SmesherService) PostSetupProviders(
	ctx context.Context,
	in *pb.PostSetupProvidersRequest,
) (*pb.PostSetupProvidersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PostConfig returns the Post protocol config.
func (s *SmesherService) PostConfig(context.Context, *emptypb.Empty) (*pb.PostConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func statusToPbStatus(status *activation.PostSetupStatus) *pb.PostSetupStatus {
	_ = "STUB: not implemented"
	return nil
}

// assuming enum values match.
