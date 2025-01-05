package parser

import (
	"fmt"
	"pratt-go/lexer"
	"pratt-go/precedence"
	"pratt-go/token"
)

type Parser struct {
	lexer   *lexer.Lexer
	curTok  *token.Token
	nextTok *token.Token

	prefixParseletMap     map[string]PrefixParselet
	infixParseletMap      map[string]InfixParselet
	operatorPrecedenceMap map[string]int
}

func NewParser(input string) *Parser {
	p := &Parser{
		lexer:                 lexer.NewLexer(input),
		prefixParseletMap:     make(map[string]PrefixParselet),
		infixParseletMap:      make(map[string]InfixParselet),
		operatorPrecedenceMap: make(map[string]int),
	}

	p.prefixParseletMap[token.IDENT] = IdentifierParselet
	p.prefixParseletMap[token.NUMBER] = NumberParselet

	p.prefixParseletMap[token.PLUS] = UnaryOperatorParselet
	p.prefixParseletMap[token.MINUS] = UnaryOperatorParselet
	p.prefixParseletMap[token.BANG] = UnaryOperatorParselet

	p.infixParseletMap[token.PLUS] = BinaryOperatorParselet
	p.infixParseletMap[token.MINUS] = BinaryOperatorParselet
	p.infixParseletMap[token.ASTERISK] = BinaryOperatorParselet

	p.operatorPrecedenceMap[token.PLUS] = precedence.Sum
	p.operatorPrecedenceMap[token.MINUS] = precedence.Sum
	p.operatorPrecedenceMap[token.ASTERISK] = precedence.Product

	p.ReadToken()
	p.ReadToken()
	return p
}

func (p *Parser) getPrecedence(tokenType string) int {
	num, ok := p.operatorPrecedenceMap[tokenType]
	if !ok {
		return 0
	}
	return num
}

func (p *Parser) ReadToken() {
	p.curTok = p.nextTok
	p.nextTok = p.lexer.NextToken()
}

func (p *Parser) ParseExpression(precedence int) (Expression, error) {
	// parseExpression parse token from `left` to `right`
	prefixParselet, ok := p.prefixParseletMap[p.curTok.Type]
	// error if we don't know how parse left most token
	if !ok {
		return nil, fmt.Errorf("Error parse token %s", p.curTok.Type)
	}

	left, err := prefixParselet(p, p.curTok)
	if err != nil {
		return nil, err
	}

	// after parse the first left expression
	// peek next token to see if this is a infix expression
	infixParselet, ok := p.infixParseletMap[p.nextTok.Type]
	if !ok {
		// if next token is not a binary operator
		// then current expression end here, return left directly
		return left, nil
	}

	for p.getPrecedence(p.nextTok.Type) > precedence {
		p.ReadToken()
		// with each loop, left get updated
		left, err = infixParselet(p, p.curTok, left)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}
