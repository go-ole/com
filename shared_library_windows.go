// +build windows,!cgo

package com

import (
	"unicode/utf16"
	"unsafe"

	syscall "golang.org/x/sys/windows"
)

var NotImplementedError = NewError(NotImplementedErrorCode)

var (
	modcombase     = syscall.NewLazyDLL("combase.dll")
	modkernel32, _ = syscall.LoadDLL("kernel32.dll")
	modoleaut32, _ = syscall.LoadDLL("oleaut32.dll")
	modmsvcrt, _   = syscall.LoadDLL("msvcrt.dll")
	moduser32, _   = syscall.LoadDLL("user32.dll")
)

// Kernel32.dll functions.
var (
	procGetUserDefaultLCID, _ = modkernel32.FindProc("GetUserDefaultLCID")
	procCopyMemory, _         = modkernel32.FindProc("RtlMoveMemory")
)

// OleAut32.dll functions.
var (
	procSysAllocString, _     = modoleaut32.FindProc("SysAllocString")
	procSysAllocStringLen, _  = modoleaut32.FindProc("SysAllocStringLen")
	procSysFreeString, _      = modoleaut32.FindProc("SysFreeString")
	procSysStringLen, _       = modoleaut32.FindProc("SysStringLen")
	procCreateDispTypeInfo, _ = modoleaut32.FindProc("CreateDispTypeInfo")
	procCreateStdDispatch, _  = modoleaut32.FindProc("CreateStdDispatch")
	procGetActiveObject, _    = modoleaut32.FindProc("GetActiveObject")
)

// User32.dll functions.
var (
	procGetMessageW, _      = moduser32.FindProc("GetMessageW")
	procDispatchMessageW, _ = moduser32.FindProc("DispatchMessageW")
)

func SysAllocString(v string) (ss *int16) {
	pss, _, _ := procSysAllocString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v))))
	ss = (*int16)(unsafe.Pointer(pss))
	return
}

func SysAllocStringLen(v string) (ss *int16) {
	utf16 := utf16.Encode([]rune(v + "\x00"))
	ptr := &utf16[0]

	pss, _, _ := procSysAllocStringLen.Call(uintptr(unsafe.Pointer(ptr)), uintptr(len(utf16)-1))
	ss = (*int16)(unsafe.Pointer(pss))
	return
}

func SysFreeString(v *int16) (err error) {
	hr, _, _ := procSysFreeString.Call(uintptr(unsafe.Pointer(v)))
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

func SysStringLen(v *int16) uint32 {
	l, _, _ := procSysStringLen.Call(uintptr(unsafe.Pointer(v)))
	return uint32(l)
}

func CreateStdDispatch(unk *iunknown.IUnknown, v uintptr, ptinfo *iunknown.IUnknown) (disp *idispatch.Dispatch, err error) {
	hr, _, _ := procCreateStdDispatch.Call(
		uintptr(unsafe.Pointer(unk)),
		v,
		uintptr(unsafe.Pointer(ptinfo)),
		uintptr(unsafe.Pointer(&disp)))
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

func CreateDispTypeInfo(idata *INTERFACEDATA) (pptinfo *iunknown.IUnknown, err error) {
	hr, _, _ := procCreateDispTypeInfo.Call(
		uintptr(unsafe.Pointer(idata)),
		uintptr(GetUserDefaultLCID()),
		uintptr(unsafe.Pointer(&pptinfo)))
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

func copyMemory(dest unsafe.Pointer, src unsafe.Pointer, length uint32) {
	procCopyMemory.Call(uintptr(dest), uintptr(src), uintptr(length))
}

func GetUserDefaultLocaleID() (lcid uint32) {
	ret, _, _ := procGetUserDefaultLCID.Call()
	lcid = uint32(ret)
	return
}

func GetMessage(msg *Msg, hwnd uint32, MsgFilterMin uint32, MsgFilterMax uint32) (ret int32, err error) {
	r0, _, err := procGetMessageW.Call(uintptr(unsafe.Pointer(msg)), uintptr(hwnd), uintptr(MsgFilterMin), uintptr(MsgFilterMax))
	ret = int32(r0)
	return
}

func DispatchMessage(msg *Msg) (ret int32) {
	r0, _, _ := procDispatchMessageW.Call(uintptr(unsafe.Pointer(msg)))
	ret = int32(r0)
	return
}
