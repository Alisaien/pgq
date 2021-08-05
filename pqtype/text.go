package pqtype

type Text string

const TextOID = 25

func (v *Text) FromBinary(src []byte) ([]byte, error) {
	err := LenCheck(src, 0)
	if err != nil {
		return nil, err
	}

	src, err = TypeCheck(src, TextOID)
	if err != nil {
		return nil, err
	}

	return v.FromPureBinary(src)
}

func (v *Text) FromPureBinary(src []byte) ([]byte, error) {
	size, src := ValueSize(src)
	if size == -1 {
		return nil, ErrNullValue
	}

	buf := make([]byte, size)
	copy(buf, src)
	*v = Text(buf)

	return src[size:], nil
}
