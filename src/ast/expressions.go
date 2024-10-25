package ast

import "github.com/Dorakokce/DoraScriptGo/src/lexer"

type NumberExpr struct {
	Value float64
}

type StringExpr struct {
	Value string
}

type SymbolExpr struct {
	Value string
}

func (n NumberExpr) expr() {}
func (n StringExpr) expr() {}
func (n SymbolExpr) expr() {}

// Complex

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

type PrefixExpr struct {
	Operator  lexer.Token
	RightExpr Expr
}

type AssignmentExpr struct {
	Assigne  Expr
	Operator lexer.Token
	RHSValue Expr
}

func (n BinaryExpr) expr()     {}
func (n PrefixExpr) expr()     {}
func (n AssignmentExpr) expr() {}
