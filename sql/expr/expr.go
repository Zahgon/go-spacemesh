// Package expr proviedes a simple SQL expression parser and builder.
// It wraps the rqlite/sql package and provides a more convenient API that contains only
// what's needed for the go-spacemesh codebase.
package expr

import (
	rsql "github.com/rqlite/sql"
)

// SQL operations.
const (
	NE     = rsql.NE     // !=
	EQ     = rsql.EQ     // =
	LE     = rsql.LE     // <=
	LT     = rsql.LT     // <
	GT     = rsql.GT     // >
	GE     = rsql.GE     // >=
	BITAND = rsql.BITAND // &
	BITOR  = rsql.BITOR  // |
	BITNOT = rsql.BITNOT // !
	LSHIFT = rsql.LSHIFT // <<
	RSHIFT = rsql.RSHIFT // >>
	PLUS   = rsql.PLUS   // +
	MINUS  = rsql.MINUS  // -
	STAR   = rsql.STAR   // *
	SLASH  = rsql.SLASH  // /
	REM    = rsql.REM    // %
	CONCAT = rsql.CONCAT // ||
	DOT    = rsql.DOT    // .
	AND    = rsql.AND
	OR     = rsql.OR
	NOT    = rsql.NOT
)

// Expr represents a parsed SQL expression.
type Expr = rsql.Expr

// Statement represents a parsed SQL statement.
type Statement = rsql.Statement

// MustParse parses an SQL expression and panics if there's an error.
func MustParse(s string) rsql.Expr { _ = "STUB: not implemented"; return *new(rsql.Expr) }

// MustParseStatement parses an SQL statement and panics if there's an error.
func MustParseStatement(s string) rsql.Statement {
	_ = "STUB: not implemented"
	return *new(rsql.Statement)
}

// MaybeAnd joins together several SQL expressions with AND, ignoring any nil exprs.
// If no non-nil expressions are passed, nil is returned.
// If a single non-nil expression is passed, that single expression is returned.
// Otherwise, the expressions are joined together with ANDs:
// a AND b AND c AND d.
func MaybeAnd(exprs ...Expr) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Ident constructs SQL identifier expression for the identifier with the specified name.
func Ident(name string) *rsql.Ident { _ = "STUB: not implemented"; return nil }

// Number constructs a number literal.
func Number(value string) *rsql.NumberLit { _ = "STUB: not implemented"; return nil }

// TableSource constructs a Source clause for SELECT statement that corresponds to
// selecting from a single table with the specified name.
func TableSource(name string) rsql.Source { _ = "STUB: not implemented"; return *new(rsql.Source) }

// Op constructs a binary expression such as x + y or x < y.
func Op(x Expr, op rsql.Token, y Expr) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Bind constructs the unnamed bind expression (?).
func Bind() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Between constructs BETWEEN expression: x BETWEEN a AND b.
func Between(x, a, b Expr) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Call constructs a call expression with specified arguments such as max(x).
func Call(name string, args ...Expr) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// CountStar returns a COUNT(*) expression.
func CountStar() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Asc constructs an ascending ORDER BY term.
func Asc(expr Expr) *rsql.OrderingTerm { _ = "STUB: not implemented"; return nil }

// Desc constructs a descedning ORDER BY term.
func Desc(expr Expr) *rsql.OrderingTerm { _ = "STUB: not implemented"; return nil }

// SelectBuilder is used to construct a SELECT statement.
type SelectBuilder struct {
	st *rsql.SelectStatement
}

// Select returns a SELECT statement builder.
func Select(columns ...any) SelectBuilder { _ = "STUB: not implemented"; return *new(SelectBuilder) }

// SelectBasedOn returns a SELECT statement builder based on the specified SELECT statement.
// The statement must be parseable, otherwise SelectBasedOn panics.
// The builder methods can be used to alter the statement.
func SelectBasedOn(st Statement) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Get returns the underlying SELECT statement.
func (sb SelectBuilder) Get() *rsql.SelectStatement {
	_ = "STUB: not implemented"

	// String returns the underlying SELECT statement as a string.
	return nil
}

func (sb SelectBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Columns sets columns in the SELECT statement.
func (sb SelectBuilder) Columns(columns ...any) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// From adds FROM clause to the SELECT statement.
func (sb SelectBuilder) From(s rsql.Source) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// From adds WHERE clause to the SELECT statement.
func (sb SelectBuilder) Where(s Expr) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// From adds ORDER BY clause to the SELECT statement.
func (sb SelectBuilder) OrderBy(terms ...*rsql.OrderingTerm) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// From adds LIMIT clause to the SELECT statement.
func (sb SelectBuilder) Limit(limit Expr) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// ColumnExpr returns nth column expression from the SELECT statement.
func ColumnExpr(st Statement, n int) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// WhereExpr returns WHERE expression from the SELECT statement.
func WhereExpr(st Statement) Expr { _ = "STUB: not implemented"; return *new(Expr) }
