package pqtype

import "encoding/binary"

type Bool bool

const (
	BoolOID  = 16
	boolSize = 1
)

func (v *Bool) FromBinary(src []byte) ([]byte, error) {
	if len(src) < valueOffset + boolSize {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != BoolOID {
		return nil, &DecodeTypeErr{expected: BoolOID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[sizeOffset:])) == -1 {
		return nil, ErrNullValue
	}

	return v.fromBinary(src[valueOffset:])
}

func (v *Bool) fromBinary(src []byte) ([]byte, error) {
	*v = src[0] == 1
	return src[boolSize:], nil
}