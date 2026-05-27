package types

// RandomBytes generates random data in bytes for testing.
func RandomBytes(size int) []byte { _ = "STUB: not implemented"; return nil }

// RandomHash generates random Hash32 for testing.
func RandomHash() Hash32 { _ = "STUB: not implemented"; return *new(Hash32) }

// RandomBeacon generates random beacon in bytes for testing.
func RandomBeacon() Beacon { _ = "STUB: not implemented"; return *new(Beacon) }

// RandomActiveSet generates a random set of ATXIDs of the specified size.
func RandomActiveSet(size int) []ATXID { _ = "STUB: not implemented"; return nil }

// RandomTXSet generates a random set of TransactionID of the specified size.
func RandomTXSet(size int) []TransactionID { _ = "STUB: not implemented"; return nil }

// RandomATXID generates a random ATXID for testing.
func RandomATXID() ATXID { _ = "STUB: not implemented"; return *new(ATXID) }

// RandomNodeID generates a random NodeID for testing.
func RandomNodeID() NodeID { _ = "STUB: not implemented"; return *new(NodeID) }

// RandomBallotID generates a random BallotID for testing.
func RandomBallotID() BallotID { _ = "STUB: not implemented"; return *new(BallotID) }

// RandomProposalID generates a random ProposalID for testing.
func RandomProposalID() ProposalID { _ = "STUB: not implemented"; return *new(ProposalID) }

// RandomBlockID generates a random ProposalID for testing.
func RandomBlockID() BlockID { _ = "STUB: not implemented"; return *new(BlockID) }

// RandomTransactionID generates a random TransactionID for testing.
func RandomTransactionID() TransactionID { _ = "STUB: not implemented"; return *new(TransactionID) }

// RandomBallot generates a Ballot with random content for testing.
func RandomBallot() *Ballot { _ = "STUB: not implemented"; return nil }

// RandomEdSignature generates a random (not necessarily valid) EdSignature for testing.
func RandomEdSignature() EdSignature { _ = "STUB: not implemented"; return *new(EdSignature) }

// RandomVrfSignature generates a random VrfSignature for testing.
func RandomVrfSignature() VrfSignature { _ = "STUB: not implemented"; return *new(VrfSignature) }
