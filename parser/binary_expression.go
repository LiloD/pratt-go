package parser

import (
	"fmt"
	"pratt-go/token"
)

type BinaryExpression struct {
	Tok   *token.Token
	Left  Expression
	Right Expression
}

func (b *BinaryExpression) expression() {}

func (b *BinaryExpression) String() string {
	return fmt.Sprintf(
		"(%s%s%s)",
		b.Left.String(),
		b.Tok.Literal,
		b.Right.String(),
	)
}

func BinaryOperatorParselet(parser *Parser, token *token.Token, left Expression) (Expression, error) {
	parser.ReadToken()
	precedence := parser.getPrecedence(token.Type)
	exp, err := parser.ParseExpression(precedence)
	if err != nil {
		return nil, fmt.Errorf("Error parse right operand of biarny operator %s: %v", token.Literal, err)
	}

	binaryExp := &BinaryExpression{
		Tok:   token,
		Left:  left,
		Right: exp,
	}

	return binaryExp, nil
}
