package pgbin

import (
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type (
	_oid struct{}
)

var (
	OID _oid
)

func (_oid) Read(iter *pgetc.Iterator) pgetc.OID {
	return pgetc.OID(Uint32.Read(iter))
}

func (_oid) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(*(*uint32)(ptr))
}
