package env_test

import (
	"testing"

	"github.com/gonstruct/validation/env"
	"github.com/stretchr/testify/assert"
)

func TestBoolean(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Setenv("BOOL_VAR_TRUE", "true")
		assert.NotPanics(t, func() { env.Boolean("BOOL_VAR_TRUE") })

		t.Setenv("BOOL_VAR_FALSE", "false")
		assert.NotPanics(t, func() { env.Boolean("BOOL_VAR_FALSE") })

		t.Setenv("BOOL_VAR_1", "1")
		assert.NotPanics(t, func() { env.Boolean("BOOL_VAR_1") })

		t.Setenv("BOOL_VAR_0", "0")
		assert.NotPanics(t, func() { env.Boolean("BOOL_VAR_0") })

		assert.PanicsWithError(t, "empty value for boolean, no default provided.", func() {
			env.Boolean("EMPTY_BOOL_VAR")
		})

		t.Setenv("INVALID_BOOL_VAR", "not_a_bool")
		assert.PanicsWithError(t, "invalid boolean format: not_a_bool", func() {
			env.Boolean("INVALID_BOOL_VAR")
		})
	})

	t.Run("optional", func(t *testing.T) {
		t.Setenv("BOOL_VAR", "true")

		if got := env.Boolean("BOOL_VAR", false); got != true {
			t.Errorf("Expected true, got %v", got)
		}

		if got := env.Boolean("EMPTY_BOOL_VAR", false); got != false {
			t.Errorf("Expected false, got %v", got)
		}
	})
}
