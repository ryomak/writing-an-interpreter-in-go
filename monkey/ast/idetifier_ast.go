package ast

import "github.com/ryomak/writing-an-interpreter-in-go/monkey/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) statementNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string {
	return i.Value
}
