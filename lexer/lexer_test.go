package lexer_test

import (
	"pratt-go/lexer"
	"pratt-go/token"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []*token.Token
	}{
		{
			name:  "test1",
			input: "(12 + 123) * 999 + foo",
			want: []*token.Token{
				{Type: token.LPARA, Literal: "("},
				{Type: token.NUMBER, Literal: "12"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: "123"},
				{Type: token.RPARA, Literal: ")"},
				{Type: token.ASTERISK, Literal: "*"},
				{Type: token.NUMBER, Literal: "999"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.IDENT, Literal: "foo"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.input)

			for i := 0; i < len(tt.want); i++ {
				got := l.NextToken()
				want := tt.want[i]

				if got.Type != want.Type || got.Literal != want.Literal {
					t.Errorf("NextToken() = %+v, want %v", got, want)
				}
			}
			got := l.NextToken()

			if got.Type != token.EOF {
				t.Errorf("NextToken() = %+v, want %v", got, token.EOF)
			}
		})
	}
}
