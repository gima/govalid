## Data validation library for `golang´

This library can be used to validate [most data types](https://godoc.org/github.com/gima/govalid/v1) supported by golang. Custom validators can be used where the supplied ones are not enough.

#### Documentation

[![godoc](https://img.shields.io/badge/godoc-reference-5976b1.svg?style=flat-square)](https://godoc.org/github.com/gima/govalid/v1) and (sorry, but for the time being) [tests](https://github.com/gima/govalid/tree/master/v1/tests)

#### Import

```go
import v "github.com/gima/govalid/v1"
```

#### Intro & Usage

Create validator:

```go
schema := v.Object(
	v.ObjKV("status", v.Boolean()),
	v.ObjKV("data", v.Object(
		v.ObjKV("token", v.Function(myValidatorFunc)),
		v.ObjKV("debug", v.Number(v.NumMin(1), v.NumMax(99999))),
		v.ObjKV("items", v.Array(v.ArrEach(v.Object(
			v.ObjKV("url", v.String(v.StrMin(1))),
			v.ObjKV("comment", v.Optional(v.String())),
		)))),
		v.ObjKV("ghost", v.Optional(v.String())),
		v.ObjKV("ghost2", v.Optional(v.String())),
		v.ObjKV("meta", v.Object(
			v.ObjKeys(v.String()),
			v.ObjValues(v.Or(v.Number(v.NumMin(.01), v.NumMax(1.1)), v.String())),
		)),
	)),
)
```

Validate some data using the schema:

```go
if path, err := schema.Validate(data); err == nil {
	t.Log("Validation passed.")
} else {
	t.Fatalf("Validation failed at %s. Error (%s)", path, err)
}
```

```go
// Example of failed validation:

// Validation failed at Object->Key[data].Value->Object->Key[debug].Value->Number.
// Error (expected (*)data convertible to float64, got bool)
```

## Greets
Idea loosely based on [js-schema](https://github.com/molnarg/js-schema), thank you.

## License

Public Domain (see UNLICENSE.txt). Mention of origin would be appreciated.

*go, golang*
