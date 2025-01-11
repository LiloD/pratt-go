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
	parser.NextToken()
	precedence := parser.getPrecedence(token.Type)
	exp, err := parser.ParseExpression(precedence)
	if err != nil {
		return nil, err
	}

	binaryExp := &BinaryExpression{
		Tok:   token,
		Left:  left,
		Right: exp,
	}

	return binaryExp, nil
}

func BinaryRightOperatorParselet(parser *Parser, token *token.Token, left Expression) (Expression, error) {
	parser.NextToken()
	precedence := parser.getPrecedence(token.Type)
	exp, err := parser.ParseExpression(precedence - 1)
	if err != nil {
		return nil, err
	}

	binaryExp := &BinaryExpression{
		Tok:   token,
		Left:  left,
		Right: exp,
	}

	return binaryExp, nil
}
