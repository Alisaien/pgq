package pqtype

import "encoding/binary"

type Text string

const TextOID = 25

func (t *Text) DecodeBinary(src []byte) ([]byte, error) {
	if len(src) < valueOffset {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != TextOID {
		return nil, &DecodeTypeErr{expected: TextOID, got: typ}
	}

	textSize := int32(binary.BigEndian.Uint32(src[sizeOffset:]))
	if textSize == -1 {
		return nil, ErrNullValue
	}

	buf := make([]byte, textSize)
	copy(buf, src[valueOffset:])
	*t = Text(buf)

	return src[valueOffset+textSize:], nil
}