package parser

import (
	"fmt"
	"src/src/ast"
	"src/src/lexer"
	"strconv"
)

func parser_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Cannot Create PRIMARY_EXPRESSION  from %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parse_expor(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]
	if !exists {
		panic(fmt.Sprintf("NUD HANDLER	EXPECTED FOR TOKEN  %s\n", lexer.TokenKindString(tokenKind)))
	}
	left := nud_fn(p)
	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("LED HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}
	return left
}

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	right := parse_expor(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}
