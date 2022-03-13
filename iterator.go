package pgq

import (
	"github.com/Alisaien/pgq/pgbin"
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type Iterator pgetc.Iterator

func NewIterator(src []byte) *Iterator {
	return (*Iterator)(pgetc.NewIterator(src))
}

func (iter *Iterator) Iterator() *pgetc.Iterator {
	return (*pgetc.Iterator)(iter)
}

func (iter *Iterator) Err() error {
	return iter.Iterator().Err()
}

func (iter *Iterator) ReportError(err error) {
	iter.Iterator().ReportError(err)
}

func (iter *Iterator) ReadBool() bool {
	return pgtyp.Bool.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadBoolPtr() *bool {
	return pgtyp.BoolPtr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadCompositeTypeHeader() uint32 {
	if err := (*pgetc.Iterator)(iter).Next(4); err != nil {
		return 0
	}
	return pgbin.Uint32.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadEnum(oid pgetc.OID) string {
	v := iter.ReadEnumN(oid)
	if v == nil {
		if iter.Err() == nil {
			iter.ReportError(pgetc.ErrNull)
		} else {
			return ""
		}
	}

	return *v
}

func (iter *Iterator) ReadEnumN(oid pgetc.OID) *string {
	id := iter.Iterator().ReadUint32()
	if id == 0 {
		return nil
	} else if id != uint32(oid) {
		iter.ReportError(pgetc.ErrUnexpectedType)
	}

	size := int32(iter.Iterator().ReadUint32())
	if size == -1 {
		return nil
	}

	if iter.Iterator().Next(int(size)) != nil {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return nil
	}

	v := string(iter.Iterator().Read())
	return &v
}

func (iter *Iterator) ReadInt() int {
	v := iter.ReadIntN()
	if v == nil {
		if iter.Err() == nil {
			iter.ReportError(pgetc.ErrNull)
		} else {
			return 0
		}
	}

	return *v
}

func (iter *Iterator) ReadIntN() *int {
	if iter.Iterator().ReadUint32() != pgtype.Int4OID {
		iter.ReportError(pgetc.ErrUnexpectedType)
		return nil
	}

	size := int32(iter.Iterator().ReadUint32())
	if size == -1 {
		return nil
	} else if size != 4 {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return nil
	}

	v := int(iter.Iterator().ReadUint32())
	return &v
}

func (iter *Iterator) ReadInt16() int16 {
	v := iter.ReadInt16N()
	if v == nil {
		if iter.Err() == nil {
			iter.ReportError(pgetc.ErrNull)
		} else {
			return 0
		}
	}

	return *v
}

func (iter *Iterator) ReadInt16N() *int16 {
	if iter.Iterator().ReadUint32() != pgtype.Int2OID {
		iter.ReportError(pgetc.ErrUnexpectedType)
		return nil
	}

	size := int32(iter.Iterator().ReadUint32())
	if size == -1 {
		return nil
	} else if size != 2 {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return nil
	}

	v := int16(iter.Iterator().ReadUint16())
	return &v
}

func (iter *Iterator) ReadInt64() int64 {
	v := iter.ReadInt64N()
	if v == nil {
		if iter.Err() == nil {
			iter.ReportError(pgetc.ErrNull)
		} else {
			return 0
		}
	}

	return *v
}

func (iter *Iterator) ReadInt64N() *int64 {
	if iter.Iterator().ReadUint32() != pgtype.Int8OID {
		iter.ReportError(pgetc.ErrUnexpectedType)
		return nil
	}

	size := int32(iter.Iterator().ReadUint32())
	if size == -1 {
		return nil
	} else if size != 8 {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return nil
	}

	v := int64(iter.Iterator().ReadUint64())
	return &v
}

func (iter *Iterator) ReadJSONB(v interface{}) {
	oid := iter.Iterator().ReadUint32()
	if oid != pgtype.JSONBOID {
		return
	}
	size := int32(iter.Iterator().ReadUint32())

	var (
		err  error
		data []byte
	)
	if size == -1 {
		data = []byte("null")
	} else {
		if b, err := iter.Iterator().ReadByte(); err != nil {
			return
		} else if b != 1 {
			iter.ReportError(pgetc.ErrUnknownVersion)
		}

		if err = iter.Iterator().Next(int(size) - 1); err != nil { // 1 byte already read
			return
		}
		data = iter.Iterator().Read()
	}

	if err = jsoniter.Unmarshal(data, v); err != nil {
		iter.ReportError(err)
	}
}

func (iter *Iterator) ReadString() string {
	return pgtyp.String.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadStringPtr() *string {
	return pgtyp.StringPtr.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadTime() time.Time {
	return pgtyp.Timestamptz.Read((*pgetc.Iterator)(iter))
}

func (iter *Iterator) ReadUUID() uuid.UUID {
	if iter.Iterator().ReadUint32() != pgtype.UUIDOID {
		iter.ReportError(pgetc.ErrUnexpectedType)
		return uuid.UUID{}
	}

	size := int32(iter.Iterator().ReadUint32())
	if size != 16 {
		iter.ReportError(pgetc.ErrInvalidSrcLength)
		return uuid.UUID{}
	}

	if iter.Iterator().Next(16) != nil {
		return uuid.UUID{}
	}

	id, _ := uuid.FromBytes(iter.Iterator().Read())
	return id
}
