package pgetc

const BoolOID = 16
type Bool bool

func (b Bool) WriteType(stream *Stream) {
	stream.WriteUint32(BoolOID)
	b.WriteValue(stream)
}

func (b Bool) WriteValue(stream *Stream) {
	stream.WriteUint32(1)
	b.WriteBinary(stream)
}

func (b Bool) WriteBinary(stream *Stream) {
	if b {
		stream.WriteByte1(1)
	} else {
		stream.WriteByte1(0)
	}
}
