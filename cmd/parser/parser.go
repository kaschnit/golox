package parser

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast/printer"
	"github.com/kaschnit/golox/pkg/parser"
	"github.com/kaschnit/golox/pkg/repl"
	"github.com/kaschnit/golox/pkg/scanner"
	"github.com/spf13/cobra"
)

type ParserFlags struct {
	interactive bool
}

var (
	flags     = &ParserFlags{}
	ParserCmd = &cobra.Command{
		Use: "parser",
		Run: runParserCmd,
	}

	astPrinter = printer.NewAstPrinter()
)

func init() {
	ParserCmd.Flags().BoolVarP(&flags.interactive, "interactive", "i", false, "Run in interactive mode.")
}

func runParserCmd(_ *cobra.Command, _ []string) {
	if flags.interactive {
		startParserRepl()
	}
}

func startParserRepl() {
	repl := repl.NewRepl(func(line string) {
		// Tokenize the input.
		scanner := scanner.NewScanner(line)
		tokens, errs := scanner.ScanAllTokens()
		if len(errs) > 0 {
			for i := 0; i < len(errs); i++ {
				fmt.Println(errs[i])
			}
			return
		}

		// Parse the input.
		parser := parser.NewParser(tokens)
		programAst, errs := parser.Parse()
		if len(errs) > 0 {
			for i := 0; i < len(errs); i++ {
				fmt.Println(errs[i])
			}
			return
		}

		// Print the AST.
		programAst.Accept(astPrinter)
	})
	repl.Start()
}
