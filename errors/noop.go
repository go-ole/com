// +build !windows

package errors

import "fmt"

func errstr(errno int) string {
	return fmt.Sprintf("error %d (FormatMessage Linux not supported)", errno)
}

// MaybeError returns error if syscall result is one.
func MaybeError(hr uintptr) error {
	return nil
}
