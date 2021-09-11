package pgq

import (
	"testing"
)

func BenchmarkStream_WriteCompositeType(b *testing.B) {
	buf := make([]byte, 0 , 128)
	RegisterOID((*Foo)(nil), 1234)
	f := &Foo{Bar: 1}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stream := NewStream(buf)
		stream.WriteCompositeType(f)
	}
}
