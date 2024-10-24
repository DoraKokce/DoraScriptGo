package main

import (
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.ds")
	source := string(bytes)
	lexer := createLexer()

	tokens := lexer.Tokenize(source)
}
