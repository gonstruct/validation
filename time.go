package validation

import (
	"fmt"
	"time"
)

// A time string must match the specified layout format.
// The layout defines the format by showing how the reference time (Mon Jan 2 15:04:05 MST 2006) would be displayed.
func Time(input, layout string, defaultValue ...time.Time) time.Time {
	if input == "" {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(fmt.Errorf("empty value for time, no default provided."))
	}

	value, err := time.Parse(layout, input)
	if err != nil {
		panic(fmt.Errorf("invalid time format: %s", input))
	}

	return value
}
