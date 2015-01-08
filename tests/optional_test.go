package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

func TestOptional(t *testing.T) {
	var np *bool
	test(t, "string", true, j.Optional(j.String(j.StrIs("a"))), "a")
	test(t, "wrong string", false, j.Optional(j.String(j.StrIs("a"))), "b")

	test(t, "nil ptr", true, j.Optional(j.String(j.StrIs("a"))), np)
}
