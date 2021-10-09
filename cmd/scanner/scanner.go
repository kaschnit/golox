package scanner

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/repl"
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

		for i := 0; i < len(tokens); i++ {
			fmt.Println(tokens[i])
		}
	})
	repl.Start()
}
