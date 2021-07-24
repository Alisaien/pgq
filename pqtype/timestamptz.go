package pqtype

import (
	"encoding/binary"
	"github.com/jackc/pgtype"
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

func (v *Timestamptz) FromBinary(src []byte) ([]byte, error) {
	if len(src) < valueOffset + timestamptzSize {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != TimestamptzOID {
		return nil, &DecodeTypeErr{expected: TimestamptzOID, got: typ}
	}

	return v.fromBinary(src[valueOffset:])
}

func (v *Timestamptz) fromBinary(src []byte) ([]byte, error) {
	microsecSinceY2K := int64(binary.BigEndian.Uint64(src))
	switch microsecSinceY2K {
	case inftyMicroSecOffset:
		v.infty = 1
	case negInftyMicroSecOffset:
		v.infty = -1
	default:
		microSecSinceUnixEpoch := microSecFromUnixEpochToY2K + microsecSinceY2K
		v.Time = time.Unix(microSecSinceUnixEpoch/1000000, (microSecSinceUnixEpoch%1000000)*1000)
		v.infty = 0
	}

	return src[timestamptzSize:], nil
}

func (v *Timestamptz) DecodeBinary(_ *pgtype.ConnInfo, src[]byte) error {
	if len(src) < valueOffset + timestamptzSize {
		return ErrInsufficientBytes
	}

	var err error
	_, err = v.fromBinary(src)
	return err
}
