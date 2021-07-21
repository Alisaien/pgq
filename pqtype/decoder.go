package pqtype

type BinaryDecoder interface {
	DecodeBinary(src []byte) ([]byte, error)
}
