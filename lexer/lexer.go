package lexer

import (
	"pratt-go/token"
)

type Lexer struct {
	input string
	ch    byte
	pos   int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
		pos:   -1,
	}

	l.nextCh() //fill the initial ch and pos
	return l
}

func (l *Lexer) NextToken() *token.Token {
	l.whitespace()
	var tok *token.Token
	switch l.ch {
	case '!':
		tok = &token.Token{Type: token.BANG, Literal: "!"}
	case '+':
		tok = &token.Token{Type: token.PLUS, Literal: "+"}
	case '-':
		tok = &token.Token{Type: token.MINUS, Literal: "-"}
	case '*':
		tok = &token.Token{Type: token.ASTERISK, Literal: "*"}
	case '(':
		tok = &token.Token{Type: token.LPARA, Literal: "("}
	case ')':
		tok = &token.Token{Type: token.RPARA, Literal: ")"}
	case 0:
		tok = &token.Token{Type: token.EOF, Literal: ""}
	default:
		if isChar(l.ch) {
			tok = &token.Token{Type: token.IDENT, Literal: l.readName()}
			return tok
		}

		if isDigit(l.ch) {
			tok = &token.Token{Type: token.NUMBER, Literal: l.readNumber()}
			return tok
		}

		tok = &token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
	}

	l.nextCh()
	return tok
}

// move the pointer and read next character
func (l *Lexer) nextCh() {
	// move the position pointer
	l.pos += 1
	// check if it's the end of the input
	if l.pos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.pos]
	}
}

func (l *Lexer) readName() string {
	start := l.pos
	for isChar(l.ch) {
		l.nextCh()
	}
	end := l.pos
	return l.input[start:end]
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for isDigit(l.ch) {
		l.nextCh()
	}
	end := l.pos
	return l.input[start:end]
}

// consume whitespaces
func (l *Lexer) whitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.nextCh()
	}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
