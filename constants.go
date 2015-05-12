package com

//go:generate stringer -type=VariantType
//go:generate stringer -type=DisplayID
//go:generate stringer -type=DispatchContext
//go:generate stringer -type=ClassContext

// VariantType is the data type of the Variant.
type VariantType uint16

// ConcurrencyModel describes how currency is handled with COM calls.
type ConcurrencyModel uint16

// DispatchContext are flags describing the context of the invoke call.
type DispatchContext int16

// ClassContext indicate what execution context class objects are to be made
// available.
//
// The contexts are used in activation calls and CoRegisterClassObject() for how
// to construct instances.
type ClassContext uint16

// DisplayID identifies methods, properties and arguments.
type DisplayID int32

// AccessControl describes whether access list is allowed or denied.
type AccessControl uint32

// CallingConvention defines the calling convention of the member function.
type CallingConvention uint8

// SafeArrayMask defines safe array.
type SafeArrayMask uint16

// KindType specifies type.
type KindType uint16

// RPCAuthenticationLevel is remote procedure call authentication levels.
type RPCAuthenticationLevel uint16

// RPCAuthenticationService is authentication service protocol.
type RPCAuthenticationService uint16

// HResultResponseCode are success or error codes.
type HResultResponseCode uint32

// Reference: https://msdn.microsoft.com/en-us/library/windows/desktop/ms682233(v=vs.85).aspx
const (
	// AccessAllowedAccessControl indicates an access-allowed entry.
	//
	// ACTRL_ACCESS_ALLOWED
	AccessAllowedAccessControl AccessControl = 0x00000000

	// AccessDeniedAccessControl indicates an access-denied entry.
	//
	// ACTRL_ACCESS_DENIED
	AccessDeniedAccessControl AccessControl = 0x10000000
)

// Reference: https://msdn.microsoft.com/en-us/library/windows/desktop/ms678435(v=vs.85).aspx
const (
	// DefaultAuthenticationLevel is default authentication level.
	//
	// RPC_C_AUTHN_LEVEL_DEFAULT
	DefaultAuthenticationLevel RPCAuthenticationLevel = iota

	// NoAuthenticationLevel performs no authentication.
	//
	// RPC_C_AUTHN_LEVEL_NONE
	NoAuthenticationLevel

	// ConnectAuthenticationLevel only authenticates when connecting to a
	// server.
	//
	// RPC_C_AUTHN_LEVEL_CONNECT
	ConnectAuthenticationLevel

	// CallAuthenticationLevel only authenticates at beginning RPC at server.
	//
	// RPC_C_AUTHN_LEVEL_CALL
	CallAuthenticationLevel

	// PKTAuthenticationLevel only authenticates when all data has been
	// received.
	//
	// RPC_C_AUTHN_LEVEL_PKT
	PKTAuthenticationLevel

	// PKTIntegrityAuthenticationLevel authenticates and verifies data has not
	// been modified.
	//
	// RPC_C_AUTHN_LEVEL_PKT_INTEGRITY
	PKTIntegrityAuthenticationLevel

	// PKTPrivacyAuthenticationLevel prevents disclose of data by encryption or
	// using secure channel.
	//
	// RPC_C_AUTHN_LEVEL_PKT_PRIVACY
	PKTPrivacyAuthenticationLevel
)

const (
	// NoAuthenticationService has no authentication.
	//
	// RPC_C_AUTHN_NONE
	NoAuthenticationService RPCAuthenticationService = 0

	// DCEPrivateAuthenticationService uses DCE private key authentication.
	//
	// RPC_C_AUTHN_DCE_PRIVATE
	DCEPrivateAuthenticationService RPCAuthenticationService = 1

	// DCEPublicAuthenticationService uses DCE public key authentication.
	//
	// Reserved for future use.
	//
	// RPC_C_AUTHN_DCE_PUBLIC
	DCEPublicAuthenticationService RPCAuthenticationService = 2

	// DECPublicAuthenticationService uses DEC public key authentication.
	//
	// Reserved for future use.
	//
	// RPC_C_AUTHN_DEC_PUBLIC
	DECPublicAuthenticationService RPCAuthenticationService = 4

	// GSSNegotiateAuthenticationService uses the Microsoft Negotiate SSP.
	//
	// RPC_C_AUTHN_GSS_NEGOTIATE
	GSSNegotiateAuthenticationService RPCAuthenticationService = 9

	// WINNTAuthenticationService uses the NTLM SSP.
	//
	// RPC_C_AUTHN_WINNT
	WINNTAuthenticationService RPCAuthenticationService = 10

	// GSSSChannelAuthenticationService uses the Schannel SSP.
	//
	// RPC_C_AUTHN_GSS_SCHANNEL
	GSSSChannelAuthenticationService RPCAuthenticationService = 14

	// GSSKerberosAuthenticationService uses the Microsoft Kerberos SSP.
	//
	// RPC_C_AUTHN_GSS_KERBEROS
	GSSKerberosAuthenticationService RPCAuthenticationService = 16

	// DPAAuthenticationService uses Distributed Password Authentication.
	//
	// RPC_C_AUTHN_DPA
	DPAAuthenticationService RPCAuthenticationService = 17

	// MSNAuthenticationService uses MSN authentication protocol.
	//
	// RPC_C_AUTHN_MSN
	MSNAuthenticationService RPCAuthenticationService = 18

	// KernelAuthenticationService is reserved and should not be used.
	//
	// RPC_C_AUTHN_KERNEL
	KernelAuthenticationService RPCAuthenticationService = 20

	// DigestAuthenticationService uses digest authentication.
	//
	// Windows XP or later.
	//
	// RPC_C_AUTHN_DIGEST
	DigestAuthenticationService RPCAuthenticationService = 21

	// NegoExtenderAuthenticationService is reserved and must not be used.
	//
	// Windows 7 or later.
	//
	// RPC_C_AUTHN_NEGO_EXTENDER
	NegoExtenderAuthenticationService RPCAuthenticationService = 30

	// PKU2UAuthenticationService is reserved and must not be used.
	//
	// RPC_C_AUTHN_PKU2U
	PKU2UAuthenticationService RPCAuthenticationService = 31

	// MessageQueueAuthenticationService provides an SSPI-compatible wrapper
	// for the Microsoft Message Queue (MSMQ) transport-level protocol.
	//
	// RPC_C_AUTHN_MQ
	MessageQueueAuthenticationService RPCAuthenticationService = 100

	// DefaultAuthenticationService is default authentication service.
	//
	// RPC_C_AUTHN_DEFAULT
	DefaultAuthenticationService RPCAuthenticationService = 0xFFFFFFFF
)

// Reference: https://msdn.microsoft.com/en-us/library/windows/desktop/ms693716(v=vs.85).aspx
const (
	// InProcessServerClassContext managed by DLL in same process as caller.
	//
	// CLSCTX_INPROC_SERVER in COM API.
	InProcessServerClassContext ClassContext = 0x1

	// InProcessHandlerClassContext managed by a DLL in the client process when
	// instances are accessed remotely.
	//
	// CLSCTX_INPROC_HANDLER in COM API.
	InProcessHandlerClassContext ClassContext = 0x2

	// LocalServerClassContext managed by a local executable loaded in a
	// separate process.
	//
	// CLSCTX_LOCAL_SERVER in COM API.
	LocalServerClassContext ClassContext = 0x4

	// InProcessServer16ClassContext is obsolete.
	//
	// CLSCTX_INPROC_SERVER16 in COM API.
	InProcessServer16ClassContext ClassContext = 0x8

	// RemoteServerClassContext remote context on a separate computer.
	//
	// CLSCTX_REMOTE_SERVER in COM API.
	RemoteServerClassContext ClassContext = 0x10

	// InProcessHandler16ClassContext is obsolete.
	//
	// CLSCTX_INPROC_HANDLER16 in COM API.
	InProcessHandler16ClassContext ClassContext = 0x20

	// DisableCodeDownloadClassContext disables downloading code from active
	// directory or Internet.
	//
	// CLSCTX_NO_CODE_DOWNLOAD in COM API.
	DisableCodeDownloadClassContext ClassContext = 0x400

	// NoCustomMarshalClassContext fail if activation uses custom marshalling.
	//
	// CLSCTX_NO_CUSTOM_MARSHAL in COM API.
	NoCustomMarshalClassContext ClassContext = 0x1000

	// EnableCodeDownloadClassContext enable downloading code from active
	// directory or Internet.
	//
	// CLSCTX_ENABLE_CODE_DOWNLOAD in COM API.
	EnableCodeDownloadClassContext ClassContext = 0x2000

	// NoFailureLogClassContext override logging of failures in
	// CoCreateInstanceEx().
	//
	// CLSCTX_NO_FAILURE_LOG in COM API.
	NoFailureLogClassContext ClassContext = 0x4000

	// DisableAAAClassContext disables launching as callers identity for this
	// activation only.
	//
	// This prevents untrusted components from gaining higher privileges.
	//
	// CLSCTX_DISABLE_AAA in COM API.
	DisableAAAClassContext ClassContext = 0x8000

	// EnableAAAClassContext enables launching as callers identity for this
	// activation only.
	//
	// CLSCTX_ENABLE_AAA in COM API.
	EnableAAAClassContext ClassContext = 0x10000

	// FromDefaultContextClassContext uses default context of current apartment.
	//
	// CLSCTX_FROM_DEFAULT_CONTEXT in COM API.
	FromDefaultContextClassContext ClassContext = 0x20000

	// ActivateServer32ClassContext activate or connect 32-bit version of
	// server, if one is registered or fail.
	//
	// CLSCTX_ACTIVATE_32_BIT_SERVER in COM API.
	ActivateServer32ClassContext ClassContext = 0x40000

	// ActivateServer64ClassContext activate or connect 64-bit version of
	// server, if one is registered or fail.
	//
	// CLSCTX_ACTIVATE_64_BIT_SERVER in COM API.
	ActivateServer64ClassContext ClassContext = 0x80000

	// EnableCloakingClassContext uses impersonation token of thread instead of
	// process token.
	//
	// Supported by Windows Vista or later.
	//
	// CLSCTX_ENABLE_CLOAKING in COM API.
	EnableCloakingClassContext ClassContext = 0x100000

	// ActivateAAAAsIUClassContext uses interactive user behavior to activate server.
	//
	// CLSCTX_ACTIVATE_AAA_AS_IU in COM API.
	ActivateAAAAsIUClassContext ClassContext = 0x800000

	// AllLocalClassContext enable all flags for local COM server.
	//
	// Supports local in process, local in process handler and local separate
	// process.
	AllLocalClassContext = InProcessServerClassContext | InProcessHandlerClassContext | LocalServerClassContext

	// AllInProcessClassContext enable all flags for local in process COM server.
	//
	// Supports local in process and local in process handler.
	AllInProcessClassContext = InProcessServerClassContext | InProcessHandlerClassContext

	// AllServerClassContext enable all flags for local in process COM server.
	//
	// Supports local in process and local in process handler.
	AllServerClassContext = InProcessServerClassContext | LocalServerClassContext | RemoteServerClassContext
)

const (
	// MultithreadedConcurrencyModel is the default concurrency model.
	//
	// Initializes thread for multithreaded object concurrency. This can not be
	// used with ApartmentThreadedConcurrencyModel.
	//
	// Allows calls to be run on any thread.
	//
	// Not to be used with UI actions or threads.
	//
	// Known as COINIT_MULTITHREADED in COM API.
	MultithreadedConcurrencyModel ConcurrencyModel = iota // 0x0

	// ApartmentThreadedConcurrencyModel initializes thread for apartment
	// threading.
	//
	// Allows for multiple threads by serializing all incoming calls and
	// requires calls to methods created by this thread run on the same thread.
	//
	// This is probably highly unsafe with Go's threading model. Given that with
	// gothreads, it is unknown which thread the gothread will run making the
	// use of gothreads unusable, unless MAX_CONCURRENT_THREADS is set to 1.
	//
	// Known as COINIT_APARTMENTTHREADED in COM API.
	ApartmentThreadedConcurrencyModel ConcurrencyModel = 1 << iota // 0x2

	// DisableOLE1DDEConcurrencyModel disables DDE for OLE1 support.
	//
	// Only useful for OLE1.
	//
	// Known as COINIT_DISABLE_OLE1DDE in COM API.
	DisableOLE1DDEConcurrencyModel ConcurrencyModel = 1 << iota // 0x4

	// SpeedOverMemoryConcurrencyModel uses more memory to increase performance.
	//
	// Known as COINIT_SPEED_OVER_MEMORY in COM API.
	SpeedOverMemoryConcurrencyModel ConcurrencyModel = 1 << iota // 0x8
)

const (
	// MethodDispatchContext invoked as method.
	//
	// If a property with the same name exists, then both this and
	// PropertyGetDispatchContext can be set.
	//
	// DISPATCH_METHOD in COM API.
	MethodDispatchContext DispatchContext = 1 << iota // 1

	// PropertyGetDispatchContext access property value.
	//
	// DISPATCH_PROPERTYGET in COM API.
	PropertyGetDispatchContext

	// PropertySetDispatchContext mutate property value assignment.
	//
	// DISPATCH_PROPERTYPUT in COM API.
	PropertySetDispatchContext

	// PropertySetRefDispatchContext mutate property by reference assignment.
	//
	// This is only valid when property accepts a reference to an object.
	//
	// DISPATCH_PROPERTYPUTREF in COM API.
	PropertySetRefDispatchContext
)

const (
	// SuccessResponseCode is S_OK
	SuccessResponseCode HResultResponseCode = 0x00000000

	// UnexpectedErrorCode is E_UNEXPECTED
	UnexpectedErrorCode HResultResponseCode = 0x8000FFFF

	// NotImplementedErrorCode is E_NOTIMPL
	NotImplementedErrorCode HResultResponseCode = 0x80004001

	// OutOfMemoryErrorCode is E_OUTOFMEMORY
	OutOfMemoryErrorCode HResultResponseCode = 0x8007000E

	// InvalidArgumentErrorCode is E_INVALIDARG
	InvalidArgumentErrorCode HResultResponseCode = 0x80070057

	// NoInterfaceErrorCode is E_NOINTERFACE
	NoInterfaceErrorCode HResultResponseCode = 0x80004002

	// PointerErrorCode is E_POINTER
	PointerErrorCode HResultResponseCode = 0x80004003

	// HandleErrorCode is E_HANDLE
	HandleErrorCode HResultResponseCode = 0x80070006

	// AbortErrorCode is E_ABORT
	AbortErrorCode HResultResponseCode = 0x80004004

	// FailureErrorCode is E_FAIL
	FailureErrorCode HResultResponseCode = 0x80004005

	// AccessDeniedErrorCode is E_ACCESSDENIED
	AccessDeniedErrorCode HResultResponseCode = 0x80070005

	// PendingErrorCode is E_PENDING
	PendingErrorCode HResultResponseCode = 0x8000000A

	// COMObjectClassStringErrorCode is CO_E_CLASSSTRING
	COMObjectClassStringErrorCode = 0x800401F3
)

const (
	// FastCallingConvention is FastCall calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_FASTCALL
	FastCallingConvention CallingConvention = iota

	// CDeclCallingConvention is CDecl calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_CDECL
	CDeclCallingConvention

	// MSCPascalCallingConvention is MSCPascal calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_MSCPASCAL
	MSCPascalCallingConvention

	// PascalCallingConvention is Pascal calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_PASCAL
	PascalCallingConvention = MSCPascalCallingConvention

	// MacPascalCallingConvention is MacPascal calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_MACPASCAL
	MacPascalCallingConvention

	// StandardCallingConvention is StdCall calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_STDCALL
	StandardCallingConvention

	// FPFastCallingConvention is FPFast calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_FPFASTCALL
	FPFastCallingConvention

	// SysCallingConvention is SysCall calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_SYSCALL
	SysCallingConvention

	// MPWCDeclCallingConvention is MPWCDECL calling convention.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_MPWCDECL
	MPWCDeclCallingConvention

	// MPWPascalCallingConvention is MPWPascal calling convention
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_MPWPASCAL
	MPWPascalCallingConvention

	// MaxCallingConvention is maximum calling convention value equal to
	// MPWPascal.
	//
	// There is no summary for this and it is assumed that you know what this
	// is.
	//
	// CC_MAX
	MaxCallingConvention = MPWPascalCallingConvention
)

const (
	// EmptyVariantType is intentionally empty variant.
	//
	// This does not mean null. It is either '0' (zero) for numeric types and
	// "" (empty string) for string types.
	//
	// VT_EMPTY in COM API.
	EmptyVariantType VariantType = 0x0

	// NullVariantType is a null variant.
	//
	// VT_NULL in COM API.
	NullVariantType VariantType = 0x1

	// Float32VariantType is 32-bit float variant.
	//
	// VT_R4 in COM API.
	Float32VariantType VariantType = 0x4

	// Float64VariantType is 64-bit float variant.
	//
	// VT_R8 in COM API.
	Float64VariantType VariantType = 0x5

	// CurrencyVariantType is a currency variant.
	//
	// VT_CY in COM API.
	CurrencyVariantType VariantType = 0x6

	// DateVariantType is a date variant.
	//
	// VT_DATE in COM API.
	DateVariantType VariantType = 0x7

	// BinaryStringVariantType is a binary string variant.
	//
	// VT_BSTR in COM API.
	BinaryStringVariantType VariantType = 0x8

	// IDispatchVariantType is a IDispatch variant.
	//
	// VT_DISPATCH in COM API.
	IDispatchVariantType VariantType = 0x9

	// ErrorVariantType is a error variant.
	//
	// VT_ERROR in COM API.
	ErrorVariantType VariantType = 0xa

	// BoolVariantType is a bool variant.
	//
	// VT_BOOL in COM API.
	BoolVariantType VariantType = 0xb

	// VariantVariantType is a variant containing a variant.
	//
	// VT_VARIANT in COM API.
	VariantVariantType VariantType = 0xc

	// IUnknownVariantType is a IUnknown variant.
	//
	// VT_UNKNOWN in COM API.
	IUnknownVariantType VariantType = 0xd

	// DecimalVariantType is a decimal variant.
	//
	// VT_DECIMAL in COM API.
	DecimalVariantType VariantType = 0xe

	// Integer8VariantType is character variant.
	//
	// VT_I1 in COM API.
	Integer8VariantType VariantType = 0x10

	// UInteger8VariantType is 8-bit unsigned integer variant.
	//
	// VT_UI1 in COM API.
	UInteger8VariantType VariantType = 0x11

	// Integer16VariantType is 16-bit integer variant.
	//
	// VT_I2 in COM API.
	Integer16VariantType VariantType = 0x2

	// UInteger16VariantType is 16-bit unsigned integer variant.
	//
	// VT_UI2 in COM API.
	UInteger16VariantType VariantType = 0x12

	// Integer32VariantType is 32-bit integer variant.
	//
	// VT_I4 in COM API.
	Integer32VariantType VariantType = 0x3

	// UInteger32VariantType is 32-bit unsigned integer variant.
	//
	// VT_UI4 in COM API.
	UInteger32VariantType VariantType = 0x13

	// Integer64VariantType is 64-bit integer variant.
	//
	// VT_I8 in COM API.
	Integer64VariantType VariantType = 0x14

	// UInteger64VariantType is 64-bit unsigned integer variant.
	//
	// VT_UI8 in COM API.
	UInteger64VariantType VariantType = 0x15

	// IntegerVariantType is 32-bit integer variant.
	//
	// VT_INT in COM API.
	IntegerVariantType VariantType = 0x16

	// UIntegerVariantType is 32-bit unsigned integer variant.
	//
	// VT_UINT in COM API.
	UIntegerVariantType VariantType = 0x17

	// VoidVariantType is a void variant.
	//
	// C-style void.
	//
	// VT_VOID in COM API.
	VoidVariantType VariantType = 0x18

	// HResultVariantType is a HRESULT variant.
	//
	// VT_HRESULT in COM API.
	HResultVariantType VariantType = 0x19

	// PointerVariantType is a pointer variant.
	//
	// VT_PTR in COM API.
	PointerVariantType VariantType = 0x1a

	// SafeArrayVariantType is a safe array variant.
	//
	// VT_SAFEARRAY in COM API.
	SafeArrayVariantType VariantType = 0x1b

	// CArrayVariantType is a C array variant.
	//
	// VT_CARRAY in COM API.
	CArrayVariantType VariantType = 0x1c

	// UserDefinedVariantType is user defined variant.
	//
	// VT_USERDEFINED in COM API.
	UserDefinedVariantType VariantType = 0x1d

	// ANSIStringVariantType is a string type variant.
	//
	// Null terminated string.
	//
	// VT_LPSTR in COM API.
	ANSIStringVariantType VariantType = 0x1e

	// UnicodeStringVariantType is a string type variant.
	//
	// Null terminated wide string.
	//
	// VT_LPWSTR in COM API.
	UnicodeStringVariantType VariantType = 0x1f

	// RecordVariantType is a record variant.
	//
	// VT_RECORD in COM API.
	RecordVariantType VariantType = 0x24

	// IntegerPointerVariantType is a pointer variant.
	//
	// VT_INT_PTR in COM API.
	IntegerPointerVariantType VariantType = 0x25

	// UIntegerPointerVariantType is a pointer variant.
	//
	// VT_UINT_PTR in COM API.
	UIntegerPointerVariantType VariantType = 0x26

	// FileTimeVariantType is a FileTime variant.
	//
	// VT_FILETIME in COM API.
	FileTimeVariantType VariantType = 0x40

	// BlobVariantType is a blob (byte array) variant.
	//
	// VT_BLOB in COM API.
	BlobVariantType VariantType = 0x41

	// StreamVariantType is a stream variant.
	//
	// VT_STREAM in COM API.
	StreamVariantType VariantType = 0x42

	// StorageVariantType is a storage variant.
	//
	// VT_STORAGE in COM API.
	StorageVariantType VariantType = 0x43

	// StreamObjectVariantType is a streaming object variant.
	//
	// VT_STREAMED_OBJECT in COM API.
	StreamObjectVariantType VariantType = 0x44

	// StoredObjectVariantType is a stored object variant.
	//
	// VT_STORED_OBJECT in COM API.
	StoredObjectVariantType VariantType = 0x45

	// BlobObjectVariantType is a blob object variant.
	//
	// VT_BLOB_OBJECT in COM API.
	BlobObjectVariantType VariantType = 0x46

	// ClipboardFormatVariantType is a clipboard format variant.
	//
	// VT_CF in COM API.
	ClipboardFormatVariantType VariantType = 0x47

	// ClassIDVariantType is a ClassID string variant.
	//
	// VT_CLSID in COM API.
	ClassIDVariantType VariantType = 0x48

	// BinaryStringBlobVariantType reserved for system use.
	//
	// VT_BSTR_BLOB in COM API.
	BinaryStringBlobVariantType VariantType = 0xfff

	// VectorVariantType is a vector variant.
	//
	// VT_VECTOR in COM API.
	VectorVariantType VariantType = 0x1000

	// ArrayVariantType is an array variant.
	//
	// VT_ARRAY in COM API.
	ArrayVariantType VariantType = 0x2000

	// ByReferenceVariantType is a by referenced masked variant.
	//
	// VT_BYREF in COM API.
	ByReferenceVariantType VariantType = 0x4000

	// ReservedVariantType is not a variant type that can be set manually.
	//
	// VT_RESERVED in COM API.
	ReservedVariantType VariantType = 0x8000

	// IllegalVariantType is not a variant type that can be set manually.
	//
	// VT_ILLEGAL in COM API.
	IllegalVariantType VariantType = 0xffff

	// IllegalMaskedVariantType is not a variant type that can be set manually.
	//
	// VT_ILLEGALMASKED in COM API.
	IllegalMaskedVariantType VariantType = 0xfff

	// TypeMaskVariantType is not a variant type that can be set manually.
	//
	// VT_TYPEMASK in COM API.
	TypeMaskVariantType VariantType = 0xfff
)

// Reference: https://msdn.microsoft.com/en-us/library/windows/desktop/ms221242(v=vs.85).aspx
const (
	// CollectDisplayID collect property. Used if calling method is an accessor.
	//
	// DISPID_COLLECT in COM API.
	CollectDisplayID DisplayID = -8

	// DestructorDisplayID C++ destructor function.
	//
	// DISPID_DESTRUCTOR in COM API.
	DestructorDisplayID DisplayID = -7

	// ConstructorDisplayID C++ constructor function.
	//
	// DISPID_DESTRUCTOR in COM API.
	ConstructorDisplayID DisplayID = -6

	// EvaluateDisplayID evaluate method.
	//
	// Implicitly invoked when ActiveX client encloses arguments in square
	// brackets.
	//
	// DISPID_EVALUATE in COM API.
	EvaluateDisplayID DisplayID = -5

	// NewEnumDisplayID _NewEnum property.
	//
	// Required for collection objects. Returns Enumerator object supporting
	// IEnumVariant.
	//
	// DISPID_NEWENUM in COM API.
	NewEnumDisplayID DisplayID = -4

	// PropertySetDisplayID parameter that recieves value of an assignment.
	//
	// DISPID_PROPERTYPUT in COM API.
	PropertySetDisplayID DisplayID = -3

	// UnknownDisplayID value returned by IDispatch.GetIDsOfNames when match
	// failed.
	//
	// DISPID_UNKNOWN in COM API.
	UnknownDisplayID DisplayID = -1

	// ValueDisplayID default member for object.
	//
	// Invoked when ActiveX client when missing property or method for object
	// name.
	//
	// DISPID_VALUE in COM API.
	ValueDisplayID DisplayID = 0
)

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms221643(v=vs.85).aspx
const (
	// EnumKindType is enumerator type.
	//
	// TKIND_ENUM in COM API.
	EnumKindType KindType = iota

	// RecordKindType is a type with no methods.
	//
	// TKIND_RECORD in COM API.
	RecordKindType

	// ModuleKindType has only static methods and data.
	//
	// TKIND_MODULE in COM API.
	ModuleKindType

	// InterfaceKindType has virtual and pure methods.
	//
	// TKIND_INTERFACE in COM API.
	InterfaceKindType

	// IDispatchKindType methods and properties can be accessed through
	// IDispatch invoke.
	//
	// TKIND_DISPATCH in COM API.
	IDispatchKindType

	// CoClassKindType has component object interfaces.
	//
	// TKIND_COCLASS in COM API.
	CoClassKindType

	// AliasKindType is an alias for another type.
	//
	// TKIND_ALIAS in COM API.
	AliasKindType

	// UnionKindType is an union type whose members all have offset of zero.
	//
	// TKIND_UNION in COM API.
	UnionKindType

	// MaxKindType is the end of enum marker.
	//
	// TKIND_MAX in COM API.
	MaxKindType
)

// Safe Array Feature Flags
const (
	// AutoSafeArrayMask is allocated on the stack.
	//
	// FADF_AUTO in COM API.
	AutoSafeArrayMask SafeArrayMask = 0x0001

	// StaticSafeArrayMask is statically allocated.
	//
	// FADF_STATIC in COM API.
	StaticSafeArrayMask SafeArrayMask = 0x0002

	// EmbeddedSafeArrayMask is embedded in a structure.
	//
	// FADF_EMBEDDED in COM API.
	EmbeddedSafeArrayMask SafeArrayMask = 0x0004

	// FixedSizeSafeArrayMask may not be resized or reallocated.
	//
	// FADF_FIXEDSIZE in COM API.
	FixedSizeSafeArrayMask SafeArrayMask = 0x0010

	// RecordSafeArrayMask is a record array.
	//
	// Contains an IRecordInfo at array offset -4.
	//
	// FADF_RECORD in COM API.
	RecordSafeArrayMask SafeArrayMask = 0x0020

	// HasInterfaceIDSafeArrayMask has IID self identifying interface.
	//
	// GUID at array offset -16.
	//
	// Requires IDispatchSafeArrayMask or IUnknownSafeArrayMask.
	//
	// FADF_HAVEIID in COM API.
	HasInterfaceIDSafeArrayMask SafeArrayMask = 0x0040

	// HasVariantSafeArrayMask is an array that has variant type.
	//
	// FADF_HAVEVARTYPE in COM API.
	HasVariantSafeArrayMask SafeArrayMask = 0x0080

	// BinaryStringSafeArrayMask is an array of binary strings.
	//
	// FADF_BSTR in COM API.
	BinaryStringSafeArrayMask SafeArrayMask = 0x0100

	// IUnknownSafeArrayMask is an array of IUnknown.
	//
	// FADF_UNKNOWN in COM API.
	IUnknownSafeArrayMask SafeArrayMask = 0x0200

	// IDispatchSafeArrayMask is an array IDispatch.
	//
	// FADF_DISPATCH in COM API.
	IDispatchSafeArrayMask SafeArrayMask = 0x0400

	// VariantSafeArrayMask is an array of Variants.
	//
	// FADF_VARIANT in COM API.
	VariantSafeArrayMask SafeArrayMask = 0x0800

	// ReservedSafeArrayMask is reserved for future use.
	//
	// FADF_RESERVED in COM API.
	ReservedSafeArrayMask SafeArrayMask = 0xF008
)
