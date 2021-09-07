package pgetc

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

var unsafeReaders = make(map[uintptr]UnsafeReader)

type UnsafeReader interface {
	ReadUnsafe(iter *Iterator) unsafe.Pointer
}

func RegisterUnsafeReader(obj interface{}, reader UnsafeReader) {
	unsafeReaders[reflect2.RTypeOf(obj)] = reader
}
