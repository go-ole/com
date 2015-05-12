// +build windows,!cgo

package com

// MaybeError is helper function for returning just the hresult and convert to
// error.
func MaybeError(hr, _ uintptr, _ error) (err error) {
	if hr != 0 {
		err = NewError(hr)
	}
	return
}
