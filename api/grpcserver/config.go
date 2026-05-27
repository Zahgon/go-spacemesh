// Package config provides configuration for GRPC and HTTP api servers
package grpcserver

import (
	"testing"
	"time"
)

type Config struct {
	PublicServices         []Service
	PublicListener         string `mapstructure:"grpc-public-listener"`
	PrivateServices        []Service
	PrivateListener        string `mapstructure:"grpc-private-listener"`
	PostServices           []Service
	PostListener           string    `mapstructure:"grpc-post-listener"`
	TLSServices            []Service `mapstructure:"grpc-tls-services"`
	TLSListener            string    `mapstructure:"grpc-tls-listener"`
	TLSCACert              string    `mapstructure:"grpc-tls-ca-cert"`
	TLSCert                string    `mapstructure:"grpc-tls-cert"`
	TLSKey                 string    `mapstructure:"grpc-tls-key"`
	GrpcSendMsgSize        int       `mapstructure:"grpc-send-msg-size"`
	GrpcRecvMsgSize        int       `mapstructure:"grpc-recv-msg-size"`
	JSONListener           string    `mapstructure:"grpc-json-listener"`
	JSONCorsAllowedOrigins []string  `mapstructure:"grpc-cors-allowed-origins"`

	SmesherStreamInterval time.Duration `mapstructure:"smesherstreaminterval"`

	DatabaseConnections int `mapstructure:"db-connections"`
}

type Service = string

const (
	// v1.
	Admin       Service = "admin"
	Debug       Service = "debug"
	GlobalState Service = "global"
	Mesh        Service = "mesh"
	Transaction Service = "transaction"
	Activation  Service = "activation"
	Smesher     Service = "smesher"
	Post        Service = "post"
	PostInfo    Service = "postInfo"
	Node        Service = "node"

	// v2alpha1.
	ActivationV2Alpha1        Service = "activation_v2alpha1"
	ActivationStreamV2Alpha1  Service = "activation_stream_v2alpha1"
	RewardV2Alpha1            Service = "reward_v2alpha1"
	RewardStreamV2Alpha1      Service = "reward_stream_v2alpha1"
	NetworkV2Alpha1           Service = "network_v2alpha1"
	NodeV2Alpha1              Service = "node_v2alpha1"
	LayerV2Alpha1             Service = "layer_v2alpha1"
	LayerStreamV2Alpha1       Service = "layer_stream_v2alpha1"
	TransactionV2Alpha1       Service = "transaction_v2alpha1"
	TransactionStreamV2Alpha1 Service = "transaction_stream_v2alpha1"
	AccountV2Alpha1           Service = "account_v2alpha1"
	MalfeasanceV2Alpha1       Service = "malfeasance_v2alpha1"
	MalfeasanceStreamV2Alpha1 Service = "malfeasance_stream_v2alpha1"

	// v2beta1.
	ActivationV2Beta1        Service = "activation_v2beta1"
	ActivationStreamV2Beta1  Service = "activation_stream_v2beta1"
	RewardV2Beta1            Service = "reward_v2beta1"
	RewardStreamV2Beta1      Service = "reward_stream_v2beta1"
	NetworkV2Beta1           Service = "network_v2beta1"
	NodeV2Beta1              Service = "node_v2beta1"
	LayerV2Beta1             Service = "layer_v2beta1"
	LayerStreamV2Beta1       Service = "layer_stream_v2beta1"
	TransactionV2Beta1       Service = "transaction_v2beta1"
	TransactionStreamV2Beta1 Service = "transaction_stream_v2beta1"
	AccountV2Beta1           Service = "account_v2beta1"
	MalfeasanceV2Beta1       Service = "malfeasance_v2beta1"
	MalfeasanceStreamV2Beta1 Service = "malfeasance_stream_v2beta1"
)

// DefaultConfig defines the default configuration options for api.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// v1

// v2alpha1

// v2beta1

// v1

// v2alpha1

// v2beta1

// DefaultTestConfig returns the default config for tests.
func DefaultTestConfig(tb testing.TB) Config { _ = "STUB: not implemented"; return *new(Config) }
