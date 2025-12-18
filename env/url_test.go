package env_test

import (
	"testing"

	"github.com/gonstruct/validation/env"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Setenv("URL_VAR", "https://example.com")
		result := env.Url("URL_VAR")
		assert.Equal(t, "https", result.Scheme)
		assert.Equal(t, "example.com", result.Host)

		t.Setenv("URL_VAR_WITH_PATH", "https://example.com/path?query=value")
		result = env.Url("URL_VAR_WITH_PATH")
		assert.Equal(t, "/path", result.Path)
		assert.Equal(t, "query=value", result.RawQuery)

		assert.PanicsWithError(t, "empty value for url, no default provided.", func() {
			env.Url("EMPTY_URL_VAR")
		})

		t.Setenv("INVALID_URL_VAR", "not-a-valid-url")
		assert.PanicsWithError(t, "invalid url format: not-a-valid-url", func() {
			env.Url("INVALID_URL_VAR")
		})
	})

	t.Run("optional", func(t *testing.T) {
		t.Setenv("URL_VAR", "https://example.com")

		got := env.Url("URL_VAR", "https://default.com")
		if got.String() != "https://example.com" {
			t.Errorf("Expected 'https://example.com', got '%s'", got.String())
		}

		got = env.Url("EMPTY_URL_VAR", "https://default.com")
		if got.String() != "https://default.com" {
			t.Errorf("Expected 'https://default.com', got '%s'", got.String())
		}
	})
}
