package main

import (
	"context"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/afero"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

const (
	SchemaVersion = "https://spacemesh.io/bootstrap.schema.json.1.0"
	confirmation  = 7
	timeout       = 5 * time.Second
)

func PersistedFilename(epoch types.EpochID, suffix string) string {
	_ = "STUB: not implemented"
	return ""
}

type Generator struct {
	logger      *zap.Logger
	fs          afero.Fs
	client      *retryablehttp.Client
	btcEndpoint string
	smEndpoint  string
}

type Opt func(*Generator)

func WithLogger(logger *zap.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithFilesystem(fs afero.Fs) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func NewGenerator(btcEndpoint, smEndpoint string, opts ...Opt) *Generator {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) SmEndpoint() string { _ = "STUB: not implemented"; return "" }

func (g *Generator) Generate(
	ctx context.Context,
	targetEpoch types.EpochID,
	genBeacon, genActiveSet bool,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// BitcoinResponse captures the only fields we care about from a bitcoin block.
type BitcoinResponse struct {
	Height uint64 `json:"height"`
	Hash   string `json:"hash"`
}

func (g *Generator) genBeacon(ctx context.Context, logger *zap.Logger) (types.Beacon, error) {
	_ = "STUB: not implemented"
	return *new(types.Beacon), nil
}

// bitcoin hash started with leading zero. we want to grab 4 LSB

func bitcoinHash(
	ctx context.Context,
	logger *zap.Logger,
	client *retryablehttp.Client,
	targetUrl string,
) (*BitcoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryBitcoin(ctx context.Context, client *retryablehttp.Client, targetUrl string) (*BitcoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// some apis got over sensitive without UA set

func getActiveSet(ctx context.Context, endpoint string, epoch types.EpochID) ([]types.ATXID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) GenUpdate(
	epoch types.EpochID,
	beacon types.Beacon,
	activeSet []types.ATXID,
	suffix string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// no leading 0x

// no leading 0x

// make sure the data is valid
