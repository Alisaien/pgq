package pgq

import (
	"sync"
	"unsafe"
)

var pool sync.Pool

func init() {
	pool = sync.Pool{
		New: func() interface{} {
			return new(buffer)
		},
	}
}

type buffer []byte

func (b *buffer) len() int {
	return len(*b)
}

func (b *buffer) write(p *buffer) *buffer {
	*b = append(*b, *p...)
	return b
}

func (b *buffer) writeString(s string) *buffer {
	*b = append(*b, s...)
	return b
}

func (b *buffer) writeSQL(qp string) {
	/*l := len(s.args)
	s.args = append(s.args, args...)
	s.buf.Grow(len(qp))

	for len(qp) > 0 {
		start := 0
		for i, c := range qp {
			s.buf.WriteRune(c)
			// look ahead to make it is a placeholder
			if c == '$' && qp[i+1] > '0' && qp[i+1] <= '9' {
				start = i + 1
				break
			}
		}
		// in this case we case traversed the whole string and are done
		if start == 0 || start == len(qp) {
			break
		}

		n := 0
		next := start
		for _, c := range qp[start:] {
			if c >= '0' && c <= '9' {
				n = n*10 + int(c-'0')
				next++
			} else {
				break
			}
		}

		s.buf.WriteString(strconv.Itoa(l + n))
		if next == len(qp) {
			break
		}

		qp = qp[next:]
	}

	s.buf.WriteString(" ")
	return s*/
}

func (b *buffer) string() string {
	return *(*string)(unsafe.Pointer(b))
}
