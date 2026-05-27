package cluster

import (
	"context"
	"errors"
	"time"

	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/config"
	"github.com/spacemeshos/go-spacemesh/systest/testcontext"
)

var errNotInitialized = errors.New("cluster: not initialized")

const (
	initBalance      = 100000000000000000
	defaultExtraData = "systest"
	certifierApp     = "certifier"
	poetApp          = "poet"
	bootnodeApp      = "boot"
	smesherApp       = "smesher"
	postServiceApp   = "postservice"
	bootstrapperApp  = "bootstrapper"
	bootstrapperPort = 80
	certifierPort    = 80
	poetPort         = 80

	poetFlags    = "poetflags"
	smesherFlags = "smesherflags"
	bsFlags      = "bsflags"
)

// MakePoetEndpoint generate a poet endpoint for the ith instance.
func MakePoetEndpoint(ith int) string { _ = "STUB: not implemented"; return "" }

func MakePoetMetricsEndpoint(testNamespace string, ith int) string {
	_ = "STUB: not implemented"
	return ""
}

func MakePoetGlobalEndpoint(testNamespace string, ith int) string {
	_ = "STUB: not implemented"
	return ""
}

// Deterministically generate poet keys for given instance.
func MakePoetKey(ith int) (ed25519.PublicKey, ed25519.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(ed25519.PublicKey), *new(ed25519.PrivateKey)
}

func BootstrapperEndpoint(ith int) string { _ = "STUB: not implemented"; return "" }

func BootstrapperGlobalEndpoint(namespace string, ith int) string {
	_ = "STUB: not implemented"
	return ""
}

// Opt is for configuring cluster.
type Opt func(c *Cluster)

// WithSmesherFlag adds smesher flag.
func WithSmesherFlag(flag DeploymentFlag) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithKeys generates n pre-funded keys.
func WithKeys(n int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithBootstrapperFlag(flag DeploymentFlag) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithBootstrapEpochs(epochs []int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

type GenAccount struct {
	Address types.Address
	Balance uint64
}

func WithGenesisBalances(gaccs ...GenAccount) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// Reuse will try to recover cluster from the given namespace, if not found
// it will create a new one.
func Reuse(cctx *testcontext.Context, opts ...Opt) (*Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReuseWait(cctx *testcontext.Context, opts ...Opt) (*Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default deploys bootnodes, one poet and the smeshers according to the cluster size.
func Default(cctx *testcontext.Context, opts ...Opt) (*Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// New initializes Cluster with options.
func New(cctx *testcontext.Context, opts ...Opt) *Cluster { _ = "STUB: not implemented"; return nil }

// Cluster for managing state of the spacemesh cluster.
type Cluster struct {
	persisted         bool
	smesherFlags      map[string]DeploymentFlag
	poetFlags         map[string]DeploymentFlag
	bootstrapperFlags map[string]DeploymentFlag
	genesis           time.Time

	accounts
	genesisBalances map[string]uint64

	bootnodes     int
	smeshers      int
	clients       []*NodeClient
	certifiers    []*NodeClient
	poets         []*NodeClient
	bootstrappers []*NodeClient
	postServices  []*NodeClient

	bootstrapEpochs []int
}

func (c *Cluster) Genesis() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *Cluster) GenesisExtraData() string { _ = "STUB: not implemented"; return "" }

// GenesisID computes id from the configuration.
func (c *Cluster) GenesisID() types.Hash20 { _ = "STUB: not implemented"; return *new(types.Hash20) }

func (c *Cluster) GoldenATX() types.ATXID { _ = "STUB: not implemented"; return *new(types.ATXID) }

func (c *Cluster) nextSmesher() int { _ = "STUB: not implemented"; return 0 }

func (c *Cluster) persist(ctx *testcontext.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cluster) persistConfigs(ctx *testcontext.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) persistFlags(ctx *testcontext.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) recoverFlags(ctx *testcontext.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) addFlag(flag DeploymentFlag) { _ = "STUB: not implemented"; return }

func (c *Cluster) addPoetFlag(flag DeploymentFlag) { _ = "STUB: not implemented"; return }

func (c *Cluster) addBootstrapperFlag(flag DeploymentFlag) { _ = "STUB: not implemented"; return }

func (c *Cluster) reuse(cctx *testcontext.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cluster) AddBootstrappers(cctx *testcontext.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// AddPoets spawns poets up to configured number of poets.
func (c *Cluster) AddPoets(cctx *testcontext.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cluster) firstFreePoetId() int { _ = "STUB: not implemented"; return 0 }

// AddCertifier spawns a single certifier with the first available id.
// Id is of form "certifier-N", where N ∈ [0, ∞).
func (c *Cluster) AddCertifier(cctx *testcontext.Context, privkey string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddPoet spawns a single poet with the first available id.
// Id is of form "poet-N", where N ∈ [0, ∞).
func (c *Cluster) AddPoet(cctx *testcontext.Context, flags ...DeploymentFlag) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) resourceControl(cctx *testcontext.Context, n int) error {
	_ = "STUB: not implemented"
	return nil
}

// maybe account for poet as well?

// AddBootnodes ...
func (c *Cluster) AddBootnodes(cctx *testcontext.Context, n int) error {
	_ = "STUB: not implemented"
	return nil
}

type SmesherDeploymentConfig struct {
	flags []DeploymentFlag
	keys  []ed25519.PrivateKey

	image          string
	noDefaultPoets bool
}

type DeploymentOpt func(cfg *SmesherDeploymentConfig)

func WithFlags(flags ...DeploymentFlag) DeploymentOpt {
	_ = "STUB: not implemented"
	return *new(DeploymentOpt)
}

func WithSmeshers(keys []ed25519.PrivateKey) DeploymentOpt {
	_ = "STUB: not implemented"
	return *new(DeploymentOpt)
}

func WithImage(image string) DeploymentOpt { _ = "STUB: not implemented"; return *new(DeploymentOpt) }

func NoDefaultPoets() DeploymentOpt { _ = "STUB: not implemented"; return *new(DeploymentOpt) }

// AddSmeshers ...
func (c *Cluster) AddSmeshers(tctx *testcontext.Context, n int, opts ...DeploymentOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) AddRemoteSmeshers(tctx *testcontext.Context, n int, opts ...DeploymentOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) AddBootstrapper(cctx *testcontext.Context, i int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) DeleteBootstrappers(cctx *testcontext.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// DeletePoets delete all poet servers.
func (c *Cluster) DeletePoets(cctx *testcontext.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) DeletePoet(cctx *testcontext.Context, i int) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteSmesher will smesher i from the cluster.
func (c *Cluster) DeleteSmesher(cctx *testcontext.Context, node *NodeClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Bootnodes returns number of the bootnodes in the cluster.
func (c *Cluster) Bootnodes() int {
	_ = "STUB: not implemented"

	// Smeshers returns number of the smeshers in the cluster.
	return 0
}

func (c *Cluster) Smeshers() int {
	_ = "STUB: not implemented"

	// Total returns total number of clients.
	return 0
}

func (c *Cluster) Total() int { _ = "STUB: not implemented"; return 0 }

// Poets returns total number of poet servers.
func (c *Cluster) Poets() int { _ = "STUB: not implemented"; return 0 }

// Poet returns client for i-th poet node.
func (c *Cluster) Poet(i int) *NodeClient {
	_ = "STUB: not implemented"

	// Client returns client for i-th node, either bootnode or smesher.
	return nil
}

func (c *Cluster) Client(i int) *NodeClient { _ = "STUB: not implemented"; return nil }

func (c *Cluster) Bootstrapper(i int) *NodeClient { _ = "STUB: not implemented"; return nil }

// Wait for i-th client to be up.
func (c *Cluster) Wait(tctx *testcontext.Context, i int) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitAll waits till (bootnode, smesher, poet, bootstrapper) pods are up.
func (c *Cluster) WaitAll(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// CloseClients closes connections to clients.
func (c *Cluster) CloseClients() { _ = "STUB: not implemented"; return }

// Account contains address and private key.
type Account struct {
	PrivateKey ed25519.PrivateKey
	Address    types.Address
}

func (a Account) String() string { _ = "STUB: not implemented"; return "" }

type accounts struct {
	keys      []*signer
	persisted bool
}

func (a *accounts) Account(i int) Account { _ = "STUB: not implemented"; return *new(Account) }

func (a *accounts) Accounts() int { _ = "STUB: not implemented"; return 0 }

func (a *accounts) Private(i int) ed25519.PrivateKey {
	_ = "STUB: not implemented"
	return *new(ed25519.PrivateKey)
}

func (a *accounts) Address(i int) types.Address {
	_ = "STUB: not implemented"
	return *new(types.Address)
}

func (a *accounts) Persist(ctx *testcontext.Context) error { _ = "STUB: not implemented"; return nil }

func (a *accounts) Recover(ctx *testcontext.Context) error { _ = "STUB: not implemented"; return nil }

type signer struct {
	Pub ed25519.PublicKey
	PK  ed25519.PrivateKey
}

func (s *signer) Address() types.Address { _ = "STUB: not implemented"; return *new(types.Address) }

func genSigners(n int) (rst []*signer) { _ = "STUB: not implemented"; return nil }

func genSigner() *signer { _ = "STUB: not implemented"; return nil }

func ExtractP2PEndpoints(tctx *testcontext.Context, nodes []*NodeClient) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func minPeers(size int) int { _ = "STUB: not implemented"; return 0 }

func persistFlags(ctx *testcontext.Context, name string, config map[string]DeploymentFlag) error {
	_ = "STUB: not implemented"
	return nil
}

func recoverFlags(ctx *testcontext.Context, name string) (map[string]DeploymentFlag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fillNetworkConfig(ctx *testcontext.Context, node *NodeClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) NodeConfig(ctx *testcontext.Context) (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
