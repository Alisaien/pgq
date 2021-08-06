package pqtype

type TypeEncoder interface {
	EncodeType([]byte) []byte
}

type ValueEncoder interface {
	EncodeValue([]byte) []byte
}

type Writer interface {
	Write([]byte) []byte
}
