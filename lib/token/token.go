package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// structural characters
	OBRACKET         = "{"
	CBRACKET         = "}"
	OPAREN           = "("
	CPAREN           = ")"
	OSQRBRACKET      = "["
	CSQRBRACKET      = "]"
	COLON            = ":"
	COMMA            = ","
	DOUBLE_BACKSLASH = "\\\\"

	IDENT = "IDENT"
	INT   = "INT"

	//values
	FALSE  = "false"
	NULL   = "null"
	TRUE   = "true"
	OBJECT = "object"
	ARRAY  = "array"
	NUMBER = "number"
	STRING = "string"

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Numbers
	BACKSLASH    byte = '\\'
	MINUS        byte = '-'
	DECIMALPOINT byte = '.'
	SPACE        byte = ' '
	QUOTE        byte = '"'
)
