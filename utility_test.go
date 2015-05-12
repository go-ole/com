package com

import (
	"fmt"
	"testing"
	"unsafe"
)

func ExamplePointerToByteArray_bytes() {
	expected := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
	actual := PointerToByteArray(uintptr(unsafe.Pointer(expected)), len(expected), TypeOf(expected))
	if expected == actual {
		fmt.Printf("Expected equals actual.")
	}
}

func TestPointerToByteArray_bytes(t *testing.T) {
	// XXX: This should be changed to use testing/quick generator.
	expected := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
	actual := PointerToByteArray(uintptr(unsafe.Pointer(expected)), len(expected), TypeOf(expected))

	if expected != actual {
		t.Fail()
	}
}

func ExamplePointerToArrayAppend_bytes() {
	expected := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
	var actual []byte

	PointerToArrayAppend(uintptr(unsafe.Pointer(expected)), len(expected), &actual)

	if expected == actual {
		fmt.Printf("Expected equals actual.")
	}
}

func TestPointerToArrayAppend_bytes(t *testing.T) {
	expected := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
	var actual []byte

	err := PointerToArrayAppend(uintptr(unsafe.Pointer(expected)), len(expected), &actual)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if expected != actual {
		t.Fail()
	}
}

func ExamplePointerToArray_bytes() {
	expected := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}

	actual, err := PointerToArray(uintptr(unsafe.Pointer(expected)), len(expected), expected)

	if expected == actual {
		fmt.Printf("Expected equals actual.")
	}
}

func TestPointerToArray_bytes(t *testing.T) {
	expected := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}

	raw, err := PointerToArray(uintptr(unsafe.Pointer(expected)), len(expected), expected)
	actual, ok := raw.([]byte)

	if !ok {
		t.Error("Failed to convert to byte array")
		t.FailNow()
	}

	if expected != actual {
		t.Fail()
	}
}
