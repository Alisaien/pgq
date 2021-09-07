package pgval

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/internal"
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

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

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		iter.Error(pgetc.ErrNull)
		return 0
	}

	if iter.Next(int(size)) != nil {
		return 0
	}

	return pgbin.OID.Read(iter)
}

func (_oid) Write(ptr unsafe.Pointer, stream *internal.Stream) {
	stream.WriteUint32(4)
	pgbin.OID.Write(ptr, stream)
}
