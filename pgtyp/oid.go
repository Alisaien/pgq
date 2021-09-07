package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/internal"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
	"unsafe"
)

const OIDOID  = 26

type (
	_oid struct{}
)

var (
	OID _oid
)

func (_oid) Read(iter *internal.Iterator) pgetc.OID {
	if iter.Next4() != nil {
		return 0
	}

	if binary.BigEndian.Uint32(iter.Read()) != OIDOID {
		iter.Error(pgetc.ErrUnexpectedType)
		return 0
	}

	return pgval.OID.Read(iter)
}

func (_oid) Write(ptr unsafe.Pointer, stream *internal.Stream) {
	stream.WriteUint32(OIDOID)
	pgval.OID.Write(ptr, stream)
}