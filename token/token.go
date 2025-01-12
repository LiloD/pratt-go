package token

// token type
const (
	IDENT    = "identifer"
	NUMBER   = "number"
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	CARET    = "^"
	SLASH    = "/"
	BANG     = "!"
	QUESTION = "?"
	COMMA    = ","
	COLON    = ":"
	LPARA    = "("
	RPARA    = ")"
	EOF      = "eof"
	ILLEGAL  = "illegal"
)

type Token struct {
	Type    string
	Literal string
}
