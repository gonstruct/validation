package validation_test

import (
	"testing"
	"time"

	"github.com/gonstruct/validation"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		assert.NotPanics(t, func() {
			got := validation.Time("2025-07-30T15:04:05Z", time.RFC3339)
			expected, _ := time.Parse(time.RFC3339, "2025-07-30T15:04:05Z")
			assert.Equal(t, expected, got)
		})

		assert.PanicsWithError(t, "empty value for time, no default provided.", func() {
			validation.Time("", time.RFC3339)
		})

		assert.PanicsWithError(t, "invalid time format: value", func() {
			validation.Time("value", time.RFC3339)
		})
	})

	t.Run("optional", func(t *testing.T) {
		defaultTime := time.Date(2024, 12, 25, 10, 30, 0, 0, time.UTC)

		if got := validation.Time("03:04PM", time.Kitchen, defaultTime); !got.Equal(time.Date(0, 1, 1, 15, 4, 0, 0, time.UTC)) {
			t.Errorf("Expected parsed Kitchen time, got '%v'", got)
		}

		if got := validation.Time("", time.Kitchen, defaultTime); !got.Equal(defaultTime) {
			t.Errorf("Expected default time %v, got '%v'", defaultTime, got)
		}
	})
}
