package env

import (
	"os"
	"strings"

	"github.com/gonstruct/validation"
)

const (
	keyValueParts = 2
)

// Map retrieves an environment variable by name and parses it as a map of key-value pairs.
// The environment variable is expected to be a comma-separated string of key=value pairs.
// Example: "authorization=something,anothervalue=this".
func Map(name string, defaultValue ...map[string]string) map[string]string {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return validation.Map(map[string]string{}, defaultValue...)
	}

	result := make(map[string]string)
	pairs := strings.Split(raw, ",")

	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		parts := strings.SplitN(pair, "=", keyValueParts)
		if len(parts) == keyValueParts {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			result[key] = value
		}
	}

	return validation.Map(result, defaultValue...)
}
