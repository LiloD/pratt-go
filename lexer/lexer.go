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

	l.readCh() // read and init
	return l
}

func (l *Lexer) NextToken() *token.Token {
	l.whitespace()
	var tok *token.Token
	switch l.ch {
	case '+':
		tok = &token.Token{Type: token.PLUS, Literal: "+"}
	case '-':
		tok = &token.Token{Type: token.MINUS, Literal: "-"}
	case '*':
		tok = &token.Token{Type: token.MULTI, Literal: "*"}
	case '(':
		tok = &token.Token{Type: token.LPARA, Literal: "("}
	case ')':
		tok = &token.Token{Type: token.RPARA, Literal: ")"}
	case 0:
		tok = &token.Token{Type: token.EOF, Literal: ""}
	default:
		if isDigit(l.ch) || isChar(l.ch) {
			tok = &token.Token{Type: token.NAME, Literal: l.readName()}
			return tok
		}

		tok = &token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
	}

	l.readCh()
	return tok
}

func (l *Lexer) readCh() {
	l.pos += 1
	if l.pos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.pos]
	}
}

// consume whitespaces
func (l *Lexer) whitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readCh()
	}
}

func (l *Lexer) readName() string {
	start := l.pos
	for isDigit(l.ch) || isChar(l.ch) {
		l.readCh()
	}
	end := l.pos
	return l.input[start:end]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
