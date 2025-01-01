package parser

import (
	"fmt"
	"pratt-go/lexer"
	"pratt-go/token"
)

type Parser struct {
	lexer   *lexer.Lexer
	curTok  *token.Token
	nextTok *token.Token

	unaryParseletMap  map[string]UnaryParselet
	binaryParseletMap map[string]BinaryParselet

	precedence map[string]int
}

func NewParser(input string) *Parser {
	p := &Parser{
		lexer:             lexer.NewLexer(input),
		unaryParseletMap:  make(map[string]UnaryParselet),
		binaryParseletMap: make(map[string]BinaryParselet),
		precedence:        make(map[string]int),
	}

	p.unaryParseletMap[token.NAME] = UnaryOperatorParselet
	p.unaryParseletMap[token.PLUS] = UnaryOperatorParselet
	p.unaryParseletMap[token.MINUS] = UnaryOperatorParselet

	p.binaryParseletMap[token.PLUS] = BinaryOperatorParselet
	p.binaryParseletMap[token.MINUS] = BinaryOperatorParselet
	p.binaryParseletMap[token.MULTI] = BinaryOperatorParselet

	p.precedence[token.PLUS] = 3
	p.precedence[token.MULTI] = 4

	p.ReadToken()
	p.ReadToken()
	return p
}

func (p *Parser) getPrecedence(tokenType string) int {
	num, ok := p.precedence[tokenType]
	if !ok {
		return 0
	}
	return num
}

func (p *Parser) ReadToken() {
	p.curTok = p.nextTok
	p.nextTok = p.lexer.NextToken()
}

// from left to right!
func (p *Parser) ParseExpression(precedence int) (Expression, error) {
	t := p.curTok
	fmt.Printf("tok: %s, precedence: %d\n", t.Literal, precedence)

	// handle unary
	unaryParselet, ok := p.unaryParseletMap[t.Type]
	if !ok {
		return nil, fmt.Errorf("Error parse token %s", t.Type)
	}

	exp, err := unaryParselet(p, t)
	if err != nil {
		return nil, err
	}

	// handle binary
	left := exp
	t = p.nextTok
	binaryParselet, ok := p.binaryParseletMap[t.Type]
	if !ok {
		// next token is not a binary operator, return directly
		return exp, nil
	}

	for precedence < p.getPrecedence(p.nextTok.Type) {
		p.ReadToken()
		exp, err = binaryParselet(p, t, left)
		if err != nil {
			return nil, err
		}
		left = exp
	}

	return exp, nil
}
