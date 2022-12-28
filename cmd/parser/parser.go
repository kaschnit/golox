package parser

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast/astutil"
	"github.com/kaschnit/golox/pkg/ast/printer"
	"github.com/kaschnit/golox/pkg/cli"
	"github.com/spf13/cobra"
)

type ParserFlags struct {
	interactive bool
}

var (
	flags     = &ParserFlags{}
	ParserCmd = &cobra.Command{
		Use:   "parser",
		Run:   runParserCmd,
		Args:  cobra.OnlyValidArgs,
		Short: "Run the golox parser",
		Long:  "Run the golox parser to produce an AST from lox source code",
	}
)

func init() {
	ParserCmd.Flags().BoolVarP(&flags.interactive, "interactive", "i", false, "Run in interactive mode.")
}

func runParserCmd(_ *cobra.Command, args []string) {
	if flags.interactive {
		startParserRepl()
	} else if len(args) > 0 {
		parseSourceFile(args[0])
	}
}

func parseSourceFile(filepath string) {
	visitor := printer.NewAstPrinter()
	err := astutil.ParseSourceFileAndVisit(filepath, visitor)
	if err != nil {
		fmt.Println(err)
	}
}

func startParserRepl() {
	visitor := printer.NewAstPrinter()
	cli.NewRepl(func(line string) {
		err := astutil.ParseLineAndVisit(line, visitor)
		if err != nil {
			fmt.Println(err)
		}
	}).Start()
}
