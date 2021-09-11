package pgq

import (
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
	"unsafe"
)

type Stream pgetc.Stream

type Streamable interface {
	WriteTo(stream *Stream)
}

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

func (stream *Stream) WriteCompositeType(oid pgetc.OID, v Streamable) {
	stream.Stream().WriteUint32(uint32(oid))
	sp := stream.Len()
	stream.Stream().WriteUint32(0)
	v.WriteTo(stream)
	stream.Stream().SetUint32(sp, uint32(stream.Len() - sp))
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
