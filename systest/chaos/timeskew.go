package chaos

import (
	"github.com/spacemeshos/go-spacemesh/systest/testcontext"
)

// Timeskew adjusts CLOCK_REALTIME on the specified pods by the offset.
func Timeskew(cctx *testcontext.Context, name, offset string, pods ...string) (Teardown, error) {
	_ = "STUB: not implemented"
	return *new(Teardown), nil
}
