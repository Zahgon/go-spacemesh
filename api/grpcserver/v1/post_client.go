package v1

import (
	"context"
	"time"

	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// postClient represents a connection to a PoST service.
//
// It uses the grpc interface of the node to send commands to the post service.
// Additionally if instructed it will start the post service and connect it to
// the node.
type postClient struct {
	con           chan<- postCommand
	queryInterval time.Duration

	closed chan struct{}
}

func newPostClient(con chan<- postCommand, queryInterval time.Duration) *postClient {
	_ = "STUB: not implemented"
	return nil
}

func (pc *postClient) Info(ctx context.Context) (*types.PostInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *postClient) Proof(ctx context.Context, challenge []byte) (*types.Post, *types.PostInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (pc *postClient) send(ctx context.Context, req *pb.NodeRequest) (*pb.ServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// send command

// receive response

func (pc *postClient) Close() error { _ = "STUB: not implemented"; return nil }
