package pgval

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type _string struct{}
type _stringPtr struct{}

var String = _string{}
var StringPtr =_stringPtr{}

func (_string) Read(iter *pgetc.Iterator) string {
	if iter.Next4() != nil {
		return ""
	}

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		iter.ReportError(pgetc.ErrNull)
		return ""
	}

	if iter.Next(int(size)) != nil {
		return ""
	}

	return pgbin.String.Read(iter)
}

func (_stringPtr) Read(iter *pgetc.Iterator) *string {
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

	val := pgbin.String.Read(iter)
	return &val
}

func (_string) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(uint32(len(*(*string)(ptr))))
	pgbin.String.Write(ptr, stream)
}
