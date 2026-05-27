package types

// ATXIDList defines ATX ID list.
type ATXIDList []ATXID

// Hash returns ATX ID list hash.
func (atxList ATXIDList) Hash() Hash32 { _ = "STUB: not implemented"; return *new(Hash32) }

// an error is never returned: https://golang.org/pkg/hash/#Hash
