package validation_test

import (
	"testing"

	"github.com/gonstruct/validation"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() { validation.String("Value") })
		assert.PanicsWithError(t, "empty value for string, no default provided.", func() { validation.String("") })
	})

	t.Run("optional", func(t *testing.T) {
		if got := validation.String("Value", "Default"); got != "Value" {
			t.Errorf("Expected 'Value', got '%s'", got)
		}

		if got := validation.String("", "Default"); got != "Default" {
			t.Errorf("Expected 'Default', got '%s'", got)
		}
	})
}
