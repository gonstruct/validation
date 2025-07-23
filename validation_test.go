package validation_test

import (
	"testing"

	"github.com/gonstruct/validation"
)

func TestPipe(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		value := validation.Pipe("test")
		if value != "test" {
			t.Errorf("Expected 'test', got '%s'", value)
		}
	})
}
