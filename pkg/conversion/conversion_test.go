package conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type randomStruct struct{}

func assertIsTruthy(t *testing.T, val interface{}) {
	assert.True(t, IsTruthy(val), "Expected %v to be truthy", val)
}

func assertIsFalsy(t *testing.T, val interface{}) {
	assert.False(t, IsTruthy(val), "Expected %v to not be truthy", val)
}

func TestIsTruthy_NilValue(t *testing.T) {
	assertIsFalsy(t, nil)
}

func TestIsTruthy_BooleanValue(t *testing.T) {
	assertIsTruthy(t, true)
	assertIsFalsy(t, false)
}

func TestIsTruthy_FloatValue_Truthy(t *testing.T) {
	assertIsTruthy(t, float64(0.5))
	assertIsTruthy(t, float32(0.5))
	assertIsTruthy(t, -0.5)
	assertIsTruthy(t, 0.00001)
	assertIsTruthy(t, -0.00001)
	assertIsTruthy(t, 40)
	assertIsTruthy(t, -40)
	assertIsTruthy(t, 99.5)
	assertIsTruthy(t, -99.5)
}

func TestIsTruthy_FloatValue_Falsy(t *testing.T) {
	assertIsFalsy(t, 0)
	assertIsFalsy(t, 0.0000)
	assertIsFalsy(t, -0)
	assertIsFalsy(t, -0.0000)
}

func TestIsTruthy_Pointer_Truthy(t *testing.T) {
	assertIsTruthy(t, &randomStruct{})

	someBool := false
	assertIsTruthy(t, &someBool)
}

func TestIsTruth_StructValue_Truthy(t *testing.T) {
	assertIsTruthy(t, randomStruct{})
}

func TestIsTruth_StringValue_Truthy(t *testing.T) {
	assertIsTruthy(t, " ")
	assertIsTruthy(t, "   ")
	assertIsTruthy(t, "\t")
	assertIsTruthy(t, "\n")
	assertIsTruthy(t, "0")
	assertIsTruthy(t, "1")
	assertIsTruthy(t, "-1")
	assertIsTruthy(t, "12345")
	assertIsTruthy(t, "hello")
	assertIsTruthy(t, "a")
	assertIsTruthy(t, "abcd")
	assertIsTruthy(t, "abcd ")
	assertIsTruthy(t, " abcd")
	assertIsTruthy(t, " abcd ")
}

func TestIsTruth_StringValue_Falsy(t *testing.T) {
	assertIsFalsy(t, "")
}

func TestToFloat_NilValue(t *testing.T) {
	result, ok := ToFloat(nil)
	assert.Zero(t, result)
	assert.False(t, ok)
}

func TestToFloat_ZeroValue(t *testing.T) {
	var val float64
	var ok bool

	val, ok = ToFloat(0)
	assert.True(t, ok)
	assert.Equal(t, 0.0, val)

	val, ok = ToFloat(float64(0.00))
	assert.True(t, ok)
	assert.Equal(t, 0.00, val)

	val, ok = ToFloat(float32(0.00))
	assert.True(t, ok)
	assert.Equal(t, 0.00, val)
}

func TestToFloat_Boolean(t *testing.T) {
	var val float64
	var ok bool

	val, ok = ToFloat(true)
	assert.True(t, ok)
	assert.Equal(t, 1.0, val)

	val, ok = ToFloat(false)
	assert.True(t, ok)
	assert.Equal(t, 0.0, val)
}

func TestToFloat_StringInvalid(t *testing.T) {
	var val float64
	var ok bool

	val, ok = ToFloat("")
	assert.False(t, ok)
	assert.Zero(t, val)

	val, ok = ToFloat(" ")
	assert.False(t, ok)
	assert.Zero(t, val)

	val, ok = ToFloat("hello")
	assert.False(t, ok)
	assert.Zero(t, val)
}
