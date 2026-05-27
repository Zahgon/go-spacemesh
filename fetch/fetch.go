// Package fetch contains mechanism to fetch Data from remote peers
package fetch

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/fetch/peers"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/p2p/server"
	"github.com/spacemeshos/go-spacemesh/proposals/store"
	"github.com/spacemeshos/go-spacemesh/sql"
)

const (
	atxProtocol       = "ax/1"
	lyrDataProtocol   = "ld/1"
	hashProtocol      = "hs/1"
	activeSetProtocol = "as/1"
	meshHashProtocol  = "mh/1"
	legacyMalProtocol = "ml/1"
	malProtocol       = "ml/2"
	OpnProtocol       = "lp/2"

	cacheSize = 1000

	RedundantPeers  = 5
	identifyTimeout = 3 * time.Second
)

var (
	// ErrExceedMaxRetries is returned when MaxRetriesForRequest attempts has been made to fetch
	// data for a hash and failed.
	ErrExceedMaxRetries = errors.New("fetch failed after max retries for request")

	errValidatorsNotSet = errors.New("validators not set")
)

// request contains all relevant Data for a single request for a specified hash.
type request struct {
	ctx       context.Context
	hash      types.Hash32   // hash is the hash of the Data requested
	hint      datastore.Hint // the hint from which database to fetch this hash
	validator dataReceiver
	promise   *promise
	retries   int
}

type promise struct {
	once      sync.Once
	completed chan struct{}
	err       error
}

var protocolMap = map[datastore.Hint]string{
	datastore.ActiveSet: activeSetProtocol,
}

type batchInfo struct {
	RequestBatch
	protocol string
	peer     p2p.Peer
}

// setID calculates the hash of all requests and sets it as this batches ID.
func (b *batchInfo) setID() { _ = "STUB: not implemented"; return }

func (b *batchInfo) toMap() map[types.Hash32]RequestMessage { _ = "STUB: not implemented"; return nil }

func (b *batchInfo) extraProtocols() []string { _ = "STUB: not implemented"; return nil }

func makeBatch(peer p2p.Peer, reqs []RequestMessage) *batchInfo {
	_ = "STUB: not implemented"
	return nil
}

type ServerConfig struct {
	Queue    int           `mapstructure:"queue"`
	Requests int           `mapstructure:"requests"`
	Interval time.Duration `mapstructure:"interval"`
}

func (s ServerConfig) ToOpts() []server.Opt { _ = "STUB: not implemented"; return nil }

// Config is the configuration file of the Fetch component.
type Config struct {
	BatchTimeout         time.Duration           `mapstructure:"batchtimeout"`
	BatchSize            int                     `mapstructure:"batchsize"`
	QueueSize            int                     `mapstructure:"queuesize"`
	MaxRetriesForRequest int                     `mapstructure:"maxretriesforrequest"`
	RequestTimeout       time.Duration           `mapstructure:"request-timeout"`
	RequestHardTimeout   time.Duration           `mapstructure:"request-hard-timeout"`
	EnableServerMetrics  bool                    `mapstructure:"servers-metrics"`
	ServersConfig        map[string]ServerConfig `mapstructure:"servers"`
	Streaming            bool                    `mapstructure:"streaming"`
	// The maximum number of concurrent requests to get ATXs.
	GetAtxsConcurrency   int64                  `mapstructure:"getatxsconcurrency"`
	DecayingTag          server.DecayingTagSpec `mapstructure:"decaying-tag"`
	LogPeerStatsInterval time.Duration          `mapstructure:"log-peer-stats-interval"`
}

func (c Config) getServerConfig(protocol string) ServerConfig {
	_ = "STUB: not implemented"
	return *new(ServerConfig)
}

// DefaultConfig is the default config for the fetch component.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// serves 1 MB of data

// serves 1 KB of data

// serves atxs, ballots, active sets
// atx - 1 KB
// ballots > 300 bytes
// often queried after receiving gossip message

// active sets (can get quite large)

// serves at most 100 hashes - 3 KB

// serves all legacy malicious ids (á 32 byte, ~2255 as of Jan 2025) - <100 KB

// serves all malicious ids (á 32 byte, 0 as of Jan 2025) - <100 KB

// 64 bytes

// randomPeer returns a random peer from current peer list.
func randomPeer(peers []p2p.Peer) p2p.Peer { _ = "STUB: not implemented"; return *new(p2p.Peer) }

// Option is a type to configure a fetcher.
type Option func(*Fetch)

// WithContext configures the shutdown context for the fetcher.
func WithContext(ctx context.Context) Option {
	_ = "STUB: not implemented"
	return *

	// TODO(mafa): fix this
	new(Option)
}

// nolint:fatcontext

// WithConfig configures the config for the fetcher.
func WithConfig(c Config) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogger configures logger for the fetcher.
func WithLogger(log *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func withServers(s map[string]requester) Option { _ = "STUB: not implemented"; return *new(Option) }

func withHost(h host) Option { _ = "STUB: not implemented"; return *new(Option) }

// Fetch is the main struct that contains network peers and logic to batch and dispatch hash fetch requests.
type Fetch struct {
	cfg    Config
	logger *zap.Logger
	bs     *datastore.BlobStore
	host   host
	peers  *peers.Peers

	servers    map[string]requester
	validators *dataValidators

	// unprocessed contains requests that are not processed
	unprocessed map[types.Hash32]*request
	// ongoing contains requests that have been processed and are waiting for responses
	ongoing      map[types.Hash32]*request
	batchTimeout *time.Ticker
	mu           sync.Mutex
	onlyOnce     sync.Once
	hashToPeers  *HashPeersCache

	shutdownCtx context.Context
	cancel      context.CancelFunc
	eg          errgroup.Group

	getAtxsLimiter limiter
}

// NewFetch creates a new Fetch struct.
func NewFetch(
	db sql.StateDatabase,
	proposals *store.Store,
	host *p2p.Host,
	peerCache *peers.Peers,
	opts ...Option,
) (*Fetch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this should never happen, since cacheSize is a constant, but just in case

// NOTE(dshulyak) this is to avoid tests refactoring.
// there is one test that covers this part.

// Make sure that the protocol list for the peer is correct.
// This is similar to what Host.NewStream does to make
// sure it is possible to use one of the specified
// protocols. If we don't do this, there may be a race causing
// some peers to be unnecessarily ignored.

func (f *Fetch) registerServer(
	host *p2p.Host,
	protocol string,
	handler server.StreamHandler,
) {
	_ = "STUB: not implemented"
	return
}

type dataValidators struct {
	atx               SyncValidator
	poet              SyncValidator
	ballot            SyncValidator
	activeset         SyncValidator
	block             SyncValidator
	proposal          SyncValidator
	txBlock           SyncValidator
	txProposal        SyncValidator
	legacyMalfeasance SyncValidator
	malfeasance       SyncValidator
}

// SetMalfeasanceProvider sets the malfeasance provider dependency.
//
// TODO(mafa): this is a hack because of a cyclic dependency between the packages
//
//	malfeasance2 -> fetcher -> datastore -> malfeasance2
func (f *Fetch) SetMalfeasanceProvider(p datastore.MalfeasanceProvider) {
	_ = "STUB: not implemented"
	return
}

// SetValidators sets the handlers to validate various mesh data fetched from peers.
func (f *Fetch) SetValidators(
	atx SyncValidator,
	poet SyncValidator,
	ballot SyncValidator,
	activeset SyncValidator,
	block SyncValidator,
	prop SyncValidator,
	txBlock SyncValidator,
	txProposal SyncValidator,
	mal SyncValidator,
	mal2 SyncValidator,
) {
	_ = "STUB: not implemented"
	return
}

// Start starts handling fetch requests.
func (f *Fetch) Start() error { _ = "STUB: not implemented"; return nil }

// Stop stops handling fetch requests.
func (f *Fetch) Stop() { _ = "STUB: not implemented"; return }

// stopped returns if we should stop.
func (f *Fetch) stopped() bool { _ = "STUB: not implemented"; return false }

// here we receive all requests for hashes for all DBs and batch them together before we send the request to peer
// there can be a priority request that will not be batched.
func (f *Fetch) loop() { _ = "STUB: not implemented"; return }

// Process the batch.

func (f *Fetch) meteredRequest(
	ctx context.Context,
	protocol string,
	peer p2p.Peer,
	req []byte,
	extraProtocols ...string,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Fetch) meteredStreamRequest(
	ctx context.Context,
	protocol string,
	peer p2p.Peer,
	req []byte,
	callback func(context.Context, io.ReadWriter) (int, error),
	extraProtocols ...string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// receive Data from message server and call response handlers accordingly.
func (f *Fetch) receiveResponse(data []byte, batch *batchInfo) { _ = "STUB: not implemented"; return }

// iterate all hash Responses

// validation fetch data recursively. offload to another goroutine

// iterate all requests that didn't return value from peer and notify
// they will be retried for MaxRetriesForRequest

func (f *Fetch) hashValidationDone(hash types.Hash32, err error) { _ = "STUB: not implemented"; return }

func (f *Fetch) failAfterRetry(hash types.Hash32) { _ = "STUB: not implemented"; return }

// first check if we have it locally from gossips

// put the request back to the unprocessed list

// this is the main function that sends the hash request to the peer.
func (f *Fetch) requestHashBatchFromPeers() { _ = "STUB: not implemented"; return }

func (f *Fetch) getUnprocessed() []RequestMessage { _ = "STUB: not implemented"; return nil }

// only send one request per hash

// move the processed requests to pending

func (f *Fetch) send(requests []RequestMessage) { _ = "STUB: not implemented"; return }

func (f *Fetch) organizeRequests(requests []RequestMessage) map[p2p.Peer][]*batchInfo {
	_ = "STUB: not implemented"
	return nil
}

// When selecting peers, provide protocol IDs so that peers that aren't yet fully
// initialized are not picked for the request, avoiding unnecessary errors.

// split every peer's requests into batches of f.cfg.BatchSize each

// Use batches of size 1 for hashes with specific protocol.
// This is currently used for active sets which are too large
// to be batched.

// streamBatch dispatches batched request messages to provided peer and
// receives the response in streaming mode.
func (f *Fetch) streamBatch(peer p2p.Peer, batch *batchInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Request is synchronous, it will return errors only if size of the bytes buffer
// is large or target peer is not connected

// iterate all requests that didn't return value from peer and notify
// they will be retried for MaxRetriesForRequest

func (f *Fetch) receiveStreamedBatch(
	ctx context.Context,
	s io.ReadWriter,
	batch *batchInfo,
	batchMap map[types.Hash32]RequestMessage,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// we make sure to read the blob before continuing

// validation fetches data recursively. offload to another goroutine

// sendBatch dispatches batched request messages to provided peer.
func (f *Fetch) sendBatch(peer p2p.Peer, batch *batchInfo) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Request is synchronous,
// it will return errors only if size of the bytes buffer is large
// or target peer is not connected

// handleHashError is called when an error occurred processing batches of the following hashes.
func (f *Fetch) handleHashError(batch *batchInfo, err error) { _ = "STUB: not implemented"; return }

// getHash is the regular buffered call to get a specific hash, using provided hash, h as hint the receiving end will
// know where to look for the hash, this function returns HashDataPromiseResult channel that will hold Data received
// or error.
func (f *Fetch) getHash(
	ctx context.Context,
	hash types.Hash32,
	h datastore.Hint,
	receiver dataReceiver,
) (*promise, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if we already have this hash locally

// Process the batch.

// RegisterPeerHashes registers provided peer for a list of hashes.
func (f *Fetch) RegisterPeerHashes(peer p2p.Peer, hashes []types.Hash32) {
	_ = "STUB: not implemented"
	return
}

func (f *Fetch) SelectBestShuffled(n int) []p2p.Peer {
	_ = "STUB: not implemented"
	// shuffle to split the load between peers with good latency.
	// and it avoids sticky behavior, when temporarily faulty peer had good latency in the past.
	return nil
}
