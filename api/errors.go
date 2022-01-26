package api

import (
	"errors"
	"fmt"
	"net/http"
)

// List of known errors that may be reused between different handlers.
var (
	ErrInvalidRequestMethod   = Error{Original: errors.New("invalid request method"), Code: http.StatusMethodNotAllowed}
	ErrBadInput               = Error{Original: errors.New("bad input"), Code: http.StatusBadRequest}
	ErrTooBigFile             = Error{Original: errors.New("file is too big"), Code: http.StatusBadRequest}
	ErrInvalidFileContentType = Error{Original: errors.New("invalid file content type"), Code: http.StatusBadRequest}
	ErrFileNotFound           = Error{Original: errors.New("file not found"), Code: http.StatusNotFound}
)

// Error extends an original error with related code.
type Error struct {
	Original error
	Code     int
}

// Error returns a string representation of an error.
func (e Error) Error() string {
	return fmt.Sprintf("err: %s; code: %d", e.Original.Error(), e.Code)
}
