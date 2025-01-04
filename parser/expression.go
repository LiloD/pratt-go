package parser

import (
	"pratt-go/token"
)

// define the precedence
const (
	_ = iota
	Sum
	Product
	Prefix
)

type Expression interface {
	String() string
	expression() // dummy
}

type PrefixParselet func(parser *Parser, token *token.Token) (Expression, error)

type InfixParselet func(parser *Parser, token *token.Token, left Expression) (Expression, error)
