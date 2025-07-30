package validation_test

import (
	"testing"
	"time"

	"github.com/gonstruct/validation"
	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() { validation.Duration("1s") })
		assert.PanicsWithError(t, "empty value for duration, no default provided.", func() { validation.Duration("") })
		assert.PanicsWithError(t, "invalid duration format: value", func() { validation.Duration("value") })
	})

	t.Run("optional", func(t *testing.T) {
		if got := validation.Duration("50s", 25*time.Second); got != 50*time.Second {
			t.Errorf("Expected 50s, got '%v'", got)
		}

		if got := validation.Duration("", 25*time.Second); got != 25*time.Second {
			t.Errorf("Expected 25s, got '%v'", got)
		}
	})
}
