package pqtype

const (
	arrayHeaderMinSize = 12
	arrayDimsSize      = 8
)

type ArrayHeader struct {
	HasNull Bool
	ElemType OID
	Dims []ArrayDims
}

type ArrayDims struct {
	Len Int4
	LowerBound Int4
}

func (ah *ArrayHeader) FromBinary(src []byte) ([]byte, error) {
	if len(src) < arrayHeaderMinSize {
		return nil, ErrInsufficientBytes
	}

	var ndim, hasNull Int4
	src, _ = ndim.fromBinary(src)
	src, _ = hasNull.fromBinary(src) // PG sends HasNull as int32
	src, _ = ah.ElemType.fromBinary(src)

	if ndim > 0 {
		ah.Dims = make([]ArrayDims, ndim)
		ah.HasNull = hasNull == 1
	}

	if len(src) < arrayHeaderMinSize+int(ndim)*arrayDimsSize {
		return nil, ErrInsufficientBytes
	}
	for i := range ah.Dims {
		src, _ = ah.Dims[i].Len.fromBinary(src)
		src, _ = ah.Dims[i].LowerBound.fromBinary(src)
	}

	return src, nil
}
