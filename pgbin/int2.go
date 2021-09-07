package pgbin

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type _int16 struct{}

var Int16 = _int16{}

func (_int16) Read(iter *pgetc.Iterator) int16 {
	return int16(binary.BigEndian.Uint16(iter.Read()))
}

func (_int16) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint16(*(*uint16)(ptr))
}
