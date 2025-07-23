package env

import (
	"os"

	"github.com/gonstruct/validation"
)

func String(name string, defaultValue ...string) string {
	return validation.String(os.Getenv(name), defaultValue...)
}
