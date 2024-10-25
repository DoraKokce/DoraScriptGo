package main

import (
	"os"       
    "https://github.com/Dorakokce/DoraScriptGo/src/lexer/"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.ds")
	source := string(bytes)

	tokens := lexer.Tokenize(source)
}
