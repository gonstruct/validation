package validation

import (
	"fmt"
)

type enumComparable[T any] interface {
	comparable
	Values() []T
	String() string
}

func Enum[T enumComparable[T]](input string, defaultValue ...T) T {
	if input == "" {
		if len(defaultValue) == 1 {
			return defaultValue[0]
		}

		panic(fmt.Errorf("empty value for enum, no default provided."))
	}

	var zero T
	for _, val := range zero.Values() {
		if val.String() == input {
			return val
		}
	}

	panic(fmt.Errorf("invalid enum value: %s, valid values are: %v", input, zero.Values()))
}
