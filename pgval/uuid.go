package pgval

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type _uuid struct{}
type _uuidPtr struct{}

var UUID = _uuid{}
var UUIDPtr =_uuidPtr{}

func (_uuid) Read(iter *pgetc.Iterator) pgetc.UUID {
	if iter.Next4() != nil {
		return pgetc.UUID{}
	}

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		iter.Error(pgetc.ErrNull)
		return pgetc.UUID{}
	}

	if iter.Next(int(size)) != nil {
		return pgetc.UUID{}
	}

	return pgbin.UUID.Read(iter)
}

func (_uuidPtr) Read(iter *pgetc.Iterator) *pgetc.UUID {
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

	val := pgbin.UUID.Read(iter)
	return &val
}

func (_uuid) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(1)
	pgbin.Bool.Write(ptr, stream)
}
