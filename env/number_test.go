package env_test

import (
	"testing"

	"github.com/gonstruct/validation/env"
	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Setenv("NUMBER_VAR", "42")
		assert.NotPanics(t, func() { env.Number("NUMBER_VAR") })

		assert.PanicsWithError(t, "empty value for number, no default provided.", func() {
			env.Number("EMPTY_NUMBER_VAR")
		})

		t.Setenv("INVALID_NUMBER_VAR", "not_a_number")
		assert.PanicsWithError(t, "invalid number format: not_a_number", func() {
			env.Number("INVALID_NUMBER_VAR")
		})
	})

	t.Run("optional", func(t *testing.T) {
		t.Setenv("NUMBER_VAR", "42")

		if got := env.Number("NUMBER_VAR", 100); got != 42 {
			t.Errorf("Expected 42, got %d", got)
		}

		if got := env.Number("EMPTY_NUMBER_VAR", 100); got != 100 {
			t.Errorf("Expected 100, got %d", got)
		}
	})
}
