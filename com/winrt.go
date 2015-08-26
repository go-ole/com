package winrt

import (
	"github.com/go-ole/com/types"

	api "github.com/go-ole/com/internal/winrt"
)

// HString is pointer to an unicode string.
type HString uintptr

// String converts HString to Go String.
func (h HString) String() string {
	return HStringToString(h)
}

// RoInitialize initializes winrt mode.
func RoInitialize(threadType uint32) error {
	return api.RoInitialize(threadType)
}

// RoActivateInstance activates winrt mode.
func RoActivateInstance(classID string) (*iinspectable.Inspectable, error) {
	return api.RoActivateInstance(classID)
}

// RoGetActivationFactory gets activation factory for winrt.
func RoGetActivationFactory(classID string, interfaceID *types.GUID) (*iinspectable.Inspectable, error) {
	return api.RoGetActivationFactory(classID, interfaceID)
}

// NewHString creates HString from Go string
func NewHString(s string) (HString, error) {
	return api.NewHString(s)
}

// DeleteHString removes the memory allocated for HString.
func DeleteHString(hstring HString) error {
	return api.DeleteHString(hstring)
}

// HStringToString converts HString type to Go string.
func HStringToString(hstring HString) string {
	return api.HStringToString(hstring)
}
