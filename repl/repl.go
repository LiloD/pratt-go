package repl

import (
	"bufio"
	"fmt"
	"io"
	"pratt-go/parser"
	"pratt-go/precedence"
	"pratt-go/utils"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := parser.NewParser(line)

		exp, err := p.ParseExpression(precedence.Lowest)

		if err != nil {
			fmt.Println(err)
		}

		utils.PrintExpression(exp)
	}
}
