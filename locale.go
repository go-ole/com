package com

import "sync"

// GetUserDefaultLocaleIDFunc implements Kernel32 GetUserDefaultLCID().
type GetUserDefaultLocaleIDFunc func() uint32

// defaultLocale stores default locale callback.
var defaultLocale locale

// locale holds a RW mutex that should prevent multiple threads or goroutines
// from clashing.
//
// It is possible that the design removes the need for this, as this should be
// set within the init function of the library package. However, should also be
// possible to override the library init and go back and forth between
// implementations.
type locale struct {
	mutex    sync.RWMutex
	callback GetUserDefaultLocaleIDFunc
}

// SetGetUserDefaultLocaleIDCallback mutates the callback for retrieving locale ID.
//
// This is to be set by the library that accesses the Windows API.
func SetGetUserDefaultLocaleIDCallback(c GetUserDefaultLocaleIDFunc) {
	defaultLocale.mutex.Lock()
	defer defaultLocale.mutex.Unlock()
	defaultLocale.callback = c
}

// UserDefaultLocaleIDCallback accesses the callback for retrieving locale ID.
func UserDefaultLocaleIDCallback() GetUserDefaultLocaleIDFunc {
	defaultLocale.mutex.RLock()
	defer defaultLocale.mutex.RUnlock()
	return defaultLocale.callback
}

// GetUserDefaultLocaleID accesses the callback to retrieve default user locale.
func GetUserDefaultLocaleID() uint32 {
	defaultLocale.mutex.RLock()
	defer defaultLocale.mutex.RUnlock()
	return defaultLocale.callback()
}
