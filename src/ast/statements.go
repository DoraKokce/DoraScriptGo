package ast

type BlockStmt struct {
	Body []Stmt
}

type ExpressionStmt struct {
	Expression Expr
}

type VarDeclStmt struct {
	VarName       string
	IsConstant    bool
	AssignedValue Expr
}

func (n BlockStmt) stmt()      {}
func (n ExpressionStmt) stmt() {}
func (n VarDeclStmt) stmt()    {}
