package types

// TYPEDESC defines data type.
type TYPEDESC struct {
	Hreftype uint32
	VT       uint16
}

// IDLDESC defines IDL info.
type IDLDESC struct {
	DwReserved uint32
	WIDLFlags  uint16
}

// TYPEATTR defines type info.
type TYPEATTR struct {
	Guid             GUID
	Lcid             uint32
	dwReserved       uint32
	MemidConstructor int32
	MemidDestructor  int32
	LpstrSchema      *uint16
	CbSizeInstance   uint32
	Typekind         int32
	CFuncs           uint16
	CVars            uint16
	CImplTypes       uint16
	CbSizeVft        uint16
	CbAlignment      uint16
	WTypeFlags       uint16
	WMajorVerNum     uint16
	WMinorVerNum     uint16
	TdescAlias       TYPEDESC
	IdldescType      IDLDESC
}
