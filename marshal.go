package pgq

import (
	"github.com/Alisaien/pgq/internal"
	"github.com/modern-go/reflect2"
)

func Marshal(val internal.Streamable) []byte {
	stream := new(internal.Stream)
	val.Write(reflect2.PtrOf(val), stream)

	return stream.Bytes()
}