package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type activationService struct {
	goldenAtx   types.ATXID
	atxProvider atxProvider
}

func NewActivationService(atxProvider atxProvider, goldenAtx types.ATXID) *activationService {
	_ = "STUB: not implemented"
	return nil
}

// RegisterService implements ServiceAPI.
func (s *activationService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *activationService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the service name.
func (s *activationService) String() string { _ = "STUB: not implemented"; return "" }

// Get implements v1.ActivationServiceServer.
func (s *activationService) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *activationService) Highest(ctx context.Context, req *emptypb.Empty) (*pb.HighestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toMalfeasancePB(nodeID types.NodeID, proof []byte, includeProof bool) *pb.MalfeasanceProof {
	_ = "STUB: not implemented"
	return nil
}
