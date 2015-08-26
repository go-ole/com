package com

import (
	"time"
	"unsafe"

	api "github.com/go-ole/com/internal/variant"
	"github.com/go-ole/com/safearray"
	"github.com/go-ole/com/types"
	"github.com/go-ole/types/idispatch"
	"github.com/go-ole/types/iunknown"
)

// NewVariant returns new variant based on type and value.
func NewVariant(vt types.VariantType, val int64) types.Variant {
	return types.Variant{VariantType: vt, Val: val}
}

// VariantByValueType creates new variant by value type.
func VariantByValueType(value interface{}) types.Variant {
	var variant types.Variant
	VariantInit(variant)

	switch vv := value.(type) {
	case bool:
		if vv {
			variant = NewVariant(types.BoolVariantType, 0xffff)
		} else {
			variant = NewVariant(types.BoolVariantType, 0)
		}
	case *bool:
		variant = NewVariant(types.BoolVariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*bool)))))
	case byte, int8:
		variant = NewVariant(types.Integer8VariantType, int64(value.(byte)))
	case *byte, *int8:
		variant = NewVariant(types.Integer8VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*byte)))))
	case uint8:
		variant = NewVariant(types.UInteger8VariantType, int64(value.(uint8)))
	case *uint8:
		variant = NewVariant(types.UInteger8VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*uint8)))))
	case int16:
		variant = NewVariant(types.Integer16VariantType, int64(value.(int16)))
	case *int16:
		variant = NewVariant(types.Integer16VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*int16)))))
	case uint16:
		variant = NewVariant(types.UInteger16VariantType, int64(value.(uint16)))
	case *uint16:
		variant = NewVariant(types.UInteger16VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*uint16)))))
	case int, int32:
		variant = NewVariant(types.Integer32VariantType, int64(value.(int)))
	case *int, *int32:
		variant = NewVariant(types.Integer32VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*int)))))
	case uint, uint32:
		variant = NewVariant(types.UInteger32VariantType, int64(value.(uint)))
	case *uint, *uint32:
		variant = NewVariant(types.UInteger32VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*uint)))))
	case int64:
		variant = NewVariant(types.Integer64VariantType, int64(value.(int64)))
	case *int64:
		variant = NewVariant(types.Integer64VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*int64)))))
	case uint64:
		variant = NewVariant(types.UInteger64VariantType, value.(int64))
	case *uint64:
		variant = NewVariant(types.UInteger64VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*uint64)))))
	case float32:
		variant = NewVariant(types.Float32VariantType, *(*int64)(unsafe.Pointer(&vv)))
	case *float32:
		variant = NewVariant(types.Float32VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*float32)))))
	case float64:
		variant = NewVariant(types.Float64VariantType, *(*int64)(unsafe.Pointer(&vv)))
	case *float64:
		variant = NewVariant(types.Float64VariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*float64)))))
	case string:
		variant = NewVariant(types.BinaryStringVariantType, int64(uintptr(unsafe.Pointer(SysAllocStringLen(value.(string))))))
	case *string:
		variant = NewVariant(types.BinaryStringVariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(value.(*string)))))
	case time.Time:
		s := vv.Format("2006-01-02 15:04:05")
		variant = NewVariant(types.BinaryStringVariantType, int64(uintptr(unsafe.Pointer(SysAllocStringLen(s)))))
	case *time.Time:
		s := vv.Format("2006-01-02 15:04:05")
		variant = NewVariant(types.BinaryStringVariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(&s))))
	case *iunknown.Unknown:
		variant = NewVariant(types.IUnknownVariantType, int64(uintptr(unsafe.Pointer(v.(*iunknown.Unknown)))))
	case **iunknown.Unknown:
		variant = NewVariant(types.IUnknownVariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(v.(**iunknown.Unknown)))))
	case *idispatch.Dispatch:
		variant = NewVariant(types.IDispatchVariantType, int64(uintptr(unsafe.Pointer(v.(*idispatch.Dispatch)))))
	case **idispatch.Dispatch:
		variant = NewVariant(types.IDispatchVariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(v.(**idispatch.Dispatch)))))
	case nil:
		variant = NewVariant(types.NullVariantType, 0)
	case *Variant:
		variant = NewVariant(types.VariantVariantType|types.ByReferenceVariantType, int64(uintptr(unsafe.Pointer(v.(*Variant)))))
	case []byte:
		safeByteArray := safearray.SafeArrayFromByteSlice(v.([]byte))
		variant = NewVariant(types.ArrayVariantType|types.UInteger8VariantType, int64(uintptr(unsafe.Pointer(safeByteArray))))
		defer VariantClear(&variant)
	default:
		panic("unknown type")
	}

	return variant
}

// VariantInit initializes memory for variant objects.
func VariantInit(v *types.Variant) error {
	return api.VariantInit(v)
}

// VariantClear clears the memory of variant objects.
func VariantClear(v *types.Variant) error {
	return api.VariantClear(v)
}

// VariantChangeType changes variant type.
func VariantChangeType(source *types.Variant, flags uint16, vt types.VariantType) *types.Variant {
	return api.VariantChangeType(source, flags, vt)
}

// VariantChangeTypeEx changes variant type.
func VariantChangeTypeEx(source *types.Variant, locale uint32, flags uint16, vt types.VariantType) *types.Variant {
	return api.VariantChangeTypeEx(source, locale, flags, vt)
}

// VariantCopy copies variant.
func VariantCopy(source *types.Variant) *types.Variant {
	return api.VariantCopy(source)
}

// VariantCopyIndirect copies variant reference.
func VariantCopyIndirect(source *types.Variant) *types.Variant {
	return api.VariantCopyIndirect(source)
}
