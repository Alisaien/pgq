package pqtype

import "encoding/binary"

//go:generate stringer -type=Regconfig

const (
	RegconfigOID  = 3734
	regconfigSize = 4
)

type Regconfig int32
const (
	Arabic     Regconfig = 13068
	Danish     Regconfig = 13070
	Dutch      Regconfig = 13072
	English    Regconfig = 13074
	Finnish    Regconfig = 13076
	French     Regconfig = 13078
	German     Regconfig = 13080
	Greek      Regconfig = 13082
	Hungarian  Regconfig = 13084
	Indonesian Regconfig = 13086
	Irish      Regconfig = 13088
	Italian    Regconfig = 13090
	Lithuanian Regconfig = 13092
	Nepali     Regconfig = 13094
	Norwegian  Regconfig = 13096
	Portuguese Regconfig = 13098
	Romanian   Regconfig = 13100
	Russian    Regconfig = 13102
	Spanish    Regconfig = 13104
	Swedish    Regconfig = 13106
	Tamil      Regconfig = 13108
	Turkish    Regconfig = 13110
)

func (v *Regconfig) FromBinary(src []byte) ([]byte, error) {
	if len(src) < valueOffset+regconfigSize {
		return nil, ErrInsufficientBytes
	}

	typ := int32(binary.BigEndian.Uint32(src))
	if typ != RegconfigOID {
		return nil, &DecodeTypeErr{expected: RegconfigOID, got: typ}
	}

	if int32(binary.BigEndian.Uint32(src[sizeOffset:])) == -1 {
		return nil, ErrNullValue
	}

	return v.fromBinary(src[valueOffset:])
}

func (v *Regconfig) fromBinary(src []byte) ([]byte, error) {
	*v = Regconfig(binary.BigEndian.Uint32(src))
	return src[int4Size:], nil
}
