package validation_test

import (
	"testing"

	"github.com/gonstruct/validation"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		result := validation.Url("https://example.com")
		assert.Equal(t, "https", result.Scheme)
		assert.Equal(t, "example.com", result.Host)

		result = validation.Url("http://example.com/path")
		assert.Equal(t, "http", result.Scheme)
		assert.Equal(t, "/path", result.Path)

		result = validation.Url("https://example.com/path?query=value")
		assert.Equal(t, "query=value", result.RawQuery)

		result = validation.Url("https://subdomain.example.com:8080/path")
		assert.Equal(t, "subdomain.example.com:8080", result.Host)

		assert.PanicsWithError(t, "empty value for url, no default provided.", func() { validation.Url("") })
		assert.PanicsWithError(t, "invalid url format: not-a-url", func() { validation.Url("not-a-url") })
		assert.PanicsWithError(t, "invalid url format: ://missing-scheme", func() { validation.Url("://missing-scheme") })
	})

	t.Run("optional", func(t *testing.T) {
		got := validation.Url("https://example.com", "https://default.com")
		if got.String() != "https://example.com" {
			t.Errorf("Expected 'https://example.com', got '%s'", got.String())
		}

		got = validation.Url("", "https://default.com")
		if got.String() != "https://default.com" {
			t.Errorf("Expected 'https://default.com', got '%s'", got.String())
		}
	})
}
