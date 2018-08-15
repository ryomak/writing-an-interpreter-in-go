package ast

import "github.com/ryomak/writing-an-interpreter-in-go/monkey/token"

type ReturnStatement struct {
	Token        token.Token
	ReturenValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
