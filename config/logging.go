package config

import "go.uber.org/zap/zapcore"

// LogEncoder defines a log encoder kind.
type LogEncoder = string

const (
	defaultLoggingLevel = zapcore.InfoLevel
	// ConsoleLogEncoder represents logging with plain text.
	ConsoleLogEncoder LogEncoder = "console"
	// JSONLogEncoder represents logging with JSON.
	JSONLogEncoder LogEncoder = "json"
)

// LoggerConfig holds the logging level for each module.
type LoggerConfig struct {
	Encoder LogEncoder `mapstructure:"log-encoder"`

	AppLoggerLevel             string `mapstructure:"app"`
	ClockLoggerLevel           string `mapstructure:"clock"`
	P2PLoggerLevel             string `mapstructure:"p2p"`
	PostLoggerLevel            string `mapstructure:"post"`
	PostServiceLoggerLevel     string `mapstructure:"postService"`
	StateDbLoggerLevel         string `mapstructure:"stateDb"`
	BeaconLoggerLevel          string `mapstructure:"beacon"`
	CachedDBLoggerLevel        string `mapstructure:"cachedDb"`
	PoetDbLoggerLevel          string `mapstructure:"poetDb"`
	TrtlLoggerLevel            string `mapstructure:"trtl"`
	AtxHandlerLevel            string `mapstructure:"atxHandler"`
	AtxBuilderLoggerLevel      string `mapstructure:"atxBuilder"`
	MeshLoggerLevel            string `mapstructure:"mesh"`
	SyncLoggerLevel            string `mapstructure:"sync"`
	HareOracleLoggerLevel      string `mapstructure:"hareOracle"`
	HareLoggerLevel            string `mapstructure:"hare"`
	BlockCertLoggerLevel       string `mapstructure:"blockCert"`
	BlockGenLoggerLevel        string `mapstructure:"blockGenerator"`
	BlockHandlerLoggerLevel    string `mapstructure:"blockHandler"`
	TxHandlerLoggerLevel       string `mapstructure:"txHandler"`
	ProposalStoreLoggerLevel   string `mapstructure:"proposalStore"`
	ProposalBuilderLoggerLevel string `mapstructure:"proposalBuilder"`
	ProposalListenerLevel      string `mapstructure:"proposalListener"`
	NipostBuilderLoggerLevel   string `mapstructure:"nipostBuilder"`
	NipostValidatorLoggerLevel string `mapstructure:"nipostValidator"`
	FetcherLoggerLevel         string `mapstructure:"fetcher"`
	TimeSyncLoggerLevel        string `mapstructure:"timesync"`
	VMLogLevel                 string `mapstructure:"vm"`
	GrpcLoggerLevel            string `mapstructure:"grpc"`
	ConStateLoggerLevel        string `mapstructure:"conState"`
	ExecutorLoggerLevel        string `mapstructure:"executor"`
	MalfeasanceLoggerLevel     string `mapstructure:"malfeasance"`
	Malfeasance2LoggerLevel    string `mapstructure:"malfeasance2"`
	BootstrapLoggerLevel       string `mapstructure:"bootstrap"`
}

func DefaultLoggingConfig() LoggerConfig { _ = "STUB: not implemented"; return *new(LoggerConfig) }
