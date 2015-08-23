# SafeArray Bindings

**Unstable! Experimental status**

[![GoDoc](https://godoc.org/github.com/go-ole/safearray?status.svg)](https://godoc.org/github.com/go-ole/safearray)

Provides functions and helpers for Windows COM SafeArray.

There are two bindings, one that loads the Windows shared libraries (DLL) for
calling Windows SafeArray API functions and another that uses cgo to include the
C header files. They are included based on your build for your project. Please
be aware that cgo will increase the executable size of the application.

SafeArray API provides bindings to the Windows API functions for handling
SafeArray structures. There are also helpers for SafeArray that abstract the API
from the developer and allow for a more Go experience.

It is recommended that developers use the helper functions and objects for
SafeArrays. The main com package will return the container structure. The COM
SafeArray structure is somewhat complicated to work with and the helpers exist
to remove the complexity from the developer.

The functions are available for those experienced with the functions. The
functions have been renamed to conform to the Go language naming standards.
Documentation does provide the original name for searches.

The SafeArray helpers only supports a limited set of array types. Currently only
provides helpers for string and byte arrays.

# Building on Linux

You will be able to build on other platforms other than Windows, but the
functions will not do anything and NOOP. If you want to build for Windows, then
you have to set the GOOS to "windows".

# Array Object

The safearray.Array object exists as a wrapper for the COM SafeArray object and
provides some helper receiver functions for converting to a Go slice.

```golang
comSafeArray := someFunctionReturnsCOMSafeArray()
safearray := &safearray.Array{comSafeArray}
```

# Convert to Go Array

Converting to a Go array has two options, you can either append to an existing
Go slice or have a new Go array returned to you.

The example below will append to an existing array. So if `bytes` had existing
data from either another COM SafeArray or from another code location, it will be
retained.

```golang
var bytes []byte

array := &safearray.Array{}
err := array.ToArrayDirect(&bytes)
```

The other option is to return an interface{} slice and convert to the correct
type. The advantage of this, is that it will check the type of the SafeArray and
build the array for you. This only works for supported Variant types and will
not work for custom objects. For user defined objects, you will need to use
`ToArrayDirect()`.

```golang
array := &safearray.Array{}
raw, err := array.ToArray()
bytes, _ := raw.([]byte)
```

There are also a few helpers for use for converting to a few common Go types. So
instead of the the above, you could just use `safearray.ToByteArray()`.

```golang
array := &safearray.Array{}
bytes, _ := safearray.ToByteArray(array)
```

A helper exists for converting to Go strings as well. Support is only limited
to binary string Variant type, but will be extended to support more in the
future.

# Convert to SafeArray

Support for converting a Go array to a COM SafeArray is limited. This will be
extended to support built-in Go types and user-defined types. Use is really
limited to COM servers that require it and very few COM interfaces need it.

**Need examples.**
