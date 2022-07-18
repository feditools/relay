package lib

import "errors"

// ErrInvalidAccountFormat is returned when a federated account is in an invalid format.
var ErrInvalidAccountFormat = errors.New("invalid account format")
