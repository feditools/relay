package kv

import "fmt"

// Error represents a database specific error.
type Error error

var (
	// ErrNil is returned when the kv value is nil.
	ErrNil Error = fmt.Errorf("nil")
)

// EncryptionError is returned when a an encryption error occurs.
type EncryptionError struct {
	message string
}

// Error returns the error message as a string.
func (e *EncryptionError) Error() string {
	return fmt.Sprintf("encryption: %s", e.message)
}

// NewEncryptionError wraps a message in an EncryptionError object.
func NewEncryptionError(msg string) Error {
	return &EncryptionError{message: msg}
}
