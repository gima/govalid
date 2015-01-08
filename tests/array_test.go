package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

func TestArray(t *testing.T) {
	var np *[]interface{}
	var nnsp = []interface{}{1, 2, 3}
	var nnap = [3]interface{}{1, 2, 3}

	test(t, "basic slice", true, j.Array(), []interface{}{1, 2, 3})
	test(t, "nil", false, j.Array(), nil)
	test(t, "nil slice pointer", false, j.Array(), np)
	test(t, "slice pointer", true, j.Array(), &nnsp)
	test(t, "non-array/slice", false, j.Array(), 3)

	test(t, "basic array ", true, j.Array(), [3]interface{}{1, 2, 3})
	test(t, "nil array pointer", true, j.Array(), [3]interface{}{1, 2, 3})
	test(t, "array pointer", true, j.Array(), &nnap)

	test(t, "int slice", true, j.Array(), []int{1, 2, 3})

	test(t, "minlen1", false, j.Array(j.ArrMin(3)), []interface{}{1, 2})
	test(t, "minlen2", true, j.Array(j.ArrMin(3)), []interface{}{1, 2, 3})
	test(t, "minlen3", true, j.Array(j.ArrMin(3)), []interface{}{1, 2, 3, 4})

	test(t, "maxlen", true, j.Array(j.ArrMax(3)), []interface{}{1, 2})
	test(t, "maxlen2", true, j.Array(j.ArrMax(3)), []interface{}{1, 2, 3})
	test(t, "maxlen3", false, j.Array(j.ArrMax(3)), []interface{}{1, 2, 3, 4})

	test(t, "combination1", false, j.Array(j.ArrMin(3), j.ArrMax(3)), []interface{}{1, 2})
	test(t, "combination2", true, j.Array(j.ArrMin(3), j.ArrMax(3)), []interface{}{1, 2, 3})
	test(t, "combination3", false, j.Array(j.ArrMin(3), j.ArrMax(3)), []interface{}{1, 2, 3, 4})

	test(t, "each1", true, j.Array(j.ArrEach(j.Number(j.NumMin(3)))), []interface{}{})
	test(t, "each2", false, j.Array(j.ArrEach(j.Number(j.NumMin(3)))), []interface{}{2, 3})
	test(t, "each3", true, j.Array(j.ArrEach(j.Number(j.NumMin(3)))), []interface{}{3, 4, 5})
}
