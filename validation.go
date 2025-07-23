package validation

import "github.com/gonstruct/validation/validators"

func Pipe[T comparable](value T, validators ...validators.Validator[any]) T {
	for _, validator := range validators {
		if err := validator.Validate(value); err != nil {
			panic(err)
		}
	}

	return value
}
