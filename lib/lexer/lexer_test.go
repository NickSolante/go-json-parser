package lexer

import (
	"go-json-parser/lib/token"
	"testing"
)

func TestLexer(t *testing.T) {
	t.Run("initial empty json{}", func(t *testing.T) {
		input := "{}"

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.OBRACKET, "{"},
			{token.CBRACKET, "}"},
		}

		l := NewLexer(input)

		for i, tt := range tests {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("test[%d] - tokenType wrong. expected =%q, got=%q", i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
			}
		}
	})

	t.Run("string hello", func(t *testing.T) {
		input := `"train"`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.STRING, "train"},
		}

		l := NewLexer(input)

		for i, tt := range tests {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("test[%d] - tokenType wrong. expected =%q, got=%q", i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
			}
		}
	})
}
