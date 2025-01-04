package parser

import (
	"fmt"
	"pratt-go/token"
)

// name
type NumberExpression struct {
	Number string
}

// name should implement Expression
func (n *NumberExpression) String() string {
	return fmt.Sprintf("(number_literal: %s)", n.Number)
}

func (n *NumberExpression) expression() {}

// NumberParselet is a prefix parselet
func NumberParselet(parser *Parser, tok *token.Token) (Expression, error) {
	return &NumberExpression{
		Number: tok.Literal,
	}, nil
}
