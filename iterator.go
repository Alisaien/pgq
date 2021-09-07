package pgq

import (
	"github.com/Alisaien/pgq/internal"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
	"unsafe"
)

func NewIterator() *Iterator {
	return new(Iterator)
}

type Iterator internal.Iterator

func (iter *Iterator) ReadBool() bool {
	return pgtyp.Bool.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadBoolPtr() *bool {
	return pgtyp.BoolPtr.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadComposite() unsafe.Pointer {
	return pgtyp.Composite.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadInt() int {
	return pgtyp.Int.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadIntPtr() *int {
	return pgtyp.IntPtr.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadInt16() int16 {
	return pgtyp.Int16.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadInt16Ptr() *int16 {
	return pgtyp.Int16Ptr.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadString() string {
	return pgtyp.String.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadStringPtr() *string {
	return pgtyp.StringPtr.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadUUID() pgetc.UUID {
	return pgtyp.UUID.Read((*internal.Iterator)(iter))
}

func (iter *Iterator) ReadUUIDPtr() *pgetc.UUID {
	return pgtyp.UUIDPtr.Read((*internal.Iterator)(iter))
}
