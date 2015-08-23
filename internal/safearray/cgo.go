// #include "OleAuto.h"
// +build windows,cgo
// XXX: This needs to be tested and completed.
// This is incomplete and will not work. Really just a skeleton.

package safearray

import "C"

import "unsafe"
import "github.com/go-ole/com/errors"
import "github.com/go-ole/com/safearray"
import "github.com/go-ole/com/types"

func AccessData(array *safearray.COMArray) (element uintptr, err error) {
	err = errors.MaybeError(C.SafeArrayAccessData(array, unsafe.Pointer(&element)))
	return
}

func UnaccessData(array *safearray.COMArray) error {
	return errors.MaybeError(C.SafeArrayUnaccessData(array))
}

func AllocateArrayData(array *safearray.COMArray) error {
	return errors.MaybeError(C.SafeArrayAllocData(array))
}

func AllocateArrayDescriptor(dimensions uint32) (array *safearray.COMArray, err error) {
	err = errors.MaybeError(C.SafeArrayAllocDescriptor(dimensions, unsafe.Pointer(&array)))
	return
}

func AllocateArrayDescriptorEx(variantType types.VariantType, dimensions uint32) (array *safearray.COMArray, err error) {
	err = errors.MaybeError(C.SafeArrayAllocDescriptorEx(uint16(variantType), dimensions, unsafe.Pointer(&array)))
	return
}

func Duplicate(original *safearray.COMArray) (array *safearray.COMArray, err error) {
	err = errors.MaybeError(C.SafeArrayCopy(original, unsafe.Pointer(&array)))
	return
}

func DuplicateData(original, duplicate *safearray.COMArray) error {
	return errors.MaybeError(C.SafeArrayCopyData(original, unsafe.Pointer(&duplicate)))
}

func CreateArray(variantType types.VariantType, dimensions uint32, bounds *safearray.Bounds) (array *safearray.COMArray, err error) {
	sa, _, err := C.SafeArrayCreate(uint16(variantType), dimensions, bounds)
	array = (*safearray.COMArray)(unsafe.Pointer(&sa))
	return
}

func CreateArrayEx(variantType types.VariantType, dimensions uint32, bounds *safearray.Bounds, extra uintptr) (array *safearray.COMArray, err error) {
	sa, _, err := C.SafeArrayCreateEx(uint16(variantType), dimensions, bounds, extra)
	array = (*safearray.COMArray)(unsafe.Pointer(sa))
	return
}

func CreateArrayVector(variantType types.VariantType, lowerBound int32, length uint32) (array *safearray.COMArray, err error) {
	sa, _, err := C.SafeArrayCreateVector(uint16(variantType), lowerBound, length)
	array = (*safearray.COMArray)(unsafe.Pointer(sa))
	return
}

func CreateArrayVectorEx(variantType types.VariantType, lowerBound int32, length uint32, extra uintptr) (array *safearray.COMArray, err error) {
	sa, _, err := C.SafeArrayCreateVectorEx(uint16(variantType), lowerBound, length, extra)
	array = (*safearray.COMArray)(unsafe.Pointer(sa))
	return
}

func Destroy(array *safearray.COMArray) error {
	return errors.MaybeError(C.SafeArrayDestroy(array))
}

func DestroyData(array *safearray.COMArray) error {
	return errors.MaybeError(C.SafeArrayDestroyData(array))
}

func DestroyDescriptor(array *safearray.COMArray) error {
	return errors.MaybeError(C.SafeArrayDestroyDescriptor(array))
}

func GetDimensions(array *safearray.COMArray) (dimensions uint32, err error) {
	l, _, err := C.SafeArrayGetDim(array)
	dimensions = *(*uint32)(unsafe.Pointer(l))
	return
}

func GetElementSize(array *safearray.COMArray) (length uint32, err error) {
	l, _, err := C.SafeArrayGetElemsize(array)
	length = *(*uint32)(unsafe.Pointer(l))
	return
}

func GetElement(array *safearray.COMArray, index int64) (element interface{}, err error) {
	err = GetElementDirect(array, index, &element)
	return
}

func GetElementDirect(array *safearray.COMArray, index int64, element interface{}) error {
	return errors.MaybeError(C.SafeArrayGetElement(array, index, unsafe.Pointer(&element)))
}

func GetInterfaceID(array *safearray.COMArray) (guid *types.GUID, err error) {
	err = errors.MaybeError(C.SafeArrayGetIID(array, unsafe.Pointer(&guid)))
	return
}

func GetLowerBound(array *safearray.COMArray, dimension uint32) (lowerBound int64, err error) {
	err = errors.MaybeError(C.SafeArrayGetLBound(array, dimension, unsafe.Pointer(&lowerBound)))
	return
}

func GetUpperBound(array *safearray.COMArray, dimension uint32) (upperBound int64, err error) {
	err = errors.MaybeError(C.SafeArrayGetUBound(array, dimension, unsafe.Pointer(&upperBound)))
	return
}

func GetVariantType(array *safearray.COMArray) (varType uint16, err error) {
	err = errors.MaybeError(C.SafeArrayGetVartype(array, unsafe.Pointer(&varType)))
	return
}

func Lock(array *safearray.COMArray) error {
	return errors.MaybeError(C.SafeArrayLock(array))
}

func Unlock(array *safearray.COMArray) error {
	return errors.MaybeError(C.SafeArrayUnlock(array))
}

func GetPointerOfIndex(array *safearray.COMArray, index int64) (ref uintptr, err error) {
	err = errors.MaybeError(C.SafeArrayPtrOfIndex(array, index, unsafe.Pointer(&ref)))
	return
}

func SetInterfaceID(array *safearray.COMArray, interfaceID *types.GUID) error {
	return errors.MaybeError(C.SafeArraySetIID(array, interfaceID))
}

func ResetDimensions(array *safearray.COMArray, bounds *safearray.Bounds) error {
	return errors.MaybeError(C.SafeArrayRedim(array, bounds))
}

func PutElement(array *safearray.COMArray, index int64, element interface{}) error {
	return errors.MaybeError(C.SafeArrayPutElement(array, index, unsafe.Pointer(&element)))
}

func GetRecordInfo(array *safearray.COMArray) (recordInfo interface{}, err error) {
	err = errors.MaybeError(C.SafeArrayGetRecordInfo(array, unsafe.Pointer(&recordInfo)))
	return
}

func SetRecordInfo(array *safearray.COMArray, recordInfo interface{}) error {
	return errors.MaybeError(C.SafeArraySetRecordInfo(array, unsafe.Pointer(recordInfo)))
}
