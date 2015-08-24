// +build !windows

package safearray

import (
	"unsafe"

	"github.com/go-ole/com/errors"
	"github.com/go-ole/com/types"
)

func AccessData(array *types.COMArray) (uintptr, error) {
	return uintptr(unsafe.Pointer(array)), errors.NotImplementedError
}

func UnaccessData(array *types.COMArray) error {
	return errors.NotImplementedError
}

func AllocateArrayData(array *types.COMArray) error {
	return errors.NotImplementedError
}

func AllocateArrayDescriptor(dimensions uint32) (*types.COMArray, error) {
	return nil, errors.NotImplementedError
}

func AllocateArrayDescriptorEx(variantType types.VariantType, dimensions uint32) (*types.COMArray, error) {
	return nil, errors.NotImplementedError
}

func Duplicate(original *types.COMArray) (*types.COMArray, error) {
	return nil, errors.NotImplementedError
}

func DuplicateData(original, duplicate *types.COMArray) error {
	return errors.NotImplementedError
}

func CreateArray(variantType types.VariantType, dimensions uint32, bounds *types.Bounds) (*types.COMArray, error) {
	return nil, errors.NotImplementedError
}

func CreateArrayEx(variantType types.VariantType, dimensions uint32, bounds *types.Bounds, extra uintptr) (*types.COMArray, error) {
	return nil, errors.NotImplementedError
}

func CreateArrayVector(variantType types.VariantType, lowerBound int32, length uint32) (*types.COMArray, error) {
	return nil, errors.NotImplementedError
}

func CreateArrayVectorEx(variantType types.VariantType, lowerBound int32, length uint32, extra uintptr) (*types.COMArray, error) {
	return nil, errors.NotImplementedError
}

func Destroy(array *types.COMArray) error {
	return errors.NotImplementedError
}

func DestroyData(array *types.COMArray) error {
	return errors.NotImplementedError
}

func DestroyDescriptor(array *types.COMArray) error {
	return errors.NotImplementedError
}

func GetDimensions(array *types.COMArray) (uint32, error) {
	return *uint32(0), errors.NotImplementedError
}

func GetElementSize(array *types.COMArray) (uint32, error) {
	return *uint32(0), errors.NotImplementedError
}

func GetElement(array *types.COMArray, index int64) (interface{}, error) {
	return nil, errors.NotImplementedError
}

func GetElementDirect(array *types.COMArray, index int64, element interface{}) error {
	return NotImplementedError
}

func GetInterfaceID(array *types.COMArray) (*types.GUID, error) {
	return nil, errors.NotImplementedError
}

func GetLowerBound(array *types.COMArray, dimension uint32) (int64, error) {
	return int64(0), errors.NotImplementedError
}

func GetUpperBound(array *types.COMArray, dimension uint32) (int64, error) {
	return int64(0), errors.NotImplementedError
}

func GetVariantType(array *types.COMArray) (types.VariantType, error) {
	return types.NullVariantType, errors.NotImplementedError
}

func Lock(array *types.COMArray) error {
	return errors.NotImplementedError
}

func Unlock(array *types.COMArray) error {
	return errors.NotImplementedError
}

func GetPointerOfIndex(array *types.COMArray, index int64) (uintptr, error) {
	return uintptr(0), errors.NotImplementedError
}

func SetInterfaceID(array *types.COMArray, interfaceID *types.GUID) error {
	return errors.NotImplementedError
}

func ResetDimensions(array *types.COMArray, bounds *types.Bounds) error {
	return errors.NotImplementedError
}

func PutElement(array *types.COMArray, index int64, element interface{}) error {
	return errors.NotImplementedError
}

func GetRecordInfo(array *types.COMArray) (interface{}, error) {
	return nil, errors.NotImplementedError
}

func SetRecordInfo(array *types.COMArray, recordInfo interface{}) (err error) {
	return errors.NotImplementedError
}
