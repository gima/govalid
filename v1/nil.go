package govalid

import (
	"fmt"

	"github.com/gima/govalid/v1/internal"
)

// Constructs a validator which validates data type to be a `nil`.
func Nil() Validator {
	return &nilValidator{}
}

type nilValidator struct {
}

func (r *nilValidator) Validate(data interface{}) (string, error) {
	if value := internal.DigValue(data); !value.IsValid() {
		return "", nil
	} else {
		return "Nil", fmt.Errorf("expected <nil>, was %v", value.Type())
	}
}
