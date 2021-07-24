package pqtype

import (
	"errors"
	"fmt"
)

type DecodeTypeErr struct {
	expected int32
	got      int32
}

func (e *DecodeTypeErr) Error() string {
	return fmt.Sprintf("expected type %d, got %d", e.expected, e.got)
}

var ErrInsufficientBytes = errors.New("insufficient bytes left in src")

var ErrInvalidSrcLength = errors.New("invalid valid length for type")

var ErrInfinity = errors.New("infinity is not a valid value")

var ErrNullValue = errors.New("expected not null value")
