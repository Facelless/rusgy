package parser

import (
	"src/src/ast"
	"src/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]
	if exists {
		return stmt_fn(p)
	}
	expression := parse_expor(p, defalt_bp)
	p.expect(lexer.SEMI_COLON)
	return ast.ExpressionStmt{
		Expression: expression,
	}
}
