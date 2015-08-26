package util

import (
	"unsafe"

	api "github.com/go-ole/com/internal/util"
	"github.com/go-ole/com/types"
	"github.com/go-ole/types/idispatch"
	"github.com/go-ole/types/iunknown"
)

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
