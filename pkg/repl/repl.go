package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type lineHandler = func(line string)

type Repl struct {
	reader  *bufio.Reader
	handler lineHandler
}

func NewRepl(handler lineHandler) *Repl {
	return NewReplWithReader(handler, os.Stdin)
}

func NewReplWithReader(handler lineHandler, reader io.Reader) *Repl {
	return &Repl{
		reader:  bufio.NewReader(reader),
		handler: handler,
	}
}

func (r *Repl) Start() {
	for {
		fmt.Print("> ")
		line, _ := r.reader.ReadString('\n')
		r.handler(line)
	}
}
