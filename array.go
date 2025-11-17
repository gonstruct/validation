package validation

import (
	"errors"
)

// Array retrieves an environment variable by name and validates it as an array of strings.
// The environment variable is expected to be a comma-separated string.
func Array[T any](value []T, defaultValue ...[]T) []T {
	if len(value) == 0 {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(errors.New("empty value for array, no default provided."))
	}

	return value
}
