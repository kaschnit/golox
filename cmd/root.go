package cmd

import (
	"fmt"
	"os"

	"github.com/kaschnit/golox/cmd/interpreter"
	"github.com/kaschnit/golox/cmd/parser"
	"github.com/kaschnit/golox/cmd/scanner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "golox",
}

func init() {
	rootCmd.AddCommand(scanner.ScannerCmd)
	rootCmd.AddCommand(parser.ParserCmd)
	rootCmd.AddCommand(interpreter.InterpreterCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
