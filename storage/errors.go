package storage

import (
	"errors"
	"fmt"
)

var (
	ErrKeyNotFound = errors.New("gogo key not found")
)

type ErrUnexpected struct {
	Cause error
}

func (e *ErrUnexpected) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Unexpected Error. Cause: %e", e.Cause)
	} else {
		return "Unexpected Error"
	}
}

func NewErrUnexpected(cause error) error {
	return &ErrUnexpected{cause}
}
