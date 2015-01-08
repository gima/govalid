package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

func TestAnd(t *testing.T) {
	test(t, "combination1", true, j.And(), nil)

	test(t, "combination2", false, j.And(j.String(j.StrMin(3)), j.String(j.StrMax(3))), "aa")
	test(t, "combination3", true, j.And(j.String(j.StrMin(3)), j.String(j.StrMax(3))), "aaa")
	test(t, "combination4", false, j.And(j.String(j.StrMin(3)), j.String(j.StrMax(3))), "aaaa")

	test(t, "combination5", false, j.And(j.String(j.StrMin(3)), j.String(j.StrMax(4))), "bb")
	test(t, "combination6", true, j.And(j.String(j.StrMin(3)), j.String(j.StrMax(4))), "bbb")
	test(t, "combination7", true, j.And(j.String(j.StrMin(3)), j.String(j.StrMax(4))), "bbbb")
	test(t, "combination8", false, j.And(j.String(j.StrMin(3)), j.String(j.StrMax(4))), "bbbbb")
}
