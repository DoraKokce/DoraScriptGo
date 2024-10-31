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
	ExplicitType  Type
}

type ObjectProperty struct {
	IsStatic bool
	Type     Type
}
type StructDeclStmt struct {
	StructName string
	Properties map[string]ObjectProperty
}

func (n BlockStmt) stmt()      {}
func (n ExpressionStmt) stmt() {}
func (n VarDeclStmt) stmt()    {}
func (n StructDeclStmt) stmt() {}
