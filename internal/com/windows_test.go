// +build windows,!cgo

package com

import (
	"fmt"
	"testing"

	"github.com/go-ole/go-ole/types/iunknown"
)

var WindowsMediaNSSManagerClassID = &GUID{0x92498132, 0x4D1A, 0x4297, [8]byte{0x9B, 0x78, 0x9E, 0x2E, 0x4B, 0xA9, 0x9C, 0x07}}

func TestComSetupAndShutDown(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Fail()
		}
	}()

	CoInitialize()
	CoUninitialize()
}

func TestComPublicSetupAndShutDown(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Fail()
		}
	}()

	CoInitialize()
	CoUninitialize()
}

func TestComExSetupAndShutDown(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Fail()
		}
	}()

	CoInitializeEx(MultithreadedConcurrencyModel)
	CoUninitialize()
}

func TestComPublicExSetupAndShutDown(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Fail()
		}
	}()

	CoInitializeEx(MultithreadedConcurrencyModel)
	CoUninitialize()
}

func TestClsidFromProgID_WindowsMediaNSSManager(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Fail()
		}
	}()

	CoInitialize()
	defer CoUninitialize()
	actual, err := ClassIDFromProgramID("WMPNSSCI.NSSManager")
	if err == nil {
		if !IsEqualGUID(WindowsMediaNSSManagerClassID, actual) {
			t.Log(err)
			t.Log(fmt.Sprintf("Actual GUID: %+v\n", actual))
			t.Fail()
		}
	}
}

func TestClsidFromString_WindowsMediaNSSManager(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Fail()
		}
	}()

	CoInitialize()
	defer CoUninitialize()
	actual, err := ClassIDFromString("{92498132-4D1A-4297-9B78-9E2E4BA99C07}")

	if !IsEqualGUID(WindowsMediaNSSManagerClassID, actual) {
		t.Log(err)
		t.Log(fmt.Sprintf("Actual GUID: %+v\n", actual))
		t.Fail()
	}
}

func TestCreateInstance_WindowsMediaNSSManager(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Fail()
		}
	}()

	CoInitialize()
	defer CoUninitialize()
	actual, err := ClassIDFromProgramID("WMPNSSCI.NSSManager")

	if err == nil {
		if !IsEqualGUID(WindowsMediaNSSManagerClassID, actual) {
			t.Log(err)
			t.Log(fmt.Sprintf("Actual GUID: %+v\n", actual))
			t.Fail()
		}

		unknown, err := CoCreateInstance(actual, iunknown.InterfaceID)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		unknown.Release()
	}
}

func TestError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Fail()
		}
	}()

	CoInitialize()
	defer CoUninitialize()
	_, err := ClassIDFromProgramID("INTERFACE-NOT-FOUND")
	if err == nil {
		t.Fatalf("should be fail", err)
	}

	switch vt := err.(type) {
	case *OleError:
	default:
		t.Fatalf("should be *OleError %t", vt)
	}
}
