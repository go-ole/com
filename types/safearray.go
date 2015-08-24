package types

import "errors"

var (
	NotSliceError              = errors.New("must be a slice")
	VariantNotImplementedError = errors.New("variant type is not implemented")
	VariantNotSupportedError   = errors.New("variant type is not supported")
	SafeArrayVectorFailedError = errors.New("could not convert slice to SafeArray")
)

// COMArray is how COM handles arrays.
type COMArray struct {
	Dimensions   uint16
	FeaturesFlag uint16
	ElementsSize uint32
	LocksAmount  uint32
	Data         uint32

	// This must hold the bytes for two Bounds objects. Use binary.Read() to
	// get the contents.
	Bounds [16]byte
}

// Bounds defines the array boundaries.
type Bounds struct {
	Elements   uint32
	LowerBound int32
}
