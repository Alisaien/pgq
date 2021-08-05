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
	if len(src) < valueOffset+oidSize {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != OIDOID {
		return nil, &DecodeTypeErr{expected: OIDOID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[sizeOffset:])) == -1 {
		return nil, ErrNullValue
	}

	return v.FromPureBinary(src[valueOffset:])
}

func (v *OID) FromPureBinary(src []byte) ([]byte, error) {
	*v = OID(binary.BigEndian.Uint32(src))
	return src[oidSize:], nil
}

func (v OID) ToPureBinary(buf []byte) []byte {
	return pgio.AppendUint32(buf, uint32(v))
}
