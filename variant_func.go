// +build !windows,!cgo

package com

// VariantInit initializes memory for variant objects.
func VariantInit(v *Variant) error {
	// NOOP
	return NotImplementedError
}

// VariantClear clears the memory of variant objects.
func VariantClear(v *Variant) error {
	// NOOP
	return NotImplementedError
}

// VariantChangeType changes variant type.
func VariantChangeType(source *Variant, flags uint16, vt VariantType) *Variant {
	// NOOP
	return source
}

// VariantChangeTypeEx changes variant type.
func VariantChangeTypeEx(source *Variant, locale uint32, flags uint16, vt VariantType) *Variant {
	// NOOP
	return source
}

// VariantCopy copies variant.
func VariantCopy(source *Variant) *Variant {
	// NOOP
	return source
}

// VariantCopyIndirect copies variant reference.
func VariantCopyIndirect(source *Variant) *Variant {
	// NOOP
	return source
}
