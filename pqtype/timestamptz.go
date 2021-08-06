package pqtype

import (
	"encoding/binary"
	"github.com/jackc/pgtype"
	"time"
)

type Timestamptz struct{ time.Time }

const (
	TimestamptzOID  = 1184
	timestamptzSize = 8

	microSecFromUnixEpochToY2K = 946684800 * 1000000
	inftyMicroSecOffset        = 9223372036854775807
	negInftyMicroSecOffset     = -9223372036854775808
)

func (v *Timestamptz) DecodeType(src []byte) ([]byte, error) {
	err := LenCheck(src, timestamptzSize)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, TimestamptzOID)
	if err != nil {
		return nil, err
	}

	return v.DecodeValue(src)
}

func (v *Timestamptz) DecodeValue(src []byte) ([]byte, error) {
	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	return v.Read(src)
}

func (v *Timestamptz) Read(src []byte) ([]byte, error) {
	microsecSinceY2K := int64(binary.BigEndian.Uint64(src))
	switch microsecSinceY2K {
	case inftyMicroSecOffset:
		return nil, ErrInfinity
	case negInftyMicroSecOffset:
		return nil, ErrInfinity
	default:
		microSecSinceUnixEpoch := microSecFromUnixEpochToY2K + microsecSinceY2K
		v.Time = time.Unix(microSecSinceUnixEpoch/1000000, (microSecSinceUnixEpoch%1000000)*1000)
	}

	return src[timestamptzSize:], nil
}

func (v *Timestamptz) DecodeBinary(_ *pgtype.ConnInfo, src []byte) error {
	if len(src) != timestamptzSize {
		return ErrInvalidSrcLength
	}

	_, err := v.Read(src)
	return err
}
