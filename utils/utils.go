package utils

import (
	"fmt"
	"pratt-go/parser"
)

func PrintExpression(exp parser.Expression) {
	printExp(exp, 0)
}

func printExp(exp parser.Expression, indent int) {
	switch exp.(type) {
	case *parser.IdentifierExpression:
		identExp := exp.(*parser.IdentifierExpression)
		fmt.Printf("identifier(%s)\n", identExp.Name)
	case *parser.NumberExpression:
		numberExp := exp.(*parser.NumberExpression)
		fmt.Printf("number_literal(%s)\n", numberExp.Number)
	case *parser.UnaryExpression:
		unaryExp := exp.(*parser.UnaryExpression)
		fmt.Printf("unary_expression(%s)\n", unaryExp.Tok.Literal)
		fmt.Printf("%soperand: ", whitespace(indent+2))
		printExp(unaryExp.Operand, indent+2)
	case *parser.BinaryExpression:
		binaryExp := exp.(*parser.BinaryExpression)
		fmt.Printf("binary_expression(%s)\n", binaryExp.Tok.Literal)
		fmt.Printf("%sleft: ", whitespace(indent+2))
		printExp(binaryExp.Left, indent+2)
		fmt.Printf("%sright: ", whitespace(indent+2))
		printExp(binaryExp.Right, indent+2)
	case *parser.ConditionalExpression:
		condExp := exp.(*parser.ConditionalExpression)
		fmt.Printf("conditional_expression\n")
		fmt.Printf("%scondition: ", whitespace(indent+2))
		printExp(condExp.Condition, indent+2)
		fmt.Printf("%sconsequence: ", whitespace(indent+2))
		printExp(condExp.Consequence, indent+2)
		fmt.Printf("%salternative: ", whitespace(indent+2))
		printExp(condExp.Alternative, indent+2)
	case *parser.CallExpression:
		callExp := exp.(*parser.CallExpression)
		fmt.Printf("call_expression\n")
		fmt.Printf("%sfunction: ", whitespace(indent+2))
		printExp(callExp.Function, indent+2)
		fmt.Printf("%sarguments:\n", whitespace(indent+2))
		for _, argExp := range callExp.Arguments {
			fmt.Printf("%s", whitespace(indent+4))
			printExp(argExp, indent+4)
		}
	case *parser.ParenthesizedExpression:
		parentExp := exp.(*parser.ParenthesizedExpression)
		fmt.Printf("parenthesized_expression\n")
		fmt.Printf("%s", whitespace(indent+2))
		printExp(parentExp.Child, indent+2)
	}
}

func whitespace(indent int) string {
	s := ""
	for i := 0; i < indent; i++ {
		s += " "
	}
	return s
}
