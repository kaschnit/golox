package interpreter

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast/interpreter/interpreterutil"
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
	err := interpreterutil.InterpretSourceFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
}

func startInterpreterRepl() {
	cli.NewRepl(func(line string) {
		err := interpreterutil.InterpretLine(line)
		if err != nil {
			fmt.Println(err)
		}
	}).Start()
}
