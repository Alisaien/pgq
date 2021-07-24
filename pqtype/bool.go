package pqtype

import "encoding/binary"

type Bool bool

const (
	BoolOID  = 16
	boolSize = 1
)

func (v *Bool) FromBinary(src []byte) ([]byte, error) {
	const size = valueOffset + boolSize

	if len(src) < size {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != BoolOID {
		return nil, &DecodeTypeErr{expected: BoolOID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[sizeOffset:])) == -1 {
		return nil, ErrNullValue
	}

	*v = src[valueOffset+1] == 1

	return src[size:], nil
}
