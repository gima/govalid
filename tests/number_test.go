package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

func TestNumber(t *testing.T) {
	var np *int
	var nnp int = 3

	test(t, "basic number", true, j.Number(), 3)
	test(t, "non-number", false, j.Number(), "")
	test(t, "nil", false, j.Number(), nil)
	test(t, "nil int pointer", false, j.Number(), np)
	test(t, "int pointer", true, j.Number(), &nnp)

	test(t, "numis", true, j.Number(j.NumIs(3)), 3)
	test(t, "!numis", false, j.Number(j.NumIs(3)), 2)

	test(t, "minlen1", false, j.Number(j.NumMin(3)), 2)
	test(t, "minlen2", true, j.Number(j.NumMin(3)), 3)
	test(t, "minlen2", true, j.Number(j.NumMin(3)), 4)

	test(t, "maxlen1", true, j.Number(j.NumMax(3)), 2)
	test(t, "maxlen2", true, j.Number(j.NumMax(3)), 3)
	test(t, "maxlen3", false, j.Number(j.NumMax(3)), 4)

	test(t, "combination1", false, j.Number(j.NumMin(3), j.NumMax(3)), 2)
	test(t, "combination2", true, j.Number(j.NumMin(3), j.NumMax(3)), 3)
	test(t, "combination3", false, j.Number(j.NumMin(3), j.NumMax(3)), 4)
}
