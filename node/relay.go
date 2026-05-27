package node

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/spacemeshos/go-spacemesh/config"
)

var relayAddrInfoCh chan peer.AddrInfo // used for testing

func runRelay(ctx context.Context, cfg *config.Config) error { _ = "STUB: not implemented"; return nil }

// Prevent testnet nodes from working on the mainnet, but
// don't use the network cookie on mainnet as this technique
// may be replaced later
