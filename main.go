package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sushil-cmd-r/expr-parser/scanner"
	"github.com/sushil-cmd-r/expr-parser/token"
)

const Propmt = "> "

func main() {
	fmt.Println("Welcome to Expression Parser!")

	bufsc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(Propmt)
		scanned := bufsc.Scan()
		if !scanned {
			break
		}

		line := bufsc.Text()
		if line == "exit" {
			fmt.Println("Exiting...")
			break
		}
		if line == "" {
			continue
		}

		lexer := scanner.New([]byte(line))

		for tok := lexer.Next(); tok.Type != token.Eof; tok = lexer.Next() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
