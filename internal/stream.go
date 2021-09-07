package internal

import (
	"encoding/binary"
	"unsafe"
)

type Streamable interface {
	Write(ptr unsafe.Pointer, stream *Stream)
}

type Stream struct {
	buf []byte
	Err error
}

func (stream  *Stream) Write(p []byte) {
	stream.buf = append(stream.buf, p...)
}

func (stream *Stream) WriteByte1(b byte) {
	stream.buf = append(stream.buf, b)
}

func (stream *Stream) WriteUint16(v uint16) {
	wp := len(stream.buf)
	stream.buf = append(stream.buf, 0, 0)
	binary.BigEndian.PutUint16(stream.buf[wp:], v)
}

func (stream *Stream) WriteUint32(v uint32) {
	wp := len(stream.buf)
	stream.buf = append(stream.buf, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(stream.buf[wp:], v)
}

func (stream *Stream) Bytes() []byte {
	return stream.buf
}

func (stream *Stream) Reset() {
	stream.buf = stream.buf[:0]
}