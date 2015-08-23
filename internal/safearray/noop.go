// +build !windows

package safearray

import "unsafe"
import "github.com/go-ole/com/errors"
import "github.com/go-ole/com/safearray"
import "github.com/go-ole/com/types"

func AccessData(array *safearray.COMArray) (uintptr, error) {
	return uintptr(unsafe.Pointer(array)), errors.NotImplementedError
}

func UnaccessData(array *safearray.COMArray) error {
	return errors.NotImplementedError
}

func AllocateArrayData(array *safearray.COMArray) error {
	return errors.NotImplementedError
}

func AllocateArrayDescriptor(dimensions uint32) (*safearray.COMArray, error) {
	return nil, errors.NotImplementedError
}

func AllocateArrayDescriptorEx(variantType types.VariantType, dimensions uint32) (*safearray.COMArray, error) {
	return nil, errors.NotImplementedError
}

func Duplicate(original *safearray.COMArray) (*safearray.COMArray, error) {
	return nil, errors.NotImplementedError
}

func DuplicateData(original, duplicate *safearray.COMArray) error {
	return errors.NotImplementedError
}

func CreateArray(variantType types.VariantType, dimensions uint32, bounds *safearray.Bounds) (*safearray.COMArray, error) {
	return nil, errors.NotImplementedError
}

func CreateArrayEx(variantType types.VariantType, dimensions uint32, bounds *safearray.Bounds, extra uintptr) (*safearray.COMArray, error) {
	return nil, errors.NotImplementedError
}

func CreateArrayVector(variantType types.VariantType, lowerBound int32, length uint32) (*safearray.COMArray, error) {
	return nil, errors.NotImplementedError
}

func CreateArrayVectorEx(variantType types.VariantType, lowerBound int32, length uint32, extra uintptr) (*safearray.COMArray, error) {
	return nil, errors.NotImplementedError
}

func Destroy(array *safearray.COMArray) error {
	return errors.NotImplementedError
}

func DestroyData(array *safearray.COMArray) error {
	return errors.NotImplementedError
}

func DestroyDescriptor(array *safearray.COMArray) error {
	return errors.NotImplementedError
}

func GetDimensions(array *safearray.COMArray) (uint32, error) {
	return *uint32(0), errors.NotImplementedError
}

func GetElementSize(array *safearray.COMArray) (uint32, error) {
	return *uint32(0), errors.NotImplementedError
}

func GetElement(array *safearray.COMArray, index int64) (interface{}, error) {
	return nil, errors.NotImplementedError
}

func GetElementDirect(array *safearray.COMArray, index int64, element interface{}) error {
	return NotImplementedError
}

func GetInterfaceID(array *safearray.COMArray) (*types.GUID, error) {
	return nil, errors.NotImplementedError
}

func GetLowerBound(array *safearray.COMArray, dimension uint32) (int64, error) {
	return int64(0), errors.NotImplementedError
}

func GetUpperBound(array *safearray.COMArray, dimension uint32) (int64, error) {
	return int64(0), errors.NotImplementedError
}

func GetVariantType(array *safearray.COMArray) (types.VariantType, error) {
	return types.NullVariantType, errors.NotImplementedError
}

func Lock(array *safearray.COMArray) error {
	return errors.NotImplementedError
}

func Unlock(array *safearray.COMArray) error {
	return errors.NotImplementedError
}

func GetPointerOfIndex(array *safearray.COMArray, index int64) (uintptr, error) {
	return uintptr(0), errors.NotImplementedError
}

func SetInterfaceID(array *safearray.COMArray, interfaceID *types.GUID) error {
	return errors.NotImplementedError
}

func ResetDimensions(array *safearray.COMArray, bounds *safearray.Bounds) error {
	return errors.NotImplementedError
}

func PutElement(array *safearray.COMArray, index int64, element interface{}) error {
	return errors.NotImplementedError
}

func GetRecordInfo(array *safearray.COMArray) (interface{}, error) {
	return nil, errors.NotImplementedError
}

func SetRecordInfo(array *safearray.COMArray, recordInfo interface{}) (err error) {
	return errors.NotImplementedError
}
