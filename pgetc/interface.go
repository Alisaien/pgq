package pgetc

import (
	"github.com/Alisaien/pgq/internal"
	"github.com/modern-go/reflect2"
	"unsafe"
)

var unsafeReaders map[uintptr]UnsafeReader

type UnsafeReader interface {
	ReadUnsafe(iter *internal.Iterator) unsafe.Pointer
}

func RegisterUnsafeReader(obj interface{}, reader UnsafeReader) {
	unsafeReaders[reflect2.RTypeOf(obj)] = reader
}
