package sql

import (
	"context"
	"sync"

	"github.com/hashicorp/golang-lru/v2/simplelru"
)

const defaultLRUCacheSize = 100

type (
	QueryCacheKind   string
	inGetValueCtxKey struct{}
)

var NullQueryCache QueryCache = (*queryCache)(nil)

type QueryCacheItemKey struct {
	Kind QueryCacheKind
	Key  string
}

// QueryCacheKey creates a key for QueryCache.
func QueryCacheKey(kind QueryCacheKind, key string) QueryCacheItemKey {
	_ = "STUB: not implemented"
	return *new(QueryCacheItemKey)
}

type (
	// UntypedRetrieveFunc retrieves a value to be cached.
	UntypedRetrieveFunc func(ctx context.Context) (any, error)
	// SliceAppender modifies slice value stored in the cache, appending the
	// specified item to it and returns the updated slice.
	SliceAppender func(s any) any
)

// QueryCache stores results of SQL queries and data derived from these results.
// Presently, the cached entries are never removed, but eventually, it might
// become an LRU cache.
type QueryCache interface {
	// IsCached returns true if the requests are being cached.
	IsCached() bool
	// GetValue retrieves the specified value from the cache. If the entry is absent
	// from cache, it's populated by calling retrieve func.
	GetValue(
		ctx context.Context,
		key QueryCacheItemKey,
		retrieve UntypedRetrieveFunc,
	) (any, error)
}

// RetrieveFunc retrieves a value to be stored in the cache.
type RetrieveFunc[T any] func() (T, error)

// IsCached returns true if the database is cached.
func IsCached(db any) bool { _ = "STUB: not implemented"; return false }

// WithCachedValue retrieves the specified value from the cache. If the entry is
// absent from the cache, it's populated by calling retrieve func.
func WithCachedValue[T any](
	ctx context.Context,
	db any,
	key QueryCacheItemKey,
	retrieve func(ctx context.Context) (T, error),
) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

type lru = simplelru.LRU[string, any]

type queryCache struct {
	sync.Mutex
	updateMtx        sync.RWMutex
	cacheSizesByKind map[QueryCacheKind]int
	caches           map[QueryCacheKind]*lru
}

var _ QueryCache = &queryCache{}

func (c *queryCache) ensureLRU(kind QueryCacheKind) *lru { _ = "STUB: not implemented"; return nil }

func (c *queryCache) get(key QueryCacheItemKey) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (c *queryCache) set(key QueryCacheItemKey, v any) { _ = "STUB: not implemented"; return }

func (c *queryCache) IsCached() bool { _ = "STUB: not implemented"; return false }

func (c *queryCache) GetValue(
	ctx context.Context,
	key QueryCacheItemKey,
	retrieve UntypedRetrieveFunc,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Avoid recursive locking from within retrieve()

// This may seem like a race, but at worst, retrieve() will be
// called several times when populating this cached entry.
// That's better than locking for the duration of retrieve(),
// which can also refer to this cache
