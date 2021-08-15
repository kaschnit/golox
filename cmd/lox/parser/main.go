// Entrypoint for the Lox parser.
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

			// Tokenize the input.
			scanner := scanner.NewScanner(line)
			tokens, errs := scanner.ScanAllTokens()
			if len(errs) > 0 {
				for i := 0; i < len(errs); i++ {
					fmt.Println(errs[i])
				}
				continue
			}

			// Parse the input.
			parser := parser.NewParser(tokens)
			programAst, errs := parser.Parse()
			if len(errs) == 0 {
				for i := 0; i < len(errs); i++ {
					fmt.Println(errs[i])
				}
				continue
			}

			// Print the AST.
			programAst.Accept(astPrinter)
		}
	}
}
