package com

import (
	"syscall"
	"unicode/utf16"
	"unsafe"

	api "github.com/go-ole/com/internal/util"
	"github.com/go-ole/com/types"
	"github.com/go-ole/types/idispatch"
	"github.com/go-ole/types/iunknown"
)

// GetInt32FromCall retrieves Int32 from syscall.
func GetInt32FromCall(obj, method uintptr) int32 {
	ret, _, _ := syscall.Syscall(method, 1, obj, 0, 0)
	return int32(ret)
}

// BytePtrToString converts byte pointer to a Go string.
func BytePtrToString(p *byte) string {
	a := (*[10000]uint8)(unsafe.Pointer(p))
	i := 0
	for a[i] != 0 {
		i++
	}
	return string(a[:i])
}

// UTF16PtrToString is alias for LpOleStrToString.
//
// Kept for compatibility reasons.
func UTF16PtrToString(p *uint16) string {
	return LpOleStrToString(p)
}

// LpOleStrToString converts COM Unicode to Go string.
func LpOleStrToString(p *uint16) string {
	if p == nil {
		return ""
	}

	length := lpOleStrLen(p)
	a := make([]uint16, length)

	ptr := unsafe.Pointer(p)

	for i := 0; i < int(length); i++ {
		a[i] = *(*uint16)(ptr)
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
	}

	return string(utf16.Decode(a))
}

// BstrToString converts COM binary string to Go string.
func BstrToString(p *uint16) string {
	if p == nil {
		return ""
	}
	length := SysStringLen((*int16)(unsafe.Pointer(p)))
	a := make([]uint16, length)

	ptr := unsafe.Pointer(p)

	for i := 0; i < int(length); i++ {
		a[i] = *(*uint16)(ptr)
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
	}
	return string(utf16.Decode(a))
}

// lpOleStrLen returns the length of Unicode string.
func lpOleStrLen(p *uint16) (length int64) {
	if p == nil {
		return 0
	}

	ptr := unsafe.Pointer(p)

	for i := 0; ; i++ {
		if 0 == *(*uint16)(ptr) {
			length = int64(i)
			break
		}
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
	}
	return
}

// GetInt32FromCall retrieves Int32 from syscall.
func GetInt32FromCall(obj, method uintptr) int32 {
	ret, _, _ := syscall.Syscall(method, 1, obj, 0, 0)
	return int32(ret)
}

func SysAllocString(v string) *int16 {
	return api.SysAllocString(v)
}

func SysAllocStringLen(v string) *int16 {
	return api.SysAllocStringLen(v)
}

func SysFreeString(v *int16) error {
	return api.SysFreeString(v)
}

func SysStringLen(v *int16) uint32 {
	return api.SysStringLen(v)
}

func CreateStdDispatch(unk *iunknown.IUnknown, v uintptr, ptinfo *iunknown.IUnknown) (*idispatch.Dispatch, error) {
	return api.CreateStdDispatch(unk, v, ptinfo)
}

func CreateDispTypeInfo(idata *INTERFACEDATA) (*iunknown.IUnknown, error) {
	return api.CreateDispTypeInfo(idata)
}

func CopyMemory(dest, src unsafe.Pointer, length uint32) {
	api.CopyMemory(dest, src, length)
}

func GetUserDefaultLocaleID() uint32 {
	return api.GetUserDefaultLocaleID()
}

func GetMessage(msg *types.Msg, hwnd, MsgFilterMin, MsgFilterMax uint32) (int32, error) {
	return api.GetMessage(msg, hwnd, MsgFilterMin, MsgFilterMax)
}

func DispatchMessage(msg *types.Msg) int32 {
	return api.DispatchMessage(msg)
}
