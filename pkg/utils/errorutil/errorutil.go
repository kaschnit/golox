package errorutil

import "fmt"

func LoxError(line int, msg string) error {
	return fmt.Errorf("[line %d] Error: %s", line, msg)
}
