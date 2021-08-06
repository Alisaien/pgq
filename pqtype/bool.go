package pqtype

import (
	"github.com/jackc/pgio"
)

type Bool bool

const (
	BoolOID  = 16
	boolSize = 1
)

func (v *Bool) DecodeType(src []byte) ([]byte, error) {
	err := LenCheck(src, boolSize)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, BoolOID)
	if err != nil {
		return nil, err
	}

	return v.DecodeValue(src)
}

func (v *Bool) DecodeValue(src []byte) ([]byte, error) {
	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	return v.Read(src)
}

func (v *Bool) Read(src []byte) ([]byte, error) {
	*v = src[0] == 1
	return src[boolSize:], nil
}

func (v Bool) EncodeType(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, BoolOID)
	return v.EncodeValue(buf)
}

func (v Bool) EncodeValue(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, boolSize)
	return v.Write(buf)
}

func (v Bool) Write(buf []byte) []byte {
	if v {
		return append(buf, 1)
	} else {
		return append(buf, 0)
	}
}
