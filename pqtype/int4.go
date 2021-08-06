package pqtype

import (
	"encoding/binary"
	"github.com/jackc/pgio"
)

type Int4 int32

const (
	Int4OID  = 23
	int4Size = 4
)

func (v *Int4) DecodeType(src []byte) ([]byte, error) {
	err := LenCheck(src, int4Size)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, Int4OID)
	if err != nil {
		return nil, err
	}

	return v.DecodeValue(src)
}

func (v *Int4) DecodeValue(src []byte) ([]byte, error) {
	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	return v.Read(src)
}

func (v *Int4) Read(src []byte) ([]byte, error) {
	*v = Int4(binary.BigEndian.Uint32(src))
	return src[int4Size:], nil
}

func (v Int4) EncodeType(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, Int4OID)
	return v.EncodeValue(buf)
}

func (v Int4) EncodeValue(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, int4Size)
	return v.Write(buf)
}

func (v Int4) Write(buf []byte) []byte {
	return pgio.AppendUint32(buf, uint32(v))
}

func DecodeInt4(src []byte) (*Int4, []byte, error) {
	err := LenCheck(src, 0)
	if err != nil {
		return nil, nil, err
	}

	src, err = TypeCheck(src, Int4OID)
	if err != nil {
		return nil, nil, err
	}

	size, src := ValueSize(src)
	if size == -1 {
		return nil, src, nil
	}

	v := new(Int4)
	src, _ = v.Read(src)
	return v, src, err
}
