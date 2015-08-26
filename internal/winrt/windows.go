// +build windows,!cgo

package winrt

import (
	"reflect"
	"syscall"
	"unicode/utf8"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/go-ole/com/errors"
	"github.com/go-ole/com/types"
	hstr "github.com/go-ole/com/winrt"
	"github.com/go-ole/types/iinspectable"
)

var modcombase = syscall.NewLazyDLL("combase.dll")

var (
	procRoInitialize              = modcombase.NewProc("RoInitialize")
	procRoActivateInstance        = modcombase.NewProc("RoActivateInstance")
	procRoGetActivationFactory    = modcombase.NewProc("RoGetActivationFactory")
	procWindowsCreateString       = modcombase.NewProc("WindowsCreateString")
	procWindowsDeleteString       = modcombase.NewProc("WindowsDeleteString")
	procWindowsGetStringRawBuffer = modcombase.NewProc("WindowsGetStringRawBuffer")
)

func RoInitialize(threadType uint32) error {
	return errors.MaybeError(procRoInitialize.Call(uintptr(threadType)))
}

func RoActivateInstance(classID string) (obj *iinspectable.Inspectable, err error) {
	hClassID, err := NewHString(classID)
	if err != nil {
		return nil, err
	}
	defer DeleteHString(hClassID)

	err = errors.MaybeError(procRoActivateInstance.Call(
		uintptr(unsafe.Pointer(hClassID)),
		uintptr(unsafe.Pointer(&obj))))
	return
}

func RoGetActivationFactory(classID string, interfaceID *types.GUID) (obj *iinspectable.Inspectable, err error) {
	hClassID, err := NewHString(classID)
	if err != nil {
		return nil, err
	}
	defer DeleteHString(hClassID)

	err = errors.MaybeError(procRoGetActivationFactory.Call(
		uintptr(unsafe.Pointer(hClassID)),
		uintptr(unsafe.Pointer(interfaceID)),
		uintptr(unsafe.Pointer(&obj))))
	return
}

func NewHString(s string) (hstring hstr.HString, err error) {
	u16 := windows.StringToUTF16Ptr(s)
	len := uint32(utf8.RuneCountInString(s))
	err = errors.MaybeError(procWindowsCreateString.Call(
		uintptr(unsafe.Pointer(u16)),
		uintptr(len),
		uintptr(unsafe.Pointer(&hstring))))
	return
}

func DeleteHString(hstring hstr.HString) error {
	return errors.MaybeError(procWindowsDeleteString.Call(uintptr(hstring)))
}

func HStringToString(hstring hstr.HString) string {
	var u16buf uintptr
	var u16len uint32
	u16buf, _, _ = procWindowsGetStringRawBuffer.Call(
		uintptr(hstring),
		uintptr(unsafe.Pointer(&u16len)))

	u16hdr := reflect.SliceHeader{Data: u16buf, Len: int(u16len), Cap: int(u16len)}
	u16 := *(*[]uint16)(unsafe.Pointer(&u16hdr))
	return syscall.UTF16ToString(u16)
}
