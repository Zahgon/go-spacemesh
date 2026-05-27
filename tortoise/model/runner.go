package model

import (
	"math/rand"
	"reflect"
	"testing"

	"go.uber.org/zap"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

type model interface {
	OnMessage(Messenger, Message)
}

func newCluster(tb testing.TB, logger *zap.Logger, rng *rand.Rand) *cluster {
	_ = "STUB: not implemented"
	return nil
}

type cluster struct {
	tb     testing.TB
	rng    *rand.Rand
	logger *zap.Logger
	models []model
}

func (r *cluster) nextid() string { _ = "STUB: not implemented"; return "" }

func (r *cluster) add(m model) *cluster { _ = "STUB: not implemented"; return nil }

func (r *cluster) addCore() *cluster { _ = "STUB: not implemented"; return nil }

func (r *cluster) addHare() *cluster { _ = "STUB: not implemented"; return nil }

func (r *cluster) addBeacon() *cluster { _ = "STUB: not implemented"; return nil }

func (r *cluster) iterate(f func(m model)) { _ = "STUB: not implemented"; return }

func newFailingRunner(c *cluster,
	messenger Messenger,
	monitors []Monitor,
	rng *rand.Rand,
	probability [2]int,
) *failingRunner {
	_ = "STUB: not implemented"
	return nil
}

type failingRunner struct {
	cluster   *cluster
	messenger Messenger
	monitors  []Monitor
	// probability that messages in failables will fail by the model
	probability [2]int
	rng         *rand.Rand
	lid         types.LayerID
	failables   map[reflect.Type]struct{}
}

func (r *failingRunner) next() { _ = "STUB: not implemented"; return }

func (r *failingRunner) failable(events ...Message) *failingRunner {
	_ = "STUB: not implemented"
	return nil
}

func (r *failingRunner) isFailable(ev Message) bool { _ = "STUB: not implemented"; return false }

func (r *failingRunner) consume() { _ = "STUB: not implemented"; return }
