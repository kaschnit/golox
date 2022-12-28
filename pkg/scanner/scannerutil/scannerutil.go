package scannerutil

import (
	"io"
	"os"

	"github.com/kaschnit/golox/pkg/scanner"
	"github.com/kaschnit/golox/pkg/token"
)

func ScanSourceFile(filepath string) ([]*token.Token, error) {
	f, err := os.Open(filepath)
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	sourceCode, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	scanner := scanner.NewScanner(string(sourceCode))
	tokens, err := scanner.ScanAllTokens()
	if err != nil {
		return nil, err
	}

	return tokens, err
}
