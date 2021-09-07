package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/internal"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
	"unsafe"
)

const BoolOID = 16

type _bool struct{}
type _boolPtr struct{}

var Bool _bool
var BoolPtr _boolPtr

func (_bool) Read(iter *internal.Iterator) bool {
	if iter.Next4() != nil {
		return false
	}

	if binary.BigEndian.Uint32(iter.Read()) != BoolOID {
		iter.Error(pgetc.ErrInvalidSrcLength)
		return false
	}

	return pgval.Bool.Read(iter)
}

func (_boolPtr) Read(iter *internal.Iterator) *bool {
	if iter.Next4() != nil {
		return nil
	}

	if binary.BigEndian.Uint32(iter.Read()) != BoolOID {
		iter.Error(pgetc.ErrInvalidSrcLength)
		return nil
	}

	return pgval.BoolPtr.Read(iter)
}

func (_bool) ReadUnsafe(iter *internal.Iterator) unsafe.Pointer {
	return unsafe.Pointer(BoolPtr.Read(iter))
}

func (_bool) Write(ptr unsafe.Pointer, stream *internal.Stream) {
	stream.WriteUint32(BoolOID)
	pgval.Bool.Write(ptr, stream)
}
