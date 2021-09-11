package pgetc

import "fmt"

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

func (iter *Iterator) Next4() error {
	return iter.Next(4)
}
