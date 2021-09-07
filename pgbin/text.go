package pgbin

import (
	"github.com/Alisaien/pgq/internal"
	"unsafe"
)

type Text string

type _string struct{}

var String = _string{}

func (_string) Read(iter *internal.Iterator) string {
	if iter.Err != nil {
		return ""
	}
	return string(iter.Read())
}

func (_string) Write(ptr unsafe.Pointer, stream *internal.Stream) {
	stream.Write([]byte(*(*string)(ptr)))
}
