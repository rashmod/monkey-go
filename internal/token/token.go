package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// SPECIAL TOKENS
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENTIFIERS AND LITERALS
	IDENT = "IDENT"
	INT   = "INT"

	// OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	// DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"

	// BRACKETS
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	LET      = "LET"
	FUNCTION = "FUNCTION"
)
