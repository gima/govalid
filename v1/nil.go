package govalid

import (
	"fmt"
	"reflect"
)

// Validates data types to be nil
func Nil() Validator {
	return &nilValidator{}
}

// validator for a nil value
type nilValidator struct {
}

// the workhorse for the nil validator
func (r *nilValidator) Validate(data interface{}) (string, error) {
	value := reflect.ValueOf(data)

	switch value.Kind() {
	case reflect.Invalid:
		return "", nil
	case reflect.Ptr:
		if !value.IsNil() {
			return "Nil", fmt.Errorf("expected <nil>, was %v", reflect.TypeOf(data))
		}
		return "", nil
	default:
		return "Nil", fmt.Errorf("expected <nil>, was %v", reflect.TypeOf(data))
	}
}
