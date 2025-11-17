package env

import (
	"os"
	"strings"

	"github.com/gonstruct/validation"
)

func Array[T any](name string, defaultValue ...[]T) []T {
	value := strings.Split(os.Getenv(name), ",")

	return validation.Array((any)(value).([]T), defaultValue...)
}
