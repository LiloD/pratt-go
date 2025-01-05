package parser

import (
	"fmt"
	"pratt-go/precedence"
	"pratt-go/token"
)

type UnaryExpression struct {
	Tok     *token.Token
	Operand Expression
}

func (u *UnaryExpression) expression() {}

func (u *UnaryExpression) String() string {
	return fmt.Sprintf(
		"(%s%s)",
		u.Tok.Literal,
		u.Operand.String(),
	)
}

func UnaryOperatorParselet(parser *Parser, tok *token.Token) (Expression, error) {
	parser.ReadToken()
	exp, err := parser.ParseExpression(precedence.Prefix)
	if err != nil {
		return nil, fmt.Errorf("Error parse operand of unary operator %s: %v", tok.Literal, err)
	}
	return &UnaryExpression{
		Tok:     tok,
		Operand: exp,
	}, nil
}
