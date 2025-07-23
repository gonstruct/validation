package _examples

import (
	v "github.com/gonstruct/validation"
	"github.com/gonstruct/validation/env"
)

type EnumType string

const (
	EnumValue1 EnumType = "value1"
	EnumValue2 EnumType = "value2"
)

func (e EnumType) Values() []EnumType {
	return []EnumType{EnumValue1, EnumValue2}
}

func (e EnumType) String() string {
	return string(e)
}

func BasicExample() {
	// Required env variables with no default values. Will panic if not set.
	env.String("example_string")
	env.Number("example_number")
	env.Enum[EnumType]("example_enum")

	// Optional env variables with default values. If not set, will return the default value.
	env.String("example_string", "default_value")
	env.Number("example_number", 42)
	env.Enum("example_enum", EnumValue1, EnumValue2)

	// Using validation functions
	v.Pipe(env.String("MY_ENV_VAR"), v.MinLength(1), v.MaxLength(100))
	v.Pipe("invalid_value", v.MinLength(25), v.MaxLength(50))
}
