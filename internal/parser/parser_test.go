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

func TestReturnStatement(tester *testing.T) {
	input := `
    return 5;
    return 10;
    return 993322;
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

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			tester.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			tester.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
			continue
		}

	}
}

func TestIdentifierExpress(tester *testing.T) {
	input := `foobar;`

	lex := lexer.New(input)
	parser := New(lex)

	program := parser.ParseProgram()
	checkParserErrors(tester, parser)

	if program == nil {
		tester.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		tester.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		tester.Errorf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		return
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		tester.Errorf("exp not *ast.Identifier. got=%T", stmt.Expression)
		return
	}

	if ident.Value != "foobar" {
		tester.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
		return
	}

	if ident.TokenLiteral() != "foobar" {
		tester.Errorf("ident.TokenLiteral not 'foobar'. got=%q", ident.TokenLiteral())
		return
	}
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
