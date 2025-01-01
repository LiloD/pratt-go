package parser

import (
	"fmt"
	"pratt-go/token"
)

type UnaryExpression struct {
	tok     *token.Token
	operand Expression
}

func (u *UnaryExpression) expression() {}

func (u *UnaryExpression) String() string {
	return fmt.Sprintf(
		"(unary_expression %s: (operand: %s))",
		u.tok.Literal,
		u.operand.String(),
	)
}

func UnaryOperatorParselet(parser *Parser, tok *token.Token) (Expression, error) {
	// name identifer is speciall, no operand needed
	if tok.Type == token.NAME {
		return &NameExpression{Name: tok.Literal}, nil
	}

	parser.ReadToken()
	exp, err := parser.ParseExpression(5)
	if err != nil {
		return nil, fmt.Errorf("Error parse operand of unary operator %s: %v", tok.Literal, err)
	}
	return &UnaryExpression{
		tok:     tok,
		operand: exp,
	}, nil
}
