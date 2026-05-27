package sim

import (
	"math/rand"
	"testing"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/datastore"
)

const (
	atxpath = "atx"
)

func newCacheDB(tb testing.TB, logger *zap.Logger, conf config) *datastore.CachedDB {
	_ = "STUB: not implemented"
	return nil
}

func intInRange(rng *rand.Rand, ints [2]int) uint32 { _ = "STUB: not implemented"; return 0 }
