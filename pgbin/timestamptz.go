package pgbin

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"time"
)

type _timestamptz struct{}

var Timestamptz _timestamptz

const (
	microSecFromUnixEpochToY2K = 946684800 * 1000000
	inftyMicroSecOffset        = 9223372036854775807
	negInftyMicroSecOffset     = -9223372036854775808
)

func (_timestamptz) Read(iter *pgetc.Iterator) time.Time {
	microsecSinceY2K := int64(binary.BigEndian.Uint64(iter.Read()))
	switch microsecSinceY2K {
	case inftyMicroSecOffset, negInftyMicroSecOffset:
		iter.Error(pgetc.ErrInfinity)
		return time.Time{}
	default:
		microSecSinceUnixEpoch := microSecFromUnixEpochToY2K + microsecSinceY2K
		return time.Unix(microSecSinceUnixEpoch/1000000, (microSecSinceUnixEpoch%1000000)*1000)
	}
}
