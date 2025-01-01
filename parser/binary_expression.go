package parser

import (
	"fmt"
	"pratt-go/token"
)

type BinaryExpression struct {
	tok   *token.Token
	left  Expression
	right Expression
}

func (b *BinaryExpression) expression() {}

func (b *BinaryExpression) String() string {
	return fmt.Sprintf(
		"(binary_expression %s: (left: %s, right: %s))",
		b.tok.Literal,
		b.left.String(),
		b.right.String(),
	)
}

func BinaryOperatorParselet(parser *Parser, token *token.Token, left Expression) (Expression, error) {
	parser.ReadToken()
	precedence := parser.getPrecedence(token.Type)
	exp, err := parser.ParseExpression(precedence) // plus 3
	if err != nil {
		return nil, fmt.Errorf("Error parse right operand of biarny operator %s: %v", token.Literal, err)
	}

	return &BinaryExpression{
		tok:   token,
		left:  left,
		right: exp,
	}, nil
}
