// +build !windows

package com

func BindMoniker(moniker uintptr, option int32, interfaceID *GUID, object interface{}) error {
	return NotImplementedError
}

func ClassIDFromProgramID(programID string) (*GUID, error) {
	return nil, NotImplementedError
}

func ClassIDFromProgramIDEx(programID string) (*GUID, error) {
	return nil, NotImplementedError
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
	return NotImplementedError
}

func CoInitializeEx(coinit ConcurrencyModel) error {
	return NotImplementedError
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
	// noop
}

func CoTaskMemRealloc() {}

func CoTestCancel() {}

func CoTreatAsClass() {}

func CoUninitialize() {
	// noop
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
