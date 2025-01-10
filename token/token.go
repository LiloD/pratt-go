package token

// token type
const (
	IDENT    = "identifer"
	NUMBER   = "number"
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	BANG     = "!"
	COMMA    = ","
	LPARA    = "("
	RPARA    = ")"
	EOF      = "eof"
	ILLEGAL  = "illegal"
)

type Token struct {
	Type    string
	Literal string
}
