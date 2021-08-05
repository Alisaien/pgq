package pqtype

import "encoding/binary"

type Int2 int16

const (
	Int2OID  = 21
	int2Size = 2
)

func (v *Int2) FromBinary(src []byte) ([]byte, error) {
	const size = ValueOffset + int2Size

	if len(src) < size {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != Int2OID {
		return nil, &DecodeTypeErr{expected: Int4OID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[SizeOffset:])) == -1 {
		return nil, ErrNullValue
	}

	*v = Int2(binary.BigEndian.Uint16(src[ValueOffset:]))

	return src[size:], nil
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
	src, err = i.FromBinary(src)
	return i, src, err
}
