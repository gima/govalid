package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

func TestString(t *testing.T) {
	var np *string
	var nnp string = "a"

	test(t, "non-string", false, j.String(), 3)
	test(t, "basic string", true, j.String(), "")
	test(t, "nil", false, j.String(), nil)
	test(t, "nil string pointer", false, j.String(), np)
	test(t, "string pointer", true, j.String(), &nnp)

	test(t, "equals", true, j.String(j.StrIs("abc")), "abc")
	test(t, "!equals", false, j.String(j.StrIs("abc")), "abd")

	test(t, "minlen1", false, j.String(j.StrMin(3)), "aa")
	test(t, "minlen2", true, j.String(j.StrMin(3)), "aaa")
	test(t, "minlen3", true, j.String(j.StrMin(3)), "aaaa")

	test(t, "maxlen1", true, j.String(j.StrMax(4)), "aaa")
	test(t, "maxlen2", true, j.String(j.StrMax(4)), "aaaa")
	test(t, "maxlen3", false, j.String(j.StrMax(4)), "aaaaa")

	test(t, "regexp1", true, j.String(j.StrRegExp("^.{3}$")), "bbb")
	test(t, "regexp2", false, j.String(j.StrRegExp("^.{3}$")), "bbbb")
	test(t, "regexp3", false, j.String(j.StrRegExp("[")), "c")

	test(t, "combination1", false, j.String(j.StrMin(3), j.StrMax(3)), "cc")
	test(t, "combination2", true, j.String(j.StrMin(3), j.StrMax(3)), "ccc")
	test(t, "combination1", false, j.String(j.StrMin(3), j.StrMax(3)), "cccc")

}
