package pgq

import (
	"github.com/Alisaien/pgq/pgetc"
	"github.com/Alisaien/pgq/pgtyp"
)

func init() {
	pgetc.RegisterUnsafeReader(false, pgtyp.Bool)
	pgetc.RegisterUnsafeReader(0, pgtyp.Int)
	pgetc.RegisterUnsafeReader(int16(0), pgtyp.Int16)
	pgetc.RegisterUnsafeReader(int32(0), pgtyp.Int32)
	pgetc.RegisterUnsafeReader("", pgtyp.String)
}
