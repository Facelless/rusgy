package parser

import (
	"src/src/ast"
	"src/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	return &parser{
		tokens: tokens, pos: 0,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	Body := make([]ast.Stmt, 0)
	p := createParser(tokens)
	for p.hasTokens() {
		Body = append(Body, parse_stmt(p))
	}
	return ast.BlockStmt{
		Body: Body,
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]

}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentToken().Kind != lexer.EOF
}
