package validation

import "github.com/gonstruct/validation/validators"

func MinLength(minLength int) validators.MinLength {
	return validators.MinLength{Min: minLength}
}

func MaxLength(maxLength int) validators.MaxLength {
	return validators.MaxLength{Max: maxLength}
}
