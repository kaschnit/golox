package scanner

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/cli"
	"github.com/kaschnit/golox/pkg/scanner"
	"github.com/spf13/cobra"
)

type ScannerFlags struct {
	interactive bool
}

var (
	flags      = &ScannerFlags{}
	ScannerCmd = &cobra.Command{
		Use:   "scanner",
		Run:   runScannerCmd,
		Args:  cobra.OnlyValidArgs,
		Short: "Run the golox scanner",
		Long:  "Run the golox scanner to produce a stream of tokens from lox source code",
	}
)

func init() {
	ScannerCmd.Flags().BoolVarP(&flags.interactive, "interactive", "i", false, "Run in interactive mode.")
}

func runScannerCmd(_ *cobra.Command, _ []string) {
	if flags.interactive {
		startScannerRepl()
	}
}

func startScannerRepl() {
	repl := cli.NewRepl(func(line string) {
		// Tokenize the input.
		scanner := scanner.NewScanner(line)
		tokens, err := scanner.ScanAllTokens()
		if err != nil {
			fmt.Println(err)
			return
		}

		for i := 0; i < len(tokens); i++ {
			fmt.Println(tokens[i])
		}
	})
	repl.Start()
}
