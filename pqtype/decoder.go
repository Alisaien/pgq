package pqtype

type Decoder interface {
	FromBinary(src []byte) ([]byte, error)
}
