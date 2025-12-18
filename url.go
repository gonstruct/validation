package validation

import (
	"fmt"
	"net/url"
)

// URL validates and parses a URL string.
// It accepts any valid URL format according to RFC 3986.
func Url(input string, defaultValue ...string) *url.URL {
	if input == "" {
		if len(defaultValue) == 1 {
			input = defaultValue[0]
		} else {
			panic(fmt.Errorf("empty value for url, no default provided."))
		}
	}

	parsed, err := url.ParseRequestURI(input)
	if err != nil {
		panic(fmt.Errorf("invalid url format: %s", input))
	}

	return parsed
}
