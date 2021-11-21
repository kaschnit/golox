package interpreter

import (
	"fmt"

	ast_interpreter "github.com/kaschnit/golox/pkg/ast/interpreter"
	"github.com/kaschnit/golox/pkg/ast/resolver"
	"github.com/kaschnit/golox/pkg/cli"
	"github.com/kaschnit/golox/pkg/cli/cli_common"
	"github.com/spf13/cobra"
)

type InterpreterFlags struct {
	interactive bool
}

var (
	flags          = &InterpreterFlags{}
	InterpreterCmd = &cobra.Command{
		Use:   "interpreter",
		Run:   runInterpreterCmd,
		Args:  cobra.OnlyValidArgs,
		Short: "Run the golox interpreter",
		Long:  "Run the golox interpreter to execute lox code",
	}
)

func init() {
	InterpreterCmd.Flags().BoolVarP(&flags.interactive, "interactive", "i", false, "Run in interactive mode.")
}

func runInterpreterCmd(_ *cobra.Command, args []string) {
	if flags.interactive {
		startInterpreterRepl()
	} else if len(args) > 0 {
		interpretSourceFile(args[0])
	} else {
		fmt.Println("No input provided. Exiting.")
	}
}

func interpretSourceFile(filepath string) {
	interp := ast_interpreter.NewAstInterpreter()
	resolv := resolver.NewAstResolver(interp)
	err := cli_common.ParseSourceFileAndVisit(filepath, resolv, interp)
	if err != nil {
		fmt.Println(err)
	}
}

func startInterpreterRepl() {
	interp := ast_interpreter.NewAstInterpreter()
	resolv := resolver.NewAstResolver(interp)
	cli.NewRepl(func(line string) {
		err := cli_common.ParseLineAndVisit(line, resolv, interp)
		if err != nil {
			fmt.Println(err)
		}
	}).Start()
}
