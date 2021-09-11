package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
	"unsafe"
)

const Int2OID  = 21

type _int16 struct{}
type _int16Ptr struct{}

var Int16 = _int16{}
var Int16Ptr =_int16Ptr{}

func (_int16) Read(iter *pgetc.Iterator) int16 {
	if iter.Next4() != nil {
		return 0
	}

	if binary.BigEndian.Uint32(iter.Read()) != Int2OID {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return 0
	}

	return pgval.Int16.Read(iter)
}

func (_int16Ptr) Read(iter *pgetc.Iterator) *int16 {
	if iter.Next4() != nil {
		return nil
	}

	if binary.BigEndian.Uint32(iter.Read()) != Int2OID {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return nil
	}

	return pgval.Int16Ptr.Read(iter)
}

func (_int16) ReadUnsafe(iter *pgetc.Iterator) unsafe.Pointer {
	return unsafe.Pointer(Int16Ptr.Read(iter))
}

func (_int16) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(Int2OID)
	pgval.Int16.Write(ptr, stream)
}
