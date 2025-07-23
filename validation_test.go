package validation

import (
	"testing"
)

func TestPipe(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		value := Pipe("test")
		if value != "test" {
			t.Errorf("Expected 'test', got '%s'", value)
		}
	})

}
