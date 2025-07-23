package env

import (
	"os"

	"github.com/gonstruct/validation"
)

func Number(name string, defaultValue ...int) int {
	return validation.Number(os.Getenv(name), defaultValue...)
}
