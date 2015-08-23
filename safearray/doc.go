/*
Package safearray provides functions and helpers for COM SafeArray.

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

You will be able to build on other platforms other than Windows, but the
functions will not do anything and NOOP. If you want to build for Windows, then
you have to set the GOOS to "windows".
*/
package safearray
