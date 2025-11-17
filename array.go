package validation

import (
	"errors"
)

func Array[T any](value []T, defaultValue ...[]T) []T {
	if len(value) == 0 {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(errors.New("empty value for array, no default provided."))
	}

	return value
}
