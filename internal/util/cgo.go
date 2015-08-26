// +build windows,cgo

package util

import (
	"unsafe"

	"github.com/go-ole/com/errors"
	"github.com/go-ole/com/types"
	"github.com/go-ole/types/idispatch"
	"github.com/go-ole/types/iunknown"
)

func SysAllocString(v string) *int16 {
	return nil
}

func SysAllocStringLen(v string) *int16 {
	return nil
}

func SysFreeString(v *int16) error {
	return errors.NotImplementedError
}

func SysStringLen(v *int16) uint32 {
	return uint32(0)
}

func CreateStdDispatch(unk *iunknown.IUnknown, v uintptr, ptinfo *iunknown.IUnknown) (*idispatch.Dispatch, error) {
	return nil, errors.NotImplementedError
}

func CreateDispTypeInfo(idata *INTERFACEDATA) (*iunknown.IUnknown, error) {
	return nil, errors.NotImplementedError
}

func CopyMemory(dest unsafe.Pointer, src unsafe.Pointer, length uint32) {
	// noop
}

func GetUserDefaultLocaleID() uint32 {
	return 0
}

func GetMessage(msg *types.Msg, hwnd, MsgFilterMin, MsgFilterMax uint32) (int32, error) {
	return 0, errors.NotImplementedError
}

func DispatchMessage(msg *types.Msg) int32 {
	return 0
}
