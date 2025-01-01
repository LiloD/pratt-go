package parser

import (
	"pratt-go/token"
)

type Expression interface {
	String() string
	expression() // dummy
}

type UnaryParselet func(parser *Parser, token *token.Token) (Expression, error)

type BinaryParselet func(parser *Parser, token *token.Token, left Expression) (Expression, error)
