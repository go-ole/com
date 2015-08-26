// +build windows,cgo

package variant

import (
	"github.com/go-ole/com/errors"
	"github.com/go-ole/com/types"
)

func VariantInit(v *types.Variant) error {
	return errors.NotImplementedError
}

func VariantClear(v *types.Variant) error {
	return errors.NotImplementedError
}

func VariantChangeType(source *types.Variant, flags uint16, vt types.VariantType) *types.Variant {
	return source
}

func VariantChangeTypeEx(source *types.Variant, locale uint32, flags uint16, vt types.VariantType) *types.Variant {
	return source
}

func VariantCopy(source *types.Variant) *types.Variant {
	return source
}

func VariantCopyIndirect(source *types.Variant) *types.Variant {
	return source
}
