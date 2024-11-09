package token

type TokenType string

const (
	// structural characters
	OBRACKET       string = "{"
	CBRACKET       string = "}"
	OPAREN         string = "("
	CPAREN         string = ")"
	OSQRBRACKET    string = "["
	CSQRBRACKET    string = "]"
	COLON          string = ":"
	DOUBLEQOUTE    string = "\""
	BACKSLASH      string = "\\"
	DOUBLBACKSLASH string = "\\\\"

	IDENT string = "IDENT"
	INT   string = "INT"

	FALSE string = "false"
	NULL  string = "null"
	TRUE  string = "true"
)

type Token struct {
	Type    TokenType
	Literal string
}
