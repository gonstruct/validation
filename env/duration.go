package env

import (
	"os"
	"time"

	"github.com/gonstruct/validation"
)

// Duration parses a duration from an environment variable.
// A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func Duration(name string, defaultValue ...time.Duration) time.Duration {
	return validation.Duration(os.Getenv(name), defaultValue...)
}
