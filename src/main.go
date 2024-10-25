package main

import (
	"os"

	"github.com/Dorakokce/DoraScriptGo/src/lexer"
	"github.com/Dorakokce/DoraScriptGo/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/04.ds")
	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
