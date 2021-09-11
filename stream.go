package pgq

import (
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
	"unsafe"
)

type Foo struct{
	Bar int
}

func (*Foo) Write(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteCompositeTypeHeader(1)
	stream.WriteInt((*Foo)(ptr).Bar)
}

type Stream pgetc.Stream

func NewStream(buf []byte) *Stream {
	return (*Stream)(pgetc.NewStream(buf))
}

func (stream *Stream) Len() int {
	return stream.Stream().Len()
}

func (stream *Stream) Bytes() []byte {
	return stream.Stream().Bytes()
}

func (stream *Stream) Stream() *pgetc.Stream {
	return (*pgetc.Stream)(stream)
}

func (stream *Stream) WriteBool(b bool) {
	pgetc.Bool(b).WriteType(stream.Stream())
}

func (stream *Stream) WriteCompositeTypeHeader(numField uint32) {
	(*pgetc.Stream)(stream).WriteUint32(numField)
}

func (stream *Stream) WriteInt(i int) {
	pgtyp.Int.Write(unsafe.Pointer(&i), (*pgetc.Stream)(stream))
}

func (stream *Stream) ReadInt16(i int16) {
	pgtyp.Int16.Write(unsafe.Pointer(&i), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteString(s string) {
	pgtyp.String.Write(unsafe.Pointer(&s), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteUUID(uuid pgetc.UUID) {
	uuid.WriteType(stream.Stream())
}
