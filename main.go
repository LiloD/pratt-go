package main

import (
	"fmt"
	"log"
	"pratt-go/parser"
	"pratt-go/utils"
)

func main() {
	// input := "a + (b + c) + d"
	// input := "(-foo(1, 3+4) + 4) * 5"
	input := "a(b)(c) + foo(bar)"
	// input := "a(b) + 100"
	// input := "!a(b)"
	// input := "!a(c+d)"
	fmt.Println(input)
	parser := parser.NewParser(input)
	exp, err := parser.ParseExpression(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(exp)
	utils.PrintExpression(exp)
}
