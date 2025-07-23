package validation_test

import (
	"testing"

	"github.com/gonstruct/validation"
	"github.com/stretchr/testify/assert"
)

func TestMinLength(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		assert.Nil(t, validation.MinLength(5).Validate("value"))
		assert.Nil(t, validation.MinLength(5).Validate(6))
	})

	t.Run("invalid", func(t *testing.T) {
		assert.Error(t, validation.MinLength(5).Validate("val"))
		assert.Error(t, validation.MinLength(5).Validate(4))
	})
}

func TestMaxLength(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		assert.Nil(t, validation.MaxLength(5).Validate("val"))
		assert.Nil(t, validation.MaxLength(5).Validate(4))
	})

	t.Run("invalid", func(t *testing.T) {
		assert.Error(t, validation.MaxLength(5).Validate("values"))
		assert.Error(t, validation.MaxLength(5).Validate(6))
	})
}
