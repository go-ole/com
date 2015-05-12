// #include "objbase.h"
// +build windows,cgo

package com

import "C"

import (
	"reflect"
	"unsafe"

	syscall "golang.org/x/sys/windows"
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
	return MaybeError(C.BindMoniker(
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
func ClassIDFromProgramID(programId string) (*GUID, error) {
	ptrProgramId, err := syscall.UTF16PtrFromString(programId)
	if err != nil {
		return
	}

	err = MaybeError(C.CLSIDFromProgID(
		uintptr(unsafe.Pointer(ptrProgramId)),
		uintptr(unsafe.Pointer(&classId))))
	return
}

// Similar to ClassIDFromProgramID(), but will attempt to install if
// COMClassStore policy is enabled.
//
// This appears to be used when the client machine is part of Active Directory
// and the policy allows for downloading missing COM components. Also, the COM
// component would have to be registered.
func ClassIDFromProgramIDEx(programId string) (*GUID, error) {
	ptrProgramId, err := syscall.UTF16PtrFromString(programId)
	if err != nil {
		return
	}

	err = MaybeError(C.CLSIDFromProgIDEx(
		uintptr(unsafe.Pointer(ptrProgramId)),
		uintptr(unsafe.Pointer(&classId))))
	return
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
	return MaybeError(C.CoInitialize(uintptr(0)))
}

func CoInitializeEx(coinit uint32) error {
	// http://msdn.microsoft.com/en-us/library/windows/desktop/ms695279(v=vs.85).aspx
	// Suggests that the first parameter is not only optional but should always be NULL.
	return MaybeError(C.CoInitializeEx(uintptr(0), uintptr(coinit)))
}

func CoUninitialize() {
	C.CoUninitialize()
}
