package ast

type BlockStmt struct {
	Body []Stmt
}

type ExpressionStmt struct {
	Expression Expr
}

func (n BlockStmt) stmt()      {}
func (n ExpressionStmt) stmt() {}
