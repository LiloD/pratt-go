package main

import (
	"fmt"
	"log"
	"pratt-go/parser"
	"pratt-go/utils"
)

func main() {
	input := "a + (b + c) + d"
	fmt.Println(input)
	parser := parser.NewParser(input)
	exp, err := parser.ParseExpression(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(exp)
	utils.PrintExpression(exp)
}
