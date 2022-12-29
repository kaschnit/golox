package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturn_ErrorString(t *testing.T) {
	returnWrapper := NewReturn("hello")
	assert.Equal(t, "RETURN hello", returnWrapper.Error())
}
