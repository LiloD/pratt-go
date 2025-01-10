package parser

import (
	"fmt"
	"pratt-go/precedence"
	"pratt-go/token"
)

type CallExpression struct {
	Function  Expression
	Arguments []Expression
}

func (c *CallExpression) expression() {}

func (c *CallExpression) String() string {
	str := fmt.Sprintf("%s(", c.Function.String())
	for i := 0; i < len(c.Arguments); i++ {
		str += c.Arguments[i].String()
		if i < len(c.Arguments)-1 {
			str += ","
		}
	}
	str += ")"
	return str
}

func CallExpressionParslet(parser *Parser, tok *token.Token, left Expression) (Expression, error) {
	// left is the function
	// move over `(`
	parser.NextToken()

	// now try to parse arguments
	var arguments []Expression

	if parser.curTok.Type != token.RPARA {
		for {
			exp, err := parser.ParseExpression(precedence.Lowest)
			if err != nil {
				return nil, err
			}
			arguments = append(arguments, exp)

			if err := parser.ExpectNextToken(token.COMMA); err != nil {
				break
			}

			// move over `,`
			parser.NextToken()
		}

		// MUST end with `)`
		err := parser.ExpectNextToken(token.RPARA)

		if err != nil {
			return nil, err
		}
	}

	callExp := &CallExpression{
		Function:  left,
		Arguments: arguments,
	}

	return callExp, nil
}
