package validators

type Validator[T comparable] interface {
	Validate(value T) error
}
