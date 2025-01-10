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
	// move over `(`
	parser.NextToken()

	// parse child expression with Lowest precedence
	exp, err := parser.ParseExpression(precedence.Lowest)
	if err != nil {
		return nil, err
	}

	// must end with `)`
	err = parser.ExpectNextToken(token.RPARA)
	if err != nil {
		return nil, err
	}

	parentExp := &ParenthesizedExpression{
		Child: exp,
	}

	return parentExp, nil
}
