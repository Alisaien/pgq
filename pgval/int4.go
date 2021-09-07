package pgval

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

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

	return pgbin.Int.Read(iter)
}

func (_intPtr) Read(iter *pgetc.Iterator) *int {
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

	val := pgbin.Int.Read(iter)
	return &val
}

func (_int32Ptr) Read(iter *pgetc.Iterator) *int32 {
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

	val := pgbin.Int32.Read(iter)
	return &val
}

func (_int) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(4)
	pgbin.Int.Write(ptr, stream)
}