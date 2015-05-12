package com

// GUID is Windows API specific GUID type.
//
// This exists to match Windows GUID type for direct passing for COM.
// Format is in xxxxxxxx-xxxx-xxxx-xxxxxxxxxxxxxxxx.
type GUID struct {
	// 4 bytes
	Data1 uint32

	// Data2 next 2 bytes
	Data2 uint16

	// Data3 next 2 bytes
	Data3 uint16

	// Data4 is last 8 bytes
	Data4 [8]byte
}

var (
	// NullInterfaceID is null Interface ID, used when no other Interface ID is known.
	NullInterfaceID = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}

	// IUnknownInterfaceID is for IUnknown interfaces.
	IUnknownInterfaceID = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

	// IDispatchInterfaceID is for IDispatch interfaces.
	IDispatchInterfaceID = &GUID{0x00020400, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

	// IConnectionPointContainerInterfaceID is for IConnectionPointContainer interfaces.
	IConnectionPointContainerInterfaceID = &GUID{0xB196B284, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}

	// IConnectionPointInterfaceID is for IConnectionPoint interfaces.
	IConnectionPointInterfaceID = &GUID{0xB196B286, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}

	// IInspectableInterfaceID is for IInspectable interfaces.
	IInspectableInterfaceID = &GUID{0xaf86e2e0, 0xb12d, 0x4c6a, [8]byte{0x9c, 0x5a, 0xd7, 0xaa, 0x65, 0x10, 0x1e, 0x90}}

	// IProvideClassInfoInterfaceID is for IProvideClassInfo interfaces.
	IProvideClassInfoInterfaceID = &GUID{0xb196b283, 0xbab4, 0x101a, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
)

// IsEqual compares current GUID with another GUID.
func (g GUID) IsEqual(guid GUID) bool {
	return IsEqualGUID(&g, &guid)
}

// IsEqualGUID compares two GUID.
//
// Not constant time comparison.
func IsEqualGUID(guid1, guid2 *GUID) bool {
	return guid1.Data1 == guid2.Data1 &&
		guid1.Data2 == guid2.Data2 &&
		guid1.Data3 == guid2.Data3 &&
		guid1.Data4[0] == guid2.Data4[0] &&
		guid1.Data4[1] == guid2.Data4[1] &&
		guid1.Data4[2] == guid2.Data4[2] &&
		guid1.Data4[3] == guid2.Data4[3] &&
		guid1.Data4[4] == guid2.Data4[4] &&
		guid1.Data4[5] == guid2.Data4[5] &&
		guid1.Data4[6] == guid2.Data4[6] &&
		guid1.Data4[7] == guid2.Data4[7]
}
