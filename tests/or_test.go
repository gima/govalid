package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

func TestOr(t *testing.T) {
	test(t, "combination1", true, j.Or(j.String(j.StrIs("a")), j.String(j.StrIs("b"))), "a")
	test(t, "combination2", true, j.Or(j.String(j.StrIs("a")), j.String(j.StrIs("b"))), "b")
	test(t, "combination3", false, j.Or(j.String(j.StrIs("a")), j.String(j.StrIs("b"))), 3)

	test(t, "combination4", true, j.Or(j.String(j.StrIs("a")), j.String(j.StrIs("b")), j.String(j.StrIs("c"))), "b")

	test(t, "combination5", true, j.Or(j.String(j.StrIs("a"))), "a")
	test(t, "combination6", false, j.Or(j.String(j.StrIs("a"))), "b")

	test(t, "combination7", true, j.Or(), nil)
}
