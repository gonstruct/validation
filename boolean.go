package validation

import (
	"fmt"
	"strconv"
)

// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
func Boolean(input string, defaultValue ...bool) bool {
	if input == "" {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(fmt.Errorf("empty value for boolean, no default provided."))
	}

	value, err := strconv.ParseBool(input)
	if err != nil {
		panic(fmt.Errorf("invalid boolean format: %s", input))
	}

	return value
}
