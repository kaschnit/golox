package interpreter

import (
	ast_interpreter "github.com/kaschnit/golox/pkg/ast/interpreter"
	"github.com/kaschnit/golox/pkg/repl"
	"github.com/kaschnit/golox/pkg/repl/repl_common"
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

func runInterpreterCmd(cmd *cobra.Command, args []string) {
	if flags.interactive {
		startInterpreterRepl()
	}
}

func startInterpreterRepl() {
	visitor := ast_interpreter.NewAstInterpreter()
	repl.NewRepl(func(line string) {
		repl_common.ParseLineAndVisit(visitor, line)
	}).Start()
}
