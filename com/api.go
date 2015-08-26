package variant

import api "github.com/go-ole/com/internal/variant"
import "github.com/go-ole/com/types"

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
