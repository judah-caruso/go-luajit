package lua

import (
	"bytes"
	"unsafe"
)

// ToCString creates a null-terminated byte slice from a Go string.
//
// The resulting slice is a new copy.
func ToCString(gostring string) []byte {
	bytes := make([]byte, len(gostring)+1)
	bytes[copy(bytes[:], []byte(gostring))] = 0
	return bytes
}

// ToGoString creates a Go string from a null-terminated byte slice.
//
// The resulting string is a new copy.
func ToGoString(cstr []byte) string {
	if cstr[len(cstr)-1] == 0 {
		cstr = cstr[:len(cstr)-1]
	}

	return string(bytes.Clone(cstr))
}

// ToGoStringPtr creates a Go string from a byte pointer and length.
//
// The resulting string is a new copy.
func ToGoStringPtr(cstr *byte, len int) string {
	return ToGoString(unsafe.Slice(cstr, len))
}
