package parser

import (
	"testing"

	"github.com/rashmod/monkey-go/internal/ast"
	"github.com/rashmod/monkey-go/internal/lexer"
)

func TestLetStatement(tester *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 838383;
    `

	lex := lexer.New(input)
	parser := New(lex)

	program := parser.ParseProgram()
	checkParserErrors(tester, parser)

	if program == nil {
		tester.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		tester.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(tester, stmt, test.expectedIdentifier) {
			return
		}

	}
}

func testLetStatement(tester *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		tester.Errorf("token literal not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		tester.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		tester.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		tester.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(tester *testing.T, parser *Parser) {
	errors := parser.Errors()

	if len(errors) == 0 {
		return
	}

	tester.Errorf("parser has %d errors", len(errors))
	for _, err := range errors {
		tester.Errorf("parser error: %s", err)
	}

	tester.FailNow()
}
