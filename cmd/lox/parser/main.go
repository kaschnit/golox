package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/kaschnit/golox/pkg/ast/printer"
	"github.com/kaschnit/golox/pkg/parser"
	"github.com/kaschnit/golox/pkg/scanner"
)

func main() {
	interactive := flag.Bool("interactive", false, "Launch the REPL")
	flag.Parse()

	astPrinter := printer.NewAstPrinter()

	if *interactive {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			line, _ := reader.ReadString('\n')

			scanner := scanner.NewScanner(line)
			tokens, scanErrs := scanner.ScanAllTokens()
			for i := 0; i < len(scanErrs); i++ {
				fmt.Println(scanErrs[i])
			}

			parser := parser.NewParser(tokens)
			programAst, parseErrs := parser.Parse()
			if len(scanErrs) == 0 {
				for i := 0; i < len(parseErrs); i++ {
					fmt.Println(parseErrs[i])
				}
			}

			programAst.Accept(astPrinter)
		}
	}
}
