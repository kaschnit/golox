package interpreter

import (
	"fmt"

	ast_interpreter "github.com/kaschnit/golox/pkg/ast/interpreter"
	"github.com/kaschnit/golox/pkg/cli"
	"github.com/spf13/cobra"
)

type InterpreterAlgorithm string

const (
	InterpreterAlgorithmAST      InterpreterAlgorithm = "ast"
	InterpreterAlgorithmByteCode InterpreterAlgorithm = "bytecode"
)

type InterpreterFlags struct {
	interactive bool
	algorithm   string
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
	InterpreterCmd.Flags().StringVarP(&flags.algorithm, "algorithm", "a", string(InterpreterAlgorithmByteCode), "The interpreter algorithm to use. One of: 'ast', 'bytecode'.")
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
	interp := ast_interpreter.NewInterpreterWrapper()
	err := interp.InterpretSourceFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
}

func startInterpreterRepl() {
	interp := ast_interpreter.NewInterpreterWrapper()
	cli.NewRepl(func(line string) {
		err := interp.InterpretLine(line)
		if err != nil {
			fmt.Println(err)
		}
	}).Start()
}
