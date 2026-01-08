package validation

import (
	"fmt"
)

// Map validates a map input and returns it if non-empty, or a default value if provided.
// If the map is empty and no default is provided, it panics.
func Map[K comparable, V any](input map[K]V, defaultValue ...map[K]V) map[K]V {
	if len(input) == 0 {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(fmt.Errorf("empty value for map, no default provided"))
	}

	return input
}
