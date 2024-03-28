package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sushil-cmd-r/expr-parser/parser"
	"github.com/sushil-cmd-r/expr-parser/scanner"
)

const Propmt = ">> "

func Start() {
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

		sc := scanner.New(line)
		p := parser.New(sc)
		expr, err := p.Parse()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(expr)
		}
	}
}
