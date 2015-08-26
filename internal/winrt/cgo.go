// +build windows,cgo

package winrt

import (
	"github.com/go-ole/com/errors"
	"github.com/go-ole/com/types"
	hstr "github.com/go-ole/com/winrt"
	"github.com/go-ole/types/iinspectable"
)

func RoInitialize(threadType uint32) error {
	return errors.NotImplementedError
}

func RoActivateInstance(classID string) (*iinspectable.Inspectable, error) {
	return nil, errors.NotImplementedError
}

func RoGetActivationFactory(classID string, interfaceID *types.GUID) (*iinspectable.Inspectable, error) {
	return nil, errors.NotImplementedError
}

func NewHString(s string) (hstr.HString, error) {
	return types.HString(0), errors.NotImplementedError
}

func DeleteHString(hstring hstr.HString) error {
	return errors.NotImplementedError
}

func HStringToString(hstring hstr.HString) string {
	return ""
}
