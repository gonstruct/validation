package validators

import "fmt"

type MinLength struct {
	Min int
}

func (v MinLength) Validate(raw any) error {
	switch value := raw.(type) {
	case string:
		if len(value) < v.Min {
			return fmt.Errorf("value must be at least %d characters long", v.Min)
		}
	case int:
		if value < v.Min {
			return fmt.Errorf("value must be at least %d", v.Min)
		}
	default:
		return fmt.Errorf("unsupported type for MinLength: %T", value)
	}

	return nil
}

type MaxLength struct {
	Max int
}

func (v MaxLength) Validate(raw any) error {
	switch value := raw.(type) {
	case string:
		if len(value) > v.Max {
			return fmt.Errorf("value must be at most %d characters long", v.Max)
		}
	case int:
		if value > v.Max {
			return fmt.Errorf("value must be at most %d", v.Max)
		}
	default:
		return fmt.Errorf("unsupported type for MaxLength: %T", value)
	}

	return nil
}
