package pqtype

import "encoding/binary"

const compositeTypeHeaderSize = 4

func DecodeCompositeTypeHeader(src []byte) (uint32, []byte, error) {
	if len(src) < compositeTypeHeaderSize {
		return 0, nil, ErrInsufficientBytes
	}

	return binary.BigEndian.Uint32(src), src[compositeTypeHeaderSize:], nil
}
