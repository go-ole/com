// +build windows,!cgo

package safearray

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
	procSafeArrayAccessData, _        = modoleaut32.FindProc("SafeArrayAccessData")
	procSafeArrayAllocData, _         = modoleaut32.FindProc("SafeArrayAllocData")
	procSafeArrayAllocDescriptor, _   = modoleaut32.FindProc("SafeArrayAllocDescriptor")
	procSafeArrayAllocDescriptorEx, _ = modoleaut32.FindProc("SafeArrayAllocDescriptorEx")
	procSafeArrayCopy, _              = modoleaut32.FindProc("SafeArrayCopy")
	procSafeArrayCopyData, _          = modoleaut32.FindProc("SafeArrayCopyData")
	procSafeArrayCreate, _            = modoleaut32.FindProc("SafeArrayCreate")
	procSafeArrayCreateEx, _          = modoleaut32.FindProc("SafeArrayCreateEx")
	procSafeArrayCreateVector, _      = modoleaut32.FindProc("SafeArrayCreateVector")
	procSafeArrayCreateVectorEx, _    = modoleaut32.FindProc("SafeArrayCreateVectorEx")
	procSafeArrayDestroy, _           = modoleaut32.FindProc("SafeArrayDestroy")
	procSafeArrayDestroyData, _       = modoleaut32.FindProc("SafeArrayDestroyData")
	procSafeArrayDestroyDescriptor, _ = modoleaut32.FindProc("SafeArrayDestroyDescriptor")
	procSafeArrayGetDim, _            = modoleaut32.FindProc("SafeArrayGetDim")
	procSafeArrayGetElement, _        = modoleaut32.FindProc("SafeArrayGetElement")
	procSafeArrayGetElemsize, _       = modoleaut32.FindProc("SafeArrayGetElemsize")
	procSafeArrayGetIID, _            = modoleaut32.FindProc("SafeArrayGetIID")
	procSafeArrayGetLBound, _         = modoleaut32.FindProc("SafeArrayGetLBound")
	procSafeArrayGetUBound, _         = modoleaut32.FindProc("SafeArrayGetUBound")
	procSafeArrayGetVartype, _        = modoleaut32.FindProc("SafeArrayGetVartype")
	procSafeArrayLock, _              = modoleaut32.FindProc("SafeArrayLock")
	procSafeArrayPtrOfIndex, _        = modoleaut32.FindProc("SafeArrayPtrOfIndex")
	procSafeArrayUnaccessData, _      = modoleaut32.FindProc("SafeArrayUnaccessData")
	procSafeArrayUnlock, _            = modoleaut32.FindProc("SafeArrayUnlock")
	procSafeArrayPutElement, _        = modoleaut32.FindProc("SafeArrayPutElement")
	procSafeArrayRedim, _             = modoleaut32.FindProc("SafeArrayRedim")
	procSafeArraySetIID, _            = modoleaut32.FindProc("SafeArraySetIID")
	procSafeArrayGetRecordInfo, _     = modoleaut32.FindProc("SafeArrayGetRecordInfo")
	procSafeArraySetRecordInfo, _     = modoleaut32.FindProc("SafeArraySetRecordInfo")
)

// AccessData returns raw array.
//
// AKA: SafeArrayAccessData in Windows API.
func AccessData(array *types.COMArray) (element uintptr, err error) {
	err = errors.HResultToError(procSafeArrayAccessData.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(unsafe.Pointer(&element))))
	return
}

// UnaccessData releases raw array.
//
// AKA: SafeArrayUnaccessData in Windows API.
func UnaccessData(array *types.COMArray) (err error) {
	err = errors.HResultToError(procSafeArrayUnaccessData.Call(uintptr(unsafe.Pointer(array))))
	return
}

// AllocateArrayData allocates SafeArray.
//
// AKA: SafeArrayAllocData in Windows API.
func AllocateArrayData(array *types.COMArray) (err error) {
	err = errors.HResultToError(procSafeArrayAllocData.Call(uintptr(unsafe.Pointer(array))))
	return
}

// AllocateArrayDescriptor allocates SafeArray.
//
// AKA: SafeArrayAllocDescriptor in Windows API.
func AllocateArrayDescriptor(dimensions uint32) (array *types.COMArray, err error) {
	err = errors.HResultToError(procSafeArrayAllocDescriptor.Call(
		uintptr(dimensions),
		uintptr(unsafe.Pointer(&array))))
	return
}

// AllocateArrayDescriptorEx allocates SafeArray.
//
// AKA: SafeArrayAllocDescriptorEx in Windows API.
func AllocateArrayDescriptorEx(variantType types.VariantType, dimensions uint32) (array *types.COMArray, err error) {
	err = errors.HResultToError(procSafeArrayAllocDescriptorEx.Call(
		uintptr(variantType),
		uintptr(dimensions),
		uintptr(unsafe.Pointer(&array))))
	return
}

// Duplicate returns copy of SafeArray.
//
// AKA: SafeArrayCopy in Windows API.
func Duplicate(original *types.COMArray) (array *types.COMArray, err error) {
	err = errors.HResultToError(procSafeArrayCopy.Call(
		uintptr(unsafe.Pointer(original)),
		uintptr(unsafe.Pointer(&array))))
	return
}

// DuplicateData duplicates SafeArray into another SafeArray object.
//
// AKA: SafeArrayCopyData in Windows API.
func DuplicateData(original, duplicate *types.COMArray) (err error) {
	err = errors.HResultToError(procSafeArrayCopyData.Call(
		uintptr(unsafe.Pointer(original)),
		uintptr(unsafe.Pointer(&duplicate))))
	return
}

// CreateArray creates SafeArray.
//
// AKA: SafeArrayCreate in Windows API.
func CreateArray(variantType types.VariantType, dimensions uint32, bounds *types.Bounds) (array *types.COMArray, err error) {
	sa, _, err := procSafeArrayCreate.Call(
		uintptr(variantType),
		uintptr(dimensions),
		uintptr(unsafe.Pointer(bounds)))
	array = (*types.COMArray)(unsafe.Pointer(&sa))
	return
}

// CreateArrayEx creates SafeArray.
//
// AKA: SafeArrayCreateEx in Windows API.
func CreateArrayEx(variantType types.VariantType, dimensions uint32, bounds *types.Bounds, extra uintptr) (array *types.COMArray, err error) {
	sa, _, err := procSafeArrayCreateEx.Call(
		uintptr(variantType),
		uintptr(dimensions),
		uintptr(unsafe.Pointer(bounds)),
		extra)
	array = (*types.COMArray)(unsafe.Pointer(sa))
	return
}

// CreateArrayVector creates SafeArray.
//
// AKA: SafeArrayCreateVector in Windows API.
func CreateArrayVector(variantType types.VariantType, lowerBound int32, length uint32) (array *types.COMArray, err error) {
	sa, _, err := procSafeArrayCreateVector.Call(
		uintptr(variantType),
		uintptr(lowerBound),
		uintptr(length))
	array = (*types.COMArray)(unsafe.Pointer(sa))
	return
}

func CreateArrayVectorEx(variantType types.VariantType, lowerBound int32, length uint32, extra uintptr) (array *types.COMArray, err error) {
	sa, _, err := procSafeArrayCreateVectorEx.Call(
		uintptr(variantType),
		uintptr(lowerBound),
		uintptr(length),
		extra)
	array = (*types.COMArray)(unsafe.Pointer(sa))
	return
}

func Destroy(array *types.COMArray) error {
	return errors.HResultToError(procSafeArrayDestroy.Call(uintptr(unsafe.Pointer(array))))
}

func DestroyData(array *types.COMArray) error {
	return errors.HResultToError(procSafeArrayDestroyData.Call(uintptr(unsafe.Pointer(array))))
}

func DestroyDescriptor(array *types.COMArray) error {
	return errors.HResultToError(procSafeArrayDestroyDescriptor.Call(uintptr(unsafe.Pointer(array))))
}

func GetDimensions(array *types.COMArray) (dimensions uint32, err error) {
	l, _, err := procSafeArrayGetDim.Call(uintptr(unsafe.Pointer(array)))
	dimensions = *(*uint32)(unsafe.Pointer(l))
	return
}

func GetElementSize(array *types.COMArray) (length uint32, err error) {
	l, _, err := procSafeArrayGetElemsize.Call(uintptr(unsafe.Pointer(array)))
	length = *(*uint32)(unsafe.Pointer(l))
	return
}

func GetElement(array *types.COMArray, index int64) (element interface{}, err error) {
	err = GetElementDirect(array, index, &element)
	return
}

func GetElementDirect(array *types.COMArray, index int64, element interface{}) (err error) {
	err = errors.HResultToError(procSafeArrayGetElement.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(index),
		uintptr(unsafe.Pointer(&element))))
	return
}

func GetInterfaceID(array *types.COMArray) (guid *types.GUID, err error) {
	err = errors.HResultToError(procSafeArrayGetIID.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(unsafe.Pointer(&guid))))
	return
}

func GetLowerBound(array *types.COMArray, dimension uint32) (lowerBound int64, err error) {
	err = errors.HResultToError(procSafeArrayGetLBound.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(dimension),
		uintptr(unsafe.Pointer(&lowerBound))))
	return
}

func GetUpperBound(array *types.COMArray, dimension uint32) (upperBound int64, err error) {
	err = errors.HResultToError(procSafeArrayGetUBound.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(dimension),
		uintptr(unsafe.Pointer(&upperBound))))
	return
}

func GetVariantType(array *types.COMArray) (varType types.VariantType, err error) {
	var vt uint16
	err = errors.HResultToError(procSafeArrayGetVartype.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(unsafe.Pointer(&vt))))
	varType = com.VariantType(vt)
	return
}

func Lock(array *types.COMArray) (err error) {
	err = errors.HResultToError(procSafeArrayLock.Call(uintptr(unsafe.Pointer(array))))
	return
}

func Unlock(array *types.COMArray) (err error) {
	err = errors.HResultToError(procSafeArrayUnlock.Call(uintptr(unsafe.Pointer(array))))
	return
}

func GetPointerOfIndex(array *types.COMArray, index int64) (ref uintptr, err error) {
	err = errors.HResultToError(procSafeArrayPtrOfIndex.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(index),
		uintptr(unsafe.Pointer(&ref))))
	return
}

func SetInterfaceID(array *types.COMArray, interfaceID *types.GUID) (err error) {
	err = errors.HResultToError(procSafeArraySetIID.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(unsafe.Pointer(interfaceID))))
	return
}

func ResetDimensions(array *types.COMArray, bounds *types.Bounds) error {
	err = errors.HResultToError(procSafeArrayRedim.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(unsafe.Pointer(bounds))))
	return
}

func PutElement(array *types.COMArray, index int64, element interface{}) (err error) {
	err = errors.HResultToError(procSafeArrayPutElement.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(index),
		uintptr(unsafe.Pointer(&element))))
	return
}

func GetRecordInfo(array *types.COMArray) (recordInfo interface{}, err error) {
	err = errors.HResultToError(procSafeArrayGetRecordInfo.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(unsafe.Pointer(&recordInfo))))
	return
}

func SetRecordInfo(array *types.COMArray, recordInfo interface{}) (err error) {
	err = errors.HResultToError(procSafeArraySetRecordInfo.Call(
		uintptr(unsafe.Pointer(array)),
		uintptr(unsafe.Pointer(recordInfo))))
	return
}
