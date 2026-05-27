package main

import (
	"context"
	"net/http"
	"regexp"
	"time"

	"github.com/spf13/afero"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

const fileRegex = "/epoch-(?P<Epoch>[0-9]+)-update-(?P<Suffix>[a-z]+)"

type NetworkParam struct {
	Genesis      time.Time
	LyrsPerEpoch uint32
	LyrDuration  time.Duration
	Offset       uint32
}

func (np *NetworkParam) updateBeaconTime(targetEpoch types.EpochID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (np *NetworkParam) updateActiveSetTime(targetEpoch types.EpochID) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Server is used to serve bootstrap update data during systest. NOT intended for production use.
// In particular, it does not protect against data loss and will serve whatever is the latest
// one on disk, even tho it's for an old epoch.
type Server struct {
	*http.Server
	eg              errgroup.Group
	logger          *zap.Logger
	fs              afero.Fs
	gen             *Generator
	genFallback     bool
	bootstrapEpochs []types.EpochID
	regex           *regexp.Regexp
}

type SrvOpt func(*Server)

func WithSrvLogger(logger *zap.Logger) SrvOpt { _ = "STUB: not implemented"; return *new(SrvOpt) }

func WithSrvFilesystem(fs afero.Fs) SrvOpt { _ = "STUB: not implemented"; return *new(SrvOpt) }

func WithBootstrapEpochs(epochs []types.EpochID) SrvOpt {
	_ = "STUB: not implemented"
	return *new(SrvOpt)
}

func NewServer(gen *Generator, fallback bool, port int, opts ...SrvOpt) *Server {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) Start(ctx context.Context, errCh chan error, params *NetworkParam) {
	_ = "STUB: not implemented"
	return
}

// start generating fallback data

func (s *Server) genWithRetry(ctx context.Context, epoch types.EpochID, maxRetries int) error {
	_ = "STUB: not implemented"
	return nil
}

// In systests, we want to be sure the nodes use the fallback data unconditionally.
// Use a fixed known value for beacon to be sure that fallback is used during testing.
func epochBeacon(epoch types.EpochID) types.Beacon {
	_ = "STUB: not implemented"
	return *new(types.Beacon)
}

func (s *Server) GenBootstrap(ctx context.Context, epoch types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) GenFallbackBeacon(epoch types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) GenFallbackActiveSet(ctx context.Context, epoch types.EpochID) error {
	_ = "STUB: not implemented"
	return nil
}

// in systests, we want to be sure the nodes use the fallback data unconditionally
// we only use half of the active set as fallback value, so we can be sure that fallback is used during testing.
func getPartialActiveSet(ctx context.Context, smEndpoint string, targetEpoch types.EpochID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// enough to allow hare to pass

func (s *Server) Stop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Server) handle(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func (s *Server) handleCheckpoint(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) servefile(f string, w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func CheckpointFilename() string { _ = "STUB: not implemented"; return "" }

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
