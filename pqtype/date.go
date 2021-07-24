package pqtype

import (
	"encoding/binary"
	"github.com/jackc/pgtype"
	"time"
)

type Date struct{ time.Time }

const (
	DateOID  = 1082
	dateSize = 4

	inftyDayOffset    = 2147483647
	negInftyDayOffset = -2147483648
)

func (v *Date) FromBinary(src []byte) ([]byte, error) {
	if len(src) < valueOffset+dateSize {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != DateOID {
		return nil, &DecodeTypeErr{expected: DateOID, got: typ}
	}

	return v.fromBinary(src[valueOffset:])
}

func (v *Date) fromBinary(src []byte) ([]byte, error) {
	dayOffset := int32(binary.BigEndian.Uint32(src))

	switch dayOffset {
	case inftyDayOffset:
		return nil, ErrInfinity
	case negInftyDayOffset:
		return nil, ErrInfinity
	default:
		v.Time = time.Date(2000, 1, int(1+dayOffset), 0, 0, 0, 0, time.UTC)
	}

	return src[dateSize:], nil
}

func (v *Date) DecodeBinary(_ *pgtype.ConnInfo, src []byte) error {
	if len(src) != dateSize {
		return ErrInvalidSrcLength
	}

	var err error
	_, err = v.fromBinary(src)
	return err
}
