package cluster

import (
	"context"
	"sync"
	"time"

	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
	"google.golang.org/grpc"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/systest/parameters"
	"github.com/spacemeshos/go-spacemesh/systest/parameters/fastnet"
	"github.com/spacemeshos/go-spacemesh/systest/testcontext"
)

var (
	certifierConfig = parameters.String(
		"certifier",
		"configuration for certifier service",
		fastnet.CertifierConfig,
	)
	poetConfig = parameters.String(
		"poet",
		"configuration for poet service",
		fastnet.PoetConfig,
	)
	smesherConfig = parameters.String(
		"smesher",
		"configuration for smesher service",
		fastnet.SmesherConfig,
	)

	smesherResources = parameters.NewParameter(
		"smesher_resources",
		"requests and limits for smesher container",
		&apiv1.ResourceRequirements{
			Requests: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("1.3"),
				apiv1.ResourceMemory: resource.MustParse("800Mi"),
			},
			Limits: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("1.3"),
				apiv1.ResourceMemory: resource.MustParse("800Mi"),
			},
		},
		toResources,
	)
	bootstrapperResources = parameters.NewParameter(
		"bootstrapper_resources",
		"requests and limits for bootstrapper container",
		&apiv1.ResourceRequirements{
			Requests: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("0.1"),
				apiv1.ResourceMemory: resource.MustParse("100Mi"),
			},
			Limits: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("0.1"),
				apiv1.ResourceMemory: resource.MustParse("100Mi"),
			},
		},
		toResources,
	)
	poetResources = parameters.NewParameter(
		"poet_resources",
		"requests and limits for poet container",
		&apiv1.ResourceRequirements{
			Requests: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("0.4"),
				apiv1.ResourceMemory: resource.MustParse("1Gi"),
			},
			Limits: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("0.4"),
				apiv1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		toResources,
	)
)

func toResources(value string) (*apiv1.ResourceRequirements, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	configDir = "/etc/config/"

	attachedCertifierConfig = "certifier.yaml"
	attachedPoetConfig      = "poet.conf"
	attachedSmesherConfig   = "smesher.json"

	certifierConfigMapName = "certifier"
	poetConfigMapName      = "poet"
	spacemeshConfigMapName = "spacemesh"

	// smeshers are split in 10 approximately equal buckets
	// to enable running chaos mesh tasks on the different parts of the cluster.
	buckets = 10
)

const (
	prometheusScrapePort = 9216
	phlareScrapePort     = 6060
)

// Node ...
type Node struct {
	Name                     string
	P2P, GRPC_PUB, GRPC_PRIV uint16
}

// P2PEndpoint returns full p2p endpoint, including identity.
func p2pEndpoint(n Node, ip, id string) string { _ = "STUB: not implemented"; return "" }

// NodeClient is a Node with attached grpc connection.
type NodeClient struct {
	session *testcontext.Context
	Node

	mu       sync.Mutex
	pubConn  *grpc.ClientConn
	privConn *grpc.ClientConn
}

func (n *NodeClient) Close() { _ = "STUB: not implemented"; return }

func (n *NodeClient) Resolve(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (n *NodeClient) ensurePubConn(ctx context.Context) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// waitForConnectionReady blocks until the connection is ready. It returns an error if the context is canceled or
// timed out.
//
// This doesn't guarantee that the connection will stay ready, but it makes it so that the test runner waits at least
// until the nodes are started before querying them.
func (n *NodeClient) waitForConnectionReady(ctx context.Context, conn *grpc.ClientConn) error {
	_ = "STUB: not implemented"
	// A blocking dial blocks until the clientConn is ready.
	return nil
}

// ctx got timeout or canceled.

func (n *NodeClient) resetPubConn(conn *grpc.ClientConn) { _ = "STUB: not implemented"; return }

func (n *NodeClient) ensurePrivConn(ctx context.Context) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NodeClient) resetPrivConn(conn *grpc.ClientConn) { _ = "STUB: not implemented"; return }

func (n *NodeClient) PubConn() *PubNodeClient { _ = "STUB: not implemented"; return nil }

type PubNodeClient struct {
	*NodeClient
}

func (n *PubNodeClient) Invoke(ctx context.Context, method string, args, reply any, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

// check for app error. this is not exhaustive.
// the goal is to reset connection if pods were redeployed and changed IP

func (n *PubNodeClient) NewStream(
	ctx context.Context,
	desc *grpc.StreamDesc,
	method string,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

func (n *NodeClient) PrivConn() *PrivNodeClient { _ = "STUB: not implemented"; return nil }

type PrivNodeClient struct {
	*NodeClient
}

func (n *PrivNodeClient) Invoke(ctx context.Context, method string, args, reply any, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

// check for app error. this is not exhaustive.
// the goal is to reset connection if pods were redeployed and changed IP

func (n *PrivNodeClient) NewStream(
	ctx context.Context,
	desc *grpc.StreamDesc,
	method string,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

func deployCertifierD(ctx *testcontext.Context, id, privkey string) (*NodeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deployPoetD(ctx *testcontext.Context, id string, flags ...DeploymentFlag) (*NodeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deployBootnodeSvc(ctx *testcontext.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func deployNodeSvc(ctx *testcontext.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func deployCertifierSvc(ctx *testcontext.Context, id string) (*apiv1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deployPoetSvc(ctx *testcontext.Context, id string) (*apiv1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPoetIdentifier(id int) string { _ = "STUB: not implemented"; return "" }

func createBootstrapperIdentifier(id int) string { _ = "STUB: not implemented"; return "" }

func decodePoetIdentifier(id string) int { _ = "STUB: not implemented"; return 0 }

// deployCertifier creates a certifier Deployment and exposes it via a Service.
// The key is passed to the certifier Pod.
func deployCertifier(ctx *testcontext.Context, id, privkey string) (*NodeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// deployPoet creates a poet Pod and exposes it via a Service.
// Flags are passed to the poet Pod as arguments.
func deployPoet(ctx *testcontext.Context, id string, flags ...DeploymentFlag) (*NodeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteServiceAndPod(ctx *testcontext.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// areContainersReady checks if all containers are ready in pod.
func areContainersReady(pod *apiv1.Pod) bool { _ = "STUB: not implemented"; return false }

func waitPod(ctx *testcontext.Context, id string) (*apiv1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nodeLabels(name, id string) map[string]string { _ = "STUB: not implemented"; return nil }

// app identifies resource kind (Node, Poet).
// It can be used to select all Pods of given kind.

// id uniquely identifies a resource (i.e. poet-0).

func labelSelector(id string) string { _ = "STUB: not implemented"; return "" }

func deployNodes(ctx *testcontext.Context, kind string, from, to int, opts ...DeploymentOpt) ([]*NodeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deployRemoteNodes(
	ctx *testcontext.Context,
	from, to int,
	goldenAtxId types.ATXID,
	opts ...DeploymentOpt,
) ([]*NodeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteNode(ctx *testcontext.Context, id string) error {
	_ = "STUB: not implemented"
	// find and delete any post services linked to this node
	return nil
}

func deployNode(
	ctx *testcontext.Context,
	id string,
	key ed25519.PrivateKey,
	image string,
	keyName string,
	labels map[string]string,
	flags []DeploymentFlag,
) error {
	_ = "STUB: not implemented"
	return nil
}

func deployPostService(
	ctx *testcontext.Context,
	id string,
	labels map[string]string,
	nodeId string,
	pubKey string,
	goldenAtxId string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// 0xffffffff = CPU Provider
// to prevent checks for mainnet compatibility

func deployBootstrapper(
	ctx *testcontext.Context,
	id string,
	bsEpochs []int,
	flags ...DeploymentFlag,
) (*NodeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deployBootstrapperSvc(ctx *testcontext.Context, id string) (*apiv1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func commaSeparatedList(epochs []int) string { _ = "STUB: not implemented"; return "" }

func deployBootstrapperD(
	ctx *testcontext.Context,
	id string,
	bsEpochs []int,
	flags ...DeploymentFlag,
) (*NodeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// empty so it generates local random beacon instead of making http queries to bitcoin explorer

// DeploymentFlag allows to configure specific flags for application binaries.
type DeploymentFlag struct {
	Name, Value string
}

// Flag returns parseable flag from Name and Value.
func (d DeploymentFlag) Flag() string { _ = "STUB: not implemented"; return "" }

func PoetEndpoints(ids ...int) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

func PostK3(k3 int) DeploymentFlag { _ = "STUB: not implemented"; return *new(DeploymentFlag) }

// MinPeers flag.
func MinPeers(target int) DeploymentFlag { _ = "STUB: not implemented"; return *new(DeploymentFlag) }

func BootstrapperUrl(endpoint string) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

func CheckpointUrl(endpoint string) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

func CheckpointLayer(restoreLayer uint32) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

const (
	genesisTimeFlag  = "--genesis-time"
	genesisExtraData = "--genesis-extra-data"
	accountsFlag     = "--accounts"
)

// GenesisTime flag.
func GenesisTime(t time.Time) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

// GenesisExtraData flag.
func GenesisExtraData(extra string) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

// Bootnodes flag.
func Bootnodes(bootnodes ...string) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

// Accounts flag.
func Accounts(accounts map[string]uint64) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

// DurationFlag is a generic duration flag.
func DurationFlag(flag string, d time.Duration) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

func PoetCertifierURL(url string) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

func PoetCertifierPubkey(key string) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

// PoetRestListen socket pair with http api.
func PoetRestListen(port int) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

func StartSmeshing(start bool) DeploymentFlag {
	_ = "STUB: not implemented"
	return *new(DeploymentFlag)
}

func GenerateFallback() DeploymentFlag { _ = "STUB: not implemented"; return *new(DeploymentFlag) }

func Bootnode() DeploymentFlag { _ = "STUB: not implemented"; return *new(DeploymentFlag) }

func PrivateNetwork() DeploymentFlag { _ = "STUB: not implemented"; return *new(DeploymentFlag) }
