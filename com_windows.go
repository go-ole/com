// +build windows,!cgo

package com

import (
	"reflect"
	"unsafe"

	syscall "golang.org/x/sys/windows"

	"github.com/jacobsantos/go-com/iunknown"
)

var (
	modole32, _ = syscall.LoadDLL("ole32.dll")
)

var (
	procBindMoniker, _                = modole32.FindProc("BindMoniker")
	procCLSIDFromProgID, _            = modole32.FindProc("CLSIDFromProgID")
	procCLSIDFromProgIDEx, _          = modole32.FindProc("CLSIDFromProgIDEx")
	procCLSIDFromString, _            = modole32.FindProc("CLSIDFromString")
	procCoAddRefServerProcess, _      = modole32.FindProc("CoAddRefServerProcess")
	procCoAllowSetForegroundWindow, _ = modole32.FindProc("CoAllowSetForegroundWindow")
	procCoInitialize, _               = modole32.FindProc("CoInitialize")
	procCoInitializeEx, _             = modole32.FindProc("CoInitializeEx")
	procCoUninitialize, _             = modole32.FindProc("CoUninitialize")
	procCoCreateInstance, _           = modole32.FindProc("CoCreateInstance")
	procCoTaskMemFree, _              = modole32.FindProc("CoTaskMemFree")
	procStringFromCLSID, _            = modole32.FindProc("StringFromCLSID")
	procStringFromIID, _              = modole32.FindProc("StringFromIID")
	procIIDFromString, _              = modole32.FindProc("IIDFromString")
)

// BindMoniker is a helper function for locating an object by its moniker and
// retrieving the pointer to the interface.
//
// moniker appears to be the address to the IMoniker object.
//
// option is ignored, because it is optional for now and must be '0'.
//
// interfaceID is the IID of the COM interface for result object.
//
// object that will be receive the object of IMoniker initialization.
//
// BindMoniker in Windows API.
func BindMoniker(moniker uintptr, option int32, interfaceID *GUID, object interface{}) error {
	return HResultToError(procBindMoniker.Call(
		moniker,
		uintptr(0),
		uintptr(unsafe.Pointer(interfaceID)),
		reflect.ValueOf(object).UnsafeAddr()))
}

// Retrieves Class Identifier with the given Program Identifier.
//
// The Programmatic Identifier must be registered, because it will be looked up
// in the Windows Registry. The registry entry has the following keys: CLSID,
// Insertable, Protocol and Shell
// (https://msdn.microsoft.com/en-us/library/dd542719(v=vs.85).aspx).
//
// programID identifies the class id with less precision and is not guaranteed
// to be unique. These are usually found in the registry under
// HKEY_LOCAL_MACHINE\SOFTWARE\Classes, usually with the format of
// "Program.Component.Version" with version being optional.
func ClassIDFromProgramID(programID string) (*GUID, error) {
	return classIDFromProgramIDHelper(procCLSIDFromProgID, programID)
}

// Similar to ClassIDFromProgramID(), but will attempt to install if
// COMClassStore policy is enabled.
//
// This appears to be used when the client machine is part of Active Directory
// and the policy allows for downloading missing COM components. Also, the COM
// component would have to be registered.
func ClassIDFromProgramIDEx(programID string) (*GUID, error) {
	return classIDFromProgramIDHelper(procCLSIDFromProgIDEx, programID)
}

// Implements ClassIDFromProgramID() and ClassIDFromProgramIDEx() functions.
func classIDFromProgramIDHelper(p *syscall.Proc, programID string) (classID *GUID, err error) {
	ptrProgramID, err := syscall.UTF16PtrFromString(programID)
	if err != nil {
		return
	}

	err = HResultToError(p.Call(
		uintptr(unsafe.Pointer(ptrProgramID)),
		uintptr(unsafe.Pointer(&classID))))
	return
}

// Retrieves Class ID from string representation.
//
// This is technically the string version of the GUID and will convert the
// string to object.
func ClassIDFromString(str string) (classID *GUID, err error) {
	ptrStr, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return
	}

	err = HResultToError(procCLSIDFromString.Call(
		uintptr(unsafe.Pointer(ptrStr)),
		uintptr(unsafe.Pointer(&classID))))
	return
}

// Adds reference to global per-process counter.
//
// This is for COM Servers, which is not currently supported.
//
// Returns current reference count, with possible error.
func CoAddRefServerProcess() (uint64, error) {
	referenceCount, _, err := procCoAddRefServerProcess.Call()
	return uint64(referenceCount), err
}

// COM Server unfocuses client application to focus to another window.
//
// This is for COM Servers, which is not currently supported.
//
// unknown is the foreground window process to gain focus.
func CoAllowSetForegroundWindow(unknown iunknown.IUnknown) error {
	return HResultToError(procCoAllowSetForegroundWindow.Call(
		uintptr(unsafe.Pointer(unknown)), uintptr(0)))
}

// Initializes COM library on current thread.
//
// MSDN documentation suggests that this function should not be called. Call
// CoInitializeEx() instead. The reason has to do with threading and this
// function is only for single-threaded apartments.
//
// That said, most users of the library have gotten away with just this
// function. If you are experiencing threading issues, then use
// CoInitializeEx().
func CoInitialize() error {
	// http://msdn.microsoft.com/en-us/library/windows/desktop/ms678543(v=vs.85).aspx
	// Suggests that no value should be passed to CoInitialized.
	return HResultToError(procCoInitialize.Call(uintptr(0)))
}

//
func CoInitializeEx(coinit ConcurrencyModel) error {
	// http://msdn.microsoft.com/en-us/library/windows/desktop/ms695279(v=vs.85).aspx
	// Suggests that the first parameter is not only optional but should always be NULL.
	return HResultToError(procCoInitializeEx.Call(uintptr(0), uintptr(uint32(coinit))))
}

func CoUninitialize() {
	procCoUninitialize.Call()
}

func CoTaskMemFree(memptr uintptr) {
	procCoTaskMemFree.Call(memptr)
}

func StringFromInterfaceID(interfaceID *GUID) (str string, err error) {
	var p *uint16
	err = HResultToError(procStringFromIID.Call(
		uintptr(unsafe.Pointer(interfaceID)),
		uintptr(unsafe.Pointer(&p))))
	str = LpOleStrToString(p)
	return
}

func StringFromClassID(classID *GUID) (str string, err error) {
	var p *uint16
	err = HResultToError(procStringFromCLSID.Call(
		uintptr(unsafe.Pointer(classID)),
		uintptr(unsafe.Pointer(&p))))
	str = LpOleStrToString(p)
	return
}

func InterfaceIDFromString(programID string) (classID *GUID, err error) {
	lpsz, err := syscall.UTF16PtrFromString(programID)
	if err != nil {
		return
	}
	err = HResultToError(procIIDFromString.Call(uintptr(unsafe.Pointer(lpsz)), uintptr(unsafe.Pointer(&classID))))
	return
}

func CoCreateInstance(classID, interfaceID *GUID, object interface{}) error {
	if interfaceID == nil {
		interfaceID = iunknown.InterfaceID
	}

	err = HResultToError(procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(classID)),
		0,
		AllServerClassContext,
		uintptr(unsafe.Pointer(interfaceID)),
		reflect.ValueOf(object).UnsafeAddr()))

	return
}

func GetActiveObject(classID, interfaceID *GUID) (unk *iunknown.Unknown, err error) {
	if interfaceID == nil {
		interfaceID = iunknown.InterfaceID
	}

	err = HResultToError(procGetActiveObject.Call(
		uintptr(unsafe.Pointer(classID)),
		uintptr(unsafe.Pointer(interfaceID)),
		uintptr(unsafe.Pointer(&unk))))

	return
}
