package ast

import (
	"github.com/rashmod/monkey-go/internal/token"
)

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (identifier *Identifier) expressionNode() {}
func (identifier *Identifier) TokenLiteral() string {
	return identifier.Token.Literal
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (letStatement *LetStatement) statementNode() {}
func (letStatement *LetStatement) TokenLiteral() string {
	return letStatement.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (returnStatement *ReturnStatement) statementNode() {}
func (returnStatement *ReturnStatement) TokenLiteral() string {
	return returnStatement.Token.Literal
}
