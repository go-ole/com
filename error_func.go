// +build !windows

package com

import "fmt"

func errstr(errno int) string {
	return fmt.Sprintf("error %d (FormatMessage Linux not supported)", errno)
}
