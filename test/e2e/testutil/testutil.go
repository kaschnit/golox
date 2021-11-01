package testutil

import (
	"fmt"
	"os/exec"
	"path"
	"runtime"

	"github.com/kaschnit/golox/test/e2e/testconst"
)

func GetProjectRoot() string {
	_, filepath, _, _ := runtime.Caller(0)
	result := path.Dir(path.Dir(path.Dir(path.Dir(filepath))))
	return result
}

func GetBuildArtifactRoot() string {
	return path.Join(GetProjectRoot(), "build")
}

func GetTestBinaryPath() string {
	return path.Join(GetBuildArtifactRoot(), testconst.TEST_BINARY_NAME)
}

func RunTestBinary(args ...string) (string, error) {
	cmd := exec.Command(GetTestBinaryPath(), args...)
	output, err := cmd.Output()
	return string(output), err
}

func InterpretTestProgram(programName string) (string, error) {
	return RunTestBinary(testconst.INTERPRETER_CMD, GetTestProgramPath(programName))
}

func GetTestProgramPath(subPath string) string {
	return path.Join(GetProjectRoot(), "test", "programs", subPath)
}

func BuildTestBinary() {
	err := exec.Command(testconst.MAKE, "-C", GetProjectRoot(), testconst.TARGET_BUILD_TEST_BINARY).Run()
	if err != nil {
		panic(fmt.Sprintf("Could not build the test binary for e2e tests: %s", err))
	}
}
