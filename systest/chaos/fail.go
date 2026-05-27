package chaos

import (
	"context"

	chaos "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"

	"github.com/spacemeshos/go-spacemesh/systest/testcontext"
)

func selectPods(pods []string) chaos.PodSelectorSpec {
	_ = "STUB: not implemented"
	return *new(chaos.PodSelectorSpec)
}

// Teardown is returned by every chaos action and executed
// by the caller once chaos needs to be stopped.
type Teardown func(context.Context) error

// Fail the list of pods and prevents them from respawning until teardown is called.
func Fail(cctx *testcontext.Context, name string, pods ...string) (Teardown, error) {
	_ = "STUB: not implemented"
	return *new(Teardown), nil
}
