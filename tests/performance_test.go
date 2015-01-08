package jsonv2_test

import (
	"encoding/json"
	"fmt"
	j "github.com/gima/jsonv2"
	"log"
	"reflect"
	"testing"
)

var (
	decoded interface{}
	schema  j.Validator
)

func init() {
	// set up raw json data
	rawJson := []byte(`
		{
    	"status": true,
      "data": {
      	"token": "CAFED00D",
	      "debug": 69306,
      	"items": [
      	  { "url": "https://news.ycombinator.com/", "comment": "clickbaits" },
          { "url": "http://golang.org/", "comment": "some darn gophers" },
          { "url": "http://www.kickstarter.com/", "comment": "opensource projects. yeah.." }
       	],
       	"ghost2": null,
       	"meta": {
       	 "g": 1,
         "xyzzy": 0.25,
         "blöö": 0.5
       	}
      }
		}`)

	// decode json
	if err := json.Unmarshal(rawJson, &decoded); err != nil {
		log.Panic("JSON parsing failed. Err =", err)
	}

	// set up a custom validator function
	myValidatorFunc := func(data interface{}) (path string, err error) {
		path = "myValidatorFunc"

		validate, ok := data.(string)
		if !ok {
			return path, fmt.Errorf("expected string, got %v", reflect.TypeOf(data))
		}

		if validate != "CAFED00D" {
			return path, fmt.Errorf("expected CAFED00D, got %s", validate)
		}

		return "", nil
	}

	// construct the schema which is used to validate data
	schema = j.Object(
		j.ObjKV("status", j.Boolean()),
		j.ObjKV("data", j.Object(
			j.ObjKV("token", j.Function(myValidatorFunc)),
			j.ObjKV("debug", j.Number(j.NumMin(1), j.NumMax(99999))),
			j.ObjKV("items", j.Array(j.ArrEach(j.Object(
				j.ObjKV("url", j.String(j.StrMin(1))),
				j.ObjKV("comment", j.Optional(j.String())),
			)))),
			j.ObjKV("ghost", j.Optional(j.String())),
			j.ObjKV("ghost2", j.Optional(j.String())),
			j.ObjKV("meta", j.Object(
				j.ObjKeys(j.String()),
				j.ObjValues(j.Or(j.Number(j.NumMin(.01), j.NumMax(1.1)), j.String())),
			)),
		)),
	)
}

func BenchmarkObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// do the actual validation
		if path, err := schema.Validate(decoded); err != nil {
			b.Fatalf("Benchmark N = %d. Failed (%s). Path: %s", b.N, err, path)
		}
	}
}
