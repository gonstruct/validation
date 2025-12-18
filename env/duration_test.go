package env_test

import (
	"testing"
	"time"

	"github.com/gonstruct/validation/env"
	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Setenv("DURATION_VAR", "5s")
		assert.NotPanics(t, func() { env.Duration("DURATION_VAR") })

		t.Setenv("DURATION_VAR_MS", "300ms")
		assert.NotPanics(t, func() { env.Duration("DURATION_VAR_MS") })

		t.Setenv("DURATION_VAR_COMPLEX", "2h45m")
		assert.NotPanics(t, func() { env.Duration("DURATION_VAR_COMPLEX") })

		assert.PanicsWithError(t, "empty value for duration, no default provided.", func() {
			env.Duration("EMPTY_DURATION_VAR")
		})

		t.Setenv("INVALID_DURATION_VAR", "not_a_duration")
		assert.PanicsWithError(t, "invalid duration format: not_a_duration", func() {
			env.Duration("INVALID_DURATION_VAR")
		})
	})

	t.Run("optional", func(t *testing.T) {
		t.Setenv("DURATION_VAR", "10s")

		if got := env.Duration("DURATION_VAR", 5*time.Second); got != 10*time.Second {
			t.Errorf("Expected 10s, got %v", got)
		}

		if got := env.Duration("EMPTY_DURATION_VAR", 5*time.Second); got != 5*time.Second {
			t.Errorf("Expected 5s, got %v", got)
		}
	})
}
