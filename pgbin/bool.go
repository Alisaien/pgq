package pgbin

import (
	"github.com/Alisaien/pgq/pgetc"
)

type _bool struct{}

var Bool = _bool{}

func (_bool) Read(iter *pgetc.Iterator) bool {
	if iter.Err() != nil {
		return false
	}
	return iter.Read()[0] == 1
}
