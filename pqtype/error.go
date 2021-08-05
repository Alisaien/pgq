package pqtype

import (
	"errors"
	"fmt"
)

type DecodeTypeErr struct {
	expected OID
	got      OID
}

func (e *DecodeTypeErr) Error() string {
	return fmt.Sprintf("expected type %d, got %d", e.expected, e.got)
}

var ErrInfinity = errors.New("infinity is not a valid value")

var ErrInsufficientBytes = errors.New("insufficient bytes left in src")

var ErrInvalidSrcLength = errors.New("invalid valid length for type")

var ErrNullValue = errors.New("expected not null value")

var ErrNumFieldMismatch = errors.New("unexpected number of fields for composite type")

var ErrTooManyDims = errors.New("too many dims in array")
