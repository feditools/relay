package db

import "fmt"

// Error represents a database specific error
type Error error

var (
	// ErrGenID is returned when creating a new ID can't be generated for a new model.
	ErrGenID Error = fmt.Errorf("can't generate id")
	// ErrNoEntries is returned when a caller expected an entry for a query, but none was found.
	ErrNoEntries Error = fmt.Errorf("no entries")
	// ErrMultipleEntries is returned when a caller expected ONE entry for a query, but multiples were found.
	ErrMultipleEntries Error = fmt.Errorf("multiple entries")
	// ErrUnknown denotes an unknown database error.
	ErrUnknown Error = fmt.Errorf("unknown error")
)

// ErrAlreadyExists is returned when a caller tries to insert a database entry that already exists in the db.
type ErrAlreadyExists struct {
	message string
}

// Error returns the error message as a string.
func (e *ErrAlreadyExists) Error() string {
	return e.message
}

// NewErrAlreadyExists wraps a message in an ErrAlreadyExists object.
func NewErrAlreadyExists(msg string) error {
	return &ErrAlreadyExists{message: msg}
}
