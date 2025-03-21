package lexer

import (
	"fmt"

	"github.com/rashmod/monkey-go/internal/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readCh()

	return l
}

func (lexer *Lexer) readCh() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

func (lexer *Lexer) peekCh() byte {
	if lexer.position >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) NextToken() token.Token {
	lexer.skipWhitespace()

	var tok token.Token

	switch lexer.ch {
	case '=':
		if lexer.peekCh() == '=' {
			ch := lexer.ch
			lexer.readCh()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(lexer.ch)}
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '*':
		tok = newToken(token.ASTERISK, lexer.ch)
	case '/':
		tok = newToken(token.SLASH, lexer.ch)
	case '!':

		if lexer.peekCh() == '=' {
			ch := lexer.ch
			lexer.readCh()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(lexer.ch)}
		} else {
			tok = newToken(token.BANG, lexer.ch)
		}

	case '<':
		tok = newToken(token.LT, lexer.ch)
	case '>':
		tok = newToken(token.GT, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readCh()

	return tok
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.isWhitespace() {
		lexer.readCh()
	}
}

func (lexer *Lexer) isWhitespace() bool {
	return lexer.ch == ' ' || lexer.ch == '\n' || lexer.ch == '\t' || lexer.ch == '\r'
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position

	for isLetter(lexer.ch) {
		lexer.readCh()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.ch) {
		lexer.readCh()
	}

	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9')
}

func newToken(toketype token.TokenType, ch byte) token.Token {
	return token.Token{Type: toketype, Literal: string(ch)}
}

func (lexer *Lexer) Print() {
	fmt.Printf("lexer: %s\n", lexer.input)
	fmt.Printf("position: %d\n", lexer.position)
	fmt.Printf("readPosition: %d\n", lexer.readPosition)
	fmt.Printf("ch: %c\n", lexer.ch)
}
