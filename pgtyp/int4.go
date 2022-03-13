package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
	"unsafe"
)

const Int4OID = 23

type _int struct{}
type _int32 struct{}
type _uint32 struct{}

type _intPtr struct{}
type _int32Ptr struct{}

var Int = _int{}
var Int32 = _int32{}
var Uint32 = _uint32{}

var IntPtr = _intPtr{}
var Int32Ptr = _int32Ptr{}

func (_int) Read(iter *pgetc.Iterator) int {
	if iter.Next(4) != nil {
		return 0
	}

	if binary.BigEndian.Uint32(iter.Read()) != Int4OID {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return 0
	}

	return pgval.Int.Read(iter)
}

func (_intPtr) Read(iter *pgetc.Iterator) *int {
	if iter.Next(4) != nil {
		return nil
	}

	if binary.BigEndian.Uint32(iter.Read()) != Int4OID {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return nil
	}

	return pgval.IntPtr.Read(iter)
}

func (_int32Ptr) Read(iter *pgetc.Iterator) *int32 {
	if iter.Next(4) != nil {
		return nil
	}

	if binary.BigEndian.Uint32(iter.Read()) != Int4OID {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return nil
	}

	return pgval.Int32Ptr.Read(iter)
}

func (_uint32) Read(iter *pgetc.Iterator) uint32 {
	if iter.Next(4) != nil {
		return 0
	}

	if binary.BigEndian.Uint32(iter.Read()) != Int4OID {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return 0
	}

	return pgval.Uint32.Read(iter)
}

func (_int) ReadUnsafe(iter *pgetc.Iterator) unsafe.Pointer {
	return unsafe.Pointer(IntPtr.Read(iter))
}

func (_int32) ReadUnsafe(iter *pgetc.Iterator) unsafe.Pointer {
	return unsafe.Pointer(Int32Ptr.Read(iter))
}

func (_int) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(Int4OID)
	pgval.Int.Write(ptr, stream)
}
