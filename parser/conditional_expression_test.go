package parser_test

import (
	"pratt-go/parser"
	"testing"
)

func TestConditionalExpression(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		input string
		want  string
	}{
		{
			input: "a?b:c",
			want:  "(a?b:c)",
		},
		{
			input: "a+1?b+2:c+3",
			want:  "((a+1)?(b+2):(c+3))",
		},
		{
			input: "a?b:c?d:e",
			want:  "(a?b:(c?d:e))",
		},
		{
			input: "a+b?c*d:e/f",
			want:  "((a+b)?(c*d):(e/f))",
		},
		{
			input: "a(b?c:d,e+f)",
			want:  "a((b?c:d),(e+f))",
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
