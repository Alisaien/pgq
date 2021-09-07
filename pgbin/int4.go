package pgbin

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/internal"
	"unsafe"
)

type _int struct{}
type _int32 struct{}
type _uint32 struct{}

var Int = _int{}
var Int32 = _int32{}
var Uint32 = _uint32{}

func (_int) Read(iter *internal.Iterator) int {
	return int(binary.BigEndian.Uint32(iter.Read()))
}

func (_int32) Read(iter *internal.Iterator) int32 {
	return int32(binary.BigEndian.Uint32(iter.Read()))
}

func (_uint32) Read(iter *internal.Iterator) uint32 {
	return binary.BigEndian.Uint32(iter.Read())
}

func (_int) Write(ptr unsafe.Pointer, stream *internal.Stream) {
	stream.WriteUint32(*(*uint32)(ptr))
}

func (_int32) Write(ptr unsafe.Pointer, stream *internal.Stream) {
	stream.WriteUint32(*(*uint32)(ptr))
}

func (_uint32) Write(ptr unsafe.Pointer, stream *internal.Stream) {
	stream.WriteUint32(*(*uint32)(ptr))
}
