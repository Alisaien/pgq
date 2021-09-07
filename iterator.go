package pgq

import (
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
	"unsafe"
)

type Iterator pgetc.Iterator

func NewIterator(src []byte) *Iterator {
	return (*Iterator)(pgetc.NewIterator(src))
}

func (iter *Iterator) ReadBool() bool {
	return pgtyp.Bool.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadBoolPtr() *bool {
	return pgtyp.BoolPtr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadComposite() unsafe.Pointer {
	return pgtyp.Composite.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadInt() int {
	return pgtyp.Int.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadIntPtr() *int {
	return pgtyp.IntPtr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadInt16() int16 {
	return pgtyp.Int16.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadInt16Ptr() *int16 {
	return pgtyp.Int16Ptr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadString() string {
	return pgtyp.String.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadStringPtr() *string {
	return pgtyp.StringPtr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadUUID() pgetc.UUID {
	return pgtyp.UUID.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadUUIDPtr() *pgetc.UUID {
	return pgtyp.UUIDPtr.Read((*pgetc.Iterator)(iter))
}