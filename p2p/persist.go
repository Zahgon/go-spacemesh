package p2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"go.uber.org/zap"
)

const connectedFile = "connected.txt"

func persist(ctx context.Context, logger *zap.Logger, h host.Host, dir string, period time.Duration) {
	_ = "STUB: not implemented"
	return
}

func writePeers(h host.Host, dir string) error { _ = "STUB: not implemented"; return nil }

func loadPeers(dir string) ([]peer.AddrInfo, error) { _ = "STUB: not implemented"; return nil, nil }
