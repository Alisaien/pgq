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

	var (
		err error
		dims Int4
	)
	src, err = dims.FromBinary(src)
	if err != nil {
		return nil, err
	}

	src, err = ah.Null.FromBinary(src)
	if err != nil {
		return nil, err
	}

	src, err = ah.OID.FromBinary(src)
	if err != nil {
		return nil, err
	}

	if dims > 0 {
		ah.Dims = make([]ArrayDims, dims)
	}

	if len(src) < arrayHeaderSize+int(dims)*arrayDimsSize {
		return nil, ErrInsufficientBytes
	}
	for i := range ah.Dims {
		src, err = ah.Dims[i].Len.FromBinary(src)
		if err != nil {
			return nil, err
		}

		src, err = ah.Dims[i].LowerBound.FromBinary(src)
		if err != nil {
			return nil, err
		}
	}

	return src, nil
}
