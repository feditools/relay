package token

import "errors"

var (
	// ErrInvalidLength is returned when a token's data is an invalid length.
	ErrInvalidLength = errors.New("invalid length")
	// ErrSaltEmpty is returned when a token's data is an invalid length.
	ErrSaltEmpty = errors.New("salt empty")
)
