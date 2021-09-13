package pgetc

import "errors"

var ErrInfinity = errors.New("infinity is not a valid value")

var ErrEOF = errors.New("insufficient bytes left in src")

var ErrInvalidSrcLength = errors.New("invalid valid length for type")

var ErrNull = errors.New("expected not null value")

var ErrNumFieldMismatch = errors.New("unexpected number of fields for composite type")

var ErrTooManyDims = errors.New("too many dims in array")

var ErrUnexpectedType = errors.New("unexpected type")

var ErrUnknownVersion = errors.New("unknown version")