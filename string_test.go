package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() { String("Value") })
		assert.PanicsWithError(t, "empty value for string, no default provided.", func() { String("") })
	})

	t.Run("optional", func(t *testing.T) {
		if got := String("Value", "Default"); got != "Value" {
			t.Errorf("Expected 'Value', got '%s'", got)
		}
		if got := String("", "Default"); got != "Default" {
			t.Errorf("Expected 'Default', got '%s'", got)
		}
	})
}
