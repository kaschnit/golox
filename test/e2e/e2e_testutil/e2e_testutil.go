package e2e_testutil

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/kaschnit/golox/pkg/ast/interpreter/interpreterutil"
	"github.com/kaschnit/golox/test/e2e/testconst"
	"github.com/kaschnit/golox/test/programs"
	"github.com/kaschnit/golox/test/testutil"
)

func GetBuildArtifactRoot() string {
	return path.Join(testutil.GetProjectRoot(), "build")
}

func GetTestBinaryPath() string {
	return path.Join(GetBuildArtifactRoot(), testconst.TEST_BINARY_NAME)
}

func RunTestBinary(args ...string) (string, error) {
	cmd := exec.Command(GetTestBinaryPath(), args...)
	output, err := cmd.Output()
	return string(output), err
}

func ScanTestProgram(programName string) (string, error) {
	return RunTestBinary(testconst.SCANNER_CMD, programs.GetPath(programName))
}

func ParseTestProgram(programName string) (string, error) {
	return RunTestBinary(testconst.PARSER_CMD, programs.GetPath(programName))
}

func InterpretTestProgram(programName string) (string, error) {
	callOutput, err := testutil.CaptureOutput(func() {
		interpreterutil.InterpretSourceFile(programs.GetPath(programName))
	})
	if err != nil {
		return callOutput, err
	}

	cliOutput, err := RunTestBinary(testconst.INTERPRETER_CMD, programs.GetPath(programName))
	if err != nil {
		return cliOutput, err
	}

	if callOutput != cliOutput {
		return "", fmt.Errorf("calling interpreter failed: expected call output (%s) to equal CLI output (%s)", callOutput, cliOutput)
	}

	return callOutput, err
}

func BuildTestBinary() {
	err := exec.Command(testconst.MAKE, "-C", testutil.GetProjectRoot(), testconst.TARGET_BUILD_TEST_BINARY).Run()
	if err != nil {
		panic(fmt.Sprintf("Could not build the test binary for e2e tests: %s", err))
	}
}
