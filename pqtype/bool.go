package pqtype

import (
	"github.com/jackc/pgio"
)

type Bool bool

const (
	BoolOID  = 16
	boolSize = 1
)

func (v *Bool) FromBinary(src []byte) ([]byte, error) {
	err := LenCheck(src, boolSize)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, BoolOID)
	if err != nil {
		return nil, err
	}

	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	return v.FromPureBinary(src)
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
