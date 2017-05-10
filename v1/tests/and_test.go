package govalid_test

import (
	"fmt"
	"testing"

	v "github.com/gima/govalid/v1"
	"github.com/stretchr/testify/require"
)

func TestAnd(t *testing.T) {
	test(t, "nil", true, v.And(), nil)

	f := testAndTestEntry

	tt := []*testAndTable{
		f(PASS, "no-args", []bool{}, []int{}),
		f(FAIL, "false", []bool{false}, []int{1}),
		f(PASS, "true", []bool{true}, []int{1}),
		f(FAIL, "false false", []bool{false, false}, []int{1, 0}),
		f(FAIL, "false true", []bool{false, true}, []int{1, 0}),
		f(FAIL, "true false", []bool{true, false}, []int{1, 1}),
		f(PASS, "true true", []bool{true, true}, []int{1, 1}),
		f(FAIL, "true true false", []bool{true, true, false}, []int{1, 1, 1}),
	}

	for _, e := range tt {
		test(t, e.name, e.pass, e.validator, nil)
		for i, expect := range e.expectCc {
			require.Equal(t, expect, e.actualCc[i], fmt.Sprintf("%s -> and[%d] -> callcount", e.name, i))
		}
	}
}

type testAndTable struct {
	name      string
	pass      bool
	expectCc  []int
	actualCc  []int
	validator v.Validator
}

func testAndForcedValidator(b bool, count *int) v.Validator {
	return v.Function(func(data interface{}) (path string, err error) {
		*count++
		if !b {
			return "forcedValidator", fmt.Errorf("forced failure")
		} else {
			return "", nil
		}
	})
}

func testAndTestEntry(pass bool, name string, forcedResults []bool, andArgCallCounts []int) *testAndTable {
	t := &testAndTable{
		name, pass, andArgCallCounts, make([]int, len(forcedResults)), nil,
	}
	var vals []v.Validator
	for i, e := range forcedResults {
		vals = append(vals, v.And(testAndForcedValidator(e, &t.actualCc[i])))
	}
	t.validator = v.And(vals...)
	return t
}
