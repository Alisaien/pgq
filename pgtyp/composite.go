package pgtyp

import (
	"encoding/binary"
	"github.com/Alisaien/pgq/internal"
	"github.com/Alisaien/pgq/pgetc"
	"unsafe"
)

var composites map[uint32]pgetc.UnsafeReader

type _composite struct{}

var Composite _composite

func (_composite) Read(iter *internal.Iterator) unsafe.Pointer {
	if iter.Next4() != nil {
		return nil
	}

	oid := binary.BigEndian.Uint32(iter.Read())
	cmp, ok := composites[oid]
	if !ok {
		iter.Error(pgetc.ErrUnknownType)
		return nil
	}

	return cmp.ReadUnsafe(iter)
}
