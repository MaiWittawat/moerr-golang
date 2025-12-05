// Package moerr provides a minimal error-construction utility.
//
// This package wraps the standard error creation patterns (errors.New and
// fmt.Errorf) in a small API to keep error creation consistent across a codebase.
//
// moerr.New      → create a simple static error message
// moerr.NewErrorf → create a formatted dynamic error message
//
// The purpose of this package is not to replace Go’s standard library but to
// provide a centralized and explicit abstraction for error instantiation,
// especially in projects that prefer a dedicated error factory for consistency.
package moerr

import "fmt"

type errorString struct {
	s string
}

func New(text string) error {
	return &errorString{
		s: text,
	}
}

func NewErrorf(format string, a ...any) error {
	return &errorString{
		s: fmt.Sprintf(format, a...),
	}
}

func (e *errorString) Error() string {
	return e.s
}
