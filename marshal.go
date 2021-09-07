package pgq

import (
	"github.com/Alisaien/pgq/pgetc"
	"github.com/modern-go/reflect2"
)

func Marshal(val pgetc.Streamable) []byte {
	stream := new(pgetc.Stream)
	val.Write(reflect2.PtrOf(val), stream)

	return stream.Bytes()
}