package pqtype

import (
	"encoding/binary"
	"time"
)

type Timestamptz struct {
	time.Time
	infty int8
}

const (
	TimestamptzOID  = 1184
	timestamptzSize = 8

	microSecFromUnixEpochToY2K = 946684800 * 1000000
	inftyMicroSecOffset        = 9223372036854775807
	negInftyMicroSecOffset     = -9223372036854775808
)

func (v *Timestamptz) DecodeBinary(src []byte) ([]byte, error) {
	const size = valueOffset + timestamptzSize
	if len(src) < size {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != TimestamptzOID {
		return nil, &DecodeTypeErr{expected: TimestamptzOID, got: typ}
	}

	microsecSinceY2K := int64(binary.BigEndian.Uint64(src))
	switch microsecSinceY2K {
	case inftyMicroSecOffset:
		v.infty = 1
	case negInftyMicroSecOffset:
		v.infty = -1
	default:
		microsecSinceUnixEpoch := microSecFromUnixEpochToY2K + microsecSinceY2K
		v.Time = time.Unix(microsecSinceUnixEpoch/1000000, (microsecSinceUnixEpoch%1000000)*1000)
		v.infty = 0
	}

	return src[size:], nil
}
