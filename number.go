package validation

import (
	"fmt"
	"strconv"
)

func Number(input string, defaultValue ...int) int {
	if input == "" {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(fmt.Errorf("empty value for number, no default provided."))
	}

	value, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Errorf("invalid number format: %s", input))
	}

	return value
}
