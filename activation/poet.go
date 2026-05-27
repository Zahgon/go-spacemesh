package activation

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql/localsql/certifier"
)

//go:generate mockgen -typed -package=activation -destination=poet_mocks.go -source=./poet.go

var (
	ErrInvalidRequest           = errors.New("invalid request")
	ErrUnauthorized             = errors.New("unauthorized")
	errCertificatesNotSupported = errors.New("poet doesn't support certificates")
	errIncompatiblePhaseShift   = errors.New("fetched poet phase_shift is incompatible with configured phase_shift")
	errCertifierNotConfigured   = errors.New("certifier service not configured")
)

type PoetPowParams struct {
	Challenge  []byte
	Difficulty uint
}

type PoetPoW struct {
	Nonce  uint64
	Params PoetPowParams
}

type PoetAuth struct {
	*PoetPoW
	*certifier.PoetCert
}

type PoetClient interface {
	Id() []byte
	Address() string

	PowParams(ctx context.Context) (*PoetPowParams, error)
	Submit(
		ctx context.Context,
		deadline time.Time,
		prefix, challenge []byte,
		signature types.EdSignature,
		nodeID types.NodeID,
		auth PoetAuth,
	) (*types.PoetRound, error)
	Proof(ctx context.Context, roundID string) (*types.PoetProofMessage, []types.Hash32, error)
	Info(ctx context.Context) (*types.PoetInfo, error)
}

// HTTPPoetClient implements PoetProvingServiceClient interface.
type HTTPPoetClient struct {
	id                    []byte
	baseURL               *url.URL
	client                *retryablehttp.Client
	submitChallengeClient *retryablehttp.Client
	logger                *zap.Logger
}

func checkRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// A wrapper around zap.Logger to make it compatible with
// retryablehttp.LeveledLogger interface.
type retryableHttpLogger struct {
	inner *zap.Logger
}

func (r retryableHttpLogger) Error(format string, args ...any) { _ = "STUB: not implemented"; return }

func (r retryableHttpLogger) Info(format string, args ...any) { _ = "STUB: not implemented"; return }

func (r retryableHttpLogger) Warn(format string, args ...any) { _ = "STUB: not implemented"; return }

// Debug seems to be the only logging level used by retryablehttp. Since it logs when it retries a request,
// and we want users to see what is happening we change the level to Info.
func (r retryableHttpLogger) Debug(format string, args ...any) { _ = "STUB: not implemented"; return }

type PoetClientOpts func(*HTTPPoetClient)

func withCustomHttpClient(client *http.Client) PoetClientOpts {
	_ = "STUB: not implemented"
	return *new(PoetClientOpts)
}

func WithLogger(logger *zap.Logger) PoetClientOpts {
	_ = "STUB: not implemented"
	return *new(PoetClientOpts)
}

func customLinearJitterBackoff(min, max time.Duration, _ int, _ *http.Response) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// NewHTTPPoetClient returns new instance of HTTPPoetClient connecting to the specified url.
func NewHTTPPoetClient(server types.PoetServer, cfg PoetConfig, opts ...PoetClientOpts) (*HTTPPoetClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *HTTPPoetClient) Id() []byte { _ = "STUB: not implemented"; return nil }

func (c *HTTPPoetClient) Address() string { _ = "STUB: not implemented"; return "" }

func (c *HTTPPoetClient) PowParams(ctx context.Context) (*PoetPowParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Submit registers a challenge in the proving service current open round.
func (c *HTTPPoetClient) Submit(
	ctx context.Context,
	deadline time.Time,
	prefix, challenge []byte,
	signature types.EdSignature,
	nodeID types.NodeID,
	auth PoetAuth,
) (*types.PoetRound, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *HTTPPoetClient) Info(ctx context.Context) (*types.PoetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Proof implements PoetProvingServiceClient.
func (c *HTTPPoetClient) Proof(ctx context.Context, roundID string) (*types.PoetProofMessage, []types.Hash32, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *HTTPPoetClient) req(
	ctx context.Context,
	method, path string,
	reqBody, resBody proto.Message,
	client *retryablehttp.Client,
) error {
	_ = "STUB: not implemented"
	return nil
}

type cachedData[T any] struct {
	mu   sync.Mutex
	data T
	exp  time.Time
	ttl  time.Duration
}

func (c *cachedData[T]) get(init func() (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// poetService is a higher-level interface to communicate with a PoET service.
// It wraps the HTTP client, adding additional functionality.
type poetService struct {
	db     poetDbAPI
	logger *zap.Logger
	client PoetClient

	requestTimeout time.Duration

	// Used to avoid concurrent requests for proof.
	gettingProof sync.Mutex
	// cached members of the last queried proof
	proofMembers map[string][]types.Hash32

	certifier certifierService

	expectedPhaseShift time.Duration
	infoCache          cachedData[*types.PoetInfo]
	powParamsCache     cachedData[*PoetPowParams]

	// Used to calculate ticks from PoetProof.LeafCount.
	tickSize uint64
}

type PoetServiceOpt func(*poetService)

func WithCertifier(certifier certifierService) PoetServiceOpt {
	_ = "STUB: not implemented"
	return *new(PoetServiceOpt)
}

func NewPoetService(
	db poetDbAPI,
	server types.PoetServer,
	cfg PoetConfig,
	logger *zap.Logger,
	tickSize uint64,
	opts ...PoetServiceOpt,
) (*poetService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPoetServiceWithClient(
	db poetDbAPI,
	client PoetClient,
	cfg PoetConfig,
	logger *zap.Logger,
	tickSize uint64,
	opts ...PoetServiceOpt,
) *poetService {
	_ = "STUB: not implemented"
	return nil
}

func (c *poetService) TickSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *poetService) verifyPhaseShiftConfiguration(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *poetService) Address() string { _ = "STUB: not implemented"; return "" }

func (c *poetService) authorize(
	ctx context.Context,
	nodeID types.NodeID,
	challenge []byte,
	logger *zap.Logger,
) (*PoetAuth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fallback to PoW
// TODO(poszu): remove this fallback once we migrate to certificates fully.

func (c *poetService) reauthorize(
	ctx context.Context,
	id types.NodeID,
	challenge []byte,
) (*PoetAuth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *poetService) Submit(
	ctx context.Context,
	deadline time.Time,
	prefix, challenge []byte,
	signature types.EdSignature,
	nodeID types.NodeID,
) (*types.PoetRound, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to obtain a certificate

func (c *poetService) Proof(ctx context.Context, roundID string) (*types.PoetProof, []types.Hash32, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *poetService) Certify(ctx context.Context, id types.NodeID) (*certifier.PoetCert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *poetService) getInfo(ctx context.Context) (*types.PoetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *poetService) powParams(ctx context.Context) (*PoetPowParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
