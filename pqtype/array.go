package pqtype

const (
	arrayHeaderSize = 12
	arrayDimsSize = 8
)

type ArrayHeader struct {
	Null Bool
	OID OID
	Dims []ArrayDims
}

type ArrayDims struct {
	Len Int4
	LowerBound Int4
}

func (ah *ArrayHeader) FromBinary(src []byte) ([]byte, error) {
	if len(src) < arrayHeaderSize {
		return nil, ErrInsufficientBytes
	}

	var dims Int4
	src, _ = dims.fromBinary(src)
	src, _ = ah.Null.fromBinary(src)
	src, _ = ah.OID.fromBinary(src)

	if dims > 0 {
		ah.Dims = make([]ArrayDims, dims)
	}

	if len(src) < arrayHeaderSize+int(dims)*arrayDimsSize {
		return nil, ErrInsufficientBytes
	}
	for i := range ah.Dims {
		src, _ = ah.Dims[i].Len.fromBinary(src)
		src, _ = ah.Dims[i].LowerBound.fromBinary(src)
	}

	return src, nil
}
