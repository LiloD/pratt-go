package parser_test

import (
	"pratt-go/parser"
	"testing"
)

func TestParser_ParseExpression(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		input string
		want  string
	}{
		{
			name:  "test1",
			input: "-a + b",
			want:  "(binary_expression +: left: (unary_expression -: (operand: (identifier: a))), right: (identifier: b)))",
		},
		{
			name:  "test2",
			input: "1+2+3",
			want:  "(binary_expression +: left: (binary_expression +: left: (number_literal: 1), right: (number_literal: 2))), right: (number_literal: 3)))",
		},
		{
			name:  "test3",
			input: "1*2+3",
			want:  "(binary_expression +: left: (binary_expression *: left: (number_literal: 1), right: (number_literal: 2))), right: (number_literal: 3)))",
		},
		{
			name:  "test4",
			input: "1+2*3",
			want:  "(binary_expression +: left: (number_literal: 1), right: (binary_expression *: left: (number_literal: 2), right: (number_literal: 3)))))",
		},
		{
			name:  "test4",
			input: "1+2*3 + 4",
			want:  "(binary_expression +: left: (binary_expression +: left: (number_literal: 1), right: (binary_expression *: left: (number_literal: 2), right: (number_literal: 3))))), right: (number_literal: 4)))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := parser.NewParser(tt.input)
			exp, err := p.ParseExpression(0)
			if err != nil {
				t.Errorf("ParseExpression() failed: %v", err)
			}
			got := exp.String()
			if got != tt.want {
				t.Errorf("ParseExpression() got v.s want:\n%s\n%s", got, tt.want)
			}
		})
	}
}
