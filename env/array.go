package env

import (
	"os"
	"strings"

	"github.com/gonstruct/validation"
)

func Array[T string](name string, defaultValue ...[]T) []T {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return validation.Array([]T{}, defaultValue...)
	}

	value := strings.Split(strings.TrimSpace(os.Getenv(name)), ",")
	return validation.Array((any)(value).([]T), defaultValue...)
}
