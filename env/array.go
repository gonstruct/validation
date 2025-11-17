package env

import (
	"os"
	"strings"

	"github.com/gonstruct/validation"
)

// Array retrieves an environment variable by name and validates it as an array of strings.
// The environment variable is expected to be a comma-separated string.
func Array[T string](name string, defaultValue ...[]T) []T {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return validation.Array([]T{}, defaultValue...)
	}

	value := strings.Split(strings.TrimSpace(os.Getenv(name)), ",")

	return validation.Array((any)(value).([]T), defaultValue...)
}
