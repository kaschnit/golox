package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/kaschnit/golox/pkg/scanner"
)

func main() {
	interactive := flag.Bool("interactive", false, "Launch the REPL")
	flag.Parse()

	if *interactive {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			line, _ := reader.ReadString('\n')
			scanner := scanner.NewScanner(line)
			tokens, _ := scanner.ScanAllTokens()
			for i := 0; i < len(tokens); i++ {
				fmt.Println(tokens[i])
			}

		}
	}
}
