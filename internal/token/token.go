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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	// COMPARISON
	BANG   = "!"
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
