package multipeer

import (
	"container/list"
	"sync"
	"time"

	"github.com/jonboulle/clockwork"
)

// syncList keeps track of recent full syncs and reports whether the node is synced, that
// is, the specified number of syncs has happened within the specified duration of time.
type syncList struct {
	mtx          sync.Mutex
	clock        clockwork.Clock
	minSyncCount int
	duration     time.Duration
	syncs        list.List
}

func newSyncList(clock clockwork.Clock, minSyncCount int, duration time.Duration) *syncList {
	_ = "STUB: not implemented"
	return nil
}

func (sl *syncList) prune(now time.Time) { _ = "STUB: not implemented"; return }

func (sl *syncList) NoteSync() { _ = "STUB: not implemented"; return }

func (sl *syncList) Synced() bool { _ = "STUB: not implemented"; return false }

func (sl *syncList) Len() int { _ = "STUB: not implemented"; return 0 }
