package validation_test

import (
	"testing"

	"github.com/gonstruct/validation"
	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() { validation.Array([]string{"value"}) })

		assert.PanicsWithError(t, "empty value for array, no default provided.", func() { validation.Array([]string{}) })
		// assert.PanicsWithError(t, "invalid array format: value", func() { validation.Array() })
	})

	t.Run("optional", func(t *testing.T) {
		if got := validation.Array([]string{"value"}, []string{"default"}); got[0] != "value" {
			t.Errorf("Expected value, got '%v'", got)
		}

		if got := validation.Array([]string{}, []string{"default"}); got[0] != "default" {
			t.Errorf("Expected default, got '%v'", got)
		}
	})
}
