package env

import (
	"os"
	"time"

	"github.com/gonstruct/validation"
)

// A time string must match the specified layout format.
// The layout defines the format by showing how the reference time (Mon Jan 2 15:04:05 MST 2006) would be displayed.
func Time(name, layout string, defaultValue ...time.Time) time.Time {
	return validation.Time(os.Getenv(name), layout, defaultValue...)
}
