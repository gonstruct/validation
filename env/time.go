package env

import (
	"os"
	"time"

	"github.com/gonstruct/validation"
)

func Time(name, layout string, defaultValue ...time.Time) time.Time {
	return validation.Time(os.Getenv(name), layout, defaultValue...)
}
