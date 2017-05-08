package govalid_test

import (
	"github.com/gima/govalid/v1/internal"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var (
	s  = "asd"
	np *string
)

func TestReflectOrIndirect(t *testing.T) {
	_, err := internal.ReflectOrIndirect(nil)
	require.Error(t, err, "@nil")

	v, err := internal.ReflectOrIndirect(s)
	require.NoError(t, err, "@value")
	require.Equal(t, reflect.String, v.Kind(), "value kind")
	require.Equal(t, s, v.Interface(), "value")

	v, err = internal.ReflectOrIndirect(&s)
	require.NoError(t, err, "@ptr")
	require.Equal(t, reflect.String, v.Kind(), "ptr kind")
	require.Equal(t, s, v.Interface(), "ptr")

	_, err = internal.ReflectOrIndirect(np)
	require.Error(t, err, "@nil ptr")
}

func TestReflectPtrOrFabricate(t *testing.T) {
	_, err := internal.ReflectPtrOrFabricate(nil)
	require.Error(t, err, "@nil")

	v, err := internal.ReflectPtrOrFabricate(s)
	require.NoError(t, err, "@value")
	require.Equal(t, reflect.Ptr, v.Kind(), "value kind: "+v.Kind().String())
	require.Equal(t, s, v.Elem().Interface(), "value")

	v, err = internal.ReflectPtrOrFabricate(np)
	require.Error(t, err, "@nil ptr")
}

func TestDigValue(t *testing.T) {
	func() {
		const id = "nil"
		v := internal.DigValue(nil)
		require.NotNil(t, v, id)
		require.False(t, v.IsValid(), id+" -> isvalid")
	}()

	func() {
		const id = "string"
		var s string = "asd"
		v := internal.DigValue(s)
		require.NotNil(t, v, id)
		require.True(t, v.IsValid(), id+" -> isvalid")
		require.Equal(t, reflect.String, v.Kind(), id+" -> kind==string")
		require.Equal(t, s, v.Interface(), id+" -> value==given")
	}()

	func() {
		const id = "nil *string"
		var ps *string = nil
		v := internal.DigValue(ps)
		require.NotNil(t, v, id)
		require.False(t, v.IsValid(), id+" -> isvalid")
	}()

	func() {
		const id = "non-nil **string to nil *string"
		var (
			ps  *string  = nil
			pps **string = &ps
		)
		v := internal.DigValue(pps)
		require.NotNil(t, v, id)
		require.True(t, v.IsValid(), id+" -> isvalid")
		require.Equal(t, reflect.Ptr, v.Type().Kind(), id+" -> type -> kind==ptr")
		require.Equal(t, reflect.String, v.Type().Elem().Kind(), id+" -> type -> kind==string")
		require.False(t, v.Elem().IsValid(), id+" -> elem -> isvalid")
	}()

	func() {
		const id = "non-nil **string to non-nil *string"
		var (
			s            = "asd"
			ps  *string  = &s
			pps **string = &ps
		)
		v := internal.DigValue(pps)
		require.NotNil(t, v, id)
		require.True(t, v.IsValid(), id+" -> isvalid")
		require.Equal(t, reflect.Ptr, v.Type().Kind(), id+" -> type -> kind==ptr")
		require.Equal(t, reflect.String, v.Type().Elem().Kind(), id+" -> type -> kind==string")
		require.True(t, v.Elem().IsValid(), id+" -> elem -> isvalid")
		require.Equal(t, s, v.Elem().Interface(), id+" -> elem -> value==given")
	}()
}
