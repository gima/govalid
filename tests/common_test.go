package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

func test(t *testing.T, title string, expectPass bool, validator j.Validator, data interface{}) {
	path, err := validator.Validate(data)

	if err == nil {

		if expectPass {
			return
		} else {
			t.Fatalf("'%s' failed (%s). Path: %s", title, "nil error, but error expected", path)
		}

	} else if err != nil {

		if !expectPass {
			return
		} else {
			t.Fatalf("'%s' failed (%s). Path: %s", title, err, path)
		}

	}
}
