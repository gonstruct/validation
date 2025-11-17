package env_test

import (
	"testing"

	"github.com/gonstruct/validation/env"
	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.PanicsWithError(t, "empty value for array, no default provided.", func() { env.Array[string]("EMPTY_VAR") })
	})

	t.Run("optional", func(t *testing.T) {
		t.Setenv("ARRAY_VARIABLE", "value")
		if got := env.Array("ARRAY_VARIABLE", []string{"default"}); got[0] != "value" {
			t.Errorf("Expected value, got '%v'", got)
		}

		t.Setenv("ARRAY_VARIABLE", "")
		if got := env.Array("ARRAY_VARIABLE", []string{"default"}); got[0] != "default" {
			t.Errorf("Expected default, got '%v'", got)
		}
	})
}
