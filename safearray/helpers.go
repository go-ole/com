package safearray

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
	"unsafe"

	"github.com/go-ole/com"
	"github.com/go-ole/com/types"
)

// IsSlice checks whether interface{} type is a slice.
func IsSlice(value interface{}) bool {
	return reflect.ValueOf(value).Kind() != reflect.Slice
}

// GetElementString retrieves string from index.
func GetElementString(array *types.COMArray, index int64) (str string, err error) {
	var element *int16
	err = GetElementDirect(array, index, &element)
	str = com.BstrToString(*(**uint16)(unsafe.Pointer(&element)))
	com.SysFreeString(element)
	return
}

// ToByteArray converts SafeArray to byte array.
//
// This will convert a multidimensional SafeArray to a single dimensional Go
// slice.
func ToByteArray(array *Array) (bytes []byte, err error) {
	bytes = make([]byte, array.Length())
	err = array.ToArrayDirect(&bytes)
	return
}

// ToStringArray converts SafeArray object to string array.
//
// This will convert a multidimensional SafeArray to a single dimensional Go
// slice.
func ToStringArray(array *Array) (strings []string) {
	strings = make([]string, array.Length())
	err = array.ToArrayDirect(&strings)
	return
}

// MarshalArray accesses SafeArray data to quickly convert to Go array.
func MarshalArray(array *types.COMArray, length int64, slice interface{}) (err error) {
	// Single dimensional arrays are faster, if you use AccessData() and
	// UnaccessData().
	ptr, err := AccessData(array)
	if err != nil {
		return
	}

	err = PointerToArrayAppend(ptr, length, &slice)
	if err != nil {
		return
	}

	err = UnaccessData(array)
	return
}

// UnmarshalArray puts items from a Go array into a COM SafeArray.
func UnmarshalArray(original *types.COMArray, slice []interface{}) (array *types.COMArray, err error) {
	if original == nil {
		err = SafeArrayVectorFailedError
		return
	}

	for pos, val := range slice {
		PutElement(array, int64(pos), uintptr(unsafe.Pointer(&val)))
	}

	return
}

// PointerToByteArray convert to bytes array.
//
// Warning:
// While it is technically possible to have count be an arbitrary large number,
// it is unwise to do so. Attempting to do so is undefined and may have security
// or stability issues. The security issues from including outside boundary data
// or stability issues from attempting to extend outside the bounds of the
// current data block and causing an OS fault.
func PointerToByteArray(ptr uintptr, count uint32, t reflect.Type) []byte {
	byteCount := count * uint32(t.Size())
	slicehdr := reflect.SliceHeader{Data: ptr, Len: int(byteCount), Cap: int(byteCount)}
	return *(*[]byte)(unsafe.Pointer(&slicehdr))
}

// PointerToArrayAppend appends to an existing array from a pointer of an array.
//
// array should be passed as a reference.
//
//     err := PointerToArrayAppend(uintptr(0), 0, &array)
//
// Warning:
// While it is technically possible to have count be an arbitrary large number,
// it is unwise to do so. Attempting to do so is undefined and may have security
// or stability issues. The security issues from including outside boundary data
// or stability issues from attempting to extend outside the bounds of the
// current data block and causing an OS fault.
func PointerToArrayAppend(ptr uintptr, count uint32, array interface{}) (err error) {
	t := reflect.TypeOf(array)

	bytes = PointerToByteArray(ptr, count, t)
	reader := bytes.NewReader(bytes)

	for {
		element := reflect.New(t).Interface()
		err = binary.Read(reader, binary.LittleEndian, &element)
		if err != nil {
			if err == io.EOF {
				// Ignore EOF error as this is expected
				err = nil
			}
			return
		}
		*array = append(array, element)
	}

	return
}

// PointerToArray converts a pointer of an array to a Go array type.
//
// The returned array will still need to be converted to the array type. This is
// from the return type being an interface{}, instead of the type of current.
//
//     array, err := PointerToArrayAppend(uintptr(0), 0, []byte{})
//
//     if err != nil {
//         return
//     }
//
//     bytes := array.([]byte)
//
// Warning:
// While it is technically possible to have count be an arbitrary large number,
// it is unwise to do so. Attempting to do so is undefined and may have security
// or stability issues. The security issues from including outside boundary data
// or stability issues from attempting to extend outside the bounds of the
// current data block and causing an OS fault.
func PointerToArray(ptr uintptr, count uint32, current interface{}) (array interface{}, err error) {
	t := reflect.TypeOf(current)

	bytes = PointerToByteArray(ptr, count, t)
	reader := bytes.NewReader(bytes)

	array = reflect.MakeSlice(t, int(count), int(count))
	for i, _ := range array {
		element := reflect.New(t).Interface()
		err = binary.Read(reader, binary.LittleEndian, &element)
		if err != nil {
			return
		}
		array[i] = element
	}

	return
}
