package ast

type SymbolType struct {
	Name string
}

type ArrayType struct {
	Underlying Type
}

func (t SymbolType) _type() {}
func (t ArrayType) _type()  {}
