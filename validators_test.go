package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinLength(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		assert.Nil(t, MinLength(5).Validate("value"))
		assert.Nil(t, MinLength(5).Validate(6))
	})

	t.Run("invalid", func(t *testing.T) {
		assert.Error(t, MinLength(5).Validate("val"))
		assert.Error(t, MinLength(5).Validate(4))
	})
}

func TestMaxLength(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		assert.Nil(t, MaxLength(5).Validate("val"))
		assert.Nil(t, MaxLength(5).Validate(4))
	})

	t.Run("invalid", func(t *testing.T) {
		assert.Error(t, MaxLength(5).Validate("values"))
		assert.Error(t, MaxLength(5).Validate(6))
	})
}
