package parser

import (
	"pratt-go/precedence"
	"pratt-go/token"
)

type ParenthesizedExpression struct {
	Child Expression
}

func (p *ParenthesizedExpression) expression() {}

func (p *ParenthesizedExpression) String() string {
	return p.Child.String()
}

func ParenthesizedParselet(parser *Parser, tok *token.Token) (Expression, error) {
	parser.ReadToken()
	exp, err := parser.ParseExpression(precedence.Lowest)
	if err != nil {
		return nil, err
	}

	err = parser.ExpectToken(token.RPARA)
	if err != nil {
		return nil, err
	}

	parentExp := &ParenthesizedExpression{
		Child: exp,
	}

	return parentExp, nil
}
