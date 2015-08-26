package types

import (
	"unsafe"

	"github.com/go-ole/com/safearray"
	api "github.com/go-ole/com/variant"
	"github.com/go-ole/types/idispatch"
	"github.com/go-ole/types/iunknown"
)

// ToIUnknown converts Variant to Unknown object.
func (v *Variant) ToIUnknown() *iunknown.Unknown {
	if v.VariantType != IUnknownVariantType {
		return nil
	}
	return (*iunknown.Unknown)(unsafe.Pointer(uintptr(v.Val)))
}

// ToIDispatch converts variant to dispatch object.
func (v *Variant) ToIDispatch() *idispatch.Dispatch {
	if v.VariantType != IDispatchVariantType {
		return nil
	}
	return (*idispatch.Dispatch)(unsafe.Pointer(uintptr(v.Val)))
}

// ToArray converts variant to SafeArray helper.
func (v *Variant) ToArray() *safearray.Array {
	if v.VariantType != SafeArrayVariantType {
		return nil
	}
	safeArray := (*COMArray)(unsafe.Pointer(uintptr(v.Val)))
	return &safearray.Array{Array: safeArray}
}

// ToString converts variant to Go string.
func (v *Variant) ToString() string {
	if v.VariantType != BinaryStringVariantType {
		return ""
	}
	return util.BstrToString(*(**uint16)(unsafe.Pointer(&v.Val)))
}

// ToCurrency converts variant to currency object.
func (v *Variant) ToCurrency() *Currency {
	if v.VariantType != CurrencyVariantType {
		return nil
	}
	return (*Currency)(unsafe.Pointer(uintptr(v.Val)))
}

// ToDecimal converts variant to decimal object.
func (v *Variant) ToDecimal() *Decimal {
	if v.VariantType != DecimalVariantType {
		return nil
	}
	return (*Decimal)(unsafe.Pointer(uintptr(v.Val)))
}

// ToVariant converts variant to another variant.
func (v *Variant) ToVariant() *Variant {
	if v.VariantType != VariantVariantType {
		return nil
	}
	return (*Variant)(unsafe.Pointer(uintptr(v.Val)))
}

// ToFileTime converts variant to FileTime object.
func (v *Variant) ToFileTime() *FileTime {
	if v.VariantType != FileTimeVariantType {
		return nil
	}
	return (*FileTime)(unsafe.Pointer(uintptr(v.Val)))
}

func (v *Variant) ToDateTime() interface{} {
	return api.VariantTimeToSystemTime(v)
}

// Clear the memory of variant object.
func (v *Variant) Clear() error {
	return api.VariantClear(v)
}

// Value returns variant value based on its type.
//
// Currently supported types: 2- and 4-byte integers, strings, bools.
// Note that 64-bit integers, datetimes, and other types are stored as strings
// and will be returned as strings.
//
// Needs to be further converted, because this returns an interface{}.
func (v *Variant) Value() interface{} {
	switch v.VariantType {
	case Integer8VariantType:
		return int8(v.Val)
	case UInteger8VariantType:
		return uint8(v.Val)
	case Integer16VariantType:
		return int16(v.Val)
	case UInteger16VariantType:
		return uint16(v.Val)
	case Integer32VariantType, IntegerVariantType:
		return int32(v.Val)
	case UInteger32VariantType, UIntegerVariantType:
		return uint32(v.Val)
	case Integer64VariantType:
		return int64(v.Val)
	case UInteger64VariantType:
		return uint64(v.Val)
	case Float32VariantType:
		return float32(v.Val)
	case Float64VariantType:
		return float64(v.Val)
	case PointerVariantType, IntegerPointerVariantType, UIntegerPointerVariantType, UserDefinedVariantType:
		return uintptr(v.Val)
	case BinaryStringVariantType:
		return v.ToString()
	case IUnknownVariantType:
		return v.ToIUnknown()
	case IDispatchVariantType:
		return v.ToIDispatch()
	case BoolVariantType:
		return v.Val != 0
	case NullVariantType:
		return nil
	case EmptyVariantType:
		return nil // TODO: This should return either 0 for ints or "" for strings.
	case VoidVariantType:
		// TODO: This might not actually mean nil, but means variant doesn't
		// take value.
		return nil
	case ErrorVariantType:
		return NewError(uintptr(v.Val))
	case HResultVariantType:
		if v.Val != 0 {
			return NewError(uintptr(v.Val))
		}
		return nil
	case SafeArrayVariantType:
		return v.ToArray()
	case ArrayVariantType:
		// ArrayVariantType is C Array, but difficult since you need to know
		// length and size.
		return v.ToArray()
	case CurrencyVariantType:
		return v.ToCurrency()
	case DecimalVariantType:
		return v.ToDecimal()
	case FileTimeVariantType:
		// Might not be applicable for Variant.
		return v.ToFileTime()
	case VariantVariantType:
		return v.ToVariant()
	case DateVariantType:
		return v.ToDateTime()
		// RecordVariantType:
		// TODO: Implement.
		// case ANSIStringVariantType:
		// Might not be applicable for Variant.
		// TODO: Convert ANSI string to Go.
		// case UnicodeStringVariantType:
		// Might not be applicable for Variant.
		// TODO: Convert Unicode string to Go.
	}

	return nil
}
