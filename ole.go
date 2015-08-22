package com

// FileTime is the datetime object for files.
//
// It has a strange time format that must be converted.
//
// Reference: https://msdn.microsoft.com/en-us/library/windows/desktop/ms724284(v=vs.85).aspx
// XXX: Complete documentation.
// XXX: Provide helpers for converting to time.Time.
type FileTime struct {
	LowDateTime  int32
	HighDateTime int32
}

// Point is 2D vector type.
type Point struct {
	X int32
	Y int32
}

// Msg is message between processes.
type Msg struct {
	Hwnd    uint32
	Message uint32
	Wparam  int32
	Lparam  int32
	Time    uint32
	Pt      Point
}
