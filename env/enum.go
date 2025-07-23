package env

import (
	"github.com/gonstruct/validation"
)

type enumComparable[T any] interface {
	comparable
	Values() []T
	String() string
}

func Enum[T enumComparable[T]](input string, defaultValue ...T) T {
	return validation.Enum(input, defaultValue...)
}
