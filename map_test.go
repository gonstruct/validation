package validation_test

import (
	"testing"

	"github.com/gonstruct/validation"
)

func TestMap(t *testing.T) {
	t.Run("returns map when non-empty", func(t *testing.T) {
		input := map[string]string{"key": "value"}

		result := validation.Map(input)
		if len(result) != 1 || result["key"] != "value" {
			t.Errorf("expected map with key=value, got %v", result)
		}
	})

	t.Run("returns default when empty and default provided", func(t *testing.T) {
		input := map[string]string{}
		defaultValue := map[string]string{"default": "value"}

		result := validation.Map(input, defaultValue)
		if len(result) != 1 || result["default"] != "value" {
			t.Errorf("expected default map, got %v", result)
		}
	})

	t.Run("panics when empty and no default", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic but did not get one")
			}
		}()

		validation.Map(map[string]string{})
	})
}
