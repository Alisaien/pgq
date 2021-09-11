package pgval

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type _int16 struct{}
type _int16Ptr struct{}

var Int16 = _int16{}
var Int16Ptr =_int16Ptr{}

func (_int16) Read(iter *pgetc.Iterator) int16 {
	if iter.Next4() != nil {
		return 0
	}

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		iter.ReportError(pgetc.ErrNull)
		return 0
	}

	if iter.Next(int(size)) != nil {
		return 0
	}

	return pgbin.Int16.Read(iter)
}

func (_int16Ptr) Read(iter *pgetc.Iterator) *int16 {
	if iter.Next4() != nil {
		return nil
	}

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		return nil
	}

	if iter.Next(int(size)) != nil {
		return nil
	}

	val := pgbin.Int16.Read(iter)
	return &val
}

func (_int16) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(2)
	pgbin.Int16.Write(ptr, stream)
}
