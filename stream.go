package pgq

import (
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
	"unsafe"
)

type Stream pgetc.Stream

func NewStream(buf []byte) *Stream {
	return (*Stream)(pgetc.NewStream(buf))
}

func (stream *Stream) WriteBool(b bool) {
	pgtyp.Bool.Write(unsafe.Pointer(&b), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteBoolPtr(b *bool) {
	pgtyp.Bool.Write(unsafe.Pointer(b), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteInt(i int) {
	pgtyp.Int.Write(unsafe.Pointer(&i), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteIntPtr(i *int) {
	pgtyp.Int.Write(unsafe.Pointer(i), (*pgetc.Stream)(stream))
}

func (stream *Stream) ReadInt16(i int16) {
	pgtyp.Int16.Write(unsafe.Pointer(&i), (*pgetc.Stream)(stream))
}

func (stream *Stream) ReadInt16Ptr(i *int16) {
	pgtyp.Int16.Write(unsafe.Pointer(i), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteString(s string) {
	pgtyp.String.Write(unsafe.Pointer(&s), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteStringPtr(s *string) {
	pgtyp.String.Write(unsafe.Pointer(s), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteUUID(uuid pgetc.UUID) {
	pgtyp.UUID.Write(unsafe.Pointer(&uuid), (*pgetc.Stream)(stream))
}
