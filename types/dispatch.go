package types

// DisplayParameter defines the arguments for method parameter.
type DisplayParameter struct {
	// Arguments array.
	Args uintptr

	// Named arguments of display ID.
	NamedArgs uintptr

	// Number of arguments.
	ArgsLength uint32

	// Number of named arguments.
	NamedArgsLength uint32
}

// PARAMDATA defines parameter data type.
//
// XXX: Documentation needs to be improved.
// XXX: Type name needs to be in Go style.
type PARAMDATA struct {
	Name *int16
	Vt   uint16
}

// METHODDATA defines method info.
//
// XXX: Documentation needs to be improved.
// XXX: Type name needs to be in Go style.
type METHODDATA struct {
	Name     *uint16
	Data     *PARAMDATA
	Dispid   int32
	Meth     uint32
	CC       int32
	CArgs    uint32
	Flags    uint16
	VtReturn uint32
}

// INTERFACEDATA defines interface info.
//
// XXX: Documentation needs to be improved.
// XXX: Type name needs to be in Go style.
type INTERFACEDATA struct {
	MethodData *METHODDATA
	CMembers   uint32
}
