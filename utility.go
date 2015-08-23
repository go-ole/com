package com

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// GetInt32FromCall retrieves Int32 from syscall.
func GetInt32FromCall(obj, method uintptr) int32 {
	ret, _, _ := syscall.Syscall(method, 1, obj, 0, 0)
	return int32(ret)
}

// BytePtrToString converts byte pointer to a Go string.
func BytePtrToString(p *byte) string {
	a := (*[10000]uint8)(unsafe.Pointer(p))
	i := 0
	for a[i] != 0 {
		i++
	}
	return string(a[:i])
}

// UTF16PtrToString is alias for LpOleStrToString.
//
// Kept for compatibility reasons.
func UTF16PtrToString(p *uint16) string {
	return LpOleStrToString(p)
}

// LpOleStrToString converts COM Unicode to Go string.
func LpOleStrToString(p *uint16) string {
	if p == nil {
		return ""
	}

	length := lpOleStrLen(p)
	a := make([]uint16, length)

	ptr := unsafe.Pointer(p)

	for i := 0; i < int(length); i++ {
		a[i] = *(*uint16)(ptr)
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
	}

	return string(utf16.Decode(a))
}

// BstrToString converts COM binary string to Go string.
func BstrToString(p *uint16) string {
	if p == nil {
		return ""
	}
	length := SysStringLen((*int16)(unsafe.Pointer(p)))
	a := make([]uint16, length)

	ptr := unsafe.Pointer(p)

	for i := 0; i < int(length); i++ {
		a[i] = *(*uint16)(ptr)
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
	}
	return string(utf16.Decode(a))
}

// lpOleStrLen returns the length of Unicode string.
func lpOleStrLen(p *uint16) (length int64) {
	if p == nil {
		return 0
	}

	ptr := unsafe.Pointer(p)

	for i := 0; ; i++ {
		if 0 == *(*uint16)(ptr) {
			length = int64(i)
			break
		}
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
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
