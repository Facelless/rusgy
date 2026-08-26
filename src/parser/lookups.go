package parser

import "src/src/ast"

type binding_power int

const (
	defalt_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *parser) ast.Stmt
type nud_handler func(p *parser) ast.Expr
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr
