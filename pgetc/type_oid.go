package pgetc

const OIDOID  = 26
type OID uint32

func (o OID) WriteType(stream *Stream) {
	stream.WriteUint32(OIDOID)
	o.WriteValue(stream)
}

func (o OID) WriteValue(stream *Stream) {
	stream.WriteUint32(4)
	o.WriteBinary(stream)
}

func (o OID) WriteBinary(stream *Stream) {
	stream.WriteUint32(uint32(o))
}
