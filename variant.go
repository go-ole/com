package com

import (
	"time"
	"unsafe"

	"github.com/go-ole/types/idispatch"
	"github.com/go-ole/types/iunknown"
)

// NewVariant returns new variant based on type and value.
func NewVariant(vt VariantType, val int64) Variant {
	return Variant{VariantType: vt, Val: val}
}

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
func (v *Variant) ToArray() *SafeArrayConversion {
	if v.VariantType != SafeArrayVariantType {
		return nil
	}
	safeArray := (*SafeArray)(unsafe.Pointer(uintptr(v.Val)))
	return &SafeArrayConversion{safeArray}
}

// ToString converts variant to Go string.
func (v *Variant) ToString() string {
	if v.VariantType != BinaryStringVariantType {
		return ""
	}
	return BstrToString(*(**uint16)(unsafe.Pointer(&v.Val)))
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

// Clear the memory of variant object.
func (v *Variant) Clear() error {
	return VariantClear(v)
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
	case SafeArrayVariantType, ArrayVariantType:
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
		//case DateVariantType:
		// TODO: use VariantTimeToSystemTime
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

// VariantByValueType creates new variant by value type.
func VariantByValueType(value interface{}) Variant {
	var variant Variant
	VariantInit(variant)

	switch vv := value.(type) {
	case bool:
		if vv {
			variant = NewVariant(BoolVariantType, 0xffff)
		} else {
			variant = NewVariant(BoolVariantType, 0)
		}
	case *bool:
		variant = NewVariant(BoolVariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*bool)))))
	case byte, int8:
		variant = NewVariant(Integer8VariantType, int64(value.(byte)))
	case *byte, *int8:
		variant = NewVariant(Integer8VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*byte)))))
	case uint8:
		variant = NewVariant(UInteger8VariantType, int64(value.(uint8)))
	case *uint8:
		variant = NewVariant(UInteger8VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*uint8)))))
	case int16:
		variant = NewVariant(Integer16VariantType, int64(value.(int16)))
	case *int16:
		variant = NewVariant(Integer16VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*int16)))))
	case uint16:
		variant = NewVariant(UInteger16VariantType, int64(value.(uint16)))
	case *uint16:
		variant = NewVariant(UInteger16VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*uint16)))))
	case int, int32:
		variant = NewVariant(Integer32VariantType, int64(value.(int)))
	case *int, *int32:
		variant = NewVariant(Integer32VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*int)))))
	case uint, uint32:
		variant = NewVariant(UInteger32VariantType, int64(value.(uint)))
	case *uint, *uint32:
		variant = NewVariant(UInteger32VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*uint)))))
	case int64:
		variant = NewVariant(Integer64VariantType, int64(value.(int64)))
	case *int64:
		variant = NewVariant(Integer64VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*int64)))))
	case uint64:
		variant = NewVariant(UInteger64VariantType, value.(int64))
	case *uint64:
		variant = NewVariant(UInteger64VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*uint64)))))
	case float32:
		variant = NewVariant(Float32VariantType, *(*int64)(unsafe.Pointer(&vv)))
	case *float32:
		variant = NewVariant(Float32VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*float32)))))
	case float64:
		variant = NewVariant(Float64VariantType, *(*int64)(unsafe.Pointer(&vv)))
	case *float64:
		variant = NewVariant(Float64VariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*float64)))))
	case string:
		variant = NewVariant(BinaryStringVariantType, int64(uintptr(unsafe.Pointer(SysAllocStringLen(value.(string))))))
	case *string:
		variant = NewVariant(BinaryStringVariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*string)))))
	case time.Time:
		s := vv.Format("2006-01-02 15:04:05")
		variant = NewVariant(BinaryStringVariantType, int64(uintptr(unsafe.Pointer(SysAllocStringLen(s)))))
	case *time.Time:
		s := vv.Format("2006-01-02 15:04:05")
		variant = NewVariant(BinaryStringVariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(&s))))
	case *iunknown.Unknown:
		variant = NewVariant(IUnknownVariantType, int64(uintptr(unsafe.Pointer(v.(*iunknown.Unknown)))))
	case **iunknown.Unknown:
		variant = NewVariant(IUnknownVariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(v.(**iunknown.Unknown)))))
	case *idispatch.Dispatch:
		variant = NewVariant(IDispatchVariantType, int64(uintptr(unsafe.Pointer(v.(*idispatch.Dispatch)))))
	case **idispatch.Dispatch:
		variant = NewVariant(IDispatchVariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(v.(**idispatch.Dispatch)))))
	case nil:
		variant = NewVariant(NullVariantType, 0)
	case *Variant:
		variant = NewVariant(VariantVariantType|ByReferenceVariantType, int64(uintptr(unsafe.Pointer(v.(*Variant)))))
	case []byte:
		safeByteArray := SafeArrayFromByteSlice(v.([]byte))
		variant = NewVariant(ArrayVariantType|UInteger8VariantType, int64(uintptr(unsafe.Pointer(safeByteArray))))
		defer VariantClear(&variant)
	default:
		panic("unknown type")
	}

	return variant
}
