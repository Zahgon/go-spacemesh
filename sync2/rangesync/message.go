package rangesync

import (
	"time"
)

// MessageType specifies the type of a sync message.
type MessageType byte

const (
	// Done message is sent to indicate the completion of the whole sync run.
	MessageTypeDone MessageType = iota
	// EndRoundMessage is sent to indicate the completion of a single round.
	MessageTypeEndRound
	// EmptySetMessage is sent to indicate that the set is empty. In response, the
	// receiving side sends its whole set using ItemBatch and RangeContents messages.
	MessageTypeEmptySet
	// EmptyRangeMessage is sent to indicate that the specified range is empty. In
	// response, the receiving side sends the contents of the range using ItemBatch
	// and RangeContents messages.
	MessageTypeEmptyRange
	// Fingerprint carries a range fingerprint. Depending on the local fingerprint of
	// the same range, the receiving side may not reply to this message (range
	// completed), may send the contents of the range using ItemBatch and
	// RangeContents messages if the range is small enough, or it can split the range
	// in two and send back the Fingerprint messages for each of resulting parts.
	MessageTypeFingerprint
	// RangeContents message is sent after ItemBatch messages to indicate what range
	// the messages belong too. If the receiving side has any items within the same
	// range, these items are sent back using ItemBatch and RangeContents messages.
	MessageTypeRangeContents
	// ItemBatchMessage is sent to carry a batch of items.
	MessageTypeItemBatch
	// ProbeMessage is sent to request the count of items and an estimate of the
	// Jaccard similarity coefficient for the specfied range.
	MessageTypeProbe
	// Sample message carries a minhash sample along with the fingerprint and
	// the fingerprint and number of items within the specied range.
	MessageTypeSample
	// Recent message is sent after ItemBatch messages to indicate that the batches
	// contains items recently added to the set.
	MessageTypeRecent
)

var messageTypes = []string{
	"done",
	"endRound",
	"emptySet",
	"emptyRange",
	"fingerprint",
	"rangeContents",
	"itemBatch",
	"probe",
	"sample",
	"recent",
}

// String implements Stringer.
func (mtype MessageType) String() string { _ = "STUB: not implemented"; return "" }

// SyncMessageToString returns string representation of a sync message.
func SyncMessageToString(m SyncMessage) string { _ = "STUB: not implemented"; return "" }

type sender struct {
	Conduit
}

func (s sender) SendFingerprint(x, y KeyBytes, fp Fingerprint, count int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s sender) SendEmptySet() error { _ = "STUB: not implemented"; return nil }

func (s sender) SendEmptyRange(x, y KeyBytes) error { _ = "STUB: not implemented"; return nil }

func (s sender) SendRangeContents(x, y KeyBytes, count int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s sender) SendChunk(items []KeyBytes) error { _ = "STUB: not implemented"; return nil }

func (s sender) SendEndRound() error { _ = "STUB: not implemented"; return nil }

func (s sender) SendDone() error { _ = "STUB: not implemented"; return nil }

func (s sender) SendProbe(x, y KeyBytes, fp Fingerprint, sampleSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s sender) SendSample(
	x, y KeyBytes,
	fp Fingerprint,
	count, sampleSize int,
	sr SeqResult,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s sender) SendRecent(since time.Time) error { _ = "STUB: not implemented"; return nil }

// "Empty" message types follow. These do not need scalegen and thus are not in wire_types.go.

type Marker struct{}

func (*Marker) X() KeyBytes              { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (*Marker) Y() KeyBytes              { _ = "STUB: not implemented"; return *new(KeyBytes) }
func (*Marker) Fingerprint() Fingerprint { _ = "STUB: not implemented"; return *new(Fingerprint) }
func (*Marker) Count() int               { _ = "STUB: not implemented"; return 0 }
func (*Marker) Keys() []KeyBytes         { _ = "STUB: not implemented"; return nil }
func (*Marker) Since() time.Time         { _ = "STUB: not implemented"; return *new(time.Time) }
func (*Marker) Sample() []MinhashSampleItem {
	_ = "STUB: not implemented"

	// DoneMessage is a SyncMessage that denotes the end of the synchronization.
	// The peer should stop any further processing after receiving this message.
	return nil
}

type DoneMessage struct{ Marker }

var _ SyncMessage = &DoneMessage{}

func (*DoneMessage) Type() MessageType {
	_ = "STUB: not implemented"
	return *

	// EndRoundMessage is a SyncMessage that denotes the end of the sync round.
	new(MessageType)
}

type EndRoundMessage struct{ Marker }

var _ SyncMessage = &EndRoundMessage{}

func (*EndRoundMessage) Type() MessageType {
	_ = "STUB: not implemented"
	return *

	// EmptySetMessage is a SyncMessage that denotes an empty set, requesting the
	// peer to send all of its items.
	new(MessageType)
}

type EmptySetMessage struct{ Marker }

var _ SyncMessage = &EmptySetMessage{}

func (*EmptySetMessage) Type() MessageType { _ = "STUB: not implemented"; return *new(MessageType) }
