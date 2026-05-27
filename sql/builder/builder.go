package builder

import (
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/sql"
)

type token string

const (
	Eq        token = "="
	NotEq     token = "!="
	Gt        token = ">"
	Gte       token = ">="
	Lt        token = "<"
	Lte       token = "<="
	In        token = "in"
	IsNotNull token = "is not null"
)

type operator string

const (
	And operator = "and"
	Or  operator = "or"
)

type field string

const (
	Epoch     field = "epoch"
	Smesher   field = "pubkey"
	Coinbase  field = "coinbase"
	Id        field = "id"
	Layer     field = "layer"
	Address   field = "address"
	Principal field = "principal"
	Proof     field = "proof"
)

type modifier string

const (
	Offset  modifier = "offset"
	Limit   modifier = "limit"
	OrderBy modifier = "order by"
	GroupBy modifier = "group by"
)

type Op struct {
	// Prefix will be added before field name
	Prefix string
	Field  field
	Token  token
	// Value will be type casted to one the expected types.
	// Operation will panic if it doesn't match any of expected.
	Value any

	Group         []Op
	GroupOperator operator

	// CustomQuery is used to add custom query. If this is set, Field and Token will be ignored.
	// This is useful for complex queries that can't be expressed with Field and Token.
	// Value will be used for custom query if it's not nil.
	// Remember about setting correct bind index for Value.
	CustomQuery string
}

type Modifier struct {
	Key modifier
	// Value will be type casted to one the expected types.
	// Modifier will panic if it doesn't match any of expected.
	Value any
}

type Operations struct {
	Filter    []Op
	Modifiers []Modifier
	StartWith string
}

func FilterEpochOnly(publish types.EpochID) Operations {
	_ = "STUB: not implemented"
	return *new(Operations)
}

func FilterFrom(operations Operations) string { _ = "STUB: not implemented"; return "" }

func BindingsFrom(operations Operations) sql.Encoder {
	_ = "STUB: not implemented"
	return *new(sql.Encoder)
}

func bindValue(stmt *sql.Statement, bindIndex int, value any) int {
	_ = "STUB: not implemented"
	return 0
}

// do nothing
