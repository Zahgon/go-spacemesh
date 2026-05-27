package grpcserver

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// ServiceAPI allows individual grpc services to register the grpc server.
type ServiceAPI interface {
	RegisterService(*grpc.Server)
	RegisterHandlerService(*runtime.ServeMux) error
	String() string
}

// Server is a very basic grpc server.
type Server struct {
	listener string
	logger   *zap.Logger
	// BoundAddress contains the address that the server bound to, useful if
	// the server uses a dynamic port. It is set during startup and can be
	// safely accessed after Start has completed (I.E. the returned channel has
	// been waited on)
	BoundAddress string
	GrpcServer   *grpc.Server
	grp          errgroup.Group
}

func unaryGrpcLogStart(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func streamingGrpcLogStart(
	srv any,
	stream grpc.ServerStream,
	_ *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	_ = "STUB: not implemented"
	return nil
}

// NewWithServices creates a new Server listening on the provided address with the given logger and config.
// Services passed in the svc slice are registered with the server.
func NewWithServices(
	listener string,
	logger *zap.Logger,
	config Config,
	svc []ServiceAPI,
	grpcOpts ...grpc.ServerOption,
) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if listener IP is in private network range

// NewTLS creates a new Server listening on the TLSListener address with the given logger and config.
// Services passed in the svc slice are registered with the server.
func NewTLS(logger *zap.Logger, config Config, svc []ServiceAPI) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// New creates and returns a new Server listening on the given address.
// The server is configured with the given logger and config. Additional grpc options can be passed.
func New(listener string, logger *zap.Logger, config Config, grpcOpts ...grpc.ServerOption) *Server {
	_ = "STUB: not implemented"
	return nil
}

// keep alive more often than once per `MinTime` will be disconnected

// Start starts the server.
func (s *Server) Start() error { _ = "STUB: not implemented"; return nil }

// Close stops the server.
func (s *Server) Close() error { _ = "STUB: not implemented"; return nil }

// GracefulStop waits for all connections to be closed before closing the
// server and returning. If there are long running stream connections then
// GracefulStop will never return. So we call it in a background thread,
// wait a bit and then call Stop which will forcefully close any remaining
// connections.
