package com

import (
	"github.com/jacobsantos/go-com/idispatch"
	"github.com/jacobsantos/go-com/iunknown"
)

// ClassIDFrom retrieves class ID whether given is program ID or application string.
func ClassIDFrom(programID string) (classID *GUID, err error) {
	classID, err = ClassIDFromProgramID(programID)
	if err != nil {
		classID, err = ClassIDFromString(programID)
		if err != nil {
			return
		}
	}
	return
}

// InterfaceIDFrom retrieves Interface ID from interface type.
//
// Only supports IUnknown and IDispatch.
//
// TODO: support more COM interfaces.
func InterfaceIDFrom(obj interface{}) *GUID {
	var interfaceID *GUID
	switch obj.(type) {
	case iunknown.Unknown, *iunknown.Unknown:
		interfaceID = IUnknownInterfaceID
	case idispatch.Dispatch, *idispatch.Dispatch:
		interfaceID = IDispatchInterfaceID
	default:
		interfaceID = IUnknownInterfaceID
	}

	return interfaceID
}

// LoadObject loads COM object if in list.
func LoadObject(obj interface{}, programID ...string) []error {
	return LoadObjectWithCallback(CreateObject, &obj, programID...)
}

// LoadActiveObject loads COM active object if in list.
func LoadActiveObject(obj interface{}, programID ...string) []error {
	return LoadObjectWithCallback(GetActiveObject, &obj, programID...)
}

// LoadObjectWithCallback loads COM object if in list with callback.
func LoadObjectWithCallback(callback func(programID string, obj interface{}) error, obj interface{}, programID ...string) (err []error) {
	tempErrors := make([]error, len(names))
	var numErrors int
	for _, name := range names {
		err := callback(name, &obj)
		if err != nil {
			tempErrors = append(tempErrors, err)
			numErrors += 1
			continue
		}
		break
	}

	copy(errors, tempErrors[0:numErrors])
	return
}

// CreateObject creates object from programID based on interface type.
//
// Only supports IUnknown and IDispatch.
//
// Program ID can be either program ID or application string.
//
// Object needs to be passed by reference.
func CreateObject(programID string, obj interface{}) (err error) {
	classID, err := ClassIDFrom(programID)
	if err != nil {
		return
	}

	err = CoCreateInstance(classID, InterfaceIDFrom(&obj), &obj)
	if err != nil {
		return
	}
	return
}

// CreateObjectWithInterfaceID creates object from programID.
//
// Program ID can be either program ID or application string.
//
// Object needs to be passed by reference.
func CreateObjectWithInterfaceID(programID string, interfaceID *GUID, obj interface{}) (err error) {
	classID, err := ClassIDFrom(programID)
	if err != nil {
		return
	}

	err = CoCreateInstance(classID, interfaceID, &obj)
	if err != nil {
		return
	}
	return
}

// GetActiveObject retrieves active object for program ID and interface ID based
// on interface type.
//
// Only supports IUnknown and IDispatch.
//
// Program ID can be either program ID or application string.
//
// Object needs to be passed by reference.
func GetActiveObject(programID string, obj interface{}) (err error) {
	classID, err := ClassIDFrom(programID)
	if err != nil {
		return
	}

	err = GetActiveObject(classID, InterfaceIDFrom(&obj), &obj)
	if err != nil {
		return
	}
	return
}

// GetActiveObjectWithInterfaceID retrieves active object for program ID and
// interface ID.
//
// Program ID can be either program ID or application string.
//
// Object needs to be passed by reference.
func GetActiveObjectWithInterfaceID(programID string, interfaceID *GUID, obj interface{}) (err error) {
	classID, err := ClassIDFrom(programID)
	if err != nil {
		return
	}

	err = GetActiveObject(classID, interfaceID, &obj)
	if err != nil {
		return
	}
	return
}
