package main

import (
	"os"
	"src/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("../archives/archive00.lg")
	tokens := lexer.Tokenize(string(bytes))
	for _, token := range tokens {
		token.Debug()
	}
}
