package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

func TestBoolean(t *testing.T) {
	var np *bool
	var nnp bool = true

	test(t, "type check: true", true, j.Boolean(), true)
	test(t, "type check: false", true, j.Boolean(), false)
	test(t, "type check: non-boolean", false, j.Boolean(), nil)

	test(t, "nil bool pointer", false, j.Boolean(), np)
	test(t, "non-nil bool pointer", true, j.Boolean(), &nnp)

	test(t, "should true value", true, j.Boolean(j.BoolIs(true)), true)
	test(t, "!should true value", false, j.Boolean(j.BoolIs(true)), false)

	test(t, "should false value", true, j.Boolean(j.BoolIs(false)), false)
	test(t, "!should false value", false, j.Boolean(j.BoolIs(false)), true)
}
