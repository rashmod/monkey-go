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

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

func (parser *Parser) registerPrefix(tok token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tok] = fn
}

func (parser *Parser) registerInfix(tok token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tok] = fn
}

func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{lex: lex, errors: []string{}}

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.registerPrefix(token.IDENT, parser.parseIdentifiers)

	parser.nextToken()
	parser.nextToken()

	return parser
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
		return parser.parseExpressionStatement()
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

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: parser.curToken}

	stmt.Expression = parser.parseExpression(LOWEST)

	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseFns[parser.curToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (parser *Parser) parseIdentifiers() ast.Expression {
	return &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
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
