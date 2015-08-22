// +build windows,!cgo

package com

import (
	"reflect"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/go-ole/go-ole/types/iunknown"
)

var (
	modole32, _ = syscall.LoadDLL("ole32.dll")
)

var (
	procBindMoniker, _                = modole32.FindProc("BindMoniker")
	procCLSIDFromProgID, _            = modole32.FindProc("CLSIDFromProgID")
	procCLSIDFromProgIDEx, _          = modole32.FindProc("CLSIDFromProgIDEx")
	procCLSIDFromString, _            = modole32.FindProc("CLSIDFromString")
	procCoAddRefServerProcess, _      = modole32.FindProc("CoAddRefServerProcess")
	procCoAllowSetForegroundWindow, _ = modole32.FindProc("CoAllowSetForegroundWindow")
	procCoInitialize, _               = modole32.FindProc("CoInitialize")
	procCoInitializeEx, _             = modole32.FindProc("CoInitializeEx")
	procCoUninitialize, _             = modole32.FindProc("CoUninitialize")
	procCoCreateInstance, _           = modole32.FindProc("CoCreateInstance")
	procCoTaskMemFree, _              = modole32.FindProc("CoTaskMemFree")
	procStringFromCLSID, _            = modole32.FindProc("StringFromCLSID")
	procStringFromIID, _              = modole32.FindProc("StringFromIID")
	procIIDFromString, _              = modole32.FindProc("IIDFromString")
)

func BindMoniker(moniker uintptr, option int32, interfaceID *GUID, object interface{}) error {
	return HResultToError(procBindMoniker.Call(
		moniker,
		uintptr(0),
		uintptr(unsafe.Pointer(interfaceID)),
		reflect.ValueOf(object).UnsafeAddr()))
}

func ClassIDFromProgramID(programID string) (*GUID, error) {
	return classIDFromProgramIDHelper(procCLSIDFromProgID, programID)
}

func ClassIDFromProgramIDEx(programID string) (*GUID, error) {
	return classIDFromProgramIDHelper(procCLSIDFromProgIDEx, programID)
}

func classIDFromProgramIDHelper(p *syscall.Proc, programID string) (classID *GUID, err error) {
	ptrProgramID, err := windows.UTF16PtrFromString(programID)
	if err != nil {
		return
	}

	err = HResultToError(p.Call(
		uintptr(unsafe.Pointer(ptrProgramID)),
		uintptr(unsafe.Pointer(&classID))))
	return
}

func ClassIDFromString(str string) (classID *GUID, err error) {
	ptrStr, err := windows.UTF16PtrFromString(str)
	if err != nil {
		return
	}

	err = HResultToError(procCLSIDFromString.Call(
		uintptr(unsafe.Pointer(ptrStr)),
		uintptr(unsafe.Pointer(&classID))))
	return
}

func CoAddRefServerProcess() (uint64, error) {
	referenceCount, _, err := procCoAddRefServerProcess.Call()
	return uint64(referenceCount), err
}

func CoAllowSetForegroundWindow(unknown iunknown.IUnknown) error {
	return HResultToError(procCoAllowSetForegroundWindow.Call(
		uintptr(unsafe.Pointer(unknown)), uintptr(0)))
}

func CoAllowUnmarshalerClassID() {}

func CoCancelCall() {}

func CoCopyProxy() {}

func CoCreateFreeThreadedMarshaler() {}

func CoCreateGUID() {}

func CoCreateInstance(classID, interfaceID *GUID, object interface{}) error {
	if interfaceID == nil {
		interfaceID = iunknown.InterfaceID
	}

	err = HResultToError(procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(classID)),
		0,
		AllServerClassContext,
		uintptr(unsafe.Pointer(interfaceID)),
		reflect.ValueOf(object).UnsafeAddr()))

	return
}

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
	return HResultToError(procCoInitialize.Call(uintptr(0)))
}

func CoInitializeEx(coinit ConcurrencyModel) error {
	// http://msdn.microsoft.com/en-us/library/windows/desktop/ms695279(v=vs.85).aspx
	// Suggests that the first parameter is not only optional but should always be NULL.
	return HResultToError(procCoInitializeEx.Call(uintptr(0), uintptr(uint32(coinit))))
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

func CoTaskMemFree(memptr uintptr) {
	procCoTaskMemFree.Call(memptr)
}

func CoTaskMemRealloc() {}

func CoTestCancel() {}

func CoTreatAsClass() {}

func CoUninitialize() {
	procCoUninitialize.Call()
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

func StringFromInterfaceID(interfaceID *GUID) (str string, err error) {
	var p *uint16
	err = HResultToError(procStringFromIID.Call(
		uintptr(unsafe.Pointer(interfaceID)),
		uintptr(unsafe.Pointer(&p))))
	str = LpOleStrToString(p)
	return
}

func StringFromClassID(classID *GUID) (str string, err error) {
	var p *uint16
	err = HResultToError(procStringFromCLSID.Call(
		uintptr(unsafe.Pointer(classID)),
		uintptr(unsafe.Pointer(&p))))
	str = LpOleStrToString(p)
	return
}

func InterfaceIDFromString(programID string) (classID *GUID, err error) {
	lpsz, err := windows.UTF16PtrFromString(programID)
	if err != nil {
		return
	}
	err = HResultToError(procIIDFromString.Call(uintptr(unsafe.Pointer(lpsz)), uintptr(unsafe.Pointer(&classID))))
	return
}

func GetActiveObject(classID, interfaceID *GUID) (unk *iunknown.Unknown, err error) {
	if interfaceID == nil {
		interfaceID = iunknown.InterfaceID
	}

	err = HResultToError(procGetActiveObject.Call(
		uintptr(unsafe.Pointer(classID)),
		uintptr(unsafe.Pointer(interfaceID)),
		uintptr(unsafe.Pointer(&unk))))

	return
}
