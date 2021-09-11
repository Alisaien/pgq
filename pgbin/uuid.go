package pgbin

import (
	"github.com/Alisaien/pgq/pgetc"
)

type _uuid struct{}

var UUID = _uuid{}

func (_uuid) Read(iter *pgetc.Iterator) pgetc.UUID {
	var uuid pgetc.UUID
	if iter.Err() != nil {
		return uuid
	}

	copy(uuid[:], iter.Read())
	return uuid
}
