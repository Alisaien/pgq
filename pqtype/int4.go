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

func (v *Int4) FromBinary(src []byte) ([]byte, error) {
	err := LenCheck(src, int4Size)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, Int4OID)
	if err != nil {
		return nil, err
	}

	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	return v.FromPureBinary(src)
}

func (v *Int4) FromPureBinary(src []byte) ([]byte, error) {
	*v = Int4(binary.BigEndian.Uint32(src))
	return src[int4Size:], nil
}

func (v Int4) ToPureBinary(buf []byte) []byte {
	return pgio.AppendUint32(buf, uint32(v))
}

func Int4FromBinary(src []byte) (*Int4, []byte, error) {
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

	i := new(Int4)
	src, err = i.FromPureBinary(src)
	return i, src, err
}
