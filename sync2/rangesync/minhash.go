package rangesync

import (
	"github.com/spacemeshos/go-scale"
)

// MinhashSampleItem represents an item of minhash sample subset.
type MinhashSampleItem uint32

func (m MinhashSampleItem) String() string { _ = "STUB: not implemented"; return "" }

func (m MinhashSampleItem) Compare(other MinhashSampleItem) int {
	_ = "STUB: not implemented"
	return 0
}

// EncodeScale implements scale.Encodable.
func (m MinhashSampleItem) EncodeScale(e *scale.Encoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DecodeScale implements scale.Decodable.
func (m *MinhashSampleItem) DecodeScale(d *scale.Decoder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MinhashSampleItemFromKeyBytes uses lower 32 bits of a hash as a MinhashSampleItem.
func MinhashSampleItemFromKeyBytes(h KeyBytes) MinhashSampleItem {
	_ = "STUB: not implemented"
	return *new(MinhashSampleItem)
}

// Sample retrieves min(count, sampleSize) items friom the ordered sequence, extracting
// MinhashSampleItem from each value.
func Sample(sr SeqResult, count, sampleSize int) ([]MinhashSampleItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CalcSim estimates the Jaccard similarity coefficient between two sets based on the
// samples, which are derived from N lowest-valued elements of each set.
// The return value is in 0..1 range, with 0 meaning almost no intersection and 1 meaning
// the sets are mostly equal.
// The precision of the estimate will suffer if none of the sets are empty and they have
// different size, with return value tending to be lower than the actual J coefficient.
func CalcSim(a, b []MinhashSampleItem) float64 { _ = "STUB: not implemented"; return 0 }

// The sample items contain the lowest bits of the actual hash values, so their
// order is initially random. We sort them here for easier comparison of the
// samples.
