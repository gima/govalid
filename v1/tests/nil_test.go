package govalid_test

import (
	"testing"

	v "github.com/gima/govalid/v1"
)

func TestNil(t *testing.T) {
	test(t, "nil", true, v.Nil(), nil)
	test(t, "non-nil", false, v.Nil(), "")
}
