package pqtype

import "encoding/binary"

type Decoder interface {
	FromBinary(src []byte) ([]byte, error)
	FromPureBinary(src []byte) ([]byte, error)
}

func LenCheck(src []byte, size int) error {
	if len(src) < valueHeaderSize+size {
		return ErrInsufficientBytes
	}
	return nil
}

func TypeCheck(src []byte, oid OID) ([]byte, error) {
	var typ OID
	src, _ = typ.FromPureBinary(src)
	if typ != oid {
		return nil, &DecodeTypeErr{expected: oid, got: typ}
	}

	return src, nil
}

func ValueSize(src []byte) (int32, []byte) {
	return int32(binary.BigEndian.Uint32(src)), src[4:]
}
