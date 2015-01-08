package jsonv2_test

import (
	j "github.com/gima/jsonv2"
	"testing"
)

type maep map[interface{}]interface{}

func TestObject(t *testing.T) {
	var np *map[interface{}]interface{}

	test(t, "basic object", true, j.Object(), maep{})
	test(t, "nil", false, j.Object(), nil)
	test(t, "nil map ptr", false, j.Object(), np)
	test(t, "non-object", false, j.Object(), 3)

	testObjKeys(t)
	testObjValues(t)
	testObjKVs(t)
}

func testObjKeys(t *testing.T) {
	counter, countingValidator := createCountingValidator()

	sch := j.Object(
		j.ObjKeys(j.String()),
		j.ObjKeys(j.Function(countingValidator)),
	)
	m := maep{
		"a": nil,
		"b": 1,
		"c": true,
	}
	test(t, "only string objkeys", true, sch, m)
	if *counter != 3 {
		t.Fatalf("key counter should be 3, got %d", *counter)
	}

	m = maep{
		"a": nil,
		1:   1,
	}
	test(t, "!only string objkeys", false, sch, m)
}

func testObjValues(t *testing.T) {
	counter, countingValidator := createCountingValidator()

	sch := j.Object(
		j.ObjValues(j.String()),
		j.ObjValues(j.Function(countingValidator)),
	)
	m := maep{
		nil:  "1",
		1:    "b",
		true: "c",
	}
	test(t, "only string objvalues", true, sch, m)
	if *counter != 3 {
		t.Fatalf("value counter should be 3, got %d", *counter)
	}

	m = maep{
		nil: "1",
		1:   1,
	}
	test(t, "!only string objvalues", false, sch, m)
}

func testObjKVs(t *testing.T) {
	counter, countingValidator := createCountingValidator()

	sch := j.Object(
		j.ObjKV(nil, j.And(j.String(j.StrIs("1")), j.Function(countingValidator))),
		j.ObjKV("1", j.And(j.String(j.StrIs("b")), j.Function(countingValidator))),
		j.ObjKV(true, j.And(j.Number(j.NumIs(3)), j.Function(countingValidator))),
	)
	m := maep{
		nil:  "1",
		"1":  "b",
		true: 3,
	}
	test(t, "mixed objkvs", true, sch, m)
	if *counter != 3 {
		t.Fatalf("value counter should be 3, got %d", *counter)
	}

	m = maep{
		nil:  "1",
		"1":  2,
		true: 3,
	}
	test(t, "!mixed objkvs", false, sch, m)

	m = maep{
		nil:  "1",
		"1":  nil,
		true: 3,
	}
	test(t, "!mixed objkvs (nil)", false, sch, m)
}
