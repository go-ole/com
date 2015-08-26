// +build windows,!cgo

package variant

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/com/errors"
	"github.com/go-ole/com/types"
)

var (
	modoleaut32, _ = syscall.LoadDLL("oleaut32.dll")
)

var (
	procVariantInit, _         = modoleaut32.FindProc("VariantInit")
	procVariantClear, _        = modoleaut32.FindProc("VariantClear")
	procVariantChangeType, _   = modoleaut32.FindProc("VariantChangeType")
	procVariantChangeTypeEx, _ = modoleaut32.FindProc("VariantChangeTypeEx")
	procVariantCopy, _         = modoleaut32.FindProc("VariantCopy")
	procVariantCopyInd, _      = modoleaut32.FindProc("VariantCopyInd")
)

func VariantInit(v *types.Variant) error {
	return errors.HResultToError(procVariantInit.Call(uintptr(unsafe.Pointer(v))))
}

func VariantClear(v *types.Variant) error {
	return errors.HResultToError(procVariantClear.Call(uintptr(unsafe.Pointer(v))))
}

func VariantChangeType(source *types.Variant, flags uint16, vt types.VariantType) (destination *types.Variant, err error) {
	err = errors.HResultToError(procVariantChangeType.Call(
		uintptr(unsafe.Pointer(&destination)),
		uintptr(unsafe.Pointer(source)),
		uintptr(flags),
		uintptr(vt)))
	return
}

func VariantChangeTypeEx(source *types.Variant, locale uint32, flags uint16, vt types.VariantType) (destination *types.Variant, err error) {
	err = errors.HResultToError(procVariantChangeTypeEx.Call(
		uintptr(unsafe.Pointer(&destination)),
		uintptr(unsafe.Pointer(source)),
		uintptr(locale),
		uintptr(flags),
		uintptr(vt)))
	return
}

func VariantCopy(source *types.Variant) (*types.Variant, error) {
	return copyVariant(procVariantCopy, source)
}

func VariantCopyIndirect(source *types.Variant) (*types.Variant, error) {
	return copyVariant(procVariantCopyInd, source)
}

func copyVariant(p *syscall.Proc, source *types.Variant) (destination *types.Variant, err error) {
	err = errors.HResultToError(p.Call(
		uintptr(unsafe.Pointer(&destination)),
		uintptr(unsafe.Pointer(source))))
	return
}
