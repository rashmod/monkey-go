package parser

import (
	"fmt"

	"github.com/rashmod/monkey-go/internal/ast"
	"github.com/rashmod/monkey-go/internal/lexer"
	"github.com/rashmod/monkey-go/internal/token"
)

type Parser struct {
	lex       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex, errors: []string{}}

	p.nextToken()
	p.nextToken()

	return p
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lex.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !parser.curTokenIs(token.EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.nextToken()
	}

	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
	case token.LET:
		return parser.parseLetStatment()
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseLetStatment() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.curToken}

	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	parser.nextToken()

	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: parser.curToken}

	parser.nextToken()

	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

func (parser *Parser) expectPeek(tok token.TokenType) bool {
	if parser.peekTokenIs(tok) {
		parser.nextToken()
		return true
	}

	parser.peekError(tok)
	return false
}

func (parser *Parser) curTokenIs(tok token.TokenType) bool {
	return parser.curToken.Type == tok
}

func (parser *Parser) peekTokenIs(tok token.TokenType) bool {
	return parser.peekToken.Type == tok
}

func (parser *Parser) peekError(tok token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tok, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) Errors() []string {
	return parser.errors
}
