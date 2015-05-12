// +build windows,!cgo

package com

import (
	"syscall"
	"unsafe"
)

var (
	procVariantInit, _         = modoleaut32.FindProc("VariantInit")
	procVariantClear, _        = modoleaut32.FindProc("VariantClear")
	procVariantChangeType, _   = modoleaut32.FindProc("VariantChangeType")
	procVariantChangeTypeEx, _ = modoleaut32.FindProc("VariantChangeTypeEx")
	procVariantCopy, _         = modoleaut32.FindProc("VariantCopy")
	procVariantCopyInd, _      = modoleaut32.FindProc("VariantCopyInd")
)

// VariantInit initializes memory for variant objects.
func VariantInit(v *Variant) error {
	return HResultToError(procVariantInit.Call(uintptr(unsafe.Pointer(v))))
}

// VariantClear clears the memory of variant objects.
func VariantClear(v *Variant) error {
	return HResultToError(procVariantClear.Call(uintptr(unsafe.Pointer(v))))
}

// VariantChangeType changes variant type.
func VariantChangeType(source *Variant, flags uint16, vt VariantType) (destination *Variant, err error) {
	err = HResultToError(procVariantChangeType.Call(
		uintptr(unsafe.Pointer(&destination)),
		uintptr(unsafe.Pointer(source)),
		uintptr(flags),
		uintptr(vt)))
	return
}

// VariantChangeTypeEx changes variant type.
func VariantChangeTypeEx(source *Variant, locale uint32, flags uint16, vt VariantType) (destination *Variant, err error) {
	err = HResultToError(procVariantChangeTypeEx.Call(
		uintptr(unsafe.Pointer(&destination)),
		uintptr(unsafe.Pointer(source)),
		uintptr(locale),
		uintptr(flags),
		uintptr(vt)))
	return
}

// VariantCopy copies variant.
func VariantCopy(source *Variant) (*Variant, error) {
	return copyVariant(procVariantCopy, source)
}

// VariantCopyIndirect copies variant reference.
func VariantCopyIndirect(source *Variant) (*Variant, error) {
	return copyVariant(procVariantCopyInd, source)
}

// Implements VariantCopyIndirect() and VariantCopy() functions.
func copyVariant(p *syscall.Proc, source *Variant) (destination *Variant, err error) {
	err = HResultToError(p.Call(
		uintptr(unsafe.Pointer(&destination)),
		uintptr(unsafe.Pointer(source))))
	return
}
