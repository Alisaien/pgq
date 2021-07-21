package pqtype

import "encoding/binary"

type Int2 int16

func (v *Int2) DecodeBinary(src []byte) ([]byte, error) {
	if len(src) < 10 {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != Int4OID {
		return nil, &DecodeTypeErr{expected: Int4OID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[4:])) == -1 {
		return nil, ErrNullValue
	}

	*v = Int2(binary.BigEndian.Uint16(src[8:]))

	return src[12:], nil
}

func Int2Null(src []byte) (*Int2, []byte, error) {
	if len(src) < 8 {
		return nil, nil, ErrInsufficientBytes
	}

	if int32(binary.BigEndian.Uint32(src[4:])) == -1 {
		return nil, src[8:], nil
	}

	i := new(Int2)
	var err error
	src, err = i.DecodeBinary(src)
	return i, src, err
}