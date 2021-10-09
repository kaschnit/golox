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
		Use: "parser",
		Run: runInterpreterCmd,
	}
)

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
