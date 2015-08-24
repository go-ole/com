package safearray

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/go-ole/com"
	"github.com/go-ole/idispatch"
	"github.com/go-ole/iunknown"
)

// This tests more than one function. It tests all of the functions needed in
// order to retrieve a SafeArray populated with Strings.
func ExampleGetElementString_quickbooks() {
	com.CoInitialize()
	defer com.CoUninitialize()

	clsid, err := com.ClassIDFromProgramID("QBXMLRP2.RequestProcessor.1")
	if err != nil {
		if err.(*com.OleError).Code() == com.COMObjectClassStringErrorCode {
			return
		}
		t.Log(err)
		t.FailNow()
	}

	var unknown *iunknown.Unknown
	var dispatch *idispatch.Dispatch

	err := com.CoCreateInstance(clsid, iunknown.InterfaceID, &unknown)
	if err != nil {
		return
	}
	defer unknown.Release()

	err := unknown.QueryInterface(idispatch.InterfaceID, &dispatch)
	if err != nil {
		return
	}

	var result *com.Variant
	_, err = idispatch.CallMethod(dispatch, "OpenConnection2", "", "Test Application 1", 1)
	if err != nil {
		return
	}

	result, err = idispatch.CallMethod(dispatch, "BeginSession", "", 2)
	if err != nil {
		return
	}

	ticket := result.ToString()

	result, err = idispatch.GetProperty(dispatch, "QBXMLVersionsForSession", ticket)
	if err != nil {
		return
	}

	//
	// Example begins.
	//

	var qbXMLVersions *COMArray
	var qbXMLVersionStrings []string
	qbXMLVersions = result.ToArray().Array

	// Release Safe Array memory
	defer Destroy(qbXMLVersions)

	// Get array bounds
	var LowerBounds int64
	var UpperBounds int64
	LowerBounds, err = GetLowerBound(qbXMLVersions, 1)
	if err != nil {
		return
	}

	UpperBounds, err = GetUpperBound(qbXMLVersions, 1)
	if err != nil {
		return
	}

	totalElements := UpperBounds - LowerBounds + 1
	qbXMLVersionStrings = make([]string, totalElements)

	for i := int64(0); i < totalElements; i++ {
		qbXMLVersionStrings[int32(i)], _ = GetElementString(qbXMLVersions, i)
	}

	//
	// Example ends.
	//

	result, err = idispatch.CallMethod(dispatch, "EndSession", ticket)
	if err != nil {
		return
	}

	result, err = idispatch.CallMethod(dispatch, "CloseConnection")
	if err != nil {
		return
	}
}

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
