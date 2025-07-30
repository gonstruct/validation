package env

import (
	"os"

	"github.com/gonstruct/validation"
)

// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
func Boolean(name string, defaultValue ...bool) bool {
	return validation.Boolean(os.Getenv(name), defaultValue...)
}
