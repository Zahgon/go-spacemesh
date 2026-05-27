package rangesync

import (
	"time"
)

//go:generate scalegen

// EmptyRangeMessage notifies the peer that it needs to send all of its items in
// the specified range.
type EmptyRangeMessage struct {
	RangeX, RangeY CompactHash
}

var _ SyncMessage = &EmptyRangeMessage{}

func (m *EmptyRangeMessage) Type() MessageType { _ = "STUB: not implemented"; return *new(MessageType) }
func (m *EmptyRangeMessage) X() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *EmptyRangeMessage) Y() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *EmptyRangeMessage) Fingerprint() Fingerprint {
	_ = "STUB: not implemented"
	return *new(Fingerprint)
}
func (m *EmptyRangeMessage) Count() int       { _ = "STUB: not implemented"; return 0 }
func (m *EmptyRangeMessage) Keys() []KeyBytes { _ = "STUB: not implemented"; return nil }
func (m *EmptyRangeMessage) Since() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (m *EmptyRangeMessage) Sample() []MinhashSampleItem {
	_ = "STUB: not implemented"

	// FingerprintMessage contains range fingerprint for comparison against the
	// peer's fingerprint of the range with the same bounds [RangeX, RangeY).
	return nil
}

type FingerprintMessage struct {
	RangeX, RangeY   CompactHash
	RangeFingerprint Fingerprint
	NumItems         uint32
}

var _ SyncMessage = &FingerprintMessage{}

func (m *FingerprintMessage) Type() MessageType {
	_ = "STUB: not implemented"
	return *new(MessageType)
}
func (m *FingerprintMessage) X() KeyBytes { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *FingerprintMessage) Y() KeyBytes { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *FingerprintMessage) Fingerprint() Fingerprint {
	_ = "STUB: not implemented"
	return *new(Fingerprint)
}
func (m *FingerprintMessage) Count() int       { _ = "STUB: not implemented"; return 0 }
func (m *FingerprintMessage) Keys() []KeyBytes { _ = "STUB: not implemented"; return nil }
func (m *FingerprintMessage) Since() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (m *FingerprintMessage) Sample() []MinhashSampleItem {
	_ = "STUB: not implemented"

	// RangeContentsMessage denotes a range for which the set of items has been sent.
	// The peer needs to send back any items it has in the same range bounded
	// by [RangeX, RangeY).
	return nil
}

type RangeContentsMessage struct {
	RangeX, RangeY CompactHash
	NumItems       uint32
}

var _ SyncMessage = &RangeContentsMessage{}

func (m *RangeContentsMessage) Type() MessageType {
	_ = "STUB: not implemented"
	return *new(MessageType)
}
func (m *RangeContentsMessage) X() KeyBytes { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *RangeContentsMessage) Y() KeyBytes { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *RangeContentsMessage) Fingerprint() Fingerprint {
	_ = "STUB: not implemented"
	return *new(Fingerprint)
}
func (m *RangeContentsMessage) Count() int       { _ = "STUB: not implemented"; return 0 }
func (m *RangeContentsMessage) Keys() []KeyBytes { _ = "STUB: not implemented"; return nil }
func (m *RangeContentsMessage) Since() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (m *RangeContentsMessage) Sample() []MinhashSampleItem {
	_ = "STUB: not implemented"

	// ItemBatchMessage denotes a batch of items to be added to the peer's set.
	return nil
}

type ItemBatchMessage struct {
	ContentKeys KeyCollection `scale:"max=1024"`
}

var _ SyncMessage = &ItemBatchMessage{}

func (m *ItemBatchMessage) Type() MessageType { _ = "STUB: not implemented"; return *new(MessageType) }
func (m *ItemBatchMessage) X() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *ItemBatchMessage) Y() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *ItemBatchMessage) Fingerprint() Fingerprint {
	_ = "STUB: not implemented"
	return *new(Fingerprint)
}
func (m *ItemBatchMessage) Count() int       { _ = "STUB: not implemented"; return 0 }
func (m *ItemBatchMessage) Keys() []KeyBytes { _ = "STUB: not implemented"; return nil }

func (m *ItemBatchMessage) Since() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (m *ItemBatchMessage) Sample() []MinhashSampleItem {
	_ = "STUB: not implemented"

	// ProbeMessage requests bounded range fingerprint and count from the peer,
	// along with a minhash sample if fingerprints differ.
	return nil
}

type ProbeMessage struct {
	RangeX, RangeY   CompactHash
	RangeFingerprint Fingerprint
	SampleSize       uint32
}

var _ SyncMessage = &ProbeMessage{}

func (m *ProbeMessage) Type() MessageType { _ = "STUB: not implemented"; return *new(MessageType) }
func (m *ProbeMessage) X() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *ProbeMessage) Y() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *ProbeMessage) Fingerprint() Fingerprint {
	_ = "STUB: not implemented"
	return *new(Fingerprint)
}
func (m *ProbeMessage) Count() int       { _ = "STUB: not implemented"; return 0 }
func (m *ProbeMessage) Keys() []KeyBytes { _ = "STUB: not implemented"; return nil }
func (m *ProbeMessage) Since() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (m *ProbeMessage) Sample() []MinhashSampleItem {
	_ = "STUB: not implemented"

	// SampleMessage is a sample of set items.
	return nil
}

type SampleMessage struct {
	RangeX, RangeY   CompactHash
	RangeFingerprint Fingerprint
	NumItems         uint32
	// NOTE: max must be in sync with maxSampleSize in hashsync/rangesync.go
	SampleItems []MinhashSampleItem `scale:"max=1000"`
}

var _ SyncMessage = &SampleMessage{}

func (m *SampleMessage) Type() MessageType { _ = "STUB: not implemented"; return *new(MessageType) }
func (m *SampleMessage) X() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *SampleMessage) Y() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *SampleMessage) Fingerprint() Fingerprint {
	_ = "STUB: not implemented"
	return *new(Fingerprint)
}
func (m *SampleMessage) Count() int                  { _ = "STUB: not implemented"; return 0 }
func (m *SampleMessage) Keys() []KeyBytes            { _ = "STUB: not implemented"; return nil }
func (m *SampleMessage) Since() time.Time            { _ = "STUB: not implemented"; return *new(time.Time) }
func (m *SampleMessage) Sample() []MinhashSampleItem { _ = "STUB: not implemented"; return nil }

// RecentMessage is a SyncMessage that denotes a set of items that have been
// added to the peer's set since the specific point in time.
type RecentMessage struct {
	SinceTime uint64
}

var _ SyncMessage = &RecentMessage{}

func (m *RecentMessage) Type() MessageType { _ = "STUB: not implemented"; return *new(MessageType) }
func (m *RecentMessage) X() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *RecentMessage) Y() KeyBytes       { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (m *RecentMessage) Fingerprint() Fingerprint {
	_ = "STUB: not implemented"
	return *new(Fingerprint)
}
func (m *RecentMessage) Count() int       { _ = "STUB: not implemented"; return 0 }
func (m *RecentMessage) Keys() []KeyBytes { _ = "STUB: not implemented"; return nil }
func (m *RecentMessage) Since() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (m *RecentMessage) Sample() []MinhashSampleItem { _ = "STUB: not implemented"; return nil }
