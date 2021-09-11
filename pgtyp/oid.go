package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
)



type (
	_oid struct{}
)

var (
	OID _oid
)

func (_oid) Read(iter *pgetc.Iterator) pgetc.OID {
	if iter.Next4() != nil {
		return 0
	}

	if binary.BigEndian.Uint32(iter.Read()) != pgetc.OIDOID {
		iter.ReportError(pgetc.ErrUnexpectedType)
		return 0
	}

	return pgval.OID.Read(iter)
}
