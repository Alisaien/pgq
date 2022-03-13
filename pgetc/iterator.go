package pgetc

import (
	"encoding/binary"
	"fmt"
)

type Iterator struct {
	err error
	csr int
	nxt int
	src []byte
}

func NewIterator(src []byte) *Iterator {
	return &Iterator{src: src}
}

func (iter *Iterator) Err() error {
	return iter.err
}

func (iter *Iterator) ReportError(err error) {
	iter.err = fmt.Errorf("%s at %d", err.Error(), iter.csr)
}

func (iter *Iterator) Read() []byte {
	csr := iter.csr
	iter.csr = iter.nxt

	return iter.src[csr:iter.nxt]
}

func (iter *Iterator) ReadByte() (byte, error) {
	if err := iter.Next(1); err != nil {
		return 0, err
	}
	return iter.Read()[0], nil
}

func (iter *Iterator) ReadUint16() uint16 {
	if iter.Next(2) != nil {
		return 0
	}
	return binary.BigEndian.Uint16(iter.Read())
}

func (iter *Iterator) ReadUint32() uint32 {
	if iter.Next(4) != nil {
		return 0
	}
	return binary.BigEndian.Uint32(iter.Read())
}

func (iter *Iterator) ReadUint64() uint64 {
	if iter.Next(8) != nil {
		return 0
	}
	return binary.BigEndian.Uint64(iter.Read())
}

// Next prepares the next n bytes for reading
func (iter *Iterator) Next(n int) error {
	if iter.Err() == nil {
		iter.nxt += n
		if iter.nxt > len(iter.src) {
			iter.ReportError(ErrEOF)
		}
	}

	return iter.Err()
}
