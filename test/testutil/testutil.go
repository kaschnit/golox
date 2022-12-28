package testutil

import (
	"io"
	"os"
	"path"
	"runtime"
)

func GetProjectRoot() string {
	_, filepath, _, _ := runtime.Caller(0)
	result := path.Dir(path.Dir(path.Dir(filepath)))
	return result
}

func CaptureOutput(handler func()) (string, error) {
	rescueStdout := os.Stdout
	rescueStderr := os.Stderr
	defer func() {
		os.Stdout = rescueStdout
		os.Stderr = rescueStderr
	}()

	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	os.Stdout = w
	os.Stderr = w

	handler()

	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
