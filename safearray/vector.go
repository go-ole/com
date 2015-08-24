package safearray

import "github.com/go-ole/com/types"

func NewVector(variantType types.VariantType, slice []interface{}) (array *types.COMArray, err error) {
	array, err = CreateArrayVector(variantType, 0, uint32(len(slice)))
	if err != nil {
		return
	}

	return UnmarshalArray(array, slice)
}

func NewVectorWithFlags(variantType types.VariantType, flags types.SafeArrayMask, slice []interface{}) (array *types.COMArray, err error) {
	array, err = CreateArrayVectorEx(variantType, 0, uint32(len(slice)), uintptr(flags))
	if err != nil {
		return
	}

	return UnmarshalArray(array, slice)
}

// FromByteArray creates SafeArray from byte array.
func FromByteArray(slice []byte) (array *COMArray, err error) {
	return NewVector(types.UInteger8VariantType, slice)
}
