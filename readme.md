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

## Exponent Operator 
Exponent Operator `^` is another interesting case, it's a binary infix operator, it has higher precedence than Product
and Sum.

Updated Precedence List:
```go
const (
	Lowest = iota
	Sum
	Product
    Exponent    // <- here!
	Prefix
	Call    
)
```

But consider `2^3^2`, it first caculates `3^2 = 9`, and then `2^9 = 512`, so it has the `right association`!

Remember other binary operators, for example `+` and expression `1 + 2 + 3`, we first caculate `1 + 2 = 3`, and then 
`3 + 3 = 6`, it has the `left association`.

We need to figure out how to handle the right association, and surprisely, it quite simple!

For normal left associative binary operator, the parserlet looks like this, we pass the precedence of the operator to the
parseExpression, inside the parseExpression call, it will return early if it encounter another same precedence operator, 
which will lead to left association.
```go
func BinaryOperatorParselet(parser *Parser, token *token.Token, left Expression) (Expression, error) {
	parser.NextToken()
	precedence := parser.getPrecedence(token.Type)
	exp, err := parser.ParseExpression(precedence)
	if err != nil {
		return nil, err
	}

	binaryExp := &BinaryExpression{
		Tok:   token,
		Left:  left,
		Right: exp,
	}

	return binaryExp, nil
}
```

To handle the right association, we simply lower the precedence value passed to parseExpression
```go
func BinaryRightOperatorParselet(parser *Parser, token *token.Token, left Expression) (Expression, error) {
	parser.NextToken()
	precedence := parser.getPrecedence(token.Type)
	exp, err := parser.ParseExpression(precedence - 1) // this is the only line changed!
	if err != nil {
		return nil, err
	}

	binaryExp := &BinaryExpression{
		Tok:   token,
		Left:  left,
		Right: exp,
	}

	return binaryExp, nil
}
```

Oh, don't forget to register with the right associative binary parserlet
```go
//...
p.infixParseletMap[token.PLUS] = BinaryOperatorParselet
p.infixParseletMap[token.MINUS] = BinaryOperatorParselet
p.infixParseletMap[token.ASTERISK] = BinaryOperatorParselet
p.infixParseletMap[token.SLASH] = BinaryOperatorParselet

p.infixParseletMap[token.CARET] = BinaryRightOperatorParselet // register ^ with Right Associative Parselet Function
```

Some example:

`2^3^2`
```
binary_expression(^)
  left: identifier(a)
  right: binary_expression(^)
    left: identifier(b)
    right: identifier(c)
```

`b+c*d^e-f/g`
```
binary_expression(-)
  left: binary_expression(+)
    left: identifier(b)
    right: binary_expression(*)
      left: identifier(c)
      right: binary_expression(^)
        left: identifier(d)
        right: identifier(e)
  right: binary_expression(/)
    left: identifier(f)
    right: identifier(g)
```


## Conditional Expression
Conditional Expression `a ? b : c` can be treat as a binary expression even though it is actually a Ternary Operator.

The precedence should be lower than any Arithmetic operator

Updated Precedence List:
```go
const (
	Lowest = iota
    Conditional     // <- here!
	Sum
	Product
    Exponent    
	Prefix
	Call    
)
```

And the implementation is quite straightforward:
```go
type ConditionalExpression struct {
	Tok         *token.Token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func ConditionalOperatorParselet(parser *Parser, tok *token.Token, left Expression) (Expression, error) {
	parser.NextToken()
	// parse consequence
	consequence, err := parser.ParseExpression(precedence.Lowest)
	if err != nil {
		return nil, err
	}

	// expect : and move over it
	parser.ExpectNextToken(token.COLON)
	parser.NextToken()

	// parse alternative
	alternative, err := parser.ParseExpression(precedence.Lowest)
	if err != nil {
		return nil, err
	}

	return &ConditionalExpression{
		Condition:   left,
		Consequence: consequence,
		Alternative: alternative,
	}, nil
}
```

Some examples:

`a?b:c`
```
conditional_expression
  condition: identifier(a)
  consequence: identifier(b)
  alternative: identifier(c)
```

`a+1?b+2:c+3`
```
conditional_expression
  condition: binary_expression(+)
    left: identifier(a)
    right: number_literal(1)
  consequence: binary_expression(+)
    left: identifier(b)
    right: number_literal(2)
  alternative: binary_expression(+)
    left: identifier(c)
    right: number_literal(3)
```

`a?b:c?d:e`
```
conditional_expression
  condition: identifier(a)
  consequence: identifier(b)
  alternative: conditional_expression
    condition: identifier(c)
    consequence: identifier(d)
    alternative: identifier(e)
```


## Questions
* Debug is hard
* Review the parse expression
