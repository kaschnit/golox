package interpreter_test

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	GO                   = "go"
	RUN                  = "run"
	PROJECT_ROOT         = "../../../"
	INTERPRETER          = "interpreter"
	SAMPLE_PROGRAMS_ROOT = "../../programs/"
)

func sourceFilePath(filename string) string {
	return fmt.Sprintf("%s%s", SAMPLE_PROGRAMS_ROOT, filename)
}

func readOutput(t *testing.T, filename string) (string, error) {
	cmd := exec.Command(GO, RUN, PROJECT_ROOT, INTERPRETER, sourceFilePath(filename))
	output, err := cmd.Output()
	return string(output), err
}

func Test_HelloWorld(t *testing.T) {
	result, err := readOutput(t, "helloworld.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Hello, world!", result)
}
