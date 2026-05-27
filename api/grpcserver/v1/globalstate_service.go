package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"google.golang.org/grpc"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

// GlobalStateService exposes global state data, output from the STF.
type GlobalStateService struct {
	mesh     meshAPI
	conState conservativeState
}

// RegisterService registers this service with a grpc server instance.
func (s *GlobalStateService) RegisterService(server *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

func (s *GlobalStateService) RegisterHandlerService(mux *runtime.ServeMux) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the name of the service.
func (s *GlobalStateService) String() string { _ = "STUB: not implemented"; return "" }

// NewGlobalStateService creates a new grpc service using config data.
func NewGlobalStateService(msh meshAPI, conState conservativeState) *GlobalStateService {
	_ = "STUB: not implemented"
	return nil
}

// GlobalStateHash returns the latest layer and its computed global state hash.
func (s *GlobalStateService) GlobalStateHash(
	context.Context,
	*pb.GlobalStateHashRequest,
) (*pb.GlobalStateHashResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *GlobalStateService) getAccount(addr types.Address) (acct *pb.Account, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Account returns current and projected counter and balance for one account.
func (s *GlobalStateService) Account(ctx context.Context, in *pb.AccountRequest) (*pb.AccountResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load data

// AccountDataQuery returns historical account data such as rewards and receipts.
func (s *GlobalStateService) AccountDataQuery(
	ctx context.Context,
	in *pb.AccountDataQueryRequest,
) (*pb.AccountDataQueryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the filter flags

// MAX RESULTS, OFFSET
// There is some code duplication here as this is implemented in other Query endpoints,
// but without generics, there's no clean way to do this for different types.

// Adjust for max results, offset

// Skip to offset, don't send more than max results

// If the offset is too high there is nothing to return (this is not an error)

// If the max results is too high, trim it. If MaxResults is zero, that means unlimited
// (since we have no way to distinguish between zero and its not being provided).

// STREAMS

// AccountDataStream exposes a stream of account-related data.
func (s *GlobalStateService) AccountDataStream(
	in *pb.AccountDataStreamRequest,
	stream pb.GlobalStateService_AccountDataStreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Subscribe to the various streams

// Apply address filter

// The Reporter service just sends us the account address. We are responsible
// for looking up the other required data here. Get the account balance and
// nonce.

// Apply address filter

// Apply address filter

// Result:      receipt.Result,

// SvmData: nil,

// GlobalStateStream exposes a stream of global data data items: rewards, receipts, account info, global state hash.
func (s *GlobalStateService) GlobalStateStream(
	in *pb.GlobalStateStreamRequest,
	stream pb.GlobalStateService_GlobalStateStreamServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Subscribe to the various streams

// Whenever new state is applied to the mesh, a new layer is reported.
// There is no separate reporting specifically for new state.

// The Reporter service just sends us the account address. We are responsible
// for looking up the other required data here. Get the account balance and
// nonce.
