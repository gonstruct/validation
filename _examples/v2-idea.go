package _examples

import (
	"github.com/gonstruct/validation/env"
)

func V2Idea() {
	env.ArrayWithOptions("ARRAY_KEY", env.WithDefault([]string{"default1", "default2"}), env.WithSeperator(";"))
}
