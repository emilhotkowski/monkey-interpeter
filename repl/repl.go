package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/emilhotkowski/monkey-interpeter/lexer"
	"github.com/emilhotkowski/monkey-interpeter/token"
)

const PROMPT = ">> "

func Start(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print(PROMPT)

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
}
