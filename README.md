## Data validation library for `golang´

This library can be used to validate [most data types](https://godoc.org/github.com/gima/jsonv2) supported by golang. Custom validators can be used where the supplied ones are not enough.

*Besides the name, there's no connection to json.*

#### Documentation

[![godoc](https://img.shields.io/badge/godoc-reference-5976b1.svg?style=flat-square)](https://godoc.org/github.com/gima/jsonv2) and (sorry, but for the time being) [tests](https://github.com/gima/jsonv2/tree/master/tests)

#### Import

```go
import j "github.com/gima/jsonv2"
```

#### Intro & Usage

Create validator:

```go
schema := j.Object(
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
