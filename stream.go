package pgq

import (
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
	"github.com/jackc/pgtype"
	jsoniter "github.com/json-iterator/go"
	"time"
	"unsafe"
)

type Stream pgetc.Stream

type Streamable interface {
	WriteTo(stream *Stream)
}

func NewStream(buf []byte) *Stream {
	return (*Stream)(pgetc.NewStream(buf))
}

func (stream *Stream) Stream() *pgetc.Stream {
	return (*pgetc.Stream)(stream)
}

func (stream *Stream) Err() error {
	return stream.Stream().Err()
}

func (stream *Stream) ReportError(err error) {
	stream.Stream().ReportError(err)
}

func (stream *Stream) Len() int {
	return stream.Stream().Len()
}

func (stream *Stream) Bytes() []byte {
	return stream.Stream().Bytes()
}

func (stream *Stream) WriteBool(b bool) {
	pgetc.Bool(b).WriteType(stream.Stream())
}

func (stream *Stream) WriteCompositeType(oid pgetc.OID, v Streamable) {
	stream.Stream().WriteUint32(uint32(oid))
	sp := stream.Len()
	stream.Stream().WriteInt32(-1)
	if v != nil {
		v.WriteTo(stream)
		stream.Stream().SetUint32(sp, uint32(stream.Len() - sp - 4)) // -4 to account for numField (4 bytes)
	}
}

func (stream *Stream) WriteEnum(oid pgetc.OID, s string) {
	stream.Stream().WriteUint32(uint32(oid))
	stream.Stream().WriteUint32(uint32(len(s)))
	stream.Stream().Write([]byte(s))
}

func (stream *Stream) WriteInt(i int) {
	pgtyp.Int.Write(unsafe.Pointer(&i), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteInt16(i int16) {
	pgtyp.Int16.Write(unsafe.Pointer(&i), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteJSONB(v interface{}) {
	stream.Stream().WriteUint32(pgtype.JSONBOID)
	sp := stream.Len()
	stream.Stream().WriteInt32(-1)

	if err := jsoniter.NewEncoder(stream.Stream()).Encode(v); err != nil {
		stream.ReportError(err)
		return
	}

	stream.Stream().SetUint32(sp, uint32(stream.Len() - sp - 4))
}

func (stream *Stream) WriteString(s string) {
	pgtyp.String.Write(unsafe.Pointer(&s), (*pgetc.Stream)(stream))
}

func (stream *Stream) WriteTime(t time.Time) {
	stream.Stream().WriteUint32(pgtype.TimestamptzOID)
	stream.Stream().WriteUint32(8)

	microsecSinceUnixEpoch := t.Unix()*1000000 + int64(t.Nanosecond())/1000
	stream.Stream().WriteUint64(uint64(microsecSinceUnixEpoch - pgetc.MicroSecFromUnixEpochToY2K))
}

func (stream *Stream) WriteUUID(uuid pgetc.UUID) {
	uuid.WriteType(stream.Stream())
}
