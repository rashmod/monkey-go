package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rashmod/monkey-go/internal/lexer"
	"github.com/rashmod/monkey-go/internal/token"
)

func Start(in io.Reader, out io.Writer) {
	PROMPT := ">> "

	scanner := bufio.NewScanner(in)

	fmt.Printf(PROMPT)
	scanned := scanner.Scan()

	if !scanned {
		return
	}

	line := scanner.Text()

	l := lexer.New(line)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
