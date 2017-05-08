package govalid_test

import (
	v "github.com/gima/govalid/v1"
	"testing"
)

func TestNil(t *testing.T) {
	var np *int = nil
	var nnp int = 3

	test(t, "nil value", true, v.Nil(), nil)
	test(t, "non-string", false, v.Nil(), 3)
	test(t, "string", false, v.Nil(), "cat's aren't nil")
	test(t, "empty object", true, v.Nil(), interface{}(nil))
	test(t, "pointer datatype", true, v.Nil(), np)
	test(t, "int pointer", false, v.Nil(), &nnp)
}
