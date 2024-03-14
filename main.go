package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sushil-cmd-r/expr-parser/parser"
	"github.com/sushil-cmd-r/expr-parser/scanner"
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

		sc := scanner.New([]byte(line))
		p := parser.New(sc)
		expr := p.Parse()
		if expr != nil {
			fmt.Println(expr)
		}
	}
}
