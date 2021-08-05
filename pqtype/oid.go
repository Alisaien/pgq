package pqtype

import (
	"encoding/binary"
	"github.com/jackc/pgio"
)

const (
	OIDOID  = 26
	oidSize = 4
)

type OID uint32

func (v *OID) FromBinary(src []byte) ([]byte, error) {
	err := LenCheck(src, oidSize)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, OIDOID)
	if err != nil {
		return nil, err
	}

	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	return v.FromPureBinary(src)
}

func (v *OID) FromPureBinary(src []byte) ([]byte, error) {
	*v = OID(binary.BigEndian.Uint32(src))
	return src[oidSize:], nil
}

func (v OID) ToPureBinary(buf []byte) []byte {
	return pgio.AppendUint32(buf, uint32(v))
}
