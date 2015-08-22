// +build 386

package com

// Variant is a COM object that stores arbitrary data.
type Variant struct {
	VariantType VariantType //  2
	wReserved1  uint16      //  4
	wReserved2  uint16      //  6
	wReserved3  uint16      //  8
	Val         int64       // 16
}
