package pgval

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/internal"
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type _bool struct{}
type _boolPtr struct{}

var Bool = _bool{}
var BoolPtr =_boolPtr{}

func (_bool) Read(iter *internal.Iterator) bool {
	if iter.Next4() != nil {
		return false
	}

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		iter.Error(pgetc.ErrNull)
		return false
	}

	if iter.Next(int(size)) != nil {
		return false
	}

	return pgbin.Bool.Read(iter)
}

func (_boolPtr) Read(iter *internal.Iterator) *bool {
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

	val := pgbin.Bool.Read(iter)
	return &val
}

func (_bool) Write(ptr unsafe.Pointer, stream *internal.Stream) {
	stream.WriteUint32(1)
	pgbin.Bool.Write(ptr, stream)
}
