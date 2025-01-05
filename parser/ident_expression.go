package parser

import (
	"pratt-go/token"
)

// name
type IdentifierExpression struct {
	Name string
}

// name should implement Expression
func (i *IdentifierExpression) String() string {
	return i.Name
}

func (i *IdentifierExpression) expression() {}

// IdentifierParselet is a prefix parselet
func IdentifierParselet(parser *Parser, tok *token.Token) (Expression, error) {
	return &IdentifierExpression{
		Name: tok.Literal,
	}, nil
}
