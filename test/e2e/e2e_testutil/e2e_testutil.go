package e2e_testutil

import (
	"fmt"
	"os/exec"
	"path"

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
	return RunTestBinary(testconst.INTERPRETER_CMD, programs.GetPath(programName))
}

func BuildTestBinary() {
	err := exec.Command(testconst.MAKE, "-C", testutil.GetProjectRoot(), testconst.TARGET_BUILD_TEST_BINARY).Run()
	if err != nil {
		panic(fmt.Sprintf("Could not build the test binary for e2e tests: %s", err))
	}
}
