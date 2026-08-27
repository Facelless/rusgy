package main

import (
	"os"
	"src/src/lexer"
	"src/src/parser"

	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("../archives/archive00.lg")
	tokens := lexer.Tokenize(string(bytes))
	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
