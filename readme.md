# Pratt in go

Pratt Parser based on [Pratt Parsers: Expression Parsing Made Easy](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/).

It starts with basic unary and binary operators like -, +, *, ! and basic identifier and number lexing and keep expanding the syntax


## Parenthesized Expression

Parenthesized Expression is a `prefix` expression, since it starts with `(`

Normally prefix operator has a higher precedence, but parenthesized expression need to parse following expression until 
meeting the token `)`.

So we need to call parseExpression with lowest precedence.

```go
type ParenthesizedExpression struct {
	Child Expression
}

func ParenthesizedParselet(parser *Parser, tok *token.Token) (Expression, error) {
	parser.ReadToken()
	exp, err := parser.ParseExpression(precedence.Lowest)
	if err != nil {
		return nil, err
	}

	err = parser.ExpectToken(token.RPARA)
	if err != nil {
		return nil, err
	}

	parentExp := &ParenthesizedExpression{
		Child: exp,
	}

	return parentExp, nil
}
```

