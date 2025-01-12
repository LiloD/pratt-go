package parser

import (
	"fmt"
	"pratt-go/precedence"
	"pratt-go/token"
)

type ConditionalExpression struct {
	Tok         *token.Token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func (c *ConditionalExpression) expression() {}

func (c *ConditionalExpression) String() string {
	return fmt.Sprintf(
		"(%s?%s:%s)",
		c.Condition.String(),
		c.Consequence.String(),
		c.Alternative.String(),
	)
}

func ConditionalOperatorParselet(parser *Parser, tok *token.Token, left Expression) (Expression, error) {
	parser.NextToken()
	// parse consequence
	consequence, err := parser.ParseExpression(precedence.Lowest)
	if err != nil {
		return nil, err
	}

	// expect : and move over it
	parser.ExpectNextToken(token.COLON)
	parser.NextToken()

	// parse alternative
	alternative, err := parser.ParseExpression(precedence.Lowest)
	if err != nil {
		return nil, err
	}

	return &ConditionalExpression{
		Condition:   left,
		Consequence: consequence,
		Alternative: alternative,
	}, nil
}
