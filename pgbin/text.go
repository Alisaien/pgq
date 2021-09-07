package pgbin

import (
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

type Text string

type _string struct{}

var String = _string{}

func (_string) Read(iter *pgetc.Iterator) string {
	if iter.Err != nil {
		return ""
	}
	return string(iter.Read())
}

func (_string) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.Write([]byte(*(*string)(ptr)))
}
