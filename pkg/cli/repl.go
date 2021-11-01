package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type lineHandler = func(line string)

// Wraps around a function to repeatedly run for REPL-like use.
type Repl struct {
	reader  *bufio.Reader
	handler lineHandler
}

// Create a Repl with a function handler that runs on each line.
func NewRepl(handler lineHandler) *Repl {
	return NewReplWithReader(handler, os.Stdin)
}

// Create a Repl with a function handler that runs on each line and a reader.
func NewReplWithReader(handler lineHandler, reader io.Reader) *Repl {
	return &Repl{
		reader:  bufio.NewReader(reader),
		handler: handler,
	}
}

// Start the REPL procedure.
func (r *Repl) Start() {
	for {
		fmt.Print("> ")
		line, _ := r.reader.ReadString('\n')
		r.handler(line)
	}
}
