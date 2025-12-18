package env_test

import (
	"testing"

	"github.com/gonstruct/validation/env"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Setenv("STRING_VAR", "value")
		assert.NotPanics(t, func() { env.String("STRING_VAR") })

		assert.PanicsWithError(t, "empty value for string, no default provided.", func() {
			env.String("EMPTY_STRING_VAR")
		})
	})

	t.Run("optional", func(t *testing.T) {
		t.Setenv("STRING_VAR", "value")

		if got := env.String("STRING_VAR", "default"); got != "value" {
			t.Errorf("Expected 'value', got '%s'", got)
		}

		if got := env.String("EMPTY_STRING_VAR", "default"); got != "default" {
			t.Errorf("Expected 'default', got '%s'", got)
		}
	})
}
