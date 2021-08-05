package pqtype

import (
	"encoding/binary"
	"github.com/jackc/pgio"
)

type Bool bool

const (
	BoolOID  = 16
	boolSize = 1
)

func (v *Bool) FromBinary(src []byte) ([]byte, error) {
	if len(src) < ValueOffset+boolSize {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != BoolOID {
		return nil, &DecodeTypeErr{expected: BoolOID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[SizeOffset:])) == -1 {
		return nil, ErrNullValue
	}

	return v.FromPureBinary(src[ValueOffset:])
}

func (v *Bool) FromPureBinary(src []byte) ([]byte, error) {
	*v = src[0] == 1
	return src[boolSize:], nil
}

func (v Bool) ToBinary(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, BoolOID)
	buf = pgio.AppendUint32(buf, boolSize)
	return v.ToPureBinary(buf)
}

func (v Bool) ToPureBinary(buf []byte) []byte {
	if v {
		return append(buf, 1)
	} else {
		return append(buf, 0)
	}
}
