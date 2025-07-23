package validation_test

import (
	"testing"

	"github.com/gonstruct/validation"
	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() { validation.Number("10") })
		assert.PanicsWithError(t, "empty value for number, no default provided.", func() { validation.Number("") })
		assert.PanicsWithError(t, "invalid number format: value", func() { validation.Number("value") })
	})

	t.Run("optional", func(t *testing.T) {
		if got := validation.Number("10", 25); got != 10 {
			t.Errorf("Expected 10, got '%v'", got)
		}

		if got := validation.Number("", 25); got != 25 {
			t.Errorf("Expected 25, got '%v'", got)
		}
	})
}
