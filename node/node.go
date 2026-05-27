// Package node contains the main executable for go-spacemesh node
package node

import (
	"context"
	"io"
	"net/http"

	"github.com/gofrs/flock"
	pyroscope "github.com/grafana/pyroscope-go"
	grpc_logsettable "github.com/grpc-ecosystem/go-grpc-middleware/logging/settable"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/spacemeshos/go-spacemesh/activation"
	"github.com/spacemeshos/go-spacemesh/api/grpcserver"
	"github.com/spacemeshos/go-spacemesh/atxsdata"
	"github.com/spacemeshos/go-spacemesh/beacon"
	"github.com/spacemeshos/go-spacemesh/blocks"
	"github.com/spacemeshos/go-spacemesh/bootstrap"
	"github.com/spacemeshos/go-spacemesh/checkpoint"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/config"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/fetch"
	"github.com/spacemeshos/go-spacemesh/hare3"
	"github.com/spacemeshos/go-spacemesh/hare3/eligibility"
	"github.com/spacemeshos/go-spacemesh/hare4"
	"github.com/spacemeshos/go-spacemesh/log"
	"github.com/spacemeshos/go-spacemesh/malfeasance"
	"github.com/spacemeshos/go-spacemesh/malfeasance2"
	"github.com/spacemeshos/go-spacemesh/mesh"
	"github.com/spacemeshos/go-spacemesh/miner"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/signing"
	"github.com/spacemeshos/go-spacemesh/sql"
	dbmetrics "github.com/spacemeshos/go-spacemesh/sql/metrics"
	"github.com/spacemeshos/go-spacemesh/syncer"
	"github.com/spacemeshos/go-spacemesh/system"
	"github.com/spacemeshos/go-spacemesh/timesync"
	"github.com/spacemeshos/go-spacemesh/timesync/peersync"
	"github.com/spacemeshos/go-spacemesh/txs"
)

const (
	genesisFileName = "genesis.json"
	dbFile          = "state.sql"

	oldLocalDbFile = "node_state.sql"
	localDbFile    = "local.sql"
)

// Logger names.
const (
	ClockLogger            = "clock"
	P2PLogger              = "p2p"
	PostLogger             = "post"
	PostServiceLogger      = "postService"
	PostInfoServiceLogger  = "postInfoService"
	StateDbLogger          = "stateDb"
	ApiStateDBLogger       = "apiStateDB"
	BeaconLogger           = "beacon"
	CachedDBLogger         = "cachedDB"
	PoetDbLogger           = "poetDb"
	TrtlLogger             = "trtl"
	ATXHandlerLogger       = "atxHandler"
	ATXBuilderLogger       = "atxBuilder"
	MeshLogger             = "mesh"
	SyncLogger             = "sync"
	HareOracleLogger       = "hareOracle"
	HareLogger             = "hare"
	BlockCertLogger        = "blockCert"
	BlockGenLogger         = "blockGenerator"
	BlockHandlerLogger     = "blockHandler"
	TxHandlerLogger        = "txHandler"
	ProposalStoreLogger    = "proposalStore"
	ProposalBuilderLogger  = "proposalBuilder"
	ProposalListenerLogger = "proposalListener"
	NipostBuilderLogger    = "nipostBuilder"
	NipostValidatorLogger  = "nipostValidator"
	Fetcher                = "fetcher"
	TimeSyncLogger         = "timesync"
	VMLogger               = "vm"
	GRPCLogger             = "grpc"
	ConStateLogger         = "conState"
	ExecutorLogger         = "executor"
	MalfeasanceLogger      = "malfeasance"
	Malfeasance2Logger     = "malfeasance2"
	BootstrapLogger        = "bootstrap"
)

func GetCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NOTE(dshulyak) this needs to be max level so that child logger can can be current level or below.
// otherwise it will fail later when child logger will try to increase level.

// os.Interrupt for all systems, especially windows, syscall.SIGTERM is mainly for docker.

// ensure all data folders exist

// Don't print usage on error from this point forward

// This blocks until the context is finished or until an error is produced

// FIXME: per https://github.com/spacemeshos/go-spacemesh/issues/3830

// versionCmd returns the current version of spacemesh.

func configure(c *cobra.Command, configPath string, conf *config.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// apply CLI args to config

var grpcLog = grpc_logsettable.ReplaceGrpcLoggerV2()

// LoadConfig loads config and preset (if provided) into the provided config.
// It first loads the preset and then overrides it with values from the config file.
func LoadConfig(cfg *config.Config, preset string, src io.Reader) error {
	_ = "STUB: not implemented"

	// read in config from src
	return nil
}

// override default config with preset if provided

// Unmarshal config file into config struct

// Disabled because it was broken for some time with `github.com/spf13/viper` `v1.19.0` and now
// previously untagged fields are used in existing configs.
// Instead of now tagging all untagged fields we will just allow them and disable fields explicitly that
// must not be configurable.
// WithIgnoreUntagged(),

func WithZeroFields() viper.DecoderConfigOption {
	_ = "STUB: not implemented"
	return *new(viper.DecoderConfigOption)
}

func WithIgnoreUntagged() viper.DecoderConfigOption {
	_ = "STUB: not implemented"
	return *new(viper.DecoderConfigOption)
}

func WithErrorUnused() viper.DecoderConfigOption {
	_ = "STUB: not implemented"
	return *new(viper.DecoderConfigOption)
}

// Option to modify an App instance.
type Option func(app *App)

// WithLog enables logger for an App.
func WithLog(logger log.Log) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConfig overwrites default App config.
func WithConfig(conf *config.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

// New creates an instance of the spacemesh app.
func New(opts ...Option) *App { _ = "STUB: not implemented"; return nil }

// TODO(mafa): this is a hack to suppress debugging logs on 0000.defaultLogger
// to fix this we should get rid of the global logger and pass app.log to all
// components that need it

// App is the cli app singleton.
type App struct {
	*cobra.Command
	fileLock            *flock.Flock
	signers             []*signing.EdSigner
	Config              *config.Config
	db                  sql.StateDatabase
	apiDB               sql.StateDatabase
	cachedDB            *datastore.CachedDB
	dbMetrics           *dbmetrics.DBMetricsCollector
	localDB             sql.LocalDatabase
	grpcPublicServer    *grpcserver.Server
	grpcPrivateServer   *grpcserver.Server
	grpcPostServer      *grpcserver.Server
	grpcTLSServer       *grpcserver.Server
	jsonAPIServer       *grpcserver.JSONHTTPServer
	grpcServices        map[grpcserver.Service]grpcserver.ServiceAPI
	pprofService        *http.Server
	profilerService     *pyroscope.Profiler
	syncer              *syncer.Syncer
	proposalBuilder     *miner.ProposalBuilder
	mesh                *mesh.Mesh
	atxsdata            *atxsdata.Data
	clock               *timesync.NodeClock
	hare3               *hare3.Hare
	hare4               *hare4.Hare
	hareResultsChan     chan hare4.ConsensusOutput
	hOracle             *eligibility.Oracle
	blockGen            *blocks.Generator
	certifier           *blocks.Certifier
	atxBuilder          *activation.Builder
	atxHandler          *activation.Handler
	txHandler           *txs.TxHandler
	validator           *activation.Validator
	edVerifier          *signing.EdVerifier
	beaconProtocol      *beacon.ProtocolDriver
	log                 log.Log
	syncLogger          log.Log
	conState            *txs.ConservativeState
	fetcher             *fetch.Fetch
	ptimesync           *peersync.Sync
	updater             *bootstrap.Updater
	poetDb              *activation.PoetDb
	postVerifier        activation.PostVerifier
	postSupervisor      *activation.PostSupervisor
	malfeasanceHandler  *malfeasance.Handler
	malfeasance2Handler *malfeasance2.Handler
	errCh               chan error

	host *p2p.Host

	loggers map[string]*zap.AtomicLevel
	started chan struct{} // this channel is closed once the app has finished starting
	eg      *errgroup.Group
}

func (app *App) loadCheckpoint(ctx context.Context) (*checkpoint.PreservedData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *App) Started() <-chan struct{} {
	_ = "STUB: not implemented"

	// Lock locks the app for exclusive use. It returns an error if the app is already locked.
	return nil
}

func (app *App) Lock() error { _ = "STUB: not implemented"; return nil }

// Unlock unlocks the app. It is a no-op if the app is not locked.
func (app *App) Unlock() { _ = "STUB: not implemented"; return }

// Initialize parses and validates the node configuration and sets up logging.
func (app *App) Initialize() error { _ = "STUB: not implemented"; return nil }

// setupLogging configured the app logging system.
func (app *App) setupLogging() { _ = "STUB: not implemented"; return }

func (app *App) getAppInfo() string { _ = "STUB: not implemented"; return "" }

// Cleanup stops all app services.
func (app *App) Cleanup(ctx context.Context) { _ = "STUB: not implemented"; return }

// Wrap the top-level logger to add context info and set the level for a
// specific module. Calling this method and will create a new logger every time
// and not re-use an existing logger with the same name.
//
// This method is not safe to be called concurrently.
func (app *App) addLogger(name string, logger log.Log) log.Log {
	_ = "STUB: not implemented"
	return *new(log.Log)
}

// SetLogLevel updates the log level of an existing logger.
func (app *App) SetLogLevel(name, loglevel string) error { _ = "STUB: not implemented"; return nil }

func (app *App) initServices(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO(dshulyak) this needs to be improved, but dependency graph is a bit complicated

// should be removed after hare4 transition is complete

// TODO(dshulyak) makes no sense. how we ended using it?

// in a remote setup we register eagerly so the atxBuilder can warn about missing connections asap.
// Any setup with more than one signer is considered a remote setup. If there is only one signer it
// is considered a remote setup if the key for the signer has not been sourced from `supervisedIDKeyFileName`.
//
// In a supervised setup the postSetupManager will register at the atxBuilder when
// it finished initializing, to avoid warning about a missing connection when the supervised post
// service isn't ready yet.

func (app *App) launchStandalone(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (app *App) listenToUpdates(ctx context.Context) { _ = "STUB: not implemented"; return }

func (app *App) startServices(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (app *App) grpcService(svc grpcserver.Service, lg log.Log) (grpcserver.ServiceAPI, error) {
	_ = "STUB: not implemented"
	return *new(grpcserver.ServiceAPI), nil
}

// StartSmeshing is only supported in a supervised setup (single signer)

// v2beta1

func (app *App) startAPIServices(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// check services for uniques across all endpoints

// start servers if at least one endpoint is defined for them

// public server needs restriction on max connection age to prevent attacks

// 0.0.0.0 isn't a valid address to connect to on windows

// supervised setup but not started

func (app *App) stopServices(ctx context.Context) { _ = "STUB: not implemented"; return }

// err is always nil

// err is always nil

// err is always nil

// err is always nil

// SetGrpcLogger unfortunately is global
// this ensures that a test-logger isn't used after the app shuts down
// by e.g. a grpc connection to the node that is still open - like in TestSpacemeshApp_NodeService

func (app *App) setupDBs(ctx context.Context, lg log.Log) error {
	_ = "STUB: not implemented"
	return nil
}

// already checked above

// Start starts the Spacemesh node and initializes all relevant services according to command line arguments provided.
func (app *App) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// app blocks until it receives a signal to exit
// this signal may come from the node or from sig-abort (ctrl-c)

func (app *App) startSynchronous(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	// notify anyone who might be listening that the app has finished starting.
	// this can be used by, e.g., app tests.
	return nil
}

// Create a contextual logger for local usage (lower-level modules will create their own contextual loggers
// using context passed down to them)

/* Setup monitoring */

// this will set the mutex profiling to sample a third of all lock events

// record block sample for every block event that takes more than 10 milliseconds

// app.Config.ProfilerURL should be the pyroscope server address
// TODO: AuthToken? no need right now since server isn't public

// by default all profilers are enabled,

/* Initialize all protocol services */

// Prevent testnet nodes from working on the mainnet, but
// don't use the network cookie on mainnet as this technique
// may be replaced later

// need post verifying service to start first

func (app *App) preserveAfterRecovery(ctx context.Context, preserved checkpoint.PreservedData) {
	_ = "STUB: not implemented"
	return
}

func (app *App) Host() *p2p.Host { _ = "STUB: not implemented"; return nil }

func decodeLoggerLevel(cfg *config.Config, name string) (zap.AtomicLevel, error) {
	_ = "STUB: not implemented"
	return *new(zap.AtomicLevel), nil
}

type tortoiseWeakCoin struct {
	db       sql.Executor
	tortoise system.Tortoise
}

func (w tortoiseWeakCoin) Set(lid types.LayerID, value bool) error {
	_ = "STUB: not implemented"
	return nil
}

func onMainNet(conf *config.Config) bool { _ = "STUB: not implemented"; return false }

// proposalConsumerHare is used for the hare3->hare4 migration
// to satisfy the proposals handler dependency on hare.
type proposalConsumerHare struct {
	hare3          *hare3.Hare
	h3DisableLayer types.LayerID
	hare4          *hare4.Hare
}

func (p *proposalConsumerHare) IsKnown(layer types.LayerID, proposal types.ProposalID) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *proposalConsumerHare) OnProposal(proposal *types.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}
