# Pratt in go

Pratt Parser based on [Pratt Parsers: Expression Parsing Made Easy](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/).

It starts with basic unary and binary operators like -, +, *, ! and basic identifier and number lexing and keep expanding the syntax

At the beginning, it only support basic unary and binary operation, and the precedence is 

```go
const (
	Lowest = iota
	Sum
	Product
	Prefix
)
```

Some basic example are:

`!a + b` - `!` is a prefix operator so it has higher precedence than sum
```
binary_expression(+)
  left: unary_expression(!)
    operand: identifier(a)
  right: identifier(b)
```

`a + b + c` - same precedence lead to left association
```
binary_expression(+)
  left: binary_expression(+)
    left: identifier(a)
    right: identifier(b)
  right: identifier(c)
```

`a + b * c` - product has higher precedence than sum
```
binary_expression(+)
  left: identifier(a)
  right: binary_expression(*)
    left: identifier(b)
    right: identifier(c)
```

## Parenthesized Expression
Parenthesized Expression is a `prefix` expression, since `it starts with the (`, so it has a higher precedence than sum 
and product operation

Parenthesized expression need to parse the inner expression `until meeting the )`. so we need to call parseExpression 
with lowest precedence in case it returned early.


Expression Struct and Parselet:
```go
type ParenthesizedExpression struct {
	Child Expression
}

func ParenthesizedParselet(parser *Parser, tok *token.Token) (Expression, error) {
	// move over `(`
	parser.NextToken()

	// parse child expression with Lowest precedence
	exp, err := parser.ParseExpression(precedence.Lowest)
	if err != nil {
		return nil, err
	}

	// must end with `)`
	err = parser.ExpectNextToken(token.RPARA)
	if err != nil {
		return nil, err
	}

	parentExp := &ParenthesizedExpression{
		Child: exp,
	}

	return parentExp, nil
}
```

Some example:

`!(a+b)`
```
unary_expression(!)
  operand: parenthesized_expression
    binary_expression(+)
      left: identifier(a)
      right: identifier(b)
```

`(-a+b)`
```
parenthesized_expression
  binary_expression(+)
    left: unary_expression(-)
      operand: identifier(a)
    right: identifier(b)
```

`a+(b+c)+d`
```
binary_expression(+)
  left: binary_expression(+)
    left: identifier(a)
    right: parenthesized_expression
      binary_expression(+)
        left: identifier(b)
        right: identifier(c)
  right: identifier(d)
```

## Call Expression
Call Expression is interesting, since it use `(` as well.

But this time `(` left parenthesis is a `infix operator`, since Call Expression starts with a `Function`(which is left
Expression), and the `(`, and the `Arguments` list(which is a list of Expression), and it must `ends with )`

But what about precedence? Well, of cause, it must be higher than Product and Sum, but consider expression `!a()`, we 
need to parse `a()` first and then the outer `!` unary prefix operator

Based on that, we need call expression to have higher precedence than any prefix expression

Updated Precedence List:
```go
const (
	Lowest = iota
	Sum
	Product
	Prefix
	Call    // <- here!
)
```

Expression Struct and Parselet:
```go
type CallExpression struct {
	Function  Expression
	Arguments []Expression
}

func CallExpressionParslet(parser *Parser, tok *token.Token, left Expression) (Expression, error) {
	// left is the function
	// move over `(`
	parser.ReadToken()

	// now try to parse arguments
	var arguments []Expression

	if parser.curTok.Type != token.RPARA {
		for {
			exp, err := parser.ParseExpression(precedence.Lowest)
			if err != nil {
				return nil, err
			}
			arguments = append(arguments, exp)

			if err := parser.ExpectToken(token.COMMA); err != nil {
				break
			}

			// move over `,`
			parser.ReadToken()
		}

		// MUST end with `)`
		err := parser.ExpectToken(token.RPARA)

		if err != nil {
			return nil, err
		}
	}

	callExp := &CallExpression{
		Function:  left,
		Arguments: arguments,
	}

	return callExp, nil
}
```
Make sure to handle 2 special cases:
1. Call Expression without a argument, like `foo()` 
2. Call Expression with only one argument, like `foo(bar)`, since it does not need to handle `,`

Some examples:
`!a(c+d)`
```
unary_expression(!)
  operand: call_expression
    function: identifier(a)
    arguments:
      binary_expression(+)
        left: identifier(c)
        right: identifier(d)
```

`a(b)(c) + foo(bar)`
```
binary_expression(+)
  left: call_expression
    function: call_expression
      function: identifier(a)
      arguments:
        identifier(b)
    arguments:
      identifier(c)
  right: call_expression
    function: identifier(foo)
    arguments:
      identifier(bar)
```

## Questions
* Debug is hard
* Review the parse expression
