package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

var statusMap = map[types.PostState]pb.PostState_State{
	types.PostStateIdle:    pb.PostState_IDLE,
	types.PostStateProving: pb.PostState_PROVING,
}

// PostInfoService provides information about connected PostServices.
type PostInfoService struct {
	states postState
}

// RegisterService registers this service with a grpc server instance.
func (s *PostInfoService) RegisterService(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *PostInfoService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of this service.
func (s *PostInfoService) String() string { _ = "STUB: not implemented"; return "" }

// NewPostInfoService creates a new instance of the post info grpc service.
func NewPostInfoService(states postState) *PostInfoService { _ = "STUB: not implemented"; return nil }

func (s *PostInfoService) PostStates(context.Context, *pb.PostStatesRequest) (*pb.PostStatesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
