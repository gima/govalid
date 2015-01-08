package jsonv2_test

import (
	"fmt"
	j "github.com/gima/jsonv2"
	"reflect"
	"strings"
	"testing"
)

func TestFunction(t *testing.T) {
	counter, countingValidator := createCountingValidator()

	test(t, "combination1", false, j.Function(dogeValidator), "cat")
	test(t, "combination2", true, j.Function(dogeValidator), "cate")
	test(t, "combination3", true, j.Function(countingValidator), "doge")
	test(t, "combination4", true, j.Function(countingValidator), "doge")

	if *counter != 2 {
		t.Fatalf("counting validator count should be 2, is %d", *counter)
	}
}

func dogeValidator(data interface{}) (path string, err error) {
	s, ok := data.(string)
	if !ok {
		return "doge-validator", fmt.Errorf("expected string, got %s", reflect.TypeOf(data))
	}

	if !strings.HasSuffix(strings.ToLower(s), "e") {
		return "doge-validator", fmt.Errorf("expected string to end in 'e'")
	}

	return "", nil
}

func createCountingValidator() (counter *int, _ j.ValidatorFunc) {
	counter = new(int)
	return counter, func(data interface{}) (path string, err error) {
		*counter++
		return "", nil
	}
}
