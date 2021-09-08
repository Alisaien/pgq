package pgetc

import "fmt"

type Iterator struct {
	Err error
	csr int
	nxt int
	src []byte
}

func NewIterator(src []byte) *Iterator {
	return &Iterator{src: src}
}

func (iter *Iterator) Error(err error) {
	iter.Err = fmt.Errorf("%s at %d", err.Error(), iter.csr)
}

func (iter *Iterator) Read() []byte {
	csr := iter.csr
	iter.csr = iter.nxt

	return iter.src[csr:iter.nxt]
}

// Next prepares the next n bytes for reading
func (iter *Iterator) Next(n int) error {
	if iter.Err == nil {
		iter.nxt += n
		if iter.nxt > len(iter.src) {
			iter.Error(ErrEOF)
		}
	}

	return iter.Err
}

func (iter *Iterator) Next4() error {
	return iter.Next(4)
}
