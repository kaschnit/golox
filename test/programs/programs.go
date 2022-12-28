package programs

import (
	"io"
	"os"
	"path"
	"runtime"
)

func GetDirectoryPath() string {
	_, filepath, _, _ := runtime.Caller(0)
	return path.Dir(filepath)
}

func GetPath(subPath string) string {
	return path.Join(GetDirectoryPath(), subPath)
}

func ReadProgramText(subPath string) string {
	filepath := GetPath(subPath)
	f, err := os.Open(filepath)
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	sourceCode, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(sourceCode)
}
