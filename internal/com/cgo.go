// #include "objbase.h"
// +build windows,cgo

package com

import "C"

import (
	"reflect"
	"unsafe"

	"golang.org/x/sys/windows"
)

func BindMoniker(moniker uintptr, option int32, interfaceID *GUID, object interface{}) error {
	return MaybeError(C.BindMoniker(
		moniker,
		uintptr(0),
		uintptr(unsafe.Pointer(interfaceID)),
		reflect.ValueOf(object).UnsafeAddr()))
}

func ClassIDFromProgramID(programId string) (*GUID, error) {
	ptrProgramId, err := windows.UTF16PtrFromString(programId)
	if err != nil {
		return
	}

	err = MaybeError(C.CLSIDFromProgID(
		uintptr(unsafe.Pointer(ptrProgramId)),
		uintptr(unsafe.Pointer(&classId))))
	return
}

func ClassIDFromProgramIDEx(programId string) (*GUID, error) {
	ptrProgramId, err := windows.UTF16PtrFromString(programId)
	if err != nil {
		return
	}

	err = MaybeError(C.CLSIDFromProgIDEx(
		uintptr(unsafe.Pointer(ptrProgramId)),
		uintptr(unsafe.Pointer(&classId))))
	return
}

func ClassIDFromString(str string) (*GUID, error) {
	return nil, NotImplementedError
}

func CoAddRefServerProcess() (uint64, error) {
	return 0, NotImplementedError
}

func CoAllowSetForegroundWindow(unknown iunknown.IUnknown) error {
	return NotImplementedError
}

func CoAllowUnmarshalerClassID() {}

func CoCancelCall() {}

func CoCopyProxy() {}

func CoCreateFreeThreadedMarshaler() {}

func CoCreateGUID() {}

func CoCreateInstance() {}

func CoCreateInstanceEx() {}

func CoCreateInstanceFromApp() {}

func CoDisableCallCancellation() {}

func CoDisconnectContext() {}

func CoDisconnectObject() {}

func CoEnableCallCancellation() {}

func CoEnterApplicationThreadLifetimeLoop() {}

func CoFileTimeNow() {}

func CoFreeAllLibraries() {}

func CoFreeLibrary() {}

func CoFreeUnusedLibrariesEx() {}

func CoGetApartmentType() {}

func CoGetApplicationThreadReference() {}

func CoGetCallContext() {}

func CoGetCallerTID() {}

func CoGetCancelObject() {}

func CoGetClassObject() {}

func CoGetContextToken() {}

func CoGetCurrentLogicalThreadID() {}

func CoGetCurrentProcess() {}

func CoGetInstanceFromFile() {}

func CoGetInstanceFromIStorage() {}

func CoGetInterceptor() {}

func CoGetInterfaceAndReleaseStream() {}

func CoGetMalloc() {}

func CoGetMarshalSizeMax() {}

func CoGetObject() {}

func CoGetObjectContext() {}

func CoGetPSClassID() {}

func CoGetStandardMarshal() {}

func CoGetStandardMarshalEx() {}

func CoGetSystemSecurityPermissions() {}

func CoGetTreatAsClass() {}

func CoHandlePriorityEventsFromMessagePump() {}

func CoImpersonateClient() {}

func CoInitialize() error {
	// http://msdn.microsoft.com/en-us/library/windows/desktop/ms678543(v=vs.85).aspx
	// Suggests that no value should be passed to CoInitialized.
	return MaybeError(C.CoInitialize(uintptr(0)))
}

func CoInitializeEx(coinit uint32) error {
	// http://msdn.microsoft.com/en-us/library/windows/desktop/ms695279(v=vs.85).aspx
	// Suggests that the first parameter is not only optional but should always be NULL.
	return MaybeError(C.CoInitializeEx(uintptr(0), uintptr(coinit)))
}

func CoInitializeSecurity() {}

func CoInstall() {}

func CoInvalidateRemoteMachineBindings() {}

func CoIsHandlerConnected() {}

func CoIsOle1Class() {}

func CoLoadLibrary() {}

func CoLockObjectExternal() {}

func CoMarshalHResult() {}

func CoMarshalInterface() {}

func CoMarshalInterThreadInterfaceInStream() {}

func CoQueryAuthenticationServices() {}

func CoQueryClientBlanket() {}

func CoQueryProxyBlanket() {}

func CoRegisterClassObject() {}

func CoRegisterInitializeSpy() {}

func CoRegisterMallocSpy() {}

func CoRegisterMessageFilter() {}

func CoRegisterPSClassID() {}

func CoRegisterSurrogate() {}

func CoReleaseMarshalData() {}

func CoReleaseServerProcess() {}

func CoResumeClassObjects() {}

func CoRevertToSelf() {}

func CoRevokeClassObject() {}

func CoRevokeInitializeSpy() {}

func CoRevokeMallocSpy() {}

func CoSetCancelObject() {}

func CoSetMessageDispatcher() {}

func CoSetProxyBlanket() {}

func CoSuspendClassObjects() {}

func CoSwitchCallContext() {}

func CoTaskMemFree(memoryPointer uintptr) {
	// todo
}

func CoTaskMemRealloc() {}

func CoTestCancel() {}

func CoTreatAsClass() {}

func CoUninitialize() {
	C.CoUninitialize()
}

func CoUnmarshalHResult() {}

func CoUnmarshalInterface() {}

func CoWaitForMultipleHandles() {}

func CreateAntiMoniker() {}

func CreateAsyncBindContext() {}

func CreateBindContext() {}

func CreateClassMoniker() {}

func CreateFileMoniker() {}

func CreateGenericComposite() {}

func CreateItemMoniker() {}

func CreateObjectRefMoniker() {}

func CreatePointerMoniker() {}

func DLLCanUnloadNow() {}

func DLLDebugObjectRPCHook() {}

func DLLGetClassObject() {}

func DLLRegisterServer() {}

func DLLUnregisterServer() {}

func GetClassFile() {}

func GetRunningObjectTable() {}

func InterfaceIDFromString() {}

func Initialize() {}

func IsAccelerator() {}

func IsEqualClassID() {}

func IsEqualGUID() {}

func IsEqualInterfaceID() {}

func MkParseDisplayName() {}

func MonikerCommonPrefixWith() {}

func MonikerRelativePathTo() {}

func OleDoAutoConvert() {}

func OleGetAutoConvert() {}

func OleGetIconOfClass() {}

func OleGetIconOfFile() {}

func OleIconToCursor() {}

func OleRegGetMiscStatus() {}

func OleRegGetUserType() {}

func OleSetAutoConvert() {}

func ProgramIDFromClassID() {}

func StringFromInterfaceID(interfaceID *GUID) (string, error) {
	return "", NotImplementedError
}

func StringFromGUID() {}

func StringFromClassID(classID *GUID) (string, error) {
	return "", NotImplementedError
}
