// +build !windows

package com // "github.com/go-ole/com"

import api "github.com/go-ole/com/internal/com"

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
	return api.BindMoniker(moniker, option, interfaceID, object)
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
	return api.ClassIDFromProgramID(programID)
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
	return api.ClassIDFromProgramIDEx(programID)
}

// ClassIDFromString retrieves Class ID from string representation.
//
// This is technically the string version of the GUID and will convert the
// string to object.
//
// CLSIDFromString in Windows API.
func ClassIDFromString(str string) (*GUID, error) {
	return api.ClassIDFromString(str)
}

// CoAddRefServerProcess adds reference to global per-process counter.
//
// This is for COM Servers, which is not currently supported.
//
// Returns current reference count, with possible error.
//
// CoAddRefServerProcess in Windows API.
func CoAddRefServerProcess() (uint64, error) {
	return api.CoAddRefServerProcess()
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
	return api.CoAllowSetForegroundWindow(unknown)
}

//
//
// CoAllowUnmarshalerCLSID in Windows API.
func CoAllowUnmarshalerClassID() {
	return api.CoAllowUnmarshalerClassID()
}

//
//
// CoCancelCall in Windows API.
func CoCancelCall() {
	return api.CoCancelCall()
}

//
//
// CoCopyProxy in Windows API.
func CoCopyProxy() {
	return api.CoCopyProxy()
}

//
//
// CoCreateFreeThreadedMarshaler in Windows API.
func CoCreateFreeThreadedMarshaler() {
	return api.CoCreateFreeThreadedMarshaler()
}

//
//
// CoCreateGUID in Windows API.
func CoCreateGUID() {
	return api.CoCreateGUID()
}

//
//
// CoCreateInstance in Windows API.
func CoCreateInstance() {
	return api.CoCreateInstance()
}

//
//
// CoCreateInstanceEx in Windows API.
func CoCreateInstanceEx() {
	return api.CoCreateInstanceEx()
}

//
//
// CoCreateInstanceFromApp in Windows API.
func CoCreateInstanceFromApp() {
	return api.CoCreateInstanceFromApp()
}

//
//
// CoDisableCallCancellation in Windows API.
func CoDisableCallCancellation() {
	return api.CoDisableCallCancellation()
}

//
//
// CoDisconnectContext in Windows API.
func CoDisconnectContext() {
	return api.CoDisconnectContext()
}

//
//
// CoDisconnectObject in Windows API.
func CoDisconnectObject() {
	return api.CoDisconnectObject()
}

//
//
// CoEnableCallCancellation in Windows API.
func CoEnableCallCancellation() {
	return api.CoEnableCallCancellation()
}

//
//
// CoEnterApplicationThreadLifetimeLoop in Windows API.
func CoEnterApplicationThreadLifetimeLoop() {
	return api.CoEnterApplicationThreadLifetimeLoop()
}

//
//
// CoFileTimeNow in Windows API.
func CoFileTimeNow() {
	return api.CoFileTimeNow()
}

//
//
// CoFreeAllLibraries in Windows API.
func CoFreeAllLibraries() {
	return api.CoFreeAllLibraries()
}

//
//
// CoFreeLibrary in Windows API.
func CoFreeLibrary() {
	return api.CoFreeLibrary()
}

//
//
// CoFreeUnusedLibrariesEx in Windows API.
func CoFreeUnusedLibrariesEx() {
	return api.CoFreeUnusedLibrariesEx()
}

//
//
// CoGetApartmentType in Windows API.
func CoGetApartmentType() {
	return api.CoGetApartmentType()
}

//
//
// CoGetApplicationThreadReference in Windows API.
func CoGetApplicationThreadReference() {
	return api.CoGetApplicationThreadReference()
}

//
//
// CoGetCallContext in Windows API.
func CoGetCallContext() {
	return api.CoGetCallContext()
}

//
//
// CoGetCallerTID in Windows API.
func CoGetCallerTID() {
	return api.CoGetCallerTID()
}

//
//
// CoGetCancelObject in Windows API.
func CoGetCancelObject() {
	return api.CoGetCancelObject()
}

//
//
// CoGetClassObject in Windows API.
func CoGetClassObject() {
	return api.CoGetClassObject()
}

//
//
// CoGetContextToken in Windows API.
func CoGetContextToken() {
	return api.CoGetContextToken()
}

//
//
// CoGetCurrentLogicalThreadId in Windows API.
func CoGetCurrentLogicalThreadID() {
	return api.CoGetCurrentLogicalThreadID()
}

//
//
// CoGetCurrentProcess in Windows API.
func CoGetCurrentProcess() {
	return api.CoGetCurrentProcess()
}

//
//
// CoGetInstanceFromFile in Windows API.
func CoGetInstanceFromFile() {
	return api.CoGetInstanceFromFile()
}

//
//
// CoGetInstanceFromIStorage in Windows API.
func CoGetInstanceFromIStorage() {
	return api.CoGetInstanceFromIStorage()
}

//
//
// CoGetInterceptor in Windows API.
func CoGetInterceptor() {
	return api.CoGetInterceptor()
}

//
//
// CoGetInterfaceAndReleaseStream in Windows API.
func CoGetInterfaceAndReleaseStream() {
	return api.CoGetInterfaceAndReleaseStream()
}

//
//
// CoGetMalloc in Windows API.
func CoGetMalloc() {
	return api.CoGetMalloc()
}

//
//
// CoGetMarshalSizeMax in Windows API.
func CoGetMarshalSizeMax() {
	return api.CoGetMarshalSizeMax()
}

//
//
// CoGetObject in Windows API.
func CoGetObject() {
	return api.CoGetObject()
}

//
//
// CoGetObjectContext in Windows API.
func CoGetObjectContext() {
	return api.CoGetObjectContext()
}

//
//
// CoGetPSClsid in Windows API.
func CoGetPSClassID() {
	return api.CoGetPSClassID()
}

//
//
// CoGetStandardMarshal in Windows API.
func CoGetStandardMarshal() {
	return api.CoGetStandardMarshal()
}

//
//
// CoGetStdMarshalEx in Windows API.
func CoGetStandardMarshalEx() {
	return api.CoGetStandardMarshalEx()
}

//
//
// CoGetSystemSecurityPermissions in Windows API.
func CoGetSystemSecurityPermissions() {
	return api.CoGetSystemSecurityPermissions()
}

//
//
// CoGetTreatAsClass in Windows API.
func CoGetTreatAsClass() {
	return api.CoGetTreatAsClass()
}

//
//
// CoHandlePriorityEventsFromMessagePump in Windows API.
func CoHandlePriorityEventsFromMessagePump() {
	return api.CoHandlePriorityEventsFromMessagePump()
}

//
//
// CoImpersonateClient in Windows API.
func CoImpersonateClient() {
	return api.CoImpersonateClient()
}

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
	return api.CoInitialize()
}

// CoInitializeEx initializes COM library with concurrency model.
//
// CoInitializeEx in Windows API.
func CoInitializeEx(coinit ConcurrencyModel) error {
	return api.CoInitializeEx(coinit)
}

//
//
// CoInitializeSecurity in Windows API.
func CoInitializeSecurity() {
	return api.CoInitializeSecurity()
}

//
//
// CoInstall in Windows API.
func CoInstall() {
	return api.CoInstall()
}

//
//
// CoInvalidateRemoteMachineBindings in Windows API.
func CoInvalidateRemoteMachineBindings() {
	return api.CoInvalidateRemoteMachineBindings()
}

//
//
// CoIsHandlerConnected in Windows API.
func CoIsHandlerConnected() {
	return api.CoIsHandlerConnected()
}

//
//
// CoIsOle1Class in Windows API.
func CoIsOle1Class() {
	return api.CoIsOle1Class()
}

//
//
// CoLoadLibrary in Windows API.
func CoLoadLibrary() {
	return api.CoLoadLibrary()
}

//
//
// CoLockObjectExternal in Windows API.
func CoLockObjectExternal() {
	return api.CoLockObjectExternal()
}

//
//
// CoMarshalHresult in Windows API.
func CoMarshalHResult() {
	return api.CoMarshalHResult()
}

//
//
// CoMarshalInterface in Windows API.
func CoMarshalInterface() {
	return api.CoMarshalInterface()
}

//
//
// CoMarshalInterThreadInterfaceInStream in Windows API.
func CoMarshalInterThreadInterfaceInStream() {
	return api.CoMarshalInterThreadInterfaceInStream()
}

//
//
// CoQueryAuthenticationServices in Windows API.
func CoQueryAuthenticationServices() {
	return api.CoQueryAuthenticationServices()
}

//
//
// CoQueryClientBlanket in Windows API.
func CoQueryClientBlanket() {
	return api.CoQueryClientBlanket()
}

//
//
// CoQueryProxyBlanket in Windows API.
func CoQueryProxyBlanket() {
	return api.CoQueryProxyBlanket()
}

//
//
// CoRegisterClassObject in Windows API.
func CoRegisterClassObject() {
	return api.CoRegisterClassObject()
}

//
//
// CoRegisterInitializeSpy in Windows API.
func CoRegisterInitializeSpy() {
	return api.CoRegisterInitializeSpy()
}

//
//
// CoRegisterMallocSpy in Windows API.
func CoRegisterMallocSpy() {
	return api.CoRegisterMallocSpy()
}

//
//
// CoRegisterMessageFilter in Windows API.
func CoRegisterMessageFilter() {
	return api.CoRegisterMessageFilter()
}

//
//
// CoRegisterPSClsid in Windows API.
func CoRegisterPSClassID() {
	return api.CoRegisterPSClassID()
}

//
//
// CoRegisterSurrogate in Windows API.
func CoRegisterSurrogate() {
	return api.CoRegisterSurrogate()
}

//
//
// CoReleaseMarshalData in Windows API.
func CoReleaseMarshalData() {
	return api.CoReleaseMarshalData()
}

//
//
// CoReleaseServerProcess in Windows API.
func CoReleaseServerProcess() {
	return api.CoReleaseServerProcess()
}

//
//
// CoResumeClassObjects in Windows API.
func CoResumeClassObjects() {
	return api.CoResumeClassObjects()
}

//
//
// CoRevertToSelf in Windows API.
func CoRevertToSelf() {
	return api.CoRevertToSelf()
}

//
//
// CoRevokeClassObject in Windows API.
func CoRevokeClassObject() {
	return api.CoRevokeClassObject()
}

//
//
// CoRevokeInitializeSpy in Windows API.
func CoRevokeInitializeSpy() {
	return api.CoRevokeInitializeSpy()
}

//
//
// CoRevokeMallocSpy in Windows API.
func CoRevokeMallocSpy() {
	return api.CoRevokeMallocSpy()
}

//
//
// CoSetCancelObject in Windows API.
func CoSetCancelObject() {
	return api.CoSetCancelObject()
}

//
//
// CoSetMessageDispatcher in Windows API.
func CoSetMessageDispatcher() {
	return api.CoSetMessageDispatcher()
}

//
//
// CoSetProxyBlanket in Windows API.
func CoSetProxyBlanket() {
	return api.CoSetProxyBlanket()
}

//
//
// CoSuspendClassObjects in Windows API.
func CoSuspendClassObjects() {
	return api.CoSuspendClassObjects()
}

//
//
// CoSwitchCallContext in Windows API.
func CoSwitchCallContext() {
	return api.CoSwitchCallContext()
}

// CoTaskMemFree frees memory pointer.
func CoTaskMemFree(memoryPointer uintptr) {
	return api.CoTaskMemFree(memoryPointer)
}

//
//
// CoTaskMemRealloc in Windows API.
func CoTaskMemRealloc() {
	return api.CoTaskMemRealloc()
}

//
//
// CoTestCancel in Windows API.
func CoTestCancel() {
	return api.CoTestCancel()
}

//
//
// CoTreatAsClass in Windows API.
func CoTreatAsClass() {
	return api.CoTreatAsClass()
}

// CoUninitialize uninitializes COM Library.
func CoUninitialize() {
	return api.CoUninitialize()
}

//
//
// CoUnmarshalHresult in Windows API.
func CoUnmarshalHResult() {
	return api.CoUnmarshalHResult()
}

//
//
// CoUnmarshalInterface in Windows API.
func CoUnmarshalInterface() {
	return api.CoUnmarshalInterface()
}

//
//
// CoWaitForMultipleHandles in Windows API.
func CoWaitForMultipleHandles() {
	return api.CoWaitForMultipleHandles()
}

//
//
// CreateAntiMoniker in Windows API.
func CreateAntiMoniker() {
	return api.CreateAntiMoniker()
}

//
//
// CreateAsyncBindCtx in Windows API.
func CreateAsyncBindContext() {
	return api.CreateAsyncBindContext()
}

//
//
// CreateBindCtx in Windows API.
func CreateBindContext() {
	return api.CreateBindContext()
}

//
//
// CreateClassMoniker in Windows API.
func CreateClassMoniker() {
	return api.CreateClassMoniker()
}

//
//
// CreateFileMoniker in Windows API.
func CreateFileMoniker() {
	return api.CreateFileMoniker()
}

//
//
// CreateGenericComposite in Windows API.
func CreateGenericComposite() {
	return api.CreateGenericComposite()
}

//
//
// CreateItemMoniker in Windows API.
func CreateItemMoniker() {
	return api.CreateItemMoniker()
}

//
//
// CreateObjrefMoniker in Windows API.
func CreateObjectRefMoniker() {
	return api.CreateObjectRefMoniker()
}

//
//
// CreatePointerMoniker in Windows API.
func CreatePointerMoniker() {
	return api.CreatePointerMoniker()
}

//
//
// DllCanUnloadNow in Windows API.
func DLLCanUnloadNow() {
	return api.DLLCanUnloadNow()
}

//
//
// DllDebugObjectRPCHook in Windows API.
func DLLDebugObjectRPCHook() {
	return api.DLLDebugObjectRPCHook()
}

//
//
// DllGetClassObject in Windows API.
func DLLGetClassObject() {
	return api.DLLGetClassObject()
}

//
//
// DllRegisterServer in Windows API.
func DLLRegisterServer() {
	return api.DLLRegisterServer()
}

//
//
// DllUnregisterServer in Windows API.
func DLLUnregisterServer() {
	return api.DLLUnregisterServer()
}

//
//
// GetClassFile in Windows API.
func GetClassFile() {
	return api.GetClassFile()
}

//
//
// GetRunningObjectTable in Windows API.
func GetRunningObjectTable() {
	return api.GetRunningObjectTable()
}

//
//
// IIDFromString in Windows API.
func InterfaceIDFromString() {
	return api.InterfaceIDFromString()
}

//
//
// Initialize in Windows API.
func Initialize() {
	return api.Initialize()
}

//
//
// IsAccelerator in Windows API.
func IsAccelerator() {
	return api.IsAccelerator()
}

//
//
// IsEqualCLSID in Windows API.
func IsEqualClassID() {
	return api.IsEqualClassID()
}

//
//
// IsEqualGUID in Windows API.
func IsEqualGUID() {
	return api.IsEqualGUID()
}

//
//
// IsEqualIID in Windows API.
func IsEqualInterfaceID() {
	return api.IsEqualInterfaceID()
}

//
//
// MkParseDisplayName() in Windows API.
func MkParseDisplayName() {
	return api.MkParseDisplayName()
}

//
//
// MonikerCommonPrefixWith() in Windows API.
func MonikerCommonPrefixWith() {
	return api.MonikerCommonPrefixWith()
}

//
//
// MonikerRelativePathTo() in Windows API.
func MonikerRelativePathTo() {
	return api.MonikerRelativePathTo()
}

//
//
// OleDoAutoConvert() in Windows API.
func OleDoAutoConvert() {
	return api.OleDoAutoConvert()
}

//
//
// OleGetAutoConvert() in Windows API.
func OleGetAutoConvert() {
	return api.OleGetAutoConvert()
}

//
//
// OleGetIconOfClass() in Windows API.
func OleGetIconOfClass() {
	return api.OleGetIconOfClass()
}

//
//
// OleGetIconOfFile() in Windows API.
func OleGetIconOfFile() {
	return api.OleGetIconOfFile()
}

//
//
// OleIconToCursor() in Windows API.
func OleIconToCursor() {
	return api.OleIconToCursor()
}

//
//
// OleRegGetMiscStatus() in Windows API.
func OleRegGetMiscStatus() {
	return api.OleRegGetMiscStatus()
}

//
//
// OleRegGetUserType() in Windows API.
func OleRegGetUserType() {
	return api.OleRegGetUserType()
}

//
//
// OleSetAutoConvert() in Windows API.
func OleSetAutoConvert() {
	return api.OleSetAutoConvert()
}

//
//
// ProgIDFromCLSID() in Windows API.
func ProgramIDFromClassID() {
	return api.ProgramIDFromClassID()
}

// StringFromInterfaceID returns GUID formated string from GUID object.
//
// StringFromIID() in Windows API.
func StringFromInterfaceID(interfaceID *GUID) (string, error) {
	return api.StringFromInterfaceID(interfaceID)
}

//
//
// StringFromGUID2() in Windows API.
func StringFromGUID() {
	return api.StringFromGUID()
}

// StringFromClassID returns GUID formated string from GUID object.
//
// StringFromCLSID() in Windows API.
func StringFromClassID(classID *GUID) (string, error) {
	return api.StringFromClassID(classID)
}
