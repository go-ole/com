package com

import (
	"fmt"

	"github.com/jacobsantos/go-com/iunknown"
)

func ExampleCoInitialize() {
	CoInitialize()
	defer CoUninitialize()

	var err error
	var unknown iunknown.Unknown

	if classID, err := ClassIDFromProgramID("msxml2.domdocument"); err != nil {
		err := CoCreateInstance(classID, iunknown.InterfaceID, &unknown)
		if err != nil {
			unknown.Release()
		}
	}
	// Output:
}

func ExampleBindMoniker() {
	// TODO: Complete.
}

func ExampleClassIDFromProgramID_msxml2() {
	iid, err := ClassIDFromProgramID("msxml2.domdocument")
	fmt.Printf("%v", iid)
	// Output:
}

func ExampleClassIDFromProgramID_msxml26() {
	iid, err := ClassIDFromProgramID("msxml2.domdocument.6.0")
	fmt.Printf("%v", iid)
	// Output:
}

func ExampleClassIDFromProgramID_eventsystem() {
	iid, err := ClassIDFromProgramID("eventsystem.eventclass.1")
	fmt.Printf("%v", iid)
	// Output:
}

func ExampleClassIDFromString_iunknown() {
	classID, err := ClassIDFromString("{00000000-0000-0000-C000-000000000046}")
	fmt.Printf("%v", classID)
	// Output:
	// GUID{
	//     0x00000000,
	//     0x0000,
	//     0x0000,
	//     [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
}

func ExampleClassIDFromString_idispatch() {
	classID, err := ClassIDFromString("{00020400-0000-0000-C000-000000000046}")
	fmt.Printf("%v", classID)
	// Output:
	// GUID{
	//     0x00020400,
	//     0x0000,
	//     0x0000,
	//     [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
}
