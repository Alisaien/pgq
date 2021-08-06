package pqtype

import "github.com/jackc/pgio"

type Text string

const TextOID = 25

func (v *Text) DecodeType(src []byte) ([]byte, error) {
	err := LenCheck(src, 0)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, TextOID)
	if err != nil {
		return nil, err
	}

	return v.DecodeValue(src)
}

func (v *Text) DecodeValue(src []byte) ([]byte, error) {
	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	v.Read(src)
	return src[size:], nil
}

func (v *Text) Read(src []byte) ([]byte, error) {
	*v = Text(src)
	return src[len(src):], nil
}

func (v Text) EncodeType(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, TextOID)
	return v.EncodeValue(buf)
}

func (v Text) EncodeValue(buf []byte) []byte {
	buf = pgio.AppendUint32(buf, uint32(len(v)))
	return v.Write(buf)
}

func (v Text) Write(buf []byte) []byte {
	return append(buf, v...)
}
