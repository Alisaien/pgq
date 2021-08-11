package pqtype

import (
	"encoding/hex"
	"fmt"
	"github.com/jackc/pgio"
	"github.com/jackc/pgtype"
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

// ----- UUID -----

const (
	UUIDOID  = 2950
	uuidSize = 16
)

type UUID [16]byte

func init() {
	jsoniter.RegisterTypeDecoder("pqtype.UUID", &UUID{})
	jsoniter.RegisterTypeEncoder("pqtype.UUID", &UUID{})
}

func (v *UUID) DecodeType(src []byte) ([]byte, error) {
	err := LenCheck(src, uuidSize)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, UUIDOID)
	if err != nil {
		return nil, err
	}

	return v.DecodeValue(src)
}

func (v *UUID) DecodeValue(src []byte) ([]byte, error) {
	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	return v.Read(src)
}

func (v *UUID) Read(src []byte) ([]byte, error) {
	copy(v[:], src)
	return src[uuidSize:], nil
}

func (v UUID) EncodeType(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, UUIDOID)
	return v.EncodeValue(buf)
}

func (v UUID) EncodeValue(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, uuidSize)
	return v.Write(buf)
}

func (v UUID) Write(buf []byte) []byte {
	return append(buf, v[:]...)
}

func (v UUID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}

func (v *UUID) UnmarshalJSON(src []byte) error {
	if len(src) < 32 {
		return fmt.Errorf("cannot parse UUID %v", src)
	}

	var err error
	*v, err = parseUUID(string(src[1:len(src)-1]))

	return err
}

func (v *UUID) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	var err error
	*(*UUID)(ptr), err = parseUUID(iter.ReadString())

	if err != nil {
		iter.ReportError("DecodeUUID", err.Error())
	}
}

func (v *UUID) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteString((*UUID)(ptr).String())
}

func (v *UUID) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (v UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", v[0:4], v[4:6], v[6:8], v[8:10], v[10:16])
}

func parseUUID(src string) ([16]byte, error) {
	var dst [16]byte

	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		return dst, fmt.Errorf("cannot parse UUID %v", src)
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return dst, err
	}

	copy(dst[:], buf)
	return dst, err
}

// ---------- UUIDArray ----------

const (
	UUIDArrayOID = 2951
)

type UUIDArray []UUID

func (v *UUIDArray) DecodeBinary(_ *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		return ErrNullValue
	}

	_, err := v.DecodeValue(src)
	return err
}

func (v *UUIDArray) DecodeType(src []byte) ([]byte, error) {
	err := LenCheck(src, arrayHeaderMinSize)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, UUIDArrayOID)
	if err != nil {
		return nil, err
	}

	return v.DecodeValue(src)
}

func (v *UUIDArray) DecodeValue(src []byte) ([]byte, error) {
	_, src = ValueSize(src)

	var header ArrayHeader
	src, err := header.FromBinary(src)
	if err != nil {
		return nil, err
	}

	if len(header.Dims) == 0 {
		*v = UUIDArray{}
		return src, nil
	}
	if len(header.Dims) > 1 {
		return nil, ErrTooManyDims
	}

	uuids := make(UUIDArray, header.Dims[0].Len)
	for i := range uuids {
		src, err = uuids[i].DecodeValue(src)
		if err != nil {
			return nil, err
		}
	}

	*v = uuids
	return src, nil
}

func (v UUIDArray) EncodeType(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, UUIDArrayOID)
	return v.EncodeValue(buf)
}

func (v UUIDArray) EncodeValue(buf []byte) []byte {
	sp := len(buf)
	buf = append(buf, 0, 0, 0, 0)
	buf = v.Write(buf)
	pgio.SetInt32(buf[sp:], int32(len(buf)-sp-4))

	return buf
}

func (v UUIDArray) Write(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, 1) // array dimensions
	buf = pgio.AppendUint32(buf, 0) // contains null
	buf = pgio.AppendUint32(buf, UUIDOID)
	buf = pgio.AppendUint32(buf, uint32(len(v)))
	buf = pgio.AppendUint32(buf, 0) // lower bound (always 0)

	for i := range v {
		buf = v[i].EncodeValue(buf)
	}

	return buf
}
