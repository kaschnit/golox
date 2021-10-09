package parser

import (
	"github.com/kaschnit/golox/pkg/ast/printer"
	"github.com/kaschnit/golox/pkg/repl"
	"github.com/kaschnit/golox/pkg/repl/repl_common"
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
	visitor := printer.NewAstPrinter()
	repl.NewRepl(func(line string) {
		repl_common.ParseLineAndVisit(visitor, line)
	}).Start()
}
