package pgq

import (
	"strconv"
	"strings"
)

type Builder struct {
	args  []interface{}
	buf   strings.Builder
	where bool
}

func (s *Builder) Where(qp string, args ...interface{}) *Builder {
	if s.where {
		s.SQL("AND "+qp, args...)
	} else {
		s.SQL("WHERE "+qp, args...)
		s.where = true
	}

	return s
}

func (s *Builder) Args() []interface{} {
	return s.args
}

func (s *Builder) String() string {
	return s.buf.String()
}

func (s *Builder) SQL(qp string, args ...interface{}) *Builder {
	l := len(s.args)
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
	return s
}
