package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type enumType string

const (
	EnumValue1 enumType = "value1"
	EnumValue2 enumType = "value2"
)

func (e enumType) Values() []enumType {
	return []enumType{EnumValue1, EnumValue2}
}

func (e enumType) String() string {
	return string(e)
}

func TestEnum(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() { Enum[enumType]("value1") })
		assert.PanicsWithError(t, "empty value for enum, no default provided.", func() { Enum[enumType]("") })
		assert.PanicsWithError(t, "invalid enum value: value3, valid values are: [value1 value2]", func() { Enum[enumType]("value3") })
	})

	t.Run("optional", func(t *testing.T) {
		if got := Enum("value1", EnumValue2); got != EnumValue1 {
			t.Errorf("Expected value1, got '%v'", got)
		}
		if got := Enum("", EnumValue2); got != EnumValue2 {
			t.Errorf("Expected value2, got '%v'", got)
		}
	})
}
