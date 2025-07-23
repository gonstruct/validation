package validation

import (
	"fmt"
)

func String(input string, defaultValue ...string) string {
	if input == "" {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(fmt.Errorf("empty value for string, no default provided."))
	}

	return input
}
