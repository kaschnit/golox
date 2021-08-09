package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRuneAlpha(t *testing.T) {
	assert.True(t, IsRuneAlpha('a'))
	assert.True(t, IsRuneAlpha('c'))
	assert.True(t, IsRuneAlpha('p'))
	assert.True(t, IsRuneAlpha('z'))
	assert.True(t, IsRuneAlpha('A'))
	assert.True(t, IsRuneAlpha('C'))
	assert.True(t, IsRuneAlpha('P'))
	assert.True(t, IsRuneAlpha('Z'))
	assert.True(t, IsRuneAlpha('_'))

	assert.False(t, IsRuneAlpha('0'))
	assert.False(t, IsRuneAlpha('1'))
	assert.False(t, IsRuneAlpha('3'))
	assert.False(t, IsRuneAlpha('9'))

	assert.False(t, IsRuneAlpha('['))
	assert.False(t, IsRuneAlpha('&'))
	assert.False(t, IsRuneAlpha('-'))
	assert.False(t, IsRuneAlpha('+'))
}

func TestIsRuneNumeric(t *testing.T) {
	assert.True(t, IsRuneNumeric('0'))
	assert.True(t, IsRuneNumeric('1'))
	assert.True(t, IsRuneNumeric('3'))
	assert.True(t, IsRuneNumeric('9'))

	assert.False(t, IsRuneNumeric('a'))
	assert.False(t, IsRuneNumeric('c'))
	assert.False(t, IsRuneNumeric('p'))
	assert.False(t, IsRuneNumeric('z'))
	assert.False(t, IsRuneNumeric('A'))
	assert.False(t, IsRuneNumeric('C'))
	assert.False(t, IsRuneNumeric('P'))
	assert.False(t, IsRuneNumeric('Z'))
	assert.False(t, IsRuneNumeric('_'))

	assert.False(t, IsRuneNumeric('['))
	assert.False(t, IsRuneNumeric('&'))
	assert.False(t, IsRuneNumeric('-'))
	assert.False(t, IsRuneNumeric('+'))
}

func TestIsRuneAlphaNumeric(t *testing.T) {
	assert.True(t, IsRuneAlphaNumeric('a'))
	assert.True(t, IsRuneAlphaNumeric('c'))
	assert.True(t, IsRuneAlphaNumeric('p'))
	assert.True(t, IsRuneAlphaNumeric('z'))
	assert.True(t, IsRuneAlphaNumeric('A'))
	assert.True(t, IsRuneAlphaNumeric('C'))
	assert.True(t, IsRuneAlphaNumeric('P'))
	assert.True(t, IsRuneAlphaNumeric('Z'))
	assert.True(t, IsRuneAlphaNumeric('_'))

	assert.True(t, IsRuneAlphaNumeric('0'))
	assert.True(t, IsRuneAlphaNumeric('1'))
	assert.True(t, IsRuneAlphaNumeric('3'))
	assert.True(t, IsRuneAlphaNumeric('9'))

	assert.False(t, IsRuneAlphaNumeric('['))
	assert.False(t, IsRuneAlphaNumeric('&'))
	assert.False(t, IsRuneAlphaNumeric('-'))
	assert.False(t, IsRuneAlphaNumeric('+'))
}
