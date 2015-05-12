// +build cgo

package com

// MaybeError returns error if syscall result is one.
func MaybeError(hr uintptr) (err error) {
	if hr != 0 {
		err = NewError(hr)
	}
	return
}
