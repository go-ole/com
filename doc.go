/*
Package com binds to the Windows COM API using either cgo or shared libraries.

You really must either link or build for Windows platform to have the library
work. The library will build on other platforms, but the function calls will
NOOP. This is in part so that the functions show up on godoc.org web site and so
that builds will work on other platforms.

How do cross platform compile is left as an exercise to the user. Documentation
is welcome from those who have been able to compile for Windows on other
platforms.

The difference between using the shared libraries and compiling using cgo is
that the shared libraries are loaded dynamically at runtime and won't be linked
with the executable. The shared library also doesn't require cgo, if it isn't
setup for the system. Cgo bindings may increase the executable size and also
requires settings for Go compiler and a C/C++ compiler installed on the system.

There is already documentation on how to setup cgo, so instruction will not be
provided. The most you will need for the shared library building is Go and build
for Windows. To build for Windows, you will need to set the GOOS environment
variable.

The COM package also provides for both 32-bit and 64-bit, so you may also need
to set the GOARCH environment variable, if you are building for another
architecture.


SafeArray API provides both straight bindings to the Windows API functions for
handling SafeArray structures. There are also helpers for SafeArray that
abstract the API from the developer and allow for a more Go experience.

It is recommended that developers use the helper functions and objects for
SafeArrays, given that the library will return the helper structure. SafeArray
structure is somewhat complicated to work with and the helpers exist to remove
the complexity from the developer.
*/
package com
