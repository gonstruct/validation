package validation

import (
	"fmt"
	"time"
)

// A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func Duration(input string, defaultValue ...time.Duration) time.Duration {
	if input == "" {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(fmt.Errorf("empty value for duration, no default provided."))
	}

	value, err := time.ParseDuration(input)
	if err != nil {
		panic(fmt.Errorf("invalid duration format: %s", input))
	}

	return value
}
