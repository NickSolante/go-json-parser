package token

type TokenType string

const (
	OBRACKET    = "{"
	CBRACKET    = "}"
	OPAREN      = "("
	CPAREN      = ")"
	OSQRBRACKET = "["
	CSQRBRACKET = "]"
)

type Token struct {
	Type    TokenType
	Literal string
}
