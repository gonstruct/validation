package env

import (
	"net/url"
	"os"

	"github.com/gonstruct/validation"
)

// URL validates and parses a URL from an environment variable.
// It accepts any valid URL format according to RFC 3986.
func Url(name string, defaultValue ...string) *url.URL {
	return validation.Url(os.Getenv(name), defaultValue...)
}
