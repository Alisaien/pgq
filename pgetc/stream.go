package pgetc

import (
	"encoding/binary"
	"fmt"
)

type Stream struct {
	buf []byte
	err error
}

func NewStream(buf []byte) *Stream {
	return &Stream{buf: buf}
}

func (stream *Stream) Err() error {
	return stream.err
}

func (stream *Stream) ReportError(err error) {
	stream.err = fmt.Errorf("%s at %d", err.Error(), len(stream.buf))
}

func (stream *Stream) Write(p []byte) (int, error) {
	stream.buf = append(stream.buf, p...)
	return len(p), nil
}

func (stream *Stream) WriteByte1(b byte) {
	stream.buf = append(stream.buf, b)
}

func (stream *Stream) WriteInt32(v int32) {
	stream.WriteUint32(uint32(v))
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

func (stream *Stream) WriteUint64(v uint64) {
	wp := len(stream.buf)
	stream.buf = append(stream.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(stream.buf[wp:], v)
}

func (stream *Stream) SetUint32(sp int, v uint32) {
	binary.BigEndian.PutUint32(stream.buf[sp:], v)
}

func (stream *Stream) Len() int {
	return len(stream.buf)
}

func (stream *Stream) Bytes() []byte {
	return stream.buf
}

func (stream *Stream) Reset() {
	stream.buf = stream.buf[:0]
}
