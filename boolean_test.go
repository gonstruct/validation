package validation_test

import (
	"testing"

	"github.com/gonstruct/validation"
	"github.com/stretchr/testify/assert"
)

func TestBoolean(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() { validation.Boolean("true") })
		assert.NotPanics(t, func() { validation.Boolean("false") })
		assert.NotPanics(t, func() { validation.Boolean("0") })
		assert.NotPanics(t, func() { validation.Boolean("1") })

		assert.PanicsWithError(t, "empty value for boolean, no default provided.", func() { validation.Boolean("") })
		assert.PanicsWithError(t, "invalid boolean format: value", func() { validation.Boolean("value") })
	})

	t.Run("optional", func(t *testing.T) {
		if got := validation.Boolean("true", false); got == false {
			t.Errorf("Expected true, got '%v'", got)
		}

		if got := validation.Boolean("", false); got == true {
			t.Errorf("Expected 25s, got '%v'", got)
		}
	})
}
