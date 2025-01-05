package parser_test

import (
	"pratt-go/parser"
	"testing"
)

func TestParseExpression(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		input string
		want  string
	}{
		{
			input: "-a + b",
			want:  "((-a)+b)",
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
