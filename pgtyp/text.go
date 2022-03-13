package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
	"unsafe"
)

const TextOID = 25

type _string struct{}
type _stringPtr struct{}

var String = _string{}
var StringPtr = _stringPtr{}

func (_string) Read(iter *pgetc.Iterator) string {
	if iter.Next(4) != nil {
		return ""
	}

	if binary.BigEndian.Uint32(iter.Read()) != TextOID {
		iter.ReportError(pgetc.ErrUnexpectedType)
		return ""
	}

	return pgval.String.Read(iter)
}

func (_stringPtr) Read(iter *pgetc.Iterator) *string {
	if iter.Next(4) != nil {
		return nil
	}

	if binary.BigEndian.Uint32(iter.Read()) != TextOID {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return nil
	}

	return pgval.StringPtr.Read(iter)
}

func (_string) ReadUnsafe(iter *pgetc.Iterator) unsafe.Pointer {
	return unsafe.Pointer(StringPtr.Read(iter))
}

func (_string) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(TextOID)
	pgval.String.Write(ptr, stream)
}
