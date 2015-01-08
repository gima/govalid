package jsonv2

import (
	"fmt"
	"github.com/gima/jsonv2/internal"
	"reflect"
)

// -----------------------------------------------------------------------------

// Object validation function. Parameter is reflection to a map.
type ObjectOpt func(m *reflect.Value) (path string, err error)

// -----------------------------------------------------------------------------

// Construct an object validator using the specified validation functions.
//
// Currently this validator supports only map data type.
func Object(opts ...ObjectOpt) Validator {
	return &objectValidator{opts}
}

// -----------------------------------------------------------------------------

// Array validator function for validating every key (of key->value pair) with the specified validator.
func ObjKeys(v Validator) ObjectOpt {
	return func(m *reflect.Value) (path string, _ error) {
		for _, key := range m.MapKeys() {
			k := key.Interface()
			if path, err := v.Validate(k); err != nil {
				return fmt.Sprintf("Key(%v)->%s", k, path), err
			}
		}
		return "", nil
	}
}

// -----------------------------------------------------------------------------

// Array validator function for validating every value (of key->value pair) with the specified validator.
func ObjValues(v Validator) ObjectOpt {
	return func(m *reflect.Value) (path string, _ error) {
		for _, key := range m.MapKeys() {
			value := m.MapIndex(key).Interface()
			if path, err := v.Validate(value); err != nil {
				return fmt.Sprintf("Key[%v].Value->%s", key.Interface(), path), err
			}
		}
		return "", nil
	}
}

// -----------------------------------------------------------------------------

// Array validator function for validating a specific key's value (of key->value pair) with the specified validator.
// If the map under validation doesn't have such key, nil is passed to the Validator (hint: Optional Validator).
//
// keys keys keys
func ObjKV(key interface{}, v Validator) ObjectOpt {
	return func(m *reflect.Value) (path string, _ error) {
		var refkey reflect.Value

		if key == nil {
			refkey = reflect.Zero(m.Type().Key())
		} else {
			refkey = reflect.ValueOf(key)
		}

		var value interface{}

		refval := m.MapIndex(refkey)
		if !refval.IsValid() {
			value = nil
		} else {
			value = refval.Interface()
		}

		if path, err := v.Validate(value); err != nil {
			return fmt.Sprintf("Key[%v].Value->%s", key, path), err
		}

		return "", nil
	}
}

// -----------------------------------------------------------------------------

// validator for an object
type objectValidator struct {
	opts []ObjectOpt
}

// -----------------------------------------------------------------------------

// the actual workhorse for object validator
func (r *objectValidator) Validate(data interface{}) (string, error) {

	value, err := internal.ReflectOrIndirect(data)
	if err != nil {
		return "Object", fmt.Errorf("expected (*)map, got %s", err)
	}

	if value.Kind() != reflect.Map {
		return "Object", fmt.Errorf("expected (*)map, got %s", value.Type())
	}

	for _, o := range r.opts {
		if path, err := o(value); err != nil {
			return "Object->" + path, err
		}
	}

	return "", nil
}

/*
func (self *Object) Validate(data *interface{}) (string, error) {
	var validate *map[string]interface{}

	switch tmp := (*data).(type) {
	case map[string]interface{}:
		validate = &tmp
	case *map[string]interface{}:
		validate = tmp
	default:
		return "Object", fmt.Errorf("expected map[string]interface{}, was %v", reflect.TypeOf(*data))
	}

	if self.Each != (ObjectEach{}) {
		// can do loop without copying?
		for key, val := range *validate {
			tmpKey := interface{}(&key)
			if desc, err := self.Each.KeyValidator.Validate(&tmpKey); err != nil {
				return fmt.Sprintf(`Object(Each, key("%s"))->%s`, key, desc), err
			}
			tmpVal := interface{}(val)
			if desc, err := self.Each.DataValidator.Validate(&tmpVal); err != nil {
				return fmt.Sprintf(`Object(Each, key("%s").Value->%s`, key, desc), err
			}
		}
		return "Object(Each)", nil
	}

	if self.Properties != nil {
		// can do loop without copying?
		for _, proofProperty := range self.Properties {
			val := interface{}((*validate)[proofProperty.Key])
			if desc, err := proofProperty.DataValidator.Validate(&val); err != nil {
				return fmt.Sprintf(`Object(Items, key("%s").Value)->%s`, proofProperty.Key, desc), err
			}
		}
		return "Object(Items)", nil
	}

	return "Object", nil
}
*/
