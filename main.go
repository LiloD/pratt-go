package main

import (
	"os"
	"pratt-go/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
