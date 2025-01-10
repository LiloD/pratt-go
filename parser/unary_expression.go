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
	parser.NextToken()
	exp, err := parser.ParseExpression(precedence.Prefix)
	if err != nil {
		return nil, err
	}
	return &UnaryExpression{
		Tok:     tok,
		Operand: exp,
	}, nil
}
