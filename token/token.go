package token

// token type
const (
	NAME    = "name"
	PLUS    = "+"
	MINUS   = "-"
	MULTI   = "*"
	LPARA   = "("
	RPARA   = ")"
	EOF     = "eof"
	ILLEGAL = "illegal"
)

type Token struct {
	Type    string
	Literal string
}
