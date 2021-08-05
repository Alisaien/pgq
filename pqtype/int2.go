package pqtype

import "encoding/binary"

type Int2 int16

const (
	Int2OID  = 21
	int2Size = 2
)

func (v *Int2) FromBinary(src []byte) ([]byte, error) {
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

func (v *Int2) FromPureBinary(src []byte) ([]byte, error) {
	*v = Int2(binary.BigEndian.Uint16(src))
	return src[int2Size:], nil
}

func Int2Null(src []byte) (*Int2, []byte, error) {
	err := LenCheck(src, 0)
	if err != nil {
		return nil, nil, err
	}

	src, err = TypeCheck(src, Int2OID)
	if err != nil {
		return nil, nil, err
	}

	size, src := ValueSize(src)
	if size == -1 {
		return nil, src, nil
	}

	i := new(Int2)
	src, err = i.FromPureBinary(src)
	return i, src, err
}
