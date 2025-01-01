package main

import (
	"fmt"
	"log"
	"pratt-go/parser"
	//    "pratt-go/lexer"
	// "pratt-go/token"
)

func main() {
	// input := "from + offset(time)"
	// lexer := lexer.NewLexer(input)
	// for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
	// 	log.Printf("tok: %+v", tok)
	// }

	// ---
	// input = "+a+b" // (binary_expression +: (left: (unary_expression +: (operand: (NAME: a))), right: (NAME: b)))
	// (binary_expression +: (left: (NAME: 1), right: (binary_expression *: (left: (NAME: 2), right: (NAME: 3)))))
	// input = "1+2*3"
	// (binary_expression *: (left: (NAME: 1), right: (binary_expression +: (left: (NAME: 2), right: (NAME: 3)))))
	// input = "1 * 2 + 3"
	// (binary_expression +: (left: (NAME: 1), right: (binary_expression +: (left: (NAME: 2), right: (NAME: 3)))))
	input := "1 + 2 + 3"
	fmt.Println(input)
	parser := parser.NewParser(input)
	exp, err := parser.ParseExpression(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(exp)
	// fmt.Println("hello pratt go")
}

// func PrintExpression(exp parser.Expression, indent int) {
// 	switch exp.(type) {
// 	case *parser.NameExpression:
// 		fmt.Printf("%s%s", buildWhitespace(indent), exp.String())
// 	case *parser.UnaryExpression:
// 		fmt.Printf("%s%s", buildWhitespace(indent), exp)
// 	case *parser.BinaryExpression:
// 	}
// }
//
// func buildWhitespace(n int) string {
// 	s := ""
// 	for i := 0; i < n; i++ {
// 		s += "  "
// 	}
// 	return s
// }
