package pqtype

import (
	"encoding/binary"
)

type Int4 int32

func (v *Int4) DecodeBinary(src []byte) ([]byte, error) {
	if len(src) < 12 {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != Int4OID {
		return nil, &DecodeTypeErr{expected: Int4OID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[4:])) == -1 {
		return nil, ErrNullValue
	}

	*v = Int4(binary.BigEndian.Uint32(src[8:]))

	return src[12:], nil
}



func Int4Null(src []byte) (*Int4, []byte, error) {
	if len(src) < 8 {
		return nil, nil, ErrInsufficientBytes
	}

	if int32(binary.BigEndian.Uint32(src[4:])) == -1 {
		return nil, src[8:], nil
	}

	i := new(Int4)
	var err error
	src, err = i.DecodeBinary(src)
	return i, src, err
}