package activation

import (
	"context"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
	certifierdb "github.com/spacemeshos/go-spacemesh/sql/localsql/certifier"
	"github.com/spacemeshos/go-spacemesh/sql/localsql/nipost"
)

type CertifierClientConfig struct {
	// Base delay between retries, scaled with the number of retries.
	RetryDelay time.Duration `mapstructure:"retry-delay"`
	// Maximum time to wait between retries
	MaxRetryDelay time.Duration `mapstructure:"max-retry-delay"`
	// Maximum number of retries
	MaxRetries int `mapstructure:"max-retries"`
}

type CertifierConfig struct {
	Client CertifierClientConfig `mapstructure:"client"`
}

func DefaultCertifierClientConfig() CertifierClientConfig {
	_ = "STUB: not implemented"
	return *new(CertifierClientConfig)
}

func DefaultCertifierConfig() CertifierConfig {
	_ = "STUB: not implemented"
	return *new(CertifierConfig)
}

type ProofToCertify struct {
	Nonce   uint32 `json:"nonce"`
	Indices []byte `json:"indices"`
	Pow     uint64 `json:"pow"`
}

type ProofToCertifyMetadata struct {
	NodeId          []byte `json:"node_id"`
	CommitmentAtxId []byte `json:"commitment_atx_id"`

	Challenge []byte `json:"challenge"`
	NumUnits  uint32 `json:"num_units"`
}

type CertifyRequest struct {
	Proof    ProofToCertify         `json:"proof"`
	Metadata ProofToCertifyMetadata `json:"metadata"`
}

type CertifyResponse struct {
	Certificate []byte `json:"certificate"`
	Signature   []byte `json:"signature"`
	PubKey      []byte `json:"pub_key"`
}

type Certifier struct {
	logger *zap.Logger
	db     sql.LocalDatabase
	client certifierClient

	certifications singleflight.Group
}

func NewCertifier(
	db sql.LocalDatabase,
	logger *zap.Logger,
	client certifierClient,
) *Certifier {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) Certificate(
	ctx context.Context,
	id types.NodeID,
	certifier *url.URL,
	pubkey []byte,
) (*certifierdb.PoetCert, error) {
	_ = "STUB: not implemented"
	// We index certs in DB by node ID and pubkey. To avoid redundant queries, we allow only 1
	// request per (nodeID, pubkey) pair to be in flight at a time.
	return nil, nil
}

func (c *Certifier) DeleteCertificate(id types.NodeID, pubkey []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type CertifierClient struct {
	client  *retryablehttp.Client
	logger  *zap.Logger
	db      sql.Executor
	localDb sql.LocalDatabase
}

type certifierClientOpts func(*CertifierClient)

func WithCertifierClientConfig(cfg CertifierClientConfig) certifierClientOpts {
	_ = "STUB: not implemented"
	return *new(certifierClientOpts)
}

func NewCertifierClient(
	db sql.Executor,
	localDb sql.LocalDatabase,
	logger *zap.Logger,
	opts ...certifierClientOpts,
) *CertifierClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *CertifierClient) obtainPostFromLastAtx(ctx context.Context, nodeId types.NodeID) (*nipost.Post, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VRF nonce is not needed

func (c *CertifierClient) obtainPost(ctx context.Context, id types.NodeID) (*nipost.Post, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no post found

func (c *CertifierClient) Certify(
	ctx context.Context,
	id types.NodeID,
	url *url.URL,
	pubkey []byte,
) (*certifierdb.PoetCert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// load NIPoST for the given ATX from the database.
func loadPost(ctx context.Context, db sql.Executor, id types.ATXID) (*types.Post, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
