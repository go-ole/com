package types

import (
	"fmt"
	"strings"
)

// ExceptionInfo defines exception info.
//
// XXX: Documentation needs to be improved.
//
// EXCEPINFO in Windows API.
type ExceptionInfo struct {
	Code              uint16
	_                 uint16
	bstrSource        *uint16
	bstrDescription   *uint16
	bstrHelpFile      *uint16
	dwHelpContext     uint32
	pvReserved        uintptr
	pfnDeferredFillIn uintptr
	scode             uint32
}

// String convert EXCEPINFO to string.
func (e *ExceptionInfo) String() string {
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
func (e *ExceptionInfo) Error() string {
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
