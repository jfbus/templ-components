// Package text renders anything (numbers, errors, ...) as text.
// It calls:
// - a strconv converter if possible
// - String() for types implementing fmt.Stringer()
// - Errors() for errors
//
// Usage:
//
//	@text.C(1) // display an integer
//
//	@text.C(err) // display an error
package text
