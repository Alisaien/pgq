package pgbin

import (
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type _bool struct{}

var Bool = _bool{}

func (_bool) Read(iter *pgetc.Iterator) bool {
	if iter.Err != nil {
		return false
	}
	return iter.ReadByte1() == 1
}

func (_bool) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	if *(*bool)(ptr) {
		stream.WriteByte1(1)
	} else {
		stream.WriteByte1(0)
	}
}
