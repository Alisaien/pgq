package pqtype

type Encoder interface {
	ToBinary([]byte) []byte
	ToPureBinary([]byte) []byte
}
