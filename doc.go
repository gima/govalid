/*
Data validation library. For introduction, visit https://github.com/gima/jsonv2/blob/master/README.md

Data type verification

The supplied Validators verify data's type implicitly; they don't need extra
parameters during their construction to do so.
	// This means, for example, that the following String validation code
	// verifies that "asd" is of string type, even though given no parameters.
	path, err := jsonv2.Strings().Validate("asd")

Supported data types

The supplied validators understand the following data types:
	bool, array, map, ptr, slice, string and any numerical type except complex
This leaves out
	complex, chan, func, interface, struct and unsafepointer

(Don't be fooled by missing interface{} support, as interface{} itself is not
passed around. Instead other data types are wrapped inside an interface{}.)

Regarding pointers

It is recommended to pass pointers to the Validators, as this avoids making
copies of data and thus avoids unnecessary garbage collection.
Make no mistake, non-pointers work perfectly fine as well.
*/
package jsonv2
