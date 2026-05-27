package multipeer

import (
	"github.com/spacemeshos/go-spacemesh/sync2/rangesync"
)

// getDelimiters generates keys that can be used as range delimiters for splitting the key
// space among the specified number of peers. maxDepth specifies maximum number of high
// non-zero bits to include in resulting keys, which helps avoiding unaligned splits which
// are more expensive for FPTree data structure.
// keyLen specifies key length in bytes.
// The function returns numPeers-1 keys. The ranges are used for split sync, with each
// range being assigned to a separate peer. The first range begins with zero-valued key
// (k0), represented by KeyBytes of length keyLen consisting entirely of zeroes.
// The ranges to scan are:
// [k0,ks[0]); [k0,ks[1]); ... [k0,ks[numPeers-2]); [ks[numPeers-2],0).
func getDelimiters(numPeers, keyLen, maxDepth int) (ks []rangesync.KeyBytes) {
	_ = "STUB: not implemented"
	return nil
}
