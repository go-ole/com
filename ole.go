package com

import (
	"fmt"
	"strings"
)

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

// EXCEPINFO defines exception info.
//
// XXX: Documentation needs to be improved.
// XXX: Type name needs to be in Go style.
type EXCEPINFO struct {
	wCode             uint16
	wReserved         uint16
	bstrSource        *uint16
	bstrDescription   *uint16
	bstrHelpFile      *uint16
	dwHelpContext     uint32
	pvReserved        uintptr
	pfnDeferredFillIn uintptr
	scode             uint32
}

// String convert EXCEPINFO to string.
func (e EXCEPINFO) String() string {
	var src, desc, hlp string
	if e.bstrSource == nil {
		src = "<nil>"
	} else {
		src = BstrToString(e.bstrSource)
	}

	if e.bstrDescription == nil {
		desc = "<nil>"
	} else {
		desc = BstrToString(e.bstrDescription)
	}

	if e.bstrHelpFile == nil {
		hlp = "<nil>"
	} else {
		hlp = BstrToString(e.bstrHelpFile)
	}

	return fmt.Sprintf(
		"wCode: %#x, bstrSource: %v, bstrDescription: %v, bstrHelpFile: %v, dwHelpContext: %#x, scode: %#x",
		e.wCode, src, desc, hlp, e.dwHelpContext, e.scode,
	)
}

// Error implements error interface and returns error string.
func (e EXCEPINFO) Error() string {
	if e.bstrDescription != nil {
		return strings.TrimSpace(BstrToString(e.bstrDescription))
	}

	src := "Unknown"
	if e.bstrSource != nil {
		src = BstrToString(e.bstrSource)
	}

	code := e.scode
	if e.wCode != 0 {
		code = uint32(e.wCode)
	}

	return fmt.Sprintf("%v: %#x", src, code)
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

// Currency is COM Currency type.
//
// This is not a float. The decimal point is between the left 15 digits and the
// right 4 digits. In base 10. Should only handle maths with other Currency
// objects.
//
// Does not currently support conversions from or to float64.
type Currency struct {
	Value int64
}

// ToString implements string interface and returns string form of currency.
func (c Currency) ToString() string {
	const placement = int64(10000)
	// Remove everything except last 4 digits.
	left := c.Value % placement
	// Shift number right 4 digits.
	right := c.Value / placement
	return fmt.Sprintf("%d.%d", right, left)
}

// Add this currency object to another Currency object.
//
// this + right
func (c Currency) Add(right Currency) Currency {
	return Currency{Value: (c.Value + right.Value)}
}

// Subtract this currency object to another Currency object.
//
// this - right
func (c Currency) Subtract(right Currency) Currency {
	return Currency{Value: (c.Value - right.Value)}
}

// Multiply this currency object to another Currency object.
//
// this * right
func (c Currency) Multiply(right Currency) Currency {
	return Currency{Value: (c.Value * right.Value)}
}

// Divide this currency object to another Currency object.
//
// this / right
func (c Currency) Divide(right Currency) Currency {
	return Currency{Value: (c.Value / right.Value)}
}

// Decimal is the COM Decimal type.
type Decimal struct {
	reserved uint16

	// Where the decimal point is, between 0 and 28.
	//
	// So if the number is 12.345, then Scale is 3. If the number is 1.2345,
	// then number is 4.
	Scale int8

	// Whether number is negative.
	// Positive number if 0. This means that the number is negative if it has
	// any value besides 0.
	Sign int8

	// Hi right side number.
	Hi uint32

	// Low is left side number.
	Low uint64
}

// DecimalNegative whether Decimal type is negative.
const DecimalNegative int8 = int8(0x80)

// FileTime is the datetime object for files.
//
// It has a strange time format that must be converted.
//
// Reference: https://msdn.microsoft.com/en-us/library/windows/desktop/ms724284(v=vs.85).aspx
// XXX: Complete documentation.
// XXX: Provide helpers for converting to time.Time.
type FileTime struct {
	LowDateTime  int32
	HighDateTime int32
}

// Point is 2D vector type.
type Point struct {
	X int32
	Y int32
}

// Msg is message between processes.
type Msg struct {
	Hwnd    uint32
	Message uint32
	Wparam  int32
	Lparam  int32
	Time    uint32
	Pt      Point
}

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
