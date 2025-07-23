package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() { Number("10") })
		assert.PanicsWithError(t, "empty value for number, no default provided.", func() { Number("") })
		assert.PanicsWithError(t, "invalid number format: value", func() { Number("value") })
	})

	t.Run("optional", func(t *testing.T) {
		if got := Number("10", 25); got != 10 {
			t.Errorf("Expected 10, got '%v'", got)
		}
		if got := Number("", 25); got != 25 {
			t.Errorf("Expected 25, got '%v'", got)
		}
	})
}
