package pqtype

const (
	ByteaOID = 17
	QCharOID = 18
	NameOID  = 19
	Int8OID  = 20


	TIDOID          = 27
	XIDOID          = 28
	CIDOID          = 29
	JSONOID         = 114
	PointOID        = 600
	LsegOID         = 601
	PathOID         = 602
	BoxOID          = 603
	PolygonOID      = 604
	LineOID         = 628
	CIDROID         = 650
	CIDRArrayOID    = 651
	Float4OID       = 700
	Float8OID       = 701
	CircleOID       = 718
	UnknownOID      = 705
	MacaddrOID      = 829
	InetOID         = 869
	BoolArrayOID    = 1000
	Int2ArrayOID    = 1005
	Int4ArrayOID    = 1007
	TextArrayOID    = 1009
	ByteaArrayOID   = 1001
	BPCharArrayOID  = 1014
	VarcharArrayOID = 1015
	Int8ArrayOID    = 1016
	Float4ArrayOID  = 1021
	Float8ArrayOID  = 1022
	ACLItemOID      = 1033
	ACLItemArrayOID = 1034
	InetArrayOID    = 1041
	BPCharOID       = 1042
	VarcharOID      = 1043

	TimeOID           = 1083
	TimestampOID      = 1114
	TimestampArrayOID = 1115
	DateArrayOID      = 1182

	TimestamptzArrayOID = 1185
	IntervalOID         = 1186
	NumericArrayOID     = 1231
	BitOID              = 1560
	VarbitOID           = 1562
	NumericOID          = 1700
	RecordOID           = 2249

	JSONBOID            = 3802
	JSONBArrayOID       = 3807
	DaterangeOID        = 3912
	Int4rangeOID        = 3904
	NumrangeOID         = 3906
	TsrangeOID          = 3908
	TsrangeArrayOID     = 3909
	TstzrangeOID        = 3910
	TstzrangeArrayOID   = 3911
	Int8rangeOID        = 3926
)

const (
	typeOffset  = 0
	sizeOffset  = typeOffset + 4
	valueOffset = sizeOffset + 4
)
