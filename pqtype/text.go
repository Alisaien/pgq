package pqtype

import "encoding/binary"

type Text string

const TextOID = 25

func (t *Text) FromBinary(src []byte) ([]byte, error) {
	if len(src) < ValueOffset {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != TextOID {
		return nil, &DecodeTypeErr{expected: TextOID, got: typ}
	}

	textSize := int32(binary.BigEndian.Uint32(src[SizeOffset:]))
	if textSize == -1 {
		return nil, ErrNullValue
	}

	buf := make([]byte, textSize)
	copy(buf, src[ValueOffset:])
	*t = Text(buf)

	return src[ValueOffset+textSize:], nil
}
