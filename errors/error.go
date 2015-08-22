package errors

// HResultResponseCode are success or error codes.
type HResultResponseCode uint32

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
	COMObjectClassStringErrorCode HResultResponseCode = 0x800401F3
)

// NotImplementedError is shortcode to OleError not implemented error type.
var NotImplementedError = NewError(NotImplementedErrorCode)

// HResultToError may return an error from a HResult or nil, no error.
func HResultToError(hr uintptr, _ uintptr, _ error) (err error) {
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

// OleError stores COM errors.
type OleError struct {
	// HResult code or pointer.
	hr uintptr

	// Error description.
	description string

	// Parent COM error.
	subError error
}

// NewError creates new error with HResult.
func NewError(hr uintptr) *OleError {
	return &OleError{hr: hr}
}

// NewErrorWithDescription creates new COM error with HResult and description.
func NewErrorWithDescription(hr uintptr, description string) *OleError {
	return &OleError{hr: hr, description: description}
}

// NewErrorWithSubError creates new COM error with parent error.
func NewErrorWithSubError(hr uintptr, description string, err error) *OleError {
	return &OleError{hr: hr, description: description, subError: err}
}

// Code is the HResult.
func (v *OleError) Code() uintptr {
	return uintptr(v.hr)
}

// String description, either manually set or format message with error code.
func (v *OleError) String() string {
	if v.description != "" {
		return errstr(int(v.hr)) + " (" + v.description + ")"
	}
	return errstr(int(v.hr))
}

// Error implements error interface.
func (v *OleError) Error() string {
	return v.String()
}

// Description retrieves error summary, if there is one.
func (v *OleError) Description() string {
	return v.description
}

// SubError returns parent error, if there is one.
func (v *OleError) SubError() error {
	return v.subError
}
