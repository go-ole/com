package com

import syscall "golang.org/x/sys/windows"

// GetInt32FromCall retrieves Int32 from syscall.
func GetInt32FromCall(obj, method uintptr) int32 {
	ret, _, _ := syscall.Syscall(method, 1, obj, 0, 0)
	return int32(ret)
}
