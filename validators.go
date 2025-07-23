package validation

import "github.com/gonstruct/validation/validators"

func MinLength(min int) validators.MinLength {
	return validators.MinLength{Min: min}
}

func MaxLength(max int) validators.MaxLength {
	return validators.MaxLength{Max: max}
}
