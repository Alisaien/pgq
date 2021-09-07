package pgetc

import (
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

type UUID [16]byte

func (v UUID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}

func (v *UUID) UnmarshalJSON(src []byte) error {
	if len(src) < 32 {
		return fmt.Errorf("cannot parse UUID %v", src)
	}

	var err error
	*v, err = parseUUID(string(src[1 : len(src)-1]))

	return err
}

func (v *UUID) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	var err error
	*(*UUID)(ptr), err = parseUUID(iter.ReadString())
	iter.ReadObject()

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
