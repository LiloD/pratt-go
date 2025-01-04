package main

import (
	"fmt"
	"log"
	"pratt-go/parser"
	"pratt-go/utils"
)

func main() {
	// ---
	input := "1 + 2 * 3 + 4"
	fmt.Println(input)
	parser := parser.NewParser(input)
	exp, err := parser.ParseExpression(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(exp)
	utils.PrintExpression(exp)
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
