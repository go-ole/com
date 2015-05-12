// +build !windows

package com

// BindMoniker is a helper function for locating an object by its moniker and
// retrieving the pointer to the interface.
//
// moniker appears to be the address to the IMoniker object.
//
// option is ignored, because it is optional for now and must be '0'.
//
// interfaceID is the IID of the COM interface for result object.
//
// object that will be receive the object of IMoniker initialization.
//
// BindMoniker in Windows API.
func BindMoniker(moniker uintptr, option int32, interfaceID *GUID, object interface{}) error {
	return NotImplementedError
}

// ClassIDFromProgramID retrieves Class Identifier with the given Program
// Identifier.
//
// The Programmatic Identifier must be registered, because it will be looked up
// in the Windows Registry. The registry entry has the following keys: CLSID,
// Insertable, Protocol and Shell
// (https://msdn.microsoft.com/en-us/library/dd542719(v=vs.85).aspx).
//
// programID identifies the class id with less precision and is not guaranteed
// to be unique. These are usually found in the registry under
// HKEY_LOCAL_MACHINE\SOFTWARE\Classes, usually with the format of
// "Program.Component.Version" with version being optional.
//
// CLSIDFromProgID in Windows API.
func ClassIDFromProgramID(programID string) (*GUID, error) {
	return nil, NotImplementedError
}

// ClassIDFromProgramIDEx is similar to ClassIDFromProgramID(), but will attempt
// to install if COMClassStore policy is enabled.
//
// This appears to be used when the client machine is part of Active Directory
// and the policy allows for downloading missing COM components. Also, the COM
// component would have to be registered.
//
// CLSIDFromProgIDEx in Windows API.
func ClassIDFromProgramIDEx(programID string) (*GUID, error) {
	return nil, NotImplementedError
}

// ClassIDFromString retrieves Class ID from string representation.
//
// This is technically the string version of the GUID and will convert the
// string to object.
//
// CLSIDFromString in Windows API.
func ClassIDFromString(str string) (*GUID, error) {
	return nil, NotImplementedError
}

// CoAddRefServerProcess adds reference to global per-process counter.
//
// This is for COM Servers, which is not currently supported.
//
// Returns current reference count, with possible error.
//
// CoAddRefServerProcess in Windows API.
func CoAddRefServerProcess() (uint64, error) {
	return 0, NotImplementedError
}

// CoAllowSetForegroundWindow COM Server unfocuses client application to focus
// to another window.
//
// This is for COM Servers, which is not currently supported.
//
// unknown is the foreground window process to gain focus.
//
// CoAllowSetForegroundWindow in Windows API.
func CoAllowSetForegroundWindow(unknown iunknown.IUnknown) error {
	return NotImplementedError
}

//
//
// CoAllowUnmarshalerCLSID in Windows API.
func CoAllowUnmarshalerClassID() {}

//
//
// CoCancelCall in Windows API.
func CoCancelCall() {}

//
//
// CoCopyProxy in Windows API.
func CoCopyProxy() {}

//
//
// CoCreateFreeThreadedMarshaler in Windows API.
func CoCreateFreeThreadedMarshaler() {}

//
//
// CoCreateGUID in Windows API.
func CoCreateGUID() {}

//
//
// CoCreateInstance in Windows API.
func CoCreateInstance() {}

//
//
// CoCreateInstanceEx in Windows API.
func CoCreateInstanceEx() {}

//
//
// CoCreateInstanceFromApp in Windows API.
func CoCreateInstanceFromApp() {}

//
//
// CoDisableCallCancellation in Windows API.
func CoDisableCallCancellation() {}

//
//
// CoDisconnectContext in Windows API.
func CoDisconnectContext() {}

//
//
// CoDisconnectObject in Windows API.
func CoDisconnectObject() {}

//
//
// CoEnableCallCancellation in Windows API.
func CoEnableCallCancellation() {}

//
//
// CoEnterApplicationThreadLifetimeLoop in Windows API.
func CoEnterApplicationThreadLifetimeLoop() {}

//
//
// CoFileTimeNow in Windows API.
func CoFileTimeNow() {}

//
//
// CoFreeAllLibraries in Windows API.
func CoFreeAllLibraries() {}

//
//
// CoFreeLibrary in Windows API.
func CoFreeLibrary() {}

//
//
// CoFreeUnusedLibrariesEx in Windows API.
func CoFreeUnusedLibrariesEx() {}

//
//
// CoGetApartmentType in Windows API.
func CoGetApartmentType() {}

//
//
// CoGetApplicationThreadReference in Windows API.
func CoGetApplicationThreadReference() {}

//
//
// CoGetCallContext in Windows API.
func CoGetCallContext() {}

//
//
// CoGetCallerTID in Windows API.
func CoGetCallerTID() {}

//
//
// CoGetCancelObject in Windows API.
func CoGetCancelObject() {}

//
//
// CoGetClassObject in Windows API.
func CoGetClassObject() {}

//
//
// CoGetContextToken in Windows API.
func CoGetContextToken() {}

//
//
// CoGetCurrentLogicalThreadId in Windows API.
func CoGetCurrentLogicalThreadID() {}

//
//
// CoGetCurrentProcess in Windows API.
func CoGetCurrentProcess() {}

//
//
// CoGetInstanceFromFile in Windows API.
func CoGetInstanceFromFile() {}

//
//
// CoGetInstanceFromIStorage in Windows API.
func CoGetInstanceFromIStorage() {}

//
//
// CoGetInterceptor in Windows API.
func CoGetInterceptor() {}

//
//
// CoGetInterfaceAndReleaseStream in Windows API.
func CoGetInterfaceAndReleaseStream() {}

//
//
// CoGetMalloc in Windows API.
func CoGetMalloc() {}

//
//
// CoGetMarshalSizeMax in Windows API.
func CoGetMarshalSizeMax() {}

//
//
// CoGetObject in Windows API.
func CoGetObject() {}

//
//
// CoGetObjectContext in Windows API.
func CoGetObjectContext() {}

//
//
// CoGetPSClsid in Windows API.
func CoGetPSClassID() {}

//
//
// CoGetStandardMarshal in Windows API.
func CoGetStandardMarshal() {}

//
//
// CoGetStdMarshalEx in Windows API.
func CoGetStandardMarshalEx() {}

//
//
// CoGetSystemSecurityPermissions in Windows API.
func CoGetSystemSecurityPermissions() {}

//
//
// CoGetTreatAsClass in Windows API.
func CoGetTreatAsClass() {}

//
//
// CoHandlePriorityEventsFromMessagePump in Windows API.
func CoHandlePriorityEventsFromMessagePump() {}

//
//
// CoImpersonateClient in Windows API.
func CoImpersonateClient() {}

// CoInitialize initializes COM library on current thread.
//
// MSDN documentation suggests that this function should not be called. Call
// CoInitializeEx() instead. The reason has to do with threading and this
// function is only for single-threaded apartments.
//
// That said, most users of the library have gotten away with just this
// function. If you are experiencing threading issues, then use
// CoInitializeEx().
//
// CoInitialize in Windows API.
func CoInitialize() error {
	return NotImplementedError
}

// CoInitializeEx initializes COM library with concurrency model.
//
// CoInitializeEx in Windows API.
func CoInitializeEx(coinit ConcurrencyModel) error {
	return LibraryCallback().CoInitializeEx(coinit)
}

//
//
// CoInitializeSecurity in Windows API.
func CoInitializeSecurity() {}

//
//
// CoInstall in Windows API.
func CoInstall() {}

//
//
// CoInvalidateRemoteMachineBindings in Windows API.
func CoInvalidateRemoteMachineBindings() {}

//
//
// CoIsHandlerConnected in Windows API.
func CoIsHandlerConnected() {}

//
//
// CoIsOle1Class in Windows API.
func CoIsOle1Class() {}

//
//
// CoLoadLibrary in Windows API.
func CoLoadLibrary() {}

//
//
// CoLockObjectExternal in Windows API.
func CoLockObjectExternal() {}

//
//
// CoMarshalHresult in Windows API.
func CoMarshalHResult() {}

//
//
// CoMarshalInterface in Windows API.
func CoMarshalInterface() {}

//
//
// CoMarshalInterThreadInterfaceInStream in Windows API.
func CoMarshalInterThreadInterfaceInStream() {}

//
//
// CoQueryAuthenticationServices in Windows API.
func CoQueryAuthenticationServices() {}

//
//
// CoQueryClientBlanket in Windows API.
func CoQueryClientBlanket() {}

//
//
// CoQueryProxyBlanket in Windows API.
func CoQueryProxyBlanket() {}

//
//
// CoRegisterClassObject in Windows API.
func CoRegisterClassObject() {}

//
//
// CoRegisterInitializeSpy in Windows API.
func CoRegisterInitializeSpy() {}

//
//
// CoRegisterMallocSpy in Windows API.
func CoRegisterMallocSpy() {}

//
//
// CoRegisterMessageFilter in Windows API.
func CoRegisterMessageFilter() {}

//
//
// CoRegisterPSClsid in Windows API.
func CoRegisterPSClassID() {}

//
//
// CoRegisterSurrogate in Windows API.
func CoRegisterSurrogate() {}

//
//
// CoReleaseMarshalData in Windows API.
func CoReleaseMarshalData() {}

//
//
// CoReleaseServerProcess in Windows API.
func CoReleaseServerProcess() {}

//
//
// CoResumeClassObjects in Windows API.
func CoResumeClassObjects() {}

//
//
// CoRevertToSelf in Windows API.
func CoRevertToSelf() {}

//
//
// CoRevokeClassObject in Windows API.
func CoRevokeClassObject() {}

//
//
// CoRevokeInitializeSpy in Windows API.
func CoRevokeInitializeSpy() {}

//
//
// CoRevokeMallocSpy in Windows API.
func CoRevokeMallocSpy() {}

//
//
// CoSetCancelObject in Windows API.
func CoSetCancelObject() {}

//
//
// CoSetMessageDispatcher in Windows API.
func CoSetMessageDispatcher() {}

//
//
// CoSetProxyBlanket in Windows API.
func CoSetProxyBlanket() {}

//
//
// CoSuspendClassObjects in Windows API.
func CoSuspendClassObjects() {}

//
//
// CoSwitchCallContext in Windows API.
func CoSwitchCallContext() {}

// CoTaskMemFree frees memory pointer.
func CoTaskMemFree(memoryPointer uintptr) {
	LibraryCallback().CoTaskMemFree(memoryPointer)
}

//
//
// CoTaskMemRealloc in Windows API.
func CoTaskMemRealloc() {}

//
//
// CoTestCancel in Windows API.
func CoTestCancel() {}

//
//
// CoTreatAsClass in Windows API.
func CoTreatAsClass() {}

// CoUninitialize uninitializes COM Library.
func CoUninitialize() {
	LibraryCallback().CoUninitialize()
}

//
//
// CoUnmarshalHresult in Windows API.
func CoUnmarshalHResult() {}

//
//
// CoUnmarshalInterface in Windows API.
func CoUnmarshalInterface() {}

//
//
// CoWaitForMultipleHandles in Windows API.
func CoWaitForMultipleHandles() {}

//
//
// CreateAntiMoniker in Windows API.
func CreateAntiMoniker() {}

//
//
// CreateAsyncBindCtx in Windows API.
func CreateAsyncBindContext() {}

//
//
// CreateBindCtx in Windows API.
func CreateBindContext() {}

//
//
// CreateClassMoniker in Windows API.
func CreateClassMoniker() {}

//
//
// CreateFileMoniker in Windows API.
func CreateFileMoniker() {}

//
//
// CreateGenericComposite in Windows API.
func CreateGenericComposite() {}

//
//
// CreateItemMoniker in Windows API.
func CreateItemMoniker() {}

//
//
// CreateObjrefMoniker in Windows API.
func CreateObjectRefMoniker() {}

//
//
// CreatePointerMoniker in Windows API.
func CreatePointerMoniker() {}

//
//
// DllCanUnloadNow in Windows API.
func DLLCanUnloadNow() {}

//
//
// DllDebugObjectRPCHook in Windows API.
func DLLDebugObjectRPCHook() {}

//
//
// DllGetClassObject in Windows API.
func DLLGetClassObject() {}

//
//
// DllRegisterServer in Windows API.
func DLLRegisterServer() {}

//
//
// DllUnregisterServer in Windows API.
func DLLUnregisterServer() {}

//
//
// GetClassFile in Windows API.
func GetClassFile() {}

//
//
// GetRunningObjectTable in Windows API.
func GetRunningObjectTable() {}

//
//
// IIDFromString in Windows API.
func InterfaceIDFromString() {}

//
//
// Initialize in Windows API.
func Initialize() {}

//
//
// IsAccelerator in Windows API.
func IsAccelerator() {}

//
//
// IsEqualCLSID in Windows API.
func IsEqualClassID() {}

//
//
// IsEqualGUID in Windows API.
func IsEqualGUID() {}

//
//
// IsEqualIID in Windows API.
func IsEqualInterfaceID() {}

//
//
// MkParseDisplayName() in Windows API.
func MkParseDisplayName() {}

//
//
// MonikerCommonPrefixWith() in Windows API.
func MonikerCommonPrefixWith() {}

//
//
// MonikerRelativePathTo() in Windows API.
func MonikerRelativePathTo() {}

//
//
// OleDoAutoConvert() in Windows API.
func OleDoAutoConvert() {}

//
//
// OleGetAutoConvert() in Windows API.
func OleGetAutoConvert() {}

//
//
// OleGetIconOfClass() in Windows API.
func OleGetIconOfClass() {}

//
//
// OleGetIconOfFile() in Windows API.
func OleGetIconOfFile() {}

//
//
// OleIconToCursor() in Windows API.
func OleIconToCursor() {}

//
//
// OleRegGetMiscStatus() in Windows API.
func OleRegGetMiscStatus() {}

//
//
// OleRegGetUserType() in Windows API.
func OleRegGetUserType() {}

//
//
// OleSetAutoConvert() in Windows API.
func OleSetAutoConvert() {}

//
//
// ProgIDFromCLSID() in Windows API.
func ProgramIDFromClassID() {}

// StringFromInterfaceID returns GUID formated string from GUID object.
//
// StringFromIID() in Windows API.
func StringFromInterfaceID(interfaceID *GUID) (string, error) {
	return LibraryCallback().StringFromInterfaceID(interfaceID)
}

//
//
// StringFromGUID2() in Windows API.
func StringFromGUID() {}

// StringFromClassID returns GUID formated string from GUID object.
//
// StringFromCLSID() in Windows API.
func StringFromClassID(classID *GUID) (string, error) {
	return LibraryCallback().StringFromClassID(classID)
}
