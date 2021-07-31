package pqtype

import (
	"encoding/binary"
)

type Int4 int32

const (
	Int4OID  = 23
	int4Size = 4
)

func (v *Int4) FromBinary(src []byte) ([]byte, error) {
	if len(src) < valueOffset + int4Size {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != Int4OID {
		return nil, &DecodeTypeErr{expected: Int4OID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[sizeOffset:])) == -1 {
		return nil, ErrNullValue
	}

	return v.fromBinary(src[valueOffset:])
}

func (v *Int4) fromBinary(src []byte) ([]byte, error) {
	*v = Int4(binary.BigEndian.Uint32(src))
	return src[int4Size:], nil
}

func Int4Null(src []byte) (*Int4, []byte, error) {
	if len(src) < 8 {
		return nil, nil, ErrInsufficientBytes
	}

	if int32(binary.BigEndian.Uint32(src[sizeOffset:])) == -1 {
		return nil, src[valueOffset:], nil
	}

	i := new(Int4)
	var err error
	src, err = i.FromBinary(src)
	return i, src, err
}
