package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
	"time"
)

const TimestamptzOID  = 1184

type _timestamptz struct{}

var Timestamptz _timestamptz

func (_timestamptz) Read(iter *pgetc.Iterator) time.Time {
	if iter.Next4() != nil {
		return time.Time{}
	}

	if binary.BigEndian.Uint32(iter.Read()) != TimestamptzOID {
		iter.Error(pgetc.ErrInvalidSrcLength)
		return time.Time{}
	}

	return pgval.Timestamptz.Read(iter)
}
