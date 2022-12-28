package programs

import (
	"io"
	"os"
	"path"
	"runtime"
)

func GetPath(subPath string) string {
	_, filepath, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filepath), subPath)
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
