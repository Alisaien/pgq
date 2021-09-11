package pgetc

import (
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

const UUIDOID = 2950
type UUID [16]byte

func (u UUID) WriteType(stream *Stream) {
	stream.WriteUint32(UUIDOID)
	u.WriteValue(stream)
}

func (u UUID) WriteValue(stream *Stream) {
	stream.WriteUint32(16)
	u.WriteBinary(stream)
}

func (u UUID) WriteBinary(stream *Stream) {
	stream.Write(u[:])
}

// --------------------------------

func (u UUID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + u.String() + `"`), nil
}

func (u *UUID) UnmarshalJSON(src []byte) error {
	if len(src) < 32 {
		return fmt.Errorf("cannot parse UUID %v", src)
	}

	var err error
	*u, err = parseUUID(string(src[1 : len(src)-1]))

	return err
}

func (u *UUID) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	var err error
	*(*UUID)(ptr), err = parseUUID(iter.ReadString())
	iter.ReadObject()

	if err != nil {
		iter.ReportError("DecodeUUID", err.Error())
	}
}

func (u *UUID) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteString((*UUID)(ptr).String())
}

func (u *UUID) IsEmpty(_ unsafe.Pointer) bool {
	return false
}

func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])
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
