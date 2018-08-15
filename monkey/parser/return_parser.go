package parser

import (
	"github.com/ryomak/writing-an-interpreter-in-go/monkey/ast"
	"github.com/ryomak/writing-an-interpreter-in-go/monkey/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
