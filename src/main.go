package main

import (
	"os"

	"github.com/Dorakokce/DoraScriptGo/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.ds")
	tokens := lexer.Tokenize(string(bytes))
	for _, tok := range tokens {
		tok.Debug()
	}
}
