package models

import "fmt"

// Error represents a logic specific error.
type Error struct {
	message string
}

// Error returns the error message as a string.
func (e *Error) Error() string {
	return e.message
}

// NewError wraps a message in an Error object.
func NewError(m string) *Error {
	return &Error{
		message: m,
	}
}

// NewErrorf wraps a message in an Error object.
func NewErrorf(m string, args ...interface{}) *Error {
	return &Error{
		message: fmt.Sprintf(m, args...),
	}
}
