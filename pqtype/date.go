package pqtype

import (
	"encoding/binary"
	"time"
)

type Date struct {
	time.Time
	infty int8
}

const (
	DateOID  = 1082
	dateSize = 4

	inftyDayOffset    = 2147483647
	negInftyDayOffset = -2147483648
)

func (v *Date) DecodeBinary(src []byte) ([]byte, error) {
	const size = valueOffset + dateSize
	if len(src) < size {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != DateOID {
		return nil, &DecodeTypeErr{expected: DateOID, got: typ}
	}

	dayOffset := int32(binary.BigEndian.Uint32(src))

	switch dayOffset {
	case inftyDayOffset:
		v.infty = 1
	case negInftyDayOffset:
		v.infty = -1
	default:
		v.Time = time.Date(2000, 1, int(1+dayOffset), 0, 0, 0, 0, time.UTC)
		v.infty = 0
	}

	return src[size:], nil
}
