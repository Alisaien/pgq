package pqtype

import (
	"encoding/binary"
)

type TypeDecoder interface {
	DecodeType(src []byte) ([]byte, error)
}

type ValueDecoder interface {
	DecodeValue(src []byte) ([]byte, error)
}

type Reader interface {
	Read(src []byte) ([]byte, error)
}

func LenCheck(src []byte, size int) error {
	if len(src) < valueHeaderSize+size {
		return ErrInsufficientBytes
	}
	return nil
}

func TypeCheck(src []byte, oid OID) ([]byte, error) {
	var typ OID
	typ.Read(src)
	if typ != oid {
		return nil, &DecodeTypeErr{expected: oid, got: typ}
	}

	return src[oidSize:], nil
}

func ValueSize(src []byte) (int32, []byte) {
	return int32(binary.BigEndian.Uint32(src)), src[4:]
}
