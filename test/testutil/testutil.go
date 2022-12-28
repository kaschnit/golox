package testutil

import (
	"path"
	"runtime"
)

func GetProjectRoot() string {
	_, filepath, _, _ := runtime.Caller(0)
	result := path.Dir(path.Dir(path.Dir(filepath)))
	return result
}
