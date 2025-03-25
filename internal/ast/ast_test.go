package ast

import (
	"testing"

	"github.com/rashmod/monkey-go/internal/token"
)

func TestString(tester *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
				Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, Value: "anotherVar"},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		tester.Errorf("program.String() wrong. got=%q", program.String())
	}
}
