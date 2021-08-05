package pqtype

type Decoder interface {
	FromBinary(src []byte) ([]byte, error)
	FromPureBinary(src []byte) ([]byte, error)
}
