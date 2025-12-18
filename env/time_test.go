package env_test

import (
	"testing"
	"time"

	"github.com/gonstruct/validation/env"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Setenv("TIME_VAR", "2025-12-18T15:04:05Z")
		assert.NotPanics(t, func() {
			got := env.Time("TIME_VAR", time.RFC3339)
			expected, _ := time.Parse(time.RFC3339, "2025-12-18T15:04:05Z")
			assert.Equal(t, expected, got)
		})

		t.Setenv("TIME_VAR_KITCHEN", "3:04PM")
		assert.NotPanics(t, func() {
			got := env.Time("TIME_VAR_KITCHEN", time.Kitchen)
			expected, _ := time.Parse(time.Kitchen, "3:04PM")
			assert.Equal(t, expected, got)
		})

		assert.PanicsWithError(t, "empty value for time, no default provided.", func() {
			env.Time("EMPTY_TIME_VAR", time.RFC3339)
		})

		t.Setenv("INVALID_TIME_VAR", "not_a_time")
		assert.PanicsWithError(t, "invalid time format: not_a_time", func() {
			env.Time("INVALID_TIME_VAR", time.RFC3339)
		})
	})

	t.Run("optional", func(t *testing.T) {
		defaultTime := time.Date(2024, 12, 25, 10, 30, 0, 0, time.UTC)

		t.Setenv("TIME_VAR", "2025-12-18T15:04:05Z")
		if got := env.Time("TIME_VAR", time.RFC3339, defaultTime); !got.Equal(time.Date(2025, 12, 18, 15, 4, 5, 0, time.UTC)) {
			t.Errorf("Expected parsed time, got %v", got)
		}

		if got := env.Time("EMPTY_TIME_VAR", time.RFC3339, defaultTime); !got.Equal(defaultTime) {
			t.Errorf("Expected default time %v, got %v", defaultTime, got)
		}
	})
}
