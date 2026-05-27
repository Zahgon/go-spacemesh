package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

const (
	retries       = 3
	retryInterval = 3 * time.Second
)

var (
	bitcoinEndpoint   string
	spacemeshEndpoint string
	genBeacon         bool
	genActiveSet      bool
	out               string
	creds             string
	serveUpdate       bool
	port              int
	epochOffset       uint32
	genFallback       bool
	dataDir           string
	logLevel          string
)

func init() {
	cmd.PersistentFlags().StringVar(&bitcoinEndpoint, "bitcoin-endpoint",
		"https://api.blockcypher.com/v1/btc/main", "URL to get bitcoin block hash")
	cmd.PersistentFlags().StringVar(&spacemeshEndpoint, "spacemesh-endpoint", "", "grpc endpoint for a spacemesh node")

	// options specific to one-time execution
	cmd.PersistentFlags().BoolVar(&genBeacon, "beacon", false, "generate beacon")
	cmd.PersistentFlags().BoolVar(&genActiveSet, "actives", false, "generate active set")
	cmd.PersistentFlags().StringVar(&out, "out", "gs://my-bucket", "gs URI for upload")
	cmd.PersistentFlags().StringVar(&creds, "creds", "", "path to gcloud credential file")

	// for systests only
	cmd.PersistentFlags().BoolVar(&serveUpdate, "serve-update",
		false, "if true, starts a http server to serve update too")
	cmd.PersistentFlags().IntVar(&port, "port",
		8080, "if starting a server, the port number to use")
	cmd.PersistentFlags().Uint32Var(&epochOffset, "epoch-offset",
		1, "number of layers before the next epoch start to publish update")
	cmd.PersistentFlags().BoolVar(&genFallback, "fallback", false,
		"in addition to bootstrap data, also generate fallback data")

	// admin
	cmd.PersistentFlags().StringVar(&dataDir, "data-dir", os.TempDir(), "directory to persist update data")
	cmd.PersistentFlags().StringVar(&logLevel, "level", "info", "logging level")
}

var cmd = &cobra.Command{
	Use:   "bootstrapper",
	Short: "generate bootstrapping data",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("epoch not specified")
		}
		var targetEpochs []types.EpochID
		epochs := strings.Split(args[0], ",")
		for _, e := range epochs {
			epoch, err := strconv.Atoi(e)
			if err != nil {
				return fmt.Errorf("cannot convert %v to epoch: %w", e, err)
			}
			targetEpochs = append(targetEpochs, types.EpochID(epoch))
		}

		lvl, err := zap.ParseAtomicLevel(logLevel)
		if err != nil {
			return err
		}

		logger, err := zap.Config{
			Level:            lvl,
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}.Build()
		if err != nil {
			return fmt.Errorf("creating logger: %w", err)
		}

		g := NewGenerator(
			bitcoinEndpoint,
			spacemeshEndpoint,
			WithLogger(logger.Named("generator")),
		)

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer cancel()
		if serveUpdate {
			srv := NewServer(g, genFallback, port,
				WithSrvFilesystem(afero.NewOsFs()),
				WithSrvLogger(logger.Named("server")),
				WithBootstrapEpochs(targetEpochs),
			)
			return runServer(ctx, srv)
		}

		if len(targetEpochs) != 1 {
			return errors.New("too many epochs specified")
		}
		// one-time execution
		if !genBeacon && !genActiveSet {
			return errors.New("no action specified via --beacon or --actives")
		}
		if genBeacon && len(bitcoinEndpoint) == 0 {
			return errors.New("missing bitcoin endpoint for beacon generation")
		}
		if genActiveSet && len(spacemeshEndpoint) == 0 {
			return errors.New("missing spacemesh endpoint for active set generation")
		}
		gsBucket, gsPath, err := parseToGsBucket(out)
		if err != nil {
			return fmt.Errorf("parse output uri %v: %w", out, err)
		}
		persisted, err := g.Generate(ctx, targetEpochs[0], genBeacon, genActiveSet)
		if err != nil {
			return err
		}
		return upload(ctx, persisted, gsBucket, gsPath)
	},
}

func upload(ctx context.Context, filename, gsBucket, gsPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseToGsBucket(gsPath string) (bucket, path string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// remove leading "/" in URL path

func runServer(ctx context.Context, srv *Server) error { _ = "STUB: not implemented"; return nil }

func queryNetworkParams(ctx context.Context, endpoint string) (*NetworkParam, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
