package pqtype

import (
	"encoding/binary"
	"time"
)

type Timestamptz time.Time

const (
	TimestamptzOID  = 1184
	timestamptzSize = 8

	microSecFromUnixEpochToY2K = 946684800 * 1000000
	inftyMicroSecOffset        = 9223372036854775807
	negInftyMicroSecOffset     = -9223372036854775808
)

func (v *Timestamptz) FromBinary(src []byte) ([]byte, error) {
	const size = valueOffset + timestamptzSize
	if len(src) < size {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != TimestamptzOID {
		return nil, &DecodeTypeErr{expected: TimestamptzOID, got: typ}
	}

	microsecSinceY2K := int64(binary.BigEndian.Uint64(src[valueOffset:]))
	switch microsecSinceY2K {
	case inftyMicroSecOffset:
		return nil, ErrInfinity
	case negInftyMicroSecOffset:
		return nil, ErrInfinity
	default:
		microSecSinceUnixEpoch := microSecFromUnixEpochToY2K + microsecSinceY2K
		*v = Timestamptz(time.Unix(microSecSinceUnixEpoch/1000000, (microSecSinceUnixEpoch%1000000)*1000))
	}

	return src[size:], nil
}
