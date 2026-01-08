package env_test

import (
	"os"
	"testing"

	"github.com/gonstruct/validation/env"
)

func TestMap(t *testing.T) {
	t.Run("parses key=value pairs", func(t *testing.T) {
		os.Setenv("TEST_MAP", "authorization=something,anothervalue=this")
		defer os.Unsetenv("TEST_MAP")

		result := env.Map("TEST_MAP")
		if len(result) != 2 {
			t.Errorf("expected 2 entries, got %d", len(result))
		}
		if result["authorization"] != "something" {
			t.Errorf("expected authorization=something, got %s", result["authorization"])
		}
		if result["anothervalue"] != "this" {
			t.Errorf("expected anothervalue=this, got %s", result["anothervalue"])
		}
	})

	t.Run("handles whitespace", func(t *testing.T) {
		os.Setenv("TEST_MAP", "key1 = value1 , key2 = value2")
		defer os.Unsetenv("TEST_MAP")

		result := env.Map("TEST_MAP")
		if result["key1"] != "value1" {
			t.Errorf("expected key1=value1, got %s", result["key1"])
		}
		if result["key2"] != "value2" {
			t.Errorf("expected key2=value2, got %s", result["key2"])
		}
	})

	t.Run("returns default when empty", func(t *testing.T) {
		os.Unsetenv("TEST_MAP")
		defaultValue := map[string]string{"default": "value"}

		result := env.Map("TEST_MAP", defaultValue)
		if len(result) != 1 || result["default"] != "value" {
			t.Errorf("expected default map, got %v", result)
		}
	})

	t.Run("panics when empty and no default", func(t *testing.T) {
		os.Unsetenv("TEST_MAP")
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic but did not get one")
			}
		}()
		env.Map("TEST_MAP")
	})

	t.Run("handles empty pairs", func(t *testing.T) {
		os.Setenv("TEST_MAP", "key1=value1,,key2=value2")
		defer os.Unsetenv("TEST_MAP")

		result := env.Map("TEST_MAP")
		if len(result) != 2 {
			t.Errorf("expected 2 entries, got %d", len(result))
		}
	})

	t.Run("ignores malformed pairs", func(t *testing.T) {
		os.Setenv("TEST_MAP", "key1=value1,invalidpair,key2=value2")
		defer os.Unsetenv("TEST_MAP")

		result := env.Map("TEST_MAP")
		if len(result) != 2 {
			t.Errorf("expected 2 entries (ignoring invalid), got %d", len(result))
		}
	})

	t.Run("handles values with equals signs", func(t *testing.T) {
		os.Setenv("TEST_MAP", "key=value=with=equals")
		defer os.Unsetenv("TEST_MAP")

		result := env.Map("TEST_MAP")
		if result["key"] != "value=with=equals" {
			t.Errorf("expected key=value=with=equals, got %s", result["key"])
		}
	})
}
