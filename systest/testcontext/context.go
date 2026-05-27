package testcontext

import (
	"context"
	"flag"
	"sync"
	"testing"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spacemeshos/go-spacemesh/systest/parameters"
)

const (
	ParamLayersPerEpoch = "layers-per-epoch"
	ParamLayerDuration  = "layer-duration"
)

var (
	configname = flag.String(
		"configname",
		"",
		"config map name. if not empty parameters will be loaded from specified configmap",
	)
	clusters = flag.Int(
		"clusters",
		1,
		"controls tests parallelization by creating multiple spacemesh clusters at the same time",
	)
	logLevel    = zap.LevelFlag("level", zap.InfoLevel, "verbosity of the logger")
	testTimeout = flag.Duration("test-timeout", 60*time.Minute, "timeout for a single test")

	tokens     chan struct{}
	initTokens sync.Once

	failed   = make(chan struct{})
	failOnce sync.Once
)

var (
	testid = parameters.String(
		"testid", "Name of the pod that runs tests.", "",
	)
	imageFlag = parameters.String(
		"image",
		"go-spacemesh image",
		"", // repo doesn't have a `latest` tag we can default to
	)
	oldImageFlag = parameters.String(
		"old-image",
		"old go-spacemesh image to test compatibility against",
		"spacemeshos/go-spacemesh-dev:v1.7.4", // repo doesn't have a `latest` tag we can default to
	)
	bsImage = parameters.String(
		"bs-image",
		"bootstrapper image",
		"", // repo doesn't have a `latest` tag we can default to
	)
	certifierImage = parameters.String(
		"certifier-image",
		"certifier service image",
		"spacemeshos/certifier-service:latest",
	)
	poetImage = parameters.String(
		"poet-image",
		"poet server image",
		"spacemeshos/poet:latest",
	)
	postServiceImage = parameters.String(
		"post-service-image",
		"post service image",
		"spacemeshos/post-service:latest",
	)
	postInitImage = parameters.String(
		"post-init-image",
		"post init image",
		"spacemeshos/postcli:latest",
	)
	namespaceFlag = parameters.String(
		"namespace",
		"namespace for the cluster. if empty every test will use random namespace",
		"",
	)
	bootstrapDuration = parameters.Duration(
		"bootstrap-duration",
		"bootstrap time is added to genesis time. "+
			"it may take longer in cloud environments due to additional resource management",
		30*time.Second,
	)
	clusterSize = parameters.Int(
		"cluster-size",
		"size of the cluster. all test must use at most this number of smeshers",
		10,
	)
	poetSize = parameters.Int(
		"poet-size", "size of the poet servers", 2,
	)
	bsSize = parameters.Int(
		"bs-size", "size of bootstrappers", 1,
	)
	storage = parameters.String(
		"storage", "<class>=<size> for the storage", "standard=1Gi",
	)
	keep = parameters.Bool(
		"keep", "if true cluster will not be removed after test is finished",
	)
	nodeSelector = parameters.NewParameter(
		"node-selector", "select where test pods will be scheduled",
		stringToString{},
		func(value string) (stringToString, error) {
			rst := stringToString{}
			if err := rst.Set(value); err != nil {
				return nil, err
			}
			return rst, nil
		},
	)
	LayerDuration = parameters.Duration(
		ParamLayerDuration, "layer duration in seconds", 5*time.Second,
	)
	LayersPerEpoch = parameters.Int(
		ParamLayersPerEpoch, "number of layers in an epoch", 4,
	)
)

const nsfile = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

const (
	keepLabel     = "keep"
	poetSizeLabel = "poet-size"
)

func rngName() string { _ = "STUB: not implemented"; return "" }

// Context must be created for every test that needs isolated cluster.
type Context struct {
	context.Context
	Client            *kubernetes.Clientset
	Parameters        *parameters.Parameters
	BootstrapDuration time.Duration
	ClusterSize       int
	BootnodeSize      int
	RemoteSize        int
	PoetSize          int
	OldSize           int
	BootstrapperSize  int
	Generic           client.Client
	TestID            string
	Keep              bool
	Namespace         string
	Image             string
	OldImage          string
	BootstrapperImage string
	CertifierImage    string
	PoetImage         string
	PostServiceImage  string
	PostInitImage     string
	Storage           struct {
		Size  string
		Class string
	}
	NodeSelector map[string]string
	Log          *zap.SugaredLogger
}

func cleanup(tb testing.TB, f func()) { _ = "STUB: not implemented"; return }

func deleteNamespace(ctx *Context) error { _ = "STUB: not implemented"; return nil }

func deployNamespace(ctx *Context) error { _ = "STUB: not implemented"; return nil }

func getStorage(p *parameters.Parameters) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func updateContext(ctx *Context) error { _ = "STUB: not implemented"; return nil }

// SkipClusterLimits will not block if there are no available tokens.
func SkipClusterLimits() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// Opt is for configuring Context.
type Opt func(*cfg)

func newCfg() *cfg { _ = "STUB: not implemented"; return nil }

type cfg struct {
	skipLimits bool
}

// New creates context for the test.
func New(t *testing.T, opts ...Opt) *Context { _ = "STUB: not implemented"; return nil }

// The default rate limiter is too slow 5qps and 10 burst, This will prevent the client from being throttled
// Change the limits to the same of kubectl and argo
// That's were those number come from
// https://github.com/kubernetes/kubernetes/pull/105520
// https://github.com/argoproj/argo-workflows/pull/11603/files

// prevent sigs.k8s.io/controller-runtime from complaining about log.SetLogger never being called

// nolint: usetesting

// 50% of smeshers are remote

// 25% of smeshers are old (use previous version of go-spacemesh)

func (c *Context) CheckFail() error { _ = "STUB: not implemented"; return nil }
