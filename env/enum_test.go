package env_test

import (
	"testing"

	"github.com/gonstruct/validation/env"
	"github.com/stretchr/testify/assert"
)

type testEnumType string

const (
	EnumOption1 testEnumType = "option1"
	EnumOption2 testEnumType = "option2"
	EnumOption3 testEnumType = "option3"
)

func (e testEnumType) Values() []testEnumType {
	return []testEnumType{EnumOption1, EnumOption2, EnumOption3}
}

func (e testEnumType) String() string {
	return string(e)
}

func TestEnum(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Setenv("ENUM_VAR", "option1")
		assert.NotPanics(t, func() { env.Enum[testEnumType]("ENUM_VAR") })

		t.Setenv("ENUM_VAR", "option2")
		assert.NotPanics(t, func() { env.Enum[testEnumType]("ENUM_VAR") })

		assert.PanicsWithError(t, "empty value for enum, no default provided.", func() {
			env.Enum[testEnumType]("EMPTY_ENUM_VAR")
		})

		t.Setenv("INVALID_ENUM_VAR", "invalid_option")
		assert.PanicsWithError(t, "invalid enum value: invalid_option, valid values are: [option1 option2 option3]", func() {
			env.Enum[testEnumType]("INVALID_ENUM_VAR")
		})
	})

	t.Run("optional", func(t *testing.T) {
		t.Setenv("ENUM_VAR", "option1")

		if got := env.Enum("ENUM_VAR", EnumOption2); got != EnumOption1 {
			t.Errorf("Expected option1, got %v", got)
		}

		if got := env.Enum("EMPTY_ENUM_VAR", EnumOption2); got != EnumOption2 {
			t.Errorf("Expected option2, got %v", got)
		}
	})
}
