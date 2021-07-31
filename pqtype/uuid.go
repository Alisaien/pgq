package pqtype

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

const (
	uuidSize     = 16
	UUIDOID      = 2950
)

type UUID [16]byte

// ----- UUID -----

func (v *UUID) FromBinary(src []byte) ([]byte, error) {
	if len(src) < valueOffset + uuidSize {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != UUIDOID {
		return nil, &DecodeTypeErr{expected: UUIDOID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[sizeOffset:])) == -1 {
		return nil, ErrNullValue
	}

	return v.fromBinary(src)
}

func (v *UUID) fromBinary(src []byte) ([]byte, error) {
	copy(v[:], src)
	return src[uuidSize:], nil
}

func (v UUID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}

func (v *UUID) UnmarshalJSON(src []byte) error {
	var err error

	*v, err = parseUUID(string(src))
	return err
}

func (v UUID) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	var err error
	*(*UUID)(ptr), err = parseUUID(iter.ReadString())

	if err != nil {
		iter.ReportError("DecodeUUID", err.Error())
	}
}

func (v UUID) Encode(_ unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteString(v.String())
}

func (v UUID) IsEmpty(_ unsafe.Pointer) bool {
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

// ----- UUIDArray -----

const (
	UUIDArrayOID = 2951
)

type UUIDArray []UUID

func (ua *UUIDArray) FromBinary(src []byte) ([]byte, error) {
	const size = valueOffset + arrayHeaderMinSize

	if len(src) < size {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != UUIDArrayOID {
		return nil, &DecodeTypeErr{expected: UUIDArrayOID, got: typ}
	}

	var (
		err    error
		header ArrayHeader
	)
	src, err = header.FromBinary(src[valueOffset:])
	if err != nil {
		return nil, err
	}

	if len(header.Dims) == 0 {
		*ua = UUIDArray{}
		return src, nil
	}
	if len(header.Dims) > 1 {
		return nil, ErrTooManyDims
	}

	uuids := make(UUIDArray, header.Dims[0].Len)
	var ln Int4
	for i := range uuids {
		src, _ = ln.fromBinary(src)
		if ln == -1 {
			return nil, ErrNullValue
		}
		src, _ = uuids[i].fromBinary(src)
	}

	*ua = uuids
	return src, nil
}
