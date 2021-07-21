package pqtype

import (
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

type UUID []byte

func (u *UUID) DecodeBinary(src []byte) ([]byte, error) {
	if len(src) < 16 {
		return nil, ErrInsufficientBytes
	}

	if *u == nil {
		*u = make([]byte, 16)
	}

	copy(*u, src)
	return src[16:], nil
}

func (u UUID) MarshalJSON() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *UUID) UnmarshalJSON(src []byte) error {
	var err error
	if *u == nil {
		*u = make([]byte, 16)
	}

	*u, err = parseUUID(string(src))
	return err
}

func (u UUID) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	var err error
	*ptr, err = parseUUID(iter.ReadString())

	if err != nil {
		iter.ReportError("DecodeUUID", err.Error())
	}
}

func (u UUID) Encode(_ unsafe.Pointer, stream *jsoniter.Stream) {
	_, _ = stream.Write(u)
}

func (u UUID) IsEmpty(ptr unsafe.Pointer) bool {
	return u == nil
}

func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])
}

func parseUUID(src string) ([]byte, error) {
	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		return nil, fmt.Errorf("cannot parse UUID %v", src)
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, 16)

	copy(dst, buf)
	return dst, err
}