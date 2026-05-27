package grpcserver

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// JSONHTTPServer is a JSON http server providing the Spacemesh API.
// It is implemented using a grpc-gateway. See https://github.com/grpc-ecosystem/grpc-gateway .
type JSONHTTPServer struct {
	listener       string
	collectMetrics bool
	logger         *zap.Logger

	// BoundAddress contains the address that the server bound to, useful if
	// the server uses a dynamic port. It is set during startup and can be
	// safely accessed after Start has completed (I.E. the returned channel has
	// been waited on)
	BoundAddress string
	server       *http.Server
	eg           errgroup.Group

	// basic CORS support
	origins []string
}

// NewJSONHTTPServer creates a new json http server.
func NewJSONHTTPServer(
	lg *zap.Logger,
	listener string,
	corsAllowedOrigins []string,
	collectMetrics bool,
) *JSONHTTPServer {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown stops the server.
func (s *JSONHTTPServer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// StartService starts the json api server and listens for status (started, stopped).
func (s *JSONHTTPServer) StartService(
	services ...ServiceAPI,
) error {
	_ = "STUB: not implemented"
	// At least one service must be enabled
	return nil
}

// setup metrics middleware

// register each individual, enabled service

// enable cors
