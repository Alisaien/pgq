package pgval

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
)

type _bool struct{}
type _boolPtr struct{}

var Bool = _bool{}
var BoolPtr =_boolPtr{}

func (_bool) Read(iter *pgetc.Iterator) bool {
	if iter.Next4() != nil {
		return false
	}

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		iter.ReportError(pgetc.ErrNull)
		return false
	}

	if iter.Next(int(size)) != nil {
		return false
	}

	return pgbin.Bool.Read(iter)
}

func (_boolPtr) Read(iter *pgetc.Iterator) *bool {
	if iter.Next4() != nil {
		return nil
	}

	size := int32(binary.BigEndian.Uint32(iter.Read()))
	if size == -1 {
		return nil
	}

	if iter.Next(int(size)) != nil {
		return nil
	}

	val := pgbin.Bool.Read(iter)
	return &val
}
