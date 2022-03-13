package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
)

type _bool struct{}
type _boolPtr struct{}

var Bool _bool
var BoolPtr _boolPtr

func (_bool) Read(iter *pgetc.Iterator) bool {
	if iter.Next(4) != nil {
		return false
	}

	if binary.BigEndian.Uint32(iter.Read()) != pgetc.BoolOID {
		iter.ReportError(pgetc.ErrUnexpectedType)
		return false
	}

	return pgval.Bool.Read(iter)
}

func (_boolPtr) Read(iter *pgetc.Iterator) *bool {
	if iter.Next(4) != nil {
		return nil
	}

	if binary.BigEndian.Uint32(iter.Read()) != pgetc.BoolOID {
		iter.ReportError(pgetc.ErrUnexpectedType)
		return nil
	}

	return pgval.BoolPtr.Read(iter)
}
