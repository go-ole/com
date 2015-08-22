package com

import "reflect"

// ServerObject is the structure for serving Go objects.
type ServerObject struct {
	VirtualTable   *interface{}
	ReferenceCount int32
	InterfaceID    *GUID
	Interface      interface{}
	FunctionMap    map[string]int32
}

// Service creates COM server reference object.
func Service(obj *interface{}, vt *interface{}, interfaceID *GUID) (*ServerObject, error) {
	rv := reflect.ValueOf(obj).Elem()
	if rv.Type().Kind() == reflect.Struct {
		service := &ServerObject{}
		service.VirtualTable = vt
		service.Interface = obj
		service.InterfaceID = interfaceID
		return service, nil
	}

	return nil, NewError(InvalidArgumentErrorCode)
}
