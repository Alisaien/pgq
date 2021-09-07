package pgbin

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type _int struct{}
type _int32 struct{}
type _uint32 struct{}

var Int = _int{}
var Int32 = _int32{}
var Uint32 = _uint32{}

func (_int) Read(iter *pgetc.Iterator) int {
	return int(binary.BigEndian.Uint32(iter.Read()))
}

func (_int32) Read(iter *pgetc.Iterator) int32 {
	return int32(binary.BigEndian.Uint32(iter.Read()))
}

func (_uint32) Read(iter *pgetc.Iterator) uint32 {
	return binary.BigEndian.Uint32(iter.Read())
}

func (_int) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(*(*uint32)(ptr))
}

func (_int32) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(*(*uint32)(ptr))
}

func (_uint32) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(*(*uint32)(ptr))
}
