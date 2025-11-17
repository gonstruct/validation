package validation

import (
	"errors"
	"fmt"
)

func Array[T any](value []T, defaultValue ...[]T) []T {
	if len(value) == 0 {
		if len(defaultValue) == 1 {
			fmt.Println("default applied")
			return defaultValue[0]
		}

		panic(errors.New("empty value for array, no default provided."))
	}

	return value
}
