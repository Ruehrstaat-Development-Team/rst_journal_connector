package errors

import "errors"

var (
	// ErrMissingField is returned when a required field is missing
	ErrMissingField = errors.New("missing field")
)
