package pgval

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"time"
)

type _timestamptz struct{}

var Timestamptz _timestamptz

func (_timestamptz) Read(iter *pgetc.Iterator) time.Time {
	if iter.Next(4) != nil {
		return time.Time{}
	}

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		iter.ReportError(pgetc.ErrNull)
		return time.Time{}
	}

	if iter.Next(int(size)) != nil {
		return time.Time{}
	}

	return pgbin.Timestamptz.Read(iter)
}
