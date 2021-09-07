package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/internal"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgval"
	"unsafe"
)

const UUIDOID = 2950

type _uuid struct{}
type _uuidPtr struct{}

var UUID = _uuid{}
var UUIDPtr = _uuidPtr{}

func (_uuid) Read(iter *pgetc.Iterator) pgetc.UUID {
	if iter.Next4() != nil {
		return pgetc.UUID{}
	}

	if binary.BigEndian.Uint32(iter.Read()) != UUIDOID {
		iter.Error(pgetc.ErrInvalidSrcLength)
		return pgetc.UUID{}
	}

	return pgval.UUID.Read(iter)
}

func (_uuidPtr) Read(iter *pgetc.Iterator) *pgetc.UUID {
	if iter.Next4() != nil {
		return nil
	}

	if binary.BigEndian.Uint32(iter.Read()) != UUIDOID {
		iter.Error(pgetc.ErrInvalidSrcLength)
		return nil
	}

	return pgval.UUIDPtr.Read(iter)
}

func (_uuid) ReadUnsafe(iter *pgetc.Iterator) unsafe.Pointer {
	return unsafe.Pointer(UUIDPtr.Read(iter))
}

func (_uuid) Write(ptr unsafe.Pointer, stream *pgetc.Stream) {
	stream.WriteUint32(UUIDOID)
	pgval.UUID.Write(ptr, stream)
}
